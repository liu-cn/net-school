package qiniu

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	"mime/multipart"
)


var (
	Bucket="liuspace"
	AccessKey="04-c10O2qR3ggEfliiIrDSXA56g39D3xwu76UL6A"
	SecretKey="r_6-SO1RHgdjt7aNCKZ_micwVYwsVo5EUtrXQBi-"
	QiniuServe="http://cdn.liubr.com/"
)

func UploadFile(file multipart.File,fileSize int64) (string,int) {
	putPolicy:=storage.PutPolicy{
		Scope: Bucket,
	}
	mac:=qbox.NewMac(AccessKey,SecretKey)

	upToken:=putPolicy.UploadToken(mac)

	cfg:=storage.Config{
		Zone: 	&storage.ZoneHuadong,
		UseCdnDomains: false,
		UseHTTPS: false,
	}
	putExtra:=storage.PutExtra{}

	formUploader:=storage.NewFormUploader(&cfg)
	ret:=storage.PutRet{}

	err:=formUploader.PutWithoutKey(context.Background(),&ret,upToken,file,fileSize,&putExtra)
	if err!=nil {
		return "",0
	}

	url:=QiniuServe+ret.Key
	return url,1
}

func GetQiniuToken(c *gin.Context) {
	putPolicy:=storage.PutPolicy{
		Scope: Bucket,
	}
	mac:=qbox.NewMac(AccessKey,SecretKey)

	upToken:=putPolicy.UploadToken(mac)
	c.JSON(200,gin.H{
		"msg":"ok",
		"qiniuToken":upToken,
	})
}