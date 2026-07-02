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

type GRPCSuiteUpsertNS struct {
	suite.Suite
	TestStruct foundation.ToolsForTest
	client     grpcClient.SGroupsNamespaceAPIClient
}

func (s *GRPCSuiteUpsertNS) BeforeAll(t provider.T) {
	s.TestStruct.NamespaceAPI = grpcClient.NewSGroupsNamespaceAPIClient(connection.GRPC(t))
}

func (s *GRPCSuiteUpsertNS) TestUpsertAddOne(t provider.T) {
	t.Parallel()
	t.Story("Добавление одного нового namespace")
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		resp, err := s.TestStruct.FoundationTestsUpsert(bodies.UpsertAddOne())

		sCtx.NewStep(steps.AssertStep)
		sCtx.Require().NoError(err)
		sCtx.Require().NotNil(resp)
		sCtx.Require().Len(resp.Namespaces, 1)
		assertNamespace(sCtx, resp.Namespaces[0], "add-namespace-1", namespaceExpectation{
			labels:      map[string]string{"add": "success"},
			annotations: map[string]string{"add": "success"},
			displayName: "New Namespace",
			comment:     "for upsert tests",
			description: "add new namespace",
		})
	})
}

func (s *GRPCSuiteUpsertNS) TestUpsertAddTwo(t provider.T) {
	t.Parallel()
	t.Story("Добавление 2 новых namespaces")
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		resp, err := s.TestStruct.FoundationTestsUpsert(bodies.UpsertAddTwo())

		sCtx.NewStep(steps.AssertStep)
		sCtx.Require().NoError(err)
		sCtx.Require().NotNil(resp)
		sCtx.Require().Len(resp.Namespaces, 2)
		assertNamespacesByName(sCtx, resp.Namespaces, map[string]namespaceExpectation{
			"add-namespace-2": {labels: map[string]string{}, annotations: map[string]string{"add": "success"}},
			"add-namespace-3": {
				labels:      map[string]string{"add": "success"},
				annotations: map[string]string{},
				displayName: "New Namespace 3",
				comment:     "for upsert tests",
				description: "add new namespace 2",
			},
		})
	})
}

func (s *GRPCSuiteUpsertNS) TestUpsertEditByUID(t provider.T) {
	t.Parallel()
	t.Story("Редактирование namespace по uid")
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		resp, err := s.TestStruct.FoundationTestsUpsert(bodies.UpsertEditByUID())

		sCtx.NewStep(steps.AssertStep)
		sCtx.Require().NoError(err)
		sCtx.Require().NotNil(resp)
		sCtx.Require().Len(resp.Namespaces, 1)
		assertNamespace(sCtx, resp.Namespaces[0], "namespace-3", namespaceExpectation{
			labels:      map[string]string{"edit": "success"},
			annotations: map[string]string{"edit": "success"},
			displayName: "edit success",
			comment:     "for upsert tests",
			description: "edit namespace",
		})
		sCtx.Assert().Equal("cd3e3c34-bf87-4787-87e8-7ba7182280c3", resp.Namespaces[0].Metadata.Uid)
	})
}

func (s *GRPCSuiteUpsertNS) TestUpsertErrorDuplicateNameWithoutUID(t provider.T) {
	t.Parallel()
	t.Story("Ошибка при добавлении, namespace с таким именем уже существует (не указан uid)")
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		_, err := s.TestStruct.FoundationTestsUpsert(bodies.UpsertErrorDuplicateNameWithoutUID())

		sCtx.NewStep(steps.AssertStep)
		utils.AssertGRPCError(sCtx, err, codes.FailedPrecondition, "for insert pass unique name")
	})
}

func (s *GRPCSuiteUpsertNS) TestUpsertErrorEmptyName(t provider.T) {
	t.Parallel()
	t.Story("Ошибка при добавлении, name пустая строка")
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		_, err := s.TestStruct.FoundationTestsUpsert(bodies.UpsertErrorEmptyName())

		sCtx.NewStep(steps.AssertStep)
		utils.AssertGRPCError(sCtx, err, codes.InvalidArgument, "at least one of UID or Name must be set")
	})
}

func (s *GRPCSuiteUpsertNS) TestUpsertErrorNameMissing(t provider.T) {
	t.Parallel()
	t.Story("Ошибка при добавлении, name отсутствует как параметр")
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		_, err := s.TestStruct.FoundationTestsUpsert(bodies.UpsertErrorNameMissing())

		sCtx.NewStep(steps.AssertStep)
		utils.AssertGRPCError(sCtx, err, codes.InvalidArgument, "at least one of UID or Name must be set")
	})
}

