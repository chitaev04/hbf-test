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
	"HBF-tests/tests/projects/server/nb/internal/bodies"
)

type GRPCSuiteDeleteNB struct {
	suite.Suite
	TestStruct foundation.ToolsForTest
	client     grpcClient.SGroupsNetworkBindingAPIClient
}

func (s *GRPCSuiteDeleteNB) BeforeAll(t provider.T) {
	s.TestStruct.NetworkBindingAPI = grpcClient.NewSGroupsNetworkBindingAPIClient(connection.GRPC(t))
}

func (s *GRPCSuiteDeleteNB) runSuccess(t provider.T, story string, body *foundation.TestCaseBodyNBDelete) {
	t.Parallel()
	t.Story(story)
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		resp, err := s.TestStruct.FoundationTestsNBDelete(body)

		sCtx.NewStep(steps.AssertStep)
		sCtx.Require().NoError(err)
		sCtx.Assert().NotNil(resp)
	})
}

func (s *GRPCSuiteDeleteNB) runError(t provider.T, story string, body *foundation.TestCaseBodyNBDelete, code codes.Code, substrings ...string) {
	t.Parallel()
	t.Story(story)
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		_, err := s.TestStruct.FoundationTestsNBDelete(body)

		sCtx.NewStep(steps.AssertStep)
		utils.AssertGRPCError(sCtx, err, code, substrings...)
	})
}

func (s *GRPCSuiteDeleteNB) TestDeleteByNameNamespace(t provider.T) {
	s.runSuccess(t, "Удаление по name+namespace", bodies.DeleteByNameNamespace())
}

func (s *GRPCSuiteDeleteNB) TestDeleteByUID(t provider.T) {
	s.runSuccess(t, "Удаление по uid", bodies.DeleteByUID())
}

func (s *GRPCSuiteDeleteNB) TestDeleteErrorNameEmpty(t provider.T) {
	s.runError(t, "Ошибка при удалении по name+namespace, name пуст", bodies.DeleteErrorNameEmpty(), codes.InvalidArgument, "at least one of UID or Name and Namespace must be set")
}

func (s *GRPCSuiteDeleteNB) TestDeleteErrorNameLeadingSpace(t provider.T) {
	s.runError(t, "Ошибка при удалении по name+namespace, пробел в начале name", bodies.DeleteErrorNameLeadingSpace(), codes.InvalidArgument, "must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteDeleteNB) TestDeleteErrorNameTrailingSpace(t provider.T) {
	s.runError(t, "Ошибка при удалении по name+namespace, пробел в конце name", bodies.DeleteErrorNameTrailingSpace(), codes.InvalidArgument, "must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteDeleteNB) TestDeleteErrorNameWithSpaces(t provider.T) {
	s.runError(t, "Ошибка при удалении по name+namespace, name с пробелами", bodies.DeleteErrorNameWithSpaces(), codes.InvalidArgument, "must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteDeleteNB) TestDeleteErrorNameNotExist(t provider.T) {
	s.runError(t, "Ошибка при удалении по name+namespace, name не существует", bodies.DeleteErrorNameNotExist(), codes.NotFound, "for delete pass existing")
}

func (s *GRPCSuiteDeleteNB) TestDeleteErrorNamespaceEmpty(t provider.T) {
	s.runError(t, "Ошибка при удалении по name+namespace, namespace пуст", bodies.DeleteErrorNamespaceEmpty(), codes.InvalidArgument, "at least one of UID or Name and Namespace must be set")
}

func (s *GRPCSuiteDeleteNB) TestDeleteErrorNamespaceLeadingSpace(t provider.T) {
	s.runError(t, "Ошибка при удалении по name+namespace, пробел в начале имени namespace", bodies.DeleteErrorNamespaceLeadingSpace(), codes.InvalidArgument, "must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteDeleteNB) TestDeleteErrorNamespaceTrailingSpace(t provider.T) {
	s.runError(t, "Ошибка при удалении по name+namespace, пробел в конце имени namespace", bodies.DeleteErrorNamespaceTrailingSpace(), codes.InvalidArgument, "must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteDeleteNB) TestDeleteErrorNamespaceWithSpaces(t provider.T) {
	s.runError(t, "Ошибка при удалении по name+namespace, имя namespace с пробелами", bodies.DeleteErrorNamespaceWithSpaces(), codes.InvalidArgument, "must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteDeleteNB) TestDeleteErrorNamespaceNotExist(t provider.T) {
	s.runError(t, "Ошибка при удалении по name+namespace, namespace не существует", bodies.DeleteErrorNamespaceNotExist(), codes.NotFound, "pass existing namespace name")
}

func (s *GRPCSuiteDeleteNB) TestDeleteErrorNamespaceMissing(t provider.T) {
	s.runError(t, "Ошибка при удалении по name+namespace, namespace отсутствует как параметр", bodies.DeleteErrorNamespaceMissing(), codes.InvalidArgument, "at least one of UID or Name and Namespace must be set")
}

func (s *GRPCSuiteDeleteNB) TestDeleteErrorNameNotExistInNamespace(t provider.T) {
	s.runError(t, "Ошибка при удалении по name+namespace, name не существует в указанном namespace", bodies.DeleteErrorNameNotExistInNamespace(), codes.NotFound, "for delete pass existing")
}

func (s *GRPCSuiteDeleteNB) TestDeleteErrorNameMissing(t provider.T) {
	s.runError(t, "Ошибка при удалении по name+namespace, name отсутствует как параметр", bodies.DeleteErrorNameMissing(), codes.InvalidArgument, "at least one of UID or Name and Namespace must be set")
}

func (s *GRPCSuiteDeleteNB) TestDeleteErrorUIDEmpty(t provider.T) {
	s.runError(t, "Ошибка при удалении по uid, пустой uid", bodies.DeleteErrorUIDEmpty(), codes.InvalidArgument, "at least one of UID or Name and Namespace must be set")
}

func (s *GRPCSuiteDeleteNB) TestDeleteErrorUIDNotExist(t provider.T) {
	s.runError(t, "Ошибка при удалении по uid, uid не существует", bodies.DeleteErrorUIDNotExist(), codes.NotFound, "for delete pass existing uid OR (name AND namespace)")
}

func (s *GRPCSuiteDeleteNB) TestDeleteErrorUIDShort(t provider.T) {
	s.runError(t, "Ошибка при удалении по uid, короткий uid", bodies.DeleteErrorUIDShort(), codes.InvalidArgument, "invalid UUID length")
}

func (s *GRPCSuiteDeleteNB) TestDeleteErrorUIDLong(t provider.T) {
	s.runError(t, "Ошибка при удалении по uid, длинный uid", bodies.DeleteErrorUIDLong(), codes.InvalidArgument, "invalid UUID length")
}

func (s *GRPCSuiteDeleteNB) TestDeleteErrorMetadataEmpty(t provider.T) {
	s.runError(t, "Ошибка при удалении, пустая метадата", bodies.DeleteErrorMetadataEmpty(), codes.InvalidArgument, "at least one of UID or Name and Namespace must be set")
}

func (s *GRPCSuiteDeleteNB) TestDeleteErrorMetadataMissing(t provider.T) {
	s.runError(t, "Ошибка при удалении, метадата отсутствует как параметр", bodies.DeleteErrorMetadataMissing(), codes.InvalidArgument, "at least one of UID or Name and Namespace must be set")
}
