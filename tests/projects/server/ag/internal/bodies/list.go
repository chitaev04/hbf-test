package bodies

import (
	sgroupsv1 "HBF-tests/pkg/api/sgroups/v1"
	"HBF-tests/tests/common/foundation"

	common "github.com/PRO-Robotech/sgroups-proto/pkg/api/common"
)

// fs Собирает FieldSelector: имя, namespace (пустая строка = не задан) и список refs
func fs(name, namespace string, refs ...*common.ResourceRef) *common.FieldSelector {
	return &common.FieldSelector{Name: name, Namespace: namespace, Refs: refs}
}

// sel Собирает ResSelector из FieldSelector и/или LabelSelector
func sel(field *common.FieldSelector, labels map[string]string) *common.ResSelector {
	return &common.ResSelector{FieldSelector: field, LabelSelector: labels}
}

// ls Собирает ResSelector только с LabelSelector
func ls(labels map[string]string) *common.ResSelector {
	return &common.ResSelector{LabelSelector: labels}
}

// ref Собирает ResourceRef
func ref(name, namespace, resType string) *common.ResourceRef {
	return &common.ResourceRef{Name: name, Namespace: namespace, ResType: resType}
}

func listAGBody(testName string, selectors ...*common.ResSelector) *foundation.TestCaseBodyAGList {
	return &foundation.TestCaseBodyAGList{
		TestName: testName,
		Req:      sgroupsv1.AddressGroupReq_List{Selectors: selectors},
	}
}

// ListAll Поиск всех AG
func ListAll() *foundation.TestCaseBodyAGList {
	return &foundation.TestCaseBodyAGList{TestName: "Поиск всех AG", Req: sgroupsv1.AddressGroupReq_List{}}
}

// ListByName Поиск AG по name
func ListByName() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск AG по name", sel(fs("ag-0", ""), nil))
}

// ListByNonExistentName Поиск AG по несуществующему name
func ListByNonExistentName() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск AG по несуществующему name", sel(fs("ag-000", ""), nil))
}

// ListByExistingAndNonExistentName Поиск AG по существующему+несуществующему name
func ListByExistingAndNonExistentName() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск AG по существующему+несуществующему name",
		sel(fs("ag-000", ""), nil), sel(fs("ag-1", ""), nil))
}

// ListByTwoNames Поиск 2 AG по names
func ListByTwoNames() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск 2 AG по names", sel(fs("ag-0", ""), nil), sel(fs("ag-1", ""), nil))
}

// ListByNamespace Поиск AG по namespace
func ListByNamespace() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск AG по namespace", sel(fs("", "namespace-0"), nil))
}

// ListByNonExistentNamespace Поиск AG по несуществующему namespace
func ListByNonExistentNamespace() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск AG по несуществующему namespace", sel(fs("", "namespace-2222"), nil))
}

// ListByExistingAndNonExistentNamespace Поиск AG по существующему+несуществующему namespace
func ListByExistingAndNonExistentNamespace() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск AG по существующему+несуществующему namespace",
		sel(fs("", "namespace-2222"), nil), sel(fs("", "namespace-0"), nil))
}

// ListByTwoNamespaces Поиск 2 AG по namespaces
func ListByTwoNamespaces() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск 2 AG по namespaces", sel(fs("", "namespace-0"), nil), sel(fs("", "namespace-2"), nil))
}

// ListByNameAndNamespace Поиск AG по name+namespace
func ListByNameAndNamespace() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск AG по name+namespace", sel(fs("ag-0", "namespace-0"), nil))
}

// ListByNameAndNamespaceNamespaceNotExist Поиск AG по name+namespace, namespace не существует
func ListByNameAndNamespaceNamespaceNotExist() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск AG по name+namespace, namespace не существует", sel(fs("ag-0", "namespace-22"), nil))
}

// ListByNameAndNamespaceNameNotExist Поиск AG по name+namespace, name не существует
func ListByNameAndNamespaceNameNotExist() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск AG по name+namespace, name не существует", sel(fs("ag-33", "namespace-0"), nil))
}

// ListByNameAndNamespaceNoneExist Поиск AG по name+namespace, ни одного из не существует
func ListByNameAndNamespaceNoneExist() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск AG по name+namespace, ни одного из не существует", sel(fs("ag-33", "namespace-22"), nil))
}

