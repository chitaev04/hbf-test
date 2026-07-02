package bodies

import (
	sgroupsv1 "HBF-tests/pkg/api/sgroups/v1"
	"HBF-tests/tests/common/foundation"

	common "github.com/PRO-Robotech/sgroups-proto/pkg/api/common"
)

// ri Собирает ResourceIdentifier
func ri(name, namespace string) *common.ResourceIdentifier {
	return &common.ResourceIdentifier{Name: name, Namespace: namespace}
}

// fs Собирает FieldSelector для service-binding
func fs(name, namespace string, ag, svc *common.ResourceIdentifier) *sgroupsv1.ServiceBindingReq_Selectors_FieldSelector {
	return &sgroupsv1.ServiceBindingReq_Selectors_FieldSelector{Name: name, Namespace: namespace, AddressGroup: ag, Service: svc}
}

// sel Собирает Selectors из FieldSelector и/или LabelSelector
func sel(field *sgroupsv1.ServiceBindingReq_Selectors_FieldSelector, labels map[string]string) *sgroupsv1.ServiceBindingReq_Selectors {
	return &sgroupsv1.ServiceBindingReq_Selectors{FieldSelector: field, LabelSelector: labels}
}

// ls Собирает Selectors только с LabelSelector
func ls(labels map[string]string) *sgroupsv1.ServiceBindingReq_Selectors {
	return &sgroupsv1.ServiceBindingReq_Selectors{LabelSelector: labels}
}

func listSBBody(testName string, selectors ...*sgroupsv1.ServiceBindingReq_Selectors) *foundation.TestCaseBodySBList {
	return &foundation.TestCaseBodySBList{
		TestName: testName,
		Req:      sgroupsv1.ServiceBindingReq_List{Selectors: selectors},
	}
}

// ListAll Поиск всех service-binding
func ListAll() *foundation.TestCaseBodySBList {
	return &foundation.TestCaseBodySBList{TestName: "Поиск всех service-binding", Req: sgroupsv1.ServiceBindingReq_List{}}
}

// ListByName Поиск service-binding по name
func ListByName() *foundation.TestCaseBodySBList {
	return listSBBody("Поиск service-binding по name", sel(fs("sb-0", "", nil, nil), nil))
}

// ListByNonExistentName Поиск service-binding по несуществующему name
func ListByNonExistentName() *foundation.TestCaseBodySBList {
	return listSBBody("Поиск service-binding по несуществующему name", sel(fs("sb-000", "", nil, nil), nil))
}

// ListByExistingAndNonExistentName Поиск service-binding по существующему+несуществующему name
func ListByExistingAndNonExistentName() *foundation.TestCaseBodySBList {
	return listSBBody("Поиск service-binding по существующему+несуществующему name",
		sel(fs("sb-000", "", nil, nil), nil), sel(fs("sb-1", "", nil, nil), nil))
}

// ListByTwoNames Поиск 2 service-binding по names
func ListByTwoNames() *foundation.TestCaseBodySBList {
	return listSBBody("Поиск 2 service-binding по names", sel(fs("sb-0", "", nil, nil), nil), sel(fs("sb-1", "", nil, nil), nil))
}

// ListByNamespace Поиск service-binding по namespace
func ListByNamespace() *foundation.TestCaseBodySBList {
	return listSBBody("Поиск service-binding по namespace", sel(fs("", "namespace-0", nil, nil), nil))
}

// ListByNonExistentNamespace Поиск service-binding по несуществующему namespace
func ListByNonExistentNamespace() *foundation.TestCaseBodySBList {
	return listSBBody("Поиск service-binding по несуществующему namespace", sel(fs("", "namespace-2222", nil, nil), nil))
}

// ListByExistingAndNonExistentNamespace Поиск service-binding по существующему+несуществующему namespace
func ListByExistingAndNonExistentNamespace() *foundation.TestCaseBodySBList {
	return listSBBody("Поиск service-binding по существующему+несуществующему namespace",
		sel(fs("", "namespace-2222", nil, nil), nil), sel(fs("", "namespace-0", nil, nil), nil))
}

// ListByTwoNamespaces Поиск 2 service-binding по namespaces
func ListByTwoNamespaces() *foundation.TestCaseBodySBList {
	return listSBBody("Поиск 2 service-binding по namespaces", sel(fs("", "namespace-0", nil, nil), nil), sel(fs("", "namespace-2", nil, nil), nil))
}

// ListByNameAndNamespace Поиск service-binding по name+namespace
func ListByNameAndNamespace() *foundation.TestCaseBodySBList {
	return listSBBody("Поиск service-binding по name+namespace", sel(fs("sb-0", "namespace-0", nil, nil), nil))
}

