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

type GRPCSuiteUpsertNB struct {
	suite.Suite
	TestStruct foundation.ToolsForTest
	client     grpcClient.SGroupsNetworkBindingAPIClient
}

func (s *GRPCSuiteUpsertNB) BeforeAll(t provider.T) {
	s.TestStruct.NetworkBindingAPI = grpcClient.NewSGroupsNetworkBindingAPIClient(connection.GRPC(t))
}

func (s *GRPCSuiteUpsertNB) runSuccess(t provider.T, story string, body *foundation.TestCaseBodyNBUpsert, name string, exp networkBindingExpectation) {
	t.Parallel()
	t.Story(story)
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		resp, err := s.TestStruct.FoundationTestsNBUpsert(body)

		sCtx.NewStep(steps.AssertStep)
		sCtx.Require().NoError(err)
		sCtx.Require().NotNil(resp)
		sCtx.Require().Len(resp.NetworkBindings, 1)
		assertNetworkBinding(sCtx, resp.NetworkBindings[0], name, exp)
	})
}

func (s *GRPCSuiteUpsertNB) runError(t provider.T, story string, body *foundation.TestCaseBodyNBUpsert, code codes.Code, substrings ...string) {
	t.Parallel()
	t.Story(story)
	t.WithNewStep(steps.WithNewStep, func(sCtx provider.StepCtx) {
		_, err := s.TestStruct.FoundationTestsNBUpsert(body)

		sCtx.NewStep(steps.AssertStep)
		utils.AssertGRPCError(sCtx, err, code, substrings...)
	})
}

func (s *GRPCSuiteUpsertNB) TestUpsertAdd(t provider.T) {
	s.runSuccess(t, "Добавление нового networkBindings", bodies.UpsertAdd(), "add-nb-1", networkBindingExpectation{
		namespace: "namespace-1", labels: map[string]string{"add": "success"}, annotations: map[string]string{"add": "success"},
		displayName: "new network binding 1", comment: "new network binding", description: "add success",
		addressGroupName: "ag-6", addressGroupNamespace: "namespace-1", networkName: "nw-6", networkNamespace: "namespace-1",
	})
}

func (s *GRPCSuiteUpsertNB) TestUpsertEdit(t provider.T) {
	s.runSuccess(t, "Редактирование networkBindings", bodies.UpsertEdit(), "nb-3", networkBindingExpectation{
		namespace: "namespace-1", uid: "507404f0-6ddc-4d15-b7bd-6c59bdff1d24",
		labels: map[string]string{"edit": "success"}, annotations: map[string]string{"edit": "success"},
		displayName: "edit network binding 1", comment: "edited network binding", description: "edit success",
		addressGroupName: "ag-3", addressGroupNamespace: "namespace-1", networkName: "nw-3", networkNamespace: "namespace-1",
	})
}

func (s *GRPCSuiteUpsertNB) TestUpsertAddSameNameDifferentNamespace(t provider.T) {
	s.runSuccess(t, "Добавление networkBindings с существующим именем, но к другому namespace", bodies.UpsertAddSameNameDifferentNamespace(), "nb-0", networkBindingExpectation{
		namespace: "namespace-1", labels: map[string]string{"add": "success"}, annotations: map[string]string{"add": "success"},
		displayName: "new network binding 1", comment: "new network binding", description: "add success",
		addressGroupName: "ag-6", addressGroupNamespace: "namespace-1", networkName: "nw-7", networkNamespace: "namespace-1",
	})
}

func (s *GRPCSuiteUpsertNB) TestUpsertErrorDisplayNameTooLong(t provider.T) {
	s.runError(t, "Ошибка при добавлении networkBindings, displayName > 63 символов", bodies.UpsertErrorDisplayNameTooLong(), codes.InvalidArgument, "DisplayName", "display name must be no longer than 63 characters")
}

func (s *GRPCSuiteUpsertNB) TestUpsertErrorEditAGAndNetworkMissing(t provider.T) {
	s.runError(t, "Ошибка при редактировании networkBindings, ag и nw отсутствуют", bodies.UpsertErrorEditAGAndNetworkMissing(), codes.InvalidArgument, "AddressGroup", "Network", "resource name is required")
}

func (s *GRPCSuiteUpsertNB) TestUpsertErrorEditNetworkImmutable(t provider.T) {
	s.runError(t, "Ошибка при редактировании networkBindings, нельзя редактировать nw", bodies.UpsertErrorEditNetworkImmutable(), codes.InvalidArgument, "network cannot be updated")
}

