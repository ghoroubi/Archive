package main

import (
	"errors"
	"hermes"
	"net/http"
)

var MyApp *hermes.App

type MyModule struct {
	hermes.Module
}

var mymodule *MyModule

func main() {

	MyApp = hermes.NewApp("conf.yml")

	MyApp.InitLogs("")
	MyApp.Logger.Level = 6
	MyApp.Router.Use(hermes.LoggerMiddleware(MyApp.Logger, []string{}))
	hermes.DisableAuth()
	//escapes
	MyApp.Router.Use(hermes.AuthMiddleware([]string{"GET:/auth/logout", "POST:/auth/changePasswordByToken", "GET:/upload/image",
		"GET:/auth/activeUser", "POST:/auth/agents", "POST:/auth/login"}))
	//upload codes
	uploadPath := MyApp.Conf.GetString("Public.Upload")
	MyApp.Router.POST("/upload/image", hermes.UploadMiddleware("file", uploadPath, Uploader))
	MyApp.Router.StaticFS("/upload", http.Dir(uploadPath))
	//http routes
	MyApp.Mount(hermes.AuthorizationModule, "/auth")
	MyApp.Mount(mymodule, "/test")

	MyApp.Run()
}

//Define Collections Variables
var userCollection *hermes.Collection
var personCollection *hermes.Collection
var skillCollection *hermes.Collection
var needCollection *hermes.Collection

//Initializing Collections

func init() {
	mymodule = &MyModule{}
}
func (mm *MyModule) Init(app *hermes.App) error {
	var err error

	userCollection, err = hermes.NewDBCollection(&User{}, app.DataSrc)
	if err != nil {
		return err
	}
	userCollection.Conf().SetAuth("Create User", "Get User", "List User", "Update User", "Delete User", "")

	personCollection, err = hermes.NewDBCollection(&Person{}, app.DataSrc)
	if err != nil {
		return err
	}
	personCollection.Conf().SetAuth("Create Person", "Get Person", "List Person", "Update Person", "Delete Person", "")

	needCollection, err = hermes.NewDBCollection(&Need{}, app.DataSrc)
	if err != nil {
		return err
	}
	needCollection.Conf().SetAuth("Create Need", "Get Need", "List Need", "Update Need", "Delete Need", "")

	skillCollection, err = hermes.NewDBCollection(&Skill{}, app.DataSrc)
	if err != nil {
		return err
	}
	skillCollection.Conf().SetAuth("Create Skill", "Get Skill", "List Skill", "Update Skill", "Delete Skill", "")

	userController := mymodule.NewController(userCollection, "/users")
	mymodule.SetCrudRoutes(userController)

	personController := mymodule.NewController(personCollection, "/persons")
	mymodule.SetCrudRoutes(personController)

	skillController := mymodule.NewController(skillCollection, "/skills")
	mymodule.SetCrudRoutes(skillController)

	needController := mymodule.NewController(needCollection, "/needs")
	mymodule.SetCrudRoutes(needController)
	AddRolePermission()
	return nil
}
func AddRolePermission() {
	usr_roleid, err := hermes.AddRole("User")
	if err != nil {
		panic(err)
	}
	add_err := hermes.AddRolePermission(usr_roleid, "List User", "Get User")
	if add_err != nil {
		panic(add_err)
	}
	admin_roleid, err := hermes.AddRole("Admin")
	if err != nil {
		panic(err)
	}
	add_err = hermes.AddRolePermission(admin_roleid, "List User", "Get User")
	if add_err != nil {
		panic(err)
	}

	roleAgent := hermes.Role_Agent{}
	roleAgent.Agent_Id = 2
	roleAgent.Role_Id = usr_roleid
	_, err = hermes.RoleAgentColl.Create(hermes.SystemToken, nil, &roleAgent)
	if err != nil {
		panic(err)
	}

}
func Uploader(url, content_type string) (interface{}, error) {

	if content_type != "image/jpeg" {
		return "", errors.New("Image should be jpeg!")
	}

	return url, nil
}
