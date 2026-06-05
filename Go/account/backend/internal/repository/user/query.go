package user
	
import (
	"database/sql"
	"errors"
	"log"
)

func (q *Queries) InsertUserInfo (nickname string,username string,password string) (err error){
	query := `INSERT INTO users (nickname,username,password_hash	)
	VALUES (?,?,?);`
	_,err = q.db.Exec(query,nickname,username,password)
	if err != nil {
		log.Println(err)
		return err
	}
	return err
}

func (q *Queries) CheckIfExist (username string) (status bool,err error) {
	query := `SELECT username FROM users WHERE username = ?;`
	row := q.db.QueryRow(query,username)
	if err = row.Scan() ; err == sql.ErrNoRows {
		return false,nil
	}
	err = errors.New("账户已存在!")
	return true,err
}

func (q *Queries) GetPasswdHash (username string) (password_hash string,err error){
	query := `SELECT password_hash FROM users WHERE username = ?;`
	row := q.db.QueryRow(query,username)
	err = row.Scan(&password_hash)
	if err == sql.ErrNoRows {
		return "",err
	}
	return password_hash,err
}
