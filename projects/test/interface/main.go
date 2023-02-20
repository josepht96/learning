package main

import "fmt"

type doctorStruct struct {
	Name      string
	MedSchool string
}

type lawyerStruct struct {
	Name      string
	LawSchool string
}

func (d *doctorStruct) PrintName() string {
	fmt.Println(d.MedSchool)
	return "string"
}

func (l *lawyerStruct) PrintName() string {
	fmt.Println(l.LawSchool)
	return "string"
}

type employeeInterface interface {
	PrintName() string
}

func printNameInterface(employee employeeInterface) {
	_ = employee.PrintName()
}

func main() {
	doctor1 := &doctorStruct{
		Name:      "Joe T",
		MedSchool: "Tufts Med",
	}
	lawyer1 := &lawyerStruct{
		Name:      "Joe T",
		LawSchool: "Tufts Law",
	}
	printNameInterface(doctor1)
	printNameInterface(lawyer1)
}
