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

// fs Собирает FieldSelector для host-binding
func fs(name, namespace string, ag, host *common.ResourceIdentifier) *sgroupsv1.HostBindingReq_Selectors_FieldSelector {
	return &sgroupsv1.HostBindingReq_Selectors_FieldSelector{Name: name, Namespace: namespace, AddressGroup: ag, Host: host}
}

// sel Собирает Selectors из FieldSelector и/или LabelSelector
func sel(field *sgroupsv1.HostBindingReq_Selectors_FieldSelector, labels map[string]string) *sgroupsv1.HostBindingReq_Selectors {
	return &sgroupsv1.HostBindingReq_Selectors{FieldSelector: field, LabelSelector: labels}
}

// ls Собирает Selectors только с LabelSelector
func ls(labels map[string]string) *sgroupsv1.HostBindingReq_Selectors {
	return &sgroupsv1.HostBindingReq_Selectors{LabelSelector: labels}
}

func listHBBody(testName string, selectors ...*sgroupsv1.HostBindingReq_Selectors) *foundation.TestCaseBodyHBList {
	return &foundation.TestCaseBodyHBList{
		TestName: testName,
		Req:      sgroupsv1.HostBindingReq_List{Selectors: selectors},
	}
}

// ListAll Поиск всех host-binding
func ListAll() *foundation.TestCaseBodyHBList {
	return &foundation.TestCaseBodyHBList{TestName: "Поиск всех host-binding", Req: sgroupsv1.HostBindingReq_List{}}
}

// ListByName Поиск host-binding по name
func ListByName() *foundation.TestCaseBodyHBList {
	return listHBBody("Поиск host-binding по name", sel(fs("hb-0", "", nil, nil), nil))
}

// ListByNonExistentName Поиск host-binding по несуществующему name
func ListByNonExistentName() *foundation.TestCaseBodyHBList {
	return listHBBody("Поиск host-binding по несуществующему name", sel(fs("host-000", "", nil, nil), nil))
}

// ListByExistingAndNonExistentName Поиск host-binding по существующему+несуществующему name
func ListByExistingAndNonExistentName() *foundation.TestCaseBodyHBList {
	return listHBBody("Поиск host-binding по существующему+несуществующему name",
		sel(fs("hb-000", "", nil, nil), nil), sel(fs("hb-1", "", nil, nil), nil))
}

// ListByTwoNames Поиск 2 host-binding по names
func ListByTwoNames() *foundation.TestCaseBodyHBList {
	return listHBBody("Поиск 2 host-binding по names", sel(fs("hb-0", "", nil, nil), nil), sel(fs("hb-1", "", nil, nil), nil))
}

// ListByNamespace Поиск host-binding по namespace
func ListByNamespace() *foundation.TestCaseBodyHBList {
	return listHBBody("Поиск host-binding по namespace", sel(fs("", "namespace-0", nil, nil), nil))
}

// ListByNonExistentNamespace Поиск host-binding по несуществующему namespace
func ListByNonExistentNamespace() *foundation.TestCaseBodyHBList {
	return listHBBody("Поиск host-binding по несуществующему namespace", sel(fs("", "namespace-2222", nil, nil), nil))
}

// ListByExistingAndNonExistentNamespace Поиск host-binding по существующему+несуществующему namespace
func ListByExistingAndNonExistentNamespace() *foundation.TestCaseBodyHBList {
	return listHBBody("Поиск host-binding по существующему+несуществующему namespace",
		sel(fs("", "namespace-2222", nil, nil), nil), sel(fs("", "namespace-0", nil, nil), nil))
}

// ListByTwoNamespaces Поиск 2 host-binding по namespaces
func ListByTwoNamespaces() *foundation.TestCaseBodyHBList {
	return listHBBody("Поиск 2 host-binding по namespaces", sel(fs("", "namespace-0", nil, nil), nil), sel(fs("", "namespace-2", nil, nil), nil))
}

// ListByNameAndNamespace Поиск host-binding по name+namespace
func ListByNameAndNamespace() *foundation.TestCaseBodyHBList {
	return listHBBody("Поиск host-binding по name+namespace", sel(fs("hb-0", "namespace-0", nil, nil), nil))
}

// ListByNameAndNamespaceNamespaceNotExist Поиск host-binding по name+namespace, namespace не существует
func ListByNameAndNamespaceNamespaceNotExist() *foundation.TestCaseBodyHBList {
	return listHBBody("Поиск host-binding по name+namespace, namespace не существует", sel(fs("ag-0", "namespace-22", nil, nil), nil))
}

