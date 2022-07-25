package main

import (
	"encoding/xml"

	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	c.XML(200, Person{FirstName: "Muhammad", LastName: "Abid"})

}

type Person struct {
	XMLName   xml.Name `xml:"person"`
	FirstName string   `xml:"firstName,attr"`
	LastName  string   `xml:"lastName,attr"`
}

func main() {
	router := gin.Default()
	router.GET("/", IndexHandler)
	router.Run()
}
