package bodies

import (
	"strings"

	sgroupsv1 "HBF-tests/pkg/api/sgroups/v1"
	"HBF-tests/tests/common/foundation"

	common "github.com/PRO-Robotech/sgroups-proto/pkg/api/common"
)

// tooLongNBName 64-символьная строка из строчных латинских букв для проверки ограничения длины в 63 символа
var tooLongNBName = strings.Repeat("a", 64)

func upsertNBBody(testName string, nbs ...*sgroupsv1.NetworkBinding) *foundation.TestCaseBodyNBUpsert {
	return &foundation.TestCaseBodyNBUpsert{
		TestName: testName,
		Req:      sgroupsv1.NetworkBindingReq_Upsert{NetworkBindings: nbs},
	}
}

func nbEntry(metadata *common.Metadata, spec *sgroupsv1.NetworkBinding_Spec) *sgroupsv1.NetworkBinding {
	return &sgroupsv1.NetworkBinding{Metadata: metadata, Spec: spec}
}

// errNameSpec Spec, общий для большинства негативных кейсов добавления network-binding по имени
func errNameSpec() *sgroupsv1.NetworkBinding_Spec {
	return &sgroupsv1.NetworkBinding_Spec{
		AddressGroup: ri("ag-3", "namespace-1"),
		Network:      ri("nw-7", "namespace-1"),
	}
}

// upsertErrName Негативный кейс с невалидным именем network-binding (namespace-1, стандартный errNameSpec)
func upsertErrName(testName, name string) *foundation.TestCaseBodyNBUpsert {
	return upsertNBBody(testName, nbEntry(&common.Metadata{Name: name, Namespace: "namespace-1"}, errNameSpec()))
}

// UpsertAdd Добавление нового networkBindings
func UpsertAdd() *foundation.TestCaseBodyNBUpsert {
	return upsertNBBody("Добавление нового networkBindings", nbEntry(
		&common.Metadata{
			Name:        "add-nb-1",
			Namespace:   "namespace-1",
			Labels:      map[string]string{"add": "success"},
			Annotations: map[string]string{"add": "success"},
		},
		&sgroupsv1.NetworkBinding_Spec{
			Comment: "new network binding", Description: "add success", DisplayName: "new network binding 1",
			AddressGroup: ri("ag-6", "namespace-1"), Network: ri("nw-6", "namespace-1"),
		},
	))
}

// UpsertEdit Редактирование networkBindings
func UpsertEdit() *foundation.TestCaseBodyNBUpsert {
	return upsertNBBody("Редактирование networkBindings", nbEntry(
		&common.Metadata{
			Name:        "nb-3",
			Namespace:   "namespace-1",
			Uid:         "507404f0-6ddc-4d15-b7bd-6c59bdff1d24",
			Labels:      map[string]string{"edit": "success"},
			Annotations: map[string]string{"edit": "success"},
		},
		&sgroupsv1.NetworkBinding_Spec{
			Comment: "edited network binding", Description: "edit success", DisplayName: "edit network binding 1",
			AddressGroup: ri("ag-3", "namespace-1"), Network: ri("nw-3", "namespace-1"),
		},
	))
}

// UpsertAddSameNameDifferentNamespace Добавление networkBindings с существующим именем, но к другому namespace
func UpsertAddSameNameDifferentNamespace() *foundation.TestCaseBodyNBUpsert {
	return upsertNBBody("Добавление networkBindings с существующим именем, но к другому namespace", nbEntry(
		&common.Metadata{
			Name:        "nb-0",
			Namespace:   "namespace-1",
			Labels:      map[string]string{"add": "success"},
			Annotations: map[string]string{"add": "success"},
		},
		&sgroupsv1.NetworkBinding_Spec{
			Comment: "new network binding", Description: "add success", DisplayName: "new network binding 1",
			AddressGroup: ri("ag-6", "namespace-1"), Network: ri("nw-7", "namespace-1"),
		},
	))
}

// UpsertErrorDisplayNameTooLong Ошибка при добавлении networkBindings, displayName > 63 символов
func UpsertErrorDisplayNameTooLong() *foundation.TestCaseBodyNBUpsert {
	return upsertNBBody("Ошибка при добавлении networkBindings, displayName > 63 символов", nbEntry(
		&common.Metadata{Name: "err-display", Namespace: "namespace-1"},
		&sgroupsv1.NetworkBinding_Spec{
			DisplayName: tooLongNBName, AddressGroup: ri("ag-1", "namespace-1"), Network: ri("nw-7", "namespace-1"),
		},
	))
}

