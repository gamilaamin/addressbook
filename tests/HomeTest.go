package tests

import (
	"github.com/revel/revel/testing"
)

type AppTest struct {
	testing.TestSuite
}

func (t *AppTest) Before() {
	println("Set up")
}
//--------------------------------------------------------------- test home.go --------------------------------------------------------------------------------------------------------------------
//-------- test Home method to render html file  -------
func (t *AppTest) TestRenderHomePage() {
	t.Get("/")
	t.AssertOk();
	t.Get("/home")
	t.AssertOk();
}
//--------- test Home method to check the host that us to make connection to cluster --------------
/*func (t *AppTest)TestHomeServer(){
	t.Get("/")
	actual:=t.Host()
	expected:="127.0.0.1:9000"
	t.AssertEqual(expected,actual)
}*/
//-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

func (t *AppTest) After() {
	println("Tear down")
}