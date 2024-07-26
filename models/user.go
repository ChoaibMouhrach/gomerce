package models

type User struct {
	Id     int     `db:"id"`
	Name   *string `db:"name"`
	Email  string  `db:"email"`
	Phone  *string `db:"phone"`
	RoleId int     `db:"role_id"`
}

func UserFirstOrCreate(email string) (*User, error) {
	var user User

	if err := DB.Get(&user, "SELECT * FROM users WHERE email=$1", email); err != nil {
		var role Role

		if err := DB.Get(&role, "SELECT * FROM roles WHERE name=$1", MEMBER); err != nil {
			return nil, err
		}

		if _, err := DB.Exec("INSERT INTO users (email, role_id) VALUES ($1, $2)", email, role.Id); err != nil {
			return nil, err
		}

		DB.Get(&user, "SELECT * FROM users WHERE email=$1", email)
	}

	return &user, nil
}
