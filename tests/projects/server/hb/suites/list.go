package suites

import (
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"

	grpcClient "HBF-tests/pkg/api/sgroups/v1"
	"HBF-tests/tests/common/connections/connection"
	"HBF-tests/tests/common/data/steps"
	"HBF-tests/tests/common/foundation"
	"HBF-tests/tests/projects/server/hb/internal/bodies"
)

type GRPCSuiteListHB struct {
	suite.Suite
	TestStruct foundation.ToolsForTest
	client     grpcClient.SGroupsHostBindingAPIClient
}

func (s *GRPCSuiteListHB) BeforeAll(t provider.T) {
	s.TestStruct.HostBindingAPI = grpcClient.NewSGroupsHostBindingAPIClient(connection.GRPC(t))
}

func (s *GRPCSuiteListHB) runList(t provider.T, story string, body *foundation.TestCaseBodyHBList, checkResp func(sCtx provider.StepCtx, resp *grpcClient.HostBindingResp_List)) {
	t.Parallel()
	t.Story(story)
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		resp, err := s.TestStruct.FoundationTestsHBList(body)

		sCtx.NewStep(steps.AssertStep)
		sCtx.Require().NoError(err)
		sCtx.Require().NotNil(resp)
		checkResp(sCtx, resp)
	})
}

func hbCheckEmpty(sCtx provider.StepCtx, resp *grpcClient.HostBindingResp_List) {
	sCtx.Assert().Empty(resp.HostBindings)
}

func hbCheckCount(n int) func(provider.StepCtx, *grpcClient.HostBindingResp_List) {
	return func(sCtx provider.StepCtx, resp *grpcClient.HostBindingResp_List) {
		sCtx.Assert().Len(resp.HostBindings, n)
	}
}

func hbCheckSingle(name string, exp hostBindingExpectation) func(provider.StepCtx, *grpcClient.HostBindingResp_List) {
	return func(sCtx provider.StepCtx, resp *grpcClient.HostBindingResp_List) {
		sCtx.Require().Len(resp.HostBindings, 1)
		assertHostBinding(sCtx, resp.HostBindings[0], name, exp)
	}
}

func hbCheckMulti(expected map[string]hostBindingExpectation) func(provider.StepCtx, *grpcClient.HostBindingResp_List) {
	return func(sCtx provider.StepCtx, resp *grpcClient.HostBindingResp_List) {
		sCtx.Assert().Len(resp.HostBindings, len(expected))
		assertHostBindingsByName(sCtx, resp.HostBindings, expected)
	}
}

func (s *GRPCSuiteListHB) TestListAll(t provider.T) {
	s.runList(t, "Поиск всех host-binding", bodies.ListAll(), hbCheckCount(6))
}

func (s *GRPCSuiteListHB) TestListByName(t provider.T) {
	s.runList(t, "Поиск host-binding по name", bodies.ListByName(), hbCheckSingle("hb-0", hb0Expectation))
}

func (s *GRPCSuiteListHB) TestListByNonExistentName(t provider.T) {
	s.runList(t, "Поиск host-binding по несуществующему name", bodies.ListByNonExistentName(), hbCheckEmpty)
}

func (s *GRPCSuiteListHB) TestListByExistingAndNonExistentName(t provider.T) {
	s.runList(t, "Поиск host-binding по существующему+несуществующему name", bodies.ListByExistingAndNonExistentName(), hbCheckSingle("hb-1", hb1Expectation))
}

func (s *GRPCSuiteListHB) TestListByTwoNames(t provider.T) {
	s.runList(t, "Поиск 2 host-binding по names", bodies.ListByTwoNames(), hbCheckMulti(hb0And1))
}

func (s *GRPCSuiteListHB) TestListByNamespace(t provider.T) {
	s.runList(t, "Поиск host-binding по namespace", bodies.ListByNamespace(), hbCheckSingle("hb-0", hb0Expectation))
}