func (s *GRPCSuiteUpsertNB) TestUpsertErrorEditNetworkImmutableDifferentNamespace(t provider.T) {
	s.runError(t, "Ошибка при редактировании networkBindings, нельзя редактировать nw из другого ns", bodies.UpsertErrorEditNetworkImmutableDifferentNamespace(), codes.InvalidArgument, "pass network and address group in the same namespace")
}

func (s *GRPCSuiteUpsertNB) TestUpsertErrorEditNetworkMissing(t provider.T) {
	s.runError(t, "Ошибка при редактировании networkBindings, nw отсутствует", bodies.UpsertErrorEditNetworkMissing(), codes.InvalidArgument, "Network", "resource name is required", "resource namespace is required")
}

func (s *GRPCSuiteUpsertNB) TestUpsertErrorEditAGImmutable(t provider.T) {
	s.runError(t, "Ошибка при редактировании networkBindings, нельзя редактировать ag", bodies.UpsertErrorEditAGImmutable(), codes.InvalidArgument, "ag cannot be updated")
}

func (s *GRPCSuiteUpsertNB) TestUpsertErrorEditAGImmutableDifferentNamespace(t provider.T) {
	s.runError(t, "Ошибка при редактировании networkBindings, нельзя редактировать ag из другого ns", bodies.UpsertErrorEditAGImmutableDifferentNamespace(), codes.InvalidArgument, "pass network and address group in the same namespace")
}

func (s *GRPCSuiteUpsertNB) TestUpsertErrorEditAGMissing(t provider.T) {
	s.runError(t, "Ошибка при редактировании networkBindings, ag отсутствует", bodies.UpsertErrorEditAGMissing(), codes.InvalidArgument, "AddressGroup", "resource name is required", "resource namespace is required")
}

func (s *GRPCSuiteUpsertNB) TestUpsertErrorBindingAlreadyExists(t provider.T) {
	s.runError(t, "Ошибка при добавлении networkBindings, такая привязка уже существует (nw-host)", bodies.UpsertErrorBindingAlreadyExists(), codes.FailedPrecondition, "binding (address_group, network) must be unique")
}

func (s *GRPCSuiteUpsertNB) TestUpsertErrorAddDuplicateNameNamespace(t provider.T) {
	s.runError(t, "Ошибка при добавлении networkBindings, nb  с таким name+namespace уже существует", bodies.UpsertErrorAddDuplicateNameNamespace(), codes.FailedPrecondition, "binding (address_group, network) must be unique")
}

func (s *GRPCSuiteUpsertNB) TestUpsertErrorNetworkNamespaceMismatch(t provider.T) {
	s.runError(t, "Ошибка при добавлении networkBindings, nw ns с не совпадает с nb ns", bodies.UpsertErrorNetworkNamespaceMismatch(), codes.InvalidArgument, "pass network and address group in the same namespace")
}

func (s *GRPCSuiteUpsertNB) TestUpsertErrorAGNamespaceMismatch(t provider.T) {
	s.runError(t, "Ошибка при добавлении networkBindings, ag ns с не совпадает с nb ns", bodies.UpsertErrorAGNamespaceMismatch(), codes.InvalidArgument, "pass network and address group in the same namespace")
}

func (s *GRPCSuiteUpsertNB) TestUpsertErrorNBNamespaceMismatchBoth(t provider.T) {
	s.runError(t, "Ошибка при добавлении networkBindings, nb ns не совпадает с ag и nw ns", bodies.UpsertErrorNBNamespaceMismatchBoth(), codes.InvalidArgument, "pass network binding namespace that matches network namespace")
}

func (s *GRPCSuiteUpsertNB) TestUpsertErrorEditLongUID(t provider.T) {
	s.runError(t, "Ошибка при редактировании networkBindings, uid длинный", bodies.UpsertErrorEditLongUID(), codes.InvalidArgument, "invalid UUID length")
}

func (s *GRPCSuiteUpsertNB) TestUpsertErrorEditShortUID(t provider.T) {
	s.runError(t, "Ошибка при редактировании networkBindings, uid короткий", bodies.UpsertErrorEditShortUID(), codes.InvalidArgument, "invalid UUID length")
}

func (s *GRPCSuiteUpsertNB) TestUpsertErrorEditUIDMismatch(t provider.T) {
	s.runError(t, "Ошибка при редактировании networkBindings, uid не соответствует указанному nb", bodies.UpsertErrorEditUIDMismatch(), codes.FailedPrecondition, "existing uid AND matching name")
}

