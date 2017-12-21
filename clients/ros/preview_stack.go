package ros

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type PreviewStackRequest struct {
	requests.RoaRequest
}

func (req *PreviewStackRequest) Invoke(client *sdk.Client) (resp *PreviewStackResponse, err error) {
	req.InitWithApiInfo("ROS", "2015-09-01", "PreviewStack", "/stacks/preview", "", "")
	req.Method = "POST"

	resp = &PreviewStackResponse{}
	err = client.DoAction(req, resp)
	return
}

type PreviewStackResponse struct {
	responses.BaseResponse
}
