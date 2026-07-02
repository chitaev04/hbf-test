package bodies

import (
	"strings"

	sgroupsv1 "HBF-tests/pkg/api/sgroups/v1"
	"HBF-tests/tests/common/foundation"

	common "github.com/PRO-Robotech/sgroups-proto/pkg/api/common"
)

// tooLongSBName 64-символьная строка из строчных латинских букв для проверки ограничения длины в 63 символа
var tooLongSBName = strings.Repeat("a", 64)

func upsertSBBody(testName string, sbs ...*sgroupsv1.ServiceBinding) *foundation.TestCaseBodySBUpsert {
	return &foundation.TestCaseBodySBUpsert{
		TestName: testName,
		Req:      sgroupsv1.ServiceBindingReq_Upsert{ServiceBindings: sbs},
	}
}

func sbEntry(metadata *common.Metadata, spec *sgroupsv1.ServiceBinding_Spec) *sgroupsv1.ServiceBinding {
	return &sgroupsv1.ServiceBinding{Metadata: metadata, Spec: spec}
}

// errNameSpec Spec, общий для большинства негативных кейсов добавления service-binding по имени
func errNameSpec() *sgroupsv1.ServiceBinding_Spec {
	return &sgroupsv1.ServiceBinding_Spec{
		AddressGroup: ri("ag-3", "namespace-1"),
		Service:      ri("svc-7", "namespace-1"),
	}
}

// upsertErrName Негативный кейс с невалидным именем service-binding (namespace-1, стандартный errNameSpec)
func upsertErrName(testName, name string) *foundation.TestCaseBodySBUpsert {
	return upsertSBBody(testName, sbEntry(&common.Metadata{Name: name, Namespace: "namespace-1"}, errNameSpec()))
}

// UpsertAdd Добавление нового serviceBindings (все в одном NS)
func UpsertAdd() *foundation.TestCaseBodySBUpsert {
	return upsertSBBody("Добавление нового serviceBindings (все в одном NS)", sbEntry(
		&common.Metadata{
			Name:        "add-sb-1",
			Namespace:   "namespace-1",
			Labels:      map[string]string{"add": "success"},
			Annotations: map[string]string{"add": "success"},
		},
		&sgroupsv1.ServiceBinding_Spec{
			Comment: "new service binding", Description: "add success", DisplayName: "new service binding 1",
			AddressGroup: ri("ag-6", "namespace-1"), Service: ri("svc-6", "namespace-1"),
		},
	))
}

// UpsertAddAGDifferentNamespace Добавление нового serviceBindings (AG в другом NS)
func UpsertAddAGDifferentNamespace() *foundation.TestCaseBodySBUpsert {
	return upsertSBBody("Добавление нового serviceBindings (AG в другом NS)", sbEntry(
		&common.Metadata{
			Name:        "add-sb-2",
			Namespace:   "namespace-1",
			Labels:      map[string]string{"add": "success"},
			Annotations: map[string]string{"add": "success"},
		},
		&sgroupsv1.ServiceBinding_Spec{
			Comment: "new service binding", Description: "add success", DisplayName: "new service binding 1",
			AddressGroup: ri("ag-2", "namespace-2"), Service: ri("svc-6", "namespace-1"),
		},
	))
}

// UpsertEdit Редактирование serviceBindings
func UpsertEdit() *foundation.TestCaseBodySBUpsert {
	return upsertSBBody("Редактирование serviceBindings", sbEntry(
		&common.Metadata{
			Name:        "sb-3",
			Namespace:   "namespace-1",
			Uid:         "a915bc00-2feb-42a5-94af-e209effbce44",
			Labels:      map[string]string{"edit": "success"},
			Annotations: map[string]string{"edit": "success"},
		},
		&sgroupsv1.ServiceBinding_Spec{
			Comment: "edited service binding", Description: "edit success", DisplayName: "edit service binding 1",
			AddressGroup: ri("ag-3", "namespace-1"), Service: ri("svc-3", "namespace-1"),
		},
	))
}

// UpsertAddSameNameDifferentNamespace Добавление serviceBindings с существующим именем, но к другому namespace
func UpsertAddSameNameDifferentNamespace() *foundation.TestCaseBodySBUpsert {
	return upsertSBBody("Добавление serviceBindings с существующим именем, но к другому namespace", sbEntry(
		&common.Metadata{
			Name:        "sb-0",
			Namespace:   "namespace-1",
			Labels:      map[string]string{"add": "success"},
			Annotations: map[string]string{"add": "success"},
		},
		&sgroupsv1.ServiceBinding_Spec{
			Comment: "new service binding", Description: "add success", DisplayName: "new service binding 1",
			AddressGroup: ri("ag-2", "namespace-2"), Service: ri("svc-7", "namespace-1"),
		},
	))
}