func (s *GRPCSuiteUpsertNB) TestUpsertErrorEditNameUIDNotExistInNamespace(t provider.T) {
	s.runError(t, "Ошибка при редактировании networkBindings, name+uid не существует в указанном namespace", bodies.UpsertErrorEditNameUIDNotExistInNamespace(), codes.FailedPrecondition, "existing uid AND matching name")
}

func (s *GRPCSuiteUpsertNB) TestUpsertErrorEditEmptyName(t provider.T) {
	s.runError(t, "Ошибка при редактировании networkBindings, name пуст", bodies.UpsertErrorEditEmptyName(), codes.InvalidArgument, "set name for both insert")
}

func (s *GRPCSuiteUpsertNB) TestUpsertErrorEditNameMissing(t provider.T) {
	s.runError(t, "Ошибка при редактировании networkBindings, name отсутствует как параметр", bodies.UpsertErrorEditNameMissing(), codes.InvalidArgument, "set name for both insert")
}

func (s *GRPCSuiteUpsertNB) TestUpsertErrorEditEmptyNamespace(t provider.T) {
	s.runError(t, "Ошибка при редактировании networkBindings, namespace пуст", bodies.UpsertErrorEditEmptyNamespace(), codes.InvalidArgument, "pass existing namespace for both insert")
}

func (s *GRPCSuiteUpsertNB) TestUpsertErrorEditNamespaceMissing(t provider.T) {
	s.runError(t, "Ошибка при редактировании networkBindings, namespace отсутствует как параметр", bodies.UpsertErrorEditNamespaceMissing(), codes.InvalidArgument, "pass existing namespace for both insert")
}

func (s *GRPCSuiteUpsertNB) TestUpsertErrorAddEmptyName(t provider.T) {
	s.runError(t, "Ошибка при добавлении networkBindings, name пуст", bodies.UpsertErrorAddEmptyName(), codes.InvalidArgument, "at least one of UID or Name and Namespace must be set")
}

