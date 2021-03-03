//GetUserWorkerHandler sends all users
func GetUserWorkerHandler(w http.ResponseWriter, r *http.Request) {
	var users = make(map[int]User)

	result := SelectAllDB(1)
	// defer result.Close()

	for result.Next() {
		var u User

		var x bool
		err := result.Scan(
			&u.UserWorkerID,
			&u.UserBossID,
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
			&u.UserImageURL,
			&x,
		)
		ErrorPrinter(err)

		users[u.UserWorkerID] = u
	}

	data, err := json.Marshal(users)
	ErrorPrinter(err)

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
	w.WriteHeader(http.StatusOK)
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
			&u.UserWorkerID,
			&u.UserBossID,
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
			&u.UserImageURL,
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

	var users []User

	for result.Next() {
		var userW User
		err := result.Scan(
			&userW.UserWorkerID,
			&userW.Name,
			&userW.LastName,
			&userW.Email,
			&userW.UserImageURL,
		)
		ErrorPrinter(err)
		fmt.Println(userW.UserWorkerID)
		users = append(users, userW)
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

//PutUserWorkerHandler updates the properties of user_boss
func PutUserWorkerHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	searchID := vars["id"]

	var userWorker User

	err := json.NewDecoder(r.Body).Decode(&userWorker)
	ErrorPrinter(err)

	fmt.Println(userWorker)

	UpdateDB(userWorker, searchID, 1)

	w.WriteHeader(http.StatusNoContent)
}

//DeleteUserWorkerHandler deactivates an user_boss
func DeleteUserWorkerHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userWorkerID := vars["id"]

	DeleteDB(userWorkerID, 1)

	w.WriteHeader(http.StatusNoContent)
}
