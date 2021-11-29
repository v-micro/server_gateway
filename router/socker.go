package router

import (
	socketio "github.com/googollee/go-socket.io"
	"github.com/labstack/echo/v4"
	"log"
)

func SocketInit(r *echo.Echo)  {
	newSocket := socketio.NewServer(nil)

	newSocket.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		log.Println("connected:", s.ID())
		return nil
	})

	newSocket.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		log.Println("notice:", msg)
		s.Emit("reply", "have "+msg)
	})

	newSocket.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
		s.SetContext(msg)
		return "recv " + msg
	})

	newSocket.OnEvent("/", "echo", func(s socketio.Conn, msg interface{}) {
		s.Emit("echo", msg)
	})

	newSocket.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})

	newSocket.OnError("/", func(s socketio.Conn, e error) {
		log.Println("meet error:", e)
	})

	newSocket.OnDisconnect("/", func(s socketio.Conn, reason string) {
		log.Println("closed", reason)
	})

	go newSocket.Serve()
	defer newSocket.Close()

	r.HideBanner = true

	r.Static("/", "./asset")
	r.Any("/socket.io/", func(context echo.Context) error {
		newSocket.ServeHTTP(context.Response(), context.Request())
		return nil
	})

}
