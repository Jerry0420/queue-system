// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	grpcServices "github.com/jerry0420/queue-system/backend/repository/grpcServices"
	mock "github.com/stretchr/testify/mock"
)

// GrpcServiceClient is an autogenerated mock type for the GrpcServiceClient type
type GrpcServiceClient struct {
	mock.Mock
}

// GenerateCSV provides a mock function with given fields: ctx, in, opts
func (_m *GrpcServiceClient) GenerateCSV(ctx context.Context, in *grpcServices.GenerateCSVRequest, opts ...grpc.CallOption) (*grpcServices.GenerateCSVResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *grpcServices.GenerateCSVResponse
	if rf, ok := ret.Get(0).(func(context.Context, *grpcServices.GenerateCSVRequest, ...grpc.CallOption) *grpcServices.GenerateCSVResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*grpcServices.GenerateCSVResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *grpcServices.GenerateCSVRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendEmail provides a mock function with given fields: ctx, in, opts
func (_m *GrpcServiceClient) SendEmail(ctx context.Context, in *grpcServices.SendEmailRequest, opts ...grpc.CallOption) (*grpcServices.SendEmailResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *grpcServices.SendEmailResponse
	if rf, ok := ret.Get(0).(func(context.Context, *grpcServices.SendEmailRequest, ...grpc.CallOption) *grpcServices.SendEmailResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*grpcServices.SendEmailResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *grpcServices.SendEmailRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
