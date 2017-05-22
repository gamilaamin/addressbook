// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tRegister struct {}
var Register tRegister


func (_ tRegister) Add(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Register.Add", args).Url
}

func (_ tRegister) AddInfo(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Register.AddInfo", args).Url
}


type tHome struct {}
var Home tHome


func (_ tHome) Home(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Home.Home", args).Url
}

func (_ tHome) Home1(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Home.Home1", args).Url
}


type tLogin struct {}
var Login tLogin


func (_ tLogin) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Login.Index", args).Url
}

func (_ tLogin) Submit(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Login.Submit", args).Url
}


type tContact struct {}
var Contact tContact


func (_ tContact) Save(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Contact.Save", args).Url
}

func (_ tContact) View(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Contact.View", args).Url
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).Url
}

func (_ tTestRunner) Suite(
		suite string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	return revel.MainRouter.Reverse("TestRunner.Suite", args).Url
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).Url
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).Url
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).Url
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).Url
}


