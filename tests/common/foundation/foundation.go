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
	NamespaceAPI SGroupsNamespace.SGroupsNamespaceAPIClient
	Trailer      metadata.MD
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