// UpsertErrorPortsOverlap Ошибка при добавлении, сервисы пересекаются портами
func UpsertErrorPortsOverlap() *foundation.TestCaseBodySBUpsert {
	return upsertSBBody("Ошибка при добавлении, сервисы пересекаются портами", sbEntry(
		&common.Metadata{
			Name:        "add-sb-3",
			Namespace:   "namespace-1",
			Labels:      map[string]string{"add": "success"},
			Annotations: map[string]string{"add": "success"},
		},
		&sgroupsv1.ServiceBinding_Spec{
			Comment: "new service binding", Description: "add success", DisplayName: "new service binding 1",
			AddressGroup: ri("ag-0", "namespace-0"), Service: ri("svc-7", "namespace-1"),
		},
	))
}

// UpsertErrorDisplayNameTooLong Ошибка при добавлении serviceBindings, displayName > 63 символов
func UpsertErrorDisplayNameTooLong() *foundation.TestCaseBodySBUpsert {
	return upsertSBBody("Ошибка при добавлении serviceBindings, displayName > 63 символов", sbEntry(
		&common.Metadata{Name: "err-display", Namespace: "namespace-1"},
		&sgroupsv1.ServiceBinding_Spec{
			DisplayName: tooLongSBName, AddressGroup: ri("ag-1", "namespace-1"), Service: ri("svc-7", "namespace-1"),
		},
	))
}

// UpsertErrorEditAGAndSVCMissing Ошибка при редактировании serviceBindings, ag и svc отсутствуют
func UpsertErrorEditAGAndSVCMissing() *foundation.TestCaseBodySBUpsert {
	return upsertSBBody("Ошибка при редактировании serviceBindings, ag и svc отсутствуют", sbEntry(
		&common.Metadata{
			Name: "sb-3", Namespace: "namespace-1", Uid: "a915bc00-2feb-42a5-94af-e209effbce44",
			Labels: map[string]string{"edit": "success"}, Annotations: map[string]string{"edit": "success"},
		},
		&sgroupsv1.ServiceBinding_Spec{},
	))
}

// editSpecBase Spec, общий для большинства негативных кейсов редактирования sb-3
func editSpecBase(ag, svc *common.ResourceIdentifier) *sgroupsv1.ServiceBinding_Spec {
	return &sgroupsv1.ServiceBinding_Spec{
		AddressGroup: ag, Service: svc,
	}
}

func editSB3(testName string, ag, svc *common.ResourceIdentifier) *foundation.TestCaseBodySBUpsert {
	return upsertSBBody(testName, sbEntry(
		&common.Metadata{
			Name: "sb-3", Namespace: "namespace-1", Uid: "a915bc00-2feb-42a5-94af-e209effbce44",
			Labels: map[string]string{"edit": "success"}, Annotations: map[string]string{"edit": "success"},
		},
		editSpecBase(ag, svc),
	))
}

// UpsertErrorEditSVCImmutable Ошибка при редактировании serviceBindings, нельзя редактировать svc
func UpsertErrorEditSVCImmutable() *foundation.TestCaseBodySBUpsert {
	return editSB3("Ошибка при редактировании serviceBindings, нельзя редактировать svc", ri("ag-3", "namespace-1"), ri("svc-1", "namespace-1"))
}

// UpsertErrorEditSVCImmutableDifferentNamespace Ошибка при редактировании serviceBindings, нельзя редактировать svc из другого ns
//
// Body в исходной Postman-коллекции содержит "service": {"name": "nw-2", "namespace": "namespace-2"} — копипаста-артефакт из NB, воспроизведён буквально
func UpsertErrorEditSVCImmutableDifferentNamespace() *foundation.TestCaseBodySBUpsert {
	return editSB3("Ошибка при редактировании serviceBindings, нельзя редактировать svc из другого ns", ri("ag-3", "namespace-1"), ri("nw-2", "namespace-2"))
}

// UpsertErrorEditSVCMissing Ошибка при редактировании serviceBindings, svc отсутствует
func UpsertErrorEditSVCMissing() *foundation.TestCaseBodySBUpsert {
	return upsertSBBody("Ошибка при редактировании serviceBindings, svc отсутствует", sbEntry(
		&common.Metadata{
			Name: "sb-3", Namespace: "namespace-1", Uid: "a915bc00-2feb-42a5-94af-e209effbce44",
			Labels: map[string]string{"edit": "success"}, Annotations: map[string]string{"edit": "success"},
		},
		&sgroupsv1.ServiceBinding_Spec{AddressGroup: ri("ag-3", "namespace-1")},
	))
}

// UpsertErrorEditAGImmutable Ошибка при редактировании serviceBindings, нельзя редактировать ag
func UpsertErrorEditAGImmutable() *foundation.TestCaseBodySBUpsert {
	return editSB3("Ошибка при редактировании serviceBindings, нельзя редактировать ag", ri("ag-6", "namespace-1"), ri("svc-3", "namespace-1"))
}

