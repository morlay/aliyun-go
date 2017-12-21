package ros

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DoActionsRequest struct {
	StackName string `position:"Path" name:"StackName"`
	StackId   string `position:"Path" name:"StackId"`
}

func (r DoActionsRequest) Invoke(client *sdk.Client) (response *DoActionsResponse, err error) {
	req := struct {
		*requests.RoaRequest
		DoActionsRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("ROS", "2015-09-01", "DoActions", "/stacks/[StackName]/[StackId]/actions", "", "")
	req.Method = "POST"

	resp := struct {
		*responses.BaseResponse
		DoActionsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DoActionsResponse

	err = client.DoAction(&req, &resp)
	return
}

type DoActionsResponse struct {
}
