package controllers

import (
	"github.com/revel/revel"
	"addressBook/app/models"
	"github.com/gocql/gocql"

)

var AccInfo models.AccountInfo

type Register struct {
	*revel.Controller
}

func (c Register) Add() revel.Result  {
	return c.Render()
}

func (c Register) AddInfo() revel.Result  {
	AccInfo.ID = gocql.TimeUUID()
	c.Params.Bind(&AccInfo.Fristname,"firstname")
	c.Params.Bind(&AccInfo.Lastname,"lastname")
	c.Params.Bind(&AccInfo.Phone,"phone")
	c.Params.Bind(&AccInfo.Mail,"mail")
	c.Params.Bind(&AccInfo.Username,"username")
	c.Params.Bind(&AccInfo.Pass1,"pass")


       res:=AccInfo.AddInfo(AccInfo.ID,AccInfo.Fristname,AccInfo.Lastname,AccInfo.Phone,AccInfo.Mail,AccInfo.Username,AccInfo.Pass1)
	if res=="SubmitAdd"{

		return c.RenderTemplate("Register/add.html")
	}else{
		 return c.RenderHtml("<p margin-top:200px;><center> plese enter your info</center></p>")

	  }
}
