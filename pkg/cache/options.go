package cache

// TODO: cc 取代为Option设计模式,可以对外提供优雅创建对象的方式
type Interface interface {
	Name() string
	Value() interface{}
}

type Option struct {
	name  string
	value interface{}
}

func New(name string, value interface{}) *Option {
	return &Option{
		name:  name,
		value: value,
	}
}

func (o *Option) Name() string {
	return o.name
}
func (o *Option) Value() interface{} {
	return o.value
}
