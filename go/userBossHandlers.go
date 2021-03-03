package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//GetUserBossHandler sends all users
func GetUserBossHandler(w http.ResponseWriter, r *http.Request) {
	var users = make(map[int]User)

	result := SelectAllDB(0)
	defer result.Close()

	for result.Next() {
		var u User

		var x bool
		err := result.Scan(
			&u.UserBossID,
			&u.Dni,
			&u.Email,
			&u.PhoneNumber,
			&u.Name,
			&u.LastName,
			&u.DateOfBirth,
			&u.UserImageURL,
			&x,
		)
		ErrorPrinter(err)

		users[u.UserBossID] = u
	}

	data, err := json.Marshal(users)
	ErrorPrinter(err)

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
	w.WriteHeader(http.StatusOK)
}

//GetUserBossByIDHandler sends the user searched by id
func GetUserBossByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	searchID := vars["id"]
	fmt.Println(searchID)
	result := SelectDB(searchID, 0)

	var u User

	var x bool
	for result.Next() {
		err := result.Scan(
			&u.UserBossID,
			&u.Dni,
			&u.Email,
			&u.PhoneNumber,
			&u.Name,
			&u.LastName,
			&u.DateOfBirth,
			&u.UserImageURL,
			&x,
		)
		ErrorPrinter(err)
	}
	data, err := json.Marshal(u)
	ErrorPrinter(err)
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
	w.WriteHeader(http.StatusOK)
}

//GetUserBossByNameHandler sends the search by name
func GetUserBossByNameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	searchString := vars["searchString"]

	result := SelectByNameDB(searchString, 0)
	defer result.Close()

	var users []User

	for result.Next() {
		var userB User
		err := result.Scan(
			&userB.UserBossID,
			&userB.Name,
			&userB.LastName,
			&userB.Email,
			&userB.UserImageURL,
		)
		ErrorPrinter(err)
		fmt.Println(userB.UserBossID)
		users = append(users, userB)
	}

	data, err := json.Marshal(users)
	ErrorPrinter(err)

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
	w.WriteHeader(http.StatusOK)
}

//PostUserBossHandler creates a new user_boss
func PostUserBossHandler(w http.ResponseWriter, r *http.Request) {
	var userBoss User
	err := json.NewDecoder(r.Body).Decode(&userBoss)
	ErrorPrinter(err)

	fmt.Println(userBoss)

	db := OpenConnectionDB()
	defer db.Close()

	InsertDB(userBoss, 0)

	w.WriteHeader(http.StatusCreated)
}

//PutUserBossHandler updates the properties of user_boss
func PutUserBossHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	searchID := vars["id"]

	var userBoss User

	err := json.NewDecoder(r.Body).Decode(&userBoss)
	ErrorPrinter(err)

	fmt.Println(userBoss)

	UpdateDB(userBoss, searchID, 0)

	w.WriteHeader(http.StatusNoContent)
}

//DeleteUserBossHandler deactivates an user_boss
func DeleteUserBossHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userBossID := vars["id"]

	DeleteDB(userBossID, 0)

	w.WriteHeader(http.StatusNoContent)
}
