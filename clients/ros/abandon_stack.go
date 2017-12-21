package ros

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AbandonStackRequest struct {
	requests.RoaRequest
	StackName string `position:"Path" name:"StackName"`
	StackId   string `position:"Path" name:"StackId"`
}

func (req *AbandonStackRequest) Invoke(client *sdk.Client) (resp *AbandonStackResponse, err error) {
	req.InitWithApiInfo("ROS", "2015-09-01", "AbandonStack", "/stacks/[StackName]/[StackId]/abandon", "", "")
	req.Method = "DELETE"

	resp = &AbandonStackResponse{}
	err = client.DoAction(req, resp)
	return
}

type AbandonStackResponse struct {
	responses.BaseResponse
}
