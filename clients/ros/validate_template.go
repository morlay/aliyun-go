package ros

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ValidateTemplateRequest struct {
	requests.RoaRequest
}

func (req *ValidateTemplateRequest) Invoke(client *sdk.Client) (resp *ValidateTemplateResponse, err error) {
	req.InitWithApiInfo("ROS", "2015-09-01", "ValidateTemplate", "/validate", "", "")
	req.Method = "POST"

	resp = &ValidateTemplateResponse{}
	err = client.DoAction(req, resp)
	return
}

type ValidateTemplateResponse struct {
	responses.BaseResponse
}
