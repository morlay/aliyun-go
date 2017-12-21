package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetProbeDataSubscriberConfigRequest struct {
	requests.RpcRequest
	Id int64 `position:"Query" name:"Id"`
}

func (req *GetProbeDataSubscriberConfigRequest) Invoke(client *sdk.Client) (resp *GetProbeDataSubscriberConfigResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "GetProbeDataSubscriberConfig", "", "")
	resp = &GetProbeDataSubscriberConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetProbeDataSubscriberConfigResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
