package dto

import (
	"GoZeroToProd/module1/task1/constants"
	"GoZeroToProd/module1/task1/types"
	"fmt"
	"slices"
)

type Person struct {
	name     string
	nickname string
	age      uint8
	sex      types.Sex
}

func (p Person) GetName() string {
	return p.name
}

func (p Person) GetNickname() string {
	return p.nickname
}

func (p *Person) ChangeAge(age uint8) {
	p.age = age
}

func (p Person) Describe() {
	var sex string

	switch p.sex {
	case types.Male:
		sex = constants.Male
	case types.Female:
		sex = constants.Female
	default:
		sex = "Other"
	}

	fmt.Printf("I`m %s (%d years old, sex %s), my nickname - %s\n", p.name, p.age, sex, p.nickname)
}

type Skalse struct {
	Person
	hobby string
}

func NewSkalse() Skalse {
	return Skalse{
		Person: Person{
			name:     "Vlad",
			nickname: "Skalse",
			age:      constants.Eighteen,
			sex:      types.Male,
		},
		hobby: "Vibe coding in `Go` 25/8",
	}
}

func (s *Skalse) ChangeHobby(hobby string) {
	s.hobby = hobby
}

func (s *Skalse) GetHobby() string {
	return s.hobby
}

type YaCode struct {
	Person
	languageStack []string
}

func NewYaCode() YaCode {
	return YaCode{
		Person: Person{
			name:     "Dima",
			nickname: "YaCode",
			age:      constants.Twenty,
			sex:      types.Male,
		},
		languageStack: []string{"Go", "Python", "Rust"},
	}
}

func (YaCode) Describe() {
	fmt.Println("2147483647 out of 32 tasks completed, 9223372036854775807 remaining")
}

func (y *YaCode) GetLanguageStack() []string {
	return slices.Clone(y.languageStack)
}

type Oleksandr struct {
	Person
	completedTasks uint64
}

func NewOleksandr() Oleksandr {
	return Oleksandr{
		Person: Person{
			name:     "Sasha",
			nickname: "OleKsandr",
			age:      constants.Nineteen,
			sex:      types.Male,
		},
		completedTasks: constants.CountCompletedTasks,
	}
}

// Oleksandr only adds completed tasks; cannot decrease them
func (o Oleksandr) UncompleteTask() {
	o.completedTasks-- // nolint
}

func (o *Oleksandr) CompleteTask() {
	o.completedTasks++
}

func (o *Oleksandr) GetCompletedTasks() uint64 {
	return o.completedTasks
}
