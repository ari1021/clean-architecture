package gateway

/*
gateway パッケージは，DB操作に対するアダプターです．
*/

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/ari1021/clean-architecture/entity"
	"github.com/ari1021/clean-architecture/usecase/port"
)

type UserRepository struct {
	conn *sql.DB
}

// NewUserRepository はUserRepositoryを返します．
func NewUserRepository(conn *sql.DB) port.UserRepository {
	return &UserRepository{
		conn: conn,
	}
}

// GetUserByID はDBからデータを取得します．
func (u *UserRepository) GetUserByID(userID string) (*entity.User, error) {
	conn := u.GetDBConn()
	row := conn.QueryRow("SELECT * FROM `user` WHERE id=?", userID)
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

// GetDBConn はconnectionを取得します．
func (u *UserRepository) GetDBConn() *sql.DB {
	return u.conn
}
