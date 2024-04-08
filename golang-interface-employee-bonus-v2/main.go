package main

import "fmt"

type Employee interface {
	GetBonus() float64
}

// Junior Employee
type Junior struct {
	Name         string
	BaseSalary   int
	WorkingMonth int
}

func (j Junior) GetBonus() float64 {
	return (1 * float64(j.BaseSalary)) * (float64(j.WorkingMonth) / 12)
}

// Senior Employee
type Senior struct {
	Name            string
	BaseSalary      int
	WorkingMonth    int
	PerformanceRate float64
}

func (s Senior) GetBonus() float64 {
	return (2*float64(s.BaseSalary))*(float64(s.WorkingMonth)/12) + (s.PerformanceRate * float64(s.BaseSalary))
}

// Manager
type Manager struct {
	Name             string
	BaseSalary       int
	WorkingMonth     int
	PerformanceRate  float64
	BonusManagerRate float64
}

func (m Manager) GetBonus() float64 {
	return (2*float64(m.BaseSalary))*(float64(m.WorkingMonth)/12) + (m.PerformanceRate * float64(m.BaseSalary)) + (m.BonusManagerRate * float64(m.BaseSalary))
}

func EmployeeBonus(employee Employee) float64 {
	return employee.GetBonus() // TODO: replace this
}

func TotalEmployeeBonus(employees []Employee) float64 {
	sum := 0.0
	for _, v := range employees {
		sum += v.GetBonus()
	}
	return sum // TODO: replace this
}

func main() {

	junior := Junior{
		Name:         "Junior 1",
		BaseSalary:   100000,
		WorkingMonth: 12,
	}
	senior := Senior{
		Name:            "Senior 1",
		BaseSalary:      100000,
		WorkingMonth:    12,
		PerformanceRate: 0.5,
	}
	manager := Manager{
		Name:             "Manager 1",
		BaseSalary:       100000,
		WorkingMonth:     12,
		PerformanceRate:  0.5,
		BonusManagerRate: 0.5,
	}

	fmt.Println(EmployeeBonus(junior))
	fmt.Println(EmployeeBonus(senior))
	fmt.Println(EmployeeBonus(manager))

}
