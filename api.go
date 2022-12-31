package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
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
			log.Println("Error Get ClientList: ", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, clientList)
	})
	{
		ts.GET("/whowhere", func(c *gin.Context) {
			whoWhere, err := teamspeak.WhoWhere(client)
			if err != nil {
				log.Println("Error Get WhoWhere:", err)
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			c.JSON(http.StatusOK, whoWhere)
		})
		ts.GET("/whowhere/short", func(c *gin.Context) {
			whoWhereShort, err := teamspeak.WhoWhereShort(client)
			if err != nil {
				log.Println("Error Get WhoWhereShort:", err)
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			c.JSON(http.StatusOK, whoWhereShort)
		})
	}

	ts.GET("/channelinfo/:id", func(c *gin.Context) {
		var channelInfo teamspeak.ChannelInfoStruct
		if number, err := strconv.Atoi(c.Param("id")); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		} else {
			channelInfo, err = teamspeak.ChannelInfo(client, number)
			if err != nil {
				log.Println("Error Get ChannelInfo:", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
		}
		c.JSON(http.StatusOK, channelInfo)
	})
	r.Run(":80")
}