func (s *GRPCSuiteListHB) TestListByNonExistentNamespace(t provider.T) {
	s.runList(t, "Поиск host-binding по несуществующему namespace", bodies.ListByNonExistentNamespace(), hbCheckEmpty)
}

func (s *GRPCSuiteListHB) TestListByExistingAndNonExistentNamespace(t provider.T) {
	s.runList(t, "Поиск host-binding по существующему+несуществующему namespace", bodies.ListByExistingAndNonExistentNamespace(), hbCheckSingle("hb-0", hb0Expectation))
}

func (s *GRPCSuiteListHB) TestListByTwoNamespaces(t provider.T) {
	s.runList(t, "Поиск 2 host-binding по namespaces", bodies.ListByTwoNamespaces(), hbCheckMulti(hb0And2))
}

func (s *GRPCSuiteListHB) TestListByNameAndNamespace(t provider.T) {
	s.runList(t, "Поиск host-binding по name+namespace", bodies.ListByNameAndNamespace(), hbCheckSingle("hb-0", hb0Expectation))
}

func (s *GRPCSuiteListHB) TestListByNameAndNamespaceNamespaceNotExist(t provider.T) {
	s.runList(t, "Поиск host-binding по name+namespace, namespace не существует", bodies.ListByNameAndNamespaceNamespaceNotExist(), hbCheckEmpty)
}

func (s *GRPCSuiteListHB) TestListByNameAndNamespaceNameNotExist(t provider.T) {
	s.runList(t, "Поиск host-binding по name+namespace, name не существует", bodies.ListByNameAndNamespaceNameNotExist(), hbCheckEmpty)
}

func (s *GRPCSuiteListHB) TestListByNameAndNamespaceNoneExist(t provider.T) {
	s.runList(t, "Поиск host-binding по name+namespace, ни одного из не существует", bodies.ListByNameAndNamespaceNoneExist(), hbCheckEmpty)
}

func (s *GRPCSuiteListHB) TestListByNameNamespaceLabels(t provider.T) {
	s.runList(t, "Поиск host-binding по name+namespace+labels", bodies.ListByNameNamespaceLabels(), hbCheckSingle("hb-2", hb2Expectation))
}

func (s *GRPCSuiteListHB) TestListByNameNamespaceLabelsNameNotExist(t provider.T) {
	s.runList(t, "Поиск host-binding по name+namespace+labels, name не существует", bodies.ListByNameNamespaceLabelsNameNotExist(), hbCheckEmpty)
}

func (s *GRPCSuiteListHB) TestListByNameNamespaceLabelsNamespaceNotExist(t provider.T) {
	s.runList(t, "Поиск host-binding по name+namespace+labels, namespace не существует", bodies.ListByNameNamespaceLabelsNamespaceNotExist(), hbCheckEmpty)
}

func (s *GRPCSuiteListHB) TestListByNameNamespaceLabelsLabelsNotExist(t provider.T) {
	s.runList(t, "Поиск host-binding по name+namespace+labels, labels не существует", bodies.ListByNameNamespaceLabelsLabelsNotExist(), hbCheckEmpty)
}

func (s *GRPCSuiteListHB) TestListByNameNamespaceLabelsNoneExist(t provider.T) {
	s.runList(t, "Поиск host-binding по name+namespace+labels, ни одного не существует", bodies.ListByNameNamespaceLabelsNoneExist(), hbCheckEmpty)
}

func (s *GRPCSuiteListHB) TestListByNamespaceLabels(t provider.T) {
	s.runList(t, "Поиск host-binding по namespace+labels", bodies.ListByNamespaceLabels(), hbCheckSingle("hb-2", hb2Expectation))
}

func (s *GRPCSuiteListHB) TestListByNamespaceLabelsNamespaceNotExist(t provider.T) {
	s.runList(t, "Поиск host-binding по namespace+labels, namespace не существует", bodies.ListByNamespaceLabelsNamespaceNotExist(), hbCheckEmpty)
}

