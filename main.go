package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ExibeTodosALunos(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"id":   "1",
		"nome": "Matheus Sasso",
	})
}

func main() {
	r := gin.Default()
	r.GET("/alunos", ExibeTodosALunos)
	r.Run()

}
