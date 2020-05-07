package telegram

import (
	"fmt"
	"github.com/aibotsoft/micro/config"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type Telegram struct {
	cfg     *config.Config
	log     *zap.SugaredLogger
	client  *resty.Client
	baseUrl string
}

func (t *Telegram) Send(msg string) {
	_, err := t.client.SetDebug(t.cfg.Telegram.Debug).R().EnableTrace().
		SetQueryParam("text", msg).
		SetQueryParam("chat_id", t.cfg.Telegram.ChatId).
		Post(t.baseUrl + "sendMessage")
	if err != nil {
		t.log.Error(err)
	}
}
func (t *Telegram) Ping() error {
	resp, err := t.client.SetDebug(t.cfg.Telegram.Debug).R().EnableTrace().Get(t.baseUrl + "getMe")
	if err != nil {
		return errors.Wrap(err, "connect to telegram error")
	}
	if resp.StatusCode() != 200 {
		return errors.Errorf("connect to telegram error")
	}
	return nil
}

func New(cfg *config.Config, log *zap.SugaredLogger) *Telegram {
	client := resty.New()
	t := &Telegram{cfg: cfg, log: log, client: client}
	t.baseUrl = fmt.Sprintf("%s/bot%s/", cfg.Telegram.Host, cfg.Telegram.Token)
	err := t.Ping()
	if err != nil {
		log.Panic(err)
	}
	return t
}
