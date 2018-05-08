package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type EnableAlarmRequest struct {
	requests.RpcRequest
	Callby_cms_owner string `position:"Query" name:"Callby_cms_owner"`
	Id               string `position:"Query" name:"Id"`
}

func (req *EnableAlarmRequest) Invoke(client *sdk.Client) (resp *EnableAlarmResponse, err error) {
	req.InitWithApiInfo("Cms", "2018-03-08", "EnableAlarm", "cms", "")
	resp = &EnableAlarmResponse{}
	err = client.DoAction(req, resp)
	return
}

type EnableAlarmResponse struct {
	responses.BaseResponse
	Success   bool
	Code      common.String
	Message   common.String
	RequestId common.String
}
