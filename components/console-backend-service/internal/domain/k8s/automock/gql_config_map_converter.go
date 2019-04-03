// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import gqlschema "github.com/kyma-project/kyma/components/console-backend-service/internal/gqlschema"

import mock "github.com/stretchr/testify/mock"
import v1 "k8s.io/api/core/v1"

// gqlConfigMapConverter is an autogenerated mock type for the gqlConfigMapConverter type
type gqlConfigMapConverter struct {
	mock.Mock
}

// GQLJSONToConfigMap provides a mock function with given fields: in
func (_m *gqlConfigMapConverter) GQLJSONToConfigMap(in gqlschema.JSON) (v1.ConfigMap, error) {
	ret := _m.Called(in)

	var r0 v1.ConfigMap
	if rf, ok := ret.Get(0).(func(gqlschema.JSON) v1.ConfigMap); ok {
		r0 = rf(in)
	} else {
		r0 = ret.Get(0).(v1.ConfigMap)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(gqlschema.JSON) error); ok {
		r1 = rf(in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ToGQL provides a mock function with given fields: in
func (_m *gqlConfigMapConverter) ToGQL(in *v1.ConfigMap) (*gqlschema.ConfigMap, error) {
	ret := _m.Called(in)

	var r0 *gqlschema.ConfigMap
	if rf, ok := ret.Get(0).(func(*v1.ConfigMap) *gqlschema.ConfigMap); ok {
		r0 = rf(in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gqlschema.ConfigMap)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*v1.ConfigMap) error); ok {
		r1 = rf(in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ToGQLs provides a mock function with given fields: in
func (_m *gqlConfigMapConverter) ToGQLs(in []*v1.ConfigMap) ([]gqlschema.ConfigMap, error) {
	ret := _m.Called(in)

	var r0 []gqlschema.ConfigMap
	if rf, ok := ret.Get(0).(func([]*v1.ConfigMap) []gqlschema.ConfigMap); ok {
		r0 = rf(in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]gqlschema.ConfigMap)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]*v1.ConfigMap) error); ok {
		r1 = rf(in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}