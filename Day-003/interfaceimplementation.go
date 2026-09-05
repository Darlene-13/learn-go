package main

import "fmt"

// Interface are implemented implicitly in go unlike in Java
func main() {

	var employee1 = fullTimeEmployee{
		name:   "Darlene Wendy",
		salary: 15000,
	}

	var employee2 = contractEmployee{
		name:         "Stacy Juma",
		hourlyPay:    12,
		hoursPerYear: 1000,
	}

	fmt.Println("==================================")
	fmt.Printf("Employee name: %s\n", employee1.getName())
	fmt.Printf("Employee salary: %v\n", employee1.getSalary())

	fmt.Println("===================================")
	fmt.Printf("Employee name: %s\n", employee2.getName())
	fmt.Printf("Employee salary: %v\n", employee2.getSalary())

	// Second option is to call a function that takes in any employee which is much quicker
	getEmployee(fullTimeEmployee{
		name:   "Darlene Wendy",
		salary: 50000,
	})

	getEmployee(contractEmployee{
		name:         "Stacy Juma",
		hoursPerYear: 1200,
		hourlyPay:    45,
	})

	getEmployee(partTimeEmployee{
		name:        "Joy Khalayi",
		hoursWorked: 1500,
		hourlyPay:   50,
	})
}

type employee interface {
	getName() string
	getSalary() int
}

type fullTimeEmployee struct {
	name   string
	salary int
}

type contractEmployee struct {
	name         string
	hourlyPay    int
	hoursPerYear int
}

type partTimeEmployee struct {
	name        string
	hourlyPay   int
	hoursWorked int
}

func (ce contractEmployee) getName() string {
	return ce.name
}

func (ce contractEmployee) getSalary() int {
	salary := ce.hourlyPay * ce.hoursPerYear
	return salary
}

func (ftemp fullTimeEmployee) getName() string {
	return ftemp.name
}

func (ftemp fullTimeEmployee) getSalary() int {
	return ftemp.salary
}

func (ptemp partTimeEmployee) getName() string {
	return ptemp.name
}

func (ptemp partTimeEmployee) getSalary() int {
	return ptemp.hourlyPay * ptemp.hoursWorked
}

// Function that takes in any employee
func getEmployee(emp employee) {
	fmt.Println("============================================")
	fmt.Printf("Employee name: %s\n", emp.getName())
	fmt.Printf("Employee salary: %v\n", emp.getSalary())
}
