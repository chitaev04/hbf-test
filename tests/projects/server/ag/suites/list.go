package suites

import (
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"google.golang.org/grpc/codes"

	grpcClient "HBF-tests/pkg/api/sgroups/v1"
	"HBF-tests/tests/common/connections/connection"
	"HBF-tests/tests/common/data/steps"
	"HBF-tests/tests/common/foundation"
	"HBF-tests/tests/common/utils"
	"HBF-tests/tests/projects/server/ag/internal/bodies"
)

type GRPCSuiteListAG struct {
	suite.Suite
	TestStruct foundation.ToolsForTest
	client     grpcClient.SGroupsAddressGroupsAPIClient
}

func (s *GRPCSuiteListAG) BeforeAll(t provider.T) {
	s.TestStruct.AddressGroupAPI = grpcClient.NewSGroupsAddressGroupsAPIClient(connection.GRPC(t))
}

// runList Общий шаблон запуска List-кейса с проверкой количества найденных AG
func (s *GRPCSuiteListAG) runList(t provider.T, story string, body *foundation.TestCaseBodyAGList, checkResp func(sCtx provider.StepCtx, resp *grpcClient.AddressGroupResp_List)) {
	t.Parallel()
	t.Story(story)
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		resp, err := s.TestStruct.FoundationTestsAGList(body)

		sCtx.NewStep(steps.AssertStep)
		sCtx.Require().NoError(err)
		sCtx.Require().NotNil(resp)
		checkResp(sCtx, resp)
	})
}

func checkEmpty(sCtx provider.StepCtx, resp *grpcClient.AddressGroupResp_List) {
	sCtx.Assert().Empty(resp.AddressGroups)
}

func checkCount(n int) func(provider.StepCtx, *grpcClient.AddressGroupResp_List) {
	return func(sCtx provider.StepCtx, resp *grpcClient.AddressGroupResp_List) {
		sCtx.Assert().Len(resp.AddressGroups, n)
	}
}

func checkSingle(name string, exp addressGroupExpectation, refs map[string][2]string) func(provider.StepCtx, *grpcClient.AddressGroupResp_List) {
	return func(sCtx provider.StepCtx, resp *grpcClient.AddressGroupResp_List) {
		sCtx.Require().Len(resp.AddressGroups, 1)
		assertAddressGroup(sCtx, resp.AddressGroups[0], name, exp)
		assertAddressGroupRefs(sCtx, resp.AddressGroups[0].Refs, refs)
	}
}

func checkSingleRefsCount(name string, exp addressGroupExpectation, refsCount int) func(provider.StepCtx, *grpcClient.AddressGroupResp_List) {
	return func(sCtx provider.StepCtx, resp *grpcClient.AddressGroupResp_List) {
		sCtx.Require().Len(resp.AddressGroups, 1)
		assertAddressGroup(sCtx, resp.AddressGroups[0], name, exp)
		sCtx.Assert().Len(resp.AddressGroups[0].Refs, refsCount)
	}
}

func checkMulti(expected map[string]addressGroupExpectation) func(provider.StepCtx, *grpcClient.AddressGroupResp_List) {
	return func(sCtx provider.StepCtx, resp *grpcClient.AddressGroupResp_List) {
		sCtx.Assert().Len(resp.AddressGroups, len(expected))
		assertAddressGroupsByName(sCtx, resp.AddressGroups, expected, agRefsCount)
	}
}

var ag0And1 = map[string]addressGroupExpectation{"ag-0": ag0Expectation, "ag-1": ag1Expectation}
var ag0And2 = map[string]addressGroupExpectation{"ag-0": ag0Expectation, "ag-2": ag2Expectation}

func (s *GRPCSuiteListAG) TestListAll(t provider.T) {
	s.runList(t, "Поиск всех AG", bodies.ListAll(), checkCount(8))
}

func (s *GRPCSuiteListAG) TestListByName(t provider.T) {
	s.runList(t, "Поиск AG по name", bodies.ListByName(), checkSingle("ag-0", ag0Expectation, ag0Refs))
}

