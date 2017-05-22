package tests

import (
	"github.com/revel/revel/testing"
	"net/url"
)

type ContactTest struct   {
	testing.TestSuite
}

func ( t *ContactTest) TestSave()  {
	//test if condition in save function
	t.PostForm("/add_contact",url.Values{
		"name1":{"mohamed"},
		"Phone":{"03216549875"},
		"phone1":{""},
	})
	t.AssertOk();
	//test else in save function
	t.PostForm("/add_contact",url.Values{
		"name1":{"mohamed"},
		"Phone":{"03216549875"},
		"phone1":{"01198765432"},
	})
	t.AssertOk();
}
func (t *ContactTest) TestView() {
	t.Post("/View","",nil)
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")
}