package data

import (
	"context"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name, omitempty"`
	LastName  string    `json:"last_name, omitempty"`
	Password  string    `json:"-"`
	Active    int       `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (user *User) GetAll() ([]*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)

	defer cancel()

	query := `SELECT * FROM users`

	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*User

	for rows.Next() {
		var user User

		err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.FirstName,
			&user.LastName,
			&user.Password,
			&user.Active,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

func (user *User) GetByEmail(email string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)

	defer cancel()

	query := `SELECT * FROM users WHERE email = $1`

	row := db.QueryRowContext(ctx, query, email)

	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Password,
		&user.Active,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (user *User) GetOne(id int) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)

	defer cancel()

	query := `SELECT * FROM users WHERE id = $1`

	row := db.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Password,
		&user.Active,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (user *User) Update() error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)

	defer cancel()

	query := `UPDATE users SET email = $1, first_name = $2, last_name = $3, password = $4, active = $5, updated_at = $6 WHERE id = $7`

	_, err := db.ExecContext(ctx, query, user.Email, user.FirstName, user.LastName, user.Password, user.Active, time.Now(), user.ID)

	if err != nil {
		return err
	}
	return nil
}

func (user *User) Delete() error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)

	defer cancel()

	query := `DELETE FROM users WHERE id = $1`

	_, err := db.ExecContext(ctx, query, user.ID)

	if err != nil {
		return err
	}
	return nil
}

func (user *User) DeleteById(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)

	defer cancel()

	query := `DELETE FROM users WHERE id = $1`

	_, err := db.ExecContext(ctx, query, id)

	if err != nil {
		return err
	}
	return nil
}

func (user *User) ResetPassword(password string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)

	defer cancel()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	query := `UPDATE users SET password = $1 WHERE id = $2`

	_, err = db.ExecContext(ctx, query, hashedPassword, user.ID)

	return err
}

func (user *User) PasswordMatches(plainText string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(plainText))

	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, err
		}
	}

	return true, nil
}
