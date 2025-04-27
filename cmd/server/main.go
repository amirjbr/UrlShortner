package main

import (
	"UrlShortner/config"
	"flag"
	"fmt"
	"os"
)

var configPath = flag.String("config", "config.json", "service configuration file")

func main() {
	flag.Parse()
	if v := os.Getenv("CONFIG_PATH"); len(v) > 0 {
		*configPath = v
	}
	c := config.MustReadConfig(*configPath)
	fmt.Println(c)
	// db := storage.NewDbInstanse(c)

	// setup Database and send config to it

	// s := gin.Default()

	// s.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	// s.Run()

}
