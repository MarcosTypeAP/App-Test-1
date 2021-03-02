package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Person is a person xd
type Person struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
}

//People is a people xd
var People = make(map[int]Person)

//User is the boss user   // = not required
type User struct {
	UserBossID        int    `json:"user_boss_id"`
	UserWorkerID      int    `json:"user_worker_id"` //
	PhoneNumber       string `json:"phone_number"`
	Country           string `json:"country"`
	Nationality       string `json:"nationality"`
	Dni               uint   `json:"dni"`
	Email             string `json:"email"`
	Name              string `json:"name"`
	LastName          string `json:"last_name"`
	DateOfBirth       string `json:"date_of_birth"`
	Gender            string `json:"gender"` //
	Address           string `json:"address"`
	VehicleID         uint8  `json:"vehicle_id"`        //
	Pc                bool   `json:"pc"`                //
	Children          bool   `json:"children"`          //
	MaritalStatusID   uint8  `json:"marital_status_id"` //
	ProfessionID      uint   `json:"profession_id"`
	PurchasedProperty bool   `json:"purchased_property"` //
	LivingPlaceID     uint8  `json:"living_place_id"`    //
	Description       string `json:"description"`        //
	Conduct           string `json:"conduct"`            //
	Ideals            string `json:"ideals"`             //
	Active            bool   `json:"active"`
}

//Users is a map of UserBoss by id
var Users = make(map[int]User)

//RunServer runs the server
func RunServer(address string, port string) {
	router := mux.NewRouter().StrictSlash(false)
	server := &http.Server{
		Addr:           address + ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	router.HandleFunc("/", HomeHandler).Methods("GET")

	router.HandleFunc("/api/userboss/{id}", GetUserBossByIDHandler).Methods("GET")
	router.HandleFunc("/api/userboss/searchbyname/{searchString}", GetUserBossByNameHandler).Methods("GET")
	router.HandleFunc("/api/userboss", PostUserBossHandler).Methods("POST")
	// router.HandleFunc("/api/userboss/{id}", PutUserBossHandler).Methods("PUT")
	router.HandleFunc("/api/userboss/{id}", DeleteUserBossHandler).Methods("DELETE")

	router.HandleFunc("/api/userworker/{id}", GetUserWorkerByIDHandler).Methods("GET")
	router.HandleFunc("/api/userworker/searchbyname/{searchString}", GetUserWorkerByNameHandler).Methods("GET")
	router.HandleFunc("/api/userworker", PostUserWorkerHandler).Methods("POST")
	router.HandleFunc("/api/userworker/{id}", DeleteUserWorkerHandler).Methods("DELETE")

	// router.HandleFunc("/api/persons", GetPersonHandler).Methods("GET")
	// router.HandleFunc("/api/persons", PostPersonHandler).Methods("POST")
	// router.HandleFunc("/api/persons/{id}", PutPersonHandler).Methods("PUT")
	// router.HandleFunc("/api/persons/{id}", DeletePersonHandler).Methods("DELETE")

	server.ListenAndServe()
}

//HomeHandler is the home page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home")
}

//GetUserBossByIDHandler sends the user searched by id
func GetUserBossByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	searchID := vars["id"]
	fmt.Println(searchID)
	result := SelectDB(searchID, 0)

	type userBoss struct {
		UserBossID  int    `json:"user_boss_id"`
		Dni         uint   `json:"dni"`
		Email       string `json:"email"`
		PhoneNumber string `json:"phone_number"`
		Name        string `json:"name"`
		LastName    string `json:"last_name"`
		DateOfBirth string `json:"date_of_birth"`
	}

	var u userBoss

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
			&x,
		)
		ErrorPrinter(err)
	}
	data, err := json.Marshal(u)
	ErrorPrinter(err)
	w.Header().Set("Access-Control-Allow-Origin", "*")
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

	type userBoss struct {
		Name     string `json:"name"`
		LastName string `json:"last_name"`
	}

	var users = make(map[int]userBoss)

	for result.Next() {
		var userB userBoss
		var userBossID int
		err := result.Scan(&userB.Name, &userB.LastName, &userBossID)
		ErrorPrinter(err)
		fmt.Println(userBossID)
		users[userBossID] = userB
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
// func PutUserBossHandler(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	searchID := vars["id"]

// 	var userBoss User

// 	err := json.NewDecoder(r.Body).Decode(&userBoss)
// 	ErrorPrinter(err)

// 	UpdateDB(userBoss, searchID, 0)

// 	w.WriteHeader(http.StatusNoContent)
// }

//DeleteUserBossHandler deactivates an user_boss
func DeleteUserBossHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userBossID := vars["id"]

	DeleteDB(userBossID, 0)

	w.WriteHeader(http.StatusNoContent)
}

