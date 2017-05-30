package main

import (
	"errors"
	//"fmt"
	"hermes"
	"net/http"
)

var App *hermes.App

func main() {

	App = hermes.NewApp("conf.yml")

	App.InitLogs("")
	App.Logger.Level = 6
	App.Router.Use(hermes.LoggerMiddleware(App.Logger, []string{}))

	hermes.DisableAuth()

	App.Router.Use(hermes.AuthMiddleware([]string{
		"GET:/auth/logout", "POST:/auth/changePasswordByToken", "GET:/upload/image",
		"GET:/auth/activeUser", "POST:/auth/agents", "POST:/auth/login"}))

	uploadPath := App.Conf.GetString("Public.Upload")
	App.Router.POST("/upload/image", hermes.UploadMiddleware("file", uploadPath, uploadImageMW))
	App.Router.StaticFS("/upload", http.Dir(uploadPath))

	App.Mount(hermes.AuthorizationModule, "/auth")

	App.Mount(mainMdl, "/test")

	App.Run()

}

var mainMdl *MMdl

//Collections Definition
var needCollection *hermes.Collection
var personCollection *hermes.Collection
var skillCollection *hermes.Collection
var userCollection *hermes.Collection

//Types Definition
type MMdl struct {
	hermes.Module
}

func init() {
	mainMdl = &MMdl{}
}

//Collections Initializing
func (mmodule *MMdl) Init(app *hermes.App) error {
	var err error
	needCollection, err = hermes.NewDBCollection(Need{}, app.DataSrc)
	if err != nil {
		return err
	}
	needCollection.Conf().SetAuth("Create Need", "Get Need", "List Need", "Update Need", "Delete Need", "")

	personCollection, err = hermes.NewDBCollection(Person{}, app.DataSrc)
	if err != nil {
		return err
	}
	personCollection.Conf().SetAuth("Create Person", "Get Person", "List Person", "Update Person", "Delete Person", "")

	skillCollection, err = hermes.NewDBCollection(Skill{}, app.DataSrc)
	if err != nil {
		return err
	}
	skillCollection.Conf().SetAuth("Create Skill", "Get Skill", "List Skill", "Update Skill", "Delete Skill", "")

	userCollection, err = hermes.NewDBCollection(User{}, app.DataSrc)
	if err != nil {
		return err
	}
	userCollection.Conf().SetAuth("Create User", "Get User", "List User", "Update User", "Delete User", "")
	//Controllers Definition
	userController := mainMdl.NewController(userCollection, "/users")
	mainMdl.SetCrudRoutes(userController)

	personController := mainMdl.NewController(personCollection, "/persons")
	mainMdl.SetCrudRoutes(personController)

	needController := mainMdl.NewController(needCollection, "/needs")
	mainMdl.SetCrudRoutes(needController)

	skillController := mainMdl.NewController(skillCollection, "/skills")
	mainMdl.SetCrudRoutes(skillController)

	addRolePersmission()
	return nil
} //End Of Init(app)

func addRolePersmission() {
	userRoleId, err := hermes.AddRole("User")
	if err != nil {
		panic(err)
	}
	errPrm := hermes.AddRolePermission(userRoleId, "List User", "Get User")
	if errPrm != nil {
		panic(errPrm)
	}
	adminRoleId, err := hermes.AddRole("Admin")
	if err != nil {
		panic(err)
	}
	errPrm = hermes.AddRolePermission(adminRoleId, "List User", "Get User")
	if errPrm != nil {
		panic(errPrm)
	}

	roleAgent := hermes.Role_Agent{}
	roleAgent.Agent_Id = 2
	roleAgent.Role_Id = userRoleId
	_, err = hermes.RoleAgentColl.Create(hermes.SystemToken, nil, &roleAgent)
	if err != nil {
		panic(err)
	}
}

func uploadImageMW(url, content_type string) (interface{}, error) {

	if content_type != "image/jpeg" {
		return "", errors.New("Image should be jpeg!")
	}

	return url, nil
}
