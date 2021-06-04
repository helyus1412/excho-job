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

	c.JSON(200, responseObject)
}
