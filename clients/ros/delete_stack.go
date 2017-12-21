package ros

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteStackRequest struct {
	StackName string `position:"Path" name:"StackName"`
	StackId   string `position:"Path" name:"StackId"`
}

func (r DeleteStackRequest) Invoke(client *sdk.Client) (response *DeleteStackResponse, err error) {
	req := struct {
		*requests.RoaRequest
		DeleteStackRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("ROS", "2015-09-01", "DeleteStack", "/stacks/[StackName]/[StackId]", "", "")
	req.Method = "DELETE"

	resp := struct {
		*responses.BaseResponse
		DeleteStackResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteStackResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteStackResponse struct {
}