func (s *GRPCSuiteListHB) TestListByNamespaceLabelsLabelsNotExist(t provider.T) {
	s.runList(t, "Поиск host-binding по namespace+labels, labels не существует", bodies.ListByNamespaceLabelsLabelsNotExist(), hbCheckEmpty)
}

func (s *GRPCSuiteListHB) TestListByNamespaceLabelsNoneExist(t provider.T) {
	s.runList(t, "Поиск host-binding по namespace+labels, ни одного из не существует", bodies.ListByNamespaceLabelsNoneExist(), hbCheckEmpty)
}

func (s *GRPCSuiteListHB) TestListByNameLabels(t provider.T) {
	s.runList(t, "Поиск host-binding по name+labels", bodies.ListByNameLabels(), hbCheckSingle("hb-2", hb2Expectation))
}

func (s *GRPCSuiteListHB) TestListByNameLabelsNameNotExist(t provider.T) {
	s.runList(t, "Поиск host-binding по name+labels, name не существует", bodies.ListByNameLabelsNameNotExist(), hbCheckEmpty)
}

func (s *GRPCSuiteListHB) TestListByNameLabelsLabelsNotExist(t provider.T) {
	s.runList(t, "Поиск host-binding по name+labels, labels не существует", bodies.ListByNameLabelsLabelsNotExist(), hbCheckEmpty)
}

func (s *GRPCSuiteListHB) TestListByNameLabelsNoneExist(t provider.T) {
	s.runList(t, "Поиск host-binding по name+labels, ни одного из не существует", bodies.ListByNameLabelsNoneExist(), hbCheckEmpty)
}

func (s *GRPCSuiteListHB) TestListByLabels(t provider.T) {
	s.runList(t, "Поиск host-binding по labels", bodies.ListByLabels(), hbCheckSingle("hb-0", hb0Expectation))
}

func (s *GRPCSuiteListHB) TestListByTwoLabels(t provider.T) {
	s.runList(t, "Поиск 2 host-binding по labels", bodies.ListByTwoLabels(), hbCheckMulti(hb0And1))
}

func (s *GRPCSuiteListHB) TestListByNonExistentLabels(t provider.T) {
	s.runList(t, "Поиск host-binding по несуществующим labels", bodies.ListByNonExistentLabels(), hbCheckEmpty)
}

func (s *GRPCSuiteListHB) TestListByExistingAndNonExistentLabels(t provider.T) {
	s.runList(t, "Поиск host-binding по сущ+несущ labels", bodies.ListByExistingAndNonExistentLabels(), hbCheckSingle("hb-0", hb0Expectation))
}

func (s *GRPCSuiteListHB) TestListByExistingAG(t provider.T) {
	s.runList(t, "Поиск host-binding по существующей AG", bodies.ListByExistingAG(), hbCheckSingle("hb-0", hb0Expectation))
}

func (s *GRPCSuiteListHB) TestListByExistingHost(t provider.T) {
	s.runList(t, "Поиск host-binding по существующему HOST", bodies.ListByExistingHost(), hbCheckSingle("hb-0", hb0Expectation))
}

func (s *GRPCSuiteListHB) TestListByExistingAndNonExistentAG(t provider.T) {
	s.runList(t, "Поиск host-binding по сущ+несущ AG", bodies.ListByExistingAndNonExistentAG(), hbCheckSingle("hb-0", hb0Expectation))
}

func (s *GRPCSuiteListHB) TestListByExistingAndNonExistentHost(t provider.T) {
	s.runList(t, "Поиск host-binding по сущ+несущ HOST", bodies.ListByExistingAndNonExistentHost(), hbCheckSingle("hb-0", hb0Expectation))
}

func (s *GRPCSuiteListHB) TestListByNonExistentAG(t provider.T) {
	s.runList(t, "Поиск host-binding по несущ AG", bodies.ListByNonExistentAG(), hbCheckEmpty)
}

