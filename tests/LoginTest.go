package tests

import(
"github.com/revel/revel/testing"

	"net/url"
)


type LoginTest struct {
	testing.TestSuite
}
//-------------------------------------------------------------- test login.go --------------------------------------------------------------------------------------------------------------------
//----------test login method to render html file ----------
func (t *LoginTest) TestLoginPage() {
	t.Get("/login")
	t.AssertOk();
}
//-------------test submit method that send data to cluster and check the return value ----------------
func (t *LoginTest) TestSuccessfulSubmitLogin() {
	//test suc login
	t.PostForm("/submit_login",url.Values{
		"user":{"Gamila"},
		"pass":{"gamila43"},
	})
	t.AssertOk();
	t.AssertContains("Contact");
}
//-------------------------------------------- test fail login---------------------------------------------
func (t *LoginTest) TestFailSubmitLogin() {
	//test fail login
	t.PostForm("/submit_login",url.Values{
		"user":{"Gamila"},
		"pass":{"gamila123"},
	})
	t.AssertOk();
	t.AssertContains("Invalid username or password!");
}

