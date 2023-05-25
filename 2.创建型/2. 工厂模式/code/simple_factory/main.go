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

// 生成实例的代理工厂
type Factory struct{}

// 生成实例的不同判断逻辑
func (f Factory) CreateOSS() OSS {
	// 获取配置文件判断是否私有部署
	isPrivate := rand.Intn(2)
	switch isPrivate {
	case 1:
		return f.createMinio()
	default:
		return f.createAliOSS()
	}
}

// createMinio 生成minio的实例
func (Factory) createMinio() *Minio {
	// (省略)获取配置文件中minio的相关配置
	return &Minio{
		EndPoint: "minio_endpoint",
		User:     "minio_user",
		Password: "minio_password",
		Bucket:   "minio_env_bucket",
	}
}

// createAliOSS 生成alioss的实例
func (Factory) createAliOSS() *AliOSS {
	return &AliOSS{
		EndPoint:  "alioss_endpoint",
		SecretId:  "alioss_secret_id",
		SecretKey: "alioss_secret_key",
		Bucket:    "alioss_env_buckect_name",
	}
}

func main() {
	// 创建实例
	oss := Factory{}.CreateOSS()
	// 业务逻辑
	oss.UploadFile("/tmp/test.txt", "tmp/test.txt")
	oss.GetFile("tmp/test.txt")
	oss.GetSignUrl("tmp/test.txt")
}
