// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	certificates "kyma-project.io/compass-runtime-agent/internal/certificates"

	config "kyma-project.io/compass-runtime-agent/internal/config"

	connector "kyma-project.io/compass-runtime-agent/internal/compass/connector"

	director "kyma-project.io/compass-runtime-agent/internal/compass/director"

	mock "github.com/stretchr/testify/mock"
)

// ClientsProvider is an autogenerated mock type for the ClientsProvider type
type ClientsProvider struct {
	mock.Mock
}

// GetConnectorCertSecuredClient provides a mock function with given fields: credentials, url
func (_m *ClientsProvider) GetConnectorCertSecuredClient(credentials certificates.ClientCredentials, url string) (connector.Client, error) {
	ret := _m.Called(credentials, url)

	var r0 connector.Client
	if rf, ok := ret.Get(0).(func(certificates.ClientCredentials, string) connector.Client); ok {
		r0 = rf(credentials, url)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(connector.Client)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(certificates.ClientCredentials, string) error); ok {
		r1 = rf(credentials, url)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetConnectorClient provides a mock function with given fields: url
func (_m *ClientsProvider) GetConnectorClient(url string) (connector.Client, error) {
	ret := _m.Called(url)

	var r0 connector.Client
	if rf, ok := ret.Get(0).(func(string) connector.Client); ok {
		r0 = rf(url)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(connector.Client)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(url)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDirectorClient provides a mock function with given fields: credentials, url, runtimeConfig
func (_m *ClientsProvider) GetDirectorClient(credentials certificates.ClientCredentials, url string, runtimeConfig config.RuntimeConfig) (director.DirectorClient, error) {
	ret := _m.Called(credentials, url, runtimeConfig)

	var r0 director.DirectorClient
	if rf, ok := ret.Get(0).(func(certificates.ClientCredentials, string, config.RuntimeConfig) director.DirectorClient); ok {
		r0 = rf(credentials, url, runtimeConfig)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(director.DirectorClient)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(certificates.ClientCredentials, string, config.RuntimeConfig) error); ok {
		r1 = rf(credentials, url, runtimeConfig)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