func (s *GRPCSuiteListAG) TestListByNonExistentName(t provider.T) {
	s.runList(t, "Поиск AG по несуществующему name", bodies.ListByNonExistentName(), checkEmpty)
}

func (s *GRPCSuiteListAG) TestListByExistingAndNonExistentName(t provider.T) {
	s.runList(t, "Поиск AG по существующему+несуществующему name", bodies.ListByExistingAndNonExistentName(), checkSingleRefsCount("ag-1", ag1Expectation, 14))
}

func (s *GRPCSuiteListAG) TestListByTwoNames(t provider.T) {
	s.runList(t, "Поиск 2 AG по names", bodies.ListByTwoNames(), checkMulti(ag0And1))
}

func (s *GRPCSuiteListAG) TestListByNamespace(t provider.T) {
	s.runList(t, "Поиск AG по namespace", bodies.ListByNamespace(), checkSingle("ag-0", ag0Expectation, ag0Refs))
}

func (s *GRPCSuiteListAG) TestListByNonExistentNamespace(t provider.T) {
	s.runList(t, "Поиск AG по несуществующему namespace", bodies.ListByNonExistentNamespace(), checkEmpty)
}

func (s *GRPCSuiteListAG) TestListByExistingAndNonExistentNamespace(t provider.T) {
	s.runList(t, "Поиск AG по существующему+несуществующему namespace", bodies.ListByExistingAndNonExistentNamespace(), checkSingle("ag-0", ag0Expectation, ag0Refs))
}

func (s *GRPCSuiteListAG) TestListByTwoNamespaces(t provider.T) {
	s.runList(t, "Поиск 2 AG по namespaces", bodies.ListByTwoNamespaces(), checkMulti(ag0And2))
}

func (s *GRPCSuiteListAG) TestListByNameAndNamespace(t provider.T) {
	s.runList(t, "Поиск AG по name+namespace", bodies.ListByNameAndNamespace(), checkSingle("ag-0", ag0Expectation, ag0Refs))
}

func (s *GRPCSuiteListAG) TestListByNameAndNamespaceNamespaceNotExist(t provider.T) {
	s.runList(t, "Поиск AG по name+namespace, namespace не существует", bodies.ListByNameAndNamespaceNamespaceNotExist(), checkEmpty)
}

func (s *GRPCSuiteListAG) TestListByNameAndNamespaceNameNotExist(t provider.T) {
	s.runList(t, "Поиск AG по name+namespace, name не существует", bodies.ListByNameAndNamespaceNameNotExist(), checkEmpty)
}

func (s *GRPCSuiteListAG) TestListByNameAndNamespaceNoneExist(t provider.T) {
	s.runList(t, "Поиск AG по name+namespace, ни одного из не существует", bodies.ListByNameAndNamespaceNoneExist(), checkEmpty)
}

func (s *GRPCSuiteListAG) TestListByNameNamespaceLabels(t provider.T) {
	s.runList(t, "Поиск AG по name+namespace+labels", bodies.ListByNameNamespaceLabels(), checkSingle("ag-0", ag0Expectation, ag0Refs))
}

func (s *GRPCSuiteListAG) TestListByNameNamespaceLabelsNameNotExist(t provider.T) {
	s.runList(t, "Поиск AG по name+namespace+labels, name не существует", bodies.ListByNameNamespaceLabelsNameNotExist(), checkEmpty)
}

func (s *GRPCSuiteListAG) TestListByNameNamespaceLabelsNamespaceNotExist(t provider.T) {
	s.runList(t, "Поиск AG по name+namespace+labels, namespace не существует", bodies.ListByNameNamespaceLabelsNamespaceNotExist(), checkEmpty)
}

func (s *GRPCSuiteListAG) TestListByNameNamespaceLabelsLabelsNotExist(t provider.T) {
	s.runList(t, "Поиск AG по name+namespace+labels, labels не существует", bodies.ListByNameNamespaceLabelsLabelsNotExist(), checkEmpty)
}

func (s *GRPCSuiteListAG) TestListByNameNamespaceLabelsNoneExist(t provider.T) {
	s.runList(t, "Поиск AG по name+namespace+labels, ни одного не существует", bodies.ListByNameNamespaceLabelsNoneExist(), checkEmpty)
}

