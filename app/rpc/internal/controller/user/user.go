package user

import (
	"context"
	"fmt"
	v1 "middle_srv/app/rpc/api/user/v1"
	"time"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
)

type Controller struct {
	v1.UnimplementedUserServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterUserServer(s.Server, &Controller{})
}

func (*Controller) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	time.Sleep(10 * time.Second)
	fmt.Println("干马爹")
	return nil, nil
}

func (*Controller) Modify(ctx context.Context, req *v1.CreateReq) (res *v1.CreateReq, err error) {
	return nil, nil
}
