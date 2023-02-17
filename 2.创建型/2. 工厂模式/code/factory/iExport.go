package factory

// 接口类
type IExport interface {
	setName(name string)
	setSuffix(name string)
	GetName() string
	GetSuffix() string
}
