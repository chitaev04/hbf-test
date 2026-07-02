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

func (s *GRPCSuiteListNS) TestListAllEmptyBody(t provider.T) {
	t.Parallel()
	t.Story("Поиск всех по пустому телу")
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		resp, err := s.TestStruct.FoundationTests(bodies.ListAllEmptyBody())

		sCtx.NewStep(steps.AssertStep)
		sCtx.Require().NoError(err)
		sCtx.Require().NotNil(resp)
		sCtx.Assert().Len(resp.Namespaces, 6)
	})
}

func (s *GRPCSuiteListNS) TestListAllEmptyFieldSelector(t provider.T) {
	t.Parallel()
	t.Story("Поиск всех по пустому объекту fieldSelector")
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		resp, err := s.TestStruct.FoundationTests(bodies.ListAllEmptyFieldSelector())

		sCtx.NewStep(steps.AssertStep)
		sCtx.Require().NoError(err)
		sCtx.Require().NotNil(resp)
		sCtx.Assert().Len(resp.Namespaces, 6)
	})
}

func (s *GRPCSuiteListNS) TestListAllEmptySelectorsArray(t provider.T) {
	t.Parallel()
	t.Story("Поиск всех по пустому массиву selectors")
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		resp, err := s.TestStruct.FoundationTests(bodies.ListAllEmptySelectorsArray())

		sCtx.NewStep(steps.AssertStep)
		sCtx.Require().NoError(err)
		sCtx.Require().NotNil(resp)
		sCtx.Assert().Len(resp.Namespaces, 6)
	})
}

func (s *GRPCSuiteListNS) TestListByName(t provider.T) {
	t.Parallel()
	t.Story("Поиск одного namespace по name")
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		resp, err := s.TestStruct.FoundationTests(bodies.ListByName())

		sCtx.NewStep(steps.AssertStep)
		sCtx.Require().NoError(err)
		sCtx.Require().NotNil(resp)
		sCtx.Require().Len(resp.Namespaces, 1)
		assertNamespace(sCtx, resp.Namespaces[0], "namespace-1", namespaceExpectation{
			labels:      map[string]string{"search": "labels"},
			annotations: map[string]string{},
		})
		sCtx.Assert().Equal("82874cc8-0711-40fa-be42-dd480c4cb550", resp.Namespaces[0].Metadata.Uid)
	})
}

func (s *GRPCSuiteListNS) TestListByNonExistentName(t provider.T) {
	t.Parallel()
	t.Story("Поиск namespace по несуществующему name")
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		resp, err := s.TestStruct.FoundationTests(bodies.ListByNonExistentName())

		sCtx.NewStep(steps.AssertStep)
		sCtx.Require().NoError(err)
		sCtx.Require().NotNil(resp)
		sCtx.Assert().Empty(resp.Namespaces)
	})
}

func (s *GRPCSuiteListNS) TestListByTwoNames(t provider.T) {
	t.Parallel()
	t.Story("Поиск 2 namespaces по names")
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		resp, err := s.TestStruct.FoundationTests(bodies.ListByTwoNames())

		sCtx.NewStep(steps.AssertStep)
		sCtx.Require().NoError(err)
		sCtx.Require().NotNil(resp)
		sCtx.Require().Len(resp.Namespaces, 2)
		assertNamespacesByName(sCtx, resp.Namespaces, map[string]namespaceExpectation{
			"namespace-1": {labels: map[string]string{"search": "labels"}, annotations: map[string]string{}},
			"namespace-2": {labels: map[string]string{"labels": "search"}, annotations: map[string]string{}},
		})
	})
}

func (s *GRPCSuiteListNS) TestListByExistingAndNonExistentName(t provider.T) {
	t.Parallel()
	t.Story("Поиск namespace по сущ+несущ name")
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		resp, err := s.TestStruct.FoundationTests(bodies.ListByExistingAndNonExistentName())

		sCtx.NewStep(steps.AssertStep)
		sCtx.Require().NoError(err)
		sCtx.Require().NotNil(resp)
		sCtx.Require().Len(resp.Namespaces, 1)
		assertNamespace(sCtx, resp.Namespaces[0], "namespace-1", namespaceExpectation{
			labels:      map[string]string{"search": "labels"},
			annotations: map[string]string{},
		})
		sCtx.Assert().Equal("82874cc8-0711-40fa-be42-dd480c4cb550", resp.Namespaces[0].Metadata.Uid)
	})
}

func (s *GRPCSuiteListNS) TestListByLabels(t provider.T) {
	t.Parallel()
	t.Story("Поиск одного namespace по labels")
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		resp, err := s.TestStruct.FoundationTests(bodies.ListByLabels())

		sCtx.NewStep(steps.AssertStep)
		sCtx.Require().NoError(err)
		sCtx.Require().NotNil(resp)
		sCtx.Require().Len(resp.Namespaces, 1)
		assertNamespace(sCtx, resp.Namespaces[0], "namespace-1", namespaceExpectation{
			labels:      map[string]string{"search": "labels"},
			annotations: map[string]string{},
		})
		sCtx.Assert().Equal("82874cc8-0711-40fa-be42-dd480c4cb550", resp.Namespaces[0].Metadata.Uid)
	})
}

