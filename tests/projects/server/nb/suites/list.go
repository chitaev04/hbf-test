package suites

import (
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"

	grpcClient "HBF-tests/pkg/api/sgroups/v1"
	"HBF-tests/tests/common/connections/connection"
	"HBF-tests/tests/common/data/steps"
	"HBF-tests/tests/common/foundation"
	"HBF-tests/tests/projects/server/nb/internal/bodies"
)

type GRPCSuiteListNB struct {
	suite.Suite
	TestStruct foundation.ToolsForTest
	client     grpcClient.SGroupsNetworkBindingAPIClient
}

func (s *GRPCSuiteListNB) BeforeAll(t provider.T) {
	s.TestStruct.NetworkBindingAPI = grpcClient.NewSGroupsNetworkBindingAPIClient(connection.GRPC(t))
}

func (s *GRPCSuiteListNB) runList(t provider.T, story string, body *foundation.TestCaseBodyNBList, checkResp func(sCtx provider.StepCtx, resp *grpcClient.NetworkBindingResp_List)) {
	t.Parallel()
	t.Story(story)
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		resp, err := s.TestStruct.FoundationTestsNBList(body)

		sCtx.NewStep(steps.AssertStep)
		sCtx.Require().NoError(err)
		sCtx.Require().NotNil(resp)
		checkResp(sCtx, resp)
	})
}

func nbCheckEmpty(sCtx provider.StepCtx, resp *grpcClient.NetworkBindingResp_List) {
	sCtx.Assert().Empty(resp.NetworkBindings)
}

func nbCheckCount(n int) func(provider.StepCtx, *grpcClient.NetworkBindingResp_List) {
	return func(sCtx provider.StepCtx, resp *grpcClient.NetworkBindingResp_List) {
		sCtx.Assert().Len(resp.NetworkBindings, n)
	}
}

func nbCheckSingle(name string, exp networkBindingExpectation) func(provider.StepCtx, *grpcClient.NetworkBindingResp_List) {
	return func(sCtx provider.StepCtx, resp *grpcClient.NetworkBindingResp_List) {
		sCtx.Require().Len(resp.NetworkBindings, 1)
		assertNetworkBinding(sCtx, resp.NetworkBindings[0], name, exp)
	}
}

func nbCheckMulti(expected map[string]networkBindingExpectation) func(provider.StepCtx, *grpcClient.NetworkBindingResp_List) {
	return func(sCtx provider.StepCtx, resp *grpcClient.NetworkBindingResp_List) {
		sCtx.Assert().Len(resp.NetworkBindings, len(expected))
		assertNetworkBindingsByName(sCtx, resp.NetworkBindings, expected)
	}
}

func (s *GRPCSuiteListNB) TestListAll(t provider.T) {
	s.runList(t, "Поиск всех network-binding", bodies.ListAll(), nbCheckCount(6))
}

func (s *GRPCSuiteListNB) TestListByName(t provider.T) {
	s.runList(t, "Поиск network-binding по name", bodies.ListByName(), nbCheckSingle("nb-0", nb0Expectation))
}

func (s *GRPCSuiteListNB) TestListByNonExistentName(t provider.T) {
	s.runList(t, "Поиск network-binding по несуществующему name", bodies.ListByNonExistentName(), nbCheckEmpty)
}

func (s *GRPCSuiteListNB) TestListByExistingAndNonExistentName(t provider.T) {
	s.runList(t, "Поиск network-binding по существующему+несуществующему name", bodies.ListByExistingAndNonExistentName(), nbCheckSingle("nb-1", nb1Expectation))
}

func (s *GRPCSuiteListNB) TestListByTwoNames(t provider.T) {
	s.runList(t, "Поиск 2 network-binding по names", bodies.ListByTwoNames(), nbCheckMulti(nb0And1))
}

func (s *GRPCSuiteListNB) TestListByNamespace(t provider.T) {
	s.runList(t, "Поиск network-binding по namespace", bodies.ListByNamespace(), nbCheckSingle("nb-0", nb0Expectation))
}