// editSpecBase Spec, общий для большинства негативных кейсов редактирования nb-3
func editSpecBase(ag, network *common.ResourceIdentifier) *sgroupsv1.NetworkBinding_Spec {
	return &sgroupsv1.NetworkBinding_Spec{
		Comment: "edited network binding", Description: "edit success", DisplayName: "edit network binding 1",
		AddressGroup: ag, Network: network,
	}
}

func editNB3(testName string, uid string, ag, network *common.ResourceIdentifier) *foundation.TestCaseBodyNBUpsert {
	return upsertNBBody(testName, nbEntry(
		&common.Metadata{
			Name: "nb-3", Namespace: "namespace-1", Uid: uid,
			Labels: map[string]string{"edit": "success"}, Annotations: map[string]string{"edit": "success"},
		},
		editSpecBase(ag, network),
	))
}

// UpsertErrorEditAGAndNetworkMissing Ошибка при редактировании networkBindings, ag и nw отсутствуют
func UpsertErrorEditAGAndNetworkMissing() *foundation.TestCaseBodyNBUpsert {
	return upsertNBBody("Ошибка при редактировании networkBindings, ag и nw отсутствуют", nbEntry(
		&common.Metadata{
			Name: "nb-3", Namespace: "namespace-1", Uid: "507404f0-6ddc-4d15-b7bd-6c59bdff1d24",
			Labels: map[string]string{"edit": "success"}, Annotations: map[string]string{"edit": "success"},
		},
		&sgroupsv1.NetworkBinding_Spec{Comment: "edited network binding", Description: "edit success", DisplayName: "edit network binding 1"},
	))
}

// UpsertErrorEditNetworkImmutable Ошибка при редактировании networkBindings, нельзя редактировать nw
func UpsertErrorEditNetworkImmutable() *foundation.TestCaseBodyNBUpsert {
	return editNB3("Ошибка при редактировании networkBindings, нельзя редактировать nw", "507404f0-6ddc-4d15-b7bd-6c59bdff1d24", ri("ag-3", "namespace-1"), ri("nw-1", "namespace-1"))
}

// UpsertErrorEditNetworkImmutableDifferentNamespace Ошибка при редактировании networkBindings, нельзя редактировать nw из другого ns
func UpsertErrorEditNetworkImmutableDifferentNamespace() *foundation.TestCaseBodyNBUpsert {
	return editNB3("Ошибка при редактировании networkBindings, нельзя редактировать nw из другого ns", "507404f0-6ddc-4d15-b7bd-6c59bdff1d24", ri("ag-3", "namespace-1"), ri("nw-2", "namespace-2"))
}

// UpsertErrorEditNetworkMissing Ошибка при редактировании networkBindings, nw отсутствует
func UpsertErrorEditNetworkMissing() *foundation.TestCaseBodyNBUpsert {
	return upsertNBBody("Ошибка при редактировании networkBindings, nw отсутствует", nbEntry(
		&common.Metadata{
			Name: "nb-3", Namespace: "namespace-1", Uid: "507404f0-6ddc-4d15-b7bd-6c59bdff1d24",
			Labels: map[string]string{"edit": "success"}, Annotations: map[string]string{"edit": "success"},
		},
		&sgroupsv1.NetworkBinding_Spec{
			Comment: "edited network binding", Description: "edit success", DisplayName: "edit network binding 1",
			AddressGroup: ri("ag-3", "namespace-1"),
		},
	))
}

// UpsertErrorEditAGImmutable Ошибка при редактировании networkBindings, нельзя редактировать ag
func UpsertErrorEditAGImmutable() *foundation.TestCaseBodyNBUpsert {
	return editNB3("Ошибка при редактировании networkBindings, нельзя редактировать ag", "507404f0-6ddc-4d15-b7bd-6c59bdff1d24", ri("ag-6", "namespace-1"), ri("nw-3", "namespace-1"))
}

// UpsertErrorEditAGImmutableDifferentNamespace Ошибка при редактировании networkBindings, нельзя редактировать ag из другого ns
func UpsertErrorEditAGImmutableDifferentNamespace() *foundation.TestCaseBodyNBUpsert {
	return editNB3("Ошибка при редактировании networkBindings, нельзя редактировать ag из другого ns", "507404f0-6ddc-4d15-b7bd-6c59bdff1d24", ri("ag-2", "namespace-2"), ri("nw-3", "namespace-1"))
}