// ListByNameAndNamespaceNamespaceNotExist Поиск service-binding по name+namespace, namespace не существует
func ListByNameAndNamespaceNamespaceNotExist() *foundation.TestCaseBodySBList {
	return listSBBody("Поиск service-binding по name+namespace, namespace не существует", sel(fs("sb-0", "namespace-22", nil, nil), nil))
}

// ListByNameAndNamespaceNameNotExist Поиск service-binding по name+namespace, name не существует
func ListByNameAndNamespaceNameNotExist() *foundation.TestCaseBodySBList {
	return listSBBody("Поиск service-binding по name+namespace, name не существует", sel(fs("sb-00", "namespace-0", nil, nil), nil))
}

// ListByNameAndNamespaceNoneExist Поиск service-binding по name+namespace, ни одного из не существует
func ListByNameAndNamespaceNoneExist() *foundation.TestCaseBodySBList {
	return listSBBody("Поиск service-binding по name+namespace, ни одного из не существует", sel(fs("sb-33", "namespace-33", nil, nil), nil))
}

// ListByNameNamespaceLabels Поиск service-binding по name+namespace+labels
func ListByNameNamespaceLabels() *foundation.TestCaseBodySBList {
	return listSBBody("Поиск service-binding по name+namespace+labels",
		sel(fs("sb-2", "namespace-2", nil, nil), map[string]string{"search": "nameLabels"}))
}

// ListByNameNamespaceLabelsNameNotExist Поиск service-binding по name+namespace+labels, name не существует
func ListByNameNamespaceLabelsNameNotExist() *foundation.TestCaseBodySBList {
	return listSBBody("Поиск service-binding по name+namespace+labels, name не существует",
		sel(fs("sb-33", "namespace-2", nil, nil), map[string]string{"search": "nameLabels"}))
}

// ListByNameNamespaceLabelsNamespaceNotExist Поиск service-binding по name+namespace+labels, namespace не существует
func ListByNameNamespaceLabelsNamespaceNotExist() *foundation.TestCaseBodySBList {
	return listSBBody("Поиск service-binding по name+namespace+labels, namespace не существует",
		sel(fs("sb-2", "namespace-22", nil, nil), map[string]string{"search": "nameLabels"}))
}

// ListByNameNamespaceLabelsLabelsNotExist Поиск service-binding по name+namespace+labels, labels не существует
func ListByNameNamespaceLabelsLabelsNotExist() *foundation.TestCaseBodySBList {
	return listSBBody("Поиск service-binding по name+namespace+labels, labels не существует",
		sel(fs("sb-2", "namespace-2", nil, nil), map[string]string{"not": "exist"}))
}

// ListByNameNamespaceLabelsNoneExist Поиск service-binding по name+namespace+labels, ни одного не существует
func ListByNameNamespaceLabelsNoneExist() *foundation.TestCaseBodySBList {
	return listSBBody("Поиск service-binding по name+namespace+labels, ни одного не существует",
		sel(fs("sb-33", "namespace-22", nil, nil), map[string]string{"not": "exist"}))
}

// ListByNamespaceLabels Поиск service-binding по namespace+labels
func ListByNamespaceLabels() *foundation.TestCaseBodySBList {
	return listSBBody("Поиск service-binding по namespace+labels", sel(fs("", "namespace-2", nil, nil), map[string]string{"search": "nameLabels"}))
}

// ListByNamespaceLabelsNamespaceNotExist Поиск service-binding по namespace+labels, namespace не существует
func ListByNamespaceLabelsNamespaceNotExist() *foundation.TestCaseBodySBList {
	return listSBBody("Поиск service-binding по namespace+labels, namespace не существует",
		sel(fs("", "namespace-22", nil, nil), map[string]string{"search": "nameLabels"}))
}

// ListByNamespaceLabelsLabelsNotExist Поиск service-binding по namespace+labels, labels не существует
func ListByNamespaceLabelsLabelsNotExist() *foundation.TestCaseBodySBList {
	return listSBBody("Поиск service-binding по namespace+labels, labels не существует",
		sel(fs("", "namespace-2", nil, nil), map[string]string{"not": "exist"}))
}

// ListByNamespaceLabelsNoneExist Поиск service-binding по namespace+labels, ни одного из не существует
func ListByNamespaceLabelsNoneExist() *foundation.TestCaseBodySBList {
	return listSBBody("Поиск service-binding по namespace+labels, ни одного из не существует",
		sel(fs("", "namespace-22", nil, nil), map[string]string{"not": "exist"}))
}

