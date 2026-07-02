package bodies

import (
	"strings"

	common "github.com/PRO-Robotech/sgroups-proto/pkg/api/common"

	sgroupsv1 "HBF-tests/pkg/api/sgroups/v1"
	"HBF-tests/tests/common/foundation"
)

// tooLongHBName 64-символьная строка из строчных латинских букв для проверки ограничения длины в 63 символа
var tooLongHBName = strings.Repeat("a", 64)

func upsertHBBody(testName string, hbs ...*sgroupsv1.HostBinding) *foundation.TestCaseBodyHBUpsert {
	return &foundation.TestCaseBodyHBUpsert{
		TestName: testName,
		Req:      sgroupsv1.HostBindingReq_Upsert{HostBindings: hbs},
	}
}

func hbEntry(metadata *common.Metadata, spec *sgroupsv1.HostBinding_Spec) *sgroupsv1.HostBinding {
	return &sgroupsv1.HostBinding{Metadata: metadata, Spec: spec}
}

// errNameSpec Spec, общий для большинства негативных кейсов добавления host-binding по имени
func errNameSpec() *sgroupsv1.HostBinding_Spec {
	return &sgroupsv1.HostBinding_Spec{
		AddressGroup: ri("ag-3", "namespace-1"),
		Host:         ri("host-7", "namespace-1"),
	}
}

// upsertErrName Негативный кейс с невалидным именем host-binding (namespace-1, стандартный errNameSpec)
func upsertErrName(testName, name string) *foundation.TestCaseBodyHBUpsert {
	return upsertHBBody(testName, hbEntry(&common.Metadata{Name: name, Namespace: "namespace-1"}, errNameSpec()))
}

// UpsertAdd Добавление нового host-binding
func UpsertAdd() *foundation.TestCaseBodyHBUpsert {
	return upsertHBBody("Добавление нового host-binding", hbEntry(
		&common.Metadata{
			Name:        "add-hb-1",
			Namespace:   "namespace-1",
			Labels:      map[string]string{"add": "success"},
			Annotations: map[string]string{"add": "success"},
		},
		&sgroupsv1.HostBinding_Spec{
			Comment: "new host binding", Description: "add success", DisplayName: "new host binding 1",
			AddressGroup: ri("ag-6", "namespace-1"), Host: ri("host-6", "namespace-1"),
		},
	))
}

// UpsertEdit Редактирование host-binging
func UpsertEdit() *foundation.TestCaseBodyHBUpsert {
	return upsertHBBody("Редактирование host-binging", hbEntry(
		&common.Metadata{
			Name:        "hb-3",
			Namespace:   "namespace-1",
			Uid:         "efb85252-c9b7-4dd5-b9be-6ce63276795c",
			Labels:      map[string]string{"edit": "success"},
			Annotations: map[string]string{"edit": "success"},
		},
		&sgroupsv1.HostBinding_Spec{
			Comment: "edited host binding", Description: "edit success", DisplayName: "edit host binding 1",
			AddressGroup: ri("ag-3", "namespace-1"), Host: ri("host-3", "namespace-1"),
		},
	))
}

// UpsertAddSameNameDifferentNamespace Добавление host-binding с существующим именем, но к другому namespace
func UpsertAddSameNameDifferentNamespace() *foundation.TestCaseBodyHBUpsert {
	return upsertHBBody("Добавление host-binding с существующим именем, но к другому namespace", hbEntry(
		&common.Metadata{
			Name:        "hb-0",
			Namespace:   "namespace-1",
			Labels:      map[string]string{"add": "success"},
			Annotations: map[string]string{"add": "success"},
		},
		&sgroupsv1.HostBinding_Spec{
			Comment: "new host binding", Description: "add success", DisplayName: "new host binding 1",
			AddressGroup: ri("ag-6", "namespace-1"), Host: ri("host-7", "namespace-1"),
		},
	))
}

// UpsertErrorDisplayNameTooLong Ошибка при добавлении host-binding, displayName > 63 символов
func UpsertErrorDisplayNameTooLong() *foundation.TestCaseBodyHBUpsert {
	return upsertHBBody("Ошибка при добавлении host-binding, displayName > 63 символов", hbEntry(
		&common.Metadata{Name: "err-display", Namespace: "namespace-1"},
		&sgroupsv1.HostBinding_Spec{
			DisplayName: tooLongHBName, AddressGroup: ri("ag-1", "namespace-1"), Host: ri("host-7", "namespace-1"),
		},
	))
}

