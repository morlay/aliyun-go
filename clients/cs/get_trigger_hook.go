package cs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetTriggerHookRequest struct {
	ClusterId string `position:"Path" name:"ClusterId"`
	ProjectId string `position:"Path" name:"ProjectId"`
}

func (r GetTriggerHookRequest) Invoke(client *sdk.Client) (response *GetTriggerHookResponse, err error) {
	req := struct {
		*requests.RoaRequest
		GetTriggerHookRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("CS", "2015-12-15", "GetTriggerHook", "/hook/trigger/[ClusterId]/[ProjectId]", "", "")
	req.Method = "GET"

	resp := struct {
		*responses.BaseResponse
		GetTriggerHookResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetTriggerHookResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetTriggerHookResponse struct {
}
