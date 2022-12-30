package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"teamspeakapi/teamspeak"
)

func main() {
	client, err := teamspeak.Connect("arc.ms", "10011", "serveradmin", "TuuyoZwn", 1)
	if err != nil {
		log.Fatalln("Error Connect Func: ", err)
	}
	r := gin.New()
	// .../v1
	v1 := r.Group("/v1")
	// .../v1/api
	api := v1.Group("/api")
	// .../v1/api/ts
	ts := api.Group("/ts")
	ts.GET("/clientlist", func(c *gin.Context) {
		clientList, err := teamspeak.ClientList(client)
		if err != nil {
			log.Fatalln("Error Get ClientList: ", err)
		}
		c.JSON(http.StatusOK, clientList)
	})
	ts.GET("/whowhere", func(c *gin.Context) {
		teamspeak.WhoWhere(client)
	})
	r.Run(":80")
}
