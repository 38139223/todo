package models

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"fmt"
)

type AliOss struct {
	endPoint string
	accessID string
	accessKey string
}
func NewAliOss(endPoint,accessId,accessKey string)(*AliOss){
	aliOss := &AliOss{}
	aliOss.endPoint = endPoint
	aliOss.accessID = accessId
	aliOss.accessKey = accessKey
	return aliOss
}
func (this *AliOss)PutFileOss(object,path string)(bool){
	client,err := oss.New(this.endPoint,this.accessID,this.accessKey)
	if err != nil {
		fmt.Println(err)
		return false
	}
	bucket,err := client.Bucket("web-index")
	if err != nil {
		fmt.Println(err)
		return false
	}
/*	fd,err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return false
	}*//*
	defer fd.Close()*/
	err = bucket.PutObjectFromFile(object,path)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}