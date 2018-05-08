package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type StopCasterRequest struct {
	requests.RpcRequest
	CasterId string `position:"Query" name:"CasterId"`
	OwnerId  int64  `position:"Query" name:"OwnerId"`
}

func (req *StopCasterRequest) Invoke(client *sdk.Client) (resp *StopCasterResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "StopCaster", "live", "")
	resp = &StopCasterResponse{}
	err = client.DoAction(req, resp)
	return
}

type StopCasterResponse struct {
	responses.BaseResponse
	RequestId common.String
}
