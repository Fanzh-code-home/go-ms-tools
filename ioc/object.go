package ioc

const (
	DEFAULT_VERSION = "v1"
)

// Object对象的包装器
type ObjectWrapper struct {
	Name           string
	Version        string
	AllowOverWrite bool
	Priority       int
	Value          Object
}

func NewObjectWrapper(obj Object) *ObjectWrapper {
	name, version := GetIocObjectUid(obj)
	return &ObjectWrapper{
		Name:           name,
		Version:        version,
		Priority:       obj.Priority(),
		AllowOverWrite: obj.AllowOverwrite(),
		Value:          obj,
	}
}
