package ros

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateStacksRequest struct {
	requests.RoaRequest
}

func (req *CreateStacksRequest) Invoke(client *sdk.Client) (resp *CreateStacksResponse, err error) {
	req.InitWithApiInfo("ROS", "2015-09-01", "CreateStacks", "/stacks", "", "")
	req.Method = "POST"

	resp = &CreateStacksResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateStacksResponse struct {
	responses.BaseResponse
}
