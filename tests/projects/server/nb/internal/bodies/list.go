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

// fs Собирает FieldSelector для network-binding
func fs(name, namespace string, ag, network *common.ResourceIdentifier) *sgroupsv1.NetworkBindingReq_Selectors_FieldSelector {
	return &sgroupsv1.NetworkBindingReq_Selectors_FieldSelector{Name: name, Namespace: namespace, AddressGroup: ag, Network: network}
}

// sel Собирает Selectors из FieldSelector и/или LabelSelector
func sel(field *sgroupsv1.NetworkBindingReq_Selectors_FieldSelector, labels map[string]string) *sgroupsv1.NetworkBindingReq_Selectors {
	return &sgroupsv1.NetworkBindingReq_Selectors{FieldSelector: field, LabelSelector: labels}
}

// ls Собирает Selectors только с LabelSelector
func ls(labels map[string]string) *sgroupsv1.NetworkBindingReq_Selectors {
	return &sgroupsv1.NetworkBindingReq_Selectors{LabelSelector: labels}
}

func listNBBody(testName string, selectors ...*sgroupsv1.NetworkBindingReq_Selectors) *foundation.TestCaseBodyNBList {
	return &foundation.TestCaseBodyNBList{
		TestName: testName,
		Req:      sgroupsv1.NetworkBindingReq_List{Selectors: selectors},
	}
}

// ListAll Поиск всех network-binding
func ListAll() *foundation.TestCaseBodyNBList {
	return &foundation.TestCaseBodyNBList{TestName: "Поиск всех network-binding", Req: sgroupsv1.NetworkBindingReq_List{}}
}

// ListByName Поиск network-binding по name
func ListByName() *foundation.TestCaseBodyNBList {
	return listNBBody("Поиск network-binding по name", sel(fs("nb-0", "", nil, nil), nil))
}

// ListByNonExistentName Поиск network-binding по несуществующему name
func ListByNonExistentName() *foundation.TestCaseBodyNBList {
	return listNBBody("Поиск network-binding по несуществующему name", sel(fs("host-000", "", nil, nil), nil))
}

// ListByExistingAndNonExistentName Поиск network-binding по существующему+несуществующему name
func ListByExistingAndNonExistentName() *foundation.TestCaseBodyNBList {
	return listNBBody("Поиск network-binding по существующему+несуществующему name",
		sel(fs("nb-000", "", nil, nil), nil), sel(fs("nb-1", "", nil, nil), nil))
}

// ListByTwoNames Поиск 2 network-binding по names
func ListByTwoNames() *foundation.TestCaseBodyNBList {
	return listNBBody("Поиск 2 network-binding по names", sel(fs("nb-0", "", nil, nil), nil), sel(fs("nb-1", "", nil, nil), nil))
}

// ListByNamespace Поиск network-binding по namespace
func ListByNamespace() *foundation.TestCaseBodyNBList {
	return listNBBody("Поиск network-binding по namespace", sel(fs("", "namespace-0", nil, nil), nil))
}

// ListByNonExistentNamespace Поиск network-binding по несуществующему namespace
func ListByNonExistentNamespace() *foundation.TestCaseBodyNBList {
	return listNBBody("Поиск network-binding по несуществующему namespace", sel(fs("", "namespace-2222", nil, nil), nil))
}

// ListByExistingAndNonExistentNamespace Поиск network-binding по существующему+несуществующему namespace
func ListByExistingAndNonExistentNamespace() *foundation.TestCaseBodyNBList {
	return listNBBody("Поиск network-binding по существующему+несуществующему namespace",
		sel(fs("", "namespace-2222", nil, nil), nil), sel(fs("", "namespace-0", nil, nil), nil))
}

// ListByTwoNamespaces Поиск 2 network-binding по namespaces
func ListByTwoNamespaces() *foundation.TestCaseBodyNBList {
	return listNBBody("Поиск 2 network-binding по namespaces", sel(fs("", "namespace-0", nil, nil), nil), sel(fs("", "namespace-2", nil, nil), nil))
}

// ListByNameAndNamespace Поиск network-binding по name+namespace
func ListByNameAndNamespace() *foundation.TestCaseBodyNBList {
	return listNBBody("Поиск network-binding по name+namespace", sel(fs("nb-0", "namespace-0", nil, nil), nil))
}

// ListByNameAndNamespaceNamespaceNotExist Поиск network-binding по name+namespace, namespace не существует
func ListByNameAndNamespaceNamespaceNotExist() *foundation.TestCaseBodyNBList {
	return listNBBody("Поиск network-binding по name+namespace, namespace не существует", sel(fs("ag-0", "namespace-22", nil, nil), nil))
}

