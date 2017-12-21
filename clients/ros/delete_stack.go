package ros

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteStackRequest struct {
	requests.RoaRequest
	StackName string `position:"Path" name:"StackName"`
	StackId   string `position:"Path" name:"StackId"`
}

func (req *DeleteStackRequest) Invoke(client *sdk.Client) (resp *DeleteStackResponse, err error) {
	req.InitWithApiInfo("ROS", "2015-09-01", "DeleteStack", "/stacks/[StackName]/[StackId]", "", "")
	req.Method = "DELETE"

	resp = &DeleteStackResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteStackResponse struct {
	responses.BaseResponse
}
