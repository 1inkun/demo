package user


type User struct {
	Id int `json:"id,omitempty"`
	Nickname string	`json:"nickname,omitempty"`
	Username string	`json:"username"`
	Password_hash string `json:"password"`
	Create_at string `json:"create_at,omitempty"`
}
