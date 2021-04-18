package handlers

import (
	"ProjectMIRRO/backend/gateway/model"
	"encoding/json"
	"errors"

	"github.com/valyala/fasthttp"
	"golang.org/x/crypto/bcrypt"
)

//LoginHandler used to sign up for a new user
func (ch *Context) LoginHandler(ctx *fasthttp.RequestCtx) {
	req := &ctx.Request
	res := &ctx.Response
	body := req.Body()

	cred := model.Credentials{}
	json.Unmarshal(body, &cred)

	user, err := ch.CheckCred(&cred)
	if err != nil {
		res.AppendBodyString("User email/password not match")
		res.SetStatusCode(403)
	}

	userByte, _ := json.Marshal(user)
	res.AppendBody(userByte)
	res.SetStatusCode(200)
}

//Check if the credential is valid
//Return nil if credential is good
func (ch *Context) CheckCred(cred *model.Credentials) (*model.User, error) {
	db := ch.UserSql
	email := cred.Email
	passError := errors.New("User email/password not match")

	//Check if user exists
	user, err := db.GetByEmail(email)
	if err != nil {
		return nil, passError
	}

	//Check pass hash match
	passHash := user.PassHash
	err = bcrypt.CompareHashAndPassword(passHash, []byte(cred.Password))
	if err != nil {
		return nil, passError
	}

	return user, nil
}
