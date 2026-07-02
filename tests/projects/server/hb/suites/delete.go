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
	"HBF-tests/tests/projects/server/hb/internal/bodies"
)

type GRPCSuiteDeleteHB struct {
	suite.Suite
	TestStruct foundation.ToolsForTest
	client     grpcClient.SGroupsHostBindingAPIClient
}

func (s *GRPCSuiteDeleteHB) BeforeAll(t provider.T) {
	s.TestStruct.HostBindingAPI = grpcClient.NewSGroupsHostBindingAPIClient(connection.GRPC(t))
}

func (s *GRPCSuiteDeleteHB) runSuccess(t provider.T, story string, body *foundation.TestCaseBodyHBDelete) {
	t.Parallel()
	t.Story(story)
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		resp, err := s.TestStruct.FoundationTestsHBDelete(body)

		sCtx.NewStep(steps.AssertStep)
		sCtx.Require().NoError(err)
		sCtx.Assert().NotNil(resp)
	})
}

func (s *GRPCSuiteDeleteHB) runError(t provider.T, story string, body *foundation.TestCaseBodyHBDelete, code codes.Code, substrings ...string) {
	t.Parallel()
	t.Story(story)
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		_, err := s.TestStruct.FoundationTestsHBDelete(body)

		sCtx.NewStep(steps.AssertStep)
		utils.AssertGRPCError(sCtx, err, code, substrings...)
	})
}

func (s *GRPCSuiteDeleteHB) TestDeleteByNameNamespace(t provider.T) {
	s.runSuccess(t, "Удаление по name+namespace", bodies.DeleteByNameNamespace())
}

func (s *GRPCSuiteDeleteHB) TestDeleteByUID(t provider.T) {
	s.runSuccess(t, "Удаление по uid", bodies.DeleteByUID())
}

func (s *GRPCSuiteDeleteHB) TestDeleteErrorNameEmpty(t provider.T) {
	s.runError(t, "Ошибка при удалении по name+namespace, name пуст", bodies.DeleteErrorNameEmpty(), codes.InvalidArgument, "at least one of UID or Name and Namespace must be set")
}

func (s *GRPCSuiteDeleteHB) TestDeleteErrorNameLeadingSpace(t provider.T) {
	s.runError(t, "Ошибка при удалении по name+namespace, пробел в начале name", bodies.DeleteErrorNameLeadingSpace(), codes.InvalidArgument, "must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteDeleteHB) TestDeleteErrorNameTrailingSpace(t provider.T) {
	s.runError(t, "Ошибка при удалении по name+namespace, пробел в конце name", bodies.DeleteErrorNameTrailingSpace(), codes.InvalidArgument, "must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteDeleteHB) TestDeleteErrorNameWithSpaces(t provider.T) {
	s.runError(t, "Ошибка при удалении по name+namespace, name с пробелами", bodies.DeleteErrorNameWithSpaces(), codes.InvalidArgument, "must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteDeleteHB) TestDeleteErrorNameNotExist(t provider.T) {
	s.runError(t, "Ошибка при удалении по name+namespace, name не существует", bodies.DeleteErrorNameNotExist(), codes.NotFound, "for delete pass existing")
}

func (s *GRPCSuiteDeleteHB) TestDeleteErrorNamespaceEmpty(t provider.T) {
	s.runError(t, "Ошибка при удалении по name+namespace, namespace пуст", bodies.DeleteErrorNamespaceEmpty(), codes.InvalidArgument, "at least one of UID or Name and Namespace must be set")
}

func (s *GRPCSuiteDeleteHB) TestDeleteErrorNamespaceLeadingSpace(t provider.T) {
	s.runError(t, "Ошибка при удалении по name+namespace, пробел в начале имени namespace", bodies.DeleteErrorNamespaceLeadingSpace(), codes.InvalidArgument, "must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteDeleteHB) TestDeleteErrorNamespaceTrailingSpace(t provider.T) {
	s.runError(t, "Ошибка при удалении по name+namespace, пробел в конце имени namespace", bodies.DeleteErrorNamespaceTrailingSpace(), codes.InvalidArgument, "must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteDeleteHB) TestDeleteErrorNamespaceWithSpaces(t provider.T) {
	s.runError(t, "Ошибка при удалении по name+namespace, имя namespace с пробелами", bodies.DeleteErrorNamespaceWithSpaces(), codes.InvalidArgument, "must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteDeleteHB) TestDeleteErrorNamespaceNotExist(t provider.T) {
	s.runError(t, "Ошибка при удалении по name+namespace, namespace не существует", bodies.DeleteErrorNamespaceNotExist(), codes.NotFound, "pass existing namespace name")
}

func (s *GRPCSuiteDeleteHB) TestDeleteErrorNamespaceMissing(t provider.T) {
	s.runError(t, "Ошибка при удалении по name+namespace, namespace отсутствует как параметр", bodies.DeleteErrorNamespaceMissing(), codes.InvalidArgument, "at least one of UID or Name and Namespace must be set")
}

func (s *GRPCSuiteDeleteHB) TestDeleteErrorNameNotExistInNamespace(t provider.T) {
	s.runError(t, "Ошибка при удалении по name+namespace, name не существует в указанном namespace", bodies.DeleteErrorNameNotExistInNamespace(), codes.NotFound, "for delete pass existing")
}

func (s *GRPCSuiteDeleteHB) TestDeleteErrorNameMissing(t provider.T) {
	s.runError(t, "Ошибка при удалении по name+namespace, name отсутствует как параметр", bodies.DeleteErrorNameMissing(), codes.InvalidArgument, "at least one of UID or Name and Namespace must be set")
}

func (s *GRPCSuiteDeleteHB) TestDeleteErrorUIDEmpty(t provider.T) {
	s.runError(t, "Ошибка при удалении по uid, пустой uid", bodies.DeleteErrorUIDEmpty(), codes.InvalidArgument, "at least one of UID or Name and Namespace must be set")
}

func (s *GRPCSuiteDeleteHB) TestDeleteErrorUIDNotExist(t provider.T) {
	s.runError(t, "Ошибка при удалении по uid, uid не существует", bodies.DeleteErrorUIDNotExist(), codes.NotFound, "for delete pass existing uid OR (name AND namespace)")
}

func (s *GRPCSuiteDeleteHB) TestDeleteErrorUIDShort(t provider.T) {
	s.runError(t, "Ошибка при удалении по uid, короткий uid", bodies.DeleteErrorUIDShort(), codes.InvalidArgument, "invalid UUID length")
}

func (s *GRPCSuiteDeleteHB) TestDeleteErrorUIDLong(t provider.T) {
	s.runError(t, "Ошибка при удалении по uid, длинный uid", bodies.DeleteErrorUIDLong(), codes.InvalidArgument, "invalid UUID length")
}

func (s *GRPCSuiteDeleteHB) TestDeleteErrorMetadataEmpty(t provider.T) {
	s.runError(t, "Ошибка при удалении, пустая метадата", bodies.DeleteErrorMetadataEmpty(), codes.InvalidArgument, "at least one of UID or Name and Namespace must be set")
}

func (s *GRPCSuiteDeleteHB) TestDeleteErrorMetadataMissing(t provider.T) {
	s.runError(t, "Ошибка при удалении, метадата отсутствует как параметр", bodies.DeleteErrorMetadataMissing(), codes.InvalidArgument, "at least one of UID or Name and Namespace must be set")
}
