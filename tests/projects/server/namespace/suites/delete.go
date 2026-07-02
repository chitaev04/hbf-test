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

type GRPCSuiteDeleteNS struct {
	suite.Suite
	TestStruct foundation.ToolsForTest
	client     grpcClient.SGroupsNamespaceAPIClient
}

func (s *GRPCSuiteDeleteNS) BeforeAll(t provider.T) {
	s.TestStruct.NamespaceAPI = grpcClient.NewSGroupsNamespaceAPIClient(connection.GRPC(t))
}

func (s *GRPCSuiteDeleteNS) TestList(t provider.T) {
	t.Parallel()
	t.Story("Проверки ручки list для Namespace")
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		resp, err := s.TestStruct.FoundationTests(bodies.CheckList())
		sCtx.NewStep(steps.AssertStep)
		sCtx.Require().Equal(resp.Namespaces[0].Metadata.Name, 123)
		sCtx.Require().NoError(err)
		sCtx.Require().NotNil(resp)
		sCtx.Assert().NotEmpty(resp.Namespaces[0].Metadata.Uid)
	})

}
