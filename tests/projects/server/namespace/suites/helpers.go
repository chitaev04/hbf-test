package suites

import (
	"github.com/ozontech/allure-go/pkg/framework/provider"

	sgroupsv1 "HBF-tests/pkg/api/sgroups/v1"
)

// namespaceExpectation Ожидаемые поля namespace для проверки в ответах Upsert/List
type namespaceExpectation struct {
	labels      map[string]string
	annotations map[string]string
	displayName string
	comment     string
	description string
}

// assertNamespace Проверка полей одного namespace
func assertNamespace(sCtx provider.StepCtx, ns *sgroupsv1.Namespace, name string, exp namespaceExpectation) {
	sCtx.Assert().Equal(name, ns.Metadata.Name)
	sCtx.Assert().NotEmpty(ns.Metadata.Uid)
	sCtx.Assert().Equal(exp.labels, ns.Metadata.Labels)
	sCtx.Assert().Equal(exp.annotations, ns.Metadata.Annotations)
	sCtx.Assert().Equal(exp.displayName, ns.Spec.DisplayName)
	sCtx.Assert().Equal(exp.comment, ns.Spec.Comment)
	sCtx.Assert().Equal(exp.description, ns.Spec.Description)
}

// assertNamespacesByName Проверка списка namespaces, сопоставленных по имени
func assertNamespacesByName(sCtx provider.StepCtx, namespaces []*sgroupsv1.Namespace, expected map[string]namespaceExpectation) {
	for _, ns := range namespaces {
		exp, ok := expected[ns.Metadata.Name]
		sCtx.Assert().True(ok, "unexpected namespace name: %s", ns.Metadata.Name)
		if !ok {
			continue
		}
		assertNamespace(sCtx, ns, ns.Metadata.Name, exp)
	}
}
