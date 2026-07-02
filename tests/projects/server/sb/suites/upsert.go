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
	"HBF-tests/tests/projects/server/sb/internal/bodies"
)

type GRPCSuiteUpsertSB struct {
	suite.Suite
	TestStruct foundation.ToolsForTest
	client     grpcClient.SGroupsServiceBindingAPIClient
}

func (s *GRPCSuiteUpsertSB) BeforeAll(t provider.T) {
	s.TestStruct.ServiceBindingAPI = grpcClient.NewSGroupsServiceBindingAPIClient(connection.GRPC(t))
}

func (s *GRPCSuiteUpsertSB) runSuccess(t provider.T, story string, body *foundation.TestCaseBodySBUpsert, name string, exp serviceBindingExpectation) {
	t.Parallel()
	t.Story(story)
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		resp, err := s.TestStruct.FoundationTestsSBUpsert(body)

		sCtx.NewStep(steps.AssertStep)
		sCtx.Require().NoError(err)
		sCtx.Require().NotNil(resp)
		sCtx.Require().Len(resp.ServiceBindings, 1)
		assertServiceBinding(sCtx, resp.ServiceBindings[0], name, exp)
	})
}

func (s *GRPCSuiteUpsertSB) runError(t provider.T, story string, body *foundation.TestCaseBodySBUpsert, code codes.Code, substrings ...string) {
	t.Parallel()
	t.Story(story)
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		_, err := s.TestStruct.FoundationTestsSBUpsert(body)

		sCtx.NewStep(steps.AssertStep)
		utils.AssertGRPCError(sCtx, err, code, substrings...)
	})
}

func (s *GRPCSuiteUpsertSB) TestUpsertAdd(t provider.T) {
	s.runSuccess(t, "Добавление нового serviceBindings (все в одном NS)", bodies.UpsertAdd(), "add-sb-1", serviceBindingExpectation{
		namespace: "namespace-1", labels: map[string]string{"add": "success"}, annotations: map[string]string{"add": "success"},
		displayName: "new service binding 1", comment: "new service binding", description: "add success",
		addressGroupName: "ag-6", addressGroupNamespace: "namespace-1", serviceName: "svc-6", serviceNamespace: "namespace-1",
	})
}

func (s *GRPCSuiteUpsertSB) TestUpsertAddAGDifferentNamespace(t provider.T) {
	s.runSuccess(t, "Добавление нового serviceBindings (AG в другом NS)", bodies.UpsertAddAGDifferentNamespace(), "add-sb-2", serviceBindingExpectation{
		namespace: "namespace-1", labels: map[string]string{"add": "success"}, annotations: map[string]string{"add": "success"},
		displayName: "new service binding 1", comment: "new service binding", description: "add success",
		addressGroupName: "ag-2", addressGroupNamespace: "namespace-2", serviceName: "svc-6", serviceNamespace: "namespace-1",
	})
}

func (s *GRPCSuiteUpsertSB) TestUpsertEdit(t provider.T) {
	s.runSuccess(t, "Редактирование serviceBindings", bodies.UpsertEdit(), "sb-3", serviceBindingExpectation{
		namespace: "namespace-1", uid: "a915bc00-2feb-42a5-94af-e209effbce44",
		labels: map[string]string{"edit": "success"}, annotations: map[string]string{"edit": "success"},
		displayName: "edit service binding 1", comment: "edited service binding", description: "edit success",
		addressGroupName: "ag-3", addressGroupNamespace: "namespace-1", serviceName: "svc-3", serviceNamespace: "namespace-1",
	})
}

func (s *GRPCSuiteUpsertSB) TestUpsertAddSameNameDifferentNamespace(t provider.T) {
	s.runSuccess(t, "Добавление serviceBindings с существующим именем, но к другому namespace", bodies.UpsertAddSameNameDifferentNamespace(), "sb-0", serviceBindingExpectation{
		namespace: "namespace-1", labels: map[string]string{"add": "success"}, annotations: map[string]string{"add": "success"},
		displayName: "new service binding 1", comment: "new service binding", description: "add success",
		addressGroupName: "ag-2", addressGroupNamespace: "namespace-2", serviceName: "svc-7", serviceNamespace: "namespace-1",
	})
}

