package main

import (
	"github.com/gin-gonic/gin"
)

func showIndex(c *gin.Context) {
	articles := getAllArticles()

}
