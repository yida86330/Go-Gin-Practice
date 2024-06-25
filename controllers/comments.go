package controllers

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "os"

    "github.com/gin-gonic/gin"
)

type Comment struct {
	User    string `json:"user"`
	Time    string `json:"time"`
	Message string `json:"message"`
}

// GetComments 獲取留言記錄
// @Summary Get Comments
// @Description Get all comments
// @Produce  json
// @Success 200 {array} Comment
// @Router /comments [get]
func GetComments(c *gin.Context) {
    file, err := os.Open("comments.txt")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer file.Close()

    data, err := ioutil.ReadAll(file)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to read comments 2"})
        return
    }

    comments := string(data)
    c.JSON(http.StatusOK, gin.H{"comments": comments})
}

// PostComment 發布留言
// @Summary Post Comment
// @Description Post a new comment
// @Accept  json
// @Produce  json
// @Param   comment     body   Comment true "Comment"
// @Success 200 {string} string "ok"
// @Router /comments [post]
func PostComment(c *gin.Context) {

    var comment Comment
    if err := c.ShouldBindJSON(&comment); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    file, err := os.OpenFile("comments.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to write comment"})
        return
    }
    defer file.Close()

    commentData := fmt.Sprintf("User: %s, Time: %s, Message: %s\n", comment.User, comment.Time, comment.Message)
    if _, err = file.WriteString(commentData); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to write comment"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"status": "comment posted"})
}