// ListByNameNamespaceLabels Поиск AG по name+namespace+labels
func ListByNameNamespaceLabels() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск AG по name+namespace+labels",
		sel(fs("ag-0", "namespace-0"), map[string]string{"labels": "search"}))
}

// ListByNameNamespaceLabelsNameNotExist Поиск AG по name+namespace+labels, name не существует
func ListByNameNamespaceLabelsNameNotExist() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск AG по name+namespace+labels, name не существует",
		sel(fs("ag-33", "namespace-2"), map[string]string{"search": "nameLabels"}))
}

// ListByNameNamespaceLabelsNamespaceNotExist Поиск AG по name+namespace+labels, namespace не существует
func ListByNameNamespaceLabelsNamespaceNotExist() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск AG по name+namespace+labels, namespace не существует",
		sel(fs("ag-2", "namespace-22"), map[string]string{"search": "nameLabels"}))
}

// ListByNameNamespaceLabelsLabelsNotExist Поиск AG по name+namespace+labels, labels не существует
func ListByNameNamespaceLabelsLabelsNotExist() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск AG по name+namespace+labels, labels не существует",
		sel(fs("ag-2", "namespace-2"), map[string]string{"not": "exist"}))
}

// ListByNameNamespaceLabelsNoneExist Поиск AG по name+namespace+labels, ни одного не существует
func ListByNameNamespaceLabelsNoneExist() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск AG по name+namespace+labels, ни одного не существует",
		sel(fs("ag-33", "namespace-22"), map[string]string{"not": "exist"}))
}

// ListByNamespaceLabels Поиск AG по namespace+labels
func ListByNamespaceLabels() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск AG по namespace+labels", sel(fs("", "namespace-0"), map[string]string{"labels": "search"}))
}

// ListByNamespaceLabelsNamespaceNotExist Поиск AG по namespace+labels, namespace не существует
func ListByNamespaceLabelsNamespaceNotExist() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск AG по namespace+labels, namespace не существует",
		sel(fs("", "namespace-22"), map[string]string{"search": "nameLabels"}))
}

// ListByNamespaceLabelsLabelsNotExist Поиск AG по namespace+labels, labels не существует
func ListByNamespaceLabelsLabelsNotExist() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск AG по namespace+labels, labels не существует",
		sel(fs("", "namespace-2"), map[string]string{"not": "exist"}))
}

// ListByNamespaceLabelsNoneExist Поиск AG по namespace+labels, ни одного из не существует
func ListByNamespaceLabelsNoneExist() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск AG по namespace+labels, ни одного из не существует",
		sel(fs("", "namespace-22"), map[string]string{"not": "exist"}))
}

// ListByNameLabels Поиск AG по name+labels
func ListByNameLabels() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск AG по name+labels", sel(fs("ag-2", ""), map[string]string{"search": "nameLabels"}))
}

// ListByNameLabelsNameNotExist Поиск AG по name+labels, name не существует
func ListByNameLabelsNameNotExist() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск AG по name+labels, name не существует",
		sel(fs("ag-22", ""), map[string]string{"search": "nameLabels"}))
}

// ListByNameLabelsLabelsNotExist Поиск AG по name+labels, labels не существует
func ListByNameLabelsLabelsNotExist() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск AG по name+labels, labels не существует",
		sel(fs("ag-2", ""), map[string]string{"not": "exist"}))
}

// ListByNameLabelsNoneExist Поиск AG по name+labels, ни одного из не существует
func ListByNameLabelsNoneExist() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск AG по name+labels, ни одного из не существует",
		sel(fs("ag-22", ""), map[string]string{"not": "exist"}))
}

// ListByLabels Поиск AG по labels
func ListByLabels() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск AG по labels", ls(map[string]string{"labels": "search"}))
}

// ListByTwoLabels Поиск 2 AG по labels
func ListByTwoLabels() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск 2 AG по labels", ls(map[string]string{"labels": "search"}), ls(map[string]string{"labels": "ns"}))
}

// ListByNonExistentLabels Поиск AG по несуществующим labels
func ListByNonExistentLabels() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск AG по несуществующим labels", ls(map[string]string{"not": "exist"}))
}

// ListByExistingAndNonExistentLabels Поиск AG по сущ+несущ labels
func ListByExistingAndNonExistentLabels() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск AG по сущ+несущ labels", ls(map[string]string{"not": "exist"}), ls(map[string]string{"labels": "search"}))
}