// UpsertErrorEditAGImmutableDifferentNamespace Ошибка при редактировании serviceBindings, нельзя редактировать ag из другого ns
func UpsertErrorEditAGImmutableDifferentNamespace() *foundation.TestCaseBodySBUpsert {
	return editSB3("Ошибка при редактировании serviceBindings, нельзя редактировать ag из другого ns", ri("ag-2", "namespace-2"), ri("svc-3", "namespace-1"))
}

// UpsertErrorEditAGMissing Ошибка при редактировании serviceBindings, ag отсутствует
func UpsertErrorEditAGMissing() *foundation.TestCaseBodySBUpsert {
	return upsertSBBody("Ошибка при редактировании serviceBindings, ag отсутствует", sbEntry(
		&common.Metadata{
			Name: "sb-3", Namespace: "namespace-1", Uid: "a915bc00-2feb-42a5-94af-e209effbce44",
			Labels: map[string]string{"edit": "success"}, Annotations: map[string]string{"edit": "success"},
		},
		&sgroupsv1.ServiceBinding_Spec{Service: ri("svc-3", "namespace-1")},
	))
}

// UpsertErrorBindingAlreadyExistsSameNamespace Ошибка при добавлении serviceBindings, такая привязка уже существует (svc-host в одном ns)
func UpsertErrorBindingAlreadyExistsSameNamespace() *foundation.TestCaseBodySBUpsert {
	return upsertSBBody("Ошибка при добавлении serviceBindings, такая привязка уже существует (svc-host в одном ns)", sbEntry(
		&common.Metadata{Name: "sb-1000", Namespace: "namespace-0"},
		&sgroupsv1.ServiceBinding_Spec{AddressGroup: ri("ag-0", "namespace-0"), Service: ri("svc-0", "namespace-0")},
	))
}

// UpsertErrorBindingAlreadyExistsDifferentNamespace Ошибка при добавлении serviceBindings, такая привязка уже существует (svc-host в разных ns)
//
// В исходной Postman-коллекции metadata.name = "nb-1000" — копипаста-артефакт из NB, воспроизведён буквально
func UpsertErrorBindingAlreadyExistsDifferentNamespace() *foundation.TestCaseBodySBUpsert {
	return upsertSBBody("Ошибка при добавлении serviceBindings, такая привязка уже существует (svc-host в разных ns)", sbEntry(
		&common.Metadata{Name: "nb-1000", Namespace: "namespace-1"},
		&sgroupsv1.ServiceBinding_Spec{AddressGroup: ri("ag-2", "namespace-2"), Service: ri("svc-6", "namespace-1")},
	))
}

// UpsertErrorAddDuplicateNameNamespace Ошибка при добавлении serviceBindings, sb с таким name+namespace уже существует
func UpsertErrorAddDuplicateNameNamespace() *foundation.TestCaseBodySBUpsert {
	return upsertSBBody("Ошибка при добавлении serviceBindings, sb с таким name+namespace уже существует", sbEntry(
		&common.Metadata{Name: "sb-0", Namespace: "namespace-0"},
		&sgroupsv1.ServiceBinding_Spec{AddressGroup: ri("ag-3", "namespace-1"), Service: ri("svc-0", "namespace-0")},
	))
}

// UpsertErrorSVCNamespaceMismatch Ошибка при добавлении serviceBindings, svc ns с не совпадает с sb ns
func UpsertErrorSVCNamespaceMismatch() *foundation.TestCaseBodySBUpsert {
	return upsertSBBody("Ошибка при добавлении serviceBindings, svc ns с не совпадает с sb ns", sbEntry(
		&common.Metadata{Name: "err-ns", Namespace: "namespace-1"},
		&sgroupsv1.ServiceBinding_Spec{AddressGroup: ri("ag-1", "namespace-1"), Service: ri("svc-2", "namespace-2")},
	))
}

// UpsertErrorAllDifferentNamespaces Ошибка при добавлении serviceBindings, все в разных NS
func UpsertErrorAllDifferentNamespaces() *foundation.TestCaseBodySBUpsert {
	return upsertSBBody("Ошибка при добавлении serviceBindings, все в разных NS", sbEntry(
		&common.Metadata{Name: "err-ns", Namespace: "namespace-3"},
		&sgroupsv1.ServiceBinding_Spec{AddressGroup: ri("ag-1", "namespace-1"), Service: ri("svc-0", "namespace-0")},
	))
}

