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

type GRPCSuiteUpsertAG struct {
	suite.Suite
	TestStruct foundation.ToolsForTest
	client     grpcClient.SGroupsAddressGroupsAPIClient
}

func (s *GRPCSuiteUpsertAG) BeforeAll(t provider.T) {
	s.TestStruct.AddressGroupAPI = grpcClient.NewSGroupsAddressGroupsAPIClient(connection.GRPC(t))
}

func (s *GRPCSuiteUpsertAG) runSuccess(t provider.T, story string, body *foundation.TestCaseBodyAGUpsert, name string, exp addressGroupExpectation) {
	t.Parallel()
	t.Story(story)
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		resp, err := s.TestStruct.FoundationTestsAGUpsert(body)

		sCtx.NewStep(steps.AssertStep)
		sCtx.Require().NoError(err)
		sCtx.Require().NotNil(resp)
		sCtx.Require().Len(resp.AddressGroups, 1)
		assertAddressGroup(sCtx, resp.AddressGroups[0], name, exp)
	})
}

func (s *GRPCSuiteUpsertAG) runError(t provider.T, story string, body *foundation.TestCaseBodyAGUpsert, code codes.Code, substrings ...string) {
	t.Parallel()
	t.Story(story)
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		_, err := s.TestStruct.FoundationTestsAGUpsert(body)

		sCtx.NewStep(steps.AssertStep)
		utils.AssertGRPCError(sCtx, err, code, substrings...)
	})
}

func (s *GRPCSuiteUpsertAG) TestUpsertAdd(t provider.T) {
	s.runSuccess(t, "Добавление новой AG", bodies.UpsertAdd(), "add-ag-0", addressGroupExpectation{
		namespace: "namespace-0", labels: map[string]string{"add": "new"}, annotations: map[string]string{"new": "add"},
		displayName: "new addressgroup-0", comment: "added success", description: "add success",
		defaultAction: "DENY", logs: true, trace: true,
	})
}

func (s *GRPCSuiteUpsertAG) TestUpsertAddSameNameDifferentNamespace(t provider.T) {
	s.runSuccess(t, "Добавление AG с существующим именем, но к другому namespace", bodies.UpsertAddSameNameDifferentNamespace(), "ag-0", addressGroupExpectation{
		namespace: "namespace-1", labels: map[string]string{"add": "existname"}, annotations: map[string]string{"new": "add"},
		displayName: "new addressgroup-0", comment: "added success name", description: "add success",
		defaultAction: "ALLOW", logs: true, trace: false,
	})
}

func (s *GRPCSuiteUpsertAG) TestUpsertEdit(t provider.T) {
	s.runSuccess(t, "Редактирование AG", bodies.UpsertEdit(), "ag-3", addressGroupExpectation{
		namespace: "namespace-1", uid: "9e9fb305-80bd-4748-ac42-fc208c398220",
		labels: map[string]string{"edit": "success"}, annotations: map[string]string{"success": "edit"},
		displayName: "edit addressgroup-0", comment: "edit success", description: "edit success",
		defaultAction: "DENY", logs: true, trace: true,
	})
}

func (s *GRPCSuiteUpsertAG) TestUpsertErrorEditUIDMismatch(t provider.T) {
	s.runError(t, "Ошибка при редактировании AG, uid не соответствует указанной AG", bodies.UpsertErrorEditUIDMismatch(), codes.FailedPrecondition, "existing uid AND matching name")
}

func (s *GRPCSuiteUpsertAG) TestUpsertErrorEditEmptyName(t provider.T) {
	s.runError(t, "Ошибка при редактировании AG, name пуст", bodies.UpsertErrorEditEmptyName(), codes.InvalidArgument, "set name for both insert")
}

func (s *GRPCSuiteUpsertAG) TestUpsertErrorEditEmptyNamespace(t provider.T) {
	s.runError(t, "Ошибка при редактировании AG, namespace пуст", bodies.UpsertErrorEditEmptyNamespace(), codes.InvalidArgument, "pass existing namespace for both insert")
}