func (s *GRPCSuiteListAG) TestListByNamespaceLabels(t provider.T) {
	s.runList(t, "Поиск AG по namespace+labels", bodies.ListByNamespaceLabels(), checkSingle("ag-0", ag0Expectation, ag0Refs))
}

func (s *GRPCSuiteListAG) TestListByNamespaceLabelsNamespaceNotExist(t provider.T) {
	s.runList(t, "Поиск AG по namespace+labels, namespace не существует", bodies.ListByNamespaceLabelsNamespaceNotExist(), checkEmpty)
}

func (s *GRPCSuiteListAG) TestListByNamespaceLabelsLabelsNotExist(t provider.T) {
	s.runList(t, "Поиск AG по namespace+labels, labels не существует", bodies.ListByNamespaceLabelsLabelsNotExist(), checkEmpty)
}

func (s *GRPCSuiteListAG) TestListByNamespaceLabelsNoneExist(t provider.T) {
	s.runList(t, "Поиск AG по namespace+labels, ни одного из не существует", bodies.ListByNamespaceLabelsNoneExist(), checkEmpty)
}

func (s *GRPCSuiteListAG) TestListByNameLabels(t provider.T) {
	s.runList(t, "Поиск AG по name+labels", bodies.ListByNameLabels(), checkSingle("ag-2", ag2Expectation, ag2Refs))
}

func (s *GRPCSuiteListAG) TestListByNameLabelsNameNotExist(t provider.T) {
	s.runList(t, "Поиск AG по name+labels, name не существует", bodies.ListByNameLabelsNameNotExist(), checkEmpty)
}

func (s *GRPCSuiteListAG) TestListByNameLabelsLabelsNotExist(t provider.T) {
	s.runList(t, "Поиск AG по name+labels, labels не существует", bodies.ListByNameLabelsLabelsNotExist(), checkEmpty)
}

func (s *GRPCSuiteListAG) TestListByNameLabelsNoneExist(t provider.T) {
	s.runList(t, "Поиск AG по name+labels, ни одного из не существует", bodies.ListByNameLabelsNoneExist(), checkEmpty)
}

func (s *GRPCSuiteListAG) TestListByLabels(t provider.T) {
	s.runList(t, "Поиск AG по labels", bodies.ListByLabels(), checkSingle("ag-0", ag0Expectation, ag0Refs))
}

func (s *GRPCSuiteListAG) TestListByTwoLabels(t provider.T) {
	s.runList(t, "Поиск 2 AG по labels", bodies.ListByTwoLabels(), checkMulti(ag0And1))
}

func (s *GRPCSuiteListAG) TestListByNonExistentLabels(t provider.T) {
	s.runList(t, "Поиск AG по несуществующим labels", bodies.ListByNonExistentLabels(), checkEmpty)
}

func (s *GRPCSuiteListAG) TestListByExistingAndNonExistentLabels(t provider.T) {
	s.runList(t, "Поиск AG по сущ+несущ labels", bodies.ListByExistingAndNonExistentLabels(), checkSingle("ag-0", ag0Expectation, ag0Refs))
}

func (s *GRPCSuiteListAG) TestListByRefSVC(t provider.T) {
	s.runList(t, "Поиск AG по ref (svc)", bodies.ListByRefSVC(), checkSingle("ag-0", ag0Expectation, ag0Refs))
}

func (s *GRPCSuiteListAG) TestListByRefSVCExistingAndNonExistentInOneObject(t provider.T) {
	s.runList(t, "Поиск AG по сущ+несущ ref  внутри одного объекта (svc)", bodies.ListByRefSVCExistingAndNonExistentInOneObject(), checkEmpty)
}

func (s *GRPCSuiteListAG) TestListByRefSVCExistingButNotBelongingInOneObject(t provider.T) {
	s.runList(t, "Поиск AG по сущ+непринадлежащей ref  внутри одного объекта (svc)", bodies.ListByRefSVCExistingButNotBelongingInOneObject(), checkEmpty)
}