func (s *GRPCSuiteListNB) TestListByNonExistentNamespace(t provider.T) {
	s.runList(t, "Поиск network-binding по несуществующему namespace", bodies.ListByNonExistentNamespace(), nbCheckEmpty)
}

func (s *GRPCSuiteListNB) TestListByExistingAndNonExistentNamespace(t provider.T) {
	s.runList(t, "Поиск network-binding по существующему+несуществующему namespace", bodies.ListByExistingAndNonExistentNamespace(), nbCheckSingle("nb-0", nb0Expectation))
}

func (s *GRPCSuiteListNB) TestListByTwoNamespaces(t provider.T) {
	s.runList(t, "Поиск 2 network-binding по namespaces", bodies.ListByTwoNamespaces(), nbCheckMulti(nb0And2))
}

func (s *GRPCSuiteListNB) TestListByNameAndNamespace(t provider.T) {
	s.runList(t, "Поиск network-binding по name+namespace", bodies.ListByNameAndNamespace(), nbCheckSingle("nb-0", nb0Expectation))
}

func (s *GRPCSuiteListNB) TestListByNameAndNamespaceNamespaceNotExist(t provider.T) {
	s.runList(t, "Поиск network-binding по name+namespace, namespace не существует", bodies.ListByNameAndNamespaceNamespaceNotExist(), nbCheckEmpty)
}

func (s *GRPCSuiteListNB) TestListByNameAndNamespaceNameNotExist(t provider.T) {
	s.runList(t, "Поиск network-binding по name+namespace, name не существует", bodies.ListByNameAndNamespaceNameNotExist(), nbCheckEmpty)
}

func (s *GRPCSuiteListNB) TestListByNameAndNamespaceNoneExist(t provider.T) {
	s.runList(t, "Поиск network-binding по name+namespace, ни одного из не существует", bodies.ListByNameAndNamespaceNoneExist(), nbCheckEmpty)
}

func (s *GRPCSuiteListNB) TestListByNameNamespaceLabels(t provider.T) {
	s.runList(t, "Поиск network-binding по name+namespace+labels", bodies.ListByNameNamespaceLabels(), nbCheckSingle("nb-2", nb2Expectation))
}

func (s *GRPCSuiteListNB) TestListByNameNamespaceLabelsNameNotExist(t provider.T) {
	s.runList(t, "Поиск network-binding по name+namespace+labels, name не существует", bodies.ListByNameNamespaceLabelsNameNotExist(), nbCheckEmpty)
}

func (s *GRPCSuiteListNB) TestListByNameNamespaceLabelsNamespaceNotExist(t provider.T) {
	s.runList(t, "Поиск network-binding по name+namespace+labels, namespace не существует", bodies.ListByNameNamespaceLabelsNamespaceNotExist(), nbCheckEmpty)
}

func (s *GRPCSuiteListNB) TestListByNameNamespaceLabelsLabelsNotExist(t provider.T) {
	s.runList(t, "Поиск network-binding по name+namespace+labels, labels не существует", bodies.ListByNameNamespaceLabelsLabelsNotExist(), nbCheckEmpty)
}

func (s *GRPCSuiteListNB) TestListByNameNamespaceLabelsNoneExist(t provider.T) {
	s.runList(t, "Поиск network-binding по name+namespace+labels, ни одного не существует", bodies.ListByNameNamespaceLabelsNoneExist(), nbCheckEmpty)
}

func (s *GRPCSuiteListNB) TestListByNamespaceLabels(t provider.T) {
	s.runList(t, "Поиск network-binding по namespace+labels", bodies.ListByNamespaceLabels(), nbCheckSingle("nb-2", nb2Expectation))
}

func (s *GRPCSuiteListNB) TestListByNamespaceLabelsNamespaceNotExist(t provider.T) {
	s.runList(t, "Поиск network-binding по namespace+labels, namespace не существует", bodies.ListByNamespaceLabelsNamespaceNotExist(), nbCheckEmpty)
}