// UpsertErrorEditLongUID Ошибка при редактировании serviceBindings, uid длинный
//
// Body в исходной Postman-коллекции содержит name="nb-3", uid="efb85252-c9b7-4dd5-b9be-6ce63276795vc" и service.name="ымс-3" —
// копипаста-артефакты из HB/NB, воспроизведены буквально; проверяется только длина uid, остальные поля роли не играют
func UpsertErrorEditLongUID() *foundation.TestCaseBodySBUpsert {
	return upsertSBBody("Ошибка при редактировании serviceBindings, uid длинный", sbEntry(
		&common.Metadata{Name: "nb-3", Namespace: "namespace-1", Uid: "efb85252-c9b7-4dd5-b9be-6ce63276795vc"},
		&sgroupsv1.ServiceBinding_Spec{AddressGroup: ri("ag-3", "namespace-1"), Service: ri("ымс-3", "namespace-1")},
	))
}

// UpsertErrorEditShortUID Ошибка при редактировании serviceBindings, uid короткий
//
// Body в исходной Postman-коллекции содержит name="nb-3", uid="efb85252-c9b7-4dd5-b9be-6ce63276795" — копипаста-артефакт из HB/NB, воспроизведён буквально
func UpsertErrorEditShortUID() *foundation.TestCaseBodySBUpsert {
	return upsertSBBody("Ошибка при редактировании serviceBindings, uid короткий", sbEntry(
		&common.Metadata{Name: "nb-3", Namespace: "namespace-1", Uid: "efb85252-c9b7-4dd5-b9be-6ce63276795"},
		&sgroupsv1.ServiceBinding_Spec{AddressGroup: ri("ag-3", "namespace-1"), Service: ri("svc-3", "namespace-1")},
	))
}

// UpsertErrorEditUIDMismatch Ошибка при редактировании serviceBindings, uid не соответствует указанному sb
func UpsertErrorEditUIDMismatch() *foundation.TestCaseBodySBUpsert {
	return upsertSBBody("Ошибка при редактировании serviceBindings, uid не соответствует указанному sb", sbEntry(
		&common.Metadata{Name: "sb-3", Namespace: "namespace-1", Uid: "7e13da25-2e49-4bae-bc5e-aa129870bc8a"},
		&sgroupsv1.ServiceBinding_Spec{AddressGroup: ri("ag-3", "namespace-1"), Service: ri("svc-3", "namespace-1")},
	))
}

// UpsertErrorEditNameUIDNotExistInNamespace Ошибка при редактировании serviceBindings, name+uid не существует в указанном namespace
func UpsertErrorEditNameUIDNotExistInNamespace() *foundation.TestCaseBodySBUpsert {
	return upsertSBBody("Ошибка при редактировании serviceBindings, name+uid не существует в указанном namespace", sbEntry(
		&common.Metadata{Name: "sb-0", Namespace: "namespace-1", Uid: "b724836e-0bbd-4cec-afa7-a8b8d45d4d6f"},
		&sgroupsv1.ServiceBinding_Spec{AddressGroup: ri("ag-3", "namespace-1"), Service: ri("svc-3", "namespace-1")},
	))
}

// UpsertErrorEditEmptyName Ошибка при редактировании serviceBindings, name пуст
func UpsertErrorEditEmptyName() *foundation.TestCaseBodySBUpsert {
	return upsertSBBody("Ошибка при редактировании serviceBindings, name пуст", sbEntry(
		&common.Metadata{Name: "", Namespace: "namespace-0", Uid: "b724836e-0bbd-4cec-afa7-a8b8d45d4d6f"},
		&sgroupsv1.ServiceBinding_Spec{AddressGroup: ri("ag-3", "namespace-1"), Service: ri("svc-3", "namespace-1")},
	))
}

// UpsertErrorEditNameMissing Ошибка при редактировании serviceBindings, name отсутствует как параметр
func UpsertErrorEditNameMissing() *foundation.TestCaseBodySBUpsert {
	return upsertSBBody("Ошибка при редактировании serviceBindings, name отсутствует как параметр", sbEntry(
		&common.Metadata{Namespace: "namespace-0", Uid: "1085d231-9a0f-4697-b864-c0a522181911"},
		&sgroupsv1.ServiceBinding_Spec{AddressGroup: ri("ag-3", "namespace-1"), Service: ri("svc-3", "namespace-1")},
	))
}

// UpsertErrorEditEmptyNamespace Ошибка при редактировании serviceBindings, namespace пуст
func UpsertErrorEditEmptyNamespace() *foundation.TestCaseBodySBUpsert {
	return upsertSBBody("Ошибка при редактировании serviceBindings, namespace пуст", sbEntry(
		&common.Metadata{Name: "sb-0", Namespace: "", Uid: "b724836e-0bbd-4cec-afa7-a8b8d45d4d6f"},
		&sgroupsv1.ServiceBinding_Spec{AddressGroup: ri("ag-3", "namespace-1"), Service: ri("svc-3", "namespace-1")},
	))
}