func (s *GRPCSuiteListAG) TestListByRefSVCExistingAndNonExistentDifferentObjects(t provider.T) {
	s.runList(t, "Поиск AG по сущ+несущ ref в разных объектах (svc)", bodies.ListByRefSVCExistingAndNonExistentDifferentObjects(), checkSingle("ag-2", ag2Expectation, ag2Refs))
}

func (s *GRPCSuiteListAG) TestListByTwoRefSVC(t provider.T) {
	s.runList(t, "Поиск 2 AG по ref (svc)", bodies.ListByTwoRefSVC(), checkMulti(ag0And1))
}

func (s *GRPCSuiteListAG) TestListByRefSVCNameNotExistInNamespace(t provider.T) {
	s.runList(t, "Поиск AG по ref, name не существует в указанном ns (svc)", bodies.ListByRefSVCNameNotExistInNamespace(), checkEmpty)
}

func (s *GRPCSuiteListAG) TestListByRefSVCNameNotExist(t provider.T) {
	s.runList(t, "Поиск AG по ref, name не существует  (svc)", bodies.ListByRefSVCNameNotExist(), checkEmpty)
}

func (s *GRPCSuiteListAG) TestListByRefSVCNamespaceNotExist(t provider.T) {
	s.runList(t, "Поиск AG по ref, namespace не существует  (svc)", bodies.ListByRefSVCNamespaceNotExist(), checkEmpty)
}

func (s *GRPCSuiteListAG) TestListByRefHost(t provider.T) {
	s.runList(t, "Поиск AG по ref (host)", bodies.ListByRefHost(), checkSingle("ag-0", ag0Expectation, ag0Refs))
}

func (s *GRPCSuiteListAG) TestListByRefHostExistingAndNonExistentInOneObject(t provider.T) {
	s.runList(t, "Поиск AG по сущ+несущ ref  внутри одного объекта (host)", bodies.ListByRefHostExistingAndNonExistentInOneObject(), checkEmpty)
}

func (s *GRPCSuiteListAG) TestListByRefHostExistingAndNonExistentDifferentObjects(t provider.T) {
	s.runList(t, "Поиск AG по сущ+несущ ref в разных объектах (host)", bodies.ListByRefHostExistingAndNonExistentDifferentObjects(), checkSingle("ag-0", ag0Expectation, ag0Refs))
}

func (s *GRPCSuiteListAG) TestListByTwoRefHost(t provider.T) {
	s.runList(t, "Поиск 2 AG по ref (hosts)", bodies.ListByTwoRefHost(), checkMulti(ag0And1))
}

func (s *GRPCSuiteListAG) TestListByRefHostNameNotExistInNamespace(t provider.T) {
	s.runList(t, "Поиск AG по ref, name не существует в указанном ns (hosts)", bodies.ListByRefHostNameNotExistInNamespace(), checkEmpty)
}

func (s *GRPCSuiteListAG) TestListByRefHostNameNotExist(t provider.T) {
	s.runList(t, "Поиск AG по ref, name не существует  (hosts)", bodies.ListByRefHostNameNotExist(), checkEmpty)
}

func (s *GRPCSuiteListAG) TestListByRefHostNamespaceNotExist(t provider.T) {
	s.runList(t, "Поиск AG по ref, namespace не существует  (hosts)", bodies.ListByRefHostNamespaceNotExist(), checkEmpty)
}

func (s *GRPCSuiteListAG) TestListByRefNetwork(t provider.T) {
	s.runList(t, "Поиск AG по ref (network)", bodies.ListByRefNetwork(), checkSingle("ag-2", ag2Expectation, ag2Refs))
}

func (s *GRPCSuiteListAG) TestListByTwoRefNetwork(t provider.T) {
	s.runList(t, "Поиск 2 AG по ref (network)", bodies.ListByTwoRefNetwork(), checkMulti(ag0And1))
}

func (s *GRPCSuiteListAG) TestListByRefNetworkExistingAndNonExistentInOneObject(t provider.T) {
	s.runList(t, "Поиск AG по сущ+несущ ref  внутри одного объекта (network)", bodies.ListByRefNetworkExistingAndNonExistentInOneObject(), checkEmpty)
}

