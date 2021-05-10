package handlers

import (
	"ProjectMIRRO/backend/gateway/model"
	"encoding/json"
	"fmt"
	"sort"

	"github.com/fatih/structs"
	"github.com/valyala/fasthttp"
)

var recommendationNum int

//Recommendation GET handler
//Take a type and a userid as inputs, return a list of recommendation products
func (ch *Context) RecoGetHandler(ctx *fasthttp.RequestCtx) {

	//Set the number of type to be recommended
	recommendationNum = 3

	//Retreive the initial info
	mongo := ch.MongoDBs.Mongo[1]
	recoType := ctx.QueryArgs().Peek("type")
	userId := string(ctx.QueryArgs().Peek("userid"))
	catergory := &model.Category{}

	//Get user type info
	collection := mongo.C(userId)
	err := collection.Find(userId).All(catergory)
	if err != nil {
		fmt.Errorf("Error retrieving user information")
		return
	}

	caterMap := structs.Map(catergory)

	//Caculate the recommendation
	highest, _ := getHighest(caterMap, recommendationNum)
	fmt.Println(highest)

	result := make(map[string]interface{})

	//Retrieve recommendations
	switch c := string(recoType); c {
	case "game":
		{
			fmt.Print("game") //TODO
		}
	case "music":
		{
			fmt.Print("music") //TODO
		}
	case "Vedio":
		{
			fmt.Print("Vedio")
			result = api.YoutebeSearch(highest)
		}
	case "book":
		{
			fmt.Print("book") //TODO
		}
	}

	recommendByte, _ := json.Marshal(result)
	ctx.Response.AppendBody(recommendByte)
	ctx.Response.SetStatusCode(200)
}

//Receive the rating of the specific type
//Implement the rating into the database
func (ch *Context) RatePostHandler(ctx *fasthttp.RequestCtx) {

	//Retrieve the initial settings
	mongo := ch.MongoDBs.Mongo[1]
	req := ctx.Request
	bodyData := req.Body()

	rate := model.Rate{}
	catergory := model.Category{}

	json.Unmarshal(bodyData, &rate)
	//ratedType := rate.Type
	userId := rate.ID

	//Get user type info
	collection := mongo.C(userId)
	err := collection.Find(userId).All(catergory)
	if err != nil {
		fmt.Errorf("Error retrieving user information")
		return
	}

	//Update rate
	if rate.Score != 0 {
		catergory[ratedType] += rate.Score - 2
	}

	ctx.Response.SetStatusCode(200)
}

//Get a slice of catergories with highest scores
func getHighest(catergory map[string]interface{}, num int) ([]string, error) {
	output := []string{}
	rate := []int{}

	//Retrieve the top three
	for _, value := range catergory {
		rate = append(rate, value.(int))
	}

	sort.Ints(rate)
	lowest := rate[num-1]
	for key, value := range catergory {
		if len(output) == num {
			break
		}

		if value.(int) >= lowest {
			output = append(output, key)
		}
	}

	return output, nil
}
