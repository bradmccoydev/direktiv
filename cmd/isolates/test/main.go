package main

import (
	"context"

	log "github.com/sirupsen/logrus"
	igrpc "github.com/vorteil/direktiv/pkg/isolates/grpc"
	"google.golang.org/grpc"
)

func main() {

	log.Infof("run client")

	// img := "vorteil/request:v2"
	// var sz int32 = 0
	// info := &igrpc.BaseInfo{
	// 	Image: &img,
	// 	Size:  &sz,
	// }

	conn, err := grpc.Dial("127.0.0.1:30234", grpc.WithInsecure())
	if err != nil {
		log.Errorf("ERR %v", err)
	}
	defer conn.Close()

	client := igrpc.NewIsolatesServiceClient(conn)

	// svn := "w-8829097305702293016"
	// sr := igrpc.UpdateIsolateRequest{
	// 	Info:        info,
	// 	ServiceName: &svn,
	// }
	//
	// _, err = client.UpdateIsolate(context.Background(), &sr)
	// if err != nil {
	// 	log.Errorf("ERR %v", err)
	// }

	// a := make(map[string]string)
	// a["direktiv.io/workflow"] = "dsdsdsssd"
	//
	// g2 := igrpc.ListIsolatesRequest{
	// 	Annotations: a,
	// }
	//
	// _, err = client.ListIsolates(context.Background(), &g2)
	// log.Infof("new client %v", client)

	// g2 := igrpc.GetIsolateRequest{
	// 	ServiceName: &svn,
	// }
	// items, err := client.GetIsolate(context.Background(), &g2)
	// if err != nil {
	// 	log.Errorf(">> %v", err)
	// }
	//
	// b1, err := json.MarshalIndent(items, "", "    ")
	// if err != nil {
	// 	log.Errorf("error marshalling new services: %v", err)
	// }
	// fmt.Printf("%s", string(b1))

	// log.Infof("new client %v", client)

	// _, err = client.CreateIsolate(context.Background(), sr)
	// _, err = client.UpdateIsolate(context.Background(), sr)

	// SetIsolateTraffic(ctx context.Context, in *SetTrafficRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)

	ll := []*igrpc.TrafficValue{}
	r1 := "w-8829097305702293016-00006"
	var p1 int64 = 30
	v1 := &igrpc.TrafficValue{
		Revision: &r1,
		Percent:  &p1,
		// Revision *string `protobuf:"bytes,1,opt,name=revision,proto3,oneof" json:"revision,omitempty"`
		// Percent  *int64  `protobuf:"varint,2,opt,name=percent,proto3,oneof" json:"percent,omitempty"`
	}

	r2 := "w-8829097305702293016-00011"
	var p2 int64 = 70
	v2 := &igrpc.TrafficValue{
		Revision: &r2,
		Percent:  &p2,
	}
	ll = append(ll, v1)
	ll = append(ll, v2)

	r := "w-8829097305702293016"
	t := igrpc.SetTrafficRequest{
		Name:    &r,
		Traffic: ll,
	}

	_, err = client.SetIsolateTraffic(context.Background(), &t)

	log.Infof("ERR %v", err)

}
