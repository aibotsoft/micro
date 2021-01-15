package config_client

import (
	"context"
	pb "github.com/aibotsoft/gen/confpb"
	"github.com/aibotsoft/micro/config"
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

func New(cfg *config.Config, log *zap.SugaredLogger) *ConfClient {
	return &ConfClient{cfg: cfg, log: log}
}
func (c *ConfClient) GetClient() pb.ConfClient {
	if c.client == nil {
		err := c.Connect()
		if err != nil {
			c.log.Error(err)
		}
	}
	return c.client
}
func (c *ConfClient) Ping(ctx context.Context) error {
	_, err := c.GetClient().Ping(ctx, &pb.PingRequest{})
	if err != nil {
		return errors.Wrap(err, "ping_error")
	}
	return nil
}

func (c *ConfClient) Connect() (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	hostPort := net.JoinHostPort(c.cfg.Service.ConfigHost, c.cfg.Service.ConfigPort)

	c.log.Infow("begin connect to config service", "")
	c.conn, err = grpc.DialContext(ctx, hostPort, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return errors.Wrapf(err, "dial_config_service_error, port: %q", hostPort)
	}
	c.client = pb.NewConfClient(c.conn)
	return nil
}
func (c *ConfClient) Close() {
	c.log.Debug("begin close conn")
	if c.conn == nil {
		return
	}
	err := c.conn.Close()
	if err != nil {
		c.log.Error(err)
	}
}
func (c *ConfClient) GetNetStatus(ctx context.Context) bool {
	resp, err := c.GetClient().GetNetStatus(ctx, &pb.GetNetStatusRequest{})
	if err != nil {
		c.log.Info(err)
		return false
	}
	return resp.GetStatus()
}

func (c *ConfClient) GetServices(ctx context.Context) ([]pb.BetService, error) {
	resp, err := c.GetClient().GetServices(ctx, &pb.GetServicesRequest{})
	if err != nil {
		return nil, errors.Wrap(err, "get services error")
	}
	return resp.GetServices(), nil
}
func (c *ConfClient) GetCurrency(ctx context.Context) ([]pb.Currency, error) {
	resp, err := c.GetClient().GetCurrency(ctx, &pb.GetCurrencyRequest{})
	if err != nil {
		return nil, errors.Wrap(err, "get currency error")
	}
	return resp.GetCurrencyList(), nil
}

func (c *ConfClient) GetGrpcAddr(ctx context.Context, serviceName string) (string, error) {
	got, err := c.GetClient().GetConfig(ctx, &pb.GetConfigRequest{ServiceName: serviceName})
	if err != nil {
		return "", errors.Wrapf(err, "get grpc addr error for service %q", c.cfg.Service.Name)
	}
	serviceConfig := got.GetServiceConfig()
	addr := net.JoinHostPort("", strconv.FormatInt(serviceConfig.GrpcPort, 10))
	return addr, nil
}
func (c *ConfClient) GetAccount(ctx context.Context, serviceName string) (pb.Account, error) {
	account, err := c.GetClient().GetAccount(ctx, &pb.GetAccountRequest{ServiceName: serviceName})
	if err != nil {
		return pb.Account{}, err
	}
	return account.GetAccount(), nil
}
