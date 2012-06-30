// Automatically generated by MockGen. DO NOT EDIT!
// Source: code.google.com/p/gomock/sample (interfaces: Index,Embed,Embedded)

package mock_sample

import (
	bufio "bufio"
	bytes "bytes"
	gomock "code.google.com/p/gomock/gomock"
	imp1 "code.google.com/p/gomock/sample/imp1"
	imp2 "code.google.com/p/gomock/sample/imp2"
	imp3 "code.google.com/p/gomock/sample/imp3"
	imp4 "code.google.com/p/gomock/sample/imp4"
	hash "hash"
	io "io"
	http "net/http"
)

// Mock of Index interface
type MockIndex struct {
	ctrl     *gomock.Controller
	recorder *_MockIndexRecorder
}

// Recorder for MockIndex (not exported)
type _MockIndexRecorder struct {
	mock *MockIndex
}

func NewMockIndex(ctrl *gomock.Controller) *MockIndex {
	mock := &MockIndex{ctrl: ctrl}
	mock.recorder = &_MockIndexRecorder{mock}
	return mock
}

func (_m *MockIndex) EXPECT() *_MockIndexRecorder {
	return _m.recorder
}

func (_m *MockIndex) Anon(_param0 string) {
	_m.ctrl.Call(_m, "Anon", _param0)
}

func (_mr *_MockIndexRecorder) Anon(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Anon", arg0)
}

func (_m *MockIndex) Chan(_param0 chan int, _param1 chan<- hash.Hash) {
	_m.ctrl.Call(_m, "Chan", _param0, _param1)
}

func (_mr *_MockIndexRecorder) Chan(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Chan", arg0, arg1)
}

func (_m *MockIndex) ConcreteRet() chan<- bool {
	ret := _m.ctrl.Call(_m, "ConcreteRet")
	ret0, _ := ret[0].(chan<- bool)
	return ret0
}

func (_mr *_MockIndexRecorder) ConcreteRet() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ConcreteRet")
}

func (_m *MockIndex) Ellip(_param0 string, _param1 ...interface{}) {
	_s := []interface{}{_param0}
	for _, _x := range _param1 {
		_s = append(_s, _x)
	}
	_m.ctrl.Call(_m, "Ellip", _s...)
}

func (_mr *_MockIndexRecorder) Ellip(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0}, arg1...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Ellip", _s...)
}

func (_m *MockIndex) EllipOnly(_param0 ...string) {
	_s := []interface{}{}
	for _, _x := range _param0 {
		_s = append(_s, _x)
	}
	_m.ctrl.Call(_m, "EllipOnly", _s...)
}

func (_mr *_MockIndexRecorder) EllipOnly(arg0 ...interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "EllipOnly", arg0...)
}

func (_m *MockIndex) ForeignFour(_param0 imp4.Imp4) {
	_m.ctrl.Call(_m, "ForeignFour", _param0)
}

func (_mr *_MockIndexRecorder) ForeignFour(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ForeignFour", arg0)
}

func (_m *MockIndex) ForeignOne(_param0 imp1.Imp1) {
	_m.ctrl.Call(_m, "ForeignOne", _param0)
}

func (_mr *_MockIndexRecorder) ForeignOne(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ForeignOne", arg0)
}

func (_m *MockIndex) ForeignThree(_param0 imp3.Imp3) {
	_m.ctrl.Call(_m, "ForeignThree", _param0)
}

func (_mr *_MockIndexRecorder) ForeignThree(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ForeignThree", arg0)
}

func (_m *MockIndex) ForeignTwo(_param0 imp2.Imp2) {
	_m.ctrl.Call(_m, "ForeignTwo", _param0)
}

func (_mr *_MockIndexRecorder) ForeignTwo(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ForeignTwo", arg0)
}

func (_m *MockIndex) Func(_param0 func(http.Request) (int, bool)) {
	_m.ctrl.Call(_m, "Func", _param0)
}

func (_mr *_MockIndexRecorder) Func(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Func", arg0)
}

func (_m *MockIndex) Get(_param0 string) interface{} {
	ret := _m.ctrl.Call(_m, "Get", _param0)
	ret0, _ := ret[0].(interface{})
	return ret0
}

func (_mr *_MockIndexRecorder) Get(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Get", arg0)
}

func (_m *MockIndex) GetTwo(_param0 string, _param1 string) (interface{}, interface{}) {
	ret := _m.ctrl.Call(_m, "GetTwo", _param0, _param1)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(interface{})
	return ret0, ret1
}

