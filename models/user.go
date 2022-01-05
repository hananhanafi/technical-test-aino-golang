package models

import (
	"database/sql"
	"fmt"
	"net/http"
	"technical-test-aino-golang/database"
	"technical-test-aino-golang/helpers"

	"github.com/go-playground/validator/v10"
)

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func GetAllUser() (Response, error) {
	var user User
	var arrUser []User
	var res Response

	db := database.CreateCon()

	sqlStatement := "SELECT * FROM users"

	rows, err := db.Query(sqlStatement)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password)

		if err != nil {
			return res, err
		}
		arrUser = append(arrUser, user)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrUser

	return res, nil
}

func StoreUser(name string, email string, password string) (Response, error) {
	var res Response
	// validate input
	v := validator.New()
	newUser := User{
		Name:     name,
		Email:    email,
		Password: password,
	}
	err := v.Struct(newUser)
	if err != nil {
		return res, err
	}

	// hashing password input
	hashedPassword, err := helpers.HashPassword(password)
	if err != nil {
		return res, err
	}

	// connect to database
	db := database.CreateCon()

	// create sql statement
	sqlStatement := "INSERT users (name, email, password) VALUES (?, ?, ?)"

	// check the sql statement
	statement, err := db.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	// execute the sql statement
	result, err := statement.Exec(newUser.Name, newUser.Email, hashedPassword)
	if err != nil {
		return res, err
	}

	// get the last inserted user ID
	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	// set response success
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"last_inserted_id": lastInsertedId,
	}

	return res, nil
}

func Login(email, password string) (Response, error) {
	var res Response

	// check login user
	user, err := CheckLogin(email, password)
	if err != nil {
		return res, err
	}

	// create JWT
	token, err := CreateToken(email)
	if err != nil {
		return res, err
	}
	// delete password value for send it as a response
	user.Password = ""
	// set response
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]interface{}{
		"token": token,
		"user":  user,
	}

	return res, nil
}

func CheckLogin(email, password string) (User, error) {

	var user User

	// connect to database
	con := database.CreateCon()

	// create sql statement
	sqlStatement := "SELECT * FROM users WHERE email = ?"
	// check and execute the sql statement
	err := con.QueryRow(sqlStatement, email).Scan(
		&user.Id, &user.Name, &user.Email, &user.Password,
	)
	if err == sql.ErrNoRows {
		fmt.Println("Email not found")
		return user, err
	}
	if err != nil {
		fmt.Println("Query error")
		return user, err
	}

	// check password input, compares with hashed passsword in database
	match, err := helpers.CheckPasswordHash(password, user.Password)
	if !match {
		fmt.Println("Hash and password doesn't match.")
		return user, err
	}

	return user, nil
}
