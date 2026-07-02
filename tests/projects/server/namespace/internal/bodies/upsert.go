package bodies

import (
	"strings"

	sgroupsv1 "HBF-tests/pkg/api/sgroups/v1"
	"HBF-tests/tests/common/foundation"

	common "github.com/PRO-Robotech/sgroups-proto/pkg/api/common"
)

// tooLongName 64-символьная строка из строчных латинских букв для проверки ограничения длины в 63 символа
var tooLongName = strings.Repeat("a", 64)

// UpsertAddOne Добавление одного нового namespace
func UpsertAddOne() *foundation.TestCaseBodyUpsert {
	return &foundation.TestCaseBodyUpsert{
		TestName: "Добавление одного нового namespace",
		Req: sgroupsv1.NamespaceReq_Upsert{
			Namespaces: []*sgroupsv1.Namespace{
				{
					Metadata: &common.Metadata{
						Name:        "add-namespace-1",
						Labels:      map[string]string{"add": "success"},
						Annotations: map[string]string{"add": "success"},
					},
					Spec: &sgroupsv1.Namespace_Spec{
						Comment:     "for upsert tests",
						Description: "add new namespace",
						DisplayName: "New Namespace",
					},
				},
			},
		},
	}
}

// UpsertAddTwo Добавление 2 новых namespaces
func UpsertAddTwo() *foundation.TestCaseBodyUpsert {
	return &foundation.TestCaseBodyUpsert{
		TestName: "Добавление 2 новых namespaces",
		Req: sgroupsv1.NamespaceReq_Upsert{
			Namespaces: []*sgroupsv1.Namespace{
				{
					Metadata: &common.Metadata{
						Name:        "add-namespace-2",
						Labels:      map[string]string{},
						Annotations: map[string]string{"add": "success"},
					},
					Spec: &sgroupsv1.Namespace_Spec{},
				},
				{
					Metadata: &common.Metadata{
						Name:        "add-namespace-3",
						Labels:      map[string]string{"add": "success"},
						Annotations: map[string]string{},
					},
					Spec: &sgroupsv1.Namespace_Spec{
						Comment:     "for upsert tests",
						Description: "add new namespace 2",
						DisplayName: "New Namespace 3",
					},
				},
			},
		},
	}
}

// UpsertEditByUID Редактирование namespace по uid
func UpsertEditByUID() *foundation.TestCaseBodyUpsert {
	return &foundation.TestCaseBodyUpsert{
		TestName: "Редактирование namespace по uid",
		Req: sgroupsv1.NamespaceReq_Upsert{
			Namespaces: []*sgroupsv1.Namespace{
				{
					Metadata: &common.Metadata{
						Name:        "namespace-3",
						Uid:         "cd3e3c34-bf87-4787-87e8-7ba7182280c3",
						Labels:      map[string]string{"edit": "success"},
						Annotations: map[string]string{"edit": "success"},
					},
					Spec: &sgroupsv1.Namespace_Spec{
						Comment:     "for upsert tests",
						Description: "edit namespace",
						DisplayName: "edit success",
					},
				},
			},
		},
	}
}

// UpsertErrorDuplicateNameWithoutUID Ошибка при добавлении, namespace с таким именем уже существует (не указан uid)
func UpsertErrorDuplicateNameWithoutUID() *foundation.TestCaseBodyUpsert {
	return &foundation.TestCaseBodyUpsert{
		TestName: "Ошибка при добавлении, namespace с таким именем уже существует (не указан uid)",
		Req: sgroupsv1.NamespaceReq_Upsert{
			Namespaces: []*sgroupsv1.Namespace{
				{
					Metadata: &common.Metadata{Name: "namespace-0"},
					Spec: &sgroupsv1.Namespace_Spec{
						Comment:     "for upsert tests",
						Description: "add new namespace",
						DisplayName: "New Namespace",
					},
				},
			},
		},
	}
}

// UpsertErrorEmptyName Ошибка при добавлении, name пустая строка
func UpsertErrorEmptyName() *foundation.TestCaseBodyUpsert {
	return &foundation.TestCaseBodyUpsert{
		TestName: "Ошибка при добавлении, name пустая строка",
		Req: sgroupsv1.NamespaceReq_Upsert{
			Namespaces: []*sgroupsv1.Namespace{
				{
					Metadata: &common.Metadata{Name: ""},
					Spec: &sgroupsv1.Namespace_Spec{
						Comment:     "for upsert tests",
						Description: "add new namespace",
						DisplayName: "New Namespace",
					},
				},
			},
		},
	}
}

