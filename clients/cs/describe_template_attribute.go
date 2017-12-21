package cs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeTemplateAttributeRequest struct {
	TemplateId string `position:"Path" name:"TemplateId"`
}

func (r DescribeTemplateAttributeRequest) Invoke(client *sdk.Client) (response *DescribeTemplateAttributeResponse, err error) {
	req := struct {
		*requests.RoaRequest
		DescribeTemplateAttributeRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("CS", "2015-12-15", "DescribeTemplateAttribute", "/templates/[TemplateId]", "", "")
	req.Method = "GET"

	resp := struct {
		*responses.BaseResponse
		DescribeTemplateAttributeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeTemplateAttributeResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeTemplateAttributeResponse struct {
}
