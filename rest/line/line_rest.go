package line

import (
	"example/homework/chatapp/service/line"
	"net/http"

	"github.com/gin-gonic/gin"
)

func apiCallBack(c *gin.Context) {
	if err := line.ParseRequest(c.Writer, c.Request); err != nil {
		c.IndentedJSON(http.StatusOK, gin.H{"status": "error", "errMsg": err})
	}
	c.IndentedJSON(http.StatusOK, gin.H{"status": "ok"})
}
