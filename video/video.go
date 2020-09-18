package video

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func Video()func(c *gin.Context){
	return func(c *gin.Context) {
		video,err:=os.Open("F:/workspace/go/web/2.mp4")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer video.Close()
		//c.Header("Content-Type","video/mp4")
		//c.Writer.Header().Add("Content-Disposition","attachment;filename=2.mp4")
		//视频点播
		c.File("./2.mp4")
		//http.ServeContent(c.Writer,c.Request,"",time.Now(),video)

	}
}

