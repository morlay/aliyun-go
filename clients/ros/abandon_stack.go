package ros

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AbandonStackRequest struct {
	StackName string `position:"Path" name:"StackName"`
	StackId   string `position:"Path" name:"StackId"`
}

func (r AbandonStackRequest) Invoke(client *sdk.Client) (response *AbandonStackResponse, err error) {
	req := struct {
		*requests.RoaRequest
		AbandonStackRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("ROS", "2015-09-01", "AbandonStack", "/stacks/[StackName]/[StackId]/abandon", "", "")
	req.Method = "DELETE"

	resp := struct {
		*responses.BaseResponse
		AbandonStackResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AbandonStackResponse

	err = client.DoAction(&req, &resp)
	return
}

type AbandonStackResponse struct {
}
