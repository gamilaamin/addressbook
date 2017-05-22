package controllers

import (
	"github.com/revel/revel"
	"addressBook/app/models"
	//"fmt"
	"github.com/gocql/gocql"
)

type Login struct {
	*revel.Controller
}

func (l Login) Index()revel.Result  {
	 return l.Render()
}
var userInfo models.Log
var ID gocql.UUID
func (l Login) Submit() revel.Result {

	var response string
	l.Params.Bind(&userInfo.User,"user")
	l.Params.Bind(&userInfo.Password,"pass")

         response= userInfo.SubmitLogin(userInfo.User,userInfo.Password)

        // fmt.Println("response",response)
	//fmt.Println("data login : %v",userInfo)
	//fmt.Println("username : %v",userInfo.User)
	//fmt.Println("password : %v",userInfo.Password)

		if response == "ok" {
			ID=userInfo.USerId(userInfo.Password)
			//fmt.Println("userID : ",ID)
			return l.RenderTemplate("Contact/Contact.html")
		} else {
			return l.RenderHtml("<p margin-top:200px;><center>Invalid username or password!</center></p>")
		}

}