// UpsertErrorEditNamespaceMissing Ошибка при редактировании serviceBindings, namespace отсутствует как параметр
//
// Body в исходной Postman-коллекции содержит service.name="nw-3" — копипаста-артефакт из NB, воспроизведён буквально
func UpsertErrorEditNamespaceMissing() *foundation.TestCaseBodySBUpsert {
	return upsertSBBody("Ошибка при редактировании serviceBindings, namespace отсутствует как параметр", sbEntry(
		&common.Metadata{Name: "sb-0", Uid: "b724836e-0bbd-4cec-afa7-a8b8d45d4d6f"},
		&sgroupsv1.ServiceBinding_Spec{AddressGroup: ri("ag-3", "namespace-1"), Service: ri("nw-3", "namespace-1")},
	))
}

// UpsertErrorAddEmptyName Ошибка при добавлении serviceBindings, name пуст
func UpsertErrorAddEmptyName() *foundation.TestCaseBodySBUpsert {
	return upsertSBBody("Ошибка при добавлении serviceBindings, name пуст", sbEntry(
		&common.Metadata{Name: "", Namespace: "namespace-1"},
		&sgroupsv1.ServiceBinding_Spec{AddressGroup: ri("ag-3", "namespace-1"), Service: ri("svc-7", "namespace-1")},
	))
}

// UpsertErrorNameCyrillic Ошибка при добавлении serviceBindings, name на кириллице
func UpsertErrorNameCyrillic() *foundation.TestCaseBodySBUpsert {
	return upsertErrName("Ошибка при добавлении serviceBindings, name на кириллице", "биндинг")
}

// UpsertErrorNameUpperCase Ошибка при добавлении serviceBindings, name капсом
func UpsertErrorNameUpperCase() *foundation.TestCaseBodySBUpsert {
	return upsertErrName("Ошибка при добавлении serviceBindings, name капсом", "BINDING")
}

// UpsertErrorNameSpecialCharsOnly Ошибка при добавлении serviceBindings, name из спец символов
func UpsertErrorNameSpecialCharsOnly() *foundation.TestCaseBodySBUpsert {
	return upsertErrName("Ошибка при добавлении serviceBindings, name из спец символов", "???!!!")
}

// UpsertErrorNameMixedCase Ошибка при добавлении serviceBindings, name капс+строчные
func UpsertErrorNameMixedCase() *foundation.TestCaseBodySBUpsert {
	return upsertErrName("Ошибка при добавлении serviceBindings, name капс+строчные", "Binding")
}

// UpsertErrorNameCyrillicLatin Ошибка при добавлении serviceBindings, name на кириллице+латинице
func UpsertErrorNameCyrillicLatin() *foundation.TestCaseBodySBUpsert {
	return upsertErrName("Ошибка при добавлении serviceBindings, name на кириллице+латинице", "binдинг")
}

// UpsertErrorNameCyrillicDigits Ошибка при добавлении serviceBindings, name на кириллице+числа
func UpsertErrorNameCyrillicDigits() *foundation.TestCaseBodySBUpsert {
	return upsertErrName("Ошибка при добавлении serviceBindings, name на кириллице+числа", "биндинг-1")
}

// UpsertErrorNameStartsWithHyphen Ошибка при добавлении serviceBindings, name начинается с тире
func UpsertErrorNameStartsWithHyphen() *foundation.TestCaseBodySBUpsert {
	return upsertErrName("Ошибка при добавлении serviceBindings, name начинается с тире", "-bind")
}

// UpsertErrorNameOnlyHyphen Ошибка при добавлении serviceBindings, name состоит только из тире
func UpsertErrorNameOnlyHyphen() *foundation.TestCaseBodySBUpsert {
	return upsertErrName("Ошибка при добавлении serviceBindings, name состоит только из тире", "-")
}

// UpsertErrorNameEndsWithHyphen Ошибка при добавлении serviceBindings, name заканчивается тире
func UpsertErrorNameEndsWithHyphen() *foundation.TestCaseBodySBUpsert {
	return upsertErrName("Ошибка при добавлении serviceBindings, name заканчивается тире", "bind-")
}

// UpsertErrorNameCyrillicSpecialChars Ошибка при добавлении serviceBindings, name на кириллице+спец.символы
func UpsertErrorNameCyrillicSpecialChars() *foundation.TestCaseBodySBUpsert {
	return upsertErrName("Ошибка при добавлении serviceBindings, name на кириллице+спец.символы", "биндинг?")
}

// UpsertErrorNameTooLong Ошибка при добавлении serviceBindings, name > 63 символов
func UpsertErrorNameTooLong() *foundation.TestCaseBodySBUpsert {
	return upsertErrName("Ошибка при добавлении serviceBindings, name > 63 символов", tooLongSBName)
}

