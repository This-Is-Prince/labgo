package copierpackage

import (
	"fmt"
	"reflect"

	"github.com/jinzhu/copier"
)

type User struct {
	Name     string
	Age      int
	IsActive bool
}

type Manager struct {
	User
	Department string
}

func indirect(reflectValue reflect.Value) reflect.Value {
	fmt.Println(reflectValue.Kind())
	for reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}
	return reflectValue
}

/*
1. Talk with old friends
2. Share your problems with team people and with friends at least 2-4 friends
3. Do not hesitate to ask for help
4. Feeling ko feel karke express karoo, guilt trip pe mtt jaoo
5. Communicate with your team members and other people, share your feelings, and ask for help
*/

func copier1(toValue interface{}, fromValue interface{}) {
	indirect(reflect.ValueOf(toValue))
	indirect(reflect.ValueOf(fromValue))
}

func Learn(run bool) {
	if !run {
		return
	}

	fmt.Println("Learning copier package")

	user := User{
		Name:     "Goku",
		Age:      30,
		IsActive: true,
	}

	manager := Manager{}

	fmt.Println("Before copying:", user, manager)
	err := copier.Copy(&manager, &user)
	if err != nil {
		fmt.Println("Error copying:", err)
		return
	}
	fmt.Println("After copying:", user, manager)

	// i := interface{}("manager")
	// i := map[string]int{"key": 1}
	i := []int{1, 2, 3}

	copier1(manager, &user)
	copier1(i, &user)
}
