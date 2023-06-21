package main

import (
	"context"
	pb "go_mod_testdemo"
	"google.golang.org/grpc"
)

type AdService struct {
}

func (this *AdService) GetBrandInfo(ctx context.Context, req *pb.GetBrandInfoReq) (*pb.GetBrandInfoRsp, error) {
	if req.GetBrandId() == 10 {
		return &pb.GetBrandInfoRsp{
			BrandId:   10,
			BrandName: "ad",
		}, nil
	}
	return &pb.GetBrandInfoRsp{
		BrandId:   req.BrandId,
		BrandName: "other",
	}, nil
}
func main() {
	// 1. 初始化grpc对象
	grpcServer := grpc.NewServer()
	// 2. 注册服务
	pb.RegisterAdServiceServer(grpcServer, new(AdService))

}
