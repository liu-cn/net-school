package download

import "github.com/gin-gonic/gin"

func Download(ctx *gin.Context)  {
	ctx.File("E:\\workspace\\Go\\net-school\\download\\apk\\aiyingyuan.apk")
}