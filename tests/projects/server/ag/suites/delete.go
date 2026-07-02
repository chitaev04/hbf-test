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

type GRPCSuiteDeleteAG struct {
	suite.Suite
	TestStruct foundation.ToolsForTest
	client     grpcClient.SGroupsAddressGroupsAPIClient
}

func (s *GRPCSuiteDeleteAG) BeforeAll(t provider.T) {
	s.TestStruct.AddressGroupAPI = grpcClient.NewSGroupsAddressGroupsAPIClient(connection.GRPC(t))
}

func (s *GRPCSuiteDeleteAG) runSuccess(t provider.T, story string, body *foundation.TestCaseBodyAGDelete) {
	t.Parallel()
	t.Story(story)
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		resp, err := s.TestStruct.FoundationTestsAGDelete(body)

		sCtx.NewStep(steps.AssertStep)
		sCtx.Require().NoError(err)
		sCtx.Assert().NotNil(resp)
	})
}

func (s *GRPCSuiteDeleteAG) runError(t provider.T, story string, body *foundation.TestCaseBodyAGDelete, code codes.Code, substrings ...string) {
	t.Parallel()
	t.Story(story)
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		_, err := s.TestStruct.FoundationTestsAGDelete(body)

		sCtx.NewStep(steps.AssertStep)
		utils.AssertGRPCError(sCtx, err, code, substrings...)
	})
}

func (s *GRPCSuiteDeleteAG) TestDeleteByNameNamespace(t provider.T) {
	s.runSuccess(t, "Удаление по name+namespace", bodies.DeleteByNameNamespace())
}

func (s *GRPCSuiteDeleteAG) TestDeleteByUID(t provider.T) {
	s.runSuccess(t, "Удаление по uid", bodies.DeleteByUID())
}

func (s *GRPCSuiteDeleteAG) TestDeleteErrorNameEmpty(t provider.T) {
	s.runError(t, "Ошибка при удалении по name+namespace, name пуст", bodies.DeleteErrorNameEmpty(), codes.InvalidArgument, "at least one of UID or Name and Namespace must be set")
}

func (s *GRPCSuiteDeleteAG) TestDeleteErrorNameLeadingSpace(t provider.T) {
	s.runError(t, "Ошибка при удалении по name+namespace, пробел в начале name", bodies.DeleteErrorNameLeadingSpace(), codes.InvalidArgument, "must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteDeleteAG) TestDeleteErrorNameTrailingSpace(t provider.T) {
	s.runError(t, "Ошибка при удалении по name+namespace, пробел в конце name", bodies.DeleteErrorNameTrailingSpace(), codes.InvalidArgument, "must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteDeleteAG) TestDeleteErrorNameWithSpaces(t provider.T) {
	s.runError(t, "Ошибка при удалении по name+namespace, name с пробелами", bodies.DeleteErrorNameWithSpaces(), codes.InvalidArgument, "must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteDeleteAG) TestDeleteErrorNameNotExist(t provider.T) {
	s.runError(t, "Ошибка при удалении по name+namespace, name не существует", bodies.DeleteErrorNameNotExist(), codes.NotFound, "for delete pass existing")
}

func (s *GRPCSuiteDeleteAG) TestDeleteErrorNamespaceEmpty(t provider.T) {
	s.runError(t, "Ошибка при удалении по name+namespace, namespace пуст", bodies.DeleteErrorNamespaceEmpty(), codes.InvalidArgument, "at least one of UID or Name and Namespace must be set")
}

func (s *GRPCSuiteDeleteAG) TestDeleteErrorNamespaceLeadingSpace(t provider.T) {
	s.runError(t, "Ошибка при удалении по name+namespace, пробел в начале имени namespace", bodies.DeleteErrorNamespaceLeadingSpace(), codes.InvalidArgument, "must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteDeleteAG) TestDeleteErrorNamespaceTrailingSpace(t provider.T) {
	s.runError(t, "Ошибка при удалении по name+namespace, пробел в конце имени namespace", bodies.DeleteErrorNamespaceTrailingSpace(), codes.InvalidArgument, "must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteDeleteAG) TestDeleteErrorNamespaceWithSpaces(t provider.T) {
	s.runError(t, "Ошибка при удалении по name+namespace, имя namespace с пробелами", bodies.DeleteErrorNamespaceWithSpaces(), codes.InvalidArgument, "must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteDeleteAG) TestDeleteErrorNamespaceNotExist(t provider.T) {
	s.runError(t, "Ошибка при удалении по name+namespace, namespace не существует", bodies.DeleteErrorNamespaceNotExist(), codes.NotFound, "pass existing namespace name")
}

func (s *GRPCSuiteDeleteAG) TestDeleteErrorNamespaceMissing(t provider.T) {
	s.runError(t, "Ошибка при удалении по name+namespace, namespace отсутствует как параметр", bodies.DeleteErrorNamespaceMissing(), codes.InvalidArgument, "at least one of UID or Name and Namespace must be set")
}

func (s *GRPCSuiteDeleteAG) TestDeleteErrorNameNotExistInNamespace(t provider.T) {
	s.runError(t, "Ошибка при удалении по name+namespace, name не существует в указанном namespace", bodies.DeleteErrorNameNotExistInNamespace(), codes.NotFound, "for delete pass existing")
}

func (s *GRPCSuiteDeleteAG) TestDeleteErrorNameMissing(t provider.T) {
	s.runError(t, "Ошибка при удалении по name+namespace, name отсутствует как параметр", bodies.DeleteErrorNameMissing(), codes.InvalidArgument, "at least one of UID or Name and Namespace must be set")
}

func (s *GRPCSuiteDeleteAG) TestDeleteErrorUIDEmpty(t provider.T) {
	s.runError(t, "Ошибка при удалении по uid, пустой uid", bodies.DeleteErrorUIDEmpty(), codes.InvalidArgument, "at least one of UID or Name and Namespace must be set")
}

func (s *GRPCSuiteDeleteAG) TestDeleteErrorUIDNotExist(t provider.T) {
	s.runError(t, "Ошибка при удалении по uid, uid не существует", bodies.DeleteErrorUIDNotExist(), codes.NotFound, "for delete pass existing uid OR (name AND namespace)")
}

func (s *GRPCSuiteDeleteAG) TestDeleteErrorUIDShort(t provider.T) {
	s.runError(t, "Ошибка при удалении по uid, короткий uid", bodies.DeleteErrorUIDShort(), codes.InvalidArgument, "invalid UUID length")
}

func (s *GRPCSuiteDeleteAG) TestDeleteErrorUIDLong(t provider.T) {
	s.runError(t, "Ошибка при удалении по uid, длинный uid", bodies.DeleteErrorUIDLong(), codes.InvalidArgument, "invalid UUID length")
}

func (s *GRPCSuiteDeleteAG) TestDeleteErrorMetadataEmpty(t provider.T) {
	s.runError(t, "Ошибка при удалении, пустая метадата", bodies.DeleteErrorMetadataEmpty(), codes.InvalidArgument, "at least one of UID or Name and Namespace must be set")
}

func (s *GRPCSuiteDeleteAG) TestDeleteErrorMetadataMissing(t provider.T) {
	s.runError(t, "Ошибка при удалении, метадата отсутствует как параметр", bodies.DeleteErrorMetadataMissing(), codes.InvalidArgument, "at least one of UID or Name and Namespace must be set")
}
