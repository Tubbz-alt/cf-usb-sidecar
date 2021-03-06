// This file was generated by counterfeiter
package provisionerfakes

import (
	"sync"

	"github.com/SUSE/cf-usb-sidecar/csm-extensions/services/dev-mongo/provisioner"
)

type FakeMongoProvisionerInterface struct {
	IsDatabaseCreatedStub        func(string) (bool, error)
	isDatabaseCreatedMutex       sync.RWMutex
	isDatabaseCreatedArgsForCall []struct {
		arg1 string
	}
	isDatabaseCreatedReturns struct {
		result1 bool
		result2 error
	}
	IsUserCreatedStub        func(string, string) (bool, error)
	isUserCreatedMutex       sync.RWMutex
	isUserCreatedArgsForCall []struct {
		arg1 string
		arg2 string
	}
	isUserCreatedReturns struct {
		result1 bool
		result2 error
	}
	CreateDatabaseStub        func(string) error
	createDatabaseMutex       sync.RWMutex
	createDatabaseArgsForCall []struct {
		arg1 string
	}
	createDatabaseReturns struct {
		result1 error
	}
	DeleteDatabaseStub        func(string) error
	deleteDatabaseMutex       sync.RWMutex
	deleteDatabaseArgsForCall []struct {
		arg1 string
	}
	deleteDatabaseReturns struct {
		result1 error
	}
	CreateUserStub        func(string, string, string) error
	createUserMutex       sync.RWMutex
	createUserArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
	}
	createUserReturns struct {
		result1 error
	}
	DeleteUserStub        func(string, string) error
	deleteUserMutex       sync.RWMutex
	deleteUserArgsForCall []struct {
		arg1 string
		arg2 string
	}
	deleteUserReturns struct {
		result1 error
	}
}

