package tests

import ("net/url"
"github.com/revel/revel/testing"
)
type RegisterTest struct {
	testing.TestSuite
}
func (t *RegisterTest) Before() {
	println("Set up")
}

//--------------------------------------------------------------------------- test Register.go ------------------------------------------------------------------------------------------------------
//test  successful register
func (t *RegisterTest)TestRegistration(){
	t.PostForm("/AddInfo",url.Values{
		"firstname":{"lobna"},
		"lastname":{"ahmed"},
		"phone":{"01089631472"},
		"mail":{"lobna@yahoo.com"},
		"username":{"lobna"},
		"pass":{"44643214"},
	})
	t.AssertOk();
	t.AssertContains("My Contact");
}
//test fail register
/*func (t *RegisterTest)TestFailRegistration(){
	t.PostForm("/AddInfo",url.Values{
		"firstname":{""},
		"lastname":{""},
		"phone":{""},
		"mail":{""},
		"username":{"afrkosh"},
		"pass":{"afrkosh45"},
	})
	t.AssertOk();
	t.AssertContains("plese fill out this field.");
}*/

//------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

