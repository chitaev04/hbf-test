package suites

import (
	"github.com/ozontech/allure-go/pkg/framework/provider"

	sgroupsv1 "HBF-tests/pkg/api/sgroups/v1"

	common "github.com/PRO-Robotech/sgroups-proto/pkg/api/common"
)

// serviceBindingExpectation Ожидаемые поля service-binding для проверки в ответах Upsert/List
type serviceBindingExpectation struct {
	namespace             string
	uid                   string
	labels                map[string]string
	annotations           map[string]string
	displayName           string
	comment               string
	description           string
	addressGroupName      string
	addressGroupNamespace string
	serviceName           string
	serviceNamespace      string
}

// serviceBindingLike Общий интерфейс ответов ServiceBinding (Upsert и List используют один и тот же тип)
type serviceBindingLike interface {
	GetMetadata() *common.Metadata
	GetSpec() *sgroupsv1.ServiceBinding_Spec
}

// assertServiceBinding Проверка полей одного service-binding
func assertServiceBinding(sCtx provider.StepCtx, sb serviceBindingLike, name string, exp serviceBindingExpectation) {
	md, spec := sb.GetMetadata(), sb.GetSpec()
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
	if exp.serviceName != "" || exp.serviceNamespace != "" {
		sCtx.Require().NotNil(spec.Service)
		sCtx.Assert().Equal(exp.serviceName, spec.Service.Name)
		sCtx.Assert().Equal(exp.serviceNamespace, spec.Service.Namespace)
	}
}

// assertServiceBindingsByName Проверка списка service-binding, сопоставленных по имени
func assertServiceBindingsByName(sCtx provider.StepCtx, sbs []*sgroupsv1.ServiceBinding, expected map[string]serviceBindingExpectation) {
	for _, sb := range sbs {
		exp, ok := expected[sb.Metadata.Name]
		sCtx.Assert().True(ok, "unexpected service-binding: %s", sb.Metadata.Name)
		if !ok {
			continue
		}
		assertServiceBinding(sCtx, sb, sb.Metadata.Name, exp)
	}
}

// Фикстуры для повторяющихся в разных кейсах service-binding (данные из тестового окружения)

var sb0Expectation = serviceBindingExpectation{
	namespace: "namespace-0", uid: "b724836e-0bbd-4cec-afa7-a8b8d45d4d6f",
	labels: map[string]string{"labels": "search"}, annotations: map[string]string{"search": "labels"},
	displayName: "service binding 0", comment: "for search by name/ns", description: "service binding for search",
	addressGroupName: "ag-0", addressGroupNamespace: "namespace-0", serviceName: "svc-0", serviceNamespace: "namespace-0",
}

var sb1Expectation = serviceBindingExpectation{
	namespace: "namespace-1", uid: "7e13da25-2e49-4bae-bc5e-aa129870bc8a",
	labels: map[string]string{"labels": "ns"}, annotations: map[string]string{"search": "name"},
	displayName: "service binding 1", comment: "for search by name/ns", description: "service binding for search",
	addressGroupName: "ag-1", addressGroupNamespace: "namespace-1", serviceName: "svc-1", serviceNamespace: "namespace-1",
}

// sb2Expectation description="network binding for search labels" — копипаста-артефакт из NB в исходных данных Postman-коллекции, воспроизведён буквально
var sb2Expectation = serviceBindingExpectation{
	namespace: "namespace-2", uid: "07f31279-b179-434c-ac45-65087d7dea91",
	labels: map[string]string{"search": "nameLabels"}, annotations: map[string]string{},
	displayName: "service binding 2", comment: "for search by labels", description: "network binding for search labels",
	addressGroupName: "ag-2", addressGroupNamespace: "namespace-2", serviceName: "svc-2", serviceNamespace: "namespace-2",
}

var sb0And1 = map[string]serviceBindingExpectation{"sb-0": sb0Expectation, "sb-1": sb1Expectation}
var sb0And2 = map[string]serviceBindingExpectation{"sb-0": sb0Expectation, "sb-2": sb2Expectation}
