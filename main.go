package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Инициализируете роутер Gin, используя Default.
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/", homePage)
	// Используете функцию GET, чтобы связать метод GET HTTP и путь /albums с функцией обработчика.
	router.GET("/albums", getAlbums)
	// Связываете метод POST по пути /albums с функцией postAlbums.
	router.POST("/albums", postAlbums)
	// Связываете путь /albums/:id с функцией getAlbumByID. В Gin двоеточие перед элементом в пути означает, что этот элемент является параметром пути.
	router.GET("/albums/:id", getAlbumByID)

	// Используйте функцию Run, чтобы подключить маршрутизатор к http.Server и запустить сервер.
	router.Run("localhost:8080")
}
