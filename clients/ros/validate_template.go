package ros

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ValidateTemplateRequest struct {
}

func (r ValidateTemplateRequest) Invoke(client *sdk.Client) (response *ValidateTemplateResponse, err error) {
	req := struct {
		*requests.RoaRequest
		ValidateTemplateRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("ROS", "2015-09-01", "ValidateTemplate", "/validate", "", "")
	req.Method = "POST"

	resp := struct {
		*responses.BaseResponse
		ValidateTemplateResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ValidateTemplateResponse

	err = client.DoAction(&req, &resp)
	return
}

type ValidateTemplateResponse struct {
}
