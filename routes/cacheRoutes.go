package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/harshsri28/apica/module"
    "github.com/harshsri28/apica/controllers"
)

func InitCacheRoutes(incomingRoutes *gin.Engine, cacheInstance *cache.LRUCache) {
    controllers.InitializeCache(cacheInstance)

    incomingRoutes.GET("/cache", controllers.GetAllCache) 
    incomingRoutes.GET("/cache/:key", controllers.GetCache)
    incomingRoutes.POST("/cache", controllers.SetCache)
    incomingRoutes.DELETE("/cache/:key", controllers.DeleteCache)
}