// UpsertErrorNameMissing Ошибка при добавлении, name отсутствует как параметр
func UpsertErrorNameMissing() *foundation.TestCaseBodyUpsert {
	return &foundation.TestCaseBodyUpsert{
		TestName: "Ошибка при добавлении, name отсутствует как параметр",
		Req: sgroupsv1.NamespaceReq_Upsert{
			Namespaces: []*sgroupsv1.Namespace{
				{
					Metadata: &common.Metadata{Labels: map[string]string{"key": "value"}},
					Spec: &sgroupsv1.Namespace_Spec{
						Comment:     "for upsert tests",
						Description: "add new namespace",
						DisplayName: "New Namespace",
					},
				},
			},
		},
	}
}

// upsertWithInvalidName вспомогательный конструктор тела запроса с невалидным именем namespace
func upsertWithInvalidName(testName, name string) *foundation.TestCaseBodyUpsert {
	return &foundation.TestCaseBodyUpsert{
		TestName: testName,
		Req: sgroupsv1.NamespaceReq_Upsert{
			Namespaces: []*sgroupsv1.Namespace{
				{
					Metadata: &common.Metadata{Name: name},
					Spec: &sgroupsv1.Namespace_Spec{
						Comment:     "for upsert tests",
						Description: "add new namespace",
						DisplayName: "New Namespace",
					},
				},
			},
		},
	}
}

// UpsertErrorNameCyrillic Ошибка при добавлении, name на кириллице
func UpsertErrorNameCyrillic() *foundation.TestCaseBodyUpsert {
	return upsertWithInvalidName("Ошибка при добавлении, name на кириллице", "неймспейс")
}

// UpsertErrorNameCyrillicDigits Ошибка при добавлении, name на кириллице+цифры
func UpsertErrorNameCyrillicDigits() *foundation.TestCaseBodyUpsert {
	return upsertWithInvalidName("Ошибка при добавлении, name на кириллице+цифры", "неймспейс-1")
}

// UpsertErrorNameCyrillicSpecialChars Ошибка при добавлении, name на кириллице+спец.символы
func UpsertErrorNameCyrillicSpecialChars() *foundation.TestCaseBodyUpsert {
	return upsertWithInvalidName("Ошибка при добавлении, name на кириллице+спец.символы", "неймспейс?!,")
}

// UpsertErrorNameLatinSpecialChars Ошибка при добавлении, name на латинице+спец.символы
func UpsertErrorNameLatinSpecialChars() *foundation.TestCaseBodyUpsert {
	return upsertWithInvalidName("Ошибка при добавлении, name на латинице+спец.символы", "namespace?!,")
}

// UpsertErrorNameUpperCase Ошибка при добавлении, name капсом
func UpsertErrorNameUpperCase() *foundation.TestCaseBodyUpsert {
	return upsertWithInvalidName("Ошибка при добавлении, name капсом", "NAMESPACE")
}

// UpsertErrorNameStartsWithHyphen Ошибка при добавлении, name начинается с дефиса
func UpsertErrorNameStartsWithHyphen() *foundation.TestCaseBodyUpsert {
	return upsertWithInvalidName("Ошибка при добавлении, name начинается с дефиса", "-namespace")
}

// UpsertErrorNameEndsWithHyphen Ошибка при добавлении, name заканчивается дефисом
func UpsertErrorNameEndsWithHyphen() *foundation.TestCaseBodyUpsert {
	return upsertWithInvalidName("Ошибка при добавлении, name заканчивается дефисом", "namespace-")
}

// UpsertErrorNameSpecialCharsOnly Ошибка при добавлении, name из спец.символов
func UpsertErrorNameSpecialCharsOnly() *foundation.TestCaseBodyUpsert {
	return upsertWithInvalidName("Ошибка при добавлении, name из спец.символов", "???!!!")
}

// UpsertErrorNameLeadingSpace Ошибка при добавлении, пробел в начале name
func UpsertErrorNameLeadingSpace() *foundation.TestCaseBodyUpsert {
	return upsertWithInvalidName("Ошибка при добавлении, пробел в начале name", " namespace")
}

// UpsertErrorNameTrailingSpace Ошибка при добавлении, пробел в конце name
func UpsertErrorNameTrailingSpace() *foundation.TestCaseBodyUpsert {
	return upsertWithInvalidName("Ошибка при добавлении, пробел в конце name", "namespace ")
}

// UpsertErrorNameWithSpaces Ошибка при добавлении, name с пробелами
func UpsertErrorNameWithSpaces() *foundation.TestCaseBodyUpsert {
	return upsertWithInvalidName("Ошибка при добавлении, name с пробелами", "namespace space")
}

// UpsertErrorNameTooLong Ошибка при добавлении, name > 63 символов
func UpsertErrorNameTooLong() *foundation.TestCaseBodyUpsert {
	return upsertWithInvalidName("Ошибка при добавлении, name > 63 символов", tooLongName)
}