func (s *GRPCSuiteUpsertSB) TestUpsertErrorPortsOverlap(t provider.T) {
	s.runError(t, "Ошибка при добавлении, сервисы пересекаются портами", bodies.UpsertErrorPortsOverlap(), codes.AlreadyExists, "ranges overlap with another bound service")
}

func (s *GRPCSuiteUpsertSB) TestUpsertErrorDisplayNameTooLong(t provider.T) {
	s.runError(t, "Ошибка при добавлении serviceBindings, displayName > 63 символов", bodies.UpsertErrorDisplayNameTooLong(), codes.InvalidArgument, "display name must be no longer than 63 characters")
}

func (s *GRPCSuiteUpsertSB) TestUpsertErrorEditAGAndSVCMissing(t provider.T) {
	s.runError(t, "Ошибка при редактировании serviceBindings, ag и svc отсутствуют", bodies.UpsertErrorEditAGAndSVCMissing(), codes.InvalidArgument, "AddressGroup", "Service", "resource name is required")
}

func (s *GRPCSuiteUpsertSB) TestUpsertErrorEditSVCImmutable(t provider.T) {
	s.runError(t, "Ошибка при редактировании serviceBindings, нельзя редактировать svc", bodies.UpsertErrorEditSVCImmutable(), codes.InvalidArgument, "service cannot be updated")
}

func (s *GRPCSuiteUpsertSB) TestUpsertErrorEditSVCImmutableDifferentNamespace(t provider.T) {
	s.runError(t, "Ошибка при редактировании serviceBindings, нельзя редактировать svc из другого ns", bodies.UpsertErrorEditSVCImmutableDifferentNamespace(), codes.InvalidArgument, "binding and service must be in the same namespace")
}

func (s *GRPCSuiteUpsertSB) TestUpsertErrorEditSVCMissing(t provider.T) {
	s.runError(t, "Ошибка при редактировании serviceBindings, svc отсутствует", bodies.UpsertErrorEditSVCMissing(), codes.InvalidArgument, "Service", "resource name is required", "resource namespace is required")
}

func (s *GRPCSuiteUpsertSB) TestUpsertErrorEditAGImmutable(t provider.T) {
	s.runError(t, "Ошибка при редактировании serviceBindings, нельзя редактировать ag", bodies.UpsertErrorEditAGImmutable(), codes.InvalidArgument, "ag cannot be updated")
}

func (s *GRPCSuiteUpsertSB) TestUpsertErrorEditAGImmutableDifferentNamespace(t provider.T) {
	s.runError(t, "Ошибка при редактировании serviceBindings, нельзя редактировать ag из другого ns", bodies.UpsertErrorEditAGImmutableDifferentNamespace(), codes.InvalidArgument, "ag cannot be updated")
}

func (s *GRPCSuiteUpsertSB) TestUpsertErrorEditAGMissing(t provider.T) {
	s.runError(t, "Ошибка при редактировании serviceBindings, ag отсутствует", bodies.UpsertErrorEditAGMissing(), codes.InvalidArgument, "AddressGroup", "resource name is required", "resource namespace is required")
}

func (s *GRPCSuiteUpsertSB) TestUpsertErrorBindingAlreadyExistsSameNamespace(t provider.T) {
	s.runError(t, "Ошибка при добавлении serviceBindings, такая привязка уже существует (svc-host в одном ns)", bodies.UpsertErrorBindingAlreadyExistsSameNamespace(), codes.FailedPrecondition, "binding: this (ag, service) pair already exists")
}

