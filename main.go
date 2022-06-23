package main

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	shell "github.com/ipfs/go-ipfs-api"
)

func main() {

	router := gin.Default()
	router.GET("/ipfs/:cid", getFromIPFS)

	router.Run("localhost:8000")
}

func getFromIPFS(c *gin.Context) {
	cid := c.Param("cid")

	// Where your local node is running on localhost:5002
	sh := shell.NewShell("localhost:5002")

	catResult, err := sh.Cat(cid)
	if err != nil {
		fmt.Println(err)
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(catResult)
	resp := buf.String()

	c.JSON(http.StatusOK, gin.H{"data": resp})

}
