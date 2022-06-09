package common

import (
	"bytes"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io/ioutil"
	"mime/multipart"
	"os"
)

func UploadOss(file *multipart.FileHeader, filename string) error {
	// 创建OSSClient实例。
	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	// 阿里云账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM用户进行API访问或日常运维，请登录RAM控制台创建RAM用户。
	client, err := oss.New("https://oss-cn-guangzhou.aliyuncs.com", "LTAI4GBDrUUiC6xYfLPQxgFg", "VxqUPo0A4QcDbEjJC8FsRIFQKQI2bF")
	if err != nil {
		os.Exit(-1)
		return err
	}

	// 填写存储空间名称，例如examplebucket。
	bucket, err := client.Bucket("dou-yin-zzz")
	if err != nil {
		os.Exit(-1)
		return err
	}


	// 将Byte数组上传至exampledir目录下的exampleobject.txt文件。
	fileContent, _ := file.Open()
	byteContent, _ := ioutil.ReadAll(fileContent)
	err = bucket.PutObject(filename, bytes.NewReader(byteContent))
	if err != nil {
		os.Exit(-1)
		return err
	}
	return nil
}