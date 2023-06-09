package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func main() {
	httpClient := http.Client{}
	router := gin.Default()

	// Expose an endpoint which will call an upstream service
	router.GET("/pet/:id", func(c *gin.Context) {
		id := c.Param("id")
		res, err := httpClient.Get("http://localhost:8080/v2/pet/" + id)
		if err != nil {
			c.JSON(502, gin.H{
				"status": err.Error(),
			})
		} else {
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					fmt.Printf("An error was encountered when closing body: %s\n", err.Error())
				}
			}(res.Body)

			b, err := io.ReadAll(res.Body)
			if err != nil {
				c.JSON(500, gin.H{
					"status": err.Error(),
				})
			} else {
				c.String(res.StatusCode, "%s", b)
			}
		}

	})

	err := router.Run("localhost:3000")
	if err != nil {
		panic("Could not run web service")
	}
}
