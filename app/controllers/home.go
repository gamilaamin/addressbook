package controllers

import (
	"github.com/revel/revel"
	"github.com/gocql/gocql"
	"log"
)
 var Session *gocql.Session
type Home struct {
	*revel.Controller
}
  var err error
func ( h Home) Home() revel.Result  {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "data"
	cluster.Consistency = gocql.Quorum
	Session , err = cluster.CreateSession()
	if err !=nil{
		log.Fatal("error open cluster")
	}
	return  h.Render()
}

func ( h Home) Home1() revel.Result  {

	return  h.RenderTemplate("Home/Home1.html")
}



