package ioc

import "context"

// Object接口，注册到ioc空间托管的对象需要实现的方法
type Object interface {
	// 对象初始化方法， 需要有初始化的方法属性
	Init() error
	// 对象的名称， 根据名称可以从空间中抽取出来对象实例
	Name() string
	// 对象版本,做版本控制使用，默认v1
	Version() string
	// 对象优先级，根据优先级 控制对象初始化的顺序
	Priority() int
	// 对象的销毁方法，服务关闭时调用
	Close(ctx context.Context) error
	// 是否允许同名对象被替换, 默认不允许被替换
	AllowOverwrite() bool
	// 对象一些元数据，对象的更多描述信息，扩展使用
	Meta() ObjectMeta
}

type ObjectMeta struct {
	// 自定义路径前缀
	CustomPathPrefix string
	// 其他扩展
	Extra map[string]string
}