// ListByNameAndNamespaceNameNotExist Поиск network-binding по name+namespace, name не существует
func ListByNameAndNamespaceNameNotExist() *foundation.TestCaseBodyNBList {
	return listNBBody("Поиск network-binding по name+namespace, name не существует", sel(fs("host-33", "namespace-0", nil, nil), nil))
}

// ListByNameAndNamespaceNoneExist Поиск network-binding по name+namespace, ни одного из не существует
func ListByNameAndNamespaceNoneExist() *foundation.TestCaseBodyNBList {
	return listNBBody("Поиск network-binding по name+namespace, ни одного из не существует", sel(fs("host-33", "namespace-22", nil, nil), nil))
}

// ListByNameNamespaceLabels Поиск network-binding по name+namespace+labels
func ListByNameNamespaceLabels() *foundation.TestCaseBodyNBList {
	return listNBBody("Поиск network-binding по name+namespace+labels",
		sel(fs("nb-2", "namespace-2", nil, nil), map[string]string{"search": "nameLabels"}))
}

// ListByNameNamespaceLabelsNameNotExist Поиск network-binding по name+namespace+labels, name не существует
func ListByNameNamespaceLabelsNameNotExist() *foundation.TestCaseBodyNBList {
	return listNBBody("Поиск network-binding по name+namespace+labels, name не существует",
		sel(fs("nb-33", "namespace-2", nil, nil), map[string]string{"search": "nameLabels"}))
}

// ListByNameNamespaceLabelsNamespaceNotExist Поиск network-binding по name+namespace+labels, namespace не существует
func ListByNameNamespaceLabelsNamespaceNotExist() *foundation.TestCaseBodyNBList {
	return listNBBody("Поиск network-binding по name+namespace+labels, namespace не существует",
		sel(fs("nb-2", "namespace-22", nil, nil), map[string]string{"search": "nameLabels"}))
}

// ListByNameNamespaceLabelsLabelsNotExist Поиск network-binding по name+namespace+labels, labels не существует
func ListByNameNamespaceLabelsLabelsNotExist() *foundation.TestCaseBodyNBList {
	return listNBBody("Поиск network-binding по name+namespace+labels, labels не существует",
		sel(fs("nb-2", "namespace-2", nil, nil), map[string]string{"not": "exist"}))
}

// ListByNameNamespaceLabelsNoneExist Поиск network-binding по name+namespace+labels, ни одного не существует
func ListByNameNamespaceLabelsNoneExist() *foundation.TestCaseBodyNBList {
	return listNBBody("Поиск network-binding по name+namespace+labels, ни одного не существует",
		sel(fs("nb-33", "namespace-22", nil, nil), map[string]string{"not": "exist"}))
}

// ListByNamespaceLabels Поиск network-binding по namespace+labels
func ListByNamespaceLabels() *foundation.TestCaseBodyNBList {
	return listNBBody("Поиск network-binding по namespace+labels", sel(fs("", "namespace-2", nil, nil), map[string]string{"search": "nameLabels"}))
}

// ListByNamespaceLabelsNamespaceNotExist Поиск network-binding по namespace+labels, namespace не существует
func ListByNamespaceLabelsNamespaceNotExist() *foundation.TestCaseBodyNBList {
	return listNBBody("Поиск network-binding по namespace+labels, namespace не существует",
		sel(fs("", "namespace-22", nil, nil), map[string]string{"search": "nameLabels"}))
}

// ListByNamespaceLabelsLabelsNotExist Поиск network-binding по namespace+labels, labels не существует
func ListByNamespaceLabelsLabelsNotExist() *foundation.TestCaseBodyNBList {
	return listNBBody("Поиск network-binding по namespace+labels, labels не существует",
		sel(fs("", "namespace-2", nil, nil), map[string]string{"not": "exist"}))
}

// ListByNamespaceLabelsNoneExist Поиск network-binding по namespace+labels, ни одного из не существует
func ListByNamespaceLabelsNoneExist() *foundation.TestCaseBodyNBList {
	return listNBBody("Поиск network-binding по namespace+labels, ни одного из не существует",
		sel(fs("", "namespace-22", nil, nil), map[string]string{"not": "exist"}))
}

// ListByNameLabels Поиск network-binding по name+labels
func ListByNameLabels() *foundation.TestCaseBodyNBList {
	return listNBBody("Поиск network-binding по name+labels", sel(fs("nb-2", "", nil, nil), map[string]string{"search": "nameLabels"}))
}