func (s *GRPCSuiteListNS) TestListByNonExistentLabels(t provider.T) {
	t.Parallel()
	t.Story("Поиск namespace по несуществующим labels")
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		resp, err := s.TestStruct.FoundationTests(bodies.ListByNonExistentLabels())

		sCtx.NewStep(steps.AssertStep)
		sCtx.Require().NoError(err)
		sCtx.Require().NotNil(resp)
		sCtx.Assert().Empty(resp.Namespaces)
	})
}

func (s *GRPCSuiteListNS) TestListByTwoLabelSelectors(t provider.T) {
	t.Parallel()
	t.Story("Поиск 2 namespaces по labels")
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		resp, err := s.TestStruct.FoundationTests(bodies.ListByTwoLabelSelectors())

		sCtx.NewStep(steps.AssertStep)
		sCtx.Require().NoError(err)
		sCtx.Require().NotNil(resp)
		sCtx.Require().Len(resp.Namespaces, 2)
		assertNamespacesByName(sCtx, resp.Namespaces, map[string]namespaceExpectation{
			"namespace-1": {labels: map[string]string{"search": "labels"}, annotations: map[string]string{}},
			"namespace-2": {labels: map[string]string{"labels": "search"}, annotations: map[string]string{}},
		})
	})
}

func (s *GRPCSuiteListNS) TestListByExistingAndNonExistentLabels(t provider.T) {
	t.Parallel()
	t.Story("Поиск namespaces по сущ+несущ labels")
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		resp, err := s.TestStruct.FoundationTests(bodies.ListByExistingAndNonExistentLabels())

		sCtx.NewStep(steps.AssertStep)
		sCtx.Require().NoError(err)
		sCtx.Require().NotNil(resp)
		sCtx.Require().Len(resp.Namespaces, 1)
		assertNamespace(sCtx, resp.Namespaces[0], "namespace-1", namespaceExpectation{
			labels:      map[string]string{"search": "labels"},
			annotations: map[string]string{},
		})
		sCtx.Assert().Equal("82874cc8-0711-40fa-be42-dd480c4cb550", resp.Namespaces[0].Metadata.Uid)
	})
}

func (s *GRPCSuiteListNS) TestListByNameAndLabels(t provider.T) {
	t.Parallel()
	t.Story("Поиск namespace по name+labels")
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		resp, err := s.TestStruct.FoundationTests(bodies.ListByNameAndLabels())

		sCtx.NewStep(steps.AssertStep)
		sCtx.Require().NoError(err)
		sCtx.Require().NotNil(resp)
		sCtx.Assert().Len(resp.Namespaces, 1)
	})
}

func (s *GRPCSuiteListNS) TestListByTwoNameAndLabels(t provider.T) {
	t.Parallel()
	t.Story("Поиск 2 namespaces по name+labels")
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		resp, err := s.TestStruct.FoundationTests(bodies.ListByTwoNameAndLabels())

		sCtx.NewStep(steps.AssertStep)
		sCtx.Require().NoError(err)
		sCtx.Require().NotNil(resp)
		sCtx.Require().Len(resp.Namespaces, 2)
		assertNamespacesByName(sCtx, resp.Namespaces, map[string]namespaceExpectation{
			"namespace-0": {labels: map[string]string{"search": "both"}, annotations: map[string]string{"search": "both"}},
			"namespace-1": {labels: map[string]string{"search": "labels"}, annotations: map[string]string{}},
		})
	})
}

func (s *GRPCSuiteListNS) TestListByNameAndLabelsNameNotExist(t provider.T) {
	t.Parallel()
	t.Story("Поиск namespaces по name+labels. Name не существует")
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		resp, err := s.TestStruct.FoundationTests(bodies.ListByNameAndLabelsNameNotExist())

		sCtx.NewStep(steps.AssertStep)
		sCtx.Require().NoError(err)
		sCtx.Require().NotNil(resp)
		sCtx.Assert().Empty(resp.Namespaces)
	})
}

func (s *GRPCSuiteListNS) TestListByNameAndLabelsLabelsNotExist(t provider.T) {
	t.Parallel()
	t.Story("Поиск namespaces по name+labels. labels не существуют")
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		resp, err := s.TestStruct.FoundationTests(bodies.ListByNameAndLabelsLabelsNotExist())

		sCtx.NewStep(steps.AssertStep)
		sCtx.Require().NoError(err)
		sCtx.Require().NotNil(resp)
		sCtx.Assert().Empty(resp.Namespaces)
	})
}
