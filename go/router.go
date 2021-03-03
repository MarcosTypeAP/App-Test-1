package main

import (
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

//User is the boss user   // = not required
type User struct {
	UserBossID        int    `json:"user_boss_id,omitempty"`
	UserWorkerID      int    `json:"user_worker_id,omitempty"` //
	PhoneNumber       string `json:"phone_number,omitempty"`
	Country           string `json:"country,omitempty"`
	Nationality       string `json:"nationality,omitempty"`
	Dni               uint   `json:"dni,omitempty"`
	Email             string `json:"email,omitempty"`
	Name              string `json:"name,omitempty"`
	LastName          string `json:"last_name,omitempty"`
	DateOfBirth       string `json:"date_of_birth,omitempty"`
	Gender            string `json:"gender,omitempty"` //
	Address           string `json:"address,omitempty"`
	VehicleID         uint8  `json:"vehicle_id,omitempty"`        //
	Pc                uint8  `json:"pc,omitempty"`                //
	Children          uint8  `json:"children,omitempty"`          //
	MaritalStatusID   uint8  `json:"marital_status_id,omitempty"` //
	ProfessionID      uint   `json:"profession_id,omitempty"`
	PurchasedProperty uint8  `json:"purchased_property,omitempty"` //
	LivingPlaceID     uint8  `json:"living_place_id,omitempty"`    //
	Description       string `json:"description,omitempty"`        //
	Conduct           string `json:"conduct,omitempty"`            //
	Ideals            string `json:"ideals,omitempty"`             //
	UserImageURL      string `json:"user_image_url,omitempty"`     //
	Active            bool   `json:"active,omitempty"`             //
}

//Users is a map of UserBoss by id
var Users = make(map[int]User)

//RunServer runs the server
func RunServer(address string, port string) {
	router := mux.NewRouter().StrictSlash(false)
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	server := &http.Server{
		Addr:           address + ":" + port,
		Handler:        handlers.CORS(headers, methods, origins)(router),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	router.HandleFunc("/", HomeHandler).Methods("GET")

	router.HandleFunc("/api/userboss", GetUserBossHandler).Methods("GET")
	router.HandleFunc("/api/userboss/{id}", GetUserBossByIDHandler).Methods("GET")
	router.HandleFunc("/api/userboss/searchbyname/{searchString}", GetUserBossByNameHandler).Methods("GET")
	router.HandleFunc("/api/userboss", PostUserBossHandler).Methods("POST")
	router.HandleFunc("/api/userboss/{id}", PutUserBossHandler).Methods("PUT")
	router.HandleFunc("/api/userboss/{id}", DeleteUserBossHandler).Methods("DELETE")

	router.HandleFunc("/api/userworker", GetUserWorkerHandler).Methods("GET")
	router.HandleFunc("/api/userworker/{id}", GetUserWorkerByIDHandler).Methods("GET")
	router.HandleFunc("/api/userworker/searchbyname/{searchString}", GetUserWorkerByNameHandler).Methods("GET")
	router.HandleFunc("/api/userworker", PostUserWorkerHandler).Methods("POST")
	router.HandleFunc("/api/userworker/{id}", PutUserWorkerHandler).Methods("PUT")
	router.HandleFunc("/api/userworker/{id}", DeleteUserWorkerHandler).Methods("DELETE")

	server.ListenAndServe()
}

//HomeHandler is the home page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Home")
	http.ServeFile(w, r, "./assets/user-default-image.png")
}
