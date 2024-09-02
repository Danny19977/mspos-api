package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/kgermando/mspos-api/controllers/area"
	"github.com/kgermando/mspos-api/controllers/asm"
	authentification "github.com/kgermando/mspos-api/controllers/auth"
	"github.com/kgermando/mspos-api/controllers/manager"
	poss "github.com/kgermando/mspos-api/controllers/pos"
	"github.com/kgermando/mspos-api/controllers/posform"
	"github.com/kgermando/mspos-api/controllers/province"
	"github.com/kgermando/mspos-api/controllers/sup"
	userlogs "github.com/kgermando/mspos-api/controllers/user_logs"
	"github.com/kgermando/mspos-api/controllers/userr"
	"github.com/kgermando/mspos-api/middlewares"
)

func Setup(app *fiber.App){

	api := app.Group("/api", logger.New())

	api.Post("/reset/:token", authentification.ResetPassword)

	// Authentification controller 
	auth := api.Group("/auth")
	auth.Post("/register", authentification.Register)
	auth.Post("/login", authentification.Login)
	auth.Get("/user", authentification.AuthUser)
	auth.Post("/forgot-password", authentification.Forgot)

	app.Use(middlewares.IsAuthenticated)

	auth.Patch("/users/profile", authentification.UpdateInfo)
	auth.Patch("/users/password", authentification.UpdatePassword)


	// Users controller
	user := api.Group("/users") 
	user.Get("/all", userr.GetUsers)
	user.Post("/create", userr.CreateUser)
	user.Get("/get/:id", userr.GetUser)
	user.Patch("/update/:id", userr.UpdateUser)
	user.Delete("/delete/:id", userr.DeleteUser)


	// Province controller
	prov := api.Group("/provinces") 
	prov.Get("/all", province.GetProvinces)
	prov.Post("/create", province.CreateProvince)
	prov.Get("/get/:id", province.GetProvince)
	prov.Patch("/update/:id", province.UpdateProvince)
	prov.Delete("/delete/:id", province.DeleteProvince)
 

	// Areas controller
	ar := api.Group("/areas") 
	ar.Get("/all", area.GetAreas)
	ar.Post("/create", area.CreateArea)
	ar.Get("/get/:id", area.GetArea)
	ar.Patch("/update/:id", area.UpdateArea)
	ar.Delete("/delete/:id", area.DeleteArea)


	// ASM controller
	as := api.Group("/asms") 
	as.Get("/all", asm.GetAsms)
	as.Post("/create", asm.CreateAsm)
	as.Get("/get/:id", asm.GetAsm)
	as.Patch("/update/:id", asm.UpdateAsm)
	as.Delete("/delete/:id", asm.DeleteAsm)


	// Manager controller
	ma := api.Group("/managers") 
	ma.Get("/all", manager.GetManagers)
	ma.Post("/create", manager.Createmanager)
	ma.Get("/get/:id", manager.GetManager)
	ma.Patch("/update/:id", manager.UpdateManager)
	ma.Delete("/delete/:id", manager.DeleteManager)


	// Posforms controller
	posf := api.Group("/posforms") 
	posf.Get("/all", posform.GetPosforms)
	posf.Post("/create", posform.CreatePosform)
	posf.Get("/get/:id", posform.GetPosform)
	posf.Patch("/update/:id", posform.UpdatePosform)
	posf.Delete("/delete/:id", posform.DeletePosform)
 

	// Sup controller
	su := api.Group("/sups") 
	su.Get("/all", sup.GetSups)
	su.Post("/create", sup.CreateSup)
	su.Get("/get/:id", sup.GetSup)
	su.Patch("/update/:id", sup.UpdateSup)
	su.Delete("/delete/:id", sup.DeleteSup)

	// Pos controller
	po := api.Group("/pos") 
	po.Get("/all", poss.GetPoss)
	po.Post("/create", poss.CreatePos)
	po.Get("/get/:id", poss.GetPos)
	po.Patch("/update/:id", poss.UpdatePos)
	po.Delete("/delete/:id", poss.DeletePos)
 
	
	// UserLogs controller
	userLog := api.Group("/users-logs") 
	userLog.Get("/all", userlogs.GetUserLogs)
	userLog.Post("/create", userlogs.CreateUserLog)
	userLog.Get("/get/:id", userlogs.GetUserLog)
	userLog.Patch("/update/:id", userlogs.UpdateUserLog)
	userLog.Delete("/delete/:id", userlogs.DeleteUserLog)
}