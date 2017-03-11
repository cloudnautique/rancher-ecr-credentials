package mocks

import client "github.com/rancher/go-rancher/client"
import mock "github.com/stretchr/testify/mock"

// RegistryCredentialOperations is an autogenerated mock type for the RegistryCredentialOperations type
type RegistryCredentialOperations struct {
	mock.Mock
}

// ActionActivate provides a mock function with given fields: _a0
func (_m *RegistryCredentialOperations) ActionActivate(_a0 *client.RegistryCredential) (*client.Credential, error) {
	ret := _m.Called(_a0)

	var r0 *client.Credential
	if rf, ok := ret.Get(0).(func(*client.RegistryCredential) *client.Credential); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.Credential)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*client.RegistryCredential) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ActionCreate provides a mock function with given fields: _a0
func (_m *RegistryCredentialOperations) ActionCreate(_a0 *client.RegistryCredential) (*client.Credential, error) {
	ret := _m.Called(_a0)

	var r0 *client.Credential
	if rf, ok := ret.Get(0).(func(*client.RegistryCredential) *client.Credential); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.Credential)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*client.RegistryCredential) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ActionDeactivate provides a mock function with given fields: _a0
func (_m *RegistryCredentialOperations) ActionDeactivate(_a0 *client.RegistryCredential) (*client.Credential, error) {
	ret := _m.Called(_a0)

	var r0 *client.Credential
	if rf, ok := ret.Get(0).(func(*client.RegistryCredential) *client.Credential); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.Credential)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*client.RegistryCredential) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ActionPurge provides a mock function with given fields: _a0
func (_m *RegistryCredentialOperations) ActionPurge(_a0 *client.RegistryCredential) (*client.Credential, error) {
	ret := _m.Called(_a0)

	var r0 *client.Credential
	if rf, ok := ret.Get(0).(func(*client.RegistryCredential) *client.Credential); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.Credential)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*client.RegistryCredential) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ActionRemove provides a mock function with given fields: _a0
func (_m *RegistryCredentialOperations) ActionRemove(_a0 *client.RegistryCredential) (*client.Credential, error) {
	ret := _m.Called(_a0)

	var r0 *client.Credential
	if rf, ok := ret.Get(0).(func(*client.RegistryCredential) *client.Credential); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.Credential)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*client.RegistryCredential) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ActionUpdate provides a mock function with given fields: _a0
func (_m *RegistryCredentialOperations) ActionUpdate(_a0 *client.RegistryCredential) (*client.Credential, error) {
	ret := _m.Called(_a0)

	var r0 *client.Credential
	if rf, ok := ret.Get(0).(func(*client.RegistryCredential) *client.Credential); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.Credential)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*client.RegistryCredential) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ById provides a mock function with given fields: id
func (_m *RegistryCredentialOperations) ById(id string) (*client.RegistryCredential, error) {
	ret := _m.Called(id)

	var r0 *client.RegistryCredential
	if rf, ok := ret.Get(0).(func(string) *client.RegistryCredential); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.RegistryCredential)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: opts
func (_m *RegistryCredentialOperations) Create(opts *client.RegistryCredential) (*client.RegistryCredential, error) {
	ret := _m.Called(opts)

	var r0 *client.RegistryCredential
	if rf, ok := ret.Get(0).(func(*client.RegistryCredential) *client.RegistryCredential); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.RegistryCredential)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*client.RegistryCredential) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: container
func (_m *RegistryCredentialOperations) Delete(container *client.RegistryCredential) error {
	ret := _m.Called(container)

	var r0 error
	if rf, ok := ret.Get(0).(func(*client.RegistryCredential) error); ok {
		r0 = rf(container)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// List provides a mock function with given fields: opts
func (_m *RegistryCredentialOperations) List(opts *client.ListOpts) (*client.RegistryCredentialCollection, error) {
	ret := _m.Called(opts)

	var r0 *client.RegistryCredentialCollection
	if rf, ok := ret.Get(0).(func(*client.ListOpts) *client.RegistryCredentialCollection); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.RegistryCredentialCollection)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*client.ListOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: existing, updates
func (_m *RegistryCredentialOperations) Update(existing *client.RegistryCredential, updates interface{}) (*client.RegistryCredential, error) {
	ret := _m.Called(existing, updates)

	var r0 *client.RegistryCredential
	if rf, ok := ret.Get(0).(func(*client.RegistryCredential, interface{}) *client.RegistryCredential); ok {
		r0 = rf(existing, updates)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.RegistryCredential)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*client.RegistryCredential, interface{}) error); ok {
		r1 = rf(existing, updates)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
