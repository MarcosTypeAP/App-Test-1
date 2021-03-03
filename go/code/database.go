package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

//OpenConnectionDB opens the connection to the database
func OpenConnectionDB() *sql.DB {
	db, err := sql.Open("mysql", ConnectionString)
	ErrorPrinter(err)
	return db
}

//InsertDB inserts values to the database
func InsertDB(user User, userType int8) {
	var table string
	var vars string
	var values string
	switch userType {
	case 0: //boss
		table = "user_boss"
		vars = "(`dni`, `email`, `phone_number`, `name`, `last_name`, `date_of_birth`"
		values = fmt.Sprintf("VALUES('%d', '%s', '%s', '%s', '%s', '%s'",
			user.Dni,
			user.Email,
			user.PhoneNumber,
			user.Name,
			user.LastName,
			user.DateOfBirth,
		)
		if user.UserImageURL != "" {
			vars += ", `user_image_url`"
			values += (", '" + user.UserImageURL + "'")
		}
		vars += ")"
		values += ")"

	case 1: //worker
		table = "user_worker"

		vars, values = checkDataInsert(user)
	}

	db := OpenConnectionDB()
	defer db.Close()
	insert, err := db.Query(
		fmt.Sprintf("INSERT INTO %s %s %s",
			table,
			vars,
			values,
		),
	)
	ErrorPrinter(err)
	defer insert.Close()
}

//SelectAllDB selects all users
func SelectAllDB(userType int8) *sql.Rows {
	var table string
	switch userType {
	case 0: //boss
		table = "user_boss"
	case 1: //worker
		table = "user_worker"
	}

	db := OpenConnectionDB()
	defer db.Close()

	result, err := db.Query(fmt.Sprintf("SELECT * FROM %s ORDER BY `name`",
		table,
	))
	ErrorPrinter(err)

	return result
}

//SelectByNameDB select values by search
func SelectByNameDB(searchString string, userType int8) *sql.Rows {
	var table string
	var userTypeID string
	switch userType {
	case 0: //boss
		table = "user_boss"
		userTypeID = "`user_boss_id`"
	case 1: //worker
		table = "user_worker"
		userTypeID = "`user_worker_id`"
	}

	db := OpenConnectionDB()
	defer db.Close()

	result, err := db.Query(fmt.Sprintf("SELECT %s, `name`, `last_name`, `email`, `user_image_url` FROM %s WHERE CONCAT(`name`, ' ', `last_name`) LIKE '%s%s%s' AND `active` = 1 ORDER BY `name`",
		userTypeID,
		table,
		"%",
		searchString,
		"%",
	))
	ErrorPrinter(err)

	return result
}

//SelectDB selects values from database
func SelectDB(userID string, userType int8) *sql.Rows {
	var table string
	var userTypeID string
	switch userType {
	case 0: //boss
		table = "user_boss"
		userTypeID = "`user_boss_id`"
	case 1: //worker
		table = "user_worker"
		userTypeID = "`user_worker_id`"
	}
	db := OpenConnectionDB()
	defer db.Close()
	result, err := db.Query(fmt.Sprintf("SELECT * FROM %s WHERE %s = %s AND `active` = 1",
		table,
		userTypeID,
		userID,
	))
	ErrorPrinter(err)
	return result
}

//UpdateDB updates the properties of an user
func UpdateDB(user User, userID string, userType int8) {
	var table string
	var userTypeID string
	var vars string

	switch userType {
	case 0: //boss
		table = "user_boss"
		userTypeID = "user_boss_id"
	case 1: //worker
		table = "user_worker"
		userTypeID = "user_worker_id"
	}

	vars = checkDataUpdate(user)

	db := OpenConnectionDB()
	defer db.Close()

	update, err := db.Query(fmt.Sprintf("UPDATE %s SET %s WHERE `%s` = '%s'",
		table,
		vars,
		userTypeID,
		userID,
	))
	ErrorPrinter(err)

	defer update.Close()
}

