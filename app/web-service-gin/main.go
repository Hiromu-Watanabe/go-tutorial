package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 構造体でJSONのフィールド名をカスタマイズしてメモリに保存
// これがなければ大文字のフィールド名になる（例）ID, Title...
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// レコードのアルバムデータのスライス
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("0.0.0.0:8080")
}

// すべてのアルバムのリストをJSONで返却
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// リクエストボディで受け取ったJSONを新規アルバムとして追加
func postAlbums(c *gin.Context) {
  var newAlbum album

  // BindJSONで受け取ったJSONをnewAlbumにバインド
  if err := c.BindJSON(&newAlbum); err != nil {
      return
  }

  // albumsスライスに新規アルバムを追加
  albums = append(albums, newAlbum)
  c.IndentedJSON(http.StatusCreated, newAlbum)
}

// リクエストパスのIDに一致するアルバムを探す
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// アルバムデータのスライスから一致するアルバムを探す
	for _, a := range albums {
			if a.ID == id {
					c.IndentedJSON(http.StatusOK, a)
					return
			}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}