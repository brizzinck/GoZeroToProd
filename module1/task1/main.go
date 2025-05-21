package main

import (
	"GoZeroToProd/module1/task1/dto"
	"fmt"
)

type Describer interface {
	Describe()
}

type Person interface {
	Describer
	GetName() string
	GetNickname() string
	ChangeAge(age uint8)
}

func main() {
	yaCode := dto.NewYaCode()
	skalse := dto.NewSkalse()
	oleksandr := dto.NewOleksandr()

	fmt.Printf("Skalse`s hobby - %s\n", skalse.GetHobby())

	skalse.ChangeHobby("Hard coding in `Rust` 25/8")

	fmt.Printf("Skalse`s hobby changed to - %s\n", skalse.GetHobby())

	oleksandr.CompleteTask()

	fmt.Printf("OleKsandr completed tasks - %d\n", oleksandr.GetCompletedTasks())

	oleksandr.UncompleteTask()

	fmt.Printf("OleKsandr completed tasks - %d\n", oleksandr.GetCompletedTasks())

	describePersons(
		[]Person{
			&yaCode,
			&skalse,
			&oleksandr,
		})
}

func describePersons(person []Person) {
	for i := range person {
		fmt.Printf("Now describing %s\n", person[i].GetName())

		person[i].Describe()
	}
}
