package mygin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ErrInfo struct {
	Code  int
	Title string
	Msg   string
	Link  string
	Btn   string
}

func ShowErrorPage(c *gin.Context, i ErrInfo) {

	c.JSON(http.StatusOK, gin.H{
		"Code":    i.Code,
		"Message": i.Msg,
	})
	c.Abort()
}
