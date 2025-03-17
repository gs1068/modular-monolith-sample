package dto

import (
	micropostpb "github.com/gs1068/modular-monolith-sample/gen/micropost/v1"
	userpb "github.com/gs1068/modular-monolith-sample/gen/user/v1"
)

type GetMicropostByUserIDOutput struct {
	User      UserOutput        `json:"user"`
	Micropost []MicropostOutput `json:"micropost"`
}

type MicropostOutput struct {
	ID      int64  `json:"id"`
	Content string `json:"content"`
}

func ConvertGetMicropostByUserIDResponse(pm []*micropostpb.Micropost, pu *userpb.User) *GetMicropostByUserIDOutput {
	var micropostOutputs []MicropostOutput
	for _, m := range pm {
		micropostOutputs = append(micropostOutputs, MicropostOutput{
			ID:      m.Id,
			Content: m.Content,
		})
	}

	return &GetMicropostByUserIDOutput{
		User: UserOutput{
			ID:    pu.Id,
			Email: pu.Email,
			Name:  pu.Name,
		},
		Micropost: micropostOutputs,
	}
}
