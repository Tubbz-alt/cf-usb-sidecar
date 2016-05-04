package handlers

import (
	"github.com/go-swagger/go-swagger/httpkit/middleware"

	"github.com/hpcloud/catalog-service-manager/src/csm_manager"

	"github.com/hpcloud/catalog-service-manager/generated/CatalogServiceManager/restapi/operations/workspace"
)

func GetWorkspace(workspaceID string) middleware.Responder {
	internalWorkspaces := csm_manager.GetWorkspace()
	wksp, err := internalWorkspaces.GetWorkspace(workspaceID)
	if err != nil {
		return workspace.NewGetWorkspaceDefault(int(*err.Code)).WithPayload(err)
	}
	return workspace.NewGetWorkspaceOK().WithPayload(wksp)
}
