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

func (s *GRPCSuiteDeleteNS) TestDeleteByName(t provider.T) {
	t.Parallel()
	t.Story("Удаление по name")
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		resp, err := s.TestStruct.FoundationTestsDelete(bodies.DeleteByName())

		sCtx.NewStep(steps.AssertStep)
		sCtx.Require().NoError(err)
		sCtx.Assert().NotNil(resp)
	})
}

func (s *GRPCSuiteDeleteNS) TestDeleteByUID(t provider.T) {
	t.Parallel()
	t.Story("Удаление по uid")
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		resp, err := s.TestStruct.FoundationTestsDelete(bodies.DeleteByUID())

		sCtx.NewStep(steps.AssertStep)
		sCtx.Require().NoError(err)
		sCtx.Assert().NotNil(resp)
	})
}

func (s *GRPCSuiteDeleteNS) TestDeleteErrorUIDNotExist(t provider.T) {
	t.Parallel()
	t.Story("Ошибка при удалении по uid, namespace с таким uid не существует")
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		_, err := s.TestStruct.FoundationTestsDelete(bodies.DeleteErrorUIDNotExist())

		sCtx.NewStep(steps.AssertStep)
		utils.AssertGRPCError(sCtx, err, codes.NotFound, "for delete you must pass existing uid OR matching name")
	})
}

func (s *GRPCSuiteDeleteNS) TestDeleteErrorUIDEmpty(t provider.T) {
	t.Parallel()
	t.Story("Ошибка при удалении по uid, uid пуст")
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		_, err := s.TestStruct.FoundationTestsDelete(bodies.DeleteErrorUIDEmpty())

		sCtx.NewStep(steps.AssertStep)
		utils.AssertGRPCError(sCtx, err, codes.InvalidArgument, "at least one of UID or Name must be set")
	})
}

func (s *GRPCSuiteDeleteNS) TestDeleteErrorUIDShort(t provider.T) {
	t.Parallel()
	t.Story("Ошибка при удалении по uid, uid короткий")
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		_, err := s.TestStruct.FoundationTestsDelete(bodies.DeleteErrorUIDShort())

		sCtx.NewStep(steps.AssertStep)
		utils.AssertGRPCError(sCtx, err, codes.InvalidArgument, "invalid UUID length")
	})
}

func (s *GRPCSuiteDeleteNS) TestDeleteErrorUIDLong(t provider.T) {
	t.Parallel()
	t.Story("Ошибка при удалении по uid, uid длинный")
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		_, err := s.TestStruct.FoundationTestsDelete(bodies.DeleteErrorUIDLong())

		sCtx.NewStep(steps.AssertStep)
		utils.AssertGRPCError(sCtx, err, codes.InvalidArgument, "invalid UUID length")
	})
}

func (s *GRPCSuiteDeleteNS) TestDeleteErrorNameNotExist(t provider.T) {
	t.Parallel()
	t.Story("Ошибка при удалении по name, namespace с таким name не существует")
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		_, err := s.TestStruct.FoundationTestsDelete(bodies.DeleteErrorNameNotExist())

		sCtx.NewStep(steps.AssertStep)
		utils.AssertGRPCError(sCtx, err, codes.NotFound, "for delete you must pass existing uid OR matching name")
	})
}

func (s *GRPCSuiteDeleteNS) TestDeleteErrorNameEmpty(t provider.T) {
	t.Parallel()
	t.Story("Ошибка при удалении по name, name пуст")
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		_, err := s.TestStruct.FoundationTestsDelete(bodies.DeleteErrorNameEmpty())

		sCtx.NewStep(steps.AssertStep)
		utils.AssertGRPCError(sCtx, err, codes.InvalidArgument, "at least one of UID or Name must be set")
	})
}

func (s *GRPCSuiteDeleteNS) TestDeleteErrorNameLeadingSpace(t provider.T) {
	t.Parallel()
	t.Story("Ошибка при удалении по name, пробел в начале name")
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		_, err := s.TestStruct.FoundationTestsDelete(bodies.DeleteErrorNameLeadingSpace())

		sCtx.NewStep(steps.AssertStep)
		utils.AssertGRPCError(sCtx, err, codes.InvalidArgument, "Name: must consist of lower-case alphanumeric characters")
	})
}

func (s *GRPCSuiteDeleteNS) TestDeleteErrorNameTrailingSpace(t provider.T) {
	t.Parallel()
	t.Story("Ошибка при удалении по name, пробел в конце name")
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		_, err := s.TestStruct.FoundationTestsDelete(bodies.DeleteErrorNameTrailingSpace())

		sCtx.NewStep(steps.AssertStep)
		utils.AssertGRPCError(sCtx, err, codes.InvalidArgument, "Name: must consist of lower-case alphanumeric characters")
	})
}

func (s *GRPCSuiteDeleteNS) TestDeleteErrorNameWithSpaces(t provider.T) {
	t.Parallel()
	t.Story("Ошибка при удалении по name, name с пробелами")
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		_, err := s.TestStruct.FoundationTestsDelete(bodies.DeleteErrorNameWithSpaces())

		sCtx.NewStep(steps.AssertStep)
		utils.AssertGRPCError(sCtx, err, codes.InvalidArgument, "Name: must consist of lower-case alphanumeric characters")
	})
}

func (s *GRPCSuiteDeleteNS) TestDeleteErrorMetadataEmpty(t provider.T) {
	t.Parallel()
	t.Story("Ошибка при удалении, metadata пустая")
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		_, err := s.TestStruct.FoundationTestsDelete(bodies.DeleteErrorMetadataEmpty())

		sCtx.NewStep(steps.AssertStep)
		utils.AssertGRPCError(sCtx, err, codes.InvalidArgument, "at least one of UID or Name must be set")
	})
}

func (s *GRPCSuiteDeleteNS) TestDeleteErrorMetadataMissing(t provider.T) {
	t.Parallel()
	t.Story("Ошибка при удалении, metadata отсутствует как параметр")
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		_, err := s.TestStruct.FoundationTestsDelete(bodies.DeleteErrorMetadataMissing())

		sCtx.NewStep(steps.AssertStep)
		utils.AssertGRPCError(sCtx, err, codes.InvalidArgument, "at least one of UID or Name must be set")
	})
}
