package cs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateTemplateRequest struct {
	requests.RoaRequest
}

func (req *CreateTemplateRequest) Invoke(client *sdk.Client) (resp *CreateTemplateResponse, err error) {
	req.InitWithApiInfo("CS", "2015-12-15", "CreateTemplate", "/templates", "", "")
	req.Method = "PUT"

	resp = &CreateTemplateResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateTemplateResponse struct {
	responses.BaseResponse
}
