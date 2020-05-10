package config_client

import (
	"context"
	pb "github.com/aibotsoft/gen/confpb"
	"github.com/aibotsoft/micro/config"
	"github.com/aibotsoft/micro/util"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"strconv"
	"time"
)

type ConfClient struct {
	cfg    *config.Config
	log    *zap.SugaredLogger
	conn   *grpc.ClientConn
	client pb.ConfClient
}

func (c *ConfClient) Close() {
	c.log.Debug("begin close conn")
	err := c.conn.Close()
	if err != nil {
		c.log.Error(err)
	}
}
func (c *ConfClient) GetGrpcAddr(ctx context.Context) (string, error) {
	got, err := c.client.GetConfig(ctx, &pb.GetConfigRequest{ServiceName: c.cfg.Service.Name})
	if err != nil {
		return "", errors.Wrapf(err, "get grpc addr error for service %q", c.cfg.Service.Name)
	}
	serviceConfig := got.GetServiceConfig()
	addr := net.JoinHostPort("", strconv.FormatInt(serviceConfig.GrpcPort, 10))
	return addr, nil
}

func (c *ConfClient) GetAccount(ctx context.Context, serviceName string) (pb.Account, error) {
	account, err := c.client.GetAccount(ctx, &pb.GetAccountRequest{ServiceName: serviceName})
	if err != nil {
		return pb.Account{}, err
	}
	return account.GetAccount(), nil
}

func New(cfg *config.Config, log *zap.SugaredLogger) *ConfClient {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	conn, err := grpc.DialContext(ctx, net.JoinHostPort("", cfg.Service.ConfigPort), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Panicw("dial to config-service error", "addr", cfg.Service.ConfigPort)
	}
	client := pb.NewConfClient(conn)
	log.Infow("begin ping to config-service", "addr", cfg.Service.ConfigPort)
	start := util.UnixMsNow()
	_, err = client.Ping(ctx, &pb.PingRequest{})
	if err != nil {
		log.Panicw("server do not response to ping", "addr", cfg.Service.ConfigPort)
	}
	log.Infow("config-serves responded to ping", "time", util.UnixMsNow()-start, "addr", cfg.Service.ConfigPort)
	return &ConfClient{cfg: cfg, log: log, client: client, conn: conn}
}
