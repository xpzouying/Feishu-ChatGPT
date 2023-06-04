package service

import (
	"context"
	"runtime/debug"

	"github.com/sirupsen/logrus"
	"github.com/xpzouying/Feishu-ChatGPT/internal/domain"

	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
)

type FeishuService struct {
	feishu *domain.FeishuEx
}

func NewFeishuService(feishu *domain.FeishuEx) *FeishuService {
	return &FeishuService{
		feishu: feishu,
	}
}

func (s *FeishuService) HandleMessageReceive(ctx context.Context, receive *larkim.P2MessageReceiveV1) error {
	defer func() {
		if r := recover(); r != nil {
			logrus.Errorf("HandleMessageReceive: recover: %v\n%s", r, debug.Stack())
		}

	}()

	larkMessage := (*domain.LarkMessage)(receive.Event.Message)

	return s.feishu.HandleMessageReceive(ctx, larkMessage)
}
