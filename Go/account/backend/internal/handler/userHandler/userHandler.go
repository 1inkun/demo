package userhandler

import (
	"backend/internal/repository/user"
	"encoding/json"
	"log"
	"net/http"
	"golang.org/x/crypto/bcrypt"
)

type userHandler struct {
	queries *user.Queries
}

func NewUserHandler (queries *user.Queries) *userHandler {
	return &userHandler{queries : queries}
}

func hashPassword (password string) (password_hash string,err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password),14)
	if err != nil {
		return "",err
	}
	return string(bytes),err
}

func checkPasswordHash (password_hash string, user_password string) (err error) {
	err = bcrypt.CompareHashAndPassword([]byte(password_hash),[]byte(user_password))
	return err
}

func (h *userHandler) Register (w http.ResponseWriter, r *http.Request){
	user := new(user.User)
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err.Error())
		http.Error(w,err.Error(),http.StatusBadGateway)
		return
	}
	// log.Println(*user)
	var (
		nickname = user.Nickname
		username = user.Username
		password = user.Password_hash
	)
	if len(nickname) == 0 {
		nickname = "请设置昵称"
	}
	log.Println(nickname,username,password)
	_,err = h.queries.CheckIfExist(username)
	if err != nil {
		log.Println(err.Error())
		http.Error(w,err.Error(),http.StatusBadRequest)
		return
	}
	password_hash,err:= hashPassword(password)
	if err != nil {
		log.Println(err)
		http.Error(w,err.Error(),http.StatusBadGateway)
		return
	}
	log.Println(password_hash)
	err = h.queries.InsertUserInfo(nickname,username,password_hash)
	if err != nil {
		log.Println(err.Error())
		http.Error(w,err.Error(),http.StatusBadGateway)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *userHandler) Login (w http.ResponseWriter, r *http.Request) {
	user := new(user.User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		log.Println(err.Error())
		http.Error(w,err.Error(),http.StatusBadGateway)
		return
	}
	var (
		username = user.Username
		password = user.Password_hash
	)
	// log.Println(password)
	password_hash,err := h.queries.GetPasswdHash(username)
	if err != nil {
		log.Println(err.Error())
		http.Error(w,err.Error(),http.StatusBadGateway)
	}
	// log.Println(password_hash)
	err = checkPasswordHash(password_hash,password)
	if err != nil {
		log.Println(err.Error())
		http.Error(w,err.Error(),http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
