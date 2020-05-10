package config_client

import (
	"context"
	pb "github.com/aibotsoft/gen/fortedpb"
	"github.com/aibotsoft/micro/config"
	"github.com/aibotsoft/micro/util"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"time"
)

func NewConfigClient(cfg *config.Config, log *zap.SugaredLogger) pb.FortedClient {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	conn, err := grpc.DialContext(ctx, net.JoinHostPort("", cfg.Service.ConfigPort), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Panicw("dial to config-service error", "addr", cfg.Service.ConfigPort)
	}
	client := pb.NewFortedClient(conn)
	log.Infow("begin ping to config-service", "addr", cfg.Service.ConfigPort)
	start := util.UnixMsNow()
	_, err = client.Ping(ctx, &pb.PingRequest{})
	if err != nil {
		log.Panicw("server do not response to ping", "addr", cfg.Service.ConfigPort)
	}
	log.Infow("config-serves responded to ping", "time", util.UnixMsNow()-start, "addr", cfg.Service.ConfigPort)
	return client
}
