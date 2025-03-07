package models

type RegisterUserResponse struct {
	AccessToken string `json:"accessToken"`
	User        *User  `json:"user"`
}

type User struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	FullName  string `json:"fullName"`
	AvatarURL string `json:"avatarUrl"`
	Bio       string `json:"bio"`
	CreateAt  string `json:"createAt"`
}