//DeleteDB deactivates an user
func DeleteDB(userID string, userType int8) {
	var table string
	var userTypeID string
	switch userType {
	case 0: //boss
		table = "user_boss"
		userTypeID = "user_boss_id"
	case 1: //worker
		table = "user_worker"
		userTypeID = "user_worker_id"
	}

	db := OpenConnectionDB()
	defer db.Close()

	delete, err := db.Query(fmt.Sprintf("UPDATE %s SET `active` = 0 WHERE %s = %s",
		table,
		userTypeID,
		userID,
	))
	ErrorPrinter(err)
	defer delete.Close()
}

func checkDataUpdate(user User) string {
	var vars string
	if user.UserWorkerID != 0 {
		if vars == "" {
			vars = "`user_worker_id` = '" + strconv.Itoa(user.UserWorkerID) + "'"
		} else {
			vars += ", `user_worker_id` = '" + strconv.Itoa(user.UserWorkerID) + "'"
		}
	}
	if user.PhoneNumber != "" {
		if vars == "" {
			vars = "`phone_number` = '" + user.PhoneNumber + "'"
		} else {
			vars += ", `phone_number` = '" + user.PhoneNumber + "'"
		}
	}
	if user.Country != "" {
		if vars == "" {
			vars = "`country` = '" + user.Country + "'"
		} else {
			vars += ", `country` = '" + user.Country + "'"
		}
	}
	if user.Nationality != "" {
		if vars == "" {
			vars = "`nationality` = '" + user.Nationality + "'"
		} else {
			vars += ", `nationality` = '" + user.Nationality + "'"
		}
	}
	if user.Dni != 0 {
		if vars == "" {
			vars = "`dni` = '" + strconv.Itoa(int(user.Dni)) + "'"
		} else {
			vars += ", `dni` = '" + strconv.Itoa(int(user.Dni)) + "'"
		}
	}
	if user.Email != "" {
		if vars == "" {
			vars = "`email` = '" + user.Email + "'"
		} else {
			vars += ", `email` = '" + user.Email + "'"
		}
	}
	if user.Name != "" {
		if vars == "" {
			vars = "`name` = '" + user.Name + "'"
		} else {
			vars += ", `name` = '" + user.Name + "'"
		}
	}
	if user.LastName != "" {
		if vars == "" {
			vars = "`last_name` = '" + user.LastName + "'"
		} else {
			vars += ", `last_name` = '" + user.LastName + "'"
		}
	}
	if user.DateOfBirth != "" {
		if vars == "" {
			vars = "`date_of_birth` = '" + user.DateOfBirth + "'"
		} else {
			vars += ", `date_of_birth` = '" + user.DateOfBirth + "'"
		}
	}
	if user.Gender != "" {
		if vars == "" {
			vars = "`gender` = '" + user.Gender + "'"
		} else {
			vars += ", `gender` = '" + user.Gender + "'"
		}
	}
	if user.Address != "" {
		if vars == "" {
			vars = "`address` = '" + user.Address + "'"
		} else {
			vars += ", `address` = '" + user.Address + "'"
		}
	}
	if user.VehicleID != 0 {
		if vars == "" {
			vars = "`vehicle_id` = '" + strconv.Itoa(int(user.VehicleID)) + "'"
		} else {
			vars += ", `vehicle_id` = '" + strconv.Itoa(int(user.VehicleID)) + "'"
		}
	}
	if user.Pc != 0 {
		if vars == "" {
			vars = "`pc` = '" + strconv.Itoa(int(user.Pc)) + "'"
		} else {
			vars += ", `pc` = '" + strconv.Itoa(int(user.Pc)) + "'"
		}
	}
	if user.Children != 0 {
		if vars == "" {
			vars = "`children` = '" + strconv.Itoa(int(user.Children)) + "'"
		} else {
			vars += ", `children` = '" + strconv.Itoa(int(user.Children)) + "'"
		}
	}
	if user.MaritalStatusID != 0 {
		if vars == "" {
			vars = "`marital_status_id` = '" + strconv.Itoa(int(user.MaritalStatusID)) + "'"
		} else {
			vars += ", `marital_status_id` = '" + strconv.Itoa(int(user.MaritalStatusID)) + "'"
		}
	}
	if user.ProfessionID != 0 {
		if vars == "" {
			vars = "`profession_id` = '" + strconv.Itoa(int(user.ProfessionID)) + "'"
		} else {
			vars += ", `profession_id` = '" + strconv.Itoa(int(user.ProfessionID)) + "'"
		}
	}
	if user.PurchasedProperty != 0 {
		if vars == "" {
			vars = "`purchased_property` = '" + strconv.Itoa(int(user.PurchasedProperty)) + "'"
		} else {
			vars += ", `purchased_property` = '" + strconv.Itoa(int(user.PurchasedProperty)) + "'"
		}
	}
	if user.LivingPlaceID != 0 {
		if vars == "" {
			vars = "`living_place_id` = '" + strconv.Itoa(int(user.LivingPlaceID)) + "'"
		} else {
			vars += ", `living_place_id` = '" + strconv.Itoa(int(user.LivingPlaceID)) + "'"
		}
	}
	if user.Description != "" {
		if vars == "" {
			vars = "`description` = '" + user.Description + "'"
		} else {
			vars += ", `description` = '" + user.Description + "'"
		}
	}
	if user.Conduct != "" {
		if vars == "" {
			vars = "`conduct` = '" + user.Conduct + "'"
		} else {
			vars += ", `conduct` = '" + user.Conduct + "'"
		}
	}
	if user.Ideals != "" {
		if vars == "" {
			vars = "`ideals` = '" + user.Ideals + "'"
		} else {
			vars += ", `ideals` = '" + user.Ideals + "'"
		}
	}
	if user.UserImageURL != "" {
		if vars == "" {
			vars = "`user_image_url` = '" + user.UserImageURL + "'"
		} else {
			vars += ", `user_image_url` = '" + user.UserImageURL + "'"
		}
	}

	return vars
}