func (_mr *_MockIndexRecorder) GetTwo(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetTwo", arg0, arg1)
}

func (_m *MockIndex) Map(_param0 map[int]hash.Hash) {
	_m.ctrl.Call(_m, "Map", _param0)
}

func (_mr *_MockIndexRecorder) Map(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Map", arg0)
}

func (_m *MockIndex) NillableRet() error {
	ret := _m.ctrl.Call(_m, "NillableRet")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockIndexRecorder) NillableRet() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NillableRet")
}

func (_m *MockIndex) Other() hash.Hash {
	ret := _m.ctrl.Call(_m, "Other")
	ret0, _ := ret[0].(hash.Hash)
	return ret0
}

func (_mr *_MockIndexRecorder) Other() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Other")
}

func (_m *MockIndex) Ptr(_param0 *int) {
	_m.ctrl.Call(_m, "Ptr", _param0)
}

func (_mr *_MockIndexRecorder) Ptr(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Ptr", arg0)
}

func (_m *MockIndex) Put(_param0 string, _param1 interface{}) {
	_m.ctrl.Call(_m, "Put", _param0, _param1)
}

func (_mr *_MockIndexRecorder) Put(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Put", arg0, arg1)
}

func (_m *MockIndex) Slice(_param0 []int) [3]int {
	ret := _m.ctrl.Call(_m, "Slice", _param0)
	ret0, _ := ret[0].([3]int)
	return ret0
}

func (_mr *_MockIndexRecorder) Slice(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Slice", arg0)
}

func (_m *MockIndex) Summary(_param0 *bytes.Buffer, _param1 io.Writer) {
	_m.ctrl.Call(_m, "Summary", _param0, _param1)
}

func (_mr *_MockIndexRecorder) Summary(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Summary", arg0, arg1)
}

// Mock of Embed interface
type MockEmbed struct {
	ctrl     *gomock.Controller
	recorder *_MockEmbedRecorder
}

// Recorder for MockEmbed (not exported)
type _MockEmbedRecorder struct {
	mock *MockEmbed
}

func NewMockEmbed(ctrl *gomock.Controller) *MockEmbed {
	mock := &MockEmbed{ctrl: ctrl}
	mock.recorder = &_MockEmbedRecorder{mock}
	return mock
}

func (_m *MockEmbed) EXPECT() *_MockEmbedRecorder {
	return _m.recorder
}

func (_m *MockEmbed) EmbeddedMethod() {
	_m.ctrl.Call(_m, "EmbeddedMethod")
}

func (_mr *_MockEmbedRecorder) EmbeddedMethod() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "EmbeddedMethod")
}

func (_m *MockEmbed) ForeignEmbeddedMethod() *bufio.Reader {
	ret := _m.ctrl.Call(_m, "ForeignEmbeddedMethod")
	ret0, _ := ret[0].(*bufio.Reader)
	return ret0
}

func (_mr *_MockEmbedRecorder) ForeignEmbeddedMethod() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ForeignEmbeddedMethod")
}

func (_m *MockEmbed) ImplicitPackage(_param0 string, _param1 imp1.ImpT, _param2 []imp1.ImpT, _param3 *imp1.ImpT, _param4 chan imp1.ImpT) {
	_m.ctrl.Call(_m, "ImplicitPackage", _param0, _param1, _param2, _param3, _param4)
}

func (_mr *_MockEmbedRecorder) ImplicitPackage(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ImplicitPackage", arg0, arg1, arg2, arg3, arg4)
}

func (_m *MockEmbed) RegularMethod() {
	_m.ctrl.Call(_m, "RegularMethod")
}

func (_mr *_MockEmbedRecorder) RegularMethod() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RegularMethod")
}

// Mock of Embedded interface
type MockEmbedded struct {
	ctrl     *gomock.Controller
	recorder *_MockEmbeddedRecorder
}

// Recorder for MockEmbedded (not exported)
type _MockEmbeddedRecorder struct {
	mock *MockEmbedded
}

func NewMockEmbedded(ctrl *gomock.Controller) *MockEmbedded {
	mock := &MockEmbedded{ctrl: ctrl}
	mock.recorder = &_MockEmbeddedRecorder{mock}
	return mock
}

func (_m *MockEmbedded) EXPECT() *_MockEmbeddedRecorder {
	return _m.recorder
}

func (_m *MockEmbedded) EmbeddedMethod() {
	_m.ctrl.Call(_m, "EmbeddedMethod")
}

func (_mr *_MockEmbeddedRecorder) EmbeddedMethod() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "EmbeddedMethod")
}
