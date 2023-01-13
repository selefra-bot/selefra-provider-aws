package mocks

import (
	"github.com/selefra/selefra-provider-aws/constants"
	context "context"
	reflect "reflect"

	firehose "github.com/aws/aws-sdk-go-v2/service/firehose"
	gomock "github.com/golang/mock/gomock"
)

type MockFirehoseClient struct {
	ctrl		*gomock.Controller
	recorder	*MockFirehoseClientMockRecorder
}

type MockFirehoseClientMockRecorder struct {
	mock *MockFirehoseClient
}

func NewMockFirehoseClient(ctrl *gomock.Controller) *MockFirehoseClient {
	mock := &MockFirehoseClient{ctrl: ctrl}
	mock.recorder = &MockFirehoseClientMockRecorder{mock}
	return mock
}

func (m *MockFirehoseClient) EXPECT() *MockFirehoseClientMockRecorder {
	return m.recorder
}

func (m *MockFirehoseClient) DescribeDeliveryStream(arg0 context.Context, arg1 *firehose.DescribeDeliveryStreamInput, arg2 ...func(*firehose.Options)) (*firehose.DescribeDeliveryStreamOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeDeliveryStream, varargs...)
	ret0, _ := ret[0].(*firehose.DescribeDeliveryStreamOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockFirehoseClientMockRecorder) DescribeDeliveryStream(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeDeliveryStream, reflect.TypeOf((*MockFirehoseClient)(nil).DescribeDeliveryStream), varargs...)
}

func (m *MockFirehoseClient) ListDeliveryStreams(arg0 context.Context, arg1 *firehose.ListDeliveryStreamsInput, arg2 ...func(*firehose.Options)) (*firehose.ListDeliveryStreamsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListDeliveryStreams, varargs...)
	ret0, _ := ret[0].(*firehose.ListDeliveryStreamsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockFirehoseClientMockRecorder) ListDeliveryStreams(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListDeliveryStreams, reflect.TypeOf((*MockFirehoseClient)(nil).ListDeliveryStreams), varargs...)
}

func (m *MockFirehoseClient) ListTagsForDeliveryStream(arg0 context.Context, arg1 *firehose.ListTagsForDeliveryStreamInput, arg2 ...func(*firehose.Options)) (*firehose.ListTagsForDeliveryStreamOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTagsForDeliveryStream, varargs...)
	ret0, _ := ret[0].(*firehose.ListTagsForDeliveryStreamOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockFirehoseClientMockRecorder) ListTagsForDeliveryStream(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTagsForDeliveryStream, reflect.TypeOf((*MockFirehoseClient)(nil).ListTagsForDeliveryStream), varargs...)
}
