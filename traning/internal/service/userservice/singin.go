package userservice

import (
	"fmt"
	"traning/internal/entity"
	"traning/internal/param"
)

func (s Service) SignUp(req param.UserRequest) (param.UserResponse, error) {

	user := entity.User{
		Name:     req.Name,
		Password: req.Password,
		Email:    req.Email,
	}

	resp, err := s.repo.CreateUser(user)
	if err != nil {
		return param.UserResponse{}, err
	}

	fmt.Println("SignInSignIn", resp)

	userResponse := param.UserResponse{
		UserInfo: param.UserInfo{
			Name:  resp.Name,
			ID:    resp.ID,
			Email: resp.Email,
		},
	}

	return userResponse, nil
}