func (s *GRPCSuiteUpsertAG) TestUpsertErrorEditShortUID(t provider.T) {
	s.runError(t, "Ошибка при редактировании AG, uid короткий", bodies.UpsertErrorEditShortUID(), codes.InvalidArgument, "invalid UUID length")
}

func (s *GRPCSuiteUpsertAG) TestUpsertErrorEditLongUID(t provider.T) {
	s.runError(t, "Ошибка при редактировании AG, uid длинный", bodies.UpsertErrorEditLongUID(), codes.InvalidArgument, "invalid UUID length")
}

func (s *GRPCSuiteUpsertAG) TestUpsertErrorAddEmptyName(t provider.T) {
	s.runError(t, "Ошибка при добавлении AG, name пуст", bodies.UpsertErrorAddEmptyName(), codes.InvalidArgument, "at least one of UID or Name and Namespace must be set")
}

func (s *GRPCSuiteUpsertAG) TestUpsertErrorAddDuplicateNameNamespace(t provider.T) {
	s.runError(t, "Ошибка при добавлении AG, AG c таким name+ns уже существует", bodies.UpsertErrorAddDuplicateNameNamespace(), codes.FailedPrecondition, "for insert pass unique (name, namespace)")
}

func (s *GRPCSuiteUpsertAG) TestUpsertErrorAddNameMissing(t provider.T) {
	s.runError(t, "Ошибка при добавлении AG, name отсутствует как параметр", bodies.UpsertErrorAddNameMissing(), codes.InvalidArgument, "at least one of UID or Name and Namespace must be set")
}