// ListByNameLabels Поиск service-binding по name+labels
func ListByNameLabels() *foundation.TestCaseBodySBList {
	return listSBBody("Поиск service-binding по name+labels", sel(fs("sb-2", "", nil, nil), map[string]string{"search": "nameLabels"}))
}

// ListByNameLabelsNameNotExist Поиск service-binding по name+labels, name не существует
func ListByNameLabelsNameNotExist() *foundation.TestCaseBodySBList {
	return listSBBody("Поиск service-binding по name+labels, name не существует",
		sel(fs("sb-22", "", nil, nil), map[string]string{"search": "nameLabels"}))
}

// ListByNameLabelsLabelsNotExist Поиск service-binding по name+labels, labels не существует
func ListByNameLabelsLabelsNotExist() *foundation.TestCaseBodySBList {
	return listSBBody("Поиск service-binding по name+labels, labels не существует",
		sel(fs("sb-2", "", nil, nil), map[string]string{"not": "exist"}))
}

// ListByNameLabelsNoneExist Поиск service-binding по name+labels, ни одного из не существует
func ListByNameLabelsNoneExist() *foundation.TestCaseBodySBList {
	return listSBBody("Поиск service-binding по name+labels, ни одного из не существует",
		sel(fs("sb-22", "", nil, nil), map[string]string{"not": "exist"}))
}

// ListByLabels Поиск service-binding по labels
func ListByLabels() *foundation.TestCaseBodySBList {
	return listSBBody("Поиск service-binding по labels", ls(map[string]string{"labels": "search"}))
}

// ListByTwoLabels Поиск 2 service-binding по labels
func ListByTwoLabels() *foundation.TestCaseBodySBList {
	return listSBBody("Поиск 2 service-binding по labels", ls(map[string]string{"labels": "search"}), ls(map[string]string{"labels": "ns"}))
}

// ListByNonExistentLabels Поиск service-binding по несуществующим labels
func ListByNonExistentLabels() *foundation.TestCaseBodySBList {
	return listSBBody("Поиск service-binding по несуществующим labels", ls(map[string]string{"not": "exist"}))
}

// ListByExistingAndNonExistentLabels Поиск service-binding по сущ+несущ labels
func ListByExistingAndNonExistentLabels() *foundation.TestCaseBodySBList {
	return listSBBody("Поиск service-binding по сущ+несущ labels", ls(map[string]string{"not": "exist"}), ls(map[string]string{"labels": "search"}))
}

// ListByExistingAG Поиск service-binding по существующей AG
func ListByExistingAG() *foundation.TestCaseBodySBList {
	return listSBBody("Поиск service-binding по существующей AG", sel(fs("", "", ri("ag-0", "namespace-0"), nil), nil))
}

// ListByExistingSVC Поиск service-binding по существующему svc
func ListByExistingSVC() *foundation.TestCaseBodySBList {
	return listSBBody("Поиск service-binding по существующему svc", sel(fs("", "", nil, ri("svc-0", "namespace-0")), nil))
}

// ListByExistingAndNonExistentAG Поиск service-binding по сущ+несущ AG
func ListByExistingAndNonExistentAG() *foundation.TestCaseBodySBList {
	return listSBBody("Поиск service-binding по сущ+несущ AG",
		sel(fs("", "", ri("ag-0", "namespace-0"), nil), nil), sel(fs("", "", ri("ag-11", "namespace-1"), nil), nil))
}

// ListByExistingAndNonExistentSVC Поиск service-binding по сущ+несущ svc
func ListByExistingAndNonExistentSVC() *foundation.TestCaseBodySBList {
	return listSBBody("Поиск service-binding по сущ+несущ svc",
		sel(fs("", "", nil, ri("svc-0", "namespace-0")), nil), sel(fs("", "", nil, ri("svc-11", "namespace-1")), nil))
}

// ListByNonExistentAG Поиск service-binding по несущ AG
func ListByNonExistentAG() *foundation.TestCaseBodySBList {
	return listSBBody("Поиск service-binding по несущ AG", sel(fs("", "", ri("ag-11", "namespace-1"), nil), nil))
}

// ListByNonExistentSVC Поиск service-binding по несущ svc
func ListByNonExistentSVC() *foundation.TestCaseBodySBList {
	return listSBBody("Поиск service-binding по несущ svc", sel(fs("", "", nil, ri("svc-11", "namespace-1")), nil))
}

// ListByTwoAG Поиск 2 service-binding по AG
func ListByTwoAG() *foundation.TestCaseBodySBList {
	return listSBBody("Поиск 2 service-binding по AG",
		sel(fs("", "", ri("ag-0", "namespace-0"), nil), nil), sel(fs("", "", ri("ag-1", "namespace-1"), nil), nil))
}

