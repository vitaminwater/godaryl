package model

import (
	"github.com/labstack/gommon/log"
	"github.com/vitaminwater/daryl/db"
	"github.com/vitaminwater/daryl/protodef"
	"golang.org/x/crypto/bcrypt"
)

type Daryl struct {
	Id       string `json:"id" db:"id" access:"s"`
	Name     string `json:"name" db:"name" access:"i,s"`
	Email    string `json:"email" db:"email" access:"i,s"`
	Password string `json:"password" db:"password" access:"i,u,s"`
}

func (d *Daryl) Insert() error {
	p, err := bcrypt.GenerateFromPassword([]byte(d.Password), -1)
	if err != nil {
		return err
	}
	d.Password = string(p)
	err = daryl_db.Insert("daryl", d)
	if err != nil {
		return err
	}
	return nil
}

func (d *Daryl) GetFromEmailAndPassword() error {
	p := d.Password
	log.Info(d)
	err := daryl_db.Get("daryl", "email", d, d)
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(d.Password), []byte(p))
	if err != nil {
		return err
	}
	return nil
}

func (d Daryl) ToProtodef() (*protodef.Daryl, error) {
	return &protodef.Daryl{
		Id:    d.Id,
		Name:  d.Name,
		Email: d.Email,
	}, nil
}

func NewDarylFromProtodef(d *protodef.Daryl) (Daryl, error) {
	return Daryl{
		Id:       d.Id,
		Name:     d.Name,
		Email:    d.Email,
		Password: d.Password,
	}, nil
}
