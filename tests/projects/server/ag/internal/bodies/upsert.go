package bodies

import (
	"strings"

	sgroupsv1 "HBF-tests/pkg/api/sgroups/v1"
	"HBF-tests/tests/common/foundation"

	common "github.com/PRO-Robotech/sgroups-proto/pkg/api/common"
)

// tooLongAGName 64-символьная строка из строчных латинских букв для проверки ограничения длины в 63 символа
var tooLongAGName = strings.Repeat("a", 64)

func upsertAGBody(testName string, ags ...*sgroupsv1.AddressGroup) *foundation.TestCaseBodyAGUpsert {
	return &foundation.TestCaseBodyAGUpsert{
		TestName: testName,
		Req:      sgroupsv1.AddressGroupReq_Upsert{AddressGroups: ags},
	}
}

func agEntry(metadata *common.Metadata, spec *sgroupsv1.AddressGroup_Spec) *sgroupsv1.AddressGroup {
	return &sgroupsv1.AddressGroup{Metadata: metadata, Spec: spec}
}

// errSpec Тело spec, общее для большинства негативных кейсов добавления AG
func errSpec() *sgroupsv1.AddressGroup_Spec {
	return &sgroupsv1.AddressGroup_Spec{
		Comment:       "err",
		Description:   "err",
		DisplayName:   "new addressgroup-0",
		Logs:          true,
		Trace:         true,
		DefaultAction: common.Action_DENY,
	}
}

// upsertErrName Негативный кейс с невалидным именем AG (namespace-0, стандартный errSpec)
func upsertErrName(testName, name string) *foundation.TestCaseBodyAGUpsert {
	return upsertAGBody(testName, agEntry(&common.Metadata{Name: name, Namespace: "namespace-0"}, errSpec()))
}

// UpsertAdd Добавление новой AG
func UpsertAdd() *foundation.TestCaseBodyAGUpsert {
	return upsertAGBody("Добавление новой AG", agEntry(
		&common.Metadata{
			Name:        "add-ag-0",
			Namespace:   "namespace-0",
			Labels:      map[string]string{"add": "new"},
			Annotations: map[string]string{"new": "add"},
		},
		&sgroupsv1.AddressGroup_Spec{
			Comment: "added success", Description: "add success", DisplayName: "new addressgroup-0",
			Logs: true, Trace: true, DefaultAction: common.Action_DENY,
		},
	))
}

// UpsertAddSameNameDifferentNamespace Добавление AG с существующим именем, но к другому namespace
func UpsertAddSameNameDifferentNamespace() *foundation.TestCaseBodyAGUpsert {
	return upsertAGBody("Добавление AG с существующим именем, но к другому namespace", agEntry(
		&common.Metadata{
			Name:        "ag-0",
			Namespace:   "namespace-1",
			Labels:      map[string]string{"add": "existname"},
			Annotations: map[string]string{"new": "add"},
		},
		&sgroupsv1.AddressGroup_Spec{
			Comment: "added success name", Description: "add success", DisplayName: "new addressgroup-0",
			Logs: true, Trace: false, DefaultAction: common.Action_ALLOW,
		},
	))
}

// UpsertEdit Редактирование AG
func UpsertEdit() *foundation.TestCaseBodyAGUpsert {
	return upsertAGBody("Редактирование AG", agEntry(
		&common.Metadata{
			Name:        "ag-3",
			Namespace:   "namespace-1",
			Uid:         "9e9fb305-80bd-4748-ac42-fc208c398220",
			Labels:      map[string]string{"edit": "success"},
			Annotations: map[string]string{"success": "edit"},
		},
		&sgroupsv1.AddressGroup_Spec{
			Comment: "edit success", Description: "edit success", DisplayName: "edit addressgroup-0",
			Logs: true, Trace: true, DefaultAction: common.Action_DENY,
		},
	))
}

