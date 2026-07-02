package suites

import (
	"github.com/ozontech/allure-go/pkg/framework/provider"

	sgroupsv1 "HBF-tests/pkg/api/sgroups/v1"

	common "github.com/PRO-Robotech/sgroups-proto/pkg/api/common"
)

// hostBindingExpectation Ожидаемые поля host-binding для проверки в ответах Upsert/List
type hostBindingExpectation struct {
	namespace             string
	uid                   string
	labels                map[string]string
	annotations           map[string]string
	displayName           string
	comment               string
	description           string
	addressGroupName      string
	addressGroupNamespace string
	hostName              string
	hostNamespace         string
}

// hostBindingLike Общий интерфейс ответов HostBinding (Upsert и List используют один и тот же тип)
type hostBindingLike interface {
	GetMetadata() *common.Metadata
	GetSpec() *sgroupsv1.HostBinding_Spec
}

// assertHostBinding Проверка полей одного host-binding
func assertHostBinding(sCtx provider.StepCtx, hb hostBindingLike, name string, exp hostBindingExpectation) {
	md, spec := hb.GetMetadata(), hb.GetSpec()
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
	if exp.hostName != "" || exp.hostNamespace != "" {
		sCtx.Require().NotNil(spec.Host)
		sCtx.Assert().Equal(exp.hostName, spec.Host.Name)
		sCtx.Assert().Equal(exp.hostNamespace, spec.Host.Namespace)
	}
}

// assertHostBindingsByName Проверка списка host-binding, сопоставленных по имени
func assertHostBindingsByName(sCtx provider.StepCtx, hbs []*sgroupsv1.HostBinding, expected map[string]hostBindingExpectation) {
	for _, hb := range hbs {
		exp, ok := expected[hb.Metadata.Name]
		sCtx.Assert().True(ok, "unexpected host-binding: %s", hb.Metadata.Name)
		if !ok {
			continue
		}
		assertHostBinding(sCtx, hb, hb.Metadata.Name, exp)
	}
}

// Фикстуры для повторяющихся в разных кейсах host-binding (данные из тестового окружения)

var hb0Expectation = hostBindingExpectation{
	namespace: "namespace-0", uid: "1085d231-9a0f-4697-b864-c0a522181911",
	labels: map[string]string{"labels": "search"}, annotations: map[string]string{"search": "labels"},
	displayName: "host binding 0", comment: "for search by name/ns", description: "host binding for search",
	addressGroupName: "ag-0", addressGroupNamespace: "namespace-0", hostName: "host-0", hostNamespace: "namespace-0",
}

var hb1Expectation = hostBindingExpectation{
	namespace: "namespace-1", uid: "af7ee13a-ed99-4d35-bf68-ba76bb5a446c",
	labels: map[string]string{"labels": "ns"}, annotations: map[string]string{"search": "name"},
	displayName: "host binding 1", comment: "for search by name/ns", description: "host binding for search",
	addressGroupName: "ag-1", addressGroupNamespace: "namespace-1", hostName: "host-1", hostNamespace: "namespace-1",
}

var hb2Expectation = hostBindingExpectation{
	namespace: "namespace-2", uid: "5bd57191-447f-4c72-a295-4e8fe5800f20",
	labels: map[string]string{"search": "nameLabels"}, annotations: map[string]string{},
	displayName: "host binding 2", comment: "for search by labels", description: "host binding for search labels",
	addressGroupName: "ag-2", addressGroupNamespace: "namespace-2", hostName: "host-2", hostNamespace: "namespace-2",
}

var hb0And1 = map[string]hostBindingExpectation{"hb-0": hb0Expectation, "hb-1": hb1Expectation}
var hb0And2 = map[string]hostBindingExpectation{"hb-0": hb0Expectation, "hb-2": hb2Expectation}
