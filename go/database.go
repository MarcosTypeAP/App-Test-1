package main

import (
	"database/sql"
	"fmt"

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
		vars, values = CheckDataInsert(user)
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

	vars = CheckDataUpdate(user)

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
