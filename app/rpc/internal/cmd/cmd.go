package cmd

import (
	"context"
	"github.com/gogf/gf/contrib/registry/consul/v2"
	"github.com/gogf/gf/v2/frame/g"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"middle_srv/app/rpc/internal/controller/user"

	//"github.com/gogf/gf/contrib/registry/etcd/v2"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			opts := []consul.Option{
				consul.WithAddress("consul1:8500"),
				consul.WithAddress("consul2:8500"),
				consul.WithAddress("consul3:8500"),
				consul.WithAddress("consul4:8500"),
			}
			registry, err := consul.New(opts...)

			if err != nil {
				g.Log().Fatal(context.Background(), err)
			}
			grpcx.Resolver.Register(registry)

			c := grpcx.Server.NewConfig()
			c.Options = append(c.Options, []grpc.ServerOption{
				grpcx.Server.ChainUnary(
					grpcx.Server.UnaryValidate,
				)}...,
			)
			s := grpcx.Server.New(c)
			
			user.Register(s)
			reflection.Register(s.Server)
			s.Run()
			return nil
		},
	}
)
