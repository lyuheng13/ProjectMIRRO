package handlers

import (
	"ProjectMIRRO/backend/gateway/model"
	"encoding/binary"
	"encoding/json"
	"fmt"

	"github.com/valyala/fasthttp"
)

//Receive userid to retrieve user info
func (ch *Context) UserGetHandler(ctx *fasthttp.RequestCtx) {
	req := &ctx.Request
	res := &ctx.Response
	userBody, err := ch.getUserBody(req)
	if err != nil {
		res.AppendBodyString("User not found")
		res.SetStatusCode(400)
	}

	userJson, err := json.Marshal(userBody)
	if err != nil {
		fmt.Println("Error marshalling body")
	}
	res.AppendBody(userJson)
	res.Header.SetContentType("application/json")
	res.SetStatusCode(200)
}

//Modify user informations to mysql
func (ch *Context) UserPatchHandler(ctx *fasthttp.RequestCtx) {
	req := &ctx.Request
	res := &ctx.Response
	_, err := ch.getUserBody(req)
	if err != nil {
		res.AppendBodyString("User not found")
		res.SetStatusCode(400)
	}
	data := make(map[string]interface{})
	err = json.Unmarshal(req.Body(), data)
	if err != nil {
		fmt.Println("Error unmarshalling")
	}

	//更新现有的数据
	//for k, v := range data {
	//	userBody[k] = data[v]
	//}
	//res.AppendBody(userJson)
	//res.Header.SetContentType("application/json")
	res.SetStatusCode(200)
}

//Delete the user
func (ch *Context) UserDeleteHandler(ctx *fasthttp.RequestCtx) {}

//Get an *fasthttp.Request as an input
//Return an *model.User struct
func (ch *Context) getUserBody(req *fasthttp.Request) (*model.User, error) {
	db := ch.UserSql

	useridByte := req.URI().QueryArgs().Peek("userid")
	userid := int64(binary.BigEndian.Uint64(useridByte))
	userBody, err := db.GetByID(userid)
	if err != nil {
		return nil, err
	}
	return userBody, nil
}
