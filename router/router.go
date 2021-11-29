package router

import (
	"github.com/labstack/echo/v4"
	"server_gateway/server"
)

func RouterInit(r *echo.Echo)  {
	// order 服务
	order := r.Group("/order")
	{
		//订单服务
		order.GET("/ping/get",server.Order{}.Get)
	}

	//// socket 长链接
	//sockerServer := socketio.NewServer(nil)
	//sockerServer.OnConnect("/", func(s socketio.Conn) error {
	//	s.SetContext("")
	//	log.Println("connected:", s.ID())
	//	return nil
	//})
	//sockerServer.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
	//	log.Println("notice:", msg)
	//	s.Emit("reply", "have "+msg)
	//})
	//sockerServer.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
	//	s.SetContext(msg)
	//	return "recv " + msg
	//})
	//sockerServer.OnEvent("/", "echo", func(s socketio.Conn, msg interface{}) {
	//	s.Emit("echo", msg)
	//})
	//sockerServer.OnEvent("/", "bye", func(s socketio.Conn) string {
	//	last := s.Context().(string)
	//	s.Emit("bye", last)
	//	s.Close()
	//	return last
	//})
	//sockerServer.OnError("/", func(s socketio.Conn, e error) {
	//	log.Println("meet error:", e)
	//})
	//sockerServer.OnDisconnect("/", func(s socketio.Conn, reason string) {
	//	log.Println("closed", reason)
	//})
	//go sockerServer.Serve()
	//defer sockerServer.Close()
	//r.Any("/socket.io/", func(context echo.Context) error {
	//	sockerServer.ServeHTTP(context.Response(), context.Request())
	//	return nil
	//})

}