// ListByRefSVC Поиск AG по ref (svc)
func ListByRefSVC() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск AG по ref (svc)", sel(fs("", "", ref("svc-0", "namespace-0", "Service")), nil))
}

// ListByRefSVCExistingAndNonExistentInOneObject Поиск AG по сущ+несущ ref внутри одного объекта (svc)
func ListByRefSVCExistingAndNonExistentInOneObject() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск AG по сущ+несущ ref  внутри одного объекта (svc)",
		sel(fs("", "", ref("svc-1", "namespace-1", "Service"), ref("svc-11", "namespace-1", "Service")), nil))
}

// ListByRefSVCExistingButNotBelongingInOneObject Поиск AG по сущ+непринадлежащей ref внутри одного объекта (svc)
func ListByRefSVCExistingButNotBelongingInOneObject() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск AG по сущ+непринадлежащей ref  внутри одного объекта (svc)",
		sel(fs("", "", ref("svc-1", "namespace-1", "Service"), ref("svc-2", "namespace-2", "Service")), nil))
}

// ListByRefSVCExistingAndNonExistentDifferentObjects Поиск AG по сущ+несущ ref в разных объектах (svc)
func ListByRefSVCExistingAndNonExistentDifferentObjects() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск AG по сущ+несущ ref в разных объектах (svc)",
		sel(fs("", "", ref("svc-2", "namespace-2", "Service")), nil),
		sel(fs("", "", ref("svc-12", "namespace-1", "Service")), nil))
}

// ListByTwoRefSVC Поиск 2 AG по ref (svc)
func ListByTwoRefSVC() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск 2 AG по ref (svc)",
		sel(fs("", "", ref("svc-0", "namespace-0", "Service")), nil),
		sel(fs("", "", ref("svc-1", "namespace-1", "Service")), nil))
}

// ListByRefSVCNameNotExistInNamespace Поиск AG по ref, name не существует в указанном ns (svc)
func ListByRefSVCNameNotExistInNamespace() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск AG по ref, name не существует в указанном ns (svc)",
		sel(fs("", "", ref("svc-2", "namespace-0", "Service")), nil))
}

// ListByRefSVCNameNotExist Поиск AG по ref, name не существует (svc)
func ListByRefSVCNameNotExist() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск AG по ref, name не существует  (svc)",
		sel(fs("", "", ref("svc-22", "namespace-0", "Service")), nil))
}

// ListByRefSVCNamespaceNotExist Поиск AG по ref, namespace не существует (svc)
func ListByRefSVCNamespaceNotExist() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск AG по ref, namespace не существует  (svc)",
		sel(fs("", "", ref("svc-2", "namespace-00", "Service")), nil))
}

// ListByRefHost Поиск AG по ref (host)
func ListByRefHost() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск AG по ref (host)", sel(fs("", "", ref("host-0", "namespace-0", "Host")), nil))
}

// ListByRefHostExistingAndNonExistentInOneObject Поиск AG по сущ+несущ ref внутри одного объекта (host)
func ListByRefHostExistingAndNonExistentInOneObject() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск AG по сущ+несущ ref  внутри одного объекта (host)",
		sel(fs("", "", ref("host-1", "namespace-1", "Host"), ref("host-11", "namespace-12", "Host")), nil))
}

// ListByRefHostExistingAndNonExistentDifferentObjects Поиск AG по сущ+несущ ref в разных объектах (host)
func ListByRefHostExistingAndNonExistentDifferentObjects() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск AG по сущ+несущ ref в разных объектах (host)",
		sel(fs("", "", ref("host-0", "namespace-0", "Host")), nil),
		sel(fs("", "", ref("host-12", "namespace-1", "Host")), nil))
}

// ListByTwoRefHost Поиск 2 AG по ref (hosts)
func ListByTwoRefHost() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск 2 AG по ref (hosts)",
		sel(fs("", "", ref("host-0", "namespace-0", "Host")), nil),
		sel(fs("", "", ref("host-1", "namespace-1", "Host")), nil))
}

// ListByRefHostNameNotExistInNamespace Поиск AG по ref, name не существует в указанном ns (hosts)
func ListByRefHostNameNotExistInNamespace() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск AG по ref, name не существует в указанном ns (hosts)",
		sel(fs("", "", ref("host-3", "namespace-0", "Host")), nil))
}

