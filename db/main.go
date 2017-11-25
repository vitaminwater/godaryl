package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/vitaminwater/daryl_db"
)

type Test struct {
	Id       int    `db:"id" access:"s"`
	Slug     string `db:"slug" access:"i,s"`
	UserName string `db:"username" access:"i,u,s"`
}

func init() {
	daryl_db.Init()
}

func main() {
	test := Test{Slug: daryl_db.UUID(), UserName: "lolname"}
	err := daryl_db.Insert("primate", &test)
	if err != nil {
		log.Fatal(err)
	}
	test.UserName = "lolname2"
	daryl_db.Update("primate", "id", test)
}
