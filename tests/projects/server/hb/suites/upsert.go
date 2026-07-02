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

type GRPCSuiteUpsertHB struct {
	suite.Suite
	TestStruct foundation.ToolsForTest
	client     grpcClient.SGroupsHostBindingAPIClient
}

func (s *GRPCSuiteUpsertHB) BeforeAll(t provider.T) {
	s.TestStruct.HostBindingAPI = grpcClient.NewSGroupsHostBindingAPIClient(connection.GRPC(t))
}

func (s *GRPCSuiteUpsertHB) runSuccess(t provider.T, story string, body *foundation.TestCaseBodyHBUpsert, name string, exp hostBindingExpectation) {
	t.Parallel()
	t.Story(story)
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		resp, err := s.TestStruct.FoundationTestsHBUpsert(body)

		sCtx.NewStep(steps.AssertStep)
		sCtx.Require().NoError(err)
		sCtx.Require().NotNil(resp)
		sCtx.Require().Len(resp.HostBindings, 1)
		assertHostBinding(sCtx, resp.HostBindings[0], name, exp)
	})
}

func (s *GRPCSuiteUpsertHB) runError(t provider.T, story string, body *foundation.TestCaseBodyHBUpsert, code codes.Code, substrings ...string) {
	t.Parallel()
	t.Story(story)
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		_, err := s.TestStruct.FoundationTestsHBUpsert(body)

		sCtx.NewStep(steps.AssertStep)
		utils.AssertGRPCError(sCtx, err, code, substrings...)
	})
}

func (s *GRPCSuiteUpsertHB) TestUpsertAdd(t provider.T) {
	s.runSuccess(t, "Добавление нового host-binding", bodies.UpsertAdd(), "add-hb-1", hostBindingExpectation{
		namespace: "namespace-1", labels: map[string]string{"add": "success"}, annotations: map[string]string{"add": "success"},
		displayName: "new host binding 1", comment: "new host binding", description: "add success",
		addressGroupName: "ag-6", addressGroupNamespace: "namespace-1", hostName: "host-6", hostNamespace: "namespace-1",
	})
}

func (s *GRPCSuiteUpsertHB) TestUpsertEdit(t provider.T) {
	s.runSuccess(t, "Редактирование host-binging", bodies.UpsertEdit(), "hb-3", hostBindingExpectation{
		namespace: "namespace-1", uid: "efb85252-c9b7-4dd5-b9be-6ce63276795c",
		labels: map[string]string{"edit": "success"}, annotations: map[string]string{"edit": "success"},
		displayName: "edit host binding 1", comment: "edited host binding", description: "edit success",
		addressGroupName: "ag-3", addressGroupNamespace: "namespace-1", hostName: "host-3", hostNamespace: "namespace-1",
	})
}

func (s *GRPCSuiteUpsertHB) TestUpsertAddSameNameDifferentNamespace(t provider.T) {
	s.runSuccess(t, "Добавление host-binding с существующим именем, но к другому namespace", bodies.UpsertAddSameNameDifferentNamespace(), "hb-0", hostBindingExpectation{
		namespace: "namespace-1", labels: map[string]string{"add": "success"}, annotations: map[string]string{"add": "success"},
		displayName: "new host binding 1", comment: "new host binding", description: "add success",
		addressGroupName: "ag-6", addressGroupNamespace: "namespace-1", hostName: "host-7", hostNamespace: "namespace-1",
	})
}

func (s *GRPCSuiteUpsertHB) TestUpsertErrorDisplayNameTooLong(t provider.T) {
	s.runError(t, "Ошибка при добавлении host-binding, displayName > 63 символов", bodies.UpsertErrorDisplayNameTooLong(), codes.InvalidArgument, "DisplayName", "display name must be no longer than 63 characters")
}

func (s *GRPCSuiteUpsertHB) TestUpsertErrorEditAGAndHostMissing(t provider.T) {
	s.runError(t, "Ошибка при редактировании host-binding,  ag и host отсутствуют", bodies.UpsertErrorEditAGAndHostMissing(), codes.InvalidArgument, "AddressGroup", "Host", "resource name is required")
}

func (s *GRPCSuiteUpsertHB) TestUpsertErrorEditHostImmutable(t provider.T) {
	s.runError(t, "Ошибка при редактировании host-binding, нельзя редактировать host", bodies.UpsertErrorEditHostImmutable(), codes.InvalidArgument, "host cannot be updated")
}

