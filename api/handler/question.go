package handler

import (
	"encoding/json"
	"excho-job/entity"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func QuestionFetch(c *gin.Context) {
	response, err := http.Get("https://opentdb.com/api.php?amount=25&category=18&difficulty=medium&type=multiple")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject entity.Response
	json.Unmarshal(responseData, &responseObject)

	// fmt.Println(responseObject.ResponseCode)
	// fmt.Println(len(responseObject.Results))

	// for i := 0; i < len(responseObject.Results); i++ {
	// 	fmt.Println(responseObject.Results[i])
	// }

	c.JSON(200, responseObject)
}
