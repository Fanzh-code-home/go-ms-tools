package ioc

import (
	"fmt"
	"reflect"
)

type defaultStore struct {
	conf  *LoadConfigRequest
	store []*NamespaceStore
}

// 带有命名空间的仓库(Store)
type NamespaceStore struct {
	// 仓库空间名称
	Namespace string
	// 空间的优先级
	Priority int
	// 空间里面的对象列表
	Items []*ObjectWrapper
}

func (ns *NamespaceStore) SetPriority(v int) *NamespaceStore {
	ns.Priority = v
	return ns
}

func (ns *NamespaceStore) Registry(v Object) {
	obj := NewObjectWrapper(v)
	oldObj, index := ns.getWithIndex(obj.Name, obj.Version)
	// 没有， 直接添加
	if oldObj == nil {
		ns.Items = append(ns.Items, obj)
		return
	}
	// 有， 允许覆盖写则直接修改
	if obj.AllowOverWrite {
		ns.Items[index] = obj
		return
	}
	// 有， 不允许覆盖写则报错，panic
	panic(fmt.Sprintf("ioc obj %s: %s has registed", obj.Name, obj.Version))
}

func (ns *NamespaceStore) Get(name string, opts ...GetOption) Object {
	opt := defaultOption().Apply(opts...)
	obj, _ := ns.getWithIndex(name, opt.version)
	if obj == nil {
		return nil
	}
	return obj.Value
}

// 根据对象类型加载对象

func (ns *NamespaceStore) Load(target any, opts ...GetOption) error {
	t := reflect.TypeOf(target)
	v := reflect.ValueOf(target)
	var obj Object
	switch t.Kind() {
	case reflect.Interface:
		objs := ns.ImplementInterface(t)
		if len(objs) > 0 {
			obj = objs[0]
		}
	default:
		obj = ns.Get(t.String(), opts...)
	}
	// 注入值
	if obj != nil {
		objValue := reflect.ValueOf(obj)
		if !(v.Kind() == reflect.Ptr && objValue.Kind() == reflect.Ptr) {
			return fmt.Errorf("target and object must both be pointers or non-pointers")
		}
		v.Elem().Set(objValue.Elem())
	}
	return nil
}

// 寻找实现了接口的对象
func (ns NamespaceStore) ImplementInterface(objType reflect.Type, opts ...GetOption) (objs []Object) {
	opt := defaultOption().Apply(opts...)
	for i := range ns.Items {
		o := ns.Items[i]
		// 断言获取的对象是否满足接口类型
		if o != nil && reflect.TypeOf(o.Value).Implements(objType) {
			if o.Version == opt.version {
				objs = append(objs, o.Value)
			}
		}
	}
	return
}

func (ns *NamespaceStore) getWithIndex(name, version string) (*ObjectWrapper, int) {
	for i := range ns.Items {
		obj := ns.Items[i]
		if obj.Name == name && obj.Version == version {
			return obj, i
		}
	}
	return nil, -1
}
