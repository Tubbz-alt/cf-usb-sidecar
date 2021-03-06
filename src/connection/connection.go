package connection

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/SUSE/cf-usb-sidecar/generated/CatalogServiceManager/models"
	"github.com/SUSE/cf-usb-sidecar/src/common"
	"github.com/SUSE/cf-usb-sidecar/src/common/utils"
	"github.com/Sirupsen/logrus"
)

// CSMConnection object for managing the connection.
type CSMConnection struct {
	Logger     *logrus.Logger
	Config     *common.ServiceManagerConfiguration
	FileHelper utils.CSMFileHelperInterface
}

// NewCSMConnection creates CSMConnection
func NewCSMConnection(logger *logrus.Logger, config *common.ServiceManagerConfiguration, fileHelper utils.CSMFileHelperInterface) *CSMConnection {
	return &CSMConnection{Logger: logger, Config: config, FileHelper: fileHelper}
}

func (c *CSMConnection) getConnectionsGetExtension(homePath string) (bool, string) {
	return c.FileHelper.GetExtension(filepath.Join(homePath, "connection", "get"))
}

func (c *CSMConnection) getConnectionsCreateExtension(homePath string) (bool, string) {
	return c.FileHelper.GetExtension(filepath.Join(homePath, "connection", "create"))
}

func (c *CSMConnection) getConnectionsDeleteExtension(homePath string) (bool, string) {
	return c.FileHelper.GetExtension(filepath.Join(homePath, "connection", "delete"))
}

//create ServiceManagerConnectionResponse from the json we received in file
func marshalResponseFromMessage(message []byte) (*models.ServiceManagerConnectionResponse, *models.Error, error) {
	connection := utils.NewConnection()
	jsonresp := utils.JsonResponse{}
	if len(message) == 0 {
		return nil, nil, errors.New("Empty response")
	}
	err := jsonresp.Unmarshal(message)

	if err != nil {
		return nil, nil, err
	}
	if strings.ToLower(jsonresp.Status) != "successful" { //the extension is giving us an error responses
		var code int64
		var message string
		if jsonresp.ErrorCode == 0 {
			code = utils.HTTP_500
		} else {
			code = int64(jsonresp.ErrorCode)
		}

		message = jsonresp.ErrorMessage

		return nil, utils.GenerateErrorResponse(&code, message), nil

	}

	connection.Details = make(map[string]interface{})
	switch t := jsonresp.Details.(type) {
	default:
		connection.Details["data"] = t
	case map[string]interface{}:
		connection.Details = jsonresp.Details.(map[string]interface{})
	}
	//connection.Details = jsonresp.Details.(map[string]interface{})
	connection.Status = &common.PROCESSING_STATUS_SUCCESSFUL
	connection.ProcessingType = &common.PROCESSING_TYPE_EXTENSION

	return &connection, nil, nil
}

func checkParamsOk(workspaceID string, connectionID string, extensionPath string) error {
	if workspaceID == "" {
		err := errors.New("workspaceID is not set")
		return err
	}
	if connectionID == "" {
		err := errors.New("connectionID is not set")
		return err
	}
	if extensionPath == "" {
		err := errors.New("extensionPath is not set")
		return err
	}
	return nil
}

func (c *CSMConnection) executeExtension(workspaceID string, connectionID string, details map[string]interface{}, extensionPath string) (*models.ServiceManagerConnectionResponse, *models.Error, error) {
	if err := checkParamsOk(workspaceID, connectionID, extensionPath); err != nil {
		return nil, nil, err
	}

	detailsStr := ""

	if details != nil {
		detailsJSON, err := json.Marshal(details)
		if err != nil {
			return nil, nil, err
		}

		detailsStr = string(detailsJSON)
	}

	c.Logger.WithFields(logrus.Fields{"workspaceID": workspaceID, "connectionID": connectionID, "extension Path": extensionPath, "details": details}).Info("executeExtension")

	if success, outputFile, output := c.FileHelper.RunExtensionFileGen(extensionPath, workspaceID, connectionID, detailsStr); success {
		c.Logger.WithFields(logrus.Fields{"extension execution status": success}).Info("executeExtension")
		c.Logger.WithFields(logrus.Fields{"extension execution Result": output}).Debug("executeExtension")

		fileContent, err := utils.ReadOutputFile(outputFile, *c.Config.DEV_MODE != "true")
		if err != nil {
			return nil, nil, errors.New(fmt.Sprintf("Error reading response from extension: %s", err.Error()))
		}
		return marshalResponseFromMessage(fileContent)
	} else {
		// extension couldn't be executed, returned an error or timedout
		//first we check for timeout (success=false,  output==nil)
		if output == "" {
			return nil, utils.GenerateErrorResponse(&utils.HTTP_408, utils.ERR_TIMEOUT), nil
		}
		//else it means that the extension did not return a zero code	 ("success = false, output != nil)
		err := errors.New(output)
		return nil, nil, err
	}
}

