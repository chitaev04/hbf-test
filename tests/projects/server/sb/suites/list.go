package suites

import (
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"

	grpcClient "HBF-tests/pkg/api/sgroups/v1"
	"HBF-tests/tests/common/connections/connection"
	"HBF-tests/tests/common/data/steps"
	"HBF-tests/tests/common/foundation"
	"HBF-tests/tests/projects/server/sb/internal/bodies"
)

type GRPCSuiteListSB struct {
	suite.Suite
	TestStruct foundation.ToolsForTest
	client     grpcClient.SGroupsServiceBindingAPIClient
}

func (s *GRPCSuiteListSB) BeforeAll(t provider.T) {
	s.TestStruct.ServiceBindingAPI = grpcClient.NewSGroupsServiceBindingAPIClient(connection.GRPC(t))
}

func (s *GRPCSuiteListSB) runList(t provider.T, story string, body *foundation.TestCaseBodySBList, checkResp func(sCtx provider.StepCtx, resp *grpcClient.ServiceBindingResp_List)) {
	t.Parallel()
	t.Story(story)
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		resp, err := s.TestStruct.FoundationTestsSBList(body)

		sCtx.NewStep(steps.AssertStep)
		sCtx.Require().NoError(err)
		sCtx.Require().NotNil(resp)
		checkResp(sCtx, resp)
	})
}

func sbCheckEmpty(sCtx provider.StepCtx, resp *grpcClient.ServiceBindingResp_List) {
	sCtx.Assert().Empty(resp.ServiceBindings)
}

func sbCheckCount(n int) func(provider.StepCtx, *grpcClient.ServiceBindingResp_List) {
	return func(sCtx provider.StepCtx, resp *grpcClient.ServiceBindingResp_List) {
		sCtx.Assert().Len(resp.ServiceBindings, n)
	}
}

func sbCheckSingle(name string, exp serviceBindingExpectation) func(provider.StepCtx, *grpcClient.ServiceBindingResp_List) {
	return func(sCtx provider.StepCtx, resp *grpcClient.ServiceBindingResp_List) {
		sCtx.Require().Len(resp.ServiceBindings, 1)
		assertServiceBinding(sCtx, resp.ServiceBindings[0], name, exp)
	}
}

func sbCheckMulti(expected map[string]serviceBindingExpectation) func(provider.StepCtx, *grpcClient.ServiceBindingResp_List) {
	return func(sCtx provider.StepCtx, resp *grpcClient.ServiceBindingResp_List) {
		sCtx.Assert().Len(resp.ServiceBindings, len(expected))
		assertServiceBindingsByName(sCtx, resp.ServiceBindings, expected)
	}
}

func (s *GRPCSuiteListSB) TestListAll(t provider.T) {
	s.runList(t, "Поиск всех service-binding", bodies.ListAll(), sbCheckCount(6))
}

func (s *GRPCSuiteListSB) TestListByName(t provider.T) {
	s.runList(t, "Поиск service-binding по name", bodies.ListByName(), sbCheckSingle("sb-0", sb0Expectation))
}

func (s *GRPCSuiteListSB) TestListByNonExistentName(t provider.T) {
	s.runList(t, "Поиск service-binding по несуществующему name", bodies.ListByNonExistentName(), sbCheckEmpty)
}

func (s *GRPCSuiteListSB) TestListByExistingAndNonExistentName(t provider.T) {
	s.runList(t, "Поиск service-binding по существующему+несуществующему name", bodies.ListByExistingAndNonExistentName(), sbCheckSingle("sb-1", sb1Expectation))
}

func (s *GRPCSuiteListSB) TestListByTwoNames(t provider.T) {
	s.runList(t, "Поиск 2 service-binding по names", bodies.ListByTwoNames(), sbCheckMulti(sb0And1))
}

func (s *GRPCSuiteListSB) TestListByNamespace(t provider.T) {
	s.runList(t, "Поиск service-binding по namespace", bodies.ListByNamespace(), sbCheckSingle("sb-0", sb0Expectation))
}

