package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func homePage(c *gin.Context) {
	c.HTML(http.StatusOK, "templates/index.html", gin.H{"title": "Home Page"})
}

// getAlbums возвращает список всех альбомов в формате JSON.
func getAlbums(c *gin.Context) {
	// Вызовете Context.IndentedJSON, чтобы сериализовать структуру в JSON и добавить ее в ответ.
	// Первый аргумент функции - это код состояния HTTP, который вы хотите отправить клиенту. Здесь вы передаете константу StatusOK из пакета net/http, чтобы указать 200 OK.
	c.IndentedJSON(http.StatusOK, albums)
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
		if s, _ := strconv.Atoi(id); a.ID == s {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