// UpsertErrorEditAGAndHostMissing Ошибка при редактировании host-binding, ag и host отсутствуют
func UpsertErrorEditAGAndHostMissing() *foundation.TestCaseBodyHBUpsert {
	return upsertHBBody("Ошибка при редактировании host-binding,  ag и host отсутствуют", hbEntry(
		&common.Metadata{
			Name: "hb-3", Namespace: "namespace-1", Uid: "efb85252-c9b7-4dd5-b9be-6ce63276795c",
			Labels: map[string]string{"edit": "success"}, Annotations: map[string]string{"edit": "success"},
		},
		&sgroupsv1.HostBinding_Spec{Comment: "edited host binding", Description: "edit success", DisplayName: "edit host binding 1"},
	))
}

// editSpecBase Spec, общий для большинства негативных кейсов редактирования hb-3
func editSpecBase(ag, host *common.ResourceIdentifier) *sgroupsv1.HostBinding_Spec {
	return &sgroupsv1.HostBinding_Spec{
		Comment: "edited host binding", Description: "edit success", DisplayName: "edit host binding 1",
		AddressGroup: ag, Host: host,
	}
}

func editHB3(testName string, uid string, ag, host *common.ResourceIdentifier) *foundation.TestCaseBodyHBUpsert {
	return upsertHBBody(testName, hbEntry(
		&common.Metadata{
			Name: "hb-3", Namespace: "namespace-1", Uid: uid,
			Labels: map[string]string{"edit": "success"}, Annotations: map[string]string{"edit": "success"},
		},
		editSpecBase(ag, host),
	))
}

// UpsertErrorEditHostImmutable Ошибка при редактировании host-binding, нельзя редактировать host
func UpsertErrorEditHostImmutable() *foundation.TestCaseBodyHBUpsert {
	return editHB3("Ошибка при редактировании host-binding, нельзя редактировать host", "efb85252-c9b7-4dd5-b9be-6ce63276795c", ri("ag-3", "namespace-1"), ri("host-1", "namespace-1"))
}

// UpsertErrorEditHostImmutableDifferentNamespace Ошибка при редактировании host-binding, нельзя редактировать host из другого ns
func UpsertErrorEditHostImmutableDifferentNamespace() *foundation.TestCaseBodyHBUpsert {
	return editHB3("Ошибка при редактировании host-binding, нельзя редактировать host из другого ns", "efb85252-c9b7-4dd5-b9be-6ce63276795c", ri("ag-3", "namespace-1"), ri("host-2", "namespace-2"))
}

// UpsertErrorEditHostMissing Ошибка при редактировании host-binding, host отсутствует
func UpsertErrorEditHostMissing() *foundation.TestCaseBodyHBUpsert {
	return upsertHBBody("Ошибка при редактировании host-binding, host отсутствует", hbEntry(
		&common.Metadata{
			Name: "hb-3", Namespace: "namespace-1", Uid: "efb85252-c9b7-4dd5-b9be-6ce63276795c",
			Labels: map[string]string{"edit": "success"}, Annotations: map[string]string{"edit": "success"},
		},
		&sgroupsv1.HostBinding_Spec{
			Comment: "edited host binding", Description: "edit success", DisplayName: "edit host binding 1",
			AddressGroup: ri("ag-3", "namespace-1"),
		},
	))
}

// UpsertErrorEditAGImmutable Ошибка при редактировании host-binding, нельзя редактировать ag
func UpsertErrorEditAGImmutable() *foundation.TestCaseBodyHBUpsert {
	return editHB3("Ошибка при редактировании host-binding, нельзя редактировать ag", "efb85252-c9b7-4dd5-b9be-6ce63276795c", ri("ag-6", "namespace-1"), ri("host-3", "namespace-1"))
}

