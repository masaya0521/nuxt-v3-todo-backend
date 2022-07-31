package gateway

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/masaya0521/go-clean-architecture-sample/entity"
	"github.com/masaya0521/go-clean-architecture-sample/usecase/port"
)

type UserRepository struct {
	conn *sql.DB
}

func NewUserRepository(conn *sql.DB) port.UserRepository{
	return &UserRepository{
		conn: conn,
	}
}

func (u *UserRepository) GetUserByID(ctx context.Context, userID string) (*entity.User, error) {
	conn := u.GetDBConn()
	row := conn.QueryRowContext(ctx, "SELECT * FROM `user` WHERE id=?", userID)
	user := entity.User{}
	err := row.Scan(&user.ID, &user.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("User Not Found. UserID = %s", userID)
		}
		log.Println(err)
		return nil, errors.New("Internal Server Error. adapter/gateway/GetUserByID")
	}
	return &user, nil
}

func (u *UserRepository) GetUser(ctx context.Context)(*[]entity.User, error) {
	users := []entity.User{}
	conn := u.GetDBConn()
	rows, err := conn.Query("SELECT * FROM `user`")
	if err != nil {
		return nil, fmt.Errorf("User Empty")
	}
	for rows.Next() {
		user := entity.User{}
		err = rows.Scan(&user.ID, &user.Name)
		if err != nil {
			return nil, fmt.Errorf("Scan failed")
		}
		users = append(users, user)
	}
	return &users, nil
}

// GetDBConn はconnectionを取得します．
func (u *UserRepository) GetDBConn() *sql.DB {
	return u.conn
}