package foundation

import (
	SGroupsNamespace "HBF-tests/pkg/api/sgroups/v1"
	"HBF-tests/tests/common/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// ToolsForTest Вспомогательная структура для тестов
type ToolsForTest struct {
	NamespaceAPI SGroupsNamespace.SGroupsNamespaceAPIClient
	Trailer      metadata.MD
}

// TestCaseBody Структура тела запроса
type TestCaseBody struct {
	TestName string
	Req      SGroupsNamespace.NamespaceReq_List
}

// sendReqDwv Отправка запроса
func (st *ToolsForTest) sendReqSGroupsNamespace(req *SGroupsNamespace.NamespaceReq_List) (*SGroupsNamespace.NamespaceResp_List, error) {
	ctx, cancel := utils.LocalCtx()
	defer cancel()
	return st.NamespaceAPI.List(ctx, req, grpc.Trailer(&st.Trailer))
}

// FoundationTests Основа теста
func (st *ToolsForTest) FoundationTests(params *TestCaseBody) (*SGroupsNamespace.NamespaceResp_List, error) {
	resp, err := st.sendReqSGroupsNamespace(&params.Req)
	if err != nil {
		return nil, err
	}
	return resp, nil

}
