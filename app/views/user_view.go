package views

import "github.com/aviabird/go-echo-seed/app/models"

type userResponse struct {
	User struct {
		Email string `json:"email"`
	} `json:"user"`
}

type userListResponse struct {
	Users []*userResponse `json:"users"`
}

func UserListResponse(users []models.User) *userListResponse {
	ulr := new(userListResponse)
	ulr.Users = make([]*userResponse, 0)
	for _, user := range users {
		ur := new(userResponse)
		ur.User.Email = user.Email
	}
	return ulr
}

func UserResponse(u *models.User) *userResponse {
	r := new(userResponse)
	r.User.Email = u.Email
	return r
}