// UpsertErrorEditUIDMismatch Ошибка при редактировании AG, uid не соответствует указанной AG
func UpsertErrorEditUIDMismatch() *foundation.TestCaseBodyAGUpsert {
	return upsertAGBody("Ошибка при редактировании AG, uid не соответствует указанной AG", agEntry(
		&common.Metadata{
			Name:        "ag-3",
			Namespace:   "namespace-1",
			Uid:         "d459b92b-0881-4166-9035-6b994ebdf798",
			Labels:      map[string]string{"edit": "success"},
			Annotations: map[string]string{"success": "edit"},
		},
		&sgroupsv1.AddressGroup_Spec{
			Comment: "edit success", Description: "edit success", DisplayName: "edit addressgroup-0",
			Logs: true, Trace: true, DefaultAction: common.Action_DENY,
		},
	))
}

// UpsertErrorEditEmptyName Ошибка при редактировании AG, name пуст
func UpsertErrorEditEmptyName() *foundation.TestCaseBodyAGUpsert {
	return upsertAGBody("Ошибка при редактировании AG, name пуст", agEntry(
		&common.Metadata{
			Name:        "",
			Namespace:   "namespace-1",
			Uid:         "d459b92b-0881-4166-9035-6b994ebdf798",
			Labels:      map[string]string{"edit": "success"},
			Annotations: map[string]string{"success": "edit"},
		},
		&sgroupsv1.AddressGroup_Spec{
			Comment: "edit success", Description: "edit success", DisplayName: "edit addressgroup-0",
			Logs: true, Trace: true, DefaultAction: common.Action_DENY,
		},
	))
}

// UpsertErrorEditEmptyNamespace Ошибка при редактировании AG, namespace пуст
func UpsertErrorEditEmptyNamespace() *foundation.TestCaseBodyAGUpsert {
	return upsertAGBody("Ошибка при редактировании AG, namespace пуст", agEntry(
		&common.Metadata{
			Name:        "ag-3",
			Namespace:   "",
			Uid:         "917b9097-2b26-420f-8829-762edbee8d79",
			Labels:      map[string]string{"edit": "success"},
			Annotations: map[string]string{"success": "edit"},
		},
		&sgroupsv1.AddressGroup_Spec{
			Comment: "edit success", Description: "edit success", DisplayName: "edit addressgroup-0",
			Logs: true, Trace: true, DefaultAction: common.Action_DENY,
		},
	))
}

// UpsertErrorEditShortUID Ошибка при редактировании AG, uid короткий
func UpsertErrorEditShortUID() *foundation.TestCaseBodyAGUpsert {
	return upsertAGBody("Ошибка при редактировании AG, uid короткий", agEntry(
		&common.Metadata{
			Name:        "ag-3",
			Namespace:   "namespace-0",
			Uid:         "917b9097-2b26-420f-8829-762edbee8d7",
			Labels:      map[string]string{"edit": "success"},
			Annotations: map[string]string{"success": "edit"},
		},
		&sgroupsv1.AddressGroup_Spec{
			Comment: "edit success", Description: "edit success", DisplayName: "edit addressgroup-0",
			Logs: true, Trace: true, DefaultAction: common.Action_DENY,
		},
	))
}

// UpsertErrorEditLongUID Ошибка при редактировании AG, uid длинный
func UpsertErrorEditLongUID() *foundation.TestCaseBodyAGUpsert {
	return upsertAGBody("Ошибка при редактировании AG, uid длинный", agEntry(
		&common.Metadata{
			Name:        "ag-3",
			Namespace:   "namespace-0",
			Uid:         "917b9097-2b26-420f-8829-762edbee8d790",
			Labels:      map[string]string{"edit": "success"},
			Annotations: map[string]string{"success": "edit"},
		},
		&sgroupsv1.AddressGroup_Spec{
			Comment: "edit success", Description: "edit success", DisplayName: "edit addressgroup-0",
			Logs: true, Trace: true, DefaultAction: common.Action_DENY,
		},
	))
}

// UpsertErrorAddEmptyName Ошибка при добавлении AG, name пуст
func UpsertErrorAddEmptyName() *foundation.TestCaseBodyAGUpsert {
	return upsertAGBody("Ошибка при добавлении AG, name пуст",
		agEntry(&common.Metadata{Name: "", Namespace: "namespace-0"}, errSpec()))
}

