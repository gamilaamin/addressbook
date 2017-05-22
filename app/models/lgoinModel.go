package models

import (
	//"fmt"
	"github.com/gocql/gocql"
	"log"
)

type Log struct {
	User,Password string
}

func(l Log) NewSession() *gocql.Session {

	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "data"
	cluster.Consistency = gocql.Quorum
	Session , err := cluster.CreateSession()
	if err !=nil{
		log.Fatal("error open cluster")
	}
	return Session
}
func (l Log) SubmitLogin(username string ,password string) string {
	var Count int
	Session:=l.NewSession()

	//fmt.Println("username : ",username)
	//fmt.Println("password : ",password)

	Session.Query("SELECT count (*) FROM account WHERE  password = ? and username = ? ", password,username).Scan(&Count)
	//fmt.Println("count :",Count)
	if Count == 1 {
		return "ok"
	} else {
		return "no"
	}

}
func (l Log) USerId( pass string)  gocql.UUID {
	var id  gocql.UUID
	Session1:=l.NewSession()
	Session1.Query("select id from account where password = ?",pass).Scan(&id)
	//fmt.Println("account id :",id)
	return id
}