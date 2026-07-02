package bodies

import (
	sgroupsv1 "HBF-tests/pkg/api/sgroups/v1"
	"HBF-tests/tests/common/foundation"

	common "github.com/PRO-Robotech/sgroups-proto/pkg/api/common"
)

func deleteHBByMetadata(testName string, metadata *common.MetadataScope) *foundation.TestCaseBodyHBDelete {
	return &foundation.TestCaseBodyHBDelete{
		TestName: testName,
		Req: sgroupsv1.HostBindingReq_Delete{
			HostBindings: []*sgroupsv1.HostBindingReq_Delete_HostBinding{
				{Metadata: metadata},
			},
		},
	}
}

// DeleteByNameNamespace Удаление по name+namespace
func DeleteByNameNamespace() *foundation.TestCaseBodyHBDelete {
	return deleteHBByMetadata("Удаление по name+namespace", &common.MetadataScope{Name: "hb-4", Namespace: "namespace-1"})
}

// DeleteByUID Удаление по uid
func DeleteByUID() *foundation.TestCaseBodyHBDelete {
	return deleteHBByMetadata("Удаление по uid", &common.MetadataScope{Uid: "13155138-92ec-49e1-8a9e-2f1faac09dcd"})
}

// DeleteErrorNameEmpty Ошибка при удалении по name+namespace, name пуст
func DeleteErrorNameEmpty() *foundation.TestCaseBodyHBDelete {
	return deleteHBByMetadata("Ошибка при удалении по name+namespace, name пуст", &common.MetadataScope{Name: "", Namespace: "namespace-4"})
}

// DeleteErrorNameLeadingSpace Ошибка при удалении по name+namespace, пробел в начале name
func DeleteErrorNameLeadingSpace() *foundation.TestCaseBodyHBDelete {
	return deleteHBByMetadata("Ошибка при удалении по name+namespace, пробел в начале name", &common.MetadataScope{Name: " hb-4", Namespace: "namespace-4"})
}

// DeleteErrorNameTrailingSpace Ошибка при удалении по name+namespace, пробел в конце name
func DeleteErrorNameTrailingSpace() *foundation.TestCaseBodyHBDelete {
	return deleteHBByMetadata("Ошибка при удалении по name+namespace, пробел в конце name", &common.MetadataScope{Name: "hb-4 ", Namespace: "namespace-4"})
}

// DeleteErrorNameWithSpaces Ошибка при удалении по name+namespace, name с пробелами
func DeleteErrorNameWithSpaces() *foundation.TestCaseBodyHBDelete {
	return deleteHBByMetadata("Ошибка при удалении по name+namespace, name с пробелами", &common.MetadataScope{Name: "hb - 4", Namespace: "namespace-4"})
}

// DeleteErrorNameNotExist Ошибка при удалении по name+namespace, name не существует
func DeleteErrorNameNotExist() *foundation.TestCaseBodyHBDelete {
	return deleteHBByMetadata("Ошибка при удалении по name+namespace, name не существует", &common.MetadataScope{Name: "hb-00", Namespace: "namespace-4"})
}

// DeleteErrorNamespaceEmpty Ошибка при удалении по name+namespace, namespace пуст
func DeleteErrorNamespaceEmpty() *foundation.TestCaseBodyHBDelete {
	return deleteHBByMetadata("Ошибка при удалении по name+namespace, namespace пуст", &common.MetadataScope{Name: "hb-00", Namespace: ""})
}

// DeleteErrorNamespaceLeadingSpace Ошибка при удалении по name+namespace, пробел в начале имени namespace
func DeleteErrorNamespaceLeadingSpace() *foundation.TestCaseBodyHBDelete {
	return deleteHBByMetadata("Ошибка при удалении по name+namespace, пробел в начале имени namespace", &common.MetadataScope{Name: "hb-0", Namespace: " namespace-0"})
}

// DeleteErrorNamespaceTrailingSpace Ошибка при удалении по name+namespace, пробел в конце имени namespace
func DeleteErrorNamespaceTrailingSpace() *foundation.TestCaseBodyHBDelete {
	return deleteHBByMetadata("Ошибка при удалении по name+namespace, пробел в конце имени namespace", &common.MetadataScope{Name: "hb-0", Namespace: "namespace-0 "})
}

