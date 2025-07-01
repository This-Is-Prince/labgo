package reflectpackage

import (
	"fmt"
	"reflect"
)

func InspectStruct(s interface{}) {
	val := reflect.ValueOf(s)

	fmt.Println(val.Kind())

	if val.Kind() != reflect.Struct {
		fmt.Println("Error: Provided interface is not a struct.")
		return
	}

	typ := val.Type()
	fmt.Printf("--- Inspecting a `%s` struct ---\n", typ.Name())

	for i := 0; i < val.NumField(); i++ {
		fieldValue := val.Field(i)
		fieldType := typ.Field(i)

		docTag := fieldType.Tag.Get("doc")
		if docTag == "" {
			docTag = "No documentation available."
		}

		if docTag := fieldType.Tag.Get("json"); docTag != "" {
			// fmt.Println("  JSON Tag:", docTag)
		}
		// Updated Printf to include the documentation from the tag.
		fmt.Printf("  Field: %-15s | Type: %-10s | Value: %-20v | Doc: %s\n",
			fieldType.Name,
			fieldValue.Kind(),
			fieldValue.Interface(),
			docTag)
	}

	fmt.Println("--------------------------------------------------")
}

type User struct {
	Id       int    `doc:"The unique identifier for the user." json:"id"`
	Name     string `doc:"User's full name."`
	IsActive bool   `doc:"Indicates if the user account is active."`
}

type Order struct {
	OrderId   string  `doc:"The unique ID for this order."`
	Amount    float64 `doc:"Total cost of the order."`
	Customer  User    `doc:"The user who placed the order."`
	IsShipped bool    // This field intentionally has no tag.
}

func Learn(run bool) {
	if !run {
		return
	}

	// name := "Goku"
	// age := uint16(9001)

	// fmt.Println("Type of name:", reflect.TypeOf(name))
	// fmt.Println("Type of age:", reflect.TypeOf(age))

	// value := reflect.ValueOf(name)
	// fmt.Println("The value is:", value)
	// fmt.Println("The type from the value is:", value.Kind())

	// value = reflect.ValueOf(age)
	// fmt.Println("The value is:", value)
	// fmt.Println("The type from the value is:", value.Kind())

	// city := "Agra"
	// fmt.Println("Original city:", city)

	// pointerToCity := reflect.ValueOf(&city)

	// element := pointerToCity.Elem()
	// element.SetString("New Delhi")

	// fmt.Println("Updated city:", city)

	user := User{
		Id:       101,
		Name:     "Vegeta",
		IsActive: true,
	}

	order := Order{
		OrderId: "XYZ-9001",
		Amount:  199.99,
		Customer: User{
			Id:       202,
			Name:     "Bulma",
			IsActive: true,
		},
		IsShipped: false,
	}

	InspectStruct(user)
	InspectStruct(order)
	fmt.Println("Inspecting structs with Pointers:")
	InspectStruct(&user)
	InspectStruct(&order)
}