func (s *GRPCSuiteListNB) TestListByNamespaceLabelsLabelsNotExist(t provider.T) {
	s.runList(t, "Поиск network-binding по namespace+labels, labels не существует", bodies.ListByNamespaceLabelsLabelsNotExist(), nbCheckEmpty)
}

func (s *GRPCSuiteListNB) TestListByNamespaceLabelsNoneExist(t provider.T) {
	s.runList(t, "Поиск network-binding по namespace+labels, ни одного из не существует", bodies.ListByNamespaceLabelsNoneExist(), nbCheckEmpty)
}

func (s *GRPCSuiteListNB) TestListByNameLabels(t provider.T) {
	s.runList(t, "Поиск network-binding по name+labels", bodies.ListByNameLabels(), nbCheckSingle("nb-2", nb2Expectation))
}

func (s *GRPCSuiteListNB) TestListByNameLabelsNameNotExist(t provider.T) {
	s.runList(t, "Поиск network-binding по name+labels, name не существует", bodies.ListByNameLabelsNameNotExist(), nbCheckEmpty)
}

func (s *GRPCSuiteListNB) TestListByNameLabelsLabelsNotExist(t provider.T) {
	s.runList(t, "Поиск network-binding по name+labels, labels не существует", bodies.ListByNameLabelsLabelsNotExist(), nbCheckEmpty)
}

func (s *GRPCSuiteListNB) TestListByNameLabelsNoneExist(t provider.T) {
	s.runList(t, "Поиск network-binding по name+labels, ни одного из не существует", bodies.ListByNameLabelsNoneExist(), nbCheckEmpty)
}

func (s *GRPCSuiteListNB) TestListByLabels(t provider.T) {
	s.runList(t, "Поиск network-binding по labels", bodies.ListByLabels(), nbCheckSingle("nb-0", nb0Expectation))
}

func (s *GRPCSuiteListNB) TestListByTwoLabels(t provider.T) {
	s.runList(t, "Поиск 2 network-binding по labels", bodies.ListByTwoLabels(), nbCheckMulti(nb0And1))
}

func (s *GRPCSuiteListNB) TestListByNonExistentLabels(t provider.T) {
	s.runList(t, "Поиск network-binding по несуществующим labels", bodies.ListByNonExistentLabels(), nbCheckEmpty)
}

func (s *GRPCSuiteListNB) TestListByExistingAndNonExistentLabels(t provider.T) {
	s.runList(t, "Поиск network-binding по сущ+несущ labels", bodies.ListByExistingAndNonExistentLabels(), nbCheckSingle("nb-0", nb0Expectation))
}

func (s *GRPCSuiteListNB) TestListByExistingAG(t provider.T) {
	s.runList(t, "Поиск network-binding по существующей AG", bodies.ListByExistingAG(), nbCheckSingle("nb-0", nb0Expectation))
}

func (s *GRPCSuiteListNB) TestListByExistingNetwork(t provider.T) {
	s.runList(t, "Поиск network-binding по существующему nw", bodies.ListByExistingNetwork(), nbCheckSingle("nb-0", nb0Expectation))
}

func (s *GRPCSuiteListNB) TestListByExistingAndNonExistentAG(t provider.T) {
	s.runList(t, "Поиск network-binding по сущ+несущ AG", bodies.ListByExistingAndNonExistentAG(), nbCheckSingle("nb-0", nb0Expectation))
}

func (s *GRPCSuiteListNB) TestListByExistingAndNonExistentNetwork(t provider.T) {
	s.runList(t, "Поиск network-binding по сущ+несущ nw", bodies.ListByExistingAndNonExistentNetwork(), nbCheckSingle("nb-0", nb0Expectation))
}

func (s *GRPCSuiteListNB) TestListByNonExistentAG(t provider.T) {
	s.runList(t, "Поиск network-binding по несущ AG", bodies.ListByNonExistentAG(), nbCheckEmpty)
}

