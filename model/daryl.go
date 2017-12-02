package model

import "github.com/vitaminwater/daryl/protodef"

type Daryl struct {
	Id       string `json:"id" db:"id" access:"s"`
	Name     string `json:"name" db:"name" access:"i,s"`
	Password string `json:"password" db:"password" access:"i,u,s"`
}

func (d Daryl) ToProtodef() (*protodef.Daryl, error) {
	return &protodef.Daryl{
		Id:   d.Id,
		Name: d.Name,
	}, nil
}

func NewDarylFromProtodef(d *protodef.Daryl) (Daryl, error) {
	return Daryl{
		Id:       d.Id,
		Name:     d.Name,
		Password: d.Password,
	}, nil
}