// UpsertErrorAddDuplicateNameNamespace Ошибка при добавлении AG, AG c таким name+ns уже существует
func UpsertErrorAddDuplicateNameNamespace() *foundation.TestCaseBodyAGUpsert {
	return upsertAGBody("Ошибка при добавлении AG, AG c таким name+ns уже существует",
		agEntry(&common.Metadata{Name: "ag-0", Namespace: "namespace-0"}, errSpec()))
}

// UpsertErrorAddNameMissing Ошибка при добавлении AG, name отсутствует как параметр
func UpsertErrorAddNameMissing() *foundation.TestCaseBodyAGUpsert {
	return upsertAGBody("Ошибка при добавлении AG, name отсутствует как параметр",
		agEntry(&common.Metadata{Namespace: "namespace-0"}, errSpec()))
}

// UpsertErrorNameCyrillic Ошибка при добавлении AG, name на кириллице
func UpsertErrorNameCyrillic() *foundation.TestCaseBodyAGUpsert {
	return upsertErrName("Ошибка при добавлении AG, name на кириллице", "адрессная")
}

// UpsertErrorNameCyrillicDigits Ошибка при добавлении AG, name на кириллице+цифры
func UpsertErrorNameCyrillicDigits() *foundation.TestCaseBodyAGUpsert {
	return upsertErrName("Ошибка при добавлении AG, name на кириллице+цифры", "адрессная-1")
}

// UpsertErrorNameCyrillicSpecialChars Ошибка при добавлении AG, name на кириллице+спец.символы
func UpsertErrorNameCyrillicSpecialChars() *foundation.TestCaseBodyAGUpsert {
	return upsertErrName("Ошибка при добавлении AG, name на кириллице+спец.символы", "адрессная?")
}

// UpsertErrorNameLatinSpecialChars Ошибка при добавлении AG, name на латинице+спец.символы
func UpsertErrorNameLatinSpecialChars() *foundation.TestCaseBodyAGUpsert {
	return upsertErrName("Ошибка при добавлении AG, name на латинице+спец.символы", "address?")
}

// UpsertErrorNameUpperCase Ошибка при добавлении AG, name капсом
func UpsertErrorNameUpperCase() *foundation.TestCaseBodyAGUpsert {
	return upsertErrName("Ошибка при добавлении AG, name капсом", "ADDRESS")
}

// UpsertErrorNameStartsWithHyphen Ошибка при добавлении AG, name начинается с дефиса
func UpsertErrorNameStartsWithHyphen() *foundation.TestCaseBodyAGUpsert {
	return upsertErrName("Ошибка при добавлении AG, name начинается с дефиса", "-ag-0")
}

// UpsertErrorNameEndsWithHyphen Ошибка при добавлении AG, name заканчивается дефисом
func UpsertErrorNameEndsWithHyphen() *foundation.TestCaseBodyAGUpsert {
	return upsertErrName("Ошибка при добавлении AG, name заканчивается дефисом", "ag-0-")
}

// UpsertErrorNameSpecialCharsOnly Ошибка при добавлении AG, name из спец.символов
func UpsertErrorNameSpecialCharsOnly() *foundation.TestCaseBodyAGUpsert {
	return upsertErrName("Ошибка при добавлении AG, name из спец.символов", "??,.;")
}

// UpsertErrorNameLeadingSpace Ошибка при добавлении AG, пробел в начале name
func UpsertErrorNameLeadingSpace() *foundation.TestCaseBodyAGUpsert {
	return upsertErrName("Ошибка при добавлении AG, пробел в начале name", " err-name")
}

// UpsertErrorNameTrailingSpace Ошибка при добавлении AG, пробел в конце name
func UpsertErrorNameTrailingSpace() *foundation.TestCaseBodyAGUpsert {
	return upsertErrName("Ошибка при добавлении AG, пробел в конце name", "err-name ")
}

// UpsertErrorNameWithSpaces Ошибка при добавлении AG, name с пробелами
func UpsertErrorNameWithSpaces() *foundation.TestCaseBodyAGUpsert {
	return upsertErrName("Ошибка при добавлении AG, name с пробелами", "err - name")
}

