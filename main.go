package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var contents = []string{}

func main() {

	app := gin.Default()

	app.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "Hello")
	})

	app.POST("/notes", func(ctx *gin.Context) {
		var message Message
		ctx.BindJSON(&message)

		contents = append(contents, message.Content)
		index := len(contents) - 1

		ctx.JSON(201, Note{Index: index, Content: message.Content})
	})

	app.POST("/notes/:index", func(ctx *gin.Context) {
		index, _ := strconv.Atoi(ctx.Param("index"))
		var message Message
		ctx.BindJSON(&message)

		contents = append(contents[:index+1], contents[index:]...)
		contents[index] = message.Content
		ctx.JSON(201, Note{Index: index, Content: message.Content})
	})

	app.GET("/notes", func(ctx *gin.Context) {
		ctx.JSON(200, contents)
	})

	app.GET("/notes/:index", func(ctx *gin.Context) {
		index, e := strconv.Atoi(ctx.Param("index"))

		if e != nil {
			ctx.String(422, "provided index was not an integer")
		}

		ctx.JSON(200, Note{Index: index, Content: contents[index]})
	})

	app.PUT("/notes/:index", func(ctx *gin.Context) {
		index, _ := strconv.Atoi(ctx.Param("index"))
		var message Message
		ctx.BindJSON(&message)

		contents[index] = message.Content

		ctx.JSON(200, Note{Index: index, Content: message.Content})
	})

	app.DELETE("/notes/:index", func(ctx *gin.Context) {
		index, _ := strconv.Atoi(ctx.Param("index"))

		contents = append(contents[:index], contents[index+1:]...) // might be a bottle neck

		ctx.String(204, "")
	})

	app.POST("/documents", func(ctx *gin.Context) {
		var message Message
		ctx.BindJSON(&message)
		docId := uuid.New().String()
		fileName := fmt.Sprintf("%s.txt", docId)
		os.WriteFile(fileName, []byte(message.Content), 0777)
		ctx.JSON(201, DocumentInfo{DocId: docId})
	})

	app.GET("/documents/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")

		fileName := fmt.Sprintf("%s.txt", id)
		file, _ := os.ReadFile(fileName)
		document := Document{DocId: id, Content: string(file)}
		ctx.JSON(200, document)
	})

	app.GET("/math/:num1/:num2/:amount", func(ctx *gin.Context) {
		num1, _ := strconv.Atoi(ctx.Param("num1"))
		num2, _ := strconv.Atoi(ctx.Param("num2"))
		amount, _ := strconv.Atoi(ctx.Param("amount"))

		var x int
		for i := 0; i < amount; i++ {
			x = num1 * num2
		}

		ctx.String(200, fmt.Sprintf("%d", x))
	})

	app.GET("factorial/:num", func(ctx *gin.Context) {
		num, _ := strconv.Atoi(ctx.Param("num"))

		var factorial int = 1
		for i := 1; i < num+1; i++ {
			factorial *= i
			fmt.Println(factorial)
		}

		ctx.String(200, fmt.Sprintf("%d", factorial))
	})

	app.GET("coordinates/:amount", func(ctx *gin.Context) {
		amount, _ := strconv.Atoi(ctx.Param("amount"))
		coordinates := []Coordinate{}

		for i := 0; i < amount; i++ {
			lattitude := rand.Float32()*180 - 90
			longitude := rand.Float32()*360 - 180
			var NS_Hemisphere string
			var EW_Hemisphere string

			if lattitude > 0 {
				NS_Hemisphere = "North"
			} else {
				NS_Hemisphere = "South"
			}

			if longitude > 0 {
				EW_Hemisphere = "East"
			} else {
				EW_Hemisphere = "West"
			}

			coordinates = append(coordinates, Coordinate{Lattitude: lattitude, Longitude: longitude, EW_Hemisphere: EW_Hemisphere, NS_Hemisphere: NS_Hemisphere})
		}

		ctx.JSON(200, coordinates)
	})

	app.Run()
}
