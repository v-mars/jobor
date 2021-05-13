package proto

import (
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"jobor/internal/proto/pb"
	"jobor/internal/response"
	"log"
	"time"
)

const (
	Address = "localhost:50051" 	// gRPC server address
)


var (
	Client  pb.TaskServiceClient
)

func NewClientGRPC() error {
	//service.WithInsecure()指定后才不会报错
	ctx, cel := context.WithTimeout(context.Background(), time.Second*3)
	defer cel()
	conn, err := grpc.DialContext(ctx,Address,grpc.WithInsecure())
	//conn, err := service.DialContext(
	//	ctx,Address, service.WithBlock(),service.WithInsecure(),
	//	)  //
	//fmt.Printf("NewClientGRPC: %++v %s\n", conn, err)
	if err != nil{
		log.Fatal("gRPC client dial error: ", err)
		return err
	}
	//defer conn.Close()

	Client = pb.NewTaskServiceClient(conn)


	/*//通过grpc 库 建立一个连接
	conn ,err := service.Dial(ADDRESS,service.WithInsecure())
	if err != nil{
		return
	}
	defer conn.Close()
	//通过刚刚的连接 生成一个client对象。
	c := pro.NewGreeterClient(conn)*/
	//调用服务端推送流
	//reqstreamData := &pb.StreamReqData{Data:"aaa"}
	//res,_ := Client.GetStream(context.Background(),reqstreamData)
	//for {
	//	aa,err := res.Recv()
	//	if err != nil {
	//		log.Println(err)
	//		break
	//	}
	//	log.Println(aa)
	//}
	////客户端 推送 流
	//putRes, _ := Client.PutStream(context.Background())
	//i := 1
	//for {
	//	i++
	//	putRes.Send(&pb.StreamReqData{Data:"ss"})
	//	time.Sleep(time.Second)
	//	if i > 10 {
	//		break
	//	}
	//}
	////服务端 客户端 双向流
	//allStr,_ := Client.AllStream(context.Background())
	//go func() {
	//	for {
	//		data,_ := allStr.Recv()
	//		log.Println(data)
	//	}
	//}()
	//
	//go func() {
	//	for {
	//		allStr.Send(&pb.StreamReqData{Data:"ssss"})
	//		time.Sleep(time.Second)
	//	}
	//}()

	//select {}

	return err
}

func ClientGRPCTest(c *gin.Context)  {
	type Param struct {
		Name       string `form:"name"`
	}
	var param Param
	if err := c.ShouldBindQuery(&param); err!=nil{
		response.ParamFailed(c, err)
		return
	}
	res, err := Client.RunRPC(c, &pb.TaskRequest{Name:param.Name})

	if err!=nil{
		response.Error(c, err)
	}else {
		response.Success(c, res)
	}
}
