package ioc

type Store interface {
	StoreUser
	StoreManage
}

type StoreManage interface {
	// 从环境变量中加载对象配置
	LoadFromEnv(prefix string) error
}

type StoreUser interface {
	// 对象注册
	Registry(obj Object)
	// 获取对象的方法
	Get(name string, opts ...GetOption) Object
	// 根据对象类型，加载对象
	Load(obj any, opts ...GetOption) error
	// 打印对象列表
	List() []string
	// 数量统计
	Count() int
	// 遍历注入的对象
	ForEach(fn func(*ObjectWrapper))
}

type GetOption func(*option)

func defaultOption() *option {
	return &option{
		version: DEFAULT_VERSION,
	}
}

type option struct {
	version string
}

func (o *option) Apply(opts ...GetOption) *option {
	for i := range opts {
		opts[i](o)
	}
	return o
}
