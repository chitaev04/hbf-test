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

type GRPCSuiteUpsetNS struct {
	suite.Suite
	TestStruct foundation.ToolsForTest
	client     grpcClient.SGroupsNamespaceAPIClient
}

func (s *GRPCSuiteUpsetNS) BeforeAll(t provider.T) {
	s.TestStruct.NamespaceAPI = grpcClient.NewSGroupsNamespaceAPIClient(connection.GRPC(t))
}

func (s *GRPCSuiteUpsetNS) TestUpset(t provider.T) {
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

func (s *GRPCSuiteUpsetNS) TestUpsetErrorCheck(t provider.T) {
	t.Parallel()
	t.Story("Проверки ручки list для Namespace")
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		resp, err := s.TestStruct.FoundationTests(bodies.CheckListError())
		sCtx.NewStep(steps.AssertStep)
		sCtx.Require().Equal(resp.Namespaces[0].Metadata.Name, 123)
		sCtx.Require().NoError(err)
		sCtx.Require().NotNil(resp)
		sCtx.Assert().NotEmpty(resp.Namespaces[0].Metadata.Uid)
	})

}
func (s *GRPCSuiteUpsetNS) TestUpsetErrorCheckEmpty(t provider.T) {
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
func (s *GRPCSuiteUpsetNS) TestUpsetErrorCheck2(t provider.T) {
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
func (s *GRPCSuiteUpsetNS) TestUpsetErrorCheck3(t provider.T) {
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
func (s *GRPCSuiteUpsetNS) TestUpsetErrorCheck4(t provider.T) {
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
func (s *GRPCSuiteUpsetNS) TestUpsetErrorCheck5(t provider.T) {
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
func (s *GRPCSuiteUpsetNS) TestUpsetErrorCheck6(t provider.T) {
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
func (s *GRPCSuiteUpsetNS) TestUpsetErrorCheck7(t provider.T) {
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
func (s *GRPCSuiteUpsetNS) TestUpsetErrorCheck8(t provider.T) {
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
func (s *GRPCSuiteUpsetNS) TestUpsetErrorCheck9(t provider.T) {
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
func (s *GRPCSuiteUpsetNS) TestUpsetErrorCheck10(t provider.T) {
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
func (s *GRPCSuiteUpsetNS) TestUpsetErrorCheck11(t provider.T) {
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
func (s *GRPCSuiteUpsetNS) TestUpsetErrorCheck12(t provider.T) {
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
func (s *GRPCSuiteUpsetNS) TestUpsetErrorCheck13(t provider.T) {
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
func (s *GRPCSuiteUpsetNS) TestUpsetErrorCheck14(t provider.T) {
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
func (s *GRPCSuiteUpsetNS) TestUpsetErrorCheck15(t provider.T) {
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
func (s *GRPCSuiteUpsetNS) TestUpsetErrorCheck16(t provider.T) {
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
func (s *GRPCSuiteUpsetNS) TestUpsetErrorCheck17(t provider.T) {
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
func (s *GRPCSuiteUpsetNS) TestUpsetErrorCheck18(t provider.T) {
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
func (s *GRPCSuiteUpsetNS) TestUpsetErrorCheck19(t provider.T) {
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