// UpsertErrorNameLeadingSpace Ошибка при добавлении serviceBindings, пробел в начале name
func UpsertErrorNameLeadingSpace() *foundation.TestCaseBodySBUpsert {
	return upsertErrName("Ошибка при добавлении serviceBindings, пробел в начале name", " bind-1")
}

// UpsertErrorNameTrailingSpace Ошибка при добавлении serviceBindings, пробел в конце name
func UpsertErrorNameTrailingSpace() *foundation.TestCaseBodySBUpsert {
	return upsertErrName("Ошибка при добавлении serviceBindings, пробел в конце name", "bind-1 ")
}

// UpsertErrorNameWithSpaces Ошибка при добавлении serviceBindings, name с пробелами
func UpsertErrorNameWithSpaces() *foundation.TestCaseBodySBUpsert {
	return upsertErrName("Ошибка при добавлении serviceBindings, name с пробелами", "bind - 1")
}

// UpsertErrorNamespaceNotExist Ошибка при добавлении serviceBindings, namespace не существует
func UpsertErrorNamespaceNotExist() *foundation.TestCaseBodySBUpsert {
	return upsertSBBody("Ошибка при добавлении serviceBindings, namespace не существует", sbEntry(
		&common.Metadata{Name: "bind", Namespace: "namespace-11"},
		&sgroupsv1.ServiceBinding_Spec{AddressGroup: ri("ag-3", "namespace-1"), Service: ri("svc-7", "namespace-1")},
	))
}

// UpsertErrorNamespaceEmpty Ошибка при добавлении serviceBindings, namespace пуст
//
// Body в исходной Postman-коллекции содержит service.name="nw-7" — копипаста-артефакт из NB, воспроизведён буквально
func UpsertErrorNamespaceEmpty() *foundation.TestCaseBodySBUpsert {
	return upsertSBBody("Ошибка при добавлении serviceBindings, namespace пуст", sbEntry(
		&common.Metadata{Name: "bind", Namespace: ""},
		&sgroupsv1.ServiceBinding_Spec{AddressGroup: ri("ag-3", "namespace-1"), Service: ri("nw-7", "namespace-1")},
	))
}

// UpsertErrorNamespaceLeadingSpace Ошибка при добавлении serviceBindings, пробел в начале имени namespace
func UpsertErrorNamespaceLeadingSpace() *foundation.TestCaseBodySBUpsert {
	return upsertSBBody("Ошибка при добавлении serviceBindings, пробел в начале имени namespace", sbEntry(
		&common.Metadata{Name: "bind", Namespace: " namespace-1"},
		&sgroupsv1.ServiceBinding_Spec{AddressGroup: ri("ag-3", "namespace-1"), Service: ri("svc-7", "namespace-1")},
	))
}

// UpsertErrorNamespaceTrailingSpace Ошибка при добавлении networkBindings, пробел в конце имени namespace
//
// Название кейса в исходной Postman-коллекции — "Ошибка при добавлении networkBindings..." — копипаста-артефакт из NB, воспроизведён буквально в t.Story
func UpsertErrorNamespaceTrailingSpace() *foundation.TestCaseBodySBUpsert {
	return upsertSBBody("Ошибка при добавлении networkBindings, пробел в конце имени namespace", sbEntry(
		&common.Metadata{Name: "bind", Namespace: "namespace-1 "},
		&sgroupsv1.ServiceBinding_Spec{AddressGroup: ri("ag-3", "namespace-1"), Service: ri("svc-7", "namespace-1")},
	))
}

// UpsertErrorNamespaceWithSpaces Ошибка при добавлении serviceBindings, namespace с пробелами
func UpsertErrorNamespaceWithSpaces() *foundation.TestCaseBodySBUpsert {
	return upsertSBBody("Ошибка при добавлении serviceBindings, namespace с пробелами", sbEntry(
		&common.Metadata{Name: "bind", Namespace: "namespace - 1"},
		&sgroupsv1.ServiceBinding_Spec{AddressGroup: ri("ag-3", "namespace-1"), Service: ri("svc-7", "namespace-1")},
	))
}

// UpsertErrorNamespaceMissing Ошибка при добавлении serviceBindings, namespace отсутствует как параметр
func UpsertErrorNamespaceMissing() *foundation.TestCaseBodySBUpsert {
	return upsertSBBody("Ошибка при добавлении serviceBindings, namespace отсутствует как параметр", sbEntry(
		&common.Metadata{Name: "bind"},
		&sgroupsv1.ServiceBinding_Spec{AddressGroup: ri("ag-3", "namespace-1"), Service: ri("svc-7", "namespace-1")},
	))
}