func (fake *FakeMongoProvisionerInterface) IsDatabaseCreated(arg1 string) (bool, error) {
	fake.isDatabaseCreatedMutex.Lock()
	fake.isDatabaseCreatedArgsForCall = append(fake.isDatabaseCreatedArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.isDatabaseCreatedMutex.Unlock()
	if fake.IsDatabaseCreatedStub != nil {
		return fake.IsDatabaseCreatedStub(arg1)
	} else {
		return fake.isDatabaseCreatedReturns.result1, fake.isDatabaseCreatedReturns.result2
	}
}

func (fake *FakeMongoProvisionerInterface) IsDatabaseCreatedCallCount() int {
	fake.isDatabaseCreatedMutex.RLock()
	defer fake.isDatabaseCreatedMutex.RUnlock()
	return len(fake.isDatabaseCreatedArgsForCall)
}

func (fake *FakeMongoProvisionerInterface) IsDatabaseCreatedArgsForCall(i int) string {
	fake.isDatabaseCreatedMutex.RLock()
	defer fake.isDatabaseCreatedMutex.RUnlock()
	return fake.isDatabaseCreatedArgsForCall[i].arg1
}

func (fake *FakeMongoProvisionerInterface) IsDatabaseCreatedReturns(result1 bool, result2 error) {
	fake.IsDatabaseCreatedStub = nil
	fake.isDatabaseCreatedReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeMongoProvisionerInterface) IsUserCreated(arg1 string, arg2 string) (bool, error) {
	fake.isUserCreatedMutex.Lock()
	fake.isUserCreatedArgsForCall = append(fake.isUserCreatedArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.isUserCreatedMutex.Unlock()
	if fake.IsUserCreatedStub != nil {
		return fake.IsUserCreatedStub(arg1, arg2)
	} else {
		return fake.isUserCreatedReturns.result1, fake.isUserCreatedReturns.result2
	}
}

func (fake *FakeMongoProvisionerInterface) IsUserCreatedCallCount() int {
	fake.isUserCreatedMutex.RLock()
	defer fake.isUserCreatedMutex.RUnlock()
	return len(fake.isUserCreatedArgsForCall)
}

func (fake *FakeMongoProvisionerInterface) IsUserCreatedArgsForCall(i int) (string, string) {
	fake.isUserCreatedMutex.RLock()
	defer fake.isUserCreatedMutex.RUnlock()
	return fake.isUserCreatedArgsForCall[i].arg1, fake.isUserCreatedArgsForCall[i].arg2
}

func (fake *FakeMongoProvisionerInterface) IsUserCreatedReturns(result1 bool, result2 error) {
	fake.IsUserCreatedStub = nil
	fake.isUserCreatedReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeMongoProvisionerInterface) CreateDatabase(arg1 string) error {
	fake.createDatabaseMutex.Lock()
	fake.createDatabaseArgsForCall = append(fake.createDatabaseArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.createDatabaseMutex.Unlock()
	if fake.CreateDatabaseStub != nil {
		return fake.CreateDatabaseStub(arg1)
	} else {
		return fake.createDatabaseReturns.result1
	}
}

func (fake *FakeMongoProvisionerInterface) CreateDatabaseCallCount() int {
	fake.createDatabaseMutex.RLock()
	defer fake.createDatabaseMutex.RUnlock()
	return len(fake.createDatabaseArgsForCall)
}

func (fake *FakeMongoProvisionerInterface) CreateDatabaseArgsForCall(i int) string {
	fake.createDatabaseMutex.RLock()
	defer fake.createDatabaseMutex.RUnlock()
	return fake.createDatabaseArgsForCall[i].arg1
}

func (fake *FakeMongoProvisionerInterface) CreateDatabaseReturns(result1 error) {
	fake.CreateDatabaseStub = nil
	fake.createDatabaseReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeMongoProvisionerInterface) DeleteDatabase(arg1 string) error {
	fake.deleteDatabaseMutex.Lock()
	fake.deleteDatabaseArgsForCall = append(fake.deleteDatabaseArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.deleteDatabaseMutex.Unlock()
	if fake.DeleteDatabaseStub != nil {
		return fake.DeleteDatabaseStub(arg1)
	} else {
		return fake.deleteDatabaseReturns.result1
	}
}

func (fake *FakeMongoProvisionerInterface) DeleteDatabaseCallCount() int {
	fake.deleteDatabaseMutex.RLock()
	defer fake.deleteDatabaseMutex.RUnlock()
	return len(fake.deleteDatabaseArgsForCall)
}

func (fake *FakeMongoProvisionerInterface) DeleteDatabaseArgsForCall(i int) string {
	fake.deleteDatabaseMutex.RLock()
	defer fake.deleteDatabaseMutex.RUnlock()
	return fake.deleteDatabaseArgsForCall[i].arg1
}

func (fake *FakeMongoProvisionerInterface) DeleteDatabaseReturns(result1 error) {
	fake.DeleteDatabaseStub = nil
	fake.deleteDatabaseReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeMongoProvisionerInterface) CreateUser(arg1 string, arg2 string, arg3 string) error {
	fake.createUserMutex.Lock()
	fake.createUserArgsForCall = append(fake.createUserArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	fake.createUserMutex.Unlock()
	if fake.CreateUserStub != nil {
		return fake.CreateUserStub(arg1, arg2, arg3)
	} else {
		return fake.createUserReturns.result1
	}
}

func (fake *FakeMongoProvisionerInterface) CreateUserCallCount() int {
	fake.createUserMutex.RLock()
	defer fake.createUserMutex.RUnlock()
	return len(fake.createUserArgsForCall)
}

func (fake *FakeMongoProvisionerInterface) CreateUserArgsForCall(i int) (string, string, string) {
	fake.createUserMutex.RLock()
	defer fake.createUserMutex.RUnlock()
	return fake.createUserArgsForCall[i].arg1, fake.createUserArgsForCall[i].arg2, fake.createUserArgsForCall[i].arg3
}

func (fake *FakeMongoProvisionerInterface) CreateUserReturns(result1 error) {
	fake.CreateUserStub = nil
	fake.createUserReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeMongoProvisionerInterface) DeleteUser(arg1 string, arg2 string) error {
	fake.deleteUserMutex.Lock()
	fake.deleteUserArgsForCall = append(fake.deleteUserArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.deleteUserMutex.Unlock()
	if fake.DeleteUserStub != nil {
		return fake.DeleteUserStub(arg1, arg2)
	} else {
		return fake.deleteUserReturns.result1
	}
}

func (fake *FakeMongoProvisionerInterface) DeleteUserCallCount() int {
	fake.deleteUserMutex.RLock()
	defer fake.deleteUserMutex.RUnlock()
	return len(fake.deleteUserArgsForCall)
}

func (fake *FakeMongoProvisionerInterface) DeleteUserArgsForCall(i int) (string, string) {
	fake.deleteUserMutex.RLock()
	defer fake.deleteUserMutex.RUnlock()
	return fake.deleteUserArgsForCall[i].arg1, fake.deleteUserArgsForCall[i].arg2
}

func (fake *FakeMongoProvisionerInterface) DeleteUserReturns(result1 error) {
	fake.DeleteUserStub = nil
	fake.deleteUserReturns = struct {
		result1 error
	}{result1}
}

var _ provisioner.MongoProvisionerInterface = new(FakeMongoProvisionerInterface)