func (s *GRPCSuiteListAG) TestListByRefNetworkExistingButNotBelongingInOneObject(t provider.T) {
	s.runList(t, "Поиск AG по сущ+непринадлежащему ref  внутри одного объекта (network)", bodies.ListByRefNetworkExistingButNotBelongingInOneObject(), checkEmpty)
}

func (s *GRPCSuiteListAG) TestListByRefNetworkExistingAndNonExistentDifferentObjects(t provider.T) {
	s.runList(t, "Поиск AG по сущ+несущ ref в разных объектах (network)", bodies.ListByRefNetworkExistingAndNonExistentDifferentObjects(), checkSingle("ag-0", ag0Expectation, ag0Refs))
}

func (s *GRPCSuiteListAG) TestListByRefNetworkNameNotExistInNamespace(t provider.T) {
	s.runList(t, "Поиск AG по ref, name не существует в указанном ns (network)", bodies.ListByRefNetworkNameNotExistInNamespace(), checkEmpty)
}

func (s *GRPCSuiteListAG) TestListByRefNetworkNameNotExist(t provider.T) {
	s.runList(t, "Поиск AG по ref, name не существует  (network)", bodies.ListByRefNetworkNameNotExist(), checkEmpty)
}

func (s *GRPCSuiteListAG) TestListByRefNetworkNamespaceNotExist(t provider.T) {
	s.runList(t, "Поиск AG по ref, namespace не существует  (network)", bodies.ListByRefNetworkNamespaceNotExist(), checkEmpty)
}

func (s *GRPCSuiteListAG) TestListByThreeRefs(t provider.T) {
	s.runList(t, "Поиск AG по  3 ref (nw+svc+host)", bodies.ListByThreeRefs(), checkSingle("ag-2", ag2Expectation, ag2Refs))
}

func (s *GRPCSuiteListAG) TestListByRefAndName(t provider.T) {
	s.runList(t, "Поиск AG по ref+name", bodies.ListByRefAndName(), checkSingle("ag-2", ag2Expectation, ag2Refs))
}

func (s *GRPCSuiteListAG) TestListByRefAndNameNameNotExist(t provider.T) {
	s.runList(t, "Поиск AG по ref+name (name не существует)", bodies.ListByRefAndNameNameNotExist(), checkEmpty)
}

func (s *GRPCSuiteListAG) TestListByRefAndNameRefNotExist(t provider.T) {
	s.runList(t, "Поиск AG по ref+name (ref не существует)", bodies.ListByRefAndNameRefNotExist(), checkEmpty)
}

func (s *GRPCSuiteListAG) TestListByRefAndNameNoneExist(t provider.T) {
	s.runList(t, "Поиск AG по ref+name (Ни одного из не существует)", bodies.ListByRefAndNameNoneExist(), checkEmpty)
}

func (s *GRPCSuiteListAG) TestListByRefAndNamespace(t provider.T) {
	s.runList(t, "Поиск AG по ref+namespace", bodies.ListByRefAndNamespace(), checkSingle("ag-2", ag2Expectation, ag2Refs))
}

func (s *GRPCSuiteListAG) TestListByRefAndNamespaceNamespaceNotExist(t provider.T) {
	s.runList(t, "Поиск AG по ref+namespace (namespace не существует)", bodies.ListByRefAndNamespaceNamespaceNotExist(), checkEmpty)
}

func (s *GRPCSuiteListAG) TestListByRefAndNamespaceRefNotExist(t provider.T) {
	s.runList(t, "Поиск AG по ref+namespace (ref не существует)", bodies.ListByRefAndNamespaceRefNotExist(), checkEmpty)
}

func (s *GRPCSuiteListAG) TestListByRefAndNamespaceNoneExist(t provider.T) {
	s.runList(t, "Поиск AG по ref+namespace (ни одного из не существует)", bodies.ListByRefAndNamespaceNoneExist(), checkEmpty)
}

func (s *GRPCSuiteListAG) TestListByRefAndLabels(t provider.T) {
	s.runList(t, "Поиск AG по ref+labels", bodies.ListByRefAndLabels(), checkSingle("ag-2", ag2Expectation, ag2Refs))
}