// ListByNameAndNamespaceNameNotExist Поиск host-binding по name+namespace, name не существует
func ListByNameAndNamespaceNameNotExist() *foundation.TestCaseBodyHBList {
	return listHBBody("Поиск host-binding по name+namespace, name не существует", sel(fs("host-33", "namespace-0", nil, nil), nil))
}

// ListByNameAndNamespaceNoneExist Поиск host-binding по name+namespace, ни одного из не существует
func ListByNameAndNamespaceNoneExist() *foundation.TestCaseBodyHBList {
	return listHBBody("Поиск host-binding по name+namespace, ни одного из не существует", sel(fs("host-33", "namespace-22", nil, nil), nil))
}

// ListByNameNamespaceLabels Поиск host-binding по name+namespace+labels
func ListByNameNamespaceLabels() *foundation.TestCaseBodyHBList {
	return listHBBody("Поиск host-binding по name+namespace+labels",
		sel(fs("hb-2", "namespace-2", nil, nil), map[string]string{"search": "nameLabels"}))
}

// ListByNameNamespaceLabelsNameNotExist Поиск host-binding по name+namespace+labels, name не существует
func ListByNameNamespaceLabelsNameNotExist() *foundation.TestCaseBodyHBList {
	return listHBBody("Поиск host-binding по name+namespace+labels, name не существует",
		sel(fs("hb-33", "namespace-2", nil, nil), map[string]string{"search": "nameLabels"}))
}

// ListByNameNamespaceLabelsNamespaceNotExist Поиск host-binding по name+namespace+labels, namespace не существует
func ListByNameNamespaceLabelsNamespaceNotExist() *foundation.TestCaseBodyHBList {
	return listHBBody("Поиск host-binding по name+namespace+labels, namespace не существует",
		sel(fs("hb-2", "namespace-22", nil, nil), map[string]string{"search": "nameLabels"}))
}

// ListByNameNamespaceLabelsLabelsNotExist Поиск host-binding по name+namespace+labels, labels не существует
func ListByNameNamespaceLabelsLabelsNotExist() *foundation.TestCaseBodyHBList {
	return listHBBody("Поиск host-binding по name+namespace+labels, labels не существует",
		sel(fs("hb-2", "namespace-2", nil, nil), map[string]string{"not": "exist"}))
}

// ListByNameNamespaceLabelsNoneExist Поиск host-binding по name+namespace+labels, ни одного не существует
func ListByNameNamespaceLabelsNoneExist() *foundation.TestCaseBodyHBList {
	return listHBBody("Поиск host-binding по name+namespace+labels, ни одного не существует",
		sel(fs("hb-33", "namespace-22", nil, nil), map[string]string{"not": "exist"}))
}

// ListByNamespaceLabels Поиск host-binding по namespace+labels
func ListByNamespaceLabels() *foundation.TestCaseBodyHBList {
	return listHBBody("Поиск host-binding по namespace+labels", sel(fs("", "namespace-2", nil, nil), map[string]string{"search": "nameLabels"}))
}

// ListByNamespaceLabelsNamespaceNotExist Поиск host-binding по namespace+labels, namespace не существует
func ListByNamespaceLabelsNamespaceNotExist() *foundation.TestCaseBodyHBList {
	return listHBBody("Поиск host-binding по namespace+labels, namespace не существует",
		sel(fs("", "namespace-22", nil, nil), map[string]string{"search": "nameLabels"}))
}

// ListByNamespaceLabelsLabelsNotExist Поиск host-binding по namespace+labels, labels не существует
func ListByNamespaceLabelsLabelsNotExist() *foundation.TestCaseBodyHBList {
	return listHBBody("Поиск host-binding по namespace+labels, labels не существует",
		sel(fs("", "namespace-2", nil, nil), map[string]string{"not": "exist"}))
}

// ListByNamespaceLabelsNoneExist Поиск host-binding по namespace+labels, ни одного из не существует
func ListByNamespaceLabelsNoneExist() *foundation.TestCaseBodyHBList {
	return listHBBody("Поиск host-binding по namespace+labels, ни одного из не существует",
		sel(fs("", "namespace-22", nil, nil), map[string]string{"not": "exist"}))
}

// ListByNameLabels Поиск host-binding по name+labels
func ListByNameLabels() *foundation.TestCaseBodyHBList {
	return listHBBody("Поиск host-binding по name+labels", sel(fs("hb-2", "", nil, nil), map[string]string{"search": "nameLabels"}))
}