// UpsertErrorSVCNotExistInNamespace Ошибка при добавлении serviceBindings, svc не существует в указанном ns
func UpsertErrorSVCNotExistInNamespace() *foundation.TestCaseBodySBUpsert {
	return upsertSBBody("Ошибка при добавлении serviceBindings, svc не существует в указанном ns", sbEntry(
		&common.Metadata{Name: "err", Namespace: "namespace-1"},
		&sgroupsv1.ServiceBinding_Spec{AddressGroup: ri("ag-3", "namespace-1"), Service: ri("svc-77", "namespace-1")},
	))
}

// UpsertErrorSVCNotInNamespaceButExists Ошибка при добавлении serviceBindings, в указанном ns нет такого svc (сам svc существует)
func UpsertErrorSVCNotInNamespaceButExists() *foundation.TestCaseBodySBUpsert {
	return upsertSBBody("Ошибка при добавлении serviceBindings, в указанном ns нет такого svc (сам svc существует)", sbEntry(
		&common.Metadata{Name: "err", Namespace: "namespace-1"},
		&sgroupsv1.ServiceBinding_Spec{AddressGroup: ri("ag-3", "namespace-1"), Service: ri("svc-2", "namespace-1")},
	))
}

// UpsertErrorSVCNameEmpty Ошибка при добавлении serviceBindings, svc name пуст
func UpsertErrorSVCNameEmpty() *foundation.TestCaseBodySBUpsert {
	return upsertSBBody("Ошибка при добавлении serviceBindings, svc name пуст", sbEntry(
		&common.Metadata{Name: "err", Namespace: "namespace-1"},
		&sgroupsv1.ServiceBinding_Spec{AddressGroup: ri("ag-3", "namespace-1"), Service: ri("", "namespace-1")},
	))
}

// UpsertErrorSVCEmpty Ошибка при добавлении serviceBindings, svc пуст (ns и name)
func UpsertErrorSVCEmpty() *foundation.TestCaseBodySBUpsert {
	return upsertSBBody("Ошибка при добавлении serviceBindings, svc пуст (ns и name)", sbEntry(
		&common.Metadata{Name: "err", Namespace: "namespace-1"},
		&sgroupsv1.ServiceBinding_Spec{AddressGroup: ri("ag-3", "namespace-1"), Service: ri("", "")},
	))
}

// UpsertErrorSVCNameMissing Ошибка при добавлении serviceBindings, в svc name отсутствует как параметр
func UpsertErrorSVCNameMissing() *foundation.TestCaseBodySBUpsert {
	return upsertSBBody("Ошибка при добавлении serviceBindings, в svc name отсутствует как параметр", sbEntry(
		&common.Metadata{Name: "err", Namespace: "namespace-1"},
		&sgroupsv1.ServiceBinding_Spec{AddressGroup: ri("ag-3", "namespace-1"), Service: ri("", "namespace-1")},
	))
}

// UpsertErrorSVCNamespaceMissing Ошибка при добавлении serviceBindings, в svc ns отсутствует как параметр
//
// Body в исходной Postman-коллекции содержит service.name="nw-1" — копипаста-артефакт из NB, воспроизведён буквально
func UpsertErrorSVCNamespaceMissing() *foundation.TestCaseBodySBUpsert {
	return upsertSBBody("Ошибка при добавлении serviceBindings, в svc ns отсутствует как параметр", sbEntry(
		&common.Metadata{Name: "err", Namespace: "namespace-1"},
		&sgroupsv1.ServiceBinding_Spec{AddressGroup: ri("ag-3", "namespace-1"), Service: ri("nw-1", "")},
	))
}

// UpsertErrorSVCEmptyObject Ошибка при добавлении serviceBindings, svc пустой объект
func UpsertErrorSVCEmptyObject() *foundation.TestCaseBodySBUpsert {
	return upsertSBBody("Ошибка при добавлении serviceBindings, svc пустой объект", sbEntry(
		&common.Metadata{Name: "err", Namespace: "namespace-1"},
		&sgroupsv1.ServiceBinding_Spec{AddressGroup: ri("ag-3", "namespace-1"), Service: &common.ResourceIdentifier{}},
	))
}

// UpsertErrorSVCMissing Ошибка при добавлении serviceBindings, svc отсутствует как параметр
func UpsertErrorSVCMissing() *foundation.TestCaseBodySBUpsert {
	return upsertSBBody("Ошибка при добавлении serviceBindings, svc отсутствует как параметр", sbEntry(
		&common.Metadata{Name: "err", Namespace: "namespace-1"},
		&sgroupsv1.ServiceBinding_Spec{AddressGroup: ri("ag-3", "namespace-1")},
	))
}

// UpsertErrorAGEmptyObject Ошибка при добавлении serviceBindings, ag пустой объект
func UpsertErrorAGEmptyObject() *foundation.TestCaseBodySBUpsert {
	return upsertSBBody("Ошибка при добавлении serviceBindings, ag пустой объект", sbEntry(
		&common.Metadata{Name: "err", Namespace: "namespace-1"},
		&sgroupsv1.ServiceBinding_Spec{AddressGroup: &common.ResourceIdentifier{}, Service: ri("svc-1", "namespace-1")},
	))
}

