package queries

import (
	"errors"
	"fmt"

	"stray-dogs/app/models"
	"stray-dogs/pkg/utils"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

/*
This will be used to call the database
*/
type UserQueries struct {
	*sqlx.DB
}

type UserTokenResponse struct {
	Token string `json:"token"`
}

/*
Query to get a list of animals using pagination.
*/
func (q *UserQueries) Login(userBody *models.UserLoginStruct) (*string, error) {
	user := models.User{}

	password := userBody.Password
	email := userBody.Email

	query := `
		SELECT * 
			FROM "public"."user" 
		WHERE 
			email = $1
		LIMIT 1`
	err := q.Get(&user, query, email)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("password: Password did not match")
	}

	token, err := utils.GenerateNewAccessToken(user)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return &token, nil
}

func (q *UserQueries) CreateAdminUser(u *models.User) error {
	password := u.Password
	hash, _ := HashPassword(password)
	query := `INSERT INTO "public".user 
		(
			"email",
			"password", 
			"name"
		) 
			VALUES 
			(
				$1,
				$2, 
				$3
			)`

	_, err := q.Exec(query,
		u.Email,
		hash,
		u.Name,
	)

	if err != nil {
		return err
	}

	return nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