// ListByNameLabelsNameNotExist Поиск host-binding по name+labels, name не существует
func ListByNameLabelsNameNotExist() *foundation.TestCaseBodyHBList {
	return listHBBody("Поиск host-binding по name+labels, name не существует",
		sel(fs("host-22", "", nil, nil), map[string]string{"search": "nameLabels"}))
}

// ListByNameLabelsLabelsNotExist Поиск host-binding по name+labels, labels не существует
func ListByNameLabelsLabelsNotExist() *foundation.TestCaseBodyHBList {
	return listHBBody("Поиск host-binding по name+labels, labels не существует",
		sel(fs("host-2", "", nil, nil), map[string]string{"not": "exist"}))
}

// ListByNameLabelsNoneExist Поиск host-binding по name+labels, ни одного из не существует
func ListByNameLabelsNoneExist() *foundation.TestCaseBodyHBList {
	return listHBBody("Поиск host-binding по name+labels, ни одного из не существует",
		sel(fs("host-22", "", nil, nil), map[string]string{"not": "exist"}))
}

// ListByLabels Поиск host-binding по labels
func ListByLabels() *foundation.TestCaseBodyHBList {
	return listHBBody("Поиск host-binding по labels", ls(map[string]string{"labels": "search"}))
}

// ListByTwoLabels Поиск 2 host-binding по labels
func ListByTwoLabels() *foundation.TestCaseBodyHBList {
	return listHBBody("Поиск 2 host-binding по labels", ls(map[string]string{"labels": "search"}), ls(map[string]string{"labels": "ns"}))
}

// ListByNonExistentLabels Поиск host-binding по несуществующим labels
func ListByNonExistentLabels() *foundation.TestCaseBodyHBList {
	return listHBBody("Поиск host-binding по несуществующим labels", ls(map[string]string{"not": "exist"}))
}

// ListByExistingAndNonExistentLabels Поиск host-binding по сущ+несущ labels
func ListByExistingAndNonExistentLabels() *foundation.TestCaseBodyHBList {
	return listHBBody("Поиск host-binding по сущ+несущ labels", ls(map[string]string{"not": "exist"}), ls(map[string]string{"labels": "search"}))
}

// ListByExistingAG Поиск host-binding по существующей AG
func ListByExistingAG() *foundation.TestCaseBodyHBList {
	return listHBBody("Поиск host-binding по существующей AG", sel(fs("", "", ri("ag-0", "namespace-0"), nil), nil))
}

// ListByExistingHost Поиск host-binding по существующему HOST
func ListByExistingHost() *foundation.TestCaseBodyHBList {
	return listHBBody("Поиск host-binding по существующему HOST", sel(fs("", "", nil, ri("host-0", "namespace-0")), nil))
}

// ListByExistingAndNonExistentAG Поиск host-binding по сущ+несущ AG
func ListByExistingAndNonExistentAG() *foundation.TestCaseBodyHBList {
	return listHBBody("Поиск host-binding по сущ+несущ AG",
		sel(fs("", "", ri("ag-0", "namespace-0"), nil), nil), sel(fs("", "", ri("ag-11", "namespace-1"), nil), nil))
}

// ListByExistingAndNonExistentHost Поиск host-binding по сущ+несущ HOST
func ListByExistingAndNonExistentHost() *foundation.TestCaseBodyHBList {
	return listHBBody("Поиск host-binding по сущ+несущ HOST",
		sel(fs("", "", nil, ri("host-0", "namespace-0")), nil), sel(fs("", "", nil, ri("host-11", "namespace-1")), nil))
}

// ListByNonExistentAG Поиск host-binding по несущ AG
func ListByNonExistentAG() *foundation.TestCaseBodyHBList {
	return listHBBody("Поиск host-binding по несущ AG", sel(fs("", "", ri("ag-11", "namespace-1"), nil), nil))
}

// ListByNonExistentHost Поиск host-binding по несущ HOST
func ListByNonExistentHost() *foundation.TestCaseBodyHBList {
	return listHBBody("Поиск host-binding по несущ HOST", sel(fs("", "", nil, ri("host-11", "namespace-1")), nil))
}

// ListByTwoAG Поиск 2 host-binding по AG
func ListByTwoAG() *foundation.TestCaseBodyHBList {
	return listHBBody("Поиск 2 host-binding по AG",
		sel(fs("", "", ri("ag-0", "namespace-0"), nil), nil), sel(fs("", "", ri("ag-1", "namespace-1"), nil), nil))
}

