package ros

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpdateStackRequest struct {
	requests.RoaRequest
	StackName string `position:"Path" name:"StackName"`
	StackId   string `position:"Path" name:"StackId"`
}

func (req *UpdateStackRequest) Invoke(client *sdk.Client) (resp *UpdateStackResponse, err error) {
	req.InitWithApiInfo("ROS", "2015-09-01", "UpdateStack", "/stacks/[StackName]/[StackId]", "", "")
	req.Method = "PUT"

	resp = &UpdateStackResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpdateStackResponse struct {
	responses.BaseResponse
}