// UpsertErrorEditAGMissing Ошибка при редактировании networkBindings, ag отсутствует
func UpsertErrorEditAGMissing() *foundation.TestCaseBodyNBUpsert {
	return upsertNBBody("Ошибка при редактировании networkBindings, ag отсутствует", nbEntry(
		&common.Metadata{
			Name: "nb-3", Namespace: "namespace-1", Uid: "507404f0-6ddc-4d15-b7bd-6c59bdff1d24",
			Labels: map[string]string{"edit": "success"}, Annotations: map[string]string{"edit": "success"},
		},
		&sgroupsv1.NetworkBinding_Spec{
			Comment: "edited network binding", Description: "edit success", DisplayName: "edit network binding 1",
			Network: ri("nw-3", "namespace-1"),
		},
	))
}

// UpsertErrorBindingAlreadyExists Ошибка при добавлении networkBindings, такая привязка уже существует (nw-host)
func UpsertErrorBindingAlreadyExists() *foundation.TestCaseBodyNBUpsert {
	return upsertNBBody("Ошибка при добавлении networkBindings, такая привязка уже существует (nw-host)", nbEntry(
		&common.Metadata{Name: "nb-1000", Namespace: "namespace-0"},
		&sgroupsv1.NetworkBinding_Spec{AddressGroup: ri("ag-0", "namespace-0"), Network: ri("nw-0", "namespace-0")},
	))
}

// UpsertErrorAddDuplicateNameNamespace Ошибка при добавлении networkBindings, nb  с таким name+namespace уже существует
func UpsertErrorAddDuplicateNameNamespace() *foundation.TestCaseBodyNBUpsert {
	return upsertNBBody("Ошибка при добавлении networkBindings, nb  с таким name+namespace уже существует", nbEntry(
		&common.Metadata{Name: "nb-0", Namespace: "namespace-0"},
		&sgroupsv1.NetworkBinding_Spec{AddressGroup: ri("ag-0", "namespace-0"), Network: ri("nw-0", "namespace-0")},
	))
}

// UpsertErrorNetworkNamespaceMismatch Ошибка при добавлении networkBindings, nw ns с не совпадает с nb ns
func UpsertErrorNetworkNamespaceMismatch() *foundation.TestCaseBodyNBUpsert {
	return upsertNBBody("Ошибка при добавлении networkBindings, nw ns с не совпадает с nb ns", nbEntry(
		&common.Metadata{Name: "err-ns", Namespace: "namespace-1"},
		&sgroupsv1.NetworkBinding_Spec{AddressGroup: ri("ag-1", "namespace-1"), Network: ri("nw-2", "namespace-2")},
	))
}

// UpsertErrorAGNamespaceMismatch Ошибка при добавлении networkBindings, ag ns с не совпадает с nb ns
func UpsertErrorAGNamespaceMismatch() *foundation.TestCaseBodyNBUpsert {
	return upsertNBBody("Ошибка при добавлении networkBindings, ag ns с не совпадает с nb ns", nbEntry(
		&common.Metadata{Name: "err-ns", Namespace: "namespace-1"},
		&sgroupsv1.NetworkBinding_Spec{AddressGroup: ri("ag-2", "namespace-2"), Network: ri("nw-1", "namespace-1")},
	))
}

// UpsertErrorNBNamespaceMismatchBoth Ошибка при добавлении networkBindings, nb ns не совпадает с ag и nw ns
func UpsertErrorNBNamespaceMismatchBoth() *foundation.TestCaseBodyNBUpsert {
	return upsertNBBody("Ошибка при добавлении networkBindings, nb ns не совпадает с ag и nw ns", nbEntry(
		&common.Metadata{Name: "err-ns", Namespace: "namespace-1"},
		&sgroupsv1.NetworkBinding_Spec{AddressGroup: ri("ag-2", "namespace-2"), Network: ri("nw-2", "namespace-2")},
	))
}

