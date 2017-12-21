package ros

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type InquiryStackRequest struct {
	requests.RoaRequest
}

func (req *InquiryStackRequest) Invoke(client *sdk.Client) (resp *InquiryStackResponse, err error) {
	req.InitWithApiInfo("ROS", "2015-09-01", "InquiryStack", "/stacks/inquiry", "", "")
	req.Method = "POST"

	resp = &InquiryStackResponse{}
	err = client.DoAction(req, resp)
	return
}

type InquiryStackResponse struct {
	responses.BaseResponse
}
