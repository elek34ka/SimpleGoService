package db

import (
	"SimpleGoService/structs"
	"database/sql"
	"log"
)

func FindAllGroups() structs.GroupsList {
	var (
		groups    structs.GroupsList
		rows, err = db.Query("SELECT * FROM `groups`")
	)
	if err != nil {
		log.Println("can't select all groups")
		return groups
	}
	defer rows.Close()
	for rows.Next() {
		var group structs.Group
		if err := rows.Scan(&group.ID, &group.Name); err != nil {
			log.Println("can't select some user")
		}
		groups = append(groups, group)
	}
	return groups
}

func FindGroupByName(name string) structs.Group {
	var group structs.Group
	row := db.QueryRow("SELECT * FROM `groups` WHERE name = ?", name)
	err := row.Scan(&group.ID, &group.Name)
	if err == sql.ErrNoRows {
		log.Println("User with name", name, " not found!")
		return structs.Group{}
	}
	return group
}

func AddGroup(name string) {
	if name == "" || FindGroupByName(name).ID > 0 {
		return
	}
	query := "INSERT INTO `groups`(name) VALUES (?)"

	prepare, err := db.Prepare(query)
	if err != nil {
		return
	}
	prepare.Exec(name)

}
