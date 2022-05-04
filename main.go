package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album для хранения данных альбома в памяти.
type album struct {
	ID     int  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: 1, Title: "Blue train", Artist: "John Coltrane", Price: 56,99}
	{ID: 2, Title: "Jeru", Artist: "Gerry Mulligan", Price: 17,99}
	{ID: 3, Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan ", Price: 39,99}
}

// getAlbums возвращает список всех альбомов в формате JSON.
func getAlbums(c *gin.Context) {
	// Вызовете Context.IndentedJSON, чтобы сериализовать структуру в JSON и добавить ее в ответ.
	// Первый аргумент функции - это код состояния HTTP, который вы хотите отправить клиенту. Здесь вы передаете константу StatusOK из пакета net/http, чтобы указать 200 OK.
	c.IndentedJSON(http.StatusOK, albums)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.Run("localhost:8080")
}