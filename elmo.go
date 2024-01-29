// class User
// class BwiCalculator
// => CalculateBwi(User User)
package main

import "fmt"

type User struct {
	Name   string
	Height float64
	Weight float64
}

type BwiCalculator struct{}

func (calc BwiCalculator) CalculateBwi(user User) float64 {
	bwi := user.Weight / (user.Height * user.Height)

	return bwi
}

func (calc BwiCalculator) AverageCalculateBwi(users []User) float64 {
	totalBwi := 0.0
	for _, user := range users {
		totalBwi += calc.CalculateBwi(user)
	}
	average := totalBwi / float64(len(users))
	return average
}

func main() {

	var users []User

	for {
		user := getUserInfo()
		users = append(users, user)

		fmt.Print("Want more? (yes/no): ")
		var choice string
		fmt.Scanln(&choice)
		if choice != "yes" {
			break
		}
	}
	calculated := BwiCalculator{}
	for _, user := range users {
		result := calculated.CalculateBwi(user)
		fmt.Printf("BWI for %s: %.2f\n", user.Name, result)
	}

	average := calculated.AverageCalculateBwi(users)
	fmt.Printf("Average BWI: %.2f\n", average)
}

func getUserInfo() User {
	var name string
	var height, weight float64

	fmt.Print("Enter Name: ")
	fmt.Scanln(&name)

	fmt.Print("Enter height: ")
	fmt.Scanln(&height)
	fmt.Print("Enter Weight: ")
	fmt.Scanln(&weight)

	return User{
		Name:   name,
		Height: height,
		Weight: weight,
	}
}