func (s *GRPCSuiteUpsertHB) TestUpsertErrorEditHostImmutableDifferentNamespace(t provider.T) {
	s.runError(t, "Ошибка при редактировании host-binding, нельзя редактировать host из другого ns", bodies.UpsertErrorEditHostImmutableDifferentNamespace(), codes.InvalidArgument, "pass host and address group in the same namespace")
}

func (s *GRPCSuiteUpsertHB) TestUpsertErrorEditHostMissing(t provider.T) {
	s.runError(t, "Ошибка при редактировании host-binding, host отсутствует", bodies.UpsertErrorEditHostMissing(), codes.InvalidArgument, "Host", "resource name is required", "resource namespace is required")
}

func (s *GRPCSuiteUpsertHB) TestUpsertErrorEditAGImmutable(t provider.T) {
	s.runError(t, "Ошибка при редактировании host-binding, нельзя редактировать ag", bodies.UpsertErrorEditAGImmutable(), codes.InvalidArgument, "ag cannot be updated")
}

func (s *GRPCSuiteUpsertHB) TestUpsertErrorEditAGImmutableDifferentNamespace(t provider.T) {
	s.runError(t, "Ошибка при редактировании host-binding, нельзя редактировать ag из другого ns", bodies.UpsertErrorEditAGImmutableDifferentNamespace(), codes.InvalidArgument, "pass host and address group in the same namespace")
}

func (s *GRPCSuiteUpsertHB) TestUpsertErrorEditAGMissing(t provider.T) {
	s.runError(t, "Ошибка при редактировании host-binding, ag отсутствует", bodies.UpsertErrorEditAGMissing(), codes.InvalidArgument, "AddressGroup", "resource name is required", "resource namespace is required")
}

func (s *GRPCSuiteUpsertHB) TestUpsertErrorBindingAlreadyExists(t provider.T) {
	s.runError(t, "Ошибка при добавлении host-binding, такая привязка уже существует (ag-host)", bodies.UpsertErrorBindingAlreadyExists(), codes.FailedPrecondition, "binding (address_group, host) must be unique")
}

func (s *GRPCSuiteUpsertHB) TestUpsertErrorAddDuplicateNameNamespace(t provider.T) {
	s.runError(t, "Ошибка при добавлении host-binding, hb  с таким name+namespace уже существует", bodies.UpsertErrorAddDuplicateNameNamespace(), codes.FailedPrecondition, "binding (address_group, host) must be unique")
}

func (s *GRPCSuiteUpsertHB) TestUpsertErrorHostNamespaceMismatch(t provider.T) {
	s.runError(t, "Ошибка при добавлении host-binding, host ns с не совпадает с hb ns", bodies.UpsertErrorHostNamespaceMismatch(), codes.InvalidArgument, "pass host and address group in the same namespace")
}

func (s *GRPCSuiteUpsertHB) TestUpsertErrorAGNamespaceMismatch(t provider.T) {
	s.runError(t, "Ошибка при добавлении host-binding, ag ns с не совпадает с hb ns", bodies.UpsertErrorAGNamespaceMismatch(), codes.InvalidArgument, "pass host and address group in the same namespace")
}

func (s *GRPCSuiteUpsertHB) TestUpsertErrorHBNamespaceMismatchBoth(t provider.T) {
	s.runError(t, "Ошибка при добавлении host-binding, hb ns не совпадает с ag и host ns", bodies.UpsertErrorHBNamespaceMismatchBoth(), codes.InvalidArgument, "pass host binding namespace that matches host namespace")
}

func (s *GRPCSuiteUpsertHB) TestUpsertErrorEditLongUID(t provider.T) {
	s.runError(t, "Ошибка при редактировании host-binding, uid длинный", bodies.UpsertErrorEditLongUID(), codes.InvalidArgument, "invalid UUID length")
}

func (s *GRPCSuiteUpsertHB) TestUpsertErrorEditShortUID(t provider.T) {
	s.runError(t, "Ошибка при редактировании host-binding, uid короткий", bodies.UpsertErrorEditShortUID(), codes.InvalidArgument, "invalid UUID length")
}

func (s *GRPCSuiteUpsertHB) TestUpsertErrorEditUIDMismatch(t provider.T) {
	s.runError(t, "Ошибка при редактировании host-binding, uid не соответствует указанному hb", bodies.UpsertErrorEditUIDMismatch(), codes.FailedPrecondition, "existing uid AND matching name")
}