// UpsertErrorEditAGImmutableDifferentNamespace Ошибка при редактировании host-binding, нельзя редактировать ag из другого ns
func UpsertErrorEditAGImmutableDifferentNamespace() *foundation.TestCaseBodyHBUpsert {
	return editHB3("Ошибка при редактировании host-binding, нельзя редактировать ag из другого ns", "efb85252-c9b7-4dd5-b9be-6ce63276795c", ri("ag-2", "namespace-2"), ri("host-3", "namespace-1"))
}

// UpsertErrorEditAGMissing Ошибка при редактировании host-binding, ag отсутствует
func UpsertErrorEditAGMissing() *foundation.TestCaseBodyHBUpsert {
	return upsertHBBody("Ошибка при редактировании host-binding, ag отсутствует", hbEntry(
		&common.Metadata{
			Name: "hb-3", Namespace: "namespace-1", Uid: "efb85252-c9b7-4dd5-b9be-6ce63276795c",
			Labels: map[string]string{"edit": "success"}, Annotations: map[string]string{"edit": "success"},
		},
		&sgroupsv1.HostBinding_Spec{
			Comment: "edited host binding", Description: "edit success", DisplayName: "edit host binding 1",
			Host: ri("host-3", "namespace-1"),
		},
	))
}

// UpsertErrorBindingAlreadyExists Ошибка при добавлении host-binding, такая привязка уже существует (ag-host)
func UpsertErrorBindingAlreadyExists() *foundation.TestCaseBodyHBUpsert {
	return upsertHBBody("Ошибка при добавлении host-binding, такая привязка уже существует (ag-host)", hbEntry(
		&common.Metadata{Name: "hb-1000", Namespace: "namespace-0"},
		&sgroupsv1.HostBinding_Spec{AddressGroup: ri("ag-0", "namespace-0"), Host: ri("host-0", "namespace-0")},
	))
}

// UpsertErrorAddDuplicateNameNamespace Ошибка при добавлении host-binding, hb  с таким name+namespace уже существует
func UpsertErrorAddDuplicateNameNamespace() *foundation.TestCaseBodyHBUpsert {
	return upsertHBBody("Ошибка при добавлении host-binding, hb  с таким name+namespace уже существует", hbEntry(
		&common.Metadata{Name: "hb-0", Namespace: "namespace-0"},
		&sgroupsv1.HostBinding_Spec{AddressGroup: ri("ag-0", "namespace-0"), Host: ri("host-0", "namespace-0")},
	))
}

// UpsertErrorHostNamespaceMismatch Ошибка при добавлении host-binding, host ns с не совпадает с hb ns
func UpsertErrorHostNamespaceMismatch() *foundation.TestCaseBodyHBUpsert {
	return upsertHBBody("Ошибка при добавлении host-binding, host ns с не совпадает с hb ns", hbEntry(
		&common.Metadata{Name: "err-ns", Namespace: "namespace-1"},
		&sgroupsv1.HostBinding_Spec{AddressGroup: ri("ag-1", "namespace-1"), Host: ri("host-2", "namespace-2")},
	))
}

// UpsertErrorAGNamespaceMismatch Ошибка при добавлении host-binding, ag ns с не совпадает с hb ns
func UpsertErrorAGNamespaceMismatch() *foundation.TestCaseBodyHBUpsert {
	return upsertHBBody("Ошибка при добавлении host-binding, ag ns с не совпадает с hb ns", hbEntry(
		&common.Metadata{Name: "err-ns", Namespace: "namespace-1"},
		&sgroupsv1.HostBinding_Spec{AddressGroup: ri("ag-2", "namespace-2"), Host: ri("host-1", "namespace-1")},
	))
}

// UpsertErrorHBNamespaceMismatchBoth Ошибка при добавлении host-binding, hb ns не совпадает с ag и host ns
func UpsertErrorHBNamespaceMismatchBoth() *foundation.TestCaseBodyHBUpsert {
	return upsertHBBody("Ошибка при добавлении host-binding, hb ns не совпадает с ag и host ns", hbEntry(
		&common.Metadata{Name: "err-ns", Namespace: "namespace-1"},
		&sgroupsv1.HostBinding_Spec{AddressGroup: ri("ag-2", "namespace-2"), Host: ri("host-2", "namespace-2")},
	))
}

