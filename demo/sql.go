package sql

import (
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID       int
	Username string
	Password string
}

func sql() {
	db, err := sql.Open("sqlite3", "./users.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create table if not exists
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, username TEXT, password TEXT)")
	if err != nil {
		log.Fatal(err)
	}

	// Insert a new user
	user := User{Username: "john", Password: "password123"}
	lastID, err := InsertUser(db, user)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("New user created with ID %d", lastID)

	// Update a user
	user.Password = "newpassword123"
	rowsAffected, err := UpdateUser(db, user)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%d rows updated", rowsAffected)

	// Delete a user
	rowsAffected, err = DeleteUser(db, user.ID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%d rows deleted", rowsAffected)

	// Select all users
	users, err := SelectAllUsers(db)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Users: %+v", users)
}

// InsertUser inserts a new user into the database and returns the last inserted ID.
func InsertUser(db *sql.DB, user User) (int64, error) {
	stmt, err := db.Prepare("INSERT INTO users(username, password) VALUES(?, ?)")
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(user.Username, user.Password)
	if err != nil {
		return 0, err
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lastID, nil
}

// UpdateUser updates an existing user in the database.
func UpdateUser(db *sql.DB, user User) (int64, error) {
	stmt, err := db.Prepare("UPDATE users SET password = ? WHERE id = ?")
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(user.Password, user.ID)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}

// DeleteUser deletes an existing user from the database.
func DeleteUser(db *sql.DB, id int) (int64, error) {
	stmt, err := db.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(id)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}

// SelectAllUsers selects all users from the database and returns a slice of User objects.
func SelectAllUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query("SELECT id, username, password FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Username, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}