// ListByNameLabelsNameNotExist Поиск network-binding по name+labels, name не существует
func ListByNameLabelsNameNotExist() *foundation.TestCaseBodyNBList {
	return listNBBody("Поиск network-binding по name+labels, name не существует",
		sel(fs("host-22", "", nil, nil), map[string]string{"search": "nameLabels"}))
}

// ListByNameLabelsLabelsNotExist Поиск network-binding по name+labels, labels не существует
func ListByNameLabelsLabelsNotExist() *foundation.TestCaseBodyNBList {
	return listNBBody("Поиск network-binding по name+labels, labels не существует",
		sel(fs("host-2", "", nil, nil), map[string]string{"not": "exist"}))
}

// ListByNameLabelsNoneExist Поиск network-binding по name+labels, ни одного из не существует
func ListByNameLabelsNoneExist() *foundation.TestCaseBodyNBList {
	return listNBBody("Поиск network-binding по name+labels, ни одного из не существует",
		sel(fs("host-22", "", nil, nil), map[string]string{"not": "exist"}))
}

// ListByLabels Поиск network-binding по labels
func ListByLabels() *foundation.TestCaseBodyNBList {
	return listNBBody("Поиск network-binding по labels", ls(map[string]string{"labels": "search"}))
}

// ListByTwoLabels Поиск 2 network-binding по labels
func ListByTwoLabels() *foundation.TestCaseBodyNBList {
	return listNBBody("Поиск 2 network-binding по labels", ls(map[string]string{"labels": "search"}), ls(map[string]string{"labels": "ns"}))
}

// ListByNonExistentLabels Поиск network-binding по несуществующим labels
func ListByNonExistentLabels() *foundation.TestCaseBodyNBList {
	return listNBBody("Поиск network-binding по несуществующим labels", ls(map[string]string{"not": "exist"}))
}

// ListByExistingAndNonExistentLabels Поиск network-binding по сущ+несущ labels
func ListByExistingAndNonExistentLabels() *foundation.TestCaseBodyNBList {
	return listNBBody("Поиск network-binding по сущ+несущ labels", ls(map[string]string{"not": "exist"}), ls(map[string]string{"labels": "search"}))
}

// ListByExistingAG Поиск network-binding по существующей AG
func ListByExistingAG() *foundation.TestCaseBodyNBList {
	return listNBBody("Поиск network-binding по существующей AG", sel(fs("", "", ri("ag-0", "namespace-0"), nil), nil))
}

// ListByExistingNetwork Поиск network-binding по существующему nw
func ListByExistingNetwork() *foundation.TestCaseBodyNBList {
	return listNBBody("Поиск network-binding по существующему nw", sel(fs("", "", nil, ri("nw-0", "namespace-0")), nil))
}

// ListByExistingAndNonExistentAG Поиск network-binding по сущ+несущ AG
func ListByExistingAndNonExistentAG() *foundation.TestCaseBodyNBList {
	return listNBBody("Поиск network-binding по сущ+несущ AG",
		sel(fs("", "", ri("ag-0", "namespace-0"), nil), nil), sel(fs("", "", ri("ag-11", "namespace-1"), nil), nil))
}

// ListByExistingAndNonExistentNetwork Поиск network-binding по сущ+несущ nw
func ListByExistingAndNonExistentNetwork() *foundation.TestCaseBodyNBList {
	return listNBBody("Поиск network-binding по сущ+несущ nw",
		sel(fs("", "", nil, ri("nw-0", "namespace-0")), nil), sel(fs("", "", nil, ri("nw-11", "namespace-1")), nil))
}

// ListByNonExistentAG Поиск network-binding по несущ AG
func ListByNonExistentAG() *foundation.TestCaseBodyNBList {
	return listNBBody("Поиск network-binding по несущ AG", sel(fs("", "", ri("ag-11", "namespace-1"), nil), nil))
}

// ListByNonExistentNetwork Поиск network-binding по несущ nw
func ListByNonExistentNetwork() *foundation.TestCaseBodyNBList {
	return listNBBody("Поиск network-binding по несущ nw", sel(fs("", "", nil, ri("nw-11", "namespace-1")), nil))
}

// ListByTwoAG Поиск 2 network-binding по AG
func ListByTwoAG() *foundation.TestCaseBodyNBList {
	return listNBBody("Поиск 2 network-binding по AG",
		sel(fs("", "", ri("ag-0", "namespace-0"), nil), nil), sel(fs("", "", ri("ag-1", "namespace-1"), nil), nil))
}

