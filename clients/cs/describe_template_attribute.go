package cs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeTemplateAttributeRequest struct {
	requests.RoaRequest
	TemplateId string `position:"Path" name:"TemplateId"`
}

func (req *DescribeTemplateAttributeRequest) Invoke(client *sdk.Client) (resp *DescribeTemplateAttributeResponse, err error) {
	req.InitWithApiInfo("CS", "2015-12-15", "DescribeTemplateAttribute", "/templates/[TemplateId]", "", "")
	req.Method = "GET"

	resp = &DescribeTemplateAttributeResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeTemplateAttributeResponse struct {
	responses.BaseResponse
}