// UpsertErrorMetadataEmpty Ошибка при добавлении, metadata пустой объект
func UpsertErrorMetadataEmpty() *foundation.TestCaseBodyUpsert {
	return &foundation.TestCaseBodyUpsert{
		TestName: "Ошибка при добавлении, metadata пустой объект",
		Req: sgroupsv1.NamespaceReq_Upsert{
			Namespaces: []*sgroupsv1.Namespace{
				{
					Metadata: &common.Metadata{},
					Spec: &sgroupsv1.Namespace_Spec{
						Comment:     "for upsert tests",
						Description: "add new namespace",
						DisplayName: "New Namespace",
					},
				},
			},
		},
	}
}

// UpsertErrorMetadataMissing Ошибка при добавлении, metadata отсутствует как параметр
func UpsertErrorMetadataMissing() *foundation.TestCaseBodyUpsert {
	return &foundation.TestCaseBodyUpsert{
		TestName: "Ошибка при добавлении, metadata отсутствует как параметр",
		Req: sgroupsv1.NamespaceReq_Upsert{
			Namespaces: []*sgroupsv1.Namespace{
				{
					Spec: &sgroupsv1.Namespace_Spec{
						Comment:     "for upsert tests",
						Description: "add new namespace",
						DisplayName: "New Namespace",
					},
				},
			},
		},
	}
}

// UpsertErrorDisplayNameTooLong Ошибка при добавлении, displayName > 63 символов
func UpsertErrorDisplayNameTooLong() *foundation.TestCaseBodyUpsert {
	return &foundation.TestCaseBodyUpsert{
		TestName: "Ошибка при добавлении, displayName > 63 символов",
		Req: sgroupsv1.NamespaceReq_Upsert{
			Namespaces: []*sgroupsv1.Namespace{
				{
					Metadata: &common.Metadata{Name: "namespace-err-display"},
					Spec: &sgroupsv1.Namespace_Spec{
						Comment:     "for upsert tests",
						Description: "add new namespace",
						DisplayName: tooLongName,
					},
				},
			},
		},
	}
}

// UpsertErrorEditNonExistentUID Ошибка при редактировании, namespace с таким uid не существует
func UpsertErrorEditNonExistentUID() *foundation.TestCaseBodyUpsert {
	return &foundation.TestCaseBodyUpsert{
		TestName: "Ошибка при редактировании, namespace с таким uid не существует",
		Req: sgroupsv1.NamespaceReq_Upsert{
			Namespaces: []*sgroupsv1.Namespace{
				{
					Metadata: &common.Metadata{
						Name:        "namespace-33",
						Uid:         "cd3e3c34-bf87-4787-87e8-7ba7182280c3",
						Labels:      map[string]string{"edit": "err"},
						Annotations: map[string]string{"edit": "err"},
					},
					Spec: &sgroupsv1.Namespace_Spec{
						Comment:     "for upsert tests",
						Description: "edit namespace",
						DisplayName: "edit err",
					},
				},
			},
		},
	}
}

// UpsertErrorEditShortUID Ошибка при редактировании, короткий uid
func UpsertErrorEditShortUID() *foundation.TestCaseBodyUpsert {
	return &foundation.TestCaseBodyUpsert{
		TestName: "Ошибка при редактировании, короткий uid",
		Req: sgroupsv1.NamespaceReq_Upsert{
			Namespaces: []*sgroupsv1.Namespace{
				{
					Metadata: &common.Metadata{
						Name:        "namespace-3",
						Uid:         "321e83ed-7dfa-40b2-bff5-f86129477c4",
						Labels:      map[string]string{"edit": "err"},
						Annotations: map[string]string{"edit": "err"},
					},
					Spec: &sgroupsv1.Namespace_Spec{
						Comment:     "for upsert tests",
						Description: "edit namespace",
						DisplayName: "edit err",
					},
				},
			},
		},
	}
}

// UpsertErrorEditLongUID Ошибка при редактировании, длинный uid
func UpsertErrorEditLongUID() *foundation.TestCaseBodyUpsert {
	return &foundation.TestCaseBodyUpsert{
		TestName: "Ошибка при редактировании, длинный uid",
		Req: sgroupsv1.NamespaceReq_Upsert{
			Namespaces: []*sgroupsv1.Namespace{
				{
					Metadata: &common.Metadata{
						Name:        "namespace-3",
						Uid:         "321e83ed-7dfa-40b2-bff5-f8614294747c4",
						Labels:      map[string]string{"edit": "err"},
						Annotations: map[string]string{"edit": "err"},
					},
					Spec: &sgroupsv1.Namespace_Spec{
						Comment:     "for upsert tests",
						Description: "edit namespace",
						DisplayName: "edit err",
					},
				},
			},
		},
	}
}
