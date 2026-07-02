package suites

import (
	"github.com/ozontech/allure-go/pkg/framework/provider"

	sgroupsv1 "HBF-tests/pkg/api/sgroups/v1"

	common "github.com/PRO-Robotech/sgroups-proto/pkg/api/common"
)

// networkBindingExpectation Ожидаемые поля network-binding для проверки в ответах Upsert/List
type networkBindingExpectation struct {
	namespace             string
	uid                   string
	labels                map[string]string
	annotations           map[string]string
	displayName           string
	comment               string
	description           string
	addressGroupName      string
	addressGroupNamespace string
	networkName           string
	networkNamespace      string
}

// networkBindingLike Общий интерфейс ответов NetworkBinding (Upsert и List используют один и тот же тип)
type networkBindingLike interface {
	GetMetadata() *common.Metadata
	GetSpec() *sgroupsv1.NetworkBinding_Spec
}

// assertNetworkBinding Проверка полей одного network-binding
func assertNetworkBinding(sCtx provider.StepCtx, nb networkBindingLike, name string, exp networkBindingExpectation) {
	md, spec := nb.GetMetadata(), nb.GetSpec()
	sCtx.Assert().Equal(name, md.Name)
	sCtx.Assert().Equal(exp.namespace, md.Namespace)
	if exp.uid != "" {
		sCtx.Assert().Equal(exp.uid, md.Uid)
	} else {
		sCtx.Assert().NotEmpty(md.Uid)
	}
	sCtx.Assert().Equal(exp.labels, md.Labels)
	sCtx.Assert().Equal(exp.annotations, md.Annotations)
	sCtx.Assert().NotEmpty(md.ResourceVersion)
	sCtx.Assert().Equal(exp.displayName, spec.DisplayName)
	sCtx.Assert().Equal(exp.comment, spec.Comment)
	sCtx.Assert().Equal(exp.description, spec.Description)
	if exp.addressGroupName != "" || exp.addressGroupNamespace != "" {
		sCtx.Require().NotNil(spec.AddressGroup)
		sCtx.Assert().Equal(exp.addressGroupName, spec.AddressGroup.Name)
		sCtx.Assert().Equal(exp.addressGroupNamespace, spec.AddressGroup.Namespace)
	}
	if exp.networkName != "" || exp.networkNamespace != "" {
		sCtx.Require().NotNil(spec.Network)
		sCtx.Assert().Equal(exp.networkName, spec.Network.Name)
		sCtx.Assert().Equal(exp.networkNamespace, spec.Network.Namespace)
	}
}

// assertNetworkBindingsByName Проверка списка network-binding, сопоставленных по имени
func assertNetworkBindingsByName(sCtx provider.StepCtx, nbs []*sgroupsv1.NetworkBinding, expected map[string]networkBindingExpectation) {
	for _, nb := range nbs {
		exp, ok := expected[nb.Metadata.Name]
		sCtx.Assert().True(ok, "unexpected network-binding: %s", nb.Metadata.Name)
		if !ok {
			continue
		}
		assertNetworkBinding(sCtx, nb, nb.Metadata.Name, exp)
	}
}

// Фикстуры для повторяющихся в разных кейсах network-binding (данные из тестового окружения)

var nb0Expectation = networkBindingExpectation{
	namespace: "namespace-0", uid: "90ef841a-f480-405d-abb0-983baf038801",
	labels: map[string]string{"labels": "search"}, annotations: map[string]string{"search": "labels"},
	displayName: "network binding 0", comment: "for search by name/ns", description: "network binding for search",
	addressGroupName: "ag-0", addressGroupNamespace: "namespace-0", networkName: "nw-0", networkNamespace: "namespace-0",
}

var nb1Expectation = networkBindingExpectation{
	namespace: "namespace-1", uid: "9b187907-29ef-4de1-a798-d6537c8a95b1",
	labels: map[string]string{"labels": "ns"}, annotations: map[string]string{"search": "name"},
	displayName: "network binding 1", comment: "for search by name/ns", description: "network binding for search",
	addressGroupName: "ag-1", addressGroupNamespace: "namespace-1", networkName: "nw-1", networkNamespace: "namespace-1",
}

var nb2Expectation = networkBindingExpectation{
	namespace: "namespace-2", uid: "fdcb56b7-94fe-412d-b7c7-e5ccf1f94085",
	labels: map[string]string{"search": "nameLabels"}, annotations: map[string]string{},
	displayName: "network binding 2", comment: "for search by labels", description: "network binding for search labels",
	addressGroupName: "ag-2", addressGroupNamespace: "namespace-2", networkName: "nw-2", networkNamespace: "namespace-2",
}

var nb0And1 = map[string]networkBindingExpectation{"nb-0": nb0Expectation, "nb-1": nb1Expectation}
var nb0And2 = map[string]networkBindingExpectation{"nb-0": nb0Expectation, "nb-2": nb2Expectation}