// UpsertErrorEditLongUID Ошибка при редактировании host-binding, uid длинный
func UpsertErrorEditLongUID() *foundation.TestCaseBodyHBUpsert {
	return editHB3("Ошибка при редактировании host-binding, uid длинный", "efb85252-c9b7-4dd5-b9be-6ce63276795vc", ri("ag-3", "namespace-1"), ri("host-3", "namespace-1"))
}

// UpsertErrorEditShortUID Ошибка при редактировании host-binding, uid короткий
func UpsertErrorEditShortUID() *foundation.TestCaseBodyHBUpsert {
	return editHB3("Ошибка при редактировании host-binding, uid короткий", "efb85252-c9b7-4dd5-b9be-6ce63276795", ri("ag-3", "namespace-1"), ri("host-3", "namespace-1"))
}

// UpsertErrorEditUIDMismatch Ошибка при редактировании host-binding, uid не соответствует указанному hb
func UpsertErrorEditUIDMismatch() *foundation.TestCaseBodyHBUpsert {
	return editHB3("Ошибка при редактировании host-binding, uid не соответствует указанному hb", "1085d231-9a0f-4697-b864-c0a522181911", ri("ag-3", "namespace-1"), ri("host-3", "namespace-1"))
}

// UpsertErrorEditNameUIDNotExistInNamespace Ошибка при редактировании host-binding, name+uid не существует в указанном namespace
func UpsertErrorEditNameUIDNotExistInNamespace() *foundation.TestCaseBodyHBUpsert {
	return upsertHBBody("Ошибка при редактировании host-binding, name+uid не существует в указанном namespace", hbEntry(
		&common.Metadata{Name: "hb-0", Namespace: "namespace-1", Uid: "1085d231-9a0f-4697-b864-c0a522181911"},
		editSpecBase(ri("ag-3", "namespace-1"), ri("host-3", "namespace-1")),
	))
}

// UpsertErrorEditEmptyName Ошибка при редактировании host-binding, name пуст
func UpsertErrorEditEmptyName() *foundation.TestCaseBodyHBUpsert {
	return upsertHBBody("Ошибка при редактировании host-binding, name пуст", hbEntry(
		&common.Metadata{Name: "", Namespace: "namespace-0", Uid: "1085d231-9a0f-4697-b864-c0a522181911"},
		editSpecBase(ri("ag-3", "namespace-1"), ri("host-3", "namespace-1")),
	))
}

// UpsertErrorEditNameMissing Ошибка при редактировании host-binding, name отсутствует как параметр
func UpsertErrorEditNameMissing() *foundation.TestCaseBodyHBUpsert {
	return upsertHBBody("Ошибка при редактировании host-binding, name отсутствует как параметр", hbEntry(
		&common.Metadata{Namespace: "namespace-0", Uid: "1085d231-9a0f-4697-b864-c0a522181911"},
		editSpecBase(ri("ag-3", "namespace-1"), ri("host-3", "namespace-1")),
	))
}

// UpsertErrorEditEmptyNamespace Ошибка при редактировании host-binding, namespace пуст
func UpsertErrorEditEmptyNamespace() *foundation.TestCaseBodyHBUpsert {
	return upsertHBBody("Ошибка при редактировании host-binding, namespace пуст", hbEntry(
		&common.Metadata{Name: "hb-0", Namespace: "", Uid: "1085d231-9a0f-4697-b864-c0a522181911"},
		editSpecBase(ri("ag-3", "namespace-1"), ri("host-3", "namespace-1")),
	))
}

// UpsertErrorEditNamespaceMissing Ошибка при редактировании host-binding, namespace отсутствует как параметр
func UpsertErrorEditNamespaceMissing() *foundation.TestCaseBodyHBUpsert {
	return upsertHBBody("Ошибка при редактировании host-binding, namespace отсутствует как параметр", hbEntry(
		&common.Metadata{Name: "hb-0", Uid: "1085d231-9a0f-4697-b864-c0a522181911"},
		&sgroupsv1.HostBinding_Spec{AddressGroup: ri("ag-3", "namespace-1"), Host: ri("host-3", "namespace-1")},
	))
}

// UpsertErrorAddEmptyName Ошибка при добавлении host-binding, name пуст
func UpsertErrorAddEmptyName() *foundation.TestCaseBodyHBUpsert {
	return upsertHBBody("Ошибка при добавлении host-binding, name пуст", hbEntry(
		&common.Metadata{Name: "", Namespace: "namespace-1"},
		&sgroupsv1.HostBinding_Spec{AddressGroup: ri("ag-3", "namespace-1"), Host: ri("host-7", "namespace-1")},
	))
}

