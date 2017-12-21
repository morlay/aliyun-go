package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DisableAlarmRequest struct {
	requests.RpcRequest
	Callby_cms_owner string `position:"Query" name:"Callby_cms_owner"`
	Id               string `position:"Query" name:"Id"`
}

func (req *DisableAlarmRequest) Invoke(client *sdk.Client) (resp *DisableAlarmResponse, err error) {
	req.InitWithApiInfo("Cms", "2017-03-01", "DisableAlarm", "cms", "")
	resp = &DisableAlarmResponse{}
	err = client.DoAction(req, resp)
	return
}

type DisableAlarmResponse struct {
	responses.BaseResponse
	Success   bool
	Code      string
	Message   string
	RequestId string
}
