package cs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetTriggerHookRequest struct {
	requests.RoaRequest
	ClusterId string `position:"Path" name:"ClusterId"`
	ProjectId string `position:"Path" name:"ProjectId"`
}

func (req *GetTriggerHookRequest) Invoke(client *sdk.Client) (resp *GetTriggerHookResponse, err error) {
	req.InitWithApiInfo("CS", "2015-12-15", "GetTriggerHook", "/hook/trigger/[ClusterId]/[ProjectId]", "", "")
	req.Method = "GET"

	resp = &GetTriggerHookResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetTriggerHookResponse struct {
	responses.BaseResponse
}