func (s *GRPCSuiteListSB) TestListByNonExistentNamespace(t provider.T) {
	s.runList(t, "Поиск service-binding по несуществующему namespace", bodies.ListByNonExistentNamespace(), sbCheckEmpty)
}

func (s *GRPCSuiteListSB) TestListByExistingAndNonExistentNamespace(t provider.T) {
	s.runList(t, "Поиск service-binding по существующему+несуществующему namespace", bodies.ListByExistingAndNonExistentNamespace(), sbCheckSingle("sb-0", sb0Expectation))
}

func (s *GRPCSuiteListSB) TestListByTwoNamespaces(t provider.T) {
	s.runList(t, "Поиск 2 service-binding по namespaces", bodies.ListByTwoNamespaces(), sbCheckMulti(sb0And2))
}

func (s *GRPCSuiteListSB) TestListByNameAndNamespace(t provider.T) {
	s.runList(t, "Поиск service-binding по name+namespace", bodies.ListByNameAndNamespace(), sbCheckSingle("sb-0", sb0Expectation))
}

func (s *GRPCSuiteListSB) TestListByNameAndNamespaceNamespaceNotExist(t provider.T) {
	s.runList(t, "Поиск service-binding по name+namespace, namespace не существует", bodies.ListByNameAndNamespaceNamespaceNotExist(), sbCheckEmpty)
}

func (s *GRPCSuiteListSB) TestListByNameAndNamespaceNameNotExist(t provider.T) {
	s.runList(t, "Поиск service-binding по name+namespace, name не существует", bodies.ListByNameAndNamespaceNameNotExist(), sbCheckEmpty)
}

func (s *GRPCSuiteListSB) TestListByNameAndNamespaceNoneExist(t provider.T) {
	s.runList(t, "Поиск service-binding по name+namespace, ни одного из не существует", bodies.ListByNameAndNamespaceNoneExist(), sbCheckEmpty)
}

func (s *GRPCSuiteListSB) TestListByNameNamespaceLabels(t provider.T) {
	s.runList(t, "Поиск service-binding по name+namespace+labels", bodies.ListByNameNamespaceLabels(), sbCheckSingle("sb-2", sb2Expectation))
}

func (s *GRPCSuiteListSB) TestListByNameNamespaceLabelsNameNotExist(t provider.T) {
	s.runList(t, "Поиск service-binding по name+namespace+labels, name не существует", bodies.ListByNameNamespaceLabelsNameNotExist(), sbCheckEmpty)
}

func (s *GRPCSuiteListSB) TestListByNameNamespaceLabelsNamespaceNotExist(t provider.T) {
	s.runList(t, "Поиск service-binding по name+namespace+labels, namespace не существует", bodies.ListByNameNamespaceLabelsNamespaceNotExist(), sbCheckEmpty)
}

func (s *GRPCSuiteListSB) TestListByNameNamespaceLabelsLabelsNotExist(t provider.T) {
	s.runList(t, "Поиск service-binding по name+namespace+labels, labels не существует", bodies.ListByNameNamespaceLabelsLabelsNotExist(), sbCheckEmpty)
}

func (s *GRPCSuiteListSB) TestListByNameNamespaceLabelsNoneExist(t provider.T) {
	s.runList(t, "Поиск service-binding по name+namespace+labels, ни одного не существует", bodies.ListByNameNamespaceLabelsNoneExist(), sbCheckEmpty)
}

func (s *GRPCSuiteListSB) TestListByNamespaceLabels(t provider.T) {
	s.runList(t, "Поиск service-binding по namespace+labels", bodies.ListByNamespaceLabels(), sbCheckSingle("sb-2", sb2Expectation))
}

func (s *GRPCSuiteListSB) TestListByNamespaceLabelsNamespaceNotExist(t provider.T) {
	s.runList(t, "Поиск service-binding по namespace+labels, namespace не существует", bodies.ListByNamespaceLabelsNamespaceNotExist(), sbCheckEmpty)
}