// ListByRefHostNameNotExist Поиск AG по ref, name не существует (hosts)
func ListByRefHostNameNotExist() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск AG по ref, name не существует  (hosts)",
		sel(fs("", "", ref("host-33", "namespace-0", "Host")), nil))
}

// ListByRefHostNamespaceNotExist Поиск AG по ref, namespace не существует (hosts)
func ListByRefHostNamespaceNotExist() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск AG по ref, namespace не существует  (hosts)",
		sel(fs("", "", ref("host-3", "namespace-000", "Host")), nil))
}

// ListByRefNetwork Поиск AG по ref (network)
func ListByRefNetwork() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск AG по ref (network)", sel(fs("", "", ref("nw-2", "namespace-2", "Network")), nil))
}

// ListByTwoRefNetwork Поиск 2 AG по ref (network)
func ListByTwoRefNetwork() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск 2 AG по ref (network)",
		sel(fs("", "", ref("nw-0", "namespace-0", "Network")), nil),
		sel(fs("", "", ref("nw-1", "namespace-1", "Network")), nil))
}

// ListByRefNetworkExistingAndNonExistentInOneObject Поиск AG по сущ+несущ ref внутри одного объекта (network)
func ListByRefNetworkExistingAndNonExistentInOneObject() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск AG по сущ+несущ ref  внутри одного объекта (network)",
		sel(fs("", "", ref("nw-1", "namespace-1", "Network"), ref("nw-1", "namespace-12", "Network")), nil))
}

// ListByRefNetworkExistingButNotBelongingInOneObject Поиск AG по сущ+непринадлежащему ref внутри одного объекта (network)
func ListByRefNetworkExistingButNotBelongingInOneObject() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск AG по сущ+непринадлежащему ref  внутри одного объекта (network)",
		sel(fs("", "", ref("nw-1", "namespace-1", "Network"), ref("nw-0", "namespace-0", "Network")), nil))
}

// ListByRefNetworkExistingAndNonExistentDifferentObjects Поиск AG по сущ+несущ ref в разных объектах (network)
func ListByRefNetworkExistingAndNonExistentDifferentObjects() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск AG по сущ+несущ ref в разных объектах (network)",
		sel(fs("", "", ref("nw-0", "namespace-0", "Network")), nil),
		sel(fs("", "", ref("nw-12", "namespace-1", "Network")), nil))
}

// ListByRefNetworkNameNotExistInNamespace Поиск AG по ref, name не существует в указанном ns (network)
func ListByRefNetworkNameNotExistInNamespace() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск AG по ref, name не существует в указанном ns (network)",
		sel(fs("", "", ref("nw-1", "namespace-2", "Network")), nil))
}

// ListByRefNetworkNameNotExist Поиск AG по ref, name не существует (network)
func ListByRefNetworkNameNotExist() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск AG по ref, name не существует  (network)",
		sel(fs("", "", ref("nw-11", "namespace-1", "Network")), nil))
}

// ListByRefNetworkNamespaceNotExist Поиск AG по ref, namespace не существует (network)
func ListByRefNetworkNamespaceNotExist() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск AG по ref, namespace не существует  (network)",
		sel(fs("", "", ref("nw-1", "namespace-10", "Network")), nil))
}

// ListByThreeRefs Поиск AG по 3 ref (nw+svc+host)
func ListByThreeRefs() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск AG по  3 ref (nw+svc+host)",
		sel(fs("", "",
			ref("host-2", "namespace-2", "Host"),
			ref("nw-2", "namespace-2", "Network"),
			ref("svc-2", "namespace-2", "Service"),
		), nil))
}

// ListByRefAndName Поиск AG по ref+name
func ListByRefAndName() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск AG по ref+name",
		sel(fs("ag-2", "",
			ref("host-2", "namespace-2", "Host"),
			ref("nw-2", "namespace-2", "Network"),
			ref("svc-2", "namespace-2", "Service"),
		), nil))
}

// ListByRefAndNameNameNotExist Поиск AG по ref+name (name не существует)
func ListByRefAndNameNameNotExist() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск AG по ref+name (name не существует)",
		sel(fs("ag-22", "",
			ref("host-2", "namespace-2", "Host"),
			ref("nw-2", "namespace-2", "Network"),
			ref("svc-2", "namespace-2", "Service"),
		), nil))
}

