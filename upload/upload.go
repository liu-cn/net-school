package upload

import (
	"fmt"
	"gin/qiniu"
	"github.com/gin-gonic/gin"
	"net/http"
)

func QiniuUpload()func(c *gin.Context){
	return func(c *gin.Context) {
		file, fileHeader, err := c.Request.FormFile("file")

		if err != nil {
			fmt.Println("文件接收失败！",err)
			fmt.Println(file,fileHeader)
			c.JSON(http.StatusOK,gin.H{
				"msg":"err",
				"url":"no",
			})
			return
		}
		defer file.Close()
		url, ok := qiniu.UploadFile(file, fileHeader.Size)
		if ok!=1 {
			fmt.Println("文件上传失败！",err)
			c.JSON(http.StatusOK,gin.H{
				"msg":"err",
				"url":url,
			})
			return
		}

		fmt.Println(url)
		c.JSON(http.StatusOK,gin.H{
			"msg":"ok",
			"url":url,
		})

	}
}

//func Upload() func(c *gin.Context){
//	return func(c *gin.Context){
//		form, err := c.MultipartForm()
//		files := form.File["file"]
//		if err != nil {
//			c.JSON(http.StatusBadRequest,gin.H{
//				"msg":"文件上传失败！",
//			})
//			return
//		}
//		path:=make([]string, len(files))
//		for i,file := range files{
//			log.Println(file.Filename)
//			dst:=fmt.Sprintf("./%s",file.Filename)
//			c.SaveUploadedFile(file,dst)
//			path[i]=dst
//		}
//		c.JSON(http.StatusOK,gin.H{
//			"msg":"文件上传成功！",
//			"path":path,
//		})
//
//
//		//formFile, err := c.FormFile("file")
//		//fmt.Println(formFile)
//		//if err != nil {
//		//	c.JSON(http.StatusBadGateway,gin.H{
//		//		"msg":"文件上传失败！",
//		//	})
//		//	return
//		//}
//		//c.SaveUploadedFile(formFile,"./"+formFile.Filename)
//		//path:="F:/workspace/go/web/"+formFile.Filename
//		//c.JSON(http.StatusOK,gin.H{
//		//	"msg":"文件上传成功！",
//		//	"img":path,
//		//})
//	}
//}





