package ms

import (
	"log"

	pb "github.com/hi20160616/fetchnews/api/fetchnews/web/v1"
	"github.com/hi20160616/fetchnews/config"
	"google.golang.org/grpc"
)

type Conn struct {
	MicroService config.MicroService
	ConnClient   *grpc.ClientConn
	FetchClient  pb.FetchnewsWebClient
}

var Conns = map[string]*Conn{}

func init() {
	if err := initPool(); err != nil {
		log.Printf("ms initPool error: %#v", err)
	}
}

func initPool() error {
	for _, v := range config.Data.MS {
		conn, err := grpc.Dial(v.Addr, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			return err
		}
		defer conn.Close()
		Conns[v.Title] = &Conn{
			MicroService: v,
			ConnClient:   conn,
			FetchClient:  pb.NewFetchnewsWebClient(conn),
		}
	}
	return nil
}

func Close() error {
	for _, c := range Conns {
		if err := c.ConnClient.Close(); err != nil {
			return err
		}
	}
	return nil
}