// UpsertErrorEditLongUID Ошибка при редактировании networkBindings, uid длинный
func UpsertErrorEditLongUID() *foundation.TestCaseBodyNBUpsert {
	return editNB3("Ошибка при редактировании networkBindings, uid длинный", "efb85252-c9b7-4dd5-b9be-6ce63276795vc", ri("ag-3", "namespace-1"), ri("nw-3", "namespace-1"))
}

// UpsertErrorEditShortUID Ошибка при редактировании networkBindings, uid короткий
func UpsertErrorEditShortUID() *foundation.TestCaseBodyNBUpsert {
	return editNB3("Ошибка при редактировании networkBindings, uid короткий", "efb85252-c9b7-4dd5-b9be-6ce63276795", ri("ag-3", "namespace-1"), ri("nw-3", "namespace-1"))
}

// UpsertErrorEditUIDMismatch Ошибка при редактировании networkBindings, uid не соответствует указанному nb
func UpsertErrorEditUIDMismatch() *foundation.TestCaseBodyNBUpsert {
	return editNB3("Ошибка при редактировании networkBindings, uid не соответствует указанному nb", "1085d231-9a0f-4697-b864-c0a522181911", ri("ag-3", "namespace-1"), ri("nw-3", "namespace-1"))
}

// UpsertErrorEditNameUIDNotExistInNamespace Ошибка при редактировании networkBindings, name+uid не существует в указанном namespace
func UpsertErrorEditNameUIDNotExistInNamespace() *foundation.TestCaseBodyNBUpsert {
	return upsertNBBody("Ошибка при редактировании networkBindings, name+uid не существует в указанном namespace", nbEntry(
		&common.Metadata{Name: "nb-0", Namespace: "namespace-1", Uid: "1085d231-9a0f-4697-b864-c0a522181911"},
		editSpecBase(ri("ag-3", "namespace-1"), ri("nw-3", "namespace-1")),
	))
}

// UpsertErrorEditEmptyName Ошибка при редактировании networkBindings, name пуст
func UpsertErrorEditEmptyName() *foundation.TestCaseBodyNBUpsert {
	return upsertNBBody("Ошибка при редактировании networkBindings, name пуст", nbEntry(
		&common.Metadata{Name: "", Namespace: "namespace-0", Uid: "1085d231-9a0f-4697-b864-c0a522181911"},
		editSpecBase(ri("ag-3", "namespace-1"), ri("nw-3", "namespace-1")),
	))
}

// UpsertErrorEditNameMissing Ошибка при редактировании networkBindings, name отсутствует как параметр
func UpsertErrorEditNameMissing() *foundation.TestCaseBodyNBUpsert {
	return upsertNBBody("Ошибка при редактировании networkBindings, name отсутствует как параметр", nbEntry(
		&common.Metadata{Namespace: "namespace-0", Uid: "1085d231-9a0f-4697-b864-c0a522181911"},
		editSpecBase(ri("ag-3", "namespace-1"), ri("nw-3", "namespace-1")),
	))
}

// UpsertErrorEditEmptyNamespace Ошибка при редактировании networkBindings, namespace пуст
func UpsertErrorEditEmptyNamespace() *foundation.TestCaseBodyNBUpsert {
	return upsertNBBody("Ошибка при редактировании networkBindings, namespace пуст", nbEntry(
		&common.Metadata{Name: "nb-0", Namespace: "", Uid: "1085d231-9a0f-4697-b864-c0a522181911"},
		editSpecBase(ri("ag-3", "namespace-1"), ri("nw-3", "namespace-1")),
	))
}

// UpsertErrorEditNamespaceMissing Ошибка при редактировании networkBindings, namespace отсутствует как параметр
func UpsertErrorEditNamespaceMissing() *foundation.TestCaseBodyNBUpsert {
	return upsertNBBody("Ошибка при редактировании networkBindings, namespace отсутствует как параметр", nbEntry(
		&common.Metadata{Name: "nb-0", Uid: "1085d231-9a0f-4697-b864-c0a522181911"},
		&sgroupsv1.NetworkBinding_Spec{AddressGroup: ri("ag-3", "namespace-1"), Network: ri("nw-3", "namespace-1")},
	))
}

// UpsertErrorAddEmptyName Ошибка при добавлении networkBindings, name пуст
func UpsertErrorAddEmptyName() *foundation.TestCaseBodyNBUpsert {
	return upsertNBBody("Ошибка при добавлении networkBindings, name пуст", nbEntry(
		&common.Metadata{Name: "", Namespace: "namespace-1"},
		&sgroupsv1.NetworkBinding_Spec{AddressGroup: ri("ag-3", "namespace-1"), Network: ri("nw-7", "namespace-1")},
	))
}