func (s *GRPCSuiteUpsertSB) TestUpsertErrorBindingAlreadyExistsDifferentNamespace(t provider.T) {
	s.runError(t, "Ошибка при добавлении serviceBindings, такая привязка уже существует (svc-host в разных ns)", bodies.UpsertErrorBindingAlreadyExistsDifferentNamespace(), codes.FailedPrecondition, "duplicate binding: this (ag, service) pair already exists")
}

func (s *GRPCSuiteUpsertSB) TestUpsertErrorAddDuplicateNameNamespace(t provider.T) {
	s.runError(t, "Ошибка при добавлении serviceBindings, sb с таким name+namespace уже существует", bodies.UpsertErrorAddDuplicateNameNamespace(), codes.FailedPrecondition, "for insert pass unique name; for update pass existing uid AND matching name")
}

func (s *GRPCSuiteUpsertSB) TestUpsertErrorSVCNamespaceMismatch(t provider.T) {
	s.runError(t, "Ошибка при добавлении serviceBindings, svc ns с не совпадает с sb ns", bodies.UpsertErrorSVCNamespaceMismatch(), codes.InvalidArgument, "binding and service must be in the same namespace")
}

func (s *GRPCSuiteUpsertSB) TestUpsertErrorAllDifferentNamespaces(t provider.T) {
	s.runError(t, "Ошибка при добавлении serviceBindings, все в разных NS", bodies.UpsertErrorAllDifferentNamespaces(), codes.InvalidArgument, "binding and service must be in the same namespace")
}

func (s *GRPCSuiteUpsertSB) TestUpsertErrorEditLongUID(t provider.T) {
	s.runError(t, "Ошибка при редактировании serviceBindings, uid длинный", bodies.UpsertErrorEditLongUID(), codes.InvalidArgument, "invalid UUID length")
}

func (s *GRPCSuiteUpsertSB) TestUpsertErrorEditShortUID(t provider.T) {
	s.runError(t, "Ошибка при редактировании serviceBindings, uid короткий", bodies.UpsertErrorEditShortUID(), codes.InvalidArgument, "invalid UUID length")
}

func (s *GRPCSuiteUpsertSB) TestUpsertErrorEditUIDMismatch(t provider.T) {
	s.runError(t, "Ошибка при редактировании serviceBindings, uid не соответствует указанному sb", bodies.UpsertErrorEditUIDMismatch(), codes.FailedPrecondition, "existing uid AND matching name")
}

func (s *GRPCSuiteUpsertSB) TestUpsertErrorEditNameUIDNotExistInNamespace(t provider.T) {
	s.runError(t, "Ошибка при редактировании serviceBindings, name+uid не существует в указанном namespace", bodies.UpsertErrorEditNameUIDNotExistInNamespace(), codes.FailedPrecondition, "existing uid AND matching name")
}

func (s *GRPCSuiteUpsertSB) TestUpsertErrorEditEmptyName(t provider.T) {
	s.runError(t, "Ошибка при редактировании serviceBindings, name пуст", bodies.UpsertErrorEditEmptyName(), codes.InvalidArgument, "set name and namespace")
}

func (s *GRPCSuiteUpsertSB) TestUpsertErrorEditNameMissing(t provider.T) {
	s.runError(t, "Ошибка при редактировании serviceBindings, name отсутствует как параметр", bodies.UpsertErrorEditNameMissing(), codes.InvalidArgument, "set name and namespace")
}

func (s *GRPCSuiteUpsertSB) TestUpsertErrorEditEmptyNamespace(t provider.T) {
	s.runError(t, "Ошибка при редактировании serviceBindings, namespace пуст", bodies.UpsertErrorEditEmptyNamespace(), codes.InvalidArgument, "set name and namespace")
}

func (s *GRPCSuiteUpsertSB) TestUpsertErrorEditNamespaceMissing(t provider.T) {
	s.runError(t, "Ошибка при редактировании serviceBindings, namespace отсутствует как параметр", bodies.UpsertErrorEditNamespaceMissing(), codes.InvalidArgument, "set name and namespace")
}