func (s *GRPCSuiteListSB) TestListByNamespaceLabelsLabelsNotExist(t provider.T) {
	s.runList(t, "Поиск service-binding по namespace+labels, labels не существует", bodies.ListByNamespaceLabelsLabelsNotExist(), sbCheckEmpty)
}

func (s *GRPCSuiteListSB) TestListByNamespaceLabelsNoneExist(t provider.T) {
	s.runList(t, "Поиск service-binding по namespace+labels, ни одного из не существует", bodies.ListByNamespaceLabelsNoneExist(), sbCheckEmpty)
}

func (s *GRPCSuiteListSB) TestListByNameLabels(t provider.T) {
	s.runList(t, "Поиск service-binding по name+labels", bodies.ListByNameLabels(), sbCheckSingle("sb-2", sb2Expectation))
}

func (s *GRPCSuiteListSB) TestListByNameLabelsNameNotExist(t provider.T) {
	s.runList(t, "Поиск service-binding по name+labels, name не существует", bodies.ListByNameLabelsNameNotExist(), sbCheckEmpty)
}

func (s *GRPCSuiteListSB) TestListByNameLabelsLabelsNotExist(t provider.T) {
	s.runList(t, "Поиск service-binding по name+labels, labels не существует", bodies.ListByNameLabelsLabelsNotExist(), sbCheckEmpty)
}

func (s *GRPCSuiteListSB) TestListByNameLabelsNoneExist(t provider.T) {
	s.runList(t, "Поиск service-binding по name+labels, ни одного из не существует", bodies.ListByNameLabelsNoneExist(), sbCheckEmpty)
}

func (s *GRPCSuiteListSB) TestListByLabels(t provider.T) {
	s.runList(t, "Поиск service-binding по labels", bodies.ListByLabels(), sbCheckSingle("sb-0", sb0Expectation))
}

func (s *GRPCSuiteListSB) TestListByTwoLabels(t provider.T) {
	s.runList(t, "Поиск 2 service-binding по labels", bodies.ListByTwoLabels(), sbCheckMulti(sb0And1))
}

func (s *GRPCSuiteListSB) TestListByNonExistentLabels(t provider.T) {
	s.runList(t, "Поиск service-binding по несуществующим labels", bodies.ListByNonExistentLabels(), sbCheckEmpty)
}

func (s *GRPCSuiteListSB) TestListByExistingAndNonExistentLabels(t provider.T) {
	s.runList(t, "Поиск service-binding по сущ+несущ labels", bodies.ListByExistingAndNonExistentLabels(), sbCheckSingle("sb-0", sb0Expectation))
}

func (s *GRPCSuiteListSB) TestListByExistingAG(t provider.T) {
	s.runList(t, "Поиск service-binding по существующей AG", bodies.ListByExistingAG(), sbCheckSingle("sb-0", sb0Expectation))
}

func (s *GRPCSuiteListSB) TestListByExistingSVC(t provider.T) {
	s.runList(t, "Поиск service-binding по существующему svc", bodies.ListByExistingSVC(), sbCheckSingle("sb-0", sb0Expectation))
}

func (s *GRPCSuiteListSB) TestListByExistingAndNonExistentAG(t provider.T) {
	s.runList(t, "Поиск service-binding по сущ+несущ AG", bodies.ListByExistingAndNonExistentAG(), sbCheckSingle("sb-0", sb0Expectation))
}

func (s *GRPCSuiteListSB) TestListByExistingAndNonExistentSVC(t provider.T) {
	s.runList(t, "Поиск service-binding по сущ+несущ svc", bodies.ListByExistingAndNonExistentSVC(), sbCheckSingle("sb-0", sb0Expectation))
}

func (s *GRPCSuiteListSB) TestListByNonExistentAG(t provider.T) {
	s.runList(t, "Поиск service-binding по несущ AG", bodies.ListByNonExistentAG(), sbCheckEmpty)
}