func (s *GRPCSuiteUpsertHB) TestUpsertErrorEditNameUIDNotExistInNamespace(t provider.T) {
	s.runError(t, "Ошибка при редактировании host-binding, name+uid не существует в указанном namespace", bodies.UpsertErrorEditNameUIDNotExistInNamespace(), codes.FailedPrecondition, "existing uid AND matching name")
}

func (s *GRPCSuiteUpsertHB) TestUpsertErrorEditEmptyName(t provider.T) {
	s.runError(t, "Ошибка при редактировании host-binding, name пуст", bodies.UpsertErrorEditEmptyName(), codes.InvalidArgument, "set name for both insert")
}

func (s *GRPCSuiteUpsertHB) TestUpsertErrorEditNameMissing(t provider.T) {
	s.runError(t, "Ошибка при редактировании host-binding, name отсутствует как параметр", bodies.UpsertErrorEditNameMissing(), codes.InvalidArgument, "set name for both insert")
}

func (s *GRPCSuiteUpsertHB) TestUpsertErrorEditEmptyNamespace(t provider.T) {
	s.runError(t, "Ошибка при редактировании host-binding, namespace пуст", bodies.UpsertErrorEditEmptyNamespace(), codes.InvalidArgument, "pass existing namespace for both insert")
}

func (s *GRPCSuiteUpsertHB) TestUpsertErrorEditNamespaceMissing(t provider.T) {
	s.runError(t, "Ошибка при редактировании host-binding, namespace отсутствует как параметр", bodies.UpsertErrorEditNamespaceMissing(), codes.InvalidArgument, "pass existing namespace for both insert")
}

func (s *GRPCSuiteUpsertHB) TestUpsertErrorAddEmptyName(t provider.T) {
	s.runError(t, "Ошибка при добавлении host-binding, name пуст", bodies.UpsertErrorAddEmptyName(), codes.InvalidArgument, "at least one of UID or Name and Namespace must be set")
}