// UpsertErrorAGMissing Ошибка при добавлении serviceBindings, ag отсутствует как параметр
func UpsertErrorAGMissing() *foundation.TestCaseBodySBUpsert {
	return upsertSBBody("Ошибка при добавлении serviceBindings, ag отсутствует как параметр", sbEntry(
		&common.Metadata{Name: "err", Namespace: "namespace-1"},
		&sgroupsv1.ServiceBinding_Spec{Service: ri("svc-3", "namespace-1")},
	))
}

// UpsertErrorAGNamespaceMissing Ошибка при добавлении serviceBindings, в ag ns отсутствует как параметр
//
// Body в исходной Postman-коллекции содержит service.name="nw-1" — копипаста-артефакт из NB, воспроизведён буквально
func UpsertErrorAGNamespaceMissing() *foundation.TestCaseBodySBUpsert {
	return upsertSBBody("Ошибка при добавлении serviceBindings, в ag ns отсутствует как параметр", sbEntry(
		&common.Metadata{Name: "err", Namespace: "namespace-1"},
		&sgroupsv1.ServiceBinding_Spec{AddressGroup: ri("ag-3", ""), Service: ri("nw-1", "namespace-1")},
	))
}

// UpsertErrorAGNameMissing Ошибка при добавлении serviceBindings, в ag name отсутствует как параметр
//
// Body в исходной Postman-коллекции содержит service.name="nw-1" — копипаста-артефакт из NB, воспроизведён буквально
func UpsertErrorAGNameMissing() *foundation.TestCaseBodySBUpsert {
	return upsertSBBody("Ошибка при добавлении serviceBindings, в ag name отсутствует как параметр", sbEntry(
		&common.Metadata{Name: "err", Namespace: "namespace-1"},
		&sgroupsv1.ServiceBinding_Spec{AddressGroup: ri("", "namespace-1"), Service: ri("nw-1", "namespace-1")},
	))
}

// UpsertErrorAGEmpty Ошибка при добавлении serviceBindings, ag пуст (ns и name)
func UpsertErrorAGEmpty() *foundation.TestCaseBodySBUpsert {
	return upsertSBBody("Ошибка при добавлении serviceBindings, ag пуст (ns и name)", sbEntry(
		&common.Metadata{Name: "err", Namespace: "namespace-1"},
		&sgroupsv1.ServiceBinding_Spec{AddressGroup: ri("", ""), Service: ri("svc-1", "namespace-1")},
	))
}

// UpsertErrorAGNameEmpty Ошибка при добавлении serviceBindings, ag name пуст
func UpsertErrorAGNameEmpty() *foundation.TestCaseBodySBUpsert {
	return upsertSBBody("Ошибка при добавлении serviceBindings, ag name пуст", sbEntry(
		&common.Metadata{Name: "err", Namespace: "namespace-1"},
		&sgroupsv1.ServiceBinding_Spec{AddressGroup: ri("", "namespace-1"), Service: ri("svc-1", "namespace-1")},
	))
}

// UpsertErrorAGNotInNamespaceButExists Ошибка при добавлении serviceBindings, в указанном ns нет такого ag (сам ag существует)
func UpsertErrorAGNotInNamespaceButExists() *foundation.TestCaseBodySBUpsert {
	return upsertSBBody("Ошибка при добавлении serviceBindings, в указанном ns нет такого ag (сам ag существует)", sbEntry(
		&common.Metadata{Name: "err", Namespace: "namespace-1"},
		&sgroupsv1.ServiceBinding_Spec{AddressGroup: ri("ag-2", "namespace-1"), Service: ri("svc-1", "namespace-1")},
	))
}

// UpsertErrorAGNotExist Ошибка при добавлении serviceBindings, ag не существует в указанном ns
func UpsertErrorAGNotExist() *foundation.TestCaseBodySBUpsert {
	return upsertSBBody("Ошибка при добавлении serviceBindings, ag не существует в указанном ns", sbEntry(
		&common.Metadata{Name: "err", Namespace: "namespace-1"},
		&sgroupsv1.ServiceBinding_Spec{AddressGroup: ri("ag-33", "namespace-1"), Service: ri("svc-7", "namespace-1")},
	))
}

// UpsertErrorSpecEmpty Ошибка при добавлении serviceBindings, spec пустой объект
func UpsertErrorSpecEmpty() *foundation.TestCaseBodySBUpsert {
	return upsertSBBody("Ошибка при добавлении serviceBindings, spec пустой объект", sbEntry(
		&common.Metadata{Name: "err", Namespace: "namespace-1"},
		&sgroupsv1.ServiceBinding_Spec{},
	))
}
