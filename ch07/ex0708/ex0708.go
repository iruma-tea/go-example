package main

import "fmt"

type Employee struct { // 従業員
	Name string
	ID   string
}

func (e Employee) Description() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

type Manager struct { // マネージャー
	Employee            // 型のみ書く(埋め込みフィールド) Name、IDが加わる
	Reports  []Employee // 部下(報告の対象者) Employeeのスライス
}

func (m Manager) FindNewEmployees() []Employee {
	newEmployees := []Employee{
		Employee{
			"石田三成",
			"13112",
		},
		Employee{
			"徳川家康",
			"13155",
		},
	}
	return newEmployees
}

func main() {
	m := Manager{
		Employee: Employee{
			Name: "豊臣秀吉",
			ID:   "12345",
		},
		Reports: []Employee{},
	}
	fmt.Printf(m.ID)
	fmt.Println(m.Description())

	m.Reports = m.FindNewEmployees()
	fmt.Println(m.Employee)
	fmt.Println(m.Reports)
}
