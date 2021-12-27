package db

import (
	"SimpleGoService/structs"
	"database/sql"
	"log"
)

func GetGroupIdByName(groupName string) int {
	row := db.QueryRow("SELECT id FROM `groups` WHERE name = ?", groupName)
	var groupID int
	error := row.Scan(&groupID)
	if error == sql.ErrNoRows {
		log.Println("Group with name", groupName, " not found!")
		return 0
	}
	return groupID
}

func GetUserIdByName(username string) int {
	row := db.QueryRow("SELECT id FROM `users` WHERE username = ?", username)
	var userId int
	error := row.Scan(&userId)
	if error == sql.ErrNoRows {
		log.Println("User with name", username, " not found!")
		return 0
	}
	return userId
}

func InsertUserInGroup(userId, groupId int) {
	query := "INSERT INTO user_in_group(user_id, group_id) VALUES (?, ?)"

	prepare, err := db.Prepare(query)
	if err != nil {
		return
	}
	prepare.Exec(userId, groupId)
}

func GetAllUsersInGroup(groupId int) []string {
	var users []string
	rows, err := db.Query("SELECT username FROM users JOIN user_in_group on users.id = user_in_group.user_id WHERE user_in_group.group_id = ?", groupId)
	if err != nil {
		log.Println("Cant select users in group ERROR!")
		return users
	}
	defer rows.Close()
	for rows.Next() {
		var user string
		if err := rows.Scan(&user); err != nil {
			log.Println("can't select some user")
		}
		users = append(users, user)
	}
	return users
}

func FindAllUsersInGroup(groupName string) []string {
	groupId := GetGroupIdByName(groupName)
	users := GetAllUsersInGroup(groupId)

	return users
}

func AddUserInGroup(userInGroup structs.UserInGroup) {
	groupId := GetGroupIdByName(userInGroup.GroupName)
	userId := GetUserIdByName(userInGroup.Username)

	InsertUserInGroup(userId, groupId)
}