//GetUserWorkerByIDHandler sends the user searched by id
func GetUserWorkerByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	searchID := vars["id"]
	fmt.Println(searchID)
	result := SelectDB(searchID, 1)

	var u User

	var x bool
	for result.Next() {
		err := result.Scan(
			&u.UserBossID,
			&u.UserWorkerID,
			&u.PhoneNumber,
			&u.Country,
			&u.Nationality,
			&u.Dni,
			&u.Email,
			&u.Name,
			&u.LastName,
			&u.DateOfBirth,
			&u.Gender,
			&u.Address,
			&u.VehicleID,
			&u.Pc,
			&u.Children,
			&u.MaritalStatusID,
			&u.ProfessionID,
			&u.PurchasedProperty,
			&u.LivingPlaceID,
			&u.Description,
			&u.Conduct,
			&u.Ideals,
			&x,
		)
		ErrorPrinter(err)
	}
	data, err := json.Marshal(u)
	ErrorPrinter(err)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
	w.WriteHeader(http.StatusOK)
}

//GetUserWorkerByNameHandler sends the search by name
func GetUserWorkerByNameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	searchString := vars["searchString"]

	result := SelectByNameDB(searchString, 1)
	defer result.Close()

	type userWorker struct {
		Name     string `json:"name"`
		LastName string `json:"last_name"`
	}

	var users = make(map[int]userWorker)

	for result.Next() {
		var userW userWorker
		var userWorkerID int
		err := result.Scan(&userW.Name, &userW.LastName, &userWorkerID)
		ErrorPrinter(err)
		fmt.Println(userWorkerID)
		users[userWorkerID] = userW
	}

	data, err := json.Marshal(users)
	ErrorPrinter(err)

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
	w.WriteHeader(http.StatusOK)
}

//PostUserWorkerHandler creates a new user_worker
func PostUserWorkerHandler(w http.ResponseWriter, r *http.Request) {
	var userWorker User
	err := json.NewDecoder(r.Body).Decode(&userWorker)
	ErrorPrinter(err)

	fmt.Println(userWorker)

	db := OpenConnectionDB()
	defer db.Close()

	InsertDB(userWorker, 1)
	w.WriteHeader(http.StatusCreated)
}

//DeleteUserWorkerHandler deactivates an user_boss
func DeleteUserWorkerHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userWorkerID := vars["id"]

	DeleteDB(userWorkerID, 1)

	w.WriteHeader(http.StatusNoContent)
}

////GetPersonHandler is GET
// func GetPersonHandler(w http.ResponseWriter, r *http.Request) {
// 	result := SelectDB()

// 	for result.Next() {
// 		var person Person
// 		var personID int
// 		err := result.Scan(&personID, &person.Name, &person.LastName)
// 		fmt.Println(personID)

// 		ErrorPrinter(err)

// 		People[personID] = person
// 	}
// 	result.Close()

// 	data, err := json.Marshal(People)
// 	ErrorPrinter(err)

// 	fmt.Println(People)

// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(data)
// 	w.WriteHeader(http.StatusOK)
// }

// //PostPersonHandler is POST
// func PostPersonHandler(w http.ResponseWriter, r *http.Request) {
// 	var people = make(map[string]Person)

// 	err := json.NewDecoder(r.Body).Decode(&people)
// 	ErrorPrinter(err)
// 	fmt.Println(people)

// 	for _, value := range people {
// 		InsertDB(value.Name, value.LastName)
// 	}
// 	w.WriteHeader(http.StatusCreated)
// }

// //PutPersonHandler is PUT
// func PutPersonHandler(w http.ResponseWriter, r *http.Request) {
// 	var person Person

// 	vars := mux.Vars(r)
// 	id, err := strconv.ParseInt(vars["id"], 10, 64)
// 	ErrorPrinter(err)

// 	err = json.NewDecoder(r.Body).Decode(&person)
// 	ErrorPrinter(err)
// 	fmt.Println(person)

// 	err1 := UpdateDB(person, id)
// 	if err1 {
// 		fmt.Println("No Content To Update")
// 	}

// 	w.WriteHeader(http.StatusNoContent)
// }

// //DeletePersonHandler is DELETE
// func DeletePersonHandler(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id, err := strconv.ParseInt(vars["id"], 10, 64) //string, base, bitSize
// 	ErrorPrinter(err)

// 	DeleteDB(id)
// 	w.WriteHeader(http.StatusNoContent)
// }