func (s *GRPCSuiteListAG) TestListByRefAndLabelsLabelsNotExist(t provider.T) {
	s.runList(t, "Поиск AG по ref+labels (labels не существуют)", bodies.ListByRefAndLabelsLabelsNotExist(), checkEmpty)
}

func (s *GRPCSuiteListAG) TestListByRefAndLabelsRefNotExist(t provider.T) {
	s.runList(t, "Поиск AG по ref+labels (ref не существуют)", bodies.ListByRefAndLabelsRefNotExist(), checkEmpty)
}

func (s *GRPCSuiteListAG) TestListByRefAndLabelsNoneExist(t provider.T) {
	s.runList(t, "Поиск AG по ref+labels (ни одного из не существует)", bodies.ListByRefAndLabelsNoneExist(), checkEmpty)
}

func (s *GRPCSuiteListAG) TestListByAllParams(t provider.T) {
	s.runList(t, "Поиск AG по по всем параметрам (name+namespace+ref+labels)", bodies.ListByAllParams(), checkSingle("ag-2", ag2Expectation, ag2Refs))
}

func (s *GRPCSuiteListAG) TestListByAllParamsOneNotExist(t provider.T) {
	s.runList(t, "Поиск AG по по всем параметрам, один из параметров не существует", bodies.ListByAllParamsOneNotExist(), checkEmpty)
}

func (s *GRPCSuiteListAG) TestListErrorUnknownResType(t provider.T) {
	t.Parallel()
	t.Story("Ошибка при поиске AG, несущетвующий resType")
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		_, err := s.TestStruct.FoundationTestsAGList(bodies.ListErrorUnknownResType())

		sCtx.NewStep(steps.AssertStep)
		utils.AssertGRPCError(sCtx, err, codes.Internal, "invalid input value for enum resource_type")
	})
}

func (s *GRPCSuiteListAG) TestListByTwoRefRule(t provider.T) {
	s.runList(t, "Поиск 2 AG по ref (rule)", bodies.ListByTwoRefRule(), checkMulti(ag0And1))
}

func (s *GRPCSuiteListAG) TestListByRefRule(t provider.T) {
	s.runList(t, "Поиск AG по ref (rule)", bodies.ListByRefRule(), checkSingle("ag-2", ag2Expectation, ag2Refs))
}

func (s *GRPCSuiteListAG) TestListByRefRuleExistingAndNonExistentInOneObject(t provider.T) {
	s.runList(t, "Поиск AG по сущ+несущ ref  внутри одного объекта (rule)", bodies.ListByRefRuleExistingAndNonExistentInOneObject(), checkEmpty)
}

func (s *GRPCSuiteListAG) TestListByRefRuleExistingButNotBelongingInOneObject(t provider.T) {
	s.runList(t, "Поиск AG по сущ+непринадлежащему ref  внутри одного объекта (rule)", bodies.ListByRefRuleExistingButNotBelongingInOneObject(), checkEmpty)
}

func (s *GRPCSuiteListAG) TestListByRefRuleExistingAndNonExistentDifferentObjects(t provider.T) {
	s.runList(t, "Поиск AG по сущ+несущ ref в разных объектах (rule)", bodies.ListByRefRuleExistingAndNonExistentDifferentObjects(), checkSingle("ag-0", ag0Expectation, ag0Refs))
}

func (s *GRPCSuiteListAG) TestListByRefRuleNameNotExistInNamespace(t provider.T) {
	s.runList(t, "Поиск AG по ref, name не существует в указанном ns (rule)", bodies.ListByRefRuleNameNotExistInNamespace(), checkEmpty)
}

func (s *GRPCSuiteListAG) TestListByRefRuleNameNotExist(t provider.T) {
	s.runList(t, "Поиск AG по ref, name не существует  (rule)", bodies.ListByRefRuleNameNotExist(), checkEmpty)
}

func (s *GRPCSuiteListAG) TestListByRefRuleNamespaceNotExist(t provider.T) {
	s.runList(t, "Поиск AG по ref, namespace не существует  (rule)", bodies.ListByRefRuleNamespaceNotExist(), checkEmpty)
}
