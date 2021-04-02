package main

import (
	"fmt"
	"os"
)

type trainee struct {
	id     int
	name   string
	gender string
}

func newTrainee(id int, name, gender string) *trainee {
	return &trainee{
		id:     id,
		name:   name,
		gender: gender,
	}
}

type traineeManage struct {
	allTrainees []*trainee
}

func newTraineeManage() *traineeManage {
	return &traineeManage{
		allTrainees: make([]*trainee, 0, 100),
	}
}

func (t *traineeManage) addTrainee(newT *trainee) {
	t.allTrainees = append(t.allTrainees, newT)
}

func (t *traineeManage) modifyTrainee(newT *trainee) {
	for i, v := range t.allTrainees {
		if newT.id == v.id {
			t.allTrainees[i] = newT
			return
		}
	}

	fmt.Printf(" No Trainee with this id %d ------\n", newT.id)
}

func (t *traineeManage) showTrainee() {
	for _, v := range t.allTrainees {
		fmt.Printf("Trainee ID: %d Name: %t gender: %t\n", v.id, v.name, v.gender)
	}
}

func getInput() *trainee {
	var (
		id     int
		name   string
		gender string
	)
	fmt.Println("Trainee information as required!!!")
	fmt.Print("Enter Trainee ID:")
	fmt.Scanf("%d\n", &id)
	fmt.Print("Enter the Trainee name:")
	fmt.Scanf("%t\n", &name)
	fmt.Print("Enter the Trainee  gender:")
	fmt.Scanf("%t\n", &gender)

	t := newTrainee(id, name, gender)
	return t
}
func showMenu() {
	fmt.Println("Trainee information system")
	fmt.Println("1. Add Trainee")
	fmt.Println("2. Edit Trainee")
	fmt.Println("3. Show all Trainee")
	fmt.Println("4. Exit ")
}

func main() {

	tm := newTraineeManage()
	for {

		showMenu()

		var input int
		fmt.Println("Enter section you want to operate:")
		fmt.Scanf("%d\n", &input)
		fmt.Println("The user entered:", input)

		switch input {
		case 1:

			t := getInput()
			tm.addTrainee(t)
		case 2:

			t := getInput()
			tm.modifyTrainee(t)
		case 3:

			tm.showTrainee()
		case 4:

			os.Exit(0)
		default:
			fmt.Println("No input, exit!!")
			os.Exit(0)
		}
	}
}