// ListByRefAndNameRefNotExist Поиск AG по ref+name (ref не существует)
func ListByRefAndNameRefNotExist() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск AG по ref+name (ref не существует)",
		sel(fs("ag-2", "",
			ref("host-2", "namespace-2", "Host"),
			ref("nw-2", "namespace-2", "Network"),
			ref("svc-0", "namespace-0", "Service"),
		), nil))
}

// ListByRefAndNameNoneExist Поиск AG по ref+name (Ни одного из не существует)
func ListByRefAndNameNoneExist() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск AG по ref+name (Ни одного из не существует)",
		sel(fs("ag-22", "",
			ref("host-02", "namespace-02", "Host"),
			ref("nw-02", "namespace-02", "Network"),
			ref("svc-02", "namespace-02", "Service"),
		), nil))
}

// ListByRefAndNamespace Поиск AG по ref+namespace
func ListByRefAndNamespace() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск AG по ref+namespace",
		sel(fs("", "namespace-2",
			ref("host-2", "namespace-2", "Host"),
			ref("nw-2", "namespace-2", "Network"),
			ref("svc-2", "namespace-2", "Service"),
		), nil))
}

// ListByRefAndNamespaceNamespaceNotExist Поиск AG по ref+namespace (namespace не существует)
func ListByRefAndNamespaceNamespaceNotExist() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск AG по ref+namespace (namespace не существует)",
		sel(fs("", "namespace-22",
			ref("host-2", "namespace-2", "Host"),
			ref("nw-2", "namespace-2", "Network"),
			ref("svc-2", "namespace-2", "Service"),
		), nil))
}

// ListByRefAndNamespaceRefNotExist Поиск AG по ref+namespace (ref не существует)
func ListByRefAndNamespaceRefNotExist() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск AG по ref+namespace (ref не существует)",
		sel(fs("", "namespace-2",
			ref("host-0", "namespace-0", "Host"),
			ref("nw-0", "namespace-0", "Network"),
			ref("svc-0", "namespace-0", "Service"),
		), nil))
}

// ListByRefAndNamespaceNoneExist Поиск AG по ref+namespace (ни одного из не существует)
func ListByRefAndNamespaceNoneExist() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск AG по ref+namespace (ни одного из не существует)",
		sel(fs("", "namespace-22",
			ref("host-02", "namespace-02", "Host"),
			ref("nw-02", "namespace-02", "Network"),
			ref("svc-02", "namespace-02", "Service"),
		), nil))
}

// ListByRefAndLabels Поиск AG по ref+labels
func ListByRefAndLabels() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск AG по ref+labels",
		sel(fs("", "",
			ref("host-2", "namespace-2", "Host"),
			ref("nw-2", "namespace-2", "Network"),
			ref("svc-2", "namespace-2", "Service"),
		), map[string]string{"search": "nameLabels"}))
}

// ListByRefAndLabelsLabelsNotExist Поиск AG по ref+labels (labels не существуют)
func ListByRefAndLabelsLabelsNotExist() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск AG по ref+labels (labels не существуют)",
		sel(fs("", "",
			ref("host-2", "namespace-2", "Host"),
			ref("nw-2", "namespace-2", "Network"),
			ref("svc-2", "namespace-2", "Service"),
		), map[string]string{"search": "nameLabels", "h": "nameLabels"}))
}

// ListByRefAndLabelsRefNotExist Поиск AG по ref+labels (ref не существуют)
func ListByRefAndLabelsRefNotExist() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск AG по ref+labels (ref не существуют)",
		sel(fs("", "",
			ref("host-02", "namespace-2", "Host"),
			ref("nw-02", "namespace-2", "Network"),
			ref("svc-02", "namespace-2", "Service"),
		), map[string]string{"search": "nameLabels"}))
}

// ListByRefAndLabelsNoneExist Поиск AG по ref+labels (ни одного из не существует)
func ListByRefAndLabelsNoneExist() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск AG по ref+labels (ни одного из не существует)",
		sel(fs("", "",
			ref("host-02", "namespace-2", "Host"),
			ref("nw-02", "namespace-2", "Network"),
			ref("svc-02", "namespace-2", "Service"),
		), map[string]string{"search": "nameLabels"}))
}

// ListByAllParams Поиск AG по всем параметрам (name+namespace+ref+labels)
func ListByAllParams() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск AG по по всем параметрам (name+namespace+ref+labels)",
		sel(fs("ag-2", "namespace-2",
			ref("host-2", "namespace-2", "Host"),
			ref("nw-2", "namespace-2", "Network"),
			ref("svc-2", "namespace-2", "Service"),
		), map[string]string{"search": "nameLabels"}))
}

