package foundation

import (
	SGroupsNamespace "HBF-tests/pkg/api/sgroups/v1"
	"HBF-tests/tests/common/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

// ToolsForTest Вспомогательная структура для тестов
type ToolsForTest struct {
	NamespaceAPI      SGroupsNamespace.SGroupsNamespaceAPIClient
	AddressGroupAPI   SGroupsNamespace.SGroupsAddressGroupsAPIClient
	HostBindingAPI    SGroupsNamespace.SGroupsHostBindingAPIClient
	NetworkBindingAPI SGroupsNamespace.SGroupsNetworkBindingAPIClient
	ServiceBindingAPI SGroupsNamespace.SGroupsServiceBindingAPIClient
	Trailer           metadata.MD
}

// TestCaseBody Структура тела запроса для List
type TestCaseBody struct {
	TestName string
	Req      SGroupsNamespace.NamespaceReq_List
}

// TestCaseBodyUpsert Структура тела запроса для Upsert
type TestCaseBodyUpsert struct {
	TestName string
	Req      SGroupsNamespace.NamespaceReq_Upsert
}

// TestCaseBodyDelete Структура тела запроса для Delete
type TestCaseBodyDelete struct {
	TestName string
	Req      SGroupsNamespace.NamespaceReq_Delete
}

// sendReqDwv Отправка запроса
func (st *ToolsForTest) sendReqSGroupsNamespace(req *SGroupsNamespace.NamespaceReq_List) (*SGroupsNamespace.NamespaceResp_List, error) {
	ctx, cancel := utils.LocalCtx()
	defer cancel()
	return st.NamespaceAPI.List(ctx, req, grpc.Trailer(&st.Trailer))
}

// FoundationTests Основа теста для List
func (st *ToolsForTest) FoundationTests(params *TestCaseBody) (*SGroupsNamespace.NamespaceResp_List, error) {
	resp, err := st.sendReqSGroupsNamespace(&params.Req)
	if err != nil {
		return nil, err
	}
	return resp, nil

}

// FoundationTestsUpsert Основа теста для Upsert
func (st *ToolsForTest) FoundationTestsUpsert(params *TestCaseBodyUpsert) (*SGroupsNamespace.NamespaceResp_Upsert, error) {
	ctx, cancel := utils.LocalCtx()
	defer cancel()
	return st.NamespaceAPI.Upsert(ctx, &params.Req, grpc.Trailer(&st.Trailer))
}

// FoundationTestsDelete Основа теста для Delete
func (st *ToolsForTest) FoundationTestsDelete(params *TestCaseBodyDelete) (*emptypb.Empty, error) {
	ctx, cancel := utils.LocalCtx()
	defer cancel()
	return st.NamespaceAPI.Delete(ctx, &params.Req, grpc.Trailer(&st.Trailer))
}

// TestCaseBodyAGList Структура тела запроса для AddressGroup List
type TestCaseBodyAGList struct {
	TestName string
	Req      SGroupsNamespace.AddressGroupReq_List
}

// TestCaseBodyAGUpsert Структура тела запроса для AddressGroup Upsert
type TestCaseBodyAGUpsert struct {
	TestName string
	Req      SGroupsNamespace.AddressGroupReq_Upsert
}

// TestCaseBodyAGDelete Структура тела запроса для AddressGroup Delete
type TestCaseBodyAGDelete struct {
	TestName string
	Req      SGroupsNamespace.AddressGroupReq_Delete
}

// FoundationTestsAGList Основа теста для AddressGroup List
func (st *ToolsForTest) FoundationTestsAGList(params *TestCaseBodyAGList) (*SGroupsNamespace.AddressGroupResp_List, error) {
	ctx, cancel := utils.LocalCtx()
	defer cancel()
	return st.AddressGroupAPI.List(ctx, &params.Req, grpc.Trailer(&st.Trailer))
}

// FoundationTestsAGUpsert Основа теста для AddressGroup Upsert
func (st *ToolsForTest) FoundationTestsAGUpsert(params *TestCaseBodyAGUpsert) (*SGroupsNamespace.AddressGroupResp_Upsert, error) {
	ctx, cancel := utils.LocalCtx()
	defer cancel()
	return st.AddressGroupAPI.Upsert(ctx, &params.Req, grpc.Trailer(&st.Trailer))
}

// FoundationTestsAGDelete Основа теста для AddressGroup Delete
func (st *ToolsForTest) FoundationTestsAGDelete(params *TestCaseBodyAGDelete) (*emptypb.Empty, error) {
	ctx, cancel := utils.LocalCtx()
	defer cancel()
	return st.AddressGroupAPI.Delete(ctx, &params.Req, grpc.Trailer(&st.Trailer))
}

// TestCaseBodyHBList Структура тела запроса для HostBinding List
type TestCaseBodyHBList struct {
	TestName string
	Req      SGroupsNamespace.HostBindingReq_List
}

// TestCaseBodyHBUpsert Структура тела запроса для HostBinding Upsert
type TestCaseBodyHBUpsert struct {
	TestName string
	Req      SGroupsNamespace.HostBindingReq_Upsert
}

// TestCaseBodyHBDelete Структура тела запроса для HostBinding Delete
type TestCaseBodyHBDelete struct {
	TestName string
	Req      SGroupsNamespace.HostBindingReq_Delete
}

// FoundationTestsHBList Основа теста для HostBinding List
func (st *ToolsForTest) FoundationTestsHBList(params *TestCaseBodyHBList) (*SGroupsNamespace.HostBindingResp_List, error) {
	ctx, cancel := utils.LocalCtx()
	defer cancel()
	return st.HostBindingAPI.List(ctx, &params.Req, grpc.Trailer(&st.Trailer))
}

// FoundationTestsHBUpsert Основа теста для HostBinding Upsert
func (st *ToolsForTest) FoundationTestsHBUpsert(params *TestCaseBodyHBUpsert) (*SGroupsNamespace.HostBindingResp_Upsert, error) {
	ctx, cancel := utils.LocalCtx()
	defer cancel()
	return st.HostBindingAPI.Upsert(ctx, &params.Req, grpc.Trailer(&st.Trailer))
}

// FoundationTestsHBDelete Основа теста для HostBinding Delete
func (st *ToolsForTest) FoundationTestsHBDelete(params *TestCaseBodyHBDelete) (*emptypb.Empty, error) {
	ctx, cancel := utils.LocalCtx()
	defer cancel()
	return st.HostBindingAPI.Delete(ctx, &params.Req, grpc.Trailer(&st.Trailer))
}

// TestCaseBodyNBList Структура тела запроса для NetworkBinding List
type TestCaseBodyNBList struct {
	TestName string
	Req      SGroupsNamespace.NetworkBindingReq_List
}

// TestCaseBodyNBUpsert Структура тела запроса для NetworkBinding Upsert
type TestCaseBodyNBUpsert struct {
	TestName string
	Req      SGroupsNamespace.NetworkBindingReq_Upsert
}

// TestCaseBodyNBDelete Структура тела запроса для NetworkBinding Delete
type TestCaseBodyNBDelete struct {
	TestName string
	Req      SGroupsNamespace.NetworkBindingReq_Delete
}

// FoundationTestsNBList Основа теста для NetworkBinding List
func (st *ToolsForTest) FoundationTestsNBList(params *TestCaseBodyNBList) (*SGroupsNamespace.NetworkBindingResp_List, error) {
	ctx, cancel := utils.LocalCtx()
	defer cancel()
	return st.NetworkBindingAPI.List(ctx, &params.Req, grpc.Trailer(&st.Trailer))
}

// FoundationTestsNBUpsert Основа теста для NetworkBinding Upsert
func (st *ToolsForTest) FoundationTestsNBUpsert(params *TestCaseBodyNBUpsert) (*SGroupsNamespace.NetworkBindingResp_Upsert, error) {
	ctx, cancel := utils.LocalCtx()
	defer cancel()
	return st.NetworkBindingAPI.Upsert(ctx, &params.Req, grpc.Trailer(&st.Trailer))
}

// FoundationTestsNBDelete Основа теста для NetworkBinding Delete
func (st *ToolsForTest) FoundationTestsNBDelete(params *TestCaseBodyNBDelete) (*emptypb.Empty, error) {
	ctx, cancel := utils.LocalCtx()
	defer cancel()
	return st.NetworkBindingAPI.Delete(ctx, &params.Req, grpc.Trailer(&st.Trailer))
}

// TestCaseBodySBList Структура тела запроса для ServiceBinding List
type TestCaseBodySBList struct {
	TestName string
	Req      SGroupsNamespace.ServiceBindingReq_List
}

// TestCaseBodySBUpsert Структура тела запроса для ServiceBinding Upsert
type TestCaseBodySBUpsert struct {
	TestName string
	Req      SGroupsNamespace.ServiceBindingReq_Upsert
}

// TestCaseBodySBDelete Структура тела запроса для ServiceBinding Delete
type TestCaseBodySBDelete struct {
	TestName string
	Req      SGroupsNamespace.ServiceBindingReq_Delete
}

// FoundationTestsSBList Основа теста для ServiceBinding List
func (st *ToolsForTest) FoundationTestsSBList(params *TestCaseBodySBList) (*SGroupsNamespace.ServiceBindingResp_List, error) {
	ctx, cancel := utils.LocalCtx()
	defer cancel()
	return st.ServiceBindingAPI.List(ctx, &params.Req, grpc.Trailer(&st.Trailer))
}

// FoundationTestsSBUpsert Основа теста для ServiceBinding Upsert
func (st *ToolsForTest) FoundationTestsSBUpsert(params *TestCaseBodySBUpsert) (*SGroupsNamespace.ServiceBindingResp_Upsert, error) {
	ctx, cancel := utils.LocalCtx()
	defer cancel()
	return st.ServiceBindingAPI.Upsert(ctx, &params.Req, grpc.Trailer(&st.Trailer))
}

// FoundationTestsSBDelete Основа теста для ServiceBinding Delete
func (st *ToolsForTest) FoundationTestsSBDelete(params *TestCaseBodySBDelete) (*emptypb.Empty, error) {
	ctx, cancel := utils.LocalCtx()
	defer cancel()
	return st.ServiceBindingAPI.Delete(ctx, &params.Req, grpc.Trailer(&st.Trailer))
}
