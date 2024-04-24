package middleware

import (
    "github.com/gin-gonic/gin"
    "log"
    "net/http"
)

func ErrorHandler() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Next()

        if len(c.Errors) > 0 {
            for _, e := range c.Errors {
                log.Println("Error:", e.Error())
            }
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
            return
        }
    }
}