// ListByAllParamsOneNotExist Поиск AG по всем параметрам, один из параметров не существует
func ListByAllParamsOneNotExist() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск AG по по всем параметрам, один из параметров не существует",
		sel(fs("ag-0", "namespace-2",
			ref("host-2", "namespace-2", "Host"),
			ref("nw-2", "namespace-2", "Network"),
			ref("svc-2", "namespace-2", "Service"),
		), map[string]string{"search": "nameLabels"}))
}

// ListErrorUnknownResType Ошибка при поиске AG, несуществующий resType
func ListErrorUnknownResType() *foundation.TestCaseBodyAGList {
	return listAGBody("Ошибка при поиске AG, несущетвующий resType",
		sel(fs("ag-0", "namespace-2",
			ref("host-2", "namespace-2", "Host"),
			ref("nw-2", "namespace-2", "Network"),
			ref("svc-2", "namespace-2", "Svc"),
		), map[string]string{"search": "nameLabels"}))
}

// ListByTwoRefRule Поиск 2 AG по ref (rule)
func ListByTwoRefRule() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск 2 AG по ref (rule)",
		sel(fs("", "", ref("rule-ag-ag-icmp-0", "namespace-0", "Ag2AgIcmpRule")), nil),
		sel(fs("", "", ref("rule-ag-icmp-0", "namespace-1", "Ag2IcmpRule")), nil))
}

// ListByRefRule Поиск AG по ref (rule)
func ListByRefRule() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск AG по ref (rule)", sel(fs("", "", ref("rule-ag-cidr-2", "namespace-2", "Ag2CidrRule")), nil))
}

// ListByRefRuleExistingAndNonExistentInOneObject Поиск AG по сущ+несущ ref внутри одного объекта (rule)
func ListByRefRuleExistingAndNonExistentInOneObject() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск AG по сущ+несущ ref  внутри одного объекта (rule)",
		sel(fs("", "", ref("rule-ag-ag-1", "namespace-1", "Ag2AgRule"), ref("rule-ag-ag-11", "namespace-1", "Ag2AgRule")), nil))
}

// ListByRefRuleExistingButNotBelongingInOneObject Поиск AG по сущ+непринадлежащему ref внутри одного объекта (rule)
func ListByRefRuleExistingButNotBelongingInOneObject() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск AG по сущ+непринадлежащему ref  внутри одного объекта (rule)",
		sel(fs("", "", ref("nw-1", "namespace-1", "Network"), ref("rule-ag-ag-icmp-2", "namespace-2", "Ag2AgIcmpRule")), nil))
}

// ListByRefRuleExistingAndNonExistentDifferentObjects Поиск AG по сущ+несущ ref в разных объектах (rule)
func ListByRefRuleExistingAndNonExistentDifferentObjects() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск AG по сущ+несущ ref в разных объектах (rule)",
		sel(fs("", "", ref("rule-ag-ag-icmp-0", "namespace-0", "Ag2AgIcmpRule")), nil),
		sel(fs("", "", ref("rule-ag-ag-icmp-00", "namespace-0", "Ag2AgIcmpRule")), nil))
}

// ListByRefRuleNameNotExistInNamespace Поиск AG по ref, name не существует в указанном ns (rule)
func ListByRefRuleNameNotExistInNamespace() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск AG по ref, name не существует в указанном ns (rule)",
		sel(fs("", "", ref("rule-ag-ag-1", "namespace-3", "Ag2AgRule")), nil))
}

// ListByRefRuleNameNotExist Поиск AG по ref, name не существует (rule)
func ListByRefRuleNameNotExist() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск AG по ref, name не существует  (rule)",
		sel(fs("", "", ref("rule-ag-ag-icmp-10", "namespace-1", "Ag2AgIcmpRule")), nil))
}

// ListByRefRuleNamespaceNotExist Поиск AG по ref, namespace не существует (rule)
func ListByRefRuleNamespaceNotExist() *foundation.TestCaseBodyAGList {
	return listAGBody("Поиск AG по ref, namespace не существует  (rule)",
		sel(fs("", "", ref("rule-ag-fqdn-2", "namespace-33", "Ag2FqdnRule")), nil))
}