func (s *GRPCSuiteUpsertSB) TestUpsertErrorAddEmptyName(t provider.T) {
	s.runError(t, "Ошибка при добавлении serviceBindings, name пуст", bodies.UpsertErrorAddEmptyName(), codes.InvalidArgument, "at least one of UID or Name and Namespace must be set")
}

func (s *GRPCSuiteUpsertSB) TestUpsertErrorNameCyrillic(t provider.T) {
	s.runError(t, "Ошибка при добавлении serviceBindings, name на кириллице", bodies.UpsertErrorNameCyrillic(), codes.InvalidArgument, "Name: must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteUpsertSB) TestUpsertErrorNameUpperCase(t provider.T) {
	s.runError(t, "Ошибка при добавлении serviceBindings, name капсом", bodies.UpsertErrorNameUpperCase(), codes.InvalidArgument, "Name: must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteUpsertSB) TestUpsertErrorNameSpecialCharsOnly(t provider.T) {
	s.runError(t, "Ошибка при добавлении serviceBindings, name из спец символов", bodies.UpsertErrorNameSpecialCharsOnly(), codes.InvalidArgument, "Name: must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteUpsertSB) TestUpsertErrorNameMixedCase(t provider.T) {
	s.runError(t, "Ошибка при добавлении serviceBindings, name капс+строчные", bodies.UpsertErrorNameMixedCase(), codes.InvalidArgument, "Name: must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteUpsertSB) TestUpsertErrorNameCyrillicLatin(t provider.T) {
	s.runError(t, "Ошибка при добавлении serviceBindings, name на кириллице+латинице", bodies.UpsertErrorNameCyrillicLatin(), codes.InvalidArgument, "Name: must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteUpsertSB) TestUpsertErrorNameCyrillicDigits(t provider.T) {
	s.runError(t, "Ошибка при добавлении serviceBindings, name на кириллице+числа", bodies.UpsertErrorNameCyrillicDigits(), codes.InvalidArgument, "Name: must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteUpsertSB) TestUpsertErrorNameStartsWithHyphen(t provider.T) {
	s.runError(t, "Ошибка при добавлении serviceBindings, name начинается с тире", bodies.UpsertErrorNameStartsWithHyphen(), codes.InvalidArgument, "Name: must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteUpsertSB) TestUpsertErrorNameOnlyHyphen(t provider.T) {
	s.runError(t, "Ошибка при добавлении serviceBindings, name состоит только из тире", bodies.UpsertErrorNameOnlyHyphen(), codes.InvalidArgument, "Name: must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteUpsertSB) TestUpsertErrorNameEndsWithHyphen(t provider.T) {
	s.runError(t, "Ошибка при добавлении serviceBindings, name заканчивается тире", bodies.UpsertErrorNameEndsWithHyphen(), codes.InvalidArgument, "Name: must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteUpsertSB) TestUpsertErrorNameCyrillicSpecialChars(t provider.T) {
	s.runError(t, "Ошибка при добавлении serviceBindings, name на кириллице+спец.символы", bodies.UpsertErrorNameCyrillicSpecialChars(), codes.InvalidArgument, "Name: must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteUpsertSB) TestUpsertErrorNameTooLong(t provider.T) {
	s.runError(t, "Ошибка при добавлении serviceBindings, name > 63 символов", bodies.UpsertErrorNameTooLong(), codes.InvalidArgument, "Name: must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteUpsertSB) TestUpsertErrorNameLeadingSpace(t provider.T) {
	s.runError(t, "Ошибка при добавлении serviceBindings, пробел в начале name", bodies.UpsertErrorNameLeadingSpace(), codes.InvalidArgument, "Name: must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteUpsertSB) TestUpsertErrorNameTrailingSpace(t provider.T) {
	s.runError(t, "Ошибка при добавлении serviceBindings, пробел в конце name", bodies.UpsertErrorNameTrailingSpace(), codes.InvalidArgument, "Name: must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteUpsertSB) TestUpsertErrorNameWithSpaces(t provider.T) {
	s.runError(t, "Ошибка при добавлении serviceBindings, name с пробелами", bodies.UpsertErrorNameWithSpaces(), codes.InvalidArgument, "Name: must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteUpsertSB) TestUpsertErrorNamespaceNotExist(t provider.T) {
	s.runError(t, "Ошибка при добавлении serviceBindings, namespace не существует", bodies.UpsertErrorNamespaceNotExist(), codes.NotFound, "pass existing namespace name")
}

func (s *GRPCSuiteUpsertSB) TestUpsertErrorNamespaceEmpty(t provider.T) {
	s.runError(t, "Ошибка при добавлении serviceBindings, namespace пуст", bodies.UpsertErrorNamespaceEmpty(), codes.InvalidArgument, "ID: at least one of UID or Name and Namespace must be set")
}

func (s *GRPCSuiteUpsertSB) TestUpsertErrorNamespaceLeadingSpace(t provider.T) {
	s.runError(t, "Ошибка при добавлении serviceBindings, пробел в начале имени namespace", bodies.UpsertErrorNamespaceLeadingSpace(), codes.InvalidArgument, "Namespace: must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteUpsertSB) TestUpsertErrorNamespaceTrailingSpace(t provider.T) {
	s.runError(t, "Ошибка при добавлении networkBindings, пробел в конце имени namespace", bodies.UpsertErrorNamespaceTrailingSpace(), codes.InvalidArgument, "Namespace: must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteUpsertSB) TestUpsertErrorNamespaceWithSpaces(t provider.T) {
	s.runError(t, "Ошибка при добавлении serviceBindings, namespace с пробелами", bodies.UpsertErrorNamespaceWithSpaces(), codes.InvalidArgument, "Namespace: must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteUpsertSB) TestUpsertErrorNamespaceMissing(t provider.T) {
	s.runError(t, "Ошибка при добавлении serviceBindings, namespace отсутствует как параметр", bodies.UpsertErrorNamespaceMissing(), codes.InvalidArgument, "at least one of UID or Name and Namespace must be set")
}

func (s *GRPCSuiteUpsertSB) TestUpsertErrorSVCNotExistInNamespace(t provider.T) {
	s.runError(t, "Ошибка при добавлении serviceBindings, svc не существует в указанном ns", bodies.UpsertErrorSVCNotExistInNamespace(), codes.NotFound, "pass existing service name and namespace")
}

func (s *GRPCSuiteUpsertSB) TestUpsertErrorSVCNotInNamespaceButExists(t provider.T) {
	s.runError(t, "Ошибка при добавлении serviceBindings, в указанном ns нет такого svc (сам svc существует)", bodies.UpsertErrorSVCNotInNamespaceButExists(), codes.NotFound, "pass existing service name and namespace")
}

func (s *GRPCSuiteUpsertSB) TestUpsertErrorSVCNameEmpty(t provider.T) {
	s.runError(t, "Ошибка при добавлении serviceBindings, svc name пуст", bodies.UpsertErrorSVCNameEmpty(), codes.InvalidArgument, "Service", "Name", "resource name is required")
}

func (s *GRPCSuiteUpsertSB) TestUpsertErrorSVCEmpty(t provider.T) {
	s.runError(t, "Ошибка при добавлении serviceBindings, svc пуст (ns и name)", bodies.UpsertErrorSVCEmpty(), codes.InvalidArgument, "Service", "Name", "Namespace", "resource name is required")
}

func (s *GRPCSuiteUpsertSB) TestUpsertErrorSVCNameMissing(t provider.T) {
	s.runError(t, "Ошибка при добавлении serviceBindings, в svc name отсутствует как параметр", bodies.UpsertErrorSVCNameMissing(), codes.InvalidArgument, "Service", "Name", "resource name is required")
}

func (s *GRPCSuiteUpsertSB) TestUpsertErrorSVCNamespaceMissing(t provider.T) {
	s.runError(t, "Ошибка при добавлении serviceBindings, в svc ns отсутствует как параметр", bodies.UpsertErrorSVCNamespaceMissing(), codes.InvalidArgument, "Service", "Namespace", "resource namespace is required")
}

func (s *GRPCSuiteUpsertSB) TestUpsertErrorSVCEmptyObject(t provider.T) {
	s.runError(t, "Ошибка при добавлении serviceBindings, svc пустой объект", bodies.UpsertErrorSVCEmptyObject(), codes.InvalidArgument, "Service", "Namespace", "resource namespace is required")
}

func (s *GRPCSuiteUpsertSB) TestUpsertErrorSVCMissing(t provider.T) {
	s.runError(t, "Ошибка при добавлении serviceBindings, svc отсутствует как параметр", bodies.UpsertErrorSVCMissing(), codes.InvalidArgument, "Service", "Namespace", "resource namespace is required")
}

func (s *GRPCSuiteUpsertSB) TestUpsertErrorAGEmptyObject(t provider.T) {
	s.runError(t, "Ошибка при добавлении serviceBindings, ag пустой объект", bodies.UpsertErrorAGEmptyObject(), codes.InvalidArgument, "AddressGroup", "Namespace", "Name", "resource namespace is required")
}

func (s *GRPCSuiteUpsertSB) TestUpsertErrorAGMissing(t provider.T) {
	s.runError(t, "Ошибка при добавлении serviceBindings, ag отсутствует как параметр", bodies.UpsertErrorAGMissing(), codes.InvalidArgument, "AddressGroup", "Namespace", "Name", "resource namespace is required")
}

func (s *GRPCSuiteUpsertSB) TestUpsertErrorAGNamespaceMissing(t provider.T) {
	s.runError(t, "Ошибка при добавлении serviceBindings, в ag ns отсутствует как параметр", bodies.UpsertErrorAGNamespaceMissing(), codes.InvalidArgument, "AddressGroup", "Namespace", "Name", "resource namespace is required")
}

func (s *GRPCSuiteUpsertSB) TestUpsertErrorAGNameMissing(t provider.T) {
	s.runError(t, "Ошибка при добавлении serviceBindings, в ag name отсутствует как параметр", bodies.UpsertErrorAGNameMissing(), codes.InvalidArgument, "AddressGroup", "Name", "resource name is required")
}

func (s *GRPCSuiteUpsertSB) TestUpsertErrorAGEmpty(t provider.T) {
	s.runError(t, "Ошибка при добавлении serviceBindings, ag пуст (ns и name)", bodies.UpsertErrorAGEmpty(), codes.InvalidArgument, "AddressGroup", "Name", "Namespace", "resource name is required")
}

func (s *GRPCSuiteUpsertSB) TestUpsertErrorAGNameEmpty(t provider.T) {
	s.runError(t, "Ошибка при добавлении serviceBindings, ag name пуст", bodies.UpsertErrorAGNameEmpty(), codes.InvalidArgument, "AddressGroup", "Name", "resource name is required")
}

func (s *GRPCSuiteUpsertSB) TestUpsertErrorAGNotInNamespaceButExists(t provider.T) {
	s.runError(t, "Ошибка при добавлении serviceBindings, в указанном ns нет такого ag (сам ag существует)", bodies.UpsertErrorAGNotInNamespaceButExists(), codes.NotFound, "pass existing address_group name and namespace")
}

func (s *GRPCSuiteUpsertSB) TestUpsertErrorAGNotExist(t provider.T) {
	s.runError(t, "Ошибка при добавлении serviceBindings, ag не существует в указанном ns", bodies.UpsertErrorAGNotExist(), codes.NotFound, "pass existing address_group name and namespace")
}

func (s *GRPCSuiteUpsertSB) TestUpsertErrorSpecEmpty(t provider.T) {
	s.runError(t, "Ошибка при добавлении serviceBindings, spec пустой объект", bodies.UpsertErrorSpecEmpty(), codes.InvalidArgument, "AddressGroup", "Service", "resource name is required")
}