// CheckExtensions checks for workspace extensions
func (c *CSMConnection) CheckExtensions() {

	_, file := c.getConnectionsGetExtension(*c.Config.MANAGER_HOME)
	c.Logger.WithFields(logrus.Fields{"Connections Get extension": file}).Info("CheckExtensions")

	_, file = c.getConnectionsCreateExtension(*c.Config.MANAGER_HOME)
	c.Logger.WithFields(logrus.Fields{"Connections Create extension": file}).Info("CheckExtensions")

	_, file = c.getConnectionsDeleteExtension(*c.Config.MANAGER_HOME)
	c.Logger.WithFields(logrus.Fields{"Connections Delete extension": file}).Info("CheckExtensions")
}

func (c *CSMConnection) executeRequest(workspaceID string, connectionID string, details map[string]interface{}, requestType string, filename string) (*models.ServiceManagerConnectionResponse, *models.Error) {
	var modelserr *models.Error
	var connection *models.ServiceManagerConnectionResponse
	var err error

	connection, modelserr, err = c.executeExtension(workspaceID, connectionID, details, filename)
	if err != nil {
		c.Logger.Error(requestType, err)
		modelserr = utils.GenerateErrorResponse(&utils.HTTP_500, err.Error())
	}

	if connection != nil {
		if connection.Details == nil {
			connection.Details = details
		} else {
			//the "data" item is added if the response of the extension is a string or nil
			if _, ok := connection.Details["data"]; ok {
				//if the response is nil, set the details
				if connection.Details["data"] == nil {
					connection.Details = details
				}
			}
		}
	}

	return connection, modelserr
}

func generateNoopResponse() *models.ServiceManagerConnectionResponse {
	resp := models.ServiceManagerConnectionResponse{
		ProcessingType: &common.PROCESSING_TYPE_NONE,
		Status:         &common.PROCESSING_STATUS_NONE,
	}
	return &resp
}

// GetConnection get connections
func (c *CSMConnection) GetConnection(workspaceID string, connectionID string) (*models.ServiceManagerConnectionResponse, *models.Error) {
	c.Logger.WithFields(logrus.Fields{"workspaceID": workspaceID, "connectionID": connectionID}).Info("GetConnection")

	serviceManagerConfig := common.NewServiceManagerConfiguration()
	exists, filename := c.getConnectionsGetExtension(*serviceManagerConfig.MANAGER_HOME)
	if !exists || filename == "" {
		c.Logger.WithFields(logrus.Fields{utils.ERR_EXTENSION_NOT_FOUND: exists}).Info("GetConnection")
		return generateNoopResponse(), nil
	}
	return c.executeRequest(workspaceID, connectionID, make(map[string]interface{}), "GetConnection", filename)
}

// CreateConnection create connections
func (c *CSMConnection) CreateConnection(workspaceID string, connectionID string, details map[string]interface{}) (*models.ServiceManagerConnectionResponse, *models.Error) {
	c.Logger.WithFields(logrus.Fields{"workspaceID": workspaceID, "connectionID": connectionID, "details": details}).Info("CreateConnection")
	serviceManagerConfig := common.NewServiceManagerConfiguration()
	exists, filename := c.getConnectionsCreateExtension(*serviceManagerConfig.MANAGER_HOME)
	if (!exists) && (serviceManagerConfig.PARAMETERS != nil) {
		connection := utils.NewConnection()
		c.Logger.WithFields(logrus.Fields{"Extension not found": exists}).Info("GetConnection")
		parametersNameList := strings.Split(*serviceManagerConfig.PARAMETERS, " ")
		c.Logger.WithFields(logrus.Fields{"Parameter List": parametersNameList}).Info("GetConnection")
		connection.Details = make(map[string]interface{})
		for _, parameterName := range parametersNameList {
			parameterValue, ok := os.LookupEnv(parameterName)
			if ok {
				connection.Details[parameterName] = parameterValue
			}
		}
		connection.ProcessingType = &common.PROCESSING_TYPE_DEFAULT
		connection.Status = &common.PROCESSING_STATUS_SUCCESSFUL
		return &connection, nil

	} else if !exists || filename == "" {
		c.Logger.WithFields(logrus.Fields{utils.ERR_EXTENSION_NOT_FOUND: exists}).Info("CreateConnection")
		return generateNoopResponse(), nil
	}
	return c.executeRequest(workspaceID, connectionID, details, "CreateConnection", filename)
}

// DeleteConnection delete connections
func (c *CSMConnection) DeleteConnection(workspaceID string, connectionID string) (*models.ServiceManagerConnectionResponse, *models.Error) {
	c.Logger.WithFields(logrus.Fields{"workspaceID": workspaceID, "connectionID": connectionID}).Info("DeleteConnection")

	serviceManagerConfig := common.NewServiceManagerConfiguration()
	exists, filename := c.getConnectionsDeleteExtension(*serviceManagerConfig.MANAGER_HOME)
	if !exists || filename == "" {
		c.Logger.WithFields(logrus.Fields{utils.ERR_EXTENSION_NOT_FOUND: exists}).Info("DeleteConnection")
		return generateNoopResponse(), nil
	}
	return c.executeRequest(workspaceID, connectionID, make(map[string]interface{}), "DeleteConnection", filename)
}