// ListByTwoNetwork Поиск 2 network-binding по nw
func ListByTwoNetwork() *foundation.TestCaseBodyNBList {
	return listNBBody("Поиск 2 network-binding по nw",
		sel(fs("", "", nil, ri("nw-0", "namespace-0")), nil), sel(fs("", "", nil, ri("nw-1", "namespace-1")), nil))
}

// ListByAGAndName Поиск network-binding по AG+name
func ListByAGAndName() *foundation.TestCaseBodyNBList {
	return listNBBody("Поиск network-binding по AG+name", sel(fs("nb-0", "", ri("ag-0", "namespace-0"), nil), nil))
}

// ListByNetworkAndName Поиск network-binding по nw+name
func ListByNetworkAndName() *foundation.TestCaseBodyNBList {
	return listNBBody("Поиск network-binding по nw+name", sel(fs("nb-0", "", nil, ri("nw-0", "namespace-0")), nil))
}

// ListByAGNameNamespace Поиск network-binding по AG+name+ns
func ListByAGNameNamespace() *foundation.TestCaseBodyNBList {
	return listNBBody("Поиск network-binding по AG+name+ns", sel(fs("nb-0", "namespace-0", ri("ag-0", "namespace-0"), nil), nil))
}

// ListByNetworkNameNamespace Поиск network-binding по nw+name+ns
func ListByNetworkNameNamespace() *foundation.TestCaseBodyNBList {
	return listNBBody("Поиск network-binding по nw+name+ns", sel(fs("nb-0", "", nil, ri("nw-0", "namespace-0")), nil))
}

// ListByAGNameNamespaceLabels Поиск network-binding по AG+name+ns+labels
func ListByAGNameNamespaceLabels() *foundation.TestCaseBodyNBList {
	return listNBBody("Поиск network-binding по AG+name+ns+labels",
		sel(fs("nb-0", "", ri("ag-0", "namespace-0"), nil), map[string]string{"labels": "search"}))
}

// ListByNetworkNameNamespaceLabels Поиск network-binding по nw+name+ns+labels
func ListByNetworkNameNamespaceLabels() *foundation.TestCaseBodyNBList {
	return listNBBody("Поиск network-binding по nw+name+ns+labels",
		sel(fs("nb-0", "", nil, ri("nw-0", "namespace-0")), map[string]string{"labels": "search"}))
}

// ListByAGNameAGNotExist Поиск network-binding по AG+name, AG не существует
func ListByAGNameAGNotExist() *foundation.TestCaseBodyNBList {
	return listNBBody("Поиск network-binding по AG+name, AG не существует", sel(fs("nb-0", "", ri("ag-11", "namespace-0"), nil), nil))
}

// ListByNetworkNameNetworkNotExist Поиск network-binding по nw+name, nw не существует
func ListByNetworkNameNetworkNotExist() *foundation.TestCaseBodyNBList {
	return listNBBody("Поиск network-binding по nw+name, nw не существует", sel(fs("nb-0", "", nil, ri("nw-11", "namespace-0")), nil))
}

// ListByAGNameNameNotExist Поиск network-binding по AG+name, name не существует
func ListByAGNameNameNotExist() *foundation.TestCaseBodyNBList {
	return listNBBody("Поиск network-binding по AG+name, name не существует", sel(fs("nb-00", "", ri("ag-0", "namespace-0"), nil), nil))
}

// ListByNetworkNameNameNotExist Поиск network-binding по nw+name, name не существует
func ListByNetworkNameNameNotExist() *foundation.TestCaseBodyNBList {
	return listNBBody("Поиск network-binding по nw+name, name не существует", sel(fs("nb-00", "", nil, ri("nw-0", "namespace-0")), nil))
}

// ListByAGNameNoneExist Поиск network-binding по AG+name, ни одного из не существует
func ListByAGNameNoneExist() *foundation.TestCaseBodyNBList {
	return listNBBody("Поиск network-binding по AG+name, ни одного из не существует", sel(fs("nb-00", "", ri("ag-00", "namespace-0"), nil), nil))
}

// ListByAGAndNetwork Поиск network-binding по AG+NW
func ListByAGAndNetwork() *foundation.TestCaseBodyNBList {
	return listNBBody("Поиск network-binding по AG+NW", sel(fs("", "", ri("ag-0", "namespace-0"), ri("nw-0", "namespace-0")), nil))
}

// ListByAllParams Поиск network-binding по всем параметрам
func ListByAllParams() *foundation.TestCaseBodyNBList {
	return listNBBody("Поиск network-binding по всем параметрам",
		sel(fs("nb-0", "namespace-0", ri("ag-0", "namespace-0"), ri("nw-0", "namespace-0")), map[string]string{"labels": "search"}))
}
