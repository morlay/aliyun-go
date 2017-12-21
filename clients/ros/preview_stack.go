package ros

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type PreviewStackRequest struct {
}

func (r PreviewStackRequest) Invoke(client *sdk.Client) (response *PreviewStackResponse, err error) {
	req := struct {
		*requests.RoaRequest
		PreviewStackRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("ROS", "2015-09-01", "PreviewStack", "/stacks/preview", "", "")
	req.Method = "POST"

	resp := struct {
		*responses.BaseResponse
		PreviewStackResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.PreviewStackResponse

	err = client.DoAction(&req, &resp)
	return
}

type PreviewStackResponse struct {
}
