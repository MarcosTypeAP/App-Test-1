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
		vars = "(`dni`, `email`, `phone_number`, `name`, `last_name`, `date_of_birth`)"
		values = fmt.Sprintf("VALUES('%d', '%s', '%s', '%s', '%s', '%s')",
			user.Dni,
			user.Email,
			user.PhoneNumber,
			user.Name,
			user.LastName,
			user.DateOfBirth,
		)
	case 1: //worker
		table = "user_worker"
		vars = "(`phone_number`, `country`, `nationality`, `dni`, `email`, `name`, `last_name`, `date_of_birth`, `address`, `profession_id`"
		values = fmt.Sprintf("VALUES('%s', '%s', '%s', '%d', '%s', '%s', '%s', '%s', '%s', '%d'",
			user.PhoneNumber,
			user.Country,
			user.Nationality,
			user.Dni,
			user.Email,
			user.Name,
			user.LastName,
			user.DateOfBirth,
			user.Address,
			user.ProfessionID,
		)
		if user.Gender != "" {
			vars += ", `gender`"
			values += (", '" + user.Gender + "'")
		}
		if user.VehicleID != 0 {
			vars += ", `vehicle_id`"
			values += (", '" + strconv.Itoa(int(user.VehicleID)) + "'")
		}

		vars += ", `pc`"
		values += (", " + strconv.FormatBool(user.Pc))

		vars += ", `children`"
		values += (", " + strconv.FormatBool(user.Children))

		if user.MaritalStatusID != 0 {
			vars += ", `marital_status_id`"
			values += (", '" + strconv.Itoa(int(user.MaritalStatusID)) + "'")
		}

		vars += ", `purchased_property`"
		values += (", " + strconv.FormatBool(user.PurchasedProperty))

		if user.LivingPlaceID != 0 {
			vars += ", `living_place_id`"
			values += (", '" + strconv.Itoa(int(user.LivingPlaceID)) + "'")
		}
		if user.Description != "" {
			vars += ", `description`"
			values += (", '" + user.Description + "'")
		}
		if user.Conduct != "" {
			vars += ", `conduct`"
			values += (", '" + user.Conduct + "'")
		}
		if user.Ideals != "" {
			vars += ", `ideals`"
			values += (", '" + user.Ideals + "'")
		}
		vars += ")"
		values += ")"
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

	x := "%"
	result, err := db.Query(fmt.Sprintf("SELECT `name`, `last_name`, %s FROM %s WHERE CONCAT(`name`, ' ', `last_name`) LIKE '%s%s%s' ORDER BY `name`",
		userTypeID,
		table,
		x,
		searchString,
		x,
	))
	ErrorPrinter(err)

	return result
}

//SelectDB gets values from database
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
	result, err := db.Query(fmt.Sprintf("SELECT * FROM %s WHERE %s = %s",
		table,
		userTypeID,
		userID,
	))
	ErrorPrinter(err)
	return result
}

// //UpdateDB updates values from database
// func UpdateDB(p Person, id int64) bool {
// 	db := OpenConnectionDB()
// 	defer db.Close()

// 	if p.Name != "" {
// 		fmt.Println("name")
// 		p.Name = fmt.Sprintf("`name` = '%s'", p.Name)
// 	}
// 	if p.LastName != "" {
// 		fmt.Println("lastName")
// 		p.LastName = fmt.Sprintf("`last_name` = '%s'", p.LastName)
// 	}
// 	if p.LastName == "" && p.Name == "" {
// 		return true
// 	}

// 	coma := ""
// 	if p.Name != "" && p.LastName != "" {
// 		coma = ","
// 	}

// 	update, err := db.Query(
// 		fmt.Sprintf("UPDATE test_1 SET %s%s%s WHERE id = %d",
// 			p.Name,
// 			coma,
// 			p.LastName,
// 			id,
// 		),
// 	)
// 	ErrorPrinter(err)
// 	defer update.Close()
// 	return false
// }

// //DeleteDB deletes values from database
// func DeleteDB(id int64) {
// 	db := OpenConnectionDB()
// 	defer db.Close()
// 	delete, err := db.Query(fmt.Sprintf("DELETE FROM test_1 WHERE id = %d", id))
// 	ErrorPrinter(err)
// 	defer delete.Close()
// }