func (s *GRPCSuiteListHB) TestListByNonExistentHost(t provider.T) {
	s.runList(t, "Поиск host-binding по несущ HOST", bodies.ListByNonExistentHost(), hbCheckEmpty)
}

func (s *GRPCSuiteListHB) TestListByTwoAG(t provider.T) {
	s.runList(t, "Поиск 2 host-binding по AG", bodies.ListByTwoAG(), hbCheckMulti(hb0And1))
}

func (s *GRPCSuiteListHB) TestListByTwoHost(t provider.T) {
	s.runList(t, "Поиск 2 host-binding по HOST", bodies.ListByTwoHost(), hbCheckMulti(hb0And1))
}

func (s *GRPCSuiteListHB) TestListByAGAndName(t provider.T) {
	s.runList(t, "Поиск host-binding по AG+name", bodies.ListByAGAndName(), hbCheckSingle("hb-0", hb0Expectation))
}

func (s *GRPCSuiteListHB) TestListByHostAndName(t provider.T) {
	s.runList(t, "Поиск host-binding по host+name", bodies.ListByHostAndName(), hbCheckSingle("hb-0", hb0Expectation))
}

func (s *GRPCSuiteListHB) TestListByAGNameNamespace(t provider.T) {
	s.runList(t, "Поиск host-binding по AG+name+ns", bodies.ListByAGNameNamespace(), hbCheckSingle("hb-0", hb0Expectation))
}

func (s *GRPCSuiteListHB) TestListByHostNameNamespace(t provider.T) {
	s.runList(t, "Поиск host-binding по host+name+ns", bodies.ListByHostNameNamespace(), hbCheckSingle("hb-0", hb0Expectation))
}

func (s *GRPCSuiteListHB) TestListByAGNameNamespaceLabels(t provider.T) {
	s.runList(t, "Поиск host-binding по AG+name+ns+labels", bodies.ListByAGNameNamespaceLabels(), hbCheckSingle("hb-0", hb0Expectation))
}

func (s *GRPCSuiteListHB) TestListByHostNameNamespaceLabels(t provider.T) {
	s.runList(t, "Поиск host-binding по host+name+ns+labels", bodies.ListByHostNameNamespaceLabels(), hbCheckSingle("hb-0", hb0Expectation))
}

func (s *GRPCSuiteListHB) TestListByAGNameAGNotExist(t provider.T) {
	s.runList(t, "Поиск host-binding по AG+name, AG не существует", bodies.ListByAGNameAGNotExist(), hbCheckEmpty)
}

func (s *GRPCSuiteListHB) TestListByHostNameHostNotExist(t provider.T) {
	s.runList(t, "Поиск host-binding по host+name, host не существует", bodies.ListByHostNameHostNotExist(), hbCheckEmpty)
}

func (s *GRPCSuiteListHB) TestListByAGNameNameNotExist(t provider.T) {
	s.runList(t, "Поиск host-binding по AG+name, name не существует", bodies.ListByAGNameNameNotExist(), hbCheckEmpty)
}

func (s *GRPCSuiteListHB) TestListByHostNameNameNotExist(t provider.T) {
	s.runList(t, "Поиск host-binding по host+name, name не существует", bodies.ListByHostNameNameNotExist(), hbCheckEmpty)
}

func (s *GRPCSuiteListHB) TestListByAGNameNoneExist(t provider.T) {
	s.runList(t, "Поиск host-binding по AG+name, ни одного из не существует", bodies.ListByAGNameNoneExist(), hbCheckEmpty)
}

func (s *GRPCSuiteListHB) TestListByAGAndHost(t provider.T) {
	s.runList(t, "Поиск host-binding по AG+HOST", bodies.ListByAGAndHost(), hbCheckSingle("hb-0", hb0Expectation))
}

func (s *GRPCSuiteListHB) TestListByAllParams(t provider.T) {
	s.runList(t, "Поиск host-binding по всем параметрам", bodies.ListByAllParams(), hbCheckSingle("hb-0", hb0Expectation))
}
