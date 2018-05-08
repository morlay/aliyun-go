package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteAlarmRequest struct {
	requests.RpcRequest
	Callby_cms_owner string `position:"Query" name:"Callby_cms_owner"`
	Id               string `position:"Query" name:"Id"`
}

func (req *DeleteAlarmRequest) Invoke(client *sdk.Client) (resp *DeleteAlarmResponse, err error) {
	req.InitWithApiInfo("Cms", "2018-03-08", "DeleteAlarm", "cms", "")
	resp = &DeleteAlarmResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteAlarmResponse struct {
	responses.BaseResponse
	Success   bool
	Code      common.String
	Message   common.String
	RequestId common.String
}
