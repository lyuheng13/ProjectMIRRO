package main

import (
	"ProjectMIRRO/backend/gateway/handlers"
	"ProjectMIRRO/backend/gateway/model"
	"fmt"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

//Create new context handler middleware
func newContextHandler() *handlers.Context {
	// Create sql connection
	dsn := "root:niupi123@tcp(localhost:3306)/niupiuser"
	db, err := model.NewMySQLStore(dsn)
	if err != nil {
		fmt.Println("Error creating new sql store:", err.Error())
		return nil
	}
	ch := &handlers.Context{}
	ch.UserSql = db

	return ch
}

func TestHandler(ctx *fasthttp.RequestCtx) {
	res := &ctx.Response
	res.AppendBodyString("Hello test")
	res.SetStatusCode(200)
}

func main() {

	// Generate new middleware
	ch := newContextHandler()

	//Routing
	router := fasthttprouter.New()
	router.POST("/user/signup", ch.SignupHandler)
	router.POST("/user/login", ch.LoginHandler)
	router.GET("/test", TestHandler)

	fmt.Println("Server Listening at port 8080...")
	if err := fasthttp.ListenAndServe("127.0.0.1:8080", router.Handler); err != nil {
		fmt.Println("Start http server failed:", err.Error())
	}
}
