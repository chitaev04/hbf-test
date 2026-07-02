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