func (s *GRPCSuiteListSB) TestListByNonExistentSVC(t provider.T) {
	s.runList(t, "Поиск service-binding по несущ svc", bodies.ListByNonExistentSVC(), sbCheckEmpty)
}

func (s *GRPCSuiteListSB) TestListByTwoAG(t provider.T) {
	s.runList(t, "Поиск 2 service-binding по AG", bodies.ListByTwoAG(), sbCheckMulti(sb0And1))
}

func (s *GRPCSuiteListSB) TestListByTwoSVC(t provider.T) {
	s.runList(t, "Поиск 2 service-binding по svc", bodies.ListByTwoSVC(), sbCheckMulti(sb0And1))
}

func (s *GRPCSuiteListSB) TestListByAGAndName(t provider.T) {
	s.runList(t, "Поиск service-binding по AG+name", bodies.ListByAGAndName(), sbCheckSingle("sb-0", sb0Expectation))
}

func (s *GRPCSuiteListSB) TestListBySVCAndName(t provider.T) {
	s.runList(t, "Поиск service-binding по svc+name", bodies.ListBySVCAndName(), sbCheckSingle("sb-0", sb0Expectation))
}

func (s *GRPCSuiteListSB) TestListByAGNameNamespace(t provider.T) {
	s.runList(t, "Поиск service-binding по AG+name+ns", bodies.ListByAGNameNamespace(), sbCheckSingle("sb-0", sb0Expectation))
}

func (s *GRPCSuiteListSB) TestListBySVCNameNamespace(t provider.T) {
	s.runList(t, "Поиск service-binding по svc+name+ns", bodies.ListBySVCNameNamespace(), sbCheckSingle("sb-0", sb0Expectation))
}

func (s *GRPCSuiteListSB) TestListByAGNameNamespaceLabels(t provider.T) {
	s.runList(t, "Поиск service-binding по AG+name+ns+labels", bodies.ListByAGNameNamespaceLabels(), sbCheckSingle("sb-0", sb0Expectation))
}

func (s *GRPCSuiteListSB) TestListBySVCNameNamespaceLabels(t provider.T) {
	s.runList(t, "Поиск service-binding по svc+name+ns+labels", bodies.ListBySVCNameNamespaceLabels(), sbCheckSingle("sb-0", sb0Expectation))
}

func (s *GRPCSuiteListSB) TestListByAGNameAGNotExist(t provider.T) {
	s.runList(t, "Поиск service-binding по AG+name, AG не существует", bodies.ListByAGNameAGNotExist(), sbCheckEmpty)
}

func (s *GRPCSuiteListSB) TestListBySVCNameSVCNotExist(t provider.T) {
	s.runList(t, "Поиск service-binding по host+name, host не существует", bodies.ListBySVCNameSVCNotExist(), sbCheckEmpty)
}

func (s *GRPCSuiteListSB) TestListByAGNameNameNotExist(t provider.T) {
	s.runList(t, "Поиск service-binding по AG+name, name не существует", bodies.ListByAGNameNameNotExist(), sbCheckEmpty)
}

func (s *GRPCSuiteListSB) TestListBySVCNameNameNotExist(t provider.T) {
	s.runList(t, "Поиск service-binding по svc+name, name не существует", bodies.ListBySVCNameNameNotExist(), sbCheckEmpty)
}

func (s *GRPCSuiteListSB) TestListByAGNameNoneExist(t provider.T) {
	s.runList(t, "Поиск service-binding по AG+name, ни одного из не существует", bodies.ListByAGNameNoneExist(), sbCheckEmpty)
}

func (s *GRPCSuiteListSB) TestListByAGAndSVC(t provider.T) {
	s.runList(t, "Поиск service-binding по AG+SVC", bodies.ListByAGAndSVC(), sbCheckSingle("sb-0", sb0Expectation))
}

func (s *GRPCSuiteListSB) TestListByAllParams(t provider.T) {
	s.runList(t, "Поиск service-binding по всем параметрам", bodies.ListByAllParams(), sbCheckSingle("sb-0", sb0Expectation))
}