func (s *GRPCSuiteUpsertHB) TestUpsertErrorNameCyrillic(t provider.T) {
	s.runError(t, "Ошибка при добавлении host-binding, name на кириллице", bodies.UpsertErrorNameCyrillic(), codes.InvalidArgument, "Name: must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteUpsertHB) TestUpsertErrorNameUpperCase(t provider.T) {
	s.runError(t, "Ошибка при добавлении host-binding, name капсом", bodies.UpsertErrorNameUpperCase(), codes.InvalidArgument, "Name: must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteUpsertHB) TestUpsertErrorNameSpecialCharsOnly(t provider.T) {
	s.runError(t, "Ошибка при добавлении host-binding, name из спец символов", bodies.UpsertErrorNameSpecialCharsOnly(), codes.InvalidArgument, "Name: must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteUpsertHB) TestUpsertErrorNameMixedCase(t provider.T) {
	s.runError(t, "Ошибка при добавлении host-binding, name капс+строчные", bodies.UpsertErrorNameMixedCase(), codes.InvalidArgument, "Name: must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteUpsertHB) TestUpsertErrorNameCyrillicLatin(t provider.T) {
	s.runError(t, "Ошибка при добавлении host-binding, name на кириллице+латинице", bodies.UpsertErrorNameCyrillicLatin(), codes.InvalidArgument, "Name: must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteUpsertHB) TestUpsertErrorNameCyrillicDigits(t provider.T) {
	s.runError(t, "Ошибка при добавлении host-binding, name на кириллице+числа", bodies.UpsertErrorNameCyrillicDigits(), codes.InvalidArgument, "Name: must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteUpsertHB) TestUpsertErrorNameStartsWithHyphen(t provider.T) {
	s.runError(t, "Ошибка при добавлении host-binding, name начинается с тире", bodies.UpsertErrorNameStartsWithHyphen(), codes.InvalidArgument, "Name: must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteUpsertHB) TestUpsertErrorNameOnlyHyphen(t provider.T) {
	s.runError(t, "Ошибка при добавлении host-binding, name состоит только из тире", bodies.UpsertErrorNameOnlyHyphen(), codes.InvalidArgument, "Name: must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteUpsertHB) TestUpsertErrorNameEndsWithHyphen(t provider.T) {
	s.runError(t, "Ошибка при добавлении host-binding, name заканчивается тире", bodies.UpsertErrorNameEndsWithHyphen(), codes.InvalidArgument, "Name: must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteUpsertHB) TestUpsertErrorNameCyrillicSpecialChars(t provider.T) {
	s.runError(t, "Ошибка при добавлении host-binding, name на кириллице+спец.символы", bodies.UpsertErrorNameCyrillicSpecialChars(), codes.InvalidArgument, "Name: must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteUpsertHB) TestUpsertErrorNameTooLong(t provider.T) {
	s.runError(t, "Ошибка при добавлении host-binding, name > 63 символов", bodies.UpsertErrorNameTooLong(), codes.InvalidArgument, "Name: must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteUpsertHB) TestUpsertErrorNameLeadingSpace(t provider.T) {
	s.runError(t, "Ошибка при добавлении host-binding, пробел в начале name", bodies.UpsertErrorNameLeadingSpace(), codes.InvalidArgument, "Name: must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteUpsertHB) TestUpsertErrorNameTrailingSpace(t provider.T) {
	s.runError(t, "Ошибка при добавлении host-binding, пробел в конце name", bodies.UpsertErrorNameTrailingSpace(), codes.InvalidArgument, "Name: must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteUpsertHB) TestUpsertErrorNameWithSpaces(t provider.T) {
	s.runError(t, "Ошибка при добавлении host-binding, name с пробелами", bodies.UpsertErrorNameWithSpaces(), codes.InvalidArgument, "Name: must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteUpsertHB) TestUpsertErrorNamespaceNotExist(t provider.T) {
	s.runError(t, "Ошибка при добавлении host-binding, namespace не существует", bodies.UpsertErrorNamespaceNotExist(), codes.NotFound, "pass existing namespace name")
}

func (s *GRPCSuiteUpsertHB) TestUpsertErrorNamespaceEmpty(t provider.T) {
	s.runError(t, "Ошибка при добавлении host-binding, namespace пуст", bodies.UpsertErrorNamespaceEmpty(), codes.InvalidArgument, "at least one of UID or Name and Namespace must be set")
}

func (s *GRPCSuiteUpsertHB) TestUpsertErrorNamespaceLeadingSpace(t provider.T) {
	s.runError(t, "Ошибка при добавлении host-binding, пробел в начале имени namespace", bodies.UpsertErrorNamespaceLeadingSpace(), codes.InvalidArgument, "Namespace: must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteUpsertHB) TestUpsertErrorNamespaceTrailingSpace(t provider.T) {
	s.runError(t, "Ошибка при добавлении host-binding, пробел в конце имени namespace", bodies.UpsertErrorNamespaceTrailingSpace(), codes.InvalidArgument, "Namespace: must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteUpsertHB) TestUpsertErrorNamespaceWithSpaces(t provider.T) {
	s.runError(t, "Ошибка при добавлении host-binding, namespace с пробелами", bodies.UpsertErrorNamespaceWithSpaces(), codes.InvalidArgument, "Namespace: must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteUpsertHB) TestUpsertErrorNamespaceMissing(t provider.T) {
	s.runError(t, "Ошибка при добавлении host-binding, namespace отсутствует как параметр", bodies.UpsertErrorNamespaceMissing(), codes.InvalidArgument, "at least one of UID or Name and Namespace must be set")
}

func (s *GRPCSuiteUpsertHB) TestUpsertErrorHostNotExistInNamespace(t provider.T) {
	s.runError(t, "Ошибка при добавлении host-binding, host не существует в указанном ns", bodies.UpsertErrorHostNotExistInNamespace(), codes.NotFound, "pass existing host name and namespace")
}

func (s *GRPCSuiteUpsertHB) TestUpsertErrorHostNotInNamespaceButExists(t provider.T) {
	s.runError(t, "Ошибка при добавлении host-binding, в указанном ns нет такого хоста (сам хост существует)", bodies.UpsertErrorHostNotInNamespaceButExists(), codes.NotFound, "pass existing host name and namespace")
}

func (s *GRPCSuiteUpsertHB) TestUpsertErrorHostNameEmpty(t provider.T) {
	s.runError(t, "Ошибка при добавлении host-binding, host name пуст", bodies.UpsertErrorHostNameEmpty(), codes.InvalidArgument, "Host", "Name", "resource name is required")
}

func (s *GRPCSuiteUpsertHB) TestUpsertErrorHostEmpty(t provider.T) {
	s.runError(t, "Ошибка при добавлении host-binding, host пуст (ns и name)", bodies.UpsertErrorHostEmpty(), codes.InvalidArgument, "Host", "Name", "Namespace", "resource name is required")
}

func (s *GRPCSuiteUpsertHB) TestUpsertErrorHostNameMissing(t provider.T) {
	s.runError(t, "Ошибка при добавлении host-binding, в host name отсутствует как параметр", bodies.UpsertErrorHostNameMissing(), codes.InvalidArgument, "Host", "Name", "resource name is required")
}

func (s *GRPCSuiteUpsertHB) TestUpsertErrorHostNamespaceMissing(t provider.T) {
	s.runError(t, "Ошибка при добавлении host-binding, в host ns отсутствует как параметр", bodies.UpsertErrorHostNamespaceMissing(), codes.InvalidArgument, "Host", "Namespace", "resource namespace is required")
}

func (s *GRPCSuiteUpsertHB) TestUpsertErrorHostEmptyObject(t provider.T) {
	s.runError(t, "Ошибка при добавлении host-binding, host пустой объект", bodies.UpsertErrorHostEmptyObject(), codes.InvalidArgument, "Host", "Namespace", "resource namespace is required")
}

func (s *GRPCSuiteUpsertHB) TestUpsertErrorHostMissing(t provider.T) {
	s.runError(t, "Ошибка при добавлении host-binding, host отсутствует как параметр", bodies.UpsertErrorHostMissing(), codes.InvalidArgument, "Host", "Namespace", "resource namespace is required")
}

func (s *GRPCSuiteUpsertHB) TestUpsertErrorAGEmptyObject(t provider.T) {
	s.runError(t, "Ошибка при добавлении host-binding, ag пустой объект", bodies.UpsertErrorAGEmptyObject(), codes.InvalidArgument, "AddressGroup", "Namespace", "Name", "resource namespace is required")
}

func (s *GRPCSuiteUpsertHB) TestUpsertErrorAGMissing(t provider.T) {
	s.runError(t, "Ошибка при добавлении host-binding, ag отсутствует как параметр", bodies.UpsertErrorAGMissing(), codes.InvalidArgument, "AddressGroup", "Namespace", "Name", "resource namespace is required")
}

func (s *GRPCSuiteUpsertHB) TestUpsertErrorAGNamespaceMissing(t provider.T) {
	s.runError(t, "Ошибка при добавлении host-binding, в ag ns отсутствует как параметр", bodies.UpsertErrorAGNamespaceMissing(), codes.InvalidArgument, "AddressGroup", "Namespace", "Name", "resource namespace is required")
}

func (s *GRPCSuiteUpsertHB) TestUpsertErrorAGNameMissing(t provider.T) {
	s.runError(t, "Ошибка при добавлении host-binding, в ag name отсутствует как параметр", bodies.UpsertErrorAGNameMissing(), codes.InvalidArgument, "AddressGroup", "Name", "resource name is required")
}

func (s *GRPCSuiteUpsertHB) TestUpsertErrorAGEmpty(t provider.T) {
	s.runError(t, "Ошибка при добавлении host-binding, ag пуст (ns и name)", bodies.UpsertErrorAGEmpty(), codes.InvalidArgument, "AddressGroup", "Name", "Namespace", "resource name is required")
}

func (s *GRPCSuiteUpsertHB) TestUpsertErrorAGNameEmpty(t provider.T) {
	s.runError(t, "Ошибка при добавлении host-binding, ag name пуст", bodies.UpsertErrorAGNameEmpty(), codes.InvalidArgument, "AddressGroup", "Name", "resource name is required")
}

func (s *GRPCSuiteUpsertHB) TestUpsertErrorAGNotInNamespaceButExists(t provider.T) {
	s.runError(t, "Ошибка при добавлении host-binding, в указанном ag нет такого ag (сам ag существует)", bodies.UpsertErrorAGNotInNamespaceButExists(), codes.NotFound, "pass existing address group name and namespace")
}

func (s *GRPCSuiteUpsertHB) TestUpsertErrorAGNotExist(t provider.T) {
	s.runError(t, "Ошибка при добавлении host-binding, ag не существует в указанном ag", bodies.UpsertErrorAGNotExist(), codes.NotFound, "pass existing address group name and namespace")
}

func (s *GRPCSuiteUpsertHB) TestUpsertErrorSpecEmpty(t provider.T) {
	s.runError(t, "Ошибка при добавлении host-binding, spec пустой объект", bodies.UpsertErrorSpecEmpty(), codes.InvalidArgument, "AddressGroup", "Host", "resource name is required")
}
