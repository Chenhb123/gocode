package main

import (
	"context"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"google.golang.org/grpc"
	"helloworld"
)

func main() {
	app := newApp()

	app.Logger().SetLevel("debug")

	app.Run(iris.TLS(":443", "server.crt", "server.key"))
}

func newApp() *iris.Application {
	app := iris.New()

	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("<h1>Index Page</h1>")
	})

	ctr1 := &myController{}
	grpcServer := grpc.NewServer()
	helloworld.RegisterGreeterServer(grpcServer, ctr1)
	mvc.New(app).Handle(ctr1, mvc.GRPC{
		Server:      grpcServer,
		ServiceName: "helloworld.Greeter",
		Strict:      false,
	})

	return app
}

type myController struct {
}

func (c *myController) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{Message: "Hello " + in.GetName()}, nil
}
