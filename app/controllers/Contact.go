package controllers

import (

	"addressBook/app/models"
	"github.com/revel/revel"
)


var ConInfo models.ContactInfo

var DataV  models.DataView


type Contact struct {
	*revel.Controller
}


func (s Contact) Save()revel.Result {

	s.Params.Bind(&ConInfo.Name,"name1")
	s.Params.Bind(&ConInfo.Phone,"Phone")
	s.Params.Bind(&ConInfo.Phone1,"phone1")

	//fmt.Println("userID: ",userInfo.USerId(userInfo.Password))

	response:=ConInfo.SaveInfo(userInfo.USerId(userInfo.Password),ConInfo.Name,ConInfo.Phone,ConInfo.Phone1,ConInfo.Session())

	//fmt.Println("response : ",response)

	if (response == "sucsessful1"||response=="sucsessful2") {
		return s.RenderTemplate("Contact/Contact.html")
	}else {
		return s.RenderHtml("<p margin-top:200px;><center> Error to add your contact!</center></p>")
	}
}

func (V Contact) View() revel.Result{
	var res []models.RecData
	//--------------------------------select all contact of this account ---------------------------------
	//fmt.Println(userInfo.User)
	//fmt.Println(userInfo.Password)

       res = DataV.View(userInfo.Password,userInfo.NewSession())

	//fmt.Println(res)
	X:=V.RenderJson(res)
	//fmt.Println("el x :",X)
        return  X
}


