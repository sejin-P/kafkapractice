package main

import (
	"context"
	"github.com/sejin-P/kafkapractice/consumer"
	"github.com/sejin-P/kafkapractice/protos/v1/empty"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"
)

const portNum = "9001"

type EmptyServer struct {
	empty.EmptyServer

	Consumer consumer.Consumer
}

func NewEmptyServer(consumer consumer.Consumer) *EmptyServer {
	return &EmptyServer{Consumer: consumer}
}

func (s *EmptyServer) GetData(_ context.Context, _ *empty.GetDataRequest) (*empty.GetDataResponse, error) {
	return &empty.GetDataResponse{}, nil
}

func main() {
	emptyServer := NewEmptyServer(nil)

	grpcServer := grpc.NewServer()
	empty.RegisterEmptyServer(grpcServer, emptyServer)
	reflection.Register(grpcServer)

	go func() {
		lis, err := net.Listen("tcp", ":"+portNum)
		if err != nil {
			log.Panic(err)
		}

		logrus.WithField("port", portNum).Info("starting consumer gRPC server")
		if err := grpcServer.Serve(lis); err != nil && err != grpc.ErrServerStopped {
			logrus.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(quit)

	<-quit

	time.Sleep(time.Duration(30000) * time.Millisecond)
	logrus.Info("Stopping consumer gRPC server")
	grpcServer.GracefulStop()
}
