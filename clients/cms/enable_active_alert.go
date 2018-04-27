package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type EnableActiveAlertRequest struct {
	requests.RpcRequest
	Product string `position:"Query" name:"Product"`
	UserId  string `position:"Query" name:"UserId"`
}

func (req *EnableActiveAlertRequest) Invoke(client *sdk.Client) (resp *EnableActiveAlertResponse, err error) {
	req.InitWithApiInfo("Cms", "2018-03-08", "EnableActiveAlert", "cms", "")
	resp = &EnableActiveAlertResponse{}
	err = client.DoAction(req, resp)
	return
}

type EnableActiveAlertResponse struct {
	responses.BaseResponse
	Success bool
	Code    string
	Message string
}
