package main

import (
	"fmt"
	"time"
)

// oss抽象类定义 同步/异步
type OSS interface {
	AsyncOSS
	SyncOSS
}

type SyncOSS interface {
	SyncUpload(filepath string) error
}
type AsyncOSS interface {
	AsyncUpload(filepath, callbackUrl string) error
}

// 关于minio的接口
type Minio struct{}

func (m Minio) SyncUpload(filepath string) error {
	fmt.Println("minio sync upload")
	return nil
}

func (m Minio) AsyncUpload(filepath, callbackUrl string) error {
	fmt.Printf("minio async upload, callback url is %s\n", callbackUrl)
	return nil
}

// 阿里只有同步上传
type AliOSS struct{}

func (a AliOSS) SyncUpload(filepath string) error {
	fmt.Println("ali sync upload")
	return nil
}

// oss同步适配异步的适配器
type OSSAdapter struct {
	oss SyncOSS
}

// 同步还是走原先的oss接口
func (o *OSSAdapter) SyncUpload(filepath string) error {
	o.oss.SyncUpload(filepath)
	return nil
}

// 异步需要通过同步进行实现适配
func (o OSSAdapter) AsyncUpload(filepath, callbackUrl string) error {
	if err := o.oss.SyncUpload(filepath); err != nil {
		return err
	}
	go o.asyncUpload(filepath, callbackUrl)
	return nil
}

func (o OSSAdapter) asyncUpload(filepath, callback string) error {
	// 另启动一个协程进行pull
	// 拿到结果后调用callback
	for {
		fmt.Println("async adapter run")
		time.Sleep(2 * time.Second)
	}
}

func main() {
	var oss OSS
	oss = &Minio{}

	oss.SyncUpload("tmp.txt")
	oss.AsyncUpload("tmp.txt", "callbackurl")

	// 阿里适配
	ali := &AliOSS{}
	oss = &OSSAdapter{ali}
	oss.SyncUpload("tmp.txt")
	oss.AsyncUpload("tmp.txt", "callbackurl")

	// output:
	// minio sync upload
	// minio async upload, callback url is callbackurl
	// ali sync upload
	// ali sync upload
}
