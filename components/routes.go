package components

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/links", getLinks)
	router.GET("/links/html", getLinksHTML)
	router.POST("/links", postLinks)
	router.GET("/links/refresh", refreshLinksHTML)
	router.GET("/links/:id", getLinksByID)
}