// UpsertErrorNameCyrillic Ошибка при добавлении host-binding, name на кириллице
func UpsertErrorNameCyrillic() *foundation.TestCaseBodyHBUpsert {
	return upsertErrName("Ошибка при добавлении host-binding, name на кириллице", "биндинг")
}

// UpsertErrorNameUpperCase Ошибка при добавлении host-binding, name капсом
func UpsertErrorNameUpperCase() *foundation.TestCaseBodyHBUpsert {
	return upsertErrName("Ошибка при добавлении host-binding, name капсом", "BINDING")
}

// UpsertErrorNameSpecialCharsOnly Ошибка при добавлении host-binding, name из спец символов
func UpsertErrorNameSpecialCharsOnly() *foundation.TestCaseBodyHBUpsert {
	return upsertErrName("Ошибка при добавлении host-binding, name из спец символов", "???!!!")
}

// UpsertErrorNameMixedCase Ошибка при добавлении host-binding, name капс+строчные
func UpsertErrorNameMixedCase() *foundation.TestCaseBodyHBUpsert {
	return upsertErrName("Ошибка при добавлении host-binding, name капс+строчные", "Binding")
}

// UpsertErrorNameCyrillicLatin Ошибка при добавлении host-binding, name на кириллице+латинице
func UpsertErrorNameCyrillicLatin() *foundation.TestCaseBodyHBUpsert {
	return upsertErrName("Ошибка при добавлении host-binding, name на кириллице+латинице", "binдинг")
}

// UpsertErrorNameCyrillicDigits Ошибка при добавлении host-binding, name на кириллице+числа
func UpsertErrorNameCyrillicDigits() *foundation.TestCaseBodyHBUpsert {
	return upsertErrName("Ошибка при добавлении host-binding, name на кириллице+числа", "биндинг-1")
}

// UpsertErrorNameStartsWithHyphen Ошибка при добавлении host-binding, name начинается с тире
func UpsertErrorNameStartsWithHyphen() *foundation.TestCaseBodyHBUpsert {
	return upsertErrName("Ошибка при добавлении host-binding, name начинается с тире", "-bind")
}

// UpsertErrorNameOnlyHyphen Ошибка при добавлении host-binding, name состоит только из тире
func UpsertErrorNameOnlyHyphen() *foundation.TestCaseBodyHBUpsert {
	return upsertErrName("Ошибка при добавлении host-binding, name состоит только из тире", "-")
}

// UpsertErrorNameEndsWithHyphen Ошибка при добавлении host-binding, name заканчивается тире
func UpsertErrorNameEndsWithHyphen() *foundation.TestCaseBodyHBUpsert {
	return upsertErrName("Ошибка при добавлении host-binding, name заканчивается тире", "bind-")
}

// UpsertErrorNameCyrillicSpecialChars Ошибка при добавлении host-binding, name на кириллице+спец.символы
func UpsertErrorNameCyrillicSpecialChars() *foundation.TestCaseBodyHBUpsert {
	return upsertErrName("Ошибка при добавлении host-binding, name на кириллице+спец.символы", "биндинг?")
}

// UpsertErrorNameTooLong Ошибка при добавлении host-binding, name > 63 символов
func UpsertErrorNameTooLong() *foundation.TestCaseBodyHBUpsert {
	return upsertErrName("Ошибка при добавлении host-binding, name > 63 символов", tooLongHBName)
}

// UpsertErrorNameLeadingSpace Ошибка при добавлении host-binding, пробел в начале name
func UpsertErrorNameLeadingSpace() *foundation.TestCaseBodyHBUpsert {
	return upsertErrName("Ошибка при добавлении host-binding, пробел в начале name", " bind-1")
}

// UpsertErrorNameTrailingSpace Ошибка при добавлении host-binding, пробел в конце name
func UpsertErrorNameTrailingSpace() *foundation.TestCaseBodyHBUpsert {
	return upsertErrName("Ошибка при добавлении host-binding, пробел в конце name", "bind-1 ")
}

