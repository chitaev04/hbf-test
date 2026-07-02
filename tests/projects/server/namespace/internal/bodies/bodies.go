package bodies

import (
	sgroupsv1 "HBF-tests/pkg/api/sgroups/v1"
	"HBF-tests/tests/common/foundation"
)

// CheckList Список всех namespace (пустой селектор → выборка всех)
func CheckList() *foundation.TestCaseBody {
	return &foundation.TestCaseBody{
		TestName: "Список всех namespace",
		Req: sgroupsv1.NamespaceReq_List{
			Selectors: []*sgroupsv1.NamespaceReq_Selector{
				{
					FieldSelector: &sgroupsv1.NamespaceReq_Selector_FieldSelector{},
				},
			},
		},
	}
}

// CheckListError Заведомо некорректный/пустой запрос для проверки ошибки валидации
func CheckListError() *foundation.TestCaseBody {
	return &foundation.TestCaseBody{
		TestName: "Список namespace с пустым телом (ожидается ошибка валидации)",
		Req:      sgroupsv1.NamespaceReq_List{},
	}
}

// ListAllEmptyBody Поиск всех по пустому телу (selectors не передан)
func ListAllEmptyBody() *foundation.TestCaseBody {
	return &foundation.TestCaseBody{
		TestName: "Поиск всех по пустому телу",
		Req:      sgroupsv1.NamespaceReq_List{},
	}
}

// ListAllEmptyFieldSelector Поиск всех по пустому объекту fieldSelector
func ListAllEmptyFieldSelector() *foundation.TestCaseBody {
	return &foundation.TestCaseBody{
		TestName: "Поиск всех по пустому объекту fieldSelector",
		Req: sgroupsv1.NamespaceReq_List{
			Selectors: []*sgroupsv1.NamespaceReq_Selector{
				{
					FieldSelector: &sgroupsv1.NamespaceReq_Selector_FieldSelector{},
				},
			},
		},
	}
}

// ListAllEmptySelectorsArray Поиск всех по пустому массиву selectors
func ListAllEmptySelectorsArray() *foundation.TestCaseBody {
	return &foundation.TestCaseBody{
		TestName: "Поиск всех по пустому массиву selectors",
		Req: sgroupsv1.NamespaceReq_List{
			Selectors: []*sgroupsv1.NamespaceReq_Selector{},
		},
	}
}

// ListByName Поиск одного namespace по name
func ListByName() *foundation.TestCaseBody {
	return &foundation.TestCaseBody{
		TestName: "Поиск одного namespace по name",
		Req: sgroupsv1.NamespaceReq_List{
			Selectors: []*sgroupsv1.NamespaceReq_Selector{
				{
					FieldSelector: &sgroupsv1.NamespaceReq_Selector_FieldSelector{Name: "namespace-1"},
				},
			},
		},
	}
}

// ListByNonExistentName Поиск namespace по несуществующему name
func ListByNonExistentName() *foundation.TestCaseBody {
	return &foundation.TestCaseBody{
		TestName: "Поиск namespace по несуществующему name",
		Req: sgroupsv1.NamespaceReq_List{
			Selectors: []*sgroupsv1.NamespaceReq_Selector{
				{
					FieldSelector: &sgroupsv1.NamespaceReq_Selector_FieldSelector{Name: "namespace-1111"},
				},
			},
		},
	}
}

// ListByTwoNames Поиск 2 namespaces по names
func ListByTwoNames() *foundation.TestCaseBody {
	return &foundation.TestCaseBody{
		TestName: "Поиск 2 namespaces по names",
		Req: sgroupsv1.NamespaceReq_List{
			Selectors: []*sgroupsv1.NamespaceReq_Selector{
				{FieldSelector: &sgroupsv1.NamespaceReq_Selector_FieldSelector{Name: "namespace-1"}},
				{FieldSelector: &sgroupsv1.NamespaceReq_Selector_FieldSelector{Name: "namespace-2"}},
			},
		},
	}
}

// ListByExistingAndNonExistentName Поиск namespace по сущ+несущ name
func ListByExistingAndNonExistentName() *foundation.TestCaseBody {
	return &foundation.TestCaseBody{
		TestName: "Поиск namespace по сущ+несущ name",
		Req: sgroupsv1.NamespaceReq_List{
			Selectors: []*sgroupsv1.NamespaceReq_Selector{
				{FieldSelector: &sgroupsv1.NamespaceReq_Selector_FieldSelector{Name: "namespace-1"}},
				{FieldSelector: &sgroupsv1.NamespaceReq_Selector_FieldSelector{Name: "namespace-222"}},
			},
		},
	}
}

