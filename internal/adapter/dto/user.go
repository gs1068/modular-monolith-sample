package dto

import proto "github.com/gs1068/modular-monolith-sample/gen/user/v1"

type UserOutput struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

func ConvertUserResponse(pu *proto.User) *UserOutput {
	return &UserOutput{
		ID:    pu.Id,
		Email: pu.Email,
		Name:  pu.Name,
	}
}
