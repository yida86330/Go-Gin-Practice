package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

// Login 用戶登入
// @Summary Login
// @Description Login to the system
// @Accept  json
// @Produce  json
// @Param   username     body   string true "Username"
// @Param   password     body   string true "Password"
// @Success 200 {string} string "ok"
// @Router /login [post]
func Login(c *gin.Context) {
    type Login struct {
        Username string `json:"username"`
		Password string `json:"password"`
    }

    var login Login
    if err := c.ShouldBindJSON(&login); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // 假設登入成功
    c.JSON(http.StatusOK, gin.H{"status": "login successful"})
}
