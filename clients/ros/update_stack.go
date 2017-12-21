package ros

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpdateStackRequest struct {
	StackName string `position:"Path" name:"StackName"`
	StackId   string `position:"Path" name:"StackId"`
}

func (r UpdateStackRequest) Invoke(client *sdk.Client) (response *UpdateStackResponse, err error) {
	req := struct {
		*requests.RoaRequest
		UpdateStackRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("ROS", "2015-09-01", "UpdateStack", "/stacks/[StackName]/[StackId]", "", "")
	req.Method = "PUT"

	resp := struct {
		*responses.BaseResponse
		UpdateStackResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.UpdateStackResponse

	err = client.DoAction(&req, &resp)
	return
}

type UpdateStackResponse struct {
}
