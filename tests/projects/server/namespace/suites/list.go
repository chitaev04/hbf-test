package suites

import (
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"

	grpcClient "HBF-tests/pkg/api/sgroups/v1"
	"HBF-tests/tests/common/connections/connection"
	"HBF-tests/tests/common/data/steps"
	"HBF-tests/tests/common/foundation"
	"HBF-tests/tests/projects/server/namespace/internal/bodies"
)

type GRPCSuiteListNS struct {
	suite.Suite
	TestStruct foundation.ToolsForTest
	client     grpcClient.SGroupsNamespaceAPIClient
}

func (s *GRPCSuiteListNS) BeforeAll(t provider.T) {
	s.TestStruct.NamespaceAPI = grpcClient.NewSGroupsNamespaceAPIClient(connection.GRPC(t))
}

func (s *GRPCSuiteListNS) TestList(t provider.T) {
	t.Parallel()
	t.Story("Проверка ручки list для Namespace. Поиск всех по пустому телу")
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		resp, err := s.TestStruct.FoundationTests(bodies.CheckList())

		sCtx.NewStep(steps.AssertStep)
		sCtx.Require().NoError(err)
		sCtx.Require().NotNil(resp)
		sCtx.Require().NotEmpty(resp.Namespaces)
		sCtx.Assert().NotEmpty(resp.Namespaces[0].Metadata.Uid)
		sCtx.Assert().Equal("add-namespace-1", resp.Namespaces[0].Metadata.Name)
	})
}