// UpsertErrorNameWithSpaces Ошибка при добавлении host-binding, name с пробелами
func UpsertErrorNameWithSpaces() *foundation.TestCaseBodyHBUpsert {
	return upsertErrName("Ошибка при добавлении host-binding, name с пробелами", "bind - 1")
}

// UpsertErrorNamespaceNotExist Ошибка при добавлении host-binding, namespace не существует
func UpsertErrorNamespaceNotExist() *foundation.TestCaseBodyHBUpsert {
	return upsertHBBody("Ошибка при добавлении host-binding, namespace не существует", hbEntry(
		&common.Metadata{Name: "bind", Namespace: "namespace-11"},
		&sgroupsv1.HostBinding_Spec{AddressGroup: ri("ag-3", "namespace-1"), Host: ri("host-7", "namespace-1")},
	))
}

// UpsertErrorNamespaceEmpty Ошибка при добавлении host-binding, namespace пуст
func UpsertErrorNamespaceEmpty() *foundation.TestCaseBodyHBUpsert {
	return upsertHBBody("Ошибка при добавлении host-binding, namespace пуст", hbEntry(
		&common.Metadata{Name: "bind", Namespace: ""},
		&sgroupsv1.HostBinding_Spec{AddressGroup: ri("ag-3", "namespace-1"), Host: ri("host-7", "namespace-1")},
	))
}

// UpsertErrorNamespaceLeadingSpace Ошибка при добавлении host-binding, пробел в начале имени namespace
func UpsertErrorNamespaceLeadingSpace() *foundation.TestCaseBodyHBUpsert {
	return upsertHBBody("Ошибка при добавлении host-binding, пробел в начале имени namespace", hbEntry(
		&common.Metadata{Name: "bind", Namespace: " namespace-1"},
		&sgroupsv1.HostBinding_Spec{AddressGroup: ri("ag-3", "namespace-1"), Host: ri("host-7", "namespace-1")},
	))
}

// UpsertErrorNamespaceTrailingSpace Ошибка при добавлении host-binding, пробел в конце имени namespace
func UpsertErrorNamespaceTrailingSpace() *foundation.TestCaseBodyHBUpsert {
	return upsertHBBody("Ошибка при добавлении host-binding, пробел в конце имени namespace", hbEntry(
		&common.Metadata{Name: "bind", Namespace: "namespace-1 "},
		&sgroupsv1.HostBinding_Spec{AddressGroup: ri("ag-3", "namespace-1"), Host: ri("host-7", "namespace-1")},
	))
}

// UpsertErrorNamespaceWithSpaces Ошибка при добавлении host-binding, namespace с пробелами
func UpsertErrorNamespaceWithSpaces() *foundation.TestCaseBodyHBUpsert {
	return upsertHBBody("Ошибка при добавлении host-binding, namespace с пробелами", hbEntry(
		&common.Metadata{Name: "bind", Namespace: "namespace - 1"},
		&sgroupsv1.HostBinding_Spec{AddressGroup: ri("ag-3", "namespace-1"), Host: ri("host-7", "namespace-1")},
	))
}

// UpsertErrorNamespaceMissing Ошибка при добавлении host-binding, namespace отсутствует как параметр
func UpsertErrorNamespaceMissing() *foundation.TestCaseBodyHBUpsert {
	return upsertHBBody("Ошибка при добавлении host-binding, namespace отсутствует как параметр", hbEntry(
		&common.Metadata{Name: "bind"},
		&sgroupsv1.HostBinding_Spec{AddressGroup: ri("ag-3", "namespace-1"), Host: ri("host-7", "namespace-1")},
	))
}

// UpsertErrorHostNotExistInNamespace Ошибка при добавлении host-binding, host не существует в указанном ns
func UpsertErrorHostNotExistInNamespace() *foundation.TestCaseBodyHBUpsert {
	return upsertHBBody("Ошибка при добавлении host-binding, host не существует в указанном ns", hbEntry(
		&common.Metadata{Name: "err", Namespace: "namespace-1"},
		&sgroupsv1.HostBinding_Spec{AddressGroup: ri("ag-3", "namespace-1"), Host: ri("host-77", "namespace-1")},
	))
}

