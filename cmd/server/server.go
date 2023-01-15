package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	// static
	engine.Static("/js", "static/js/")
	engine.Static("/css", "static/css/")
	// engine.Static("/wasm", "static/wasm/")
	engine.LoadHTMLGlob("static/*.html")

	engine.GET("/", Index)                // ホームページ
	engine.GET("/register", RegisterPage) // アカウント登録用ページ
	engine.POST("/register", Index)       // アカウント登録
	engine.GET("/login", LoginPage)       // ログイン用ページ
	engine.POST("/login", Login)          // ログイン
	engine.POST("/add", AddContents)      // コンテンツ追加

	engine.Run(":3000")
}

// ページ用
// ホームページ
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

// ログインページ
func LoginPage(c *gin.Context) {

}

// アカウント登録ページ
func RegisterPage(c *gin.Context) {

}

// API用
// ログイン
func Login(c *gin.Context) {

}

// アカウント登録
func Register(c *gin.Context) {

}

// コンテンツ登録
func AddContents(c *gin.Context) {

}
