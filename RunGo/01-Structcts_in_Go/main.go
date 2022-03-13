package main

import (
	"fmt"
)

/* Declaring a struct type */
/*
type StructName struct{
	field1 fieldType1
	field2 fieldType2
}
*/

type Employee struct {
	firstName, lastName string
	salary              int
	fullTime            bool
}

func main() {
	fmt.Println("--------Structures in Go--------")

	// 1. Creating a struct
	/*
		a) struct keyword is built-in `type`.
		b) Employee is `struct type`
		c) ross is `struct`
	*/
	var ross Employee // Uninitialized struct, it contains default zero values of fields.
	// it is also called `empty struct`
	fmt.Println(ross)

	// 2. Getting and setting struct fields
	// we can get and set struct fields using . (dot) operator
	ross.firstName = "ross"
	ross.lastName = "Bing"
	ross.salary = 1200
	ross.fullTime = true

	fmt.Println("ross.firstName =", ross.firstName)
	fmt.Println("ross.lastName =", ross.lastName)
	fmt.Println("ross.salary =", ross.salary)
	fmt.Println("ross.fullTime =", ross.fullTime)

	// 3. Initializing a struct
	ross1 := Employee{
		firstName: "ross1",
		lastName:  "Bing1",
		fullTime:  true,
		salary:    1201,
	}
	ross2 := Employee{
		firstName: "ross2",
		lastName:  "Bing2",
	}
	ross3 := Employee{"ross3", "Bing3", 1203, true}
	fmt.Println(ross1, ross2, ross3)

	// 4. Anonymous struct
	monica := struct {
		firstName, lastName string
		salary              int
		fullTime            bool
	}{
		firstName: "Monica",
		lastName:  "Geller",
		salary:    2222,
	}
	monica1 := struct {
		firstName, lastName string
		salary              int
		fullTime            bool
	}{
		"Monica1",
		"Geller1",
		22221,
		false,
	}
	fmt.Println(monica, monica1)

	// 5. Pointer to a struct
	ross4 := &Employee{
		firstName: "ross4",
		lastName:  "Bing4",
		salary:    1204,
		fullTime:  true,
	}
	// Both are fine
	fmt.Println("firstName", ross4.firstName)
	fmt.Println("firstName", (*ross4).firstName)

	// 6.Anonymous fields
	type Data struct {
		string
		int
		bool
	}
	sample1 := Data{"Monday", 1200, true}
	sample1.bool = false
	fmt.Println(sample1.string, sample1.int, sample1.bool)
	/* Mixing some anonymous fields with named fields */
	/*
		type Employee struct {
			firstName, lastName string
			salary              int
			bool                // anonymous field
		}
	*/

	// 7. Nested struct
	type Salary struct {
		basic     int
		insurance int
		allowance int
	}

	type Employee1 struct {
		firstName, lastName string
		salary              Salary
		bool
	}
	me := Employee1{
		firstName: "Me",
		lastName:  "He", bool: true,
		salary: Salary{1100, 50, 50},
	}
	fmt.Println(me)
	fmt.Println("Me's basic salary", me.salary.basic)

	// 8. Promoted fields

	type Employee2 struct {
		firstName, lastName string
		Salary              // anonymous field
		bool                //anonymous field
	}
	me2 := Employee2{
		firstName: "Me2",
		lastName:  "He2", bool: true,
		Salary: Salary{1100, 50, 50},
	}
	fmt.Println(me2)
	fmt.Println("Me2's basic salary", me2.Salary.basic)
	/*
		But the cool thing about Go is that, when we use an anonymous nested struct,
		all the nested struct fields are automatically available on parent struct.
		This is called field promotion.
	*/
	me2.basic = 1200
	me2.insurance = 0
	me2.allowance = 0
	fmt.Println("Me2's basic salary", me2.basic)
	fmt.Println("Me2 is", me2)

	// 10. Exported fields
	/*
		me3 := Employee3{FirstName: "Me3", salary: 12}
		fmt.Println(me3)
		me4 := src.Employee4{FirstName: "Me4", salary: 14} // ERROR, unknown field salary in struct literal, salary field is not exported
		fmt.Println(me4)
	*/

	// 11. Function fields
	type FullNameType func(string, string) string

	type Employee5 struct {
		FirstName, LastName string
		FullName            FullNameType
	}
	e := Employee5{
		FirstName: "Employee",
		LastName:  "Singh",
		FullName: func(firstName, lastName string) string {
			return firstName + " " + lastName
		},
	}
	fmt.Println(e.FullName(e.FirstName, e.LastName))
}
