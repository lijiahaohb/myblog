package model

import (
	"context"
	"ginblog/utils"
	"ginblog/utils/errmsg"
	"mime/multipart"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

var AccessKey = utils.AccessKey
var SecretKey = utils.SecretKey
var Bucket = utils.Bucket
var ImgUrl = utils.QiniuServer

func UploadFile(file multipart.File, fileSize int64) (string, int) {
	putPolicy := storage.PutPolicy{
		Scope: Bucket,
	}
	mac := qbox.NewMac(AccessKey, SecretKey)
	upToken := putPolicy.UploadToken(mac)

	config := storage.Config{
		Zone: &storage.ZoneHuabei,
		// 是否使用https域名
		UseHTTPS: false,
		// 上传是否使用CDN上传加速
		UseCdnDomains: false,
	}

	putExtra := storage.PutExtra{}

	formUploader := storage.NewFormUploader(&config)

	ret := storage.PutRet{}

	err := formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &putExtra)
	if err != nil {
		return "", errmsg.ERROR
	}
	url := ImgUrl + ret.Key
	return url, errmsg.SUCCESS

}
