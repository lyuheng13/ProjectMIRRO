package model

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// GetByType is an enumerate for GetBy* functions implemented
// by MySQLStore structs
type GetByType string

// MySQLStore is a user.Store backed by MySQL
type MySQLStore struct {
	Database *sql.DB
}

// NewMySQLStore constructs a new MySQLStore, and returns an error
// if there is a problem along the way.
func NewMySQLStore(dataSourceName string) (*MySQLStore, error) {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	return &MySQLStore{db}, nil
}

// getByProvidedType gets a specific user given the provided type.
// This requires the GetByType to be "unique" in the database.
func (ms *MySQLStore) getByProvidedType(t string, arg interface{}) (*User, error) {
	sel := string("select ID, Email, PassHash, UserName, FirstName, LastName, PhotoURL from Users where " + t + " = ?")

	rows, err := ms.Database.Query(sel, arg)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	user := &User{}

	// Should never have more than one row, so only grab one
	rows.Next()
	if err := rows.Scan(
		&user.ID,
		&user.Email,
		&user.PassHash,
		&user.UserName,
		&user.FirstName,
		&user.LastName,
		&user.PhotoURL); err != nil {
		return nil, err
	}
	return user, nil
}

//GetByID returns the User with the given ID
func (ms *MySQLStore) GetByID(id int64) (*User, error) {
	return ms.getByProvidedType("ID", id)
}

//GetByEmail returns the User with the given email
func (ms *MySQLStore) GetByEmail(email string) (*User, error) {
	return ms.getByProvidedType("Email", email)
}

//GetByUserName returns the User with the given Username
func (ms *MySQLStore) GetByUserName(username string) (*User, error) {
	return ms.getByProvidedType("UserName", username)
}

//Insert inserts the user into the database, and returns
//the newly-inserted User, complete with the DBMS-assigned ID
func (ms *MySQLStore) Insert(user *User) (*User, error) {
	ins := "insert into Users(Email, PassHash, UserName, FirstName, LastName, PhotoURL) values (?,?,?,?,?,?)"
	fmt.Println(ms)
	fmt.Println(ms.Database)
	res, err := ms.Database.Exec(ins, user.Email, user.PassHash, user.UserName,
		user.FirstName, user.LastName, user.PhotoURL)
	if err != nil {
		return nil, err
	}

	lid, lidErr := res.LastInsertId()
	if lidErr != nil {
		return nil, lidErr
	}

	user.ID = lid
	return user, nil
}

//Update applies UserUpdates to the given user ID
//and returns the newly-updated user
func (ms *MySQLStore) Update(id int64, updates *Updates) (*User, error) {
	// Assumes updates ALWAYS includes FirstName and LastName
	upd := "update Users set FirstName = ?, LastName = ? where ID = ?"
	res, err := ms.Database.Exec(upd, updates.FirstName, updates.LastName, id)
	if err != nil {
		return nil, err
	}

	rowsAffected, rowsAffectedErr := res.RowsAffected()
	if rowsAffectedErr != nil {
		return nil, rowsAffectedErr
	}

	if rowsAffected != 1 {
		return nil, errors.New("User not found")
	}

	// Get the user using GetByID
	return ms.GetByID(id)
}

//Delete deletes the user with the given ID
func (ms *MySQLStore) Delete(id int64) error {
	del := "delete from Users where ID = ?"
	res, err := ms.Database.Exec(del, id)
	if err != nil {
		return err
	}

	rowsAffected, rowsAffectedErr := res.RowsAffected()
	if rowsAffectedErr != nil {
		return rowsAffectedErr
	}

	if rowsAffected != 1 {
		return errors.New("User not found")
	}
	return nil
}

// LogSignIn logs sign-in attempts
func (ms *MySQLStore) LogSignIn(user *User, dateTime time.Time, clientIP string) error {
	insertQuery := "insert into LogInfo(UserID,LogTime,IpAddress) values(?,?,?)"
	_, err := ms.Database.Exec(insertQuery, user.ID, dateTime, clientIP)
	if err != nil {
		return fmt.Errorf("error logging sign in: %v", err)
	}
	return nil
}
