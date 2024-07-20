package main

import (
    "github.com/gin-gonic/gin"
    "github.com/harshsri28/apica/module"
    "github.com/harshsri28/apica/routes"
)

func main() {
    port := "3000"

    lruCache := cache.NewLRUCache(10)
    router := gin.New()
    router.Use(gin.Logger())

    routes.InitCacheRoutes(router,lruCache)

    router.Run(":" + port)
}
