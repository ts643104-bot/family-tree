package main

import (
	"fmt"
	"reflect"
)


type User struct {
	Name  string
	Age   int
	Admin bool
}

func main() {

	u := User{Name: "Taha", Age: 25, Admin: true}

	
	v := reflect.ValueOf(u)
	t := reflect.TypeOf(u)

	fmt.Println("--- معلومات عامة ---")
	fmt.Printf("النوع (Type): %v\n", t)      // هيطبع main.User
	fmt.Printf("التصنيف (Kind): %v\n", t.Kind()) // هيطبع struct

	fmt.Println("\n--- فحص محتويات الـ Struct حقل بحقل ---")


	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i) 
		value := v.Field(i) 

		fmt.Printf("اسم الحقل: %-6s | نوعه: %-8v | قيمته: %v\n", 
			field.Name, field.Type, value)
	}
}