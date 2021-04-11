package handlers

import (
	"ProjectMIRRO/backend/gateway/model"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/valyala/fasthttp"
	"golang.org/x/crypto/bcrypt"
)

// Singup handler:

//SignupHandler used to sign up for a new user
func (ch *Context) SignupHandler(ctx *fasthttp.RequestCtx) {
	req := &ctx.Request
	res := &ctx.Response

	//Unmarshalling the data
	body := req.Body()
	data := model.NewUser{}
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println("Error unmarshalling new user struct:", err.Error())
		return
	}

	//Check if the request is valid
	if err := ch.CheckValidSingup(&data); err != nil {
		res.AppendBodyString(err.Error())
		res.SetStatusCode(400)
	}

	if err := ch.InsertNewUser(&data); err != nil {
		fmt.Println("Error inserting new user")
		return
	}
	res.SetStatusCode(200)
}

//Check if the sign up post has enough info
func (ch *Context) CheckValidSingup(user *model.NewUser) error {
	//db := ch.UserSql
	//Check all the requested info are filled
	/*
		v := reflect.ValueOf(user)
		for i := 0; i < v.NumField(); i++ {
			if v.Field(i).IsNil() {
				return errors.New("Fields empty")
			}
		}
	*/

	//Check if the user email already exists
	/*
		email := user.Email
		_, err := db.GetByEmail(email)
		if err != nil {
			return errors.New("User already exists")
		}
	*/

	//Check passwordConf
	if user.Password != user.PasswordConf {
		return errors.New("Password not matched")
	}

	return nil
}

//Insert the newuser into the mysql db
func (ch *Context) InsertNewUser(newUser *model.NewUser) error {
	//Passing newUser struct to User struct
	user := model.User{}
	user.Email = newUser.Email
	user.UserName = newUser.UserName
	user.FirstName = newUser.FirstName
	user.LastName = newUser.LastName
	user.PhotoURL = ""
	user.PassHash = hashPass([]byte(newUser.Password))

	//Insert into db
	db := ch.UserSql
	_, err := db.Insert(&user)
	if err != nil {
		return err
	}
	return nil
}

//Generate password hash
func hashPass(pass []byte) []byte {
	hash, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Generate passhash error:", err.Error())
	}
	return hash
}

// Login hanlder:

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
