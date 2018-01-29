package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DisableActiveAlertRequest struct {
	requests.RpcRequest
	Product string `position:"Query" name:"Product"`
	UserId  string `position:"Query" name:"UserId"`
}

func (req *DisableActiveAlertRequest) Invoke(client *sdk.Client) (resp *DisableActiveAlertResponse, err error) {
	req.InitWithApiInfo("Cms", "2017-03-01", "DisableActiveAlert", "cms", "")
	resp = &DisableActiveAlertResponse{}
	err = client.DoAction(req, resp)
	return
}

type DisableActiveAlertResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Code      string
	Message   string
}
