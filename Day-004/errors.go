package main

import (
	"errors"
	"fmt"
)

// Create a user struct
type User struct {
	name     string
	age      int
	isActive bool
}

// User repository
type UserRepository struct {
	users []User
}

// Function to create a user profile
func (r *UserRepository) createUser(name string, age int) (User, error) {
	if name == "" {
		return User{}, errors.New("name cannot be empty")
	}

	if age < 0 {
		return User{}, errors.New("age cannot be negative")
	}

	user := User{
		name:     name,
		age:      age,
		isActive: true,
	}

	r.users = append(r.users, user)

	return user, nil
}

// Function to get a user
func (r *UserRepository) getUser(name string) (User, error) {
	for _, user := range r.users {
		if user.name == name {
			return user, nil
		}
	}

	return User{}, errors.New("user not found")
}

// Function to get user profile
func getUserProfile(user User) (User, error) {
	if user.name == "" {
		return User{}, errors.New("invalid user")
	}

	return user, nil
}

func main() {

	repository := UserRepository{}

	// Create user
	user, err := repository.createUser("Darlene", 23)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("User created:", user)

	// Get user
	foundUser, err := repository.getUser("Darlene")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Get profile
	userProfile, err := getUserProfile(foundUser)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("User profile:", userProfile)
}
