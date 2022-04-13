package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Welcomepage(context *gin.Context) {

	//call the HTML Method of the context to render the template

	context.HTML(
		//setup the status of the template
		http.StatusOK,

		//use which template or deploy which template
		"index.html",
		gin.H{
			"title": "Home page"})

}

func SignUp() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}