// UpsertErrorNameCyrillic Ошибка при добавлении networkBindings, name на кириллице
func UpsertErrorNameCyrillic() *foundation.TestCaseBodyNBUpsert {
	return upsertErrName("Ошибка при добавлении networkBindings, name на кириллице", "биндинг")
}

// UpsertErrorNameUpperCase Ошибка при добавлении networkBindings, name капсом
func UpsertErrorNameUpperCase() *foundation.TestCaseBodyNBUpsert {
	return upsertErrName("Ошибка при добавлении networkBindings, name капсом", "BINDING")
}

// UpsertErrorNameSpecialCharsOnly Ошибка при добавлении networkBindings, name из спец символов
func UpsertErrorNameSpecialCharsOnly() *foundation.TestCaseBodyNBUpsert {
	return upsertErrName("Ошибка при добавлении networkBindings, name из спец символов", "???!!!")
}

// UpsertErrorNameMixedCase Ошибка при добавлении networkBindings, name капс+строчные
func UpsertErrorNameMixedCase() *foundation.TestCaseBodyNBUpsert {
	return upsertErrName("Ошибка при добавлении networkBindings, name капс+строчные", "Binding")
}

// UpsertErrorNameCyrillicLatin Ошибка при добавлении networkBindings, name на кириллице+латинице
func UpsertErrorNameCyrillicLatin() *foundation.TestCaseBodyNBUpsert {
	return upsertErrName("Ошибка при добавлении networkBindings, name на кириллице+латинице", "binдинг")
}

// UpsertErrorNameCyrillicDigits Ошибка при добавлении networkBindings, name на кириллице+числа
func UpsertErrorNameCyrillicDigits() *foundation.TestCaseBodyNBUpsert {
	return upsertErrName("Ошибка при добавлении networkBindings, name на кириллице+числа", "биндинг-1")
}

// UpsertErrorNameStartsWithHyphen Ошибка при добавлении networkBindings, name начинается с тире
func UpsertErrorNameStartsWithHyphen() *foundation.TestCaseBodyNBUpsert {
	return upsertErrName("Ошибка при добавлении networkBindings, name начинается с тире", "-bind")
}

// UpsertErrorNameOnlyHyphen Ошибка при добавлении networkBindings, name состоит только из тире
func UpsertErrorNameOnlyHyphen() *foundation.TestCaseBodyNBUpsert {
	return upsertErrName("Ошибка при добавлении networkBindings, name состоит только из тире", "-")
}

// UpsertErrorNameEndsWithHyphen Ошибка при добавлении networkBindings, name заканчивается тире
func UpsertErrorNameEndsWithHyphen() *foundation.TestCaseBodyNBUpsert {
	return upsertErrName("Ошибка при добавлении networkBindings, name заканчивается тире", "bind-")
}

// UpsertErrorNameCyrillicSpecialChars Ошибка при добавлении networkBindings, name на кириллице+спец.символы
func UpsertErrorNameCyrillicSpecialChars() *foundation.TestCaseBodyNBUpsert {
	return upsertErrName("Ошибка при добавлении networkBindings, name на кириллице+спец.символы", "биндинг?")
}

// UpsertErrorNameTooLong Ошибка при добавлении networkBindings, name > 63 символов
func UpsertErrorNameTooLong() *foundation.TestCaseBodyNBUpsert {
	return upsertErrName("Ошибка при добавлении networkBindings, name > 63 символов", tooLongNBName)
}

// UpsertErrorNameLeadingSpace Ошибка при добавлении networkBindings, пробел в начале name
func UpsertErrorNameLeadingSpace() *foundation.TestCaseBodyNBUpsert {
	return upsertErrName("Ошибка при добавлении networkBindings, пробел в начале name", " bind-1")
}

// UpsertErrorNameTrailingSpace Ошибка при добавлении networkBindings, пробел в конце name
func UpsertErrorNameTrailingSpace() *foundation.TestCaseBodyNBUpsert {
	return upsertErrName("Ошибка при добавлении networkBindings, пробел в конце name", "bind-1 ")
}