// UpsertErrorHostNotInNamespaceButExists Ошибка при добавлении host-binding, в указанном ns нет такого хоста (сам хост существует)
func UpsertErrorHostNotInNamespaceButExists() *foundation.TestCaseBodyHBUpsert {
	return upsertHBBody("Ошибка при добавлении host-binding, в указанном ns нет такого хоста (сам хост существует)", hbEntry(
		&common.Metadata{Name: "err", Namespace: "namespace-1"},
		&sgroupsv1.HostBinding_Spec{AddressGroup: ri("ag-3", "namespace-1"), Host: ri("host-2", "namespace-1")},
	))
}

// UpsertErrorHostNameEmpty Ошибка при добавлении host-binding, host name пуст
func UpsertErrorHostNameEmpty() *foundation.TestCaseBodyHBUpsert {
	return upsertHBBody("Ошибка при добавлении host-binding, host name пуст", hbEntry(
		&common.Metadata{Name: "err", Namespace: "namespace-1"},
		&sgroupsv1.HostBinding_Spec{AddressGroup: ri("ag-3", "namespace-1"), Host: ri("", "namespace-1")},
	))
}

// UpsertErrorHostEmpty Ошибка при добавлении host-binding, host пуст (ns и name)
func UpsertErrorHostEmpty() *foundation.TestCaseBodyHBUpsert {
	return upsertHBBody("Ошибка при добавлении host-binding, host пуст (ns и name)", hbEntry(
		&common.Metadata{Name: "err", Namespace: "namespace-1"},
		&sgroupsv1.HostBinding_Spec{AddressGroup: ri("ag-3", "namespace-1"), Host: ri("", "")},
	))
}

// UpsertErrorHostNameMissing Ошибка при добавлении host-binding, в host name отсутствует как параметр
func UpsertErrorHostNameMissing() *foundation.TestCaseBodyHBUpsert {
	return upsertHBBody("Ошибка при добавлении host-binding, в host name отсутствует как параметр", hbEntry(
		&common.Metadata{Name: "err", Namespace: "namespace-1"},
		&sgroupsv1.HostBinding_Spec{AddressGroup: ri("ag-3", "namespace-1"), Host: ri("", "namespace-1")},
	))
}

// UpsertErrorHostNamespaceMissing Ошибка при добавлении host-binding, в host ns отсутствует как параметр
func UpsertErrorHostNamespaceMissing() *foundation.TestCaseBodyHBUpsert {
	return upsertHBBody("Ошибка при добавлении host-binding, в host ns отсутствует как параметр", hbEntry(
		&common.Metadata{Name: "err", Namespace: "namespace-1"},
		&sgroupsv1.HostBinding_Spec{AddressGroup: ri("ag-3", "namespace-1"), Host: ri("host-1", "")},
	))
}

// UpsertErrorHostEmptyObject Ошибка при добавлении host-binding, host пустой объект
func UpsertErrorHostEmptyObject() *foundation.TestCaseBodyHBUpsert {
	return upsertHBBody("Ошибка при добавлении host-binding, host пустой объект", hbEntry(
		&common.Metadata{Name: "err", Namespace: "namespace-1"},
		&sgroupsv1.HostBinding_Spec{AddressGroup: ri("ag-3", "namespace-1"), Host: ri("host-1", "")},
	))
}

// UpsertErrorHostMissing Ошибка при добавлении host-binding, host отсутствует как параметр
func UpsertErrorHostMissing() *foundation.TestCaseBodyHBUpsert {
	return upsertHBBody("Ошибка при добавлении host-binding, host отсутствует как параметр", hbEntry(
		&common.Metadata{Name: "err", Namespace: "namespace-1"},
		&sgroupsv1.HostBinding_Spec{AddressGroup: ri("ag-3", "namespace-1")},
	))
}

// UpsertErrorAGEmptyObject Ошибка при добавлении host-binding, ag пустой объект
func UpsertErrorAGEmptyObject() *foundation.TestCaseBodyHBUpsert {
	return upsertHBBody("Ошибка при добавлении host-binding, ag пустой объект", hbEntry(
		&common.Metadata{Name: "err", Namespace: "namespace-1"},
		&sgroupsv1.HostBinding_Spec{AddressGroup: &common.ResourceIdentifier{}, Host: ri("host-1", "namespace-1")},
	))
}

