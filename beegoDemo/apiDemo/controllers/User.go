package controllers

import (
	"log"

	"apiDemo/models"
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (u *UserController) Post() {
	defer func() {
		err := recover()
		if err != nil {
			u.Ctx.WriteString("Error")
		}
	}()
	name := u.GetString("name")
	age, err := u.GetInt("age")
	if name == "" || err != nil {
		log.Fatalf("The Request Data Is Not Correct")
	} else {
		models.InsertUser(name, age)
	}
}
