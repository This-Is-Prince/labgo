package main

import "fmt"

type Employee struct {
	firstName, lastName string
	salary              int
	fullTime            bool
}

type Salary struct {
	basic     int
	insurance int
	allowance int
}

type EmployeeWithSalary struct {
	firstName, lastName string
	Salary
}

func main() {
	fmt.Println("Structures in GO! (structs)")

	/*
		type StructName struct {
			field1 fieldType1
			field2 fieldType2
			// ...
		}
	*/

	var ross Employee
	fmt.Println(ross)
	ross.firstName = "Ross"
	ross.lastName = "bing"
	ross.salary = 1200
	ross.fullTime = true

	fmt.Println("ross.firstName =", ross.firstName)
	fmt.Println("ross.lastName =", ross.lastName)
	fmt.Println("ross.salary =", ross.salary)
	fmt.Println("ross.fullTime =", ross.fullTime)

	ross1 := Employee{
		firstName: "Ross",
		lastName:  "Bing",
		salary:    1200,
	}
	fmt.Println(ross1)

	ross2 := Employee{"Ross", "bing", 1300, true}
	fmt.Println(ross2)

	monica := struct {
		firstName, lastName string
		salary              int
		fullTime            bool
	}{
		firstName: "Monica",
		lastName:  "Geller",
		salary:    1300,
		fullTime:  true,
	}

	fmt.Println(monica)

	ross3 := &Employee{
		firstName: "Ross",
		lastName:  "Bing",
		salary:    1200,
	}
	fmt.Println(ross3, ross3.firstName)

	kriti := EmployeeWithSalary{
		firstName: "Kriti",
		lastName:  "Ahuja",
		Salary:    Salary{1100, 50, 50},
	}
	fmt.Println(kriti.Salary.basic, kriti.basic)
}
