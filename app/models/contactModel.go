package models

import (
	"github.com/gocql/gocql"
	"log"
)

type ContactInfo struct {
	C_ID int
	Phone,Name ,Phone1 string
}
type DataView struct {
	Name ,Phone,Phone1 string
}
type RecData struct {
	Name ,Phone,Phone1 string
}

func (c ContactInfo)Session()* gocql.Session {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "data"
	cluster.Consistency = gocql.Quorum
	Session , err := cluster.CreateSession()
	if err !=nil{
		log.Fatal("error open cluster")
	}
	return Session
}
func (c ContactInfo) SaveInfo(id gocql.UUID,name ,ph1 ,ph2 string,se *gocql.Session) string {
	//fmt.Println( "account id :",id)
	var error1,error2 error
	//fmt.Println("phone2 : ",ph2)
	if ph2 == "" {
		//fmt.Println("case1 here !!!!!!!!!!!!")
		error1 = se.Query("INSERT into contact_info (account_id ,contact_id,name,phone,phone1) values (?,?,?,?,?)", id, gocql.TimeUUID(),name, ph1,"--").Exec()
		if error1 != nil {
			return "error1"
		}else {
			return "sucsessful1"
		}
	}else if ph2 !=""{
		//fmt.Println("case2 here !!!!!!!!!!!!!!!!!!!!!!!!!!!! ")
		error2 = se.Query("INSERT into  contact_info (account_id ,contact_id,name,phone,phone1) values(?,?,?,?,?)",id,gocql.TimeUUID(),name,ph1,ph2).Exec()
		if error2 != nil {
			return "error2"
		}else {
			return "sucsessful2"
		}
	}else
	{
		return " no data to add !1"
	}



}
func (d DataView) View(password string,sess *gocql.Session)[]RecData {

	var a []RecData
	var id gocql.UUID
	var x RecData
	//fmt.Println("user : ",username)
	//fmt.Println("pass : ",password)
	sess.Query("select id from account where password=?",password).Scan(&id)
	//fmt.Println("account id  : ",id)

	iter:=sess.Query("select name , phone, phone1 from contact_info where account_id =? ",id).Iter()
	for iter.Scan(&x.Name,&x.Phone,&x.Phone1){
		a=append(a,x)
	}
	//fmt.Println("UserContact : ",x)
	return a
}
