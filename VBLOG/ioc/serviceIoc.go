package ioc

import "fmt"

var ServiceIocList = map[string]ServiceObjIoc{}

// 注册ioc
func RegistryServiceIoc(obj ServiceObjIoc) {
	ServiceIocList[obj.Name()] = obj

}

// 获取iocbyname
func GetServiceIocByName(name string) ServiceObjIoc {
	ioc, ok := ServiceIocList[name]
	if !ok {
		message := fmt.Sprintf("%s ioc  not found", name)
		panic(message)
	}
	return ioc
}

// 初始化ioc，一般用于ioc对象自己初始化，比如数据库连接池建立等

func InitServiceIoc() error {
	for _, ioc := range ServiceIocList {
		ioc.Init()
	}
	return nil
}

func ShowServiceIoc() (names []string) {
	for k := range ServiceIocList {
		names = append(names, k)
	}
	return names
}
