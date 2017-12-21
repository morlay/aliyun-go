package ros

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateStacksRequest struct {
}

func (r CreateStacksRequest) Invoke(client *sdk.Client) (response *CreateStacksResponse, err error) {
	req := struct {
		*requests.RoaRequest
		CreateStacksRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("ROS", "2015-09-01", "CreateStacks", "/stacks", "", "")
	req.Method = "POST"

	resp := struct {
		*responses.BaseResponse
		CreateStacksResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateStacksResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateStacksResponse struct {
}
