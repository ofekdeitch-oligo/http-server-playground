package main

import "testing"

func Test_Array1_Sum(t *testing.T) {
	arr := make([]int, 3)

	arr[0] = 1
	arr[1] = 2
	arr[2] = 3

	Test{t}.Expect(sum(arr)).ToEqual(6)
}

func Test_Array2_Sum(t *testing.T) {
	arr := make([]int, 3)

	arr[0] = 1
	arr[1] = -2
	arr[2] = 3

	Test{t}.Expect(sum(arr)).ToEqual(2)
}

func Test_Array1_Max(t *testing.T) {
	arr := make([]int, 4)

	arr[0] = 1
	arr[1] = 2
	arr[2] = 5
	arr[3] = 3

	Test{t}.Expect(max(arr)).ToEqual(5)
}

func Test_Array1_MaxBy(t *testing.T) {
	person1 := Person{name: "John", age: 20}
	person2 := Person{name: "Tod", age: 31}
	person3 := Person{name: "Sharon", age: 14}

	arr := make([]Person, 3)
	arr[0] = person1
	arr[1] = person2
	arr[2] = person3

	Test{t}.Expect(maxBy(arr, func(person Person) int {
		return person.age
	})).ToEqual(31)
}

func Test_Array1_SumBy(t *testing.T) {
	person1 := Person{name: "John", age: 20}
	person2 := Person{name: "Tod", age: 31}
	person3 := Person{name: "Sharon", age: 14}

	arr := make([]Person, 3)
	arr[0] = person1
	arr[1] = person2
	arr[2] = person3

	Test{t}.Expect(sumBy(arr, func(person Person) int {
		return person.age
	})).ToEqual(65)
}

func Test_Map1(t *testing.T) {
	person1 := Person{name: "John", age: 20}
	person2 := Person{name: "Tod", age: 31}
	person3 := Person{name: "Sharon", age: 14}

	arr := make([]Person, 3)
	arr[0] = person1
	arr[1] = person2
	arr[2] = person3

	names := Map(arr, func(p Person) string { return p.name })

	Test{t}.Expect(names[0]).ToEqual("John")
}

type Person struct {
	name string
	age  int
}