func checkDataInsert(user User) (string, string) {
	var vars string
	var values string

	if user.UserWorkerID != 0 {
		if vars == "" {
			vars = "(`user_worker_id`"
			values = "('" + strconv.Itoa(user.UserWorkerID) + "'"
		} else {
			vars += ", `user_worker_id`"
			values += ", '" + strconv.Itoa(user.UserWorkerID) + "'"
		}
	}
	if user.PhoneNumber != "" {
		if vars == "" {
			vars = "(`phone_number`"
			values = "('" + user.PhoneNumber + "'"
		} else {
			vars += ", `phone_number`"
			values += ", '" + user.PhoneNumber + "'"
		}
	}
	if user.Country != "" {
		if vars == "" {
			vars = "(`country`"
			values = "('" + user.Country + "'"
		} else {
			vars += ", `country`"
			values += ", '" + user.Country + "'"
		}
	}
	if user.Nationality != "" {
		if vars == "" {
			vars = "(`nationality`"
			values = "('" + user.Nationality + "'"
		} else {
			vars += ", `nationality`"
			values += ", '" + user.Nationality + "'"
		}
	}
	if user.Dni != 0 {
		if vars == "" {
			vars = "(`dni`"
			values = "('" + strconv.Itoa(int(user.Dni)) + "'"
		} else {
			vars += ", `dni`"
			values += ", '" + strconv.Itoa(int(user.Dni)) + "'"
		}
	}
	if user.Email != "" {
		if vars == "" {
			vars = "(`email`"
			values = "('" + user.Email + "'"
		} else {
			vars += ", `email`"
			values += ", '" + user.Email + "'"
		}
	}
	if user.Name != "" {
		if vars == "" {
			vars = "(`name`"
			values = "('" + user.Name + "'"
		} else {
			vars += ", `name`"
			values += ", '" + user.Name + "'"
		}
	}
	if user.LastName != "" {
		if vars == "" {
			vars = "(`last_name`"
			values = "('" + user.LastName + "'"
		} else {
			vars += ", `last_name`"
			values += ", '" + user.LastName + "'"
		}
	}
	if user.DateOfBirth != "" {
		if vars == "" {
			vars = "(`date_of_birth`"
			values = "('" + user.DateOfBirth + "'"
		} else {
			vars += ", `date_of_birth`"
			values += ", '" + user.DateOfBirth + "'"
		}
	}
	if user.Gender != "" {
		if vars == "" {
			vars = "(`gender`"
			values = "('" + user.Gender + "'"
		} else {
			vars += ", `gender`"
			values += ", '" + user.Gender + "'"
		}
	}
	if user.Address != "" {
		if vars == "" {
			vars = "(`address`"
			values = "('" + user.Address + "'"
		} else {
			vars += ", `address`"
			values += ", '" + user.Address + "'"
		}
	}
	if user.VehicleID != 0 {
		if vars == "" {
			vars = "(`vehicle_id`"
			values = "('" + strconv.Itoa(int(user.VehicleID)) + "'"
		} else {
			vars += ", `vehicle_id`"
			values += ", '" + strconv.Itoa(int(user.VehicleID)) + "'"
		}
	}
	if user.Pc != 0 {
		if vars == "" {
			vars = "(`pc`"
			values = "('" + strconv.Itoa(int(user.Pc)) + "'"
		} else {
			vars += ", `pc`"
			values += ", '" + strconv.Itoa(int(user.Pc)) + "'"
		}
	}
	if user.Children != 0 {
		if vars == "" {
			vars = "(`children`"
			values = "('" + strconv.Itoa(int(user.Children)) + "'"
		} else {
			vars += ", `children`"
			values += ", '" + strconv.Itoa(int(user.Children)) + "'"
		}
	}
	if user.MaritalStatusID != 0 {
		if vars == "" {
			vars = "(`marital_status_id`"
			values = "('" + strconv.Itoa(int(user.MaritalStatusID)) + "'"
		} else {
			vars += ", `marital_status_id`"
			values += ", '" + strconv.Itoa(int(user.MaritalStatusID)) + "'"
		}
	}
	if user.ProfessionID != 0 {
		if vars == "" {
			vars = "(`profession_id`"
			values = "('" + strconv.Itoa(int(user.ProfessionID)) + "'"
		} else {
			vars += ", `profession_id`"
			values += ", '" + strconv.Itoa(int(user.ProfessionID)) + "'"
		}
	}
	if user.PurchasedProperty != 0 {
		if vars == "" {
			vars = "(`purchased_property`"
			values = "('" + strconv.Itoa(int(user.PurchasedProperty)) + "'"
		} else {
			vars += ", `purchased_property`"
			values += ", '" + strconv.Itoa(int(user.PurchasedProperty)) + "'"
		}
	}
	if user.LivingPlaceID != 0 {
		if vars == "" {
			vars = "(`living_place_id`"
			values = "('" + strconv.Itoa(int(user.LivingPlaceID)) + "'"
		} else {
			vars += ", `living_place_id`"
			values += ", '" + strconv.Itoa(int(user.LivingPlaceID)) + "'"
		}
	}
	if user.Description != "" {
		if vars == "" {
			vars = "(`description`"
			values = "('" + user.Description + "'"
		} else {
			vars += ", `description`"
			values += ", '" + user.Description + "'"
		}
	}
	if user.Conduct != "" {
		if vars == "" {
			vars = "(`conduct`"
			values = "('" + user.Conduct + "'"
		} else {
			vars += ", `conduct`"
			values += ", '" + user.Conduct + "'"
		}
	}
	if user.Ideals != "" {
		if vars == "" {
			vars = "(`ideals`"
			values = "('" + user.Ideals + "'"
		} else {
			vars += ", `ideals`"
			values += ", '" + user.Ideals + "'"
		}
	}
	if user.UserImageURL != "" {
		if vars == "" {
			vars = "(`user_image_url`"
			values = "('" + user.UserImageURL + "'"
		} else {
			vars += ", `user_image_url`"
			values += ", '" + user.UserImageURL + "'"
		}
	}

	vars += ")"
	values += ")"

	return vars, values
}
