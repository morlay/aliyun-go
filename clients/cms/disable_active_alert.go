package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DisableActiveAlertRequest struct {
	requests.RpcRequest
	Product string `position:"Query" name:"Product"`
	UserId  string `position:"Query" name:"UserId"`
}

func (req *DisableActiveAlertRequest) Invoke(client *sdk.Client) (resp *DisableActiveAlertResponse, err error) {
	req.InitWithApiInfo("Cms", "2018-03-08", "DisableActiveAlert", "cms", "")
	resp = &DisableActiveAlertResponse{}
	err = client.DoAction(req, resp)
	return
}

type DisableActiveAlertResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
	Code      common.String
	Message   common.String
}
