package bodies

import (
	sgroupsv1 "HBF-tests/pkg/api/sgroups/v1"
	"HBF-tests/tests/common/foundation"
)

// deleteByMetadata вспомогательный конструктор тела запроса на удаление namespace по метаданным
func deleteByMetadata(testName string, metadata *sgroupsv1.NamespaceReq_Delete_MetadataScope) *foundation.TestCaseBodyDelete {
	return &foundation.TestCaseBodyDelete{
		TestName: testName,
		Req: sgroupsv1.NamespaceReq_Delete{
			Namespaces: []*sgroupsv1.NamespaceReq_Delete_Namespace{
				{Metadata: metadata},
			},
		},
	}
}

// DeleteByName Удаление по name
func DeleteByName() *foundation.TestCaseBodyDelete {
	return deleteByMetadata("Удаление по name", &sgroupsv1.NamespaceReq_Delete_MetadataScope{Name: "namespace-4"})
}

// DeleteByUID Удаление по uid
func DeleteByUID() *foundation.TestCaseBodyDelete {
	return deleteByMetadata("Удаление по uid", &sgroupsv1.NamespaceReq_Delete_MetadataScope{Uid: "1e7c1f61-c021-4e53-8d53-79ee323264b9"})
}

// DeleteErrorUIDNotExist Ошибка при удалении по uid, namespace с таким uid не существует
func DeleteErrorUIDNotExist() *foundation.TestCaseBodyDelete {
	return deleteByMetadata("Ошибка при удалении по uid, namespace с таким uid не существует", &sgroupsv1.NamespaceReq_Delete_MetadataScope{Uid: "39fd4d11-b7c2-4eae-ab46-9ccfb5ea7a26"})
}

// DeleteErrorUIDEmpty Ошибка при удалении по uid, uid пуст
func DeleteErrorUIDEmpty() *foundation.TestCaseBodyDelete {
	return deleteByMetadata("Ошибка при удалении по uid, uid пуст", &sgroupsv1.NamespaceReq_Delete_MetadataScope{Uid: ""})
}

// DeleteErrorUIDShort Ошибка при удалении по uid, uid короткий
func DeleteErrorUIDShort() *foundation.TestCaseBodyDelete {
	return deleteByMetadata("Ошибка при удалении по uid, uid короткий", &sgroupsv1.NamespaceReq_Delete_MetadataScope{Uid: "5887ddd3-2449-4905-8c45-20c808fcbd3"})
}

// DeleteErrorUIDLong Ошибка при удалении по uid, uid длинный
func DeleteErrorUIDLong() *foundation.TestCaseBodyDelete {
	return deleteByMetadata("Ошибка при удалении по uid, uid длинный", &sgroupsv1.NamespaceReq_Delete_MetadataScope{Uid: "5887ddd3-2449-4905-8c45-20c808fcbd563"})
}

// DeleteErrorNameNotExist Ошибка при удалении по name, namespace с таким name не существует
func DeleteErrorNameNotExist() *foundation.TestCaseBodyDelete {
	return deleteByMetadata("Ошибка при удалении по name, namespace с таким name не существует", &sgroupsv1.NamespaceReq_Delete_MetadataScope{Name: "err-name"})
}

// DeleteErrorNameEmpty Ошибка при удалении по name, name пуст
func DeleteErrorNameEmpty() *foundation.TestCaseBodyDelete {
	return deleteByMetadata("Ошибка при удалении по name, name пуст", &sgroupsv1.NamespaceReq_Delete_MetadataScope{Name: ""})
}

// DeleteErrorNameLeadingSpace Ошибка при удалении по name, пробел в начале name
func DeleteErrorNameLeadingSpace() *foundation.TestCaseBodyDelete {
	return deleteByMetadata("Ошибка при удалении по name, пробел в начале name", &sgroupsv1.NamespaceReq_Delete_MetadataScope{Name: " err-name"})
}

// DeleteErrorNameTrailingSpace Ошибка при удалении по name, пробел в конце name
func DeleteErrorNameTrailingSpace() *foundation.TestCaseBodyDelete {
	return deleteByMetadata("Ошибка при удалении по name, пробел в конце name", &sgroupsv1.NamespaceReq_Delete_MetadataScope{Name: "err-name "})
}

// DeleteErrorNameWithSpaces Ошибка при удалении по name, name с пробелами
func DeleteErrorNameWithSpaces() *foundation.TestCaseBodyDelete {
	return deleteByMetadata("Ошибка при удалении по name, name с пробелами", &sgroupsv1.NamespaceReq_Delete_MetadataScope{Name: "err - name"})
}

// DeleteErrorMetadataEmpty Ошибка при удалении, metadata пустая
func DeleteErrorMetadataEmpty() *foundation.TestCaseBodyDelete {
	return deleteByMetadata("Ошибка при удалении, metadata пустая", &sgroupsv1.NamespaceReq_Delete_MetadataScope{})
}

// DeleteErrorMetadataMissing Ошибка при удалении, metadata отсутствует как параметр
func DeleteErrorMetadataMissing() *foundation.TestCaseBodyDelete {
	return &foundation.TestCaseBodyDelete{
		TestName: "Ошибка при удалении, metadata отсутствует как параметр",
		Req: sgroupsv1.NamespaceReq_Delete{
			Namespaces: []*sgroupsv1.NamespaceReq_Delete_Namespace{
				{},
			},
		},
	}
}
