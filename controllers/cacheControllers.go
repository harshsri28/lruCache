package controllers

import (
    "github.com/gin-gonic/gin"
    "github.com/harshsri28/apica/module"
    "net/http"
    "time"
)

var lruCache *cache.LRUCache

func InitializeCache(cacheInstance *cache.LRUCache) {
    lruCache = cacheInstance
    lruCache.StartExpirationRoutine()
}


func GetAllCache(c *gin.Context) {
    allItems := lruCache.GetAll()
    c.JSON(http.StatusOK, gin.H{"items": allItems})
}

func GetCache(c *gin.Context) {
    key := c.Param("key")
    if value, found := lruCache.Get(key); found {
        c.JSON(http.StatusOK, gin.H{"key": key, "value": value})
    } else {
        c.JSON(http.StatusNotFound, gin.H{"error": "key not found"})
    }
}

func SetCache(c *gin.Context) {
    var requestBody struct {
        Key       string `json:"key" binding:"required"`
        Value     string `json:"value" binding:"required"`
        Duration  int    `json:"duration" binding:"required"`
    }
    if err := c.ShouldBindJSON(&requestBody); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
        return
    }

    duration := time.Duration(requestBody.Duration) * time.Second
    lruCache.Set(requestBody.Key, requestBody.Value, duration)
    c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func DeleteCache(c *gin.Context) {
    key := c.Param("key")
    lruCache.Delete(key)
    c.JSON(http.StatusOK, gin.H{"status": "success"})
}