// UpsertErrorAGMissing Ошибка при добавлении host-binding, ag отсутствует как параметр
func UpsertErrorAGMissing() *foundation.TestCaseBodyHBUpsert {
	return upsertHBBody("Ошибка при добавлении host-binding, ag отсутствует как параметр", hbEntry(
		&common.Metadata{Name: "err", Namespace: "namespace-1"},
		&sgroupsv1.HostBinding_Spec{Host: ri("host-3", "namespace-1")},
	))
}

// UpsertErrorAGNamespaceMissing Ошибка при добавлении host-binding, в ag ns отсутствует как параметр
func UpsertErrorAGNamespaceMissing() *foundation.TestCaseBodyHBUpsert {
	return upsertHBBody("Ошибка при добавлении host-binding, в ag ns отсутствует как параметр", hbEntry(
		&common.Metadata{Name: "err", Namespace: "namespace-1"},
		&sgroupsv1.HostBinding_Spec{AddressGroup: ri("ag-3", ""), Host: ri("host-1", "namespace-1")},
	))
}

// UpsertErrorAGNameMissing Ошибка при добавлении host-binding, в ag name отсутствует как параметр
func UpsertErrorAGNameMissing() *foundation.TestCaseBodyHBUpsert {
	return upsertHBBody("Ошибка при добавлении host-binding, в ag name отсутствует как параметр", hbEntry(
		&common.Metadata{Name: "err", Namespace: "namespace-1"},
		&sgroupsv1.HostBinding_Spec{AddressGroup: ri("", "namespace-1"), Host: ri("host-1", "namespace-1")},
	))
}

// UpsertErrorAGEmpty Ошибка при добавлении host-binding, ag пуст (ns и name)
func UpsertErrorAGEmpty() *foundation.TestCaseBodyHBUpsert {
	return upsertHBBody("Ошибка при добавлении host-binding, ag пуст (ns и name)", hbEntry(
		&common.Metadata{Name: "err", Namespace: "namespace-1"},
		&sgroupsv1.HostBinding_Spec{AddressGroup: ri("", ""), Host: ri("host-1", "namespace-1")},
	))
}

// UpsertErrorAGNameEmpty Ошибка при добавлении host-binding, ag name пуст
func UpsertErrorAGNameEmpty() *foundation.TestCaseBodyHBUpsert {
	return upsertHBBody("Ошибка при добавлении host-binding, ag name пуст", hbEntry(
		&common.Metadata{Name: "err", Namespace: "namespace-1"},
		&sgroupsv1.HostBinding_Spec{AddressGroup: ri("", "namespace-1"), Host: ri("host-1", "namespace-1")},
	))
}

// UpsertErrorAGNotInNamespaceButExists Ошибка при добавлении host-binding, в указанном ag нет такого ag (сам ag существует)
func UpsertErrorAGNotInNamespaceButExists() *foundation.TestCaseBodyHBUpsert {
	return upsertHBBody("Ошибка при добавлении host-binding, в указанном ag нет такого ag (сам ag существует)", hbEntry(
		&common.Metadata{Name: "err", Namespace: "namespace-1"},
		&sgroupsv1.HostBinding_Spec{AddressGroup: ri("ag-2", "namespace-1"), Host: ri("host-1", "namespace-1")},
	))
}

// UpsertErrorAGNotExist Ошибка при добавлении host-binding, ag не существует в указанном ag
func UpsertErrorAGNotExist() *foundation.TestCaseBodyHBUpsert {
	return upsertHBBody("Ошибка при добавлении host-binding, ag не существует в указанном ag", hbEntry(
		&common.Metadata{Name: "err", Namespace: "namespace-1"},
		&sgroupsv1.HostBinding_Spec{AddressGroup: ri("ag-33", "namespace-1"), Host: ri("host-7", "namespace-1")},
	))
}

// UpsertErrorSpecEmpty Ошибка при добавлении host-binding, spec пустой объект
func UpsertErrorSpecEmpty() *foundation.TestCaseBodyHBUpsert {
	return upsertHBBody("Ошибка при добавлении host-binding, spec пустой объект", hbEntry(
		&common.Metadata{Name: "err", Namespace: "namespace-1"},
		&sgroupsv1.HostBinding_Spec{},
	))
}
