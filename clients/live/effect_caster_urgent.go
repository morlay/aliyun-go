package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type EffectCasterUrgentRequest struct {
	requests.RpcRequest
	CasterId string `position:"Query" name:"CasterId"`
	SceneId  string `position:"Query" name:"SceneId"`
	OwnerId  int64  `position:"Query" name:"OwnerId"`
}

func (req *EffectCasterUrgentRequest) Invoke(client *sdk.Client) (resp *EffectCasterUrgentResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "EffectCasterUrgent", "live", "")
	resp = &EffectCasterUrgentResponse{}
	err = client.DoAction(req, resp)
	return
}

type EffectCasterUrgentResponse struct {
	responses.BaseResponse
	RequestId common.String
}
