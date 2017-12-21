package ros

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type InquiryStackRequest struct {
}

func (r InquiryStackRequest) Invoke(client *sdk.Client) (response *InquiryStackResponse, err error) {
	req := struct {
		*requests.RoaRequest
		InquiryStackRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("ROS", "2015-09-01", "InquiryStack", "/stacks/inquiry", "", "")
	req.Method = "POST"

	resp := struct {
		*responses.BaseResponse
		InquiryStackResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.InquiryStackResponse

	err = client.DoAction(&req, &resp)
	return
}

type InquiryStackResponse struct {
}
