package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("static", "/static")

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})

	router.POST("/", homePostRouteFunc)

	router.GET("/login", func(ctx *gin.Context) {
		var _ = sessions.Default(ctx)
		ctx.HTML(http.StatusOK, "login.html", nil)
	})

	router.GET("/store", func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		storeName := session.Get("storename")
		storeId := session.Get("storeid")

		if storeName == nil || storeId == nil {
			ctx.Redirect(http.StatusMovedPermanently, "/login")
		}

		ctx.HTML(http.StatusOK, "store.html", gin.H{
			"storeName": storeName,
			"storeId":   storeId,
		})
	})

	router.Run()
}

func homePostRouteFunc(ctx *gin.Context) {
	session := sessions.Default(ctx)
	storeName := ctx.PostForm("storename")
	storeId := uuid.New().String()

	session.Set("storename", storeName)
	session.Set("storeid", storeId)

	session.Save()
	ctx.Redirect(http.StatusMovedPermanently, "/store")
}