// ListByTwoHost Поиск 2 host-binding по HOST
func ListByTwoHost() *foundation.TestCaseBodyHBList {
	return listHBBody("Поиск 2 host-binding по HOST",
		sel(fs("", "", nil, ri("host-0", "namespace-0")), nil), sel(fs("", "", nil, ri("host-1", "namespace-1")), nil))
}

// ListByAGAndName Поиск host-binding по AG+name
func ListByAGAndName() *foundation.TestCaseBodyHBList {
	return listHBBody("Поиск host-binding по AG+name", sel(fs("hb-0", "", ri("ag-0", "namespace-0"), nil), nil))
}

// ListByHostAndName Поиск host-binding по host+name
func ListByHostAndName() *foundation.TestCaseBodyHBList {
	return listHBBody("Поиск host-binding по host+name", sel(fs("hb-0", "", nil, ri("host-0", "namespace-0")), nil))
}

// ListByAGNameNamespace Поиск host-binding по AG+name+ns
func ListByAGNameNamespace() *foundation.TestCaseBodyHBList {
	return listHBBody("Поиск host-binding по AG+name+ns", sel(fs("hb-0", "namespace-0", ri("ag-0", "namespace-0"), nil), nil))
}

// ListByHostNameNamespace Поиск host-binding по host+name+ns
func ListByHostNameNamespace() *foundation.TestCaseBodyHBList {
	return listHBBody("Поиск host-binding по host+name+ns", sel(fs("hb-0", "", nil, ri("host-0", "namespace-0")), nil))
}

// ListByAGNameNamespaceLabels Поиск host-binding по AG+name+ns+labels
func ListByAGNameNamespaceLabels() *foundation.TestCaseBodyHBList {
	return listHBBody("Поиск host-binding по AG+name+ns+labels",
		sel(fs("hb-0", "", ri("ag-0", "namespace-0"), nil), map[string]string{"labels": "search"}))
}

// ListByHostNameNamespaceLabels Поиск host-binding по host+name+ns+labels
func ListByHostNameNamespaceLabels() *foundation.TestCaseBodyHBList {
	return listHBBody("Поиск host-binding по host+name+ns+labels",
		sel(fs("hb-0", "", nil, ri("host-0", "namespace-0")), map[string]string{"labels": "search"}))
}

// ListByAGNameAGNotExist Поиск host-binding по AG+name, AG не существует
func ListByAGNameAGNotExist() *foundation.TestCaseBodyHBList {
	return listHBBody("Поиск host-binding по AG+name, AG не существует", sel(fs("hb-0", "", ri("ag-11", "namespace-0"), nil), nil))
}

// ListByHostNameHostNotExist Поиск host-binding по host+name, host не существует
func ListByHostNameHostNotExist() *foundation.TestCaseBodyHBList {
	return listHBBody("Поиск host-binding по host+name, host не существует", sel(fs("hb-0", "", nil, ri("host-11", "namespace-0")), nil))
}

// ListByAGNameNameNotExist Поиск host-binding по AG+name, name не существует
func ListByAGNameNameNotExist() *foundation.TestCaseBodyHBList {
	return listHBBody("Поиск host-binding по AG+name, name не существует", sel(fs("hb-00", "", ri("ag-0", "namespace-0"), nil), nil))
}

// ListByHostNameNameNotExist Поиск host-binding по host+name, name не существует
func ListByHostNameNameNotExist() *foundation.TestCaseBodyHBList {
	return listHBBody("Поиск host-binding по host+name, name не существует", sel(fs("hb-00", "", nil, ri("host-0", "namespace-0")), nil))
}

// ListByAGNameNoneExist Поиск host-binding по AG+name, ни одного из не существует
func ListByAGNameNoneExist() *foundation.TestCaseBodyHBList {
	return listHBBody("Поиск host-binding по AG+name, ни одного из не существует", sel(fs("hb-00", "", ri("ag-00", "namespace-0"), nil), nil))
}

// ListByAGAndHost Поиск host-binding по AG+HOST
func ListByAGAndHost() *foundation.TestCaseBodyHBList {
	return listHBBody("Поиск host-binding по AG+HOST", sel(fs("", "", ri("ag-0", "namespace-0"), ri("host-0", "namespace-0")), nil))
}

// ListByAllParams Поиск host-binding по всем параметрам
func ListByAllParams() *foundation.TestCaseBodyHBList {
	return listHBBody("Поиск host-binding по всем параметрам",
		sel(fs("hb-0", "namespace-0", ri("ag-0", "namespace-0"), ri("host-0", "namespace-0")), map[string]string{"labels": "search"}))
}
