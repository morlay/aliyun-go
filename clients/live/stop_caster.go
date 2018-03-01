package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type StopCasterRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	CasterId      string `position:"Query" name:"CasterId"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	Version       string `position:"Query" name:"Version"`
}

func (req *StopCasterRequest) Invoke(client *sdk.Client) (resp *StopCasterResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "StopCaster", "live", "")
	resp = &StopCasterResponse{}
	err = client.DoAction(req, resp)
	return
}

type StopCasterResponse struct {
	responses.BaseResponse
	RequestId string
}
