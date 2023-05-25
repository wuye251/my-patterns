package main

import (
	"fmt"
	"io"
	"math/rand"
)

// oss 接口指定实现类需要提供哪些能力
type OSS interface {
	UploadFile(localFilePath, objectKey string) error
	GetFile(objectKey string) (io.Reader, error)
	GetSignUrl(objectKey string) (string, error)
}

type AliOSS struct {
	EndPoint  string // 服务地址
	SecretId  string // 秘钥id
	SecretKey string // 秘钥key
	Bucket    string // 桶
}

// UploadFile 上传文件到阿里oss
func (AliOSS) UploadFile(localFilePath, objectKey string) error {
	fmt.Println("alioss upload file")
	return nil
}

// GetFile 下载/获取阿里oss文件
func (AliOSS) GetFile(objectKey string) (io.Reader, error) {
	fmt.Println("alioss get file")
	return nil, nil
}

// GetSignUrl 获取阿里oss临时访问地址
func (AliOSS) GetSignUrl(objectKey string) (string, error) {
	fmt.Println("alioss get sign url")
	return "", nil
}

type Minio struct {
	EndPoint string // 服务地址
	User     string // 用户
	Password string // 密码
	Bucket   string // 桶
}

// UploadFile 上传文件到minio
func (Minio) UploadFile(localFilePath, objectKey string) error {
	fmt.Println("minio upload file")
	return nil
}

// GetFile 下载/获取minio文件
func (Minio) GetFile(objectKey string) (io.Reader, error) {
	fmt.Println("minio get file")
	return nil, nil
}

// GetSignUrl 获取minio临时访问地址
func (Minio) GetSignUrl(objectKey string) (string, error) {
	fmt.Println("minio get sign url")
	return "", nil
}

// 生成实例抽象工厂
type Factory interface {
	CreateOSS() OSS
}

// 每个产品提供一个创建工厂
type AliOSSFactory struct{}

// CreateOSS  生成alioss 实例的工厂
func (AliOSSFactory) CreateOSS() *AliOSS {
	// (省略)获取配置文件中minio的相关配置
	return &AliOSS{
		EndPoint:  "alioss_endpoint",
		SecretId:  "alioss_secret_id",
		SecretKey: "alioss_secret_key",
		Bucket:    "alioss_env_buckect_name",
	}
}

type MinioFactory struct{}

// CreateOSS 生成minio的实例工厂
func (MinioFactory) CreateOSS() *Minio {
	// (省略)获取配置文件中minio的相关配置
	return &Minio{
		EndPoint: "minio_endpoint",
		User:     "minio_user",
		Password: "minio_password",
		Bucket:   "minio_env_bucket",
	}
}

func main() {
	// 获取配置文件判断是否私有部署
	var oss OSS
	isPrivate := rand.Intn(2)
	switch isPrivate {
	case 1:
		oss = MinioFactory{}.CreateOSS()
	default:
		oss = AliOSSFactory{}.CreateOSS()
	}

	// 业务逻辑
	oss.UploadFile("/tmp/test.txt", "tmp/test.txt")
	oss.GetFile("tmp/test.txt")
	oss.GetSignUrl("tmp/test.txt")
}
