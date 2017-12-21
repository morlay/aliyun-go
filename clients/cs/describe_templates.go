package cs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeTemplatesRequest struct {
	requests.RoaRequest
}

func (req *DescribeTemplatesRequest) Invoke(client *sdk.Client) (resp *DescribeTemplatesResponse, err error) {
	req.InitWithApiInfo("CS", "2015-12-15", "DescribeTemplates", "/templates", "", "")
	req.Method = "GET"

	resp = &DescribeTemplatesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeTemplatesResponse struct {
	responses.BaseResponse
}