// UpsertErrorNameWithSpaces Ошибка при добавлении networkBindings, name с пробелами
func UpsertErrorNameWithSpaces() *foundation.TestCaseBodyNBUpsert {
	return upsertErrName("Ошибка при добавлении networkBindings, name с пробелами", "bind - 1")
}

// UpsertErrorNamespaceNotExist Ошибка при добавлении networkBindings, namespace не существует
func UpsertErrorNamespaceNotExist() *foundation.TestCaseBodyNBUpsert {
	return upsertNBBody("Ошибка при добавлении networkBindings, namespace не существует", nbEntry(
		&common.Metadata{Name: "bind", Namespace: "namespace-11"},
		&sgroupsv1.NetworkBinding_Spec{AddressGroup: ri("ag-3", "namespace-1"), Network: ri("nw-7", "namespace-1")},
	))
}

// UpsertErrorNamespaceEmpty Ошибка при добавлении networkBindings, namespace пуст
func UpsertErrorNamespaceEmpty() *foundation.TestCaseBodyNBUpsert {
	return upsertNBBody("Ошибка при добавлении networkBindings, namespace пуст", nbEntry(
		&common.Metadata{Name: "bind", Namespace: ""},
		&sgroupsv1.NetworkBinding_Spec{AddressGroup: ri("ag-3", "namespace-1"), Network: ri("nw-7", "namespace-1")},
	))
}

// UpsertErrorNamespaceLeadingSpace Ошибка при добавлении networkBindings, пробел в начале имени namespace
func UpsertErrorNamespaceLeadingSpace() *foundation.TestCaseBodyNBUpsert {
	return upsertNBBody("Ошибка при добавлении networkBindings, пробел в начале имени namespace", nbEntry(
		&common.Metadata{Name: "bind", Namespace: " namespace-1"},
		&sgroupsv1.NetworkBinding_Spec{AddressGroup: ri("ag-3", "namespace-1"), Network: ri("nw-7", "namespace-1")},
	))
}

// UpsertErrorNamespaceTrailingSpace Ошибка при добавлении networkBindings, пробел в конце имени namespace
func UpsertErrorNamespaceTrailingSpace() *foundation.TestCaseBodyNBUpsert {
	return upsertNBBody("Ошибка при добавлении networkBindings, пробел в конце имени namespace", nbEntry(
		&common.Metadata{Name: "bind", Namespace: "namespace-1 "},
		&sgroupsv1.NetworkBinding_Spec{AddressGroup: ri("ag-3", "namespace-1"), Network: ri("nw-7", "namespace-1")},
	))
}

// UpsertErrorNamespaceWithSpaces Ошибка при добавлении networkBindings, namespace с пробелами
func UpsertErrorNamespaceWithSpaces() *foundation.TestCaseBodyNBUpsert {
	return upsertNBBody("Ошибка при добавлении networkBindings, namespace с пробелами", nbEntry(
		&common.Metadata{Name: "bind", Namespace: "namespace - 1"},
		&sgroupsv1.NetworkBinding_Spec{AddressGroup: ri("ag-3", "namespace-1"), Network: ri("nw-7", "namespace-1")},
	))
}

// UpsertErrorNamespaceMissing Ошибка при добавлении networkBindings, namespace отсутствует как параметр
func UpsertErrorNamespaceMissing() *foundation.TestCaseBodyNBUpsert {
	return upsertNBBody("Ошибка при добавлении networkBindings, namespace отсутствует как параметр", nbEntry(
		&common.Metadata{Name: "bind"},
		&sgroupsv1.NetworkBinding_Spec{AddressGroup: ri("ag-3", "namespace-1"), Network: ri("nw-7", "namespace-1")},
	))
}

// UpsertErrorNetworkNotExistInNamespace Ошибка при добавлении networkBindings, nw не существует в указанном ns
func UpsertErrorNetworkNotExistInNamespace() *foundation.TestCaseBodyNBUpsert {
	return upsertNBBody("Ошибка при добавлении networkBindings, nw не существует в указанном ns", nbEntry(
		&common.Metadata{Name: "err", Namespace: "namespace-1"},
		&sgroupsv1.NetworkBinding_Spec{AddressGroup: ri("ag-3", "namespace-1"), Network: ri("nw-77", "namespace-1")},
	))
}