// DeleteErrorNamespaceWithSpaces Ошибка при удалении по name+namespace, имя namespace с пробелами
func DeleteErrorNamespaceWithSpaces() *foundation.TestCaseBodyHBDelete {
	return deleteHBByMetadata("Ошибка при удалении по name+namespace, имя namespace с пробелами", &common.MetadataScope{Name: "hb-0", Namespace: "namespace - 0"})
}

// DeleteErrorNamespaceNotExist Ошибка при удалении по name+namespace, namespace не существует
func DeleteErrorNamespaceNotExist() *foundation.TestCaseBodyHBDelete {
	return deleteHBByMetadata("Ошибка при удалении по name+namespace, namespace не существует", &common.MetadataScope{Name: "hb-0", Namespace: "not-exst"})
}

// DeleteErrorNamespaceMissing Ошибка при удалении по name+namespace, namespace отсутствует как параметр
func DeleteErrorNamespaceMissing() *foundation.TestCaseBodyHBDelete {
	return deleteHBByMetadata("Ошибка при удалении по name+namespace, namespace отсутствует как параметр", &common.MetadataScope{Name: "hb-0"})
}

// DeleteErrorNameNotExistInNamespace Ошибка при удалении по name+namespace, name не существует в указанном namespace
func DeleteErrorNameNotExistInNamespace() *foundation.TestCaseBodyHBDelete {
	return deleteHBByMetadata("Ошибка при удалении по name+namespace, name не существует в указанном namespace", &common.MetadataScope{Name: "hb-0", Namespace: "namespace-4"})
}

// DeleteErrorNameMissing Ошибка при удалении по name+namespace, name отсутствует как параметр
func DeleteErrorNameMissing() *foundation.TestCaseBodyHBDelete {
	return deleteHBByMetadata("Ошибка при удалении по name+namespace, name отсутствует как параметр", &common.MetadataScope{Namespace: "namespace-4"})
}

// DeleteErrorUIDEmpty Ошибка при удалении по uid, пустой uid
func DeleteErrorUIDEmpty() *foundation.TestCaseBodyHBDelete {
	return deleteHBByMetadata("Ошибка при удалении по uid, пустой uid", &common.MetadataScope{Uid: ""})
}

// DeleteErrorUIDNotExist Ошибка при удалении по uid, uid не существует
func DeleteErrorUIDNotExist() *foundation.TestCaseBodyHBDelete {
	return deleteHBByMetadata("Ошибка при удалении по uid, uid не существует", &common.MetadataScope{Uid: "368240f3-4866-4cbf-ad3f-6c55b48d4b92"})
}

// DeleteErrorUIDShort Ошибка при удалении по uid, короткий uid
func DeleteErrorUIDShort() *foundation.TestCaseBodyHBDelete {
	return deleteHBByMetadata("Ошибка при удалении по uid, короткий uid", &common.MetadataScope{Uid: "368240f3-4866-4cbf-ad3f-6c55b48d92"})
}

// DeleteErrorUIDLong Ошибка при удалении по uid, длинный uid
func DeleteErrorUIDLong() *foundation.TestCaseBodyHBDelete {
	return deleteHBByMetadata("Ошибка при удалении по uid, длинный uid", &common.MetadataScope{Uid: "368240f3-4866-4cbf-ad3f-6c55b48djkl92"})
}

// DeleteErrorMetadataEmpty Ошибка при удалении, пустая метадата
func DeleteErrorMetadataEmpty() *foundation.TestCaseBodyHBDelete {
	return deleteHBByMetadata("Ошибка при удалении, пустая метадата", &common.MetadataScope{})
}

// DeleteErrorMetadataMissing Ошибка при удалении, метадата отсутствует как параметр
func DeleteErrorMetadataMissing() *foundation.TestCaseBodyHBDelete {
	return &foundation.TestCaseBodyHBDelete{
		TestName: "Ошибка при удалении, метадата отсутствует как параметр",
		Req: sgroupsv1.HostBindingReq_Delete{
			HostBindings: []*sgroupsv1.HostBindingReq_Delete_HostBinding{{}},
		},
	}
}
