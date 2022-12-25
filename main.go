package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

type RequestURI struct {
	ImageId string `uri:"id"`
}

type Response struct {
	Res string `json:"res"`
}

type UploadRequestBody struct {
	no    string
	image string
}

func GenerateUUID() uuid.UUID {
	gen := uuid.NewGen()
	result, err := gen.NewV4()
	if err != nil {
		panic(err)
	}
	return result
}

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("./html/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	r.GET("/images/:id", func(c *gin.Context) {
		reqURI := RequestURI{}
		if err := c.ShouldBindUri(&reqURI); err != nil {
			fmt.Print(err)
			c.JSON(http.StatusBadRequest, Response{Res: "잘못된 요청입니다!"})
			return
		}

		c.File("./files/" + reqURI.ImageId)
	})

	r.POST("/images/", func(c *gin.Context) {
		requestBody := UploadRequestBody{}

		fmt.Println(requestBody.no)
		fmt.Println(requestBody.image)

		if err := c.ShouldBind(&requestBody); err != nil {
			fmt.Print(err)
			c.JSON(http.StatusBadRequest, Response{Res: "잘못된 요청입니다!"})
			return
		}

		c.JSON(http.StatusOK, gin.H{})
	})

	r.Run()
}
