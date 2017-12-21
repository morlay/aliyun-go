package ros

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DoActionsRequest struct {
	requests.RoaRequest
	StackName string `position:"Path" name:"StackName"`
	StackId   string `position:"Path" name:"StackId"`
}

func (req *DoActionsRequest) Invoke(client *sdk.Client) (resp *DoActionsResponse, err error) {
	req.InitWithApiInfo("ROS", "2015-09-01", "DoActions", "/stacks/[StackName]/[StackId]/actions", "", "")
	req.Method = "POST"

	resp = &DoActionsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DoActionsResponse struct {
	responses.BaseResponse
}