// ListByLabels Поиск одного namespace по labels
func ListByLabels() *foundation.TestCaseBody {
	return &foundation.TestCaseBody{
		TestName: "Поиск одного namespace по labels",
		Req: sgroupsv1.NamespaceReq_List{
			Selectors: []*sgroupsv1.NamespaceReq_Selector{
				{LabelSelector: map[string]string{"search": "labels"}},
			},
		},
	}
}

// ListByNonExistentLabels Поиск namespace по несуществующим labels
func ListByNonExistentLabels() *foundation.TestCaseBody {
	return &foundation.TestCaseBody{
		TestName: "Поиск namespace по несуществующим labels",
		Req: sgroupsv1.NamespaceReq_List{
			Selectors: []*sgroupsv1.NamespaceReq_Selector{
				{LabelSelector: map[string]string{"not": "exist"}},
			},
		},
	}
}

// ListByTwoLabelSelectors Поиск 2 namespaces по labels
func ListByTwoLabelSelectors() *foundation.TestCaseBody {
	return &foundation.TestCaseBody{
		TestName: "Поиск 2 namespaces по labels",
		Req: sgroupsv1.NamespaceReq_List{
			Selectors: []*sgroupsv1.NamespaceReq_Selector{
				{LabelSelector: map[string]string{"search": "labels"}},
				{LabelSelector: map[string]string{"labels": "search"}},
			},
		},
	}
}

// ListByExistingAndNonExistentLabels Поиск namespaces по сущ+несущ labels
func ListByExistingAndNonExistentLabels() *foundation.TestCaseBody {
	return &foundation.TestCaseBody{
		TestName: "Поиск namespaces по сущ+несущ labels",
		Req: sgroupsv1.NamespaceReq_List{
			Selectors: []*sgroupsv1.NamespaceReq_Selector{
				{LabelSelector: map[string]string{"search": "labels"}},
				{LabelSelector: map[string]string{"not": "exist"}},
			},
		},
	}
}

// ListByNameAndLabels Поиск namespace по name+labels
func ListByNameAndLabels() *foundation.TestCaseBody {
	return &foundation.TestCaseBody{
		TestName: "Поиск namespace по name+labels",
		Req: sgroupsv1.NamespaceReq_List{
			Selectors: []*sgroupsv1.NamespaceReq_Selector{
				{
					FieldSelector: &sgroupsv1.NamespaceReq_Selector_FieldSelector{Name: "namespace-0"},
					LabelSelector: map[string]string{"search": "both"},
				},
			},
		},
	}
}

// ListByTwoNameAndLabels Поиск 2 namespaces по name+labels
func ListByTwoNameAndLabels() *foundation.TestCaseBody {
	return &foundation.TestCaseBody{
		TestName: "Поиск 2 namespaces по name+labels",
		Req: sgroupsv1.NamespaceReq_List{
			Selectors: []*sgroupsv1.NamespaceReq_Selector{
				{
					FieldSelector: &sgroupsv1.NamespaceReq_Selector_FieldSelector{Name: "namespace-0"},
					LabelSelector: map[string]string{"search": "both"},
				},
				{
					FieldSelector: &sgroupsv1.NamespaceReq_Selector_FieldSelector{Name: "namespace-1"},
					LabelSelector: map[string]string{"search": "labels"},
				},
			},
		},
	}
}

// ListByNameAndLabelsNameNotExist Поиск namespaces по name+labels. Name не существует
func ListByNameAndLabelsNameNotExist() *foundation.TestCaseBody {
	return &foundation.TestCaseBody{
		TestName: "Поиск namespaces по name+labels. Name не существует",
		Req: sgroupsv1.NamespaceReq_List{
			Selectors: []*sgroupsv1.NamespaceReq_Selector{
				{
					FieldSelector: &sgroupsv1.NamespaceReq_Selector_FieldSelector{Name: "n-0"},
					LabelSelector: map[string]string{"search": "both"},
				},
			},
		},
	}
}

// ListByNameAndLabelsLabelsNotExist Поиск namespaces по name+labels. labels не существуют
func ListByNameAndLabelsLabelsNotExist() *foundation.TestCaseBody {
	return &foundation.TestCaseBody{
		TestName: "Поиск namespaces по name+labels. labels не существуют",
		Req: sgroupsv1.NamespaceReq_List{
			Selectors: []*sgroupsv1.NamespaceReq_Selector{
				{
					FieldSelector: &sgroupsv1.NamespaceReq_Selector_FieldSelector{Name: "namespace-0"},
					LabelSelector: map[string]string{"not": "exist"},
				},
			},
		},
	}
}