// UpsertErrorNameTooLong Ошибка при добавлении AG, name > 63 символов
func UpsertErrorNameTooLong() *foundation.TestCaseBodyAGUpsert {
	return upsertErrName("Ошибка при добавлении AG, name > 63 символов", tooLongAGName)
}

// UpsertErrorNamespaceNotExist Ошибка при добавлении AG, namespace не существует
func UpsertErrorNamespaceNotExist() *foundation.TestCaseBodyAGUpsert {
	return upsertAGBody("Ошибка при добавлении AG, namespace не существует",
		agEntry(&common.Metadata{Name: "err-namespace", Namespace: "not-exist"}, errSpec()))
}

// UpsertErrorNamespaceLeadingSpace Ошибка при добавлении AG, пробел в начале имени namespace
func UpsertErrorNamespaceLeadingSpace() *foundation.TestCaseBodyAGUpsert {
	return upsertAGBody("Ошибка при добавлении AG, пробел в начале имени namespace",
		agEntry(&common.Metadata{Name: "err-namespace", Namespace: " namespace-0"}, errSpec()))
}

// UpsertErrorNamespaceTrailingSpace Ошибка при добавлении AG, пробел в конце имени namespace
func UpsertErrorNamespaceTrailingSpace() *foundation.TestCaseBodyAGUpsert {
	return upsertAGBody("Ошибка при добавлении AG, пробел в конце имени namespace",
		agEntry(&common.Metadata{Name: "err-namespace", Namespace: "namespace-0 "}, errSpec()))
}

// UpsertErrorNamespaceWithSpaces Ошибка при добавлении AG, имя namespace с пробелами
func UpsertErrorNamespaceWithSpaces() *foundation.TestCaseBodyAGUpsert {
	return upsertAGBody("Ошибка при добавлении AG, имя namespace с пробелами",
		agEntry(&common.Metadata{Name: "err-namespace", Namespace: "namespace - 0"}, errSpec()))
}

// UpsertErrorNamespaceEmpty Ошибка при добавлении AG, namespace пуст
func UpsertErrorNamespaceEmpty() *foundation.TestCaseBodyAGUpsert {
	return upsertAGBody("Ошибка при добавлении AG, namespace пуст",
		agEntry(&common.Metadata{Name: "err-namespace", Namespace: ""}, errSpec()))
}

// UpsertErrorNamespaceMissing Ошибка при добавлении AG, namespace отсутствует как параметр
func UpsertErrorNamespaceMissing() *foundation.TestCaseBodyAGUpsert {
	return upsertAGBody("Ошибка при добавлении AG, namespace отсутствует как параметр",
		agEntry(&common.Metadata{Name: "err-namespace"}, errSpec()))
}

// UpsertErrorAddExistingName Ошибка при добавлении AG, AG с таким именем существует
func UpsertErrorAddExistingName() *foundation.TestCaseBodyAGUpsert {
	return upsertAGBody("Ошибка при добавлении AG, AG с таким именем существует",
		agEntry(&common.Metadata{Name: "ag-0", Namespace: "namespace-0"}, errSpec()))
}

// UpsertErrorDisplayNameTooLong Ошибка при добавлении AG, displayName > 63 символов
func UpsertErrorDisplayNameTooLong() *foundation.TestCaseBodyAGUpsert {
	return upsertAGBody("Ошибка при добавлении AG, displayName > 63 символов", agEntry(
		&common.Metadata{Name: "err-display", Namespace: "namespace-0"},
		&sgroupsv1.AddressGroup_Spec{
			Comment: "err", Description: "err", DisplayName: tooLongAGName,
			Logs: true, Trace: true, DefaultAction: common.Action_DENY,
		},
	))
}

// UpsertErrorDefaultActionMissing Ошибка при добавлении AG, defaultAction отсутствует как параметр
func UpsertErrorDefaultActionMissing() *foundation.TestCaseBodyAGUpsert {
	return upsertAGBody("Ошибка при добавлении AG, defaultAction отсутствует как параметр", agEntry(
		&common.Metadata{Name: "err-display", Namespace: "namespace-0"},
		&sgroupsv1.AddressGroup_Spec{Comment: "err", Description: "err", Logs: true, Trace: true},
	))
}
