package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

type Users struct {
	Id       string
	Password string
}

type Urls struct {
	Id     int
	UserId string
	Url    string
}

var dbname string = "link.db"
var dbtype string = "sqlite3"

func main() {
	//DB初期化
	DbInit()

	engine := gin.Default()
	// static
	engine.Static("/js", "static/js/")
	engine.Static("/css", "static/css/")
	// engine.Static("/wasm", "static/wasm/")
	engine.LoadHTMLGlob("templates/*.html")

	engine.GET("/", Index) // ホームページ
	engine.GET("/:id", UserPage)
	// engine.GET("/register", RegisterPage) // アカウント登録用ページ
	// engine.POST("/register", Index)       // アカウント登録
	// engine.GET("/login", LoginPage)       // ログイン用ページ
	// engine.POST("/login", Login)          // ログイン
	// engine.POST("/add", AddContents)      // コンテンツ追加

	engine.Run(":3000")
}

// ページ用
// ホームページ
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

// 個人ページ
func UserPage(c *gin.Context) {
	id := c.Param("id")

	db, err := sql.Open(dbtype, dbname)
	if err != nil {
		log.Print(err)
		// ここでエラーページに飛ばす
	}
	defer db.Close()

	sql := "select * from users where id = ?"
	// stmt, _ := db.Prepare(sql)
	// if err != nil {
	// 	log.Fatal(err)
	// 	// to error page.
	// }
	// defer stmt.Close()

	// get password
	u := Users{}
	err = db.QueryRow(sql, id).Scan(&u.Id, &u.Password)
	if err != nil {
		log.Print(err)
		// to error page.
	}

	// if u.Id != id {
	// 	log.Print("")
	// }

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

// DB操作
func DbInit() {
	db, err := sql.Open(dbtype, dbname)
	if err != nil {
		log.Fatal(err)
		// ここでエラーページに飛ばす
	}
	defer db.Close()

	//テーブル作成
	sql := " create table users if not exists ( id text primary key, password text );"
	if _, err = db.Exec(sql); err != nil {
		log.Fatal(err)
	}

	sql = " create table urls if not exists ( id integer primary key, userid text , url text );"
	if _, err = db.Exec(sql); err != nil {
		log.Fatal(err)
	}
}