// UpsertErrorNetworkNotInNamespaceButExists Ошибка при добавлении networkBindings, в указанном ns нет такого nw (сам nw существует)
func UpsertErrorNetworkNotInNamespaceButExists() *foundation.TestCaseBodyNBUpsert {
	return upsertNBBody("Ошибка при добавлении networkBindings, в указанном ns нет такого nw (сам nw существует)", nbEntry(
		&common.Metadata{Name: "err", Namespace: "namespace-1"},
		&sgroupsv1.NetworkBinding_Spec{AddressGroup: ri("ag-3", "namespace-1"), Network: ri("nw-2", "namespace-1")},
	))
}

// UpsertErrorNetworkNameEmpty Ошибка при добавлении networkBindings, nw name пуст
func UpsertErrorNetworkNameEmpty() *foundation.TestCaseBodyNBUpsert {
	return upsertNBBody("Ошибка при добавлении networkBindings, nw name пуст", nbEntry(
		&common.Metadata{Name: "err", Namespace: "namespace-1"},
		&sgroupsv1.NetworkBinding_Spec{AddressGroup: ri("ag-3", "namespace-1"), Network: ri("", "namespace-1")},
	))
}

// UpsertErrorNetworkEmpty Ошибка при добавлении networkBindings, nw пуст (ns и name)
func UpsertErrorNetworkEmpty() *foundation.TestCaseBodyNBUpsert {
	return upsertNBBody("Ошибка при добавлении networkBindings, nw пуст (ns и name)", nbEntry(
		&common.Metadata{Name: "err", Namespace: "namespace-1"},
		&sgroupsv1.NetworkBinding_Spec{AddressGroup: ri("ag-3", "namespace-1"), Network: ri("", "")},
	))
}

// UpsertErrorNetworkNameMissing Ошибка при добавлении networkBindings, в nw name отсутствует как параметр
func UpsertErrorNetworkNameMissing() *foundation.TestCaseBodyNBUpsert {
	return upsertNBBody("Ошибка при добавлении networkBindings, в nw name отсутствует как параметр", nbEntry(
		&common.Metadata{Name: "err", Namespace: "namespace-1"},
		&sgroupsv1.NetworkBinding_Spec{AddressGroup: ri("ag-3", "namespace-1"), Network: ri("", "namespace-1")},
	))
}

// UpsertErrorNetworkNamespaceMissing Ошибка при добавлении networkBindings, в nw ns отсутствует как параметр
func UpsertErrorNetworkNamespaceMissing() *foundation.TestCaseBodyNBUpsert {
	return upsertNBBody("Ошибка при добавлении networkBindings, в nw ns отсутствует как параметр", nbEntry(
		&common.Metadata{Name: "err", Namespace: "namespace-1"},
		&sgroupsv1.NetworkBinding_Spec{AddressGroup: ri("ag-3", "namespace-1"), Network: ri("nw-1", "")},
	))
}

// UpsertErrorNetworkEmptyObject Ошибка при добавлении networkBindings, nw пустой объект
func UpsertErrorNetworkEmptyObject() *foundation.TestCaseBodyNBUpsert {
	return upsertNBBody("Ошибка при добавлении networkBindings, nw пустой объект", nbEntry(
		&common.Metadata{Name: "err", Namespace: "namespace-1"},
		&sgroupsv1.NetworkBinding_Spec{AddressGroup: ri("ag-3", "namespace-1"), Network: &common.ResourceIdentifier{}},
	))
}

// UpsertErrorNetworkMissing Ошибка при добавлении networkBindings, nw отсутствует как параметр
func UpsertErrorNetworkMissing() *foundation.TestCaseBodyNBUpsert {
	return upsertNBBody("Ошибка при добавлении networkBindings, nw отсутствует как параметр", nbEntry(
		&common.Metadata{Name: "err", Namespace: "namespace-1"},
		&sgroupsv1.NetworkBinding_Spec{AddressGroup: ri("ag-3", "namespace-1")},
	))
}

// UpsertErrorAGEmptyObject Ошибка при добавлении networkBindings, ag пустой объект
func UpsertErrorAGEmptyObject() *foundation.TestCaseBodyNBUpsert {
	return upsertNBBody("Ошибка при добавлении networkBindings, ag пустой объект", nbEntry(
		&common.Metadata{Name: "err", Namespace: "namespace-1"},
		&sgroupsv1.NetworkBinding_Spec{AddressGroup: &common.ResourceIdentifier{}, Network: ri("nw-1", "namespace-1")},
	))
}

