package structs

type User struct {
	ID       int
	Username string
	Debt     int
}

type UsersList []User
