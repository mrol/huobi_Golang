package market

import (
	"github.com/mrol/huobi_golang/pkg/model/base"
)

type SubscribeDepthResponse struct {
	base.WebSocketResponseBase
	Data *Depth
	Tick *Depth
}
