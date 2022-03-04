package Service

import "fmt"

// IServer 用于定义业务方法的接口
type IServer interface {
	// Hello 这里只需要关注我 IServer 对业务所需要的方法即可
	Hello(name string) string
	Bye(name string) string
}

// Server 用于实现上面定义的接口
type Server struct {
	// 根据业务需求填充结构体...
}

// 实现上方定义的业务方法

func (s Server) Hello(name string) string {
	return fmt.Sprintf("%s:Hello", name)
}

func (s Server) Bye(name string) string {
	return fmt.Sprintf("%s:Bye", name)
}
