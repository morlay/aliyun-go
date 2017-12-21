package cs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateTemplateRequest struct {
}

func (r CreateTemplateRequest) Invoke(client *sdk.Client) (response *CreateTemplateResponse, err error) {
	req := struct {
		*requests.RoaRequest
		CreateTemplateRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("CS", "2015-12-15", "CreateTemplate", "/templates", "", "")
	req.Method = "PUT"

	resp := struct {
		*responses.BaseResponse
		CreateTemplateResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateTemplateResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateTemplateResponse struct {
}