// ListByTwoSVC Поиск 2 service-binding по svc
func ListByTwoSVC() *foundation.TestCaseBodySBList {
	return listSBBody("Поиск 2 service-binding по svc",
		sel(fs("", "", nil, ri("svc-0", "namespace-0")), nil), sel(fs("", "", nil, ri("svc-1", "namespace-1")), nil))
}

// ListByAGAndName Поиск service-binding по AG+name
func ListByAGAndName() *foundation.TestCaseBodySBList {
	return listSBBody("Поиск service-binding по AG+name", sel(fs("sb-0", "", ri("ag-0", "namespace-0"), nil), nil))
}

// ListBySVCAndName Поиск service-binding по svc+name
func ListBySVCAndName() *foundation.TestCaseBodySBList {
	return listSBBody("Поиск service-binding по svc+name", sel(fs("sb-0", "", nil, ri("svc-0", "namespace-0")), nil))
}

// ListByAGNameNamespace Поиск service-binding по AG+name+ns
func ListByAGNameNamespace() *foundation.TestCaseBodySBList {
	return listSBBody("Поиск service-binding по AG+name+ns", sel(fs("sb-0", "namespace-0", ri("ag-0", "namespace-0"), nil), nil))
}

// ListBySVCNameNamespace Поиск service-binding по svc+name+ns
func ListBySVCNameNamespace() *foundation.TestCaseBodySBList {
	return listSBBody("Поиск service-binding по svc+name+ns", sel(fs("sb-0", "", nil, ri("svc-0", "namespace-0")), nil))
}

// ListByAGNameNamespaceLabels Поиск service-binding по AG+name+ns+labels
func ListByAGNameNamespaceLabels() *foundation.TestCaseBodySBList {
	return listSBBody("Поиск service-binding по AG+name+ns+labels",
		sel(fs("sb-0", "", ri("ag-0", "namespace-0"), nil), map[string]string{"labels": "search"}))
}

// ListBySVCNameNamespaceLabels Поиск service-binding по svc+name+ns+labels
func ListBySVCNameNamespaceLabels() *foundation.TestCaseBodySBList {
	return listSBBody("Поиск service-binding по svc+name+ns+labels",
		sel(fs("sb-0", "", nil, ri("svc-0", "namespace-0")), map[string]string{"labels": "search"}))
}

// ListByAGNameAGNotExist Поиск service-binding по AG+name, AG не существует
func ListByAGNameAGNotExist() *foundation.TestCaseBodySBList {
	return listSBBody("Поиск service-binding по AG+name, AG не существует", sel(fs("sb-0", "", ri("ag-11", "namespace-0"), nil), nil))
}

// ListBySVCNameSVCNotExist Поиск service-binding по host+name, host не существует
func ListBySVCNameSVCNotExist() *foundation.TestCaseBodySBList {
	return listSBBody("Поиск service-binding по host+name, host не существует", sel(fs("sb-0", "", nil, ri("svc-11", "namespace-0")), nil))
}

// ListByAGNameNameNotExist Поиск service-binding по AG+name, name не существует
func ListByAGNameNameNotExist() *foundation.TestCaseBodySBList {
	return listSBBody("Поиск service-binding по AG+name, name не существует", sel(fs("sb-00", "", ri("ag-0", "namespace-0"), nil), nil))
}

// ListBySVCNameNameNotExist Поиск service-binding по svc+name, name не существует
func ListBySVCNameNameNotExist() *foundation.TestCaseBodySBList {
	return listSBBody("Поиск service-binding по svc+name, name не существует", sel(fs("sb-00", "", nil, ri("svc-0", "namespace-0")), nil))
}

// ListByAGNameNoneExist Поиск service-binding по AG+name, ни одного из не существует
func ListByAGNameNoneExist() *foundation.TestCaseBodySBList {
	return listSBBody("Поиск service-binding по AG+name, ни одного из не существует", sel(fs("sb-00", "", ri("ag-00", "namespace-0"), nil), nil))
}

// ListByAGAndSVC Поиск service-binding по AG+SVC
func ListByAGAndSVC() *foundation.TestCaseBodySBList {
	return listSBBody("Поиск service-binding по AG+SVC", sel(fs("", "", ri("ag-0", "namespace-0"), ri("svc-0", "namespace-0")), nil))
}

// ListByAllParams Поиск service-binding по всем параметрам
func ListByAllParams() *foundation.TestCaseBodySBList {
	return listSBBody("Поиск service-binding по всем параметрам",
		sel(fs("sb-0", "namespace-0", ri("ag-0", "namespace-0"), ri("svc-0", "namespace-0")), map[string]string{"labels": "search"}))
}