func (s *GRPCSuiteUpsertNB) TestUpsertErrorNameCyrillic(t provider.T) {
	s.runError(t, "Ошибка при добавлении networkBindings, name на кириллице", bodies.UpsertErrorNameCyrillic(), codes.InvalidArgument, "Name: must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteUpsertNB) TestUpsertErrorNameUpperCase(t provider.T) {
	s.runError(t, "Ошибка при добавлении networkBindings, name капсом", bodies.UpsertErrorNameUpperCase(), codes.InvalidArgument, "Name: must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteUpsertNB) TestUpsertErrorNameSpecialCharsOnly(t provider.T) {
	s.runError(t, "Ошибка при добавлении networkBindings, name из спец символов", bodies.UpsertErrorNameSpecialCharsOnly(), codes.InvalidArgument, "Name: must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteUpsertNB) TestUpsertErrorNameMixedCase(t provider.T) {
	s.runError(t, "Ошибка при добавлении networkBindings, name капс+строчные", bodies.UpsertErrorNameMixedCase(), codes.InvalidArgument, "Name: must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteUpsertNB) TestUpsertErrorNameCyrillicLatin(t provider.T) {
	s.runError(t, "Ошибка при добавлении networkBindings, name на кириллице+латинице", bodies.UpsertErrorNameCyrillicLatin(), codes.InvalidArgument, "Name: must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteUpsertNB) TestUpsertErrorNameCyrillicDigits(t provider.T) {
	s.runError(t, "Ошибка при добавлении networkBindings, name на кириллице+числа", bodies.UpsertErrorNameCyrillicDigits(), codes.InvalidArgument, "Name: must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteUpsertNB) TestUpsertErrorNameStartsWithHyphen(t provider.T) {
	s.runError(t, "Ошибка при добавлении networkBindings, name начинается с тире", bodies.UpsertErrorNameStartsWithHyphen(), codes.InvalidArgument, "Name: must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteUpsertNB) TestUpsertErrorNameOnlyHyphen(t provider.T) {
	s.runError(t, "Ошибка при добавлении networkBindings, name состоит только из тире", bodies.UpsertErrorNameOnlyHyphen(), codes.InvalidArgument, "Name: must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteUpsertNB) TestUpsertErrorNameEndsWithHyphen(t provider.T) {
	s.runError(t, "Ошибка при добавлении networkBindings, name заканчивается тире", bodies.UpsertErrorNameEndsWithHyphen(), codes.InvalidArgument, "Name: must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteUpsertNB) TestUpsertErrorNameCyrillicSpecialChars(t provider.T) {
	s.runError(t, "Ошибка при добавлении networkBindings, name на кириллице+спец.символы", bodies.UpsertErrorNameCyrillicSpecialChars(), codes.InvalidArgument, "Name: must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteUpsertNB) TestUpsertErrorNameTooLong(t provider.T) {
	s.runError(t, "Ошибка при добавлении networkBindings, name > 63 символов", bodies.UpsertErrorNameTooLong(), codes.InvalidArgument, "Name: must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteUpsertNB) TestUpsertErrorNameLeadingSpace(t provider.T) {
	s.runError(t, "Ошибка при добавлении networkBindings, пробел в начале name", bodies.UpsertErrorNameLeadingSpace(), codes.InvalidArgument, "Name: must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteUpsertNB) TestUpsertErrorNameTrailingSpace(t provider.T) {
	s.runError(t, "Ошибка при добавлении networkBindings, пробел в конце name", bodies.UpsertErrorNameTrailingSpace(), codes.InvalidArgument, "Name: must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteUpsertNB) TestUpsertErrorNameWithSpaces(t provider.T) {
	s.runError(t, "Ошибка при добавлении networkBindings, name с пробелами", bodies.UpsertErrorNameWithSpaces(), codes.InvalidArgument, "Name: must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteUpsertNB) TestUpsertErrorNamespaceNotExist(t provider.T) {
	s.runError(t, "Ошибка при добавлении networkBindings, namespace не существует", bodies.UpsertErrorNamespaceNotExist(), codes.NotFound, "pass existing namespace name")
}

func (s *GRPCSuiteUpsertNB) TestUpsertErrorNamespaceEmpty(t provider.T) {
	s.runError(t, "Ошибка при добавлении networkBindings, namespace пуст", bodies.UpsertErrorNamespaceEmpty(), codes.InvalidArgument, "at least one of UID or Name and Namespace must be set")
}

func (s *GRPCSuiteUpsertNB) TestUpsertErrorNamespaceLeadingSpace(t provider.T) {
	s.runError(t, "Ошибка при добавлении networkBindings, пробел в начале имени namespace", bodies.UpsertErrorNamespaceLeadingSpace(), codes.InvalidArgument, "Namespace: must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteUpsertNB) TestUpsertErrorNamespaceTrailingSpace(t provider.T) {
	s.runError(t, "Ошибка при добавлении networkBindings, пробел в конце имени namespace", bodies.UpsertErrorNamespaceTrailingSpace(), codes.InvalidArgument, "Namespace: must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteUpsertNB) TestUpsertErrorNamespaceWithSpaces(t provider.T) {
	s.runError(t, "Ошибка при добавлении networkBindings, namespace с пробелами", bodies.UpsertErrorNamespaceWithSpaces(), codes.InvalidArgument, "Namespace: must consist of lower-case alphanumeric characters")
}

func (s *GRPCSuiteUpsertNB) TestUpsertErrorNamespaceMissing(t provider.T) {
	s.runError(t, "Ошибка при добавлении networkBindings, namespace отсутствует как параметр", bodies.UpsertErrorNamespaceMissing(), codes.InvalidArgument, "at least one of UID or Name and Namespace must be set")
}

func (s *GRPCSuiteUpsertNB) TestUpsertErrorNetworkNotExistInNamespace(t provider.T) {
	s.runError(t, "Ошибка при добавлении networkBindings, nw не существует в указанном ns", bodies.UpsertErrorNetworkNotExistInNamespace(), codes.NotFound, "pass existing network name and namespace")
}

func (s *GRPCSuiteUpsertNB) TestUpsertErrorNetworkNotInNamespaceButExists(t provider.T) {
	s.runError(t, "Ошибка при добавлении networkBindings, в указанном ns нет такого nw (сам nw существует)", bodies.UpsertErrorNetworkNotInNamespaceButExists(), codes.NotFound, "pass existing network name and namespace")
}

func (s *GRPCSuiteUpsertNB) TestUpsertErrorNetworkNameEmpty(t provider.T) {
	s.runError(t, "Ошибка при добавлении networkBindings, nw name пуст", bodies.UpsertErrorNetworkNameEmpty(), codes.InvalidArgument, "Network", "Name", "resource name is required")
}

func (s *GRPCSuiteUpsertNB) TestUpsertErrorNetworkEmpty(t provider.T) {
	s.runError(t, "Ошибка при добавлении networkBindings, nw пуст (ns и name)", bodies.UpsertErrorNetworkEmpty(), codes.InvalidArgument, "Network", "Name", "Namespace", "resource name is required")
}

func (s *GRPCSuiteUpsertNB) TestUpsertErrorNetworkNameMissing(t provider.T) {
	s.runError(t, "Ошибка при добавлении networkBindings, в nw name отсутствует как параметр", bodies.UpsertErrorNetworkNameMissing(), codes.InvalidArgument, "Network", "Name", "resource name is required")
}

func (s *GRPCSuiteUpsertNB) TestUpsertErrorNetworkNamespaceMissing(t provider.T) {
	s.runError(t, "Ошибка при добавлении networkBindings, в nw ns отсутствует как параметр", bodies.UpsertErrorNetworkNamespaceMissing(), codes.InvalidArgument, "Network", "Namespace", "resource namespace is required")
}

func (s *GRPCSuiteUpsertNB) TestUpsertErrorNetworkEmptyObject(t provider.T) {
	s.runError(t, "Ошибка при добавлении networkBindings, nw пустой объект", bodies.UpsertErrorNetworkEmptyObject(), codes.InvalidArgument, "Network", "Namespace", "resource namespace is required")
}

func (s *GRPCSuiteUpsertNB) TestUpsertErrorNetworkMissing(t provider.T) {
	s.runError(t, "Ошибка при добавлении networkBindings, nw отсутствует как параметр", bodies.UpsertErrorNetworkMissing(), codes.InvalidArgument, "Network", "Namespace", "resource namespace is required")
}

func (s *GRPCSuiteUpsertNB) TestUpsertErrorAGEmptyObject(t provider.T) {
	s.runError(t, "Ошибка при добавлении networkBindings, ag пустой объект", bodies.UpsertErrorAGEmptyObject(), codes.InvalidArgument, "AddressGroup", "Namespace", "Name", "resource namespace is required")
}

func (s *GRPCSuiteUpsertNB) TestUpsertErrorAGMissing(t provider.T) {
	s.runError(t, "Ошибка при добавлении networkBindings, ag отсутствует как параметр", bodies.UpsertErrorAGMissing(), codes.InvalidArgument, "AddressGroup", "Namespace", "Name", "resource namespace is required")
}

func (s *GRPCSuiteUpsertNB) TestUpsertErrorAGNamespaceMissing(t provider.T) {
	s.runError(t, "Ошибка при добавлении networkBindings, в ag ns отсутствует как параметр", bodies.UpsertErrorAGNamespaceMissing(), codes.InvalidArgument, "AddressGroup", "Namespace", "Name", "resource namespace is required")
}

func (s *GRPCSuiteUpsertNB) TestUpsertErrorAGNameMissing(t provider.T) {
	s.runError(t, "Ошибка при добавлении networkBindings, в ag name отсутствует как параметр", bodies.UpsertErrorAGNameMissing(), codes.InvalidArgument, "AddressGroup", "Name", "resource name is required")
}

func (s *GRPCSuiteUpsertNB) TestUpsertErrorAGEmpty(t provider.T) {
	s.runError(t, "Ошибка при добавлении networkBindings, ag пуст (ns и name)", bodies.UpsertErrorAGEmpty(), codes.InvalidArgument, "AddressGroup", "Name", "Namespace", "resource name is required")
}

func (s *GRPCSuiteUpsertNB) TestUpsertErrorAGNameEmpty(t provider.T) {
	s.runError(t, "Ошибка при добавлении networkBindings, ag name пуст", bodies.UpsertErrorAGNameEmpty(), codes.InvalidArgument, "AddressGroup", "Name", "resource name is required")
}

func (s *GRPCSuiteUpsertNB) TestUpsertErrorAGNotInNamespaceButExists(t provider.T) {
	s.runError(t, "Ошибка при добавлении networkBindings, в указанном ag нет такого ag (сам ag существует)", bodies.UpsertErrorAGNotInNamespaceButExists(), codes.NotFound, "pass existing address group name and namespace")
}

func (s *GRPCSuiteUpsertNB) TestUpsertErrorAGNotExist(t provider.T) {
	s.runError(t, "Ошибка при добавлении networkBindings, ag не существует в указанном ns", bodies.UpsertErrorAGNotExist(), codes.NotFound, "pass existing address group name and namespace")
}

func (s *GRPCSuiteUpsertNB) TestUpsertErrorSpecEmpty(t provider.T) {
	s.runError(t, "Ошибка при добавлении networkBindings, spec пустой объект", bodies.UpsertErrorSpecEmpty(), codes.InvalidArgument, "AddressGroup", "Network", "resource name is required")
}
