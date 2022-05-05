package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album для хранения данных альбома в памяти.
type album struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: 1, Title: "Blue train", Artist: "John Coltrane", Price: 56.99},
	{ID: 2, Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: 3, Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan ", Price: 39.99},
}

// getAlbums возвращает список всех альбомов в формате JSON.
func getAlbums(c *gin.Context) {
	// Вызовете Context.IndentedJSON, чтобы сериализовать структуру в JSON и добавить ее в ответ.
	// Первый аргумент функции - это код состояния HTTP, который вы хотите отправить клиенту. Здесь вы передаете константу StatusOK из пакета net/http, чтобы указать 200 OK.
	c.IndentedJSON(http.StatusOK, albums)
}

func main() {
	// Инициализируете роутер Gin, используя Default.
	router := gin.Default()

	// Используете функцию GET, чтобы связать метод GET HTTP и путь /albums с функцией обработчика.
	router.GET("/albums", getAlbums)
	// Связываете метод POST по пути /albums с функцией postAlbums.
	router.POST("/albums", postAlbums)

	// Используйте функцию Run, чтобы подключить маршрутизатор к http.Server и запустить сервер.
	router.Run("localhost:8080")
}

// postAlbums добавляет альбом из JSON, полученного в теле запроса.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Вызов BindJSON для привязки полученного JSON к newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Добавляем в срез новый альбом.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID находит альбом,
// значение идентификатора которого совпадает с
// параметром id, отправленным клиентом,
// затем возвращает этот альбом в качестве ответа
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Перебираем список альбомов в поисках альбома,
	// значение идентификатора которого соответствует параметру.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
