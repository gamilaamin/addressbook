package models

import (
	"github.com/gocql/gocql"
	//"fmt"
)

type AccountInfo struct {
	ID  gocql.UUID
	Fristname ,Lastname,Mail,Phone,Username,Pass1 string

}

func (a AccountInfo) AddInfo(id gocql.UUID,fristname,lastname,phone,mail,username,pass string) string {
	var err  error

	/*fmt.Println("id : ",id)
	fmt.Println("fristname : ",fristname)
	fmt.Println("lastname : ",lastname)
	fmt.Println("phone : ",phone)
	fmt.Println("mail : ",mail)
	fmt.Println("username : ",username)
	fmt.Println("pass : ",pass)*/

	if ( fristname != "" && lastname != "" && phone != "" && mail != "" && username != "" && pass != "" ) {
		cluster := gocql.NewCluster("127.0.0.1")
		cluster.Keyspace = "data"
		cluster.Consistency = gocql.Quorum
		Session, _ := cluster.CreateSession()

		//insert data in account table

		err = Session.Query("INSERT into account  (id,password,fristname,lastname,phone,mail,username) values(?,?,?,?,?,?,?)", id, pass, fristname, lastname, phone, mail, username).Exec()

	}
	if err != nil {
		return "errorInAdd"
	} else {
		return "SubmitAdd"

	}
}