func (s *GRPCSuiteUpsertNS) TestUpsertErrorNameCyrillic(t provider.T) {
	t.Parallel()
	t.Story("Ошибка при добавлении, name на кириллице")
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		_, err := s.TestStruct.FoundationTestsUpsert(bodies.UpsertErrorNameCyrillic())

		sCtx.NewStep(steps.AssertStep)
		utils.AssertGRPCError(sCtx, err, codes.InvalidArgument,
			"must start and end with an alphanumeric character",
			"Name: must consist of lower-case alphanumeric characters")
	})
}

func (s *GRPCSuiteUpsertNS) TestUpsertErrorNameCyrillicDigits(t provider.T) {
	t.Parallel()
	t.Story("Ошибка при добавлении, name на кириллице+цифры")
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		_, err := s.TestStruct.FoundationTestsUpsert(bodies.UpsertErrorNameCyrillicDigits())

		sCtx.NewStep(steps.AssertStep)
		utils.AssertGRPCError(sCtx, err, codes.InvalidArgument,
			"must start and end with an alphanumeric character",
			"Name: must consist of lower-case alphanumeric characters")
	})
}

func (s *GRPCSuiteUpsertNS) TestUpsertErrorNameCyrillicSpecialChars(t provider.T) {
	t.Parallel()
	t.Story("Ошибка при добавлении, name на кириллице+спец.символы")
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		_, err := s.TestStruct.FoundationTestsUpsert(bodies.UpsertErrorNameCyrillicSpecialChars())

		sCtx.NewStep(steps.AssertStep)
		utils.AssertGRPCError(sCtx, err, codes.InvalidArgument,
			"must start and end with an alphanumeric character",
			"Name: must consist of lower-case alphanumeric characters")
	})
}

func (s *GRPCSuiteUpsertNS) TestUpsertErrorNameLatinSpecialChars(t provider.T) {
	t.Parallel()
	t.Story("Ошибка при добавлении, name на латинице+спец.символы")
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		_, err := s.TestStruct.FoundationTestsUpsert(bodies.UpsertErrorNameLatinSpecialChars())

		sCtx.NewStep(steps.AssertStep)
		utils.AssertGRPCError(sCtx, err, codes.InvalidArgument,
			"must start and end with an alphanumeric character",
			"Name: must consist of lower-case alphanumeric characters")
	})
}

func (s *GRPCSuiteUpsertNS) TestUpsertErrorNameUpperCase(t provider.T) {
	t.Parallel()
	t.Story("Ошибка при добавлении, name капсом")
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		_, err := s.TestStruct.FoundationTestsUpsert(bodies.UpsertErrorNameUpperCase())

		sCtx.NewStep(steps.AssertStep)
		utils.AssertGRPCError(sCtx, err, codes.InvalidArgument,
			"must start and end with an alphanumeric character",
			"Name: must consist of lower-case alphanumeric characters")
	})
}

func (s *GRPCSuiteUpsertNS) TestUpsertErrorNameStartsWithHyphen(t provider.T) {
	t.Parallel()
	t.Story("Ошибка при добавлении, name начинается с дефиса")
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		_, err := s.TestStruct.FoundationTestsUpsert(bodies.UpsertErrorNameStartsWithHyphen())

		sCtx.NewStep(steps.AssertStep)
		utils.AssertGRPCError(sCtx, err, codes.InvalidArgument,
			"must start and end with an alphanumeric character",
			"Name: must consist of lower-case alphanumeric characters")
	})
}

func (s *GRPCSuiteUpsertNS) TestUpsertErrorNameEndsWithHyphen(t provider.T) {
	t.Parallel()
	t.Story("Ошибка при добавлении, name заканчивается дефисом")
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		_, err := s.TestStruct.FoundationTestsUpsert(bodies.UpsertErrorNameEndsWithHyphen())

		sCtx.NewStep(steps.AssertStep)
		utils.AssertGRPCError(sCtx, err, codes.InvalidArgument,
			"must start and end with an alphanumeric character",
			"Name: must consist of lower-case alphanumeric characters")
	})
}

func (s *GRPCSuiteUpsertNS) TestUpsertErrorNameSpecialCharsOnly(t provider.T) {
	t.Parallel()
	t.Story("Ошибка при добавлении, name из спец.символов")
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		_, err := s.TestStruct.FoundationTestsUpsert(bodies.UpsertErrorNameSpecialCharsOnly())

		sCtx.NewStep(steps.AssertStep)
		utils.AssertGRPCError(sCtx, err, codes.InvalidArgument,
			"must start and end with an alphanumeric character",
			"Name: must consist of lower-case alphanumeric characters")
	})
}

func (s *GRPCSuiteUpsertNS) TestUpsertErrorNameLeadingSpace(t provider.T) {
	t.Parallel()
	t.Story("Ошибка при добавлении, пробел в начале name")
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		_, err := s.TestStruct.FoundationTestsUpsert(bodies.UpsertErrorNameLeadingSpace())

		sCtx.NewStep(steps.AssertStep)
		utils.AssertGRPCError(sCtx, err, codes.InvalidArgument,
			"must start and end with an alphanumeric character",
			"Name: must consist of lower-case alphanumeric characters")
	})
}

func (s *GRPCSuiteUpsertNS) TestUpsertErrorNameTrailingSpace(t provider.T) {
	t.Parallel()
	t.Story("Ошибка при добавлении, пробел в конце name")
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		_, err := s.TestStruct.FoundationTestsUpsert(bodies.UpsertErrorNameTrailingSpace())

		sCtx.NewStep(steps.AssertStep)
		utils.AssertGRPCError(sCtx, err, codes.InvalidArgument,
			"must start and end with an alphanumeric character",
			"Name: must consist of lower-case alphanumeric characters")
	})
}

func (s *GRPCSuiteUpsertNS) TestUpsertErrorNameWithSpaces(t provider.T) {
	t.Parallel()
	t.Story("Ошибка при добавлении, name с пробелами")
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		_, err := s.TestStruct.FoundationTestsUpsert(bodies.UpsertErrorNameWithSpaces())

		sCtx.NewStep(steps.AssertStep)
		utils.AssertGRPCError(sCtx, err, codes.InvalidArgument,
			"must start and end with an alphanumeric character",
			"Name: must consist of lower-case alphanumeric characters")
	})
}

func (s *GRPCSuiteUpsertNS) TestUpsertErrorNameTooLong(t provider.T) {
	t.Parallel()
	t.Story("Ошибка при добавлении, name > 63 символов")
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		_, err := s.TestStruct.FoundationTestsUpsert(bodies.UpsertErrorNameTooLong())

		sCtx.NewStep(steps.AssertStep)
		utils.AssertGRPCError(sCtx, err, codes.InvalidArgument, "be no longer than 63 characters")
	})
}

func (s *GRPCSuiteUpsertNS) TestUpsertErrorMetadataEmpty(t provider.T) {
	t.Parallel()
	t.Story("Ошибка при добавлении, metadata пустой объект")
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		_, err := s.TestStruct.FoundationTestsUpsert(bodies.UpsertErrorMetadataEmpty())

		sCtx.NewStep(steps.AssertStep)
		utils.AssertGRPCError(sCtx, err, codes.InvalidArgument, "at least one of UID or Name must be set")
	})
}

func (s *GRPCSuiteUpsertNS) TestUpsertErrorMetadataMissing(t provider.T) {
	t.Parallel()
	t.Story("Ошибка при добавлении, metadata отсутствует как параметр")
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		_, err := s.TestStruct.FoundationTestsUpsert(bodies.UpsertErrorMetadataMissing())

		sCtx.NewStep(steps.AssertStep)
		utils.AssertGRPCError(sCtx, err, codes.InvalidArgument, "at least one of UID or Name must be set")
	})
}

func (s *GRPCSuiteUpsertNS) TestUpsertErrorDisplayNameTooLong(t provider.T) {
	t.Parallel()
	t.Story("Ошибка при добавлении, displayName > 63 символов")
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		_, err := s.TestStruct.FoundationTestsUpsert(bodies.UpsertErrorDisplayNameTooLong())

		sCtx.NewStep(steps.AssertStep)
		utils.AssertGRPCError(sCtx, err, codes.InvalidArgument, "display name must be no longer than 63 characters")
	})
}

func (s *GRPCSuiteUpsertNS) TestUpsertErrorEditNonExistentUID(t provider.T) {
	t.Parallel()
	t.Story("Ошибка при редактировании, namespace с таким uid не существует")
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		_, err := s.TestStruct.FoundationTestsUpsert(bodies.UpsertErrorEditNonExistentUID())

		sCtx.NewStep(steps.AssertStep)
		utils.AssertGRPCError(sCtx, err, codes.FailedPrecondition, "for update you must pass existing uid AND matching name")
	})
}

func (s *GRPCSuiteUpsertNS) TestUpsertErrorEditShortUID(t provider.T) {
	t.Parallel()
	t.Story("Ошибка при редактировании, короткий uid")
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		_, err := s.TestStruct.FoundationTestsUpsert(bodies.UpsertErrorEditShortUID())

		sCtx.NewStep(steps.AssertStep)
		utils.AssertGRPCError(sCtx, err, codes.InvalidArgument, "invalid UUID length")
	})
}

func (s *GRPCSuiteUpsertNS) TestUpsertErrorEditLongUID(t provider.T) {
	t.Parallel()
	t.Story("Ошибка при редактировании, длинный uid")
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		_, err := s.TestStruct.FoundationTestsUpsert(bodies.UpsertErrorEditLongUID())

		sCtx.NewStep(steps.AssertStep)
		utils.AssertGRPCError(sCtx, err, codes.InvalidArgument, "invalid UUID length")
	})
}
