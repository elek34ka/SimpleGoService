package db

import (
	"SimpleGoService/structs"
	"database/sql"
	"log"
)

func FindAllUsers() structs.UsersList {
	var (
		users     structs.UsersList
		rows, err = db.Query("SELECT * FROM users")
	)
	if err != nil {
		log.Println("can't select all users")
		return users
	}
	defer rows.Close()
	for rows.Next() {
		var user structs.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Debt); err != nil {
			log.Println("can't select some user")
		}
		users = append(users, user)
	}
	return users
}

func FindUserByName(username string) structs.User {
	var user structs.User
	row := db.QueryRow("SELECT * FROM users WHERE username = ?", username)
	err := row.Scan(&user.ID, &user.Username, &user.Debt)
	if err == sql.ErrNoRows {
		log.Println("User with name", username, " not found!")
		return structs.User{}
	}
	return user
}

func AddUser(username string) {
	if username == "" || FindGroupByName(username).ID > 0 {
		return
	}
	query := "INSERT INTO users(username, debt) VALUES (?, 0)"

	prepare, err := db.Prepare(query)
	if err != nil {
		return
	}
	prepare.Exec(username)

}