func (s *GRPCSuiteUpsertAG) TestUpsertErrorNameCyrillic(t provider.T) {
	s.runError(t, "Ошибка при добавлении AG, name на кириллице", bodies.UpsertErrorNameCyrillic(), codes.InvalidArgument, "Name: must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteUpsertAG) TestUpsertErrorNameCyrillicDigits(t provider.T) {
	s.runError(t, "Ошибка при добавлении AG, name на кириллице+цифры", bodies.UpsertErrorNameCyrillicDigits(), codes.InvalidArgument, "Name: must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteUpsertAG) TestUpsertErrorNameCyrillicSpecialChars(t provider.T) {
	s.runError(t, "Ошибка при добавлении AG, name на кириллице+спец.символы", bodies.UpsertErrorNameCyrillicSpecialChars(), codes.InvalidArgument, "Name: must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteUpsertAG) TestUpsertErrorNameLatinSpecialChars(t provider.T) {
	s.runError(t, "Ошибка при добавлении AG, name на латинице+спец.символы", bodies.UpsertErrorNameLatinSpecialChars(), codes.InvalidArgument, "Name: must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteUpsertAG) TestUpsertErrorNameUpperCase(t provider.T) {
	s.runError(t, "Ошибка при добавлении AG, name капсом", bodies.UpsertErrorNameUpperCase(), codes.InvalidArgument, "Name: must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteUpsertAG) TestUpsertErrorNameStartsWithHyphen(t provider.T) {
	s.runError(t, "Ошибка при добавлении AG, name начинается с дефиса", bodies.UpsertErrorNameStartsWithHyphen(), codes.InvalidArgument, "Name: must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteUpsertAG) TestUpsertErrorNameEndsWithHyphen(t provider.T) {
	s.runError(t, "Ошибка при добавлении AG, name заканчивается дефисом", bodies.UpsertErrorNameEndsWithHyphen(), codes.InvalidArgument, "Name: must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteUpsertAG) TestUpsertErrorNameSpecialCharsOnly(t provider.T) {
	s.runError(t, "Ошибка при добавлении AG, name из спец.символов", bodies.UpsertErrorNameSpecialCharsOnly(), codes.InvalidArgument, "Name: must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteUpsertAG) TestUpsertErrorNameLeadingSpace(t provider.T) {
	s.runError(t, "Ошибка при добавлении AG, пробел в начале name", bodies.UpsertErrorNameLeadingSpace(), codes.InvalidArgument, "Name: must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteUpsertAG) TestUpsertErrorNameTrailingSpace(t provider.T) {
	s.runError(t, "Ошибка при добавлении AG, пробел в конце name", bodies.UpsertErrorNameTrailingSpace(), codes.InvalidArgument, "Name: must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteUpsertAG) TestUpsertErrorNameWithSpaces(t provider.T) {
	s.runError(t, "Ошибка при добавлении AG, name с пробелами", bodies.UpsertErrorNameWithSpaces(), codes.InvalidArgument, "Name: must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteUpsertAG) TestUpsertErrorNameTooLong(t provider.T) {
	s.runError(t, "Ошибка при добавлении AG, name > 63 символов", bodies.UpsertErrorNameTooLong(), codes.InvalidArgument, "Name: must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteUpsertAG) TestUpsertErrorNamespaceNotExist(t provider.T) {
	s.runError(t, "Ошибка при добавлении AG, namespace не существует", bodies.UpsertErrorNamespaceNotExist(), codes.NotFound, "pass existing namespace name")
}

func (s *GRPCSuiteUpsertAG) TestUpsertErrorNamespaceLeadingSpace(t provider.T) {
	s.runError(t, "Ошибка при добавлении AG, пробел в начале имени namespace", bodies.UpsertErrorNamespaceLeadingSpace(), codes.InvalidArgument, "Namespace: must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteUpsertAG) TestUpsertErrorNamespaceTrailingSpace(t provider.T) {
	s.runError(t, "Ошибка при добавлении AG, пробел в конце имени namespace", bodies.UpsertErrorNamespaceTrailingSpace(), codes.InvalidArgument, "Namespace: must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteUpsertAG) TestUpsertErrorNamespaceWithSpaces(t provider.T) {
	s.runError(t, "Ошибка при добавлении AG, имя namespace с пробелами", bodies.UpsertErrorNamespaceWithSpaces(), codes.InvalidArgument, "Namespace: must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteUpsertAG) TestUpsertErrorNamespaceEmpty(t provider.T) {
	s.runError(t, "Ошибка при добавлении AG, namespace пуст", bodies.UpsertErrorNamespaceEmpty(), codes.InvalidArgument, "at least one of UID or Name and Namespace must be set")
}

func (s *GRPCSuiteUpsertAG) TestUpsertErrorNamespaceMissing(t provider.T) {
	s.runError(t, "Ошибка при добавлении AG, namespace отсутствует как параметр", bodies.UpsertErrorNamespaceMissing(), codes.InvalidArgument, "at least one of UID or Name and Namespace must be set")
}

func (s *GRPCSuiteUpsertAG) TestUpsertErrorAddExistingName(t provider.T) {
	s.runError(t, "Ошибка при добавлении AG, AG с таким именем существует", bodies.UpsertErrorAddExistingName(), codes.FailedPrecondition, "for insert pass unique (name, namespace); for update pass existing uid AND matching name AND namespace")
}

func (s *GRPCSuiteUpsertAG) TestUpsertErrorDisplayNameTooLong(t provider.T) {
	s.runError(t, "Ошибка при добавлении AG, displayName > 63 символов", bodies.UpsertErrorDisplayNameTooLong(), codes.InvalidArgument, "display name must be no longer than 63 characters")
}

func (s *GRPCSuiteUpsertAG) TestUpsertErrorDefaultActionMissing(t provider.T) {
	s.runError(t, "Ошибка при добавлении AG, defaultAction отсутствует как параметр", bodies.UpsertErrorDefaultActionMissing(), codes.InvalidArgument, "DefaultAction", "ALLOW,DENY")
}