// UpsertErrorAGMissing Ошибка при добавлении networkBindings, ag отсутствует как параметр
func UpsertErrorAGMissing() *foundation.TestCaseBodyNBUpsert {
	return upsertNBBody("Ошибка при добавлении networkBindings, ag отсутствует как параметр", nbEntry(
		&common.Metadata{Name: "err", Namespace: "namespace-1"},
		&sgroupsv1.NetworkBinding_Spec{Network: ri("nw-3", "namespace-1")},
	))
}

// UpsertErrorAGNamespaceMissing Ошибка при добавлении networkBindings, в ag ns отсутствует как параметр
func UpsertErrorAGNamespaceMissing() *foundation.TestCaseBodyNBUpsert {
	return upsertNBBody("Ошибка при добавлении networkBindings, в ag ns отсутствует как параметр", nbEntry(
		&common.Metadata{Name: "err", Namespace: "namespace-1"},
		&sgroupsv1.NetworkBinding_Spec{AddressGroup: ri("ag-3", ""), Network: ri("nw-1", "namespace-1")},
	))
}

// UpsertErrorAGNameMissing Ошибка при добавлении networkBindings, в ag name отсутствует как параметр
func UpsertErrorAGNameMissing() *foundation.TestCaseBodyNBUpsert {
	return upsertNBBody("Ошибка при добавлении networkBindings, в ag name отсутствует как параметр", nbEntry(
		&common.Metadata{Name: "err", Namespace: "namespace-1"},
		&sgroupsv1.NetworkBinding_Spec{AddressGroup: ri("", "namespace-1"), Network: ri("nw-1", "namespace-1")},
	))
}

// UpsertErrorAGEmpty Ошибка при добавлении networkBindings, ag пуст (ns и name)
func UpsertErrorAGEmpty() *foundation.TestCaseBodyNBUpsert {
	return upsertNBBody("Ошибка при добавлении networkBindings, ag пуст (ns и name)", nbEntry(
		&common.Metadata{Name: "err", Namespace: "namespace-1"},
		&sgroupsv1.NetworkBinding_Spec{AddressGroup: ri("", ""), Network: ri("nw-1", "namespace-1")},
	))
}

// UpsertErrorAGNameEmpty Ошибка при добавлении networkBindings, ag name пуст
func UpsertErrorAGNameEmpty() *foundation.TestCaseBodyNBUpsert {
	return upsertNBBody("Ошибка при добавлении networkBindings, ag name пуст", nbEntry(
		&common.Metadata{Name: "err", Namespace: "namespace-1"},
		&sgroupsv1.NetworkBinding_Spec{AddressGroup: ri("", "namespace-1"), Network: ri("nw-1", "namespace-1")},
	))
}

// UpsertErrorAGNotInNamespaceButExists Ошибка при добавлении networkBindings, в указанном ag нет такого ag (сам ag существует)
func UpsertErrorAGNotInNamespaceButExists() *foundation.TestCaseBodyNBUpsert {
	return upsertNBBody("Ошибка при добавлении networkBindings, в указанном ag нет такого ag (сам ag существует)", nbEntry(
		&common.Metadata{Name: "err", Namespace: "namespace-1"},
		&sgroupsv1.NetworkBinding_Spec{AddressGroup: ri("ag-2", "namespace-1"), Network: ri("nw-1", "namespace-1")},
	))
}

// UpsertErrorAGNotExist Ошибка при добавлении networkBindings, ag не существует в указанном ns
func UpsertErrorAGNotExist() *foundation.TestCaseBodyNBUpsert {
	return upsertNBBody("Ошибка при добавлении networkBindings, ag не существует в указанном ns", nbEntry(
		&common.Metadata{Name: "err", Namespace: "namespace-1"},
		&sgroupsv1.NetworkBinding_Spec{AddressGroup: ri("ag-33", "namespace-1"), Network: ri("nw-7", "namespace-1")},
	))
}

// UpsertErrorSpecEmpty Ошибка при добавлении networkBindings, spec пустой объект
func UpsertErrorSpecEmpty() *foundation.TestCaseBodyNBUpsert {
	return upsertNBBody("Ошибка при добавлении networkBindings, spec пустой объект", nbEntry(
		&common.Metadata{Name: "err", Namespace: "namespace-1"},
		&sgroupsv1.NetworkBinding_Spec{},
	))
}
