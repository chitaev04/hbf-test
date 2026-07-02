package suites

import (
	"github.com/ozontech/allure-go/pkg/framework/provider"

	sgroupsv1 "HBF-tests/pkg/api/sgroups/v1"

	common "github.com/PRO-Robotech/sgroups-proto/pkg/api/common"
)

// addressGroupExpectation Ожидаемые поля address group для проверки в ответах Upsert/List
type addressGroupExpectation struct {
	namespace     string
	uid           string
	labels        map[string]string
	annotations   map[string]string
	displayName   string
	comment       string
	description   string
	defaultAction string
	logs          bool
	trace         bool
}

// addressGroupLike Общий интерфейс ответов AddressGroup (Upsert) и AddressGroupExt (List)
type addressGroupLike interface {
	GetMetadata() *common.Metadata
	GetSpec() *sgroupsv1.AddressGroup_Spec
}

// assertAddressGroup Проверка полей одной address group
func assertAddressGroup(sCtx provider.StepCtx, ag addressGroupLike, name string, exp addressGroupExpectation) {
	md, spec := ag.GetMetadata(), ag.GetSpec()
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
	sCtx.Assert().Equal(exp.defaultAction, spec.DefaultAction.String())
	sCtx.Assert().Equal(exp.logs, spec.Logs)
	sCtx.Assert().Equal(exp.trace, spec.Trace)
}

// assertAddressGroupsByName Проверка списка address groups, сопоставленных по имени, с проверкой количества refs
func assertAddressGroupsByName(sCtx provider.StepCtx, ags []*sgroupsv1.AddressGroupResp_AddressGroupExt, expected map[string]addressGroupExpectation, refsCount map[string]int) {
	for _, ag := range ags {
		exp, ok := expected[ag.Metadata.Name]
		sCtx.Assert().True(ok, "unexpected address group: %s", ag.Metadata.Name)
		if !ok {
			continue
		}
		assertAddressGroup(sCtx, ag, ag.Metadata.Name, exp)
		if cnt, ok := refsCount[ag.Metadata.Name]; ok {
			sCtx.Assert().Len(ag.Refs, cnt)
		}
	}
}

// assertAddressGroupRefs Проверка refs одной address group по testData-подобной карте (ключ resType, либо resType_namespace при неоднозначности)
func assertAddressGroupRefs(sCtx provider.StepCtx, refs []*common.ResourceRef, expected map[string][2]string) {
	for _, r := range refs {
		key := r.ResType
		exp, ok := expected[key]
		if !ok {
			key = r.ResType + "_" + r.Namespace
			exp, ok = expected[key]
		}
		sCtx.Assert().True(ok, "unexpected ref: resType=%s name=%s namespace=%s", r.ResType, r.Name, r.Namespace)
		if !ok {
			continue
		}
		sCtx.Assert().Equal(exp[0], r.Name)
		sCtx.Assert().Equal(exp[1], r.Namespace)
	}
}

// Фикстуры для повторяющихся в разных кейсах address group (данные из тестового окружения)

var ag0Expectation = addressGroupExpectation{
	namespace: "namespace-0", uid: "d459b92b-0881-4166-9035-6b994ebdf798",
	labels: map[string]string{"labels": "search"}, annotations: map[string]string{"search": "labels"},
	displayName: "Address Group", comment: "for search by name/ns", description: "address group for search",
	defaultAction: "DENY", logs: true, trace: true,
}

var ag1Expectation = addressGroupExpectation{
	namespace: "namespace-1", uid: "37817691-8a8b-4344-8593-8a78ec3ce329",
	labels: map[string]string{"labels": "ns"}, annotations: map[string]string{"search": "name"},
	displayName: "Address Group 1", comment: "for search by name+ns", description: "address group for search",
	defaultAction: "ALLOW", logs: false, trace: false,
}

var ag2Expectation = addressGroupExpectation{
	namespace: "namespace-2", uid: "ce6a67d4-c2fa-484f-ae74-85fcd9e65a21",
	labels: map[string]string{"search": "nameLabels"}, annotations: map[string]string{},
	displayName: "Address Group 2", comment: "for search by labels", description: "address group for search labels",
	defaultAction: "DENY", logs: false, trace: true,
}

var ag0Refs = map[string][2]string{
	"Host":          {"host-0", "namespace-0"},
	"Network":       {"nw-0", "namespace-0"},
	"Service":       {"svc-0", "namespace-0"},
	"Ag2AgIcmpRule": {"rule-ag-ag-icmp-0", "namespace-0"},
}

var ag2Refs = map[string][2]string{
	"Host":                      {"host-2", "namespace-2"},
	"Network":                   {"nw-2", "namespace-2"},
	"Service":                   {"svc-2", "namespace-2"},
	"Ag2AgRule_namespace-1":     {"rule-ag-ag-2", "namespace-1"},
	"Ag2AgRule_namespace-2":     {"rule-ag-ag-2", "namespace-2"},
	"Ag2AgIcmpRule_namespace-1": {"rule-ag-ag-icmp-2", "namespace-1"},
	"Ag2AgIcmpRule_namespace-2": {"rule-ag-ag-icmp-2", "namespace-2"},
	"Ag2IcmpRule":               {"rule-ag-icmp-2", "namespace-2"},
	"Ag2CidrRule":               {"rule-ag-cidr-2", "namespace-2"},
	"Ag2CidrIcmpRule":           {"rule-ag-cidr-icmp-2", "namespace-2"},
	"Ag2FqdnRule":               {"rule-ag-fqdn-2", "namespace-2"},
}

// agRefsCount Количество refs у фикстур ag-0/ag-1/ag-2, используется в мульти-кейсах
var agRefsCount = map[string]int{"ag-0": 4, "ag-1": 14, "ag-2": 11}
