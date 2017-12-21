package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteAlarmRequest struct {
	requests.RpcRequest
	Callby_cms_owner string `position:"Query" name:"Callby_cms_owner"`
	Id               string `position:"Query" name:"Id"`
}

func (req *DeleteAlarmRequest) Invoke(client *sdk.Client) (resp *DeleteAlarmResponse, err error) {
	req.InitWithApiInfo("Cms", "2017-03-01", "DeleteAlarm", "cms", "")
	resp = &DeleteAlarmResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteAlarmResponse struct {
	responses.BaseResponse
	Success   bool
	Code      string
	Message   string
	RequestId string
}