func (s *GRPCSuiteListNB) TestListByNonExistentNetwork(t provider.T) {
	s.runList(t, "Поиск network-binding по несущ nw", bodies.ListByNonExistentNetwork(), nbCheckEmpty)
}

func (s *GRPCSuiteListNB) TestListByTwoAG(t provider.T) {
	s.runList(t, "Поиск 2 network-binding по AG", bodies.ListByTwoAG(), nbCheckMulti(nb0And1))
}

func (s *GRPCSuiteListNB) TestListByTwoNetwork(t provider.T) {
	s.runList(t, "Поиск 2 network-binding по nw", bodies.ListByTwoNetwork(), nbCheckMulti(nb0And1))
}

func (s *GRPCSuiteListNB) TestListByAGAndName(t provider.T) {
	s.runList(t, "Поиск network-binding по AG+name", bodies.ListByAGAndName(), nbCheckSingle("nb-0", nb0Expectation))
}

func (s *GRPCSuiteListNB) TestListByNetworkAndName(t provider.T) {
	s.runList(t, "Поиск network-binding по nw+name", bodies.ListByNetworkAndName(), nbCheckSingle("nb-0", nb0Expectation))
}

func (s *GRPCSuiteListNB) TestListByAGNameNamespace(t provider.T) {
	s.runList(t, "Поиск network-binding по AG+name+ns", bodies.ListByAGNameNamespace(), nbCheckSingle("nb-0", nb0Expectation))
}

func (s *GRPCSuiteListNB) TestListByNetworkNameNamespace(t provider.T) {
	s.runList(t, "Поиск network-binding по nw+name+ns", bodies.ListByNetworkNameNamespace(), nbCheckSingle("nb-0", nb0Expectation))
}

func (s *GRPCSuiteListNB) TestListByAGNameNamespaceLabels(t provider.T) {
	s.runList(t, "Поиск network-binding по AG+name+ns+labels", bodies.ListByAGNameNamespaceLabels(), nbCheckSingle("nb-0", nb0Expectation))
}

func (s *GRPCSuiteListNB) TestListByNetworkNameNamespaceLabels(t provider.T) {
	s.runList(t, "Поиск network-binding по nw+name+ns+labels", bodies.ListByNetworkNameNamespaceLabels(), nbCheckSingle("nb-0", nb0Expectation))
}

func (s *GRPCSuiteListNB) TestListByAGNameAGNotExist(t provider.T) {
	s.runList(t, "Поиск network-binding по AG+name, AG не существует", bodies.ListByAGNameAGNotExist(), nbCheckEmpty)
}

func (s *GRPCSuiteListNB) TestListByNetworkNameNetworkNotExist(t provider.T) {
	s.runList(t, "Поиск network-binding по nw+name, nw не существует", bodies.ListByNetworkNameNetworkNotExist(), nbCheckEmpty)
}

func (s *GRPCSuiteListNB) TestListByAGNameNameNotExist(t provider.T) {
	s.runList(t, "Поиск network-binding по AG+name, name не существует", bodies.ListByAGNameNameNotExist(), nbCheckEmpty)
}

func (s *GRPCSuiteListNB) TestListByNetworkNameNameNotExist(t provider.T) {
	s.runList(t, "Поиск network-binding по nw+name, name не существует", bodies.ListByNetworkNameNameNotExist(), nbCheckEmpty)
}

func (s *GRPCSuiteListNB) TestListByAGNameNoneExist(t provider.T) {
	s.runList(t, "Поиск network-binding по AG+name, ни одного из не существует", bodies.ListByAGNameNoneExist(), nbCheckEmpty)
}

func (s *GRPCSuiteListNB) TestListByAGAndNetwork(t provider.T) {
	s.runList(t, "Поиск network-binding по AG+NW", bodies.ListByAGAndNetwork(), nbCheckSingle("nb-0", nb0Expectation))
}

func (s *GRPCSuiteListNB) TestListByAllParams(t provider.T) {
	s.runList(t, "Поиск network-binding по всем параметрам", bodies.ListByAllParams(), nbCheckSingle("nb-0", nb0Expectation))
}
