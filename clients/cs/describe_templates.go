package cs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeTemplatesRequest struct {
}

func (r DescribeTemplatesRequest) Invoke(client *sdk.Client) (response *DescribeTemplatesResponse, err error) {
	req := struct {
		*requests.RoaRequest
		DescribeTemplatesRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("CS", "2015-12-15", "DescribeTemplates", "/templates", "", "")
	req.Method = "GET"

	resp := struct {
		*responses.BaseResponse
		DescribeTemplatesResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeTemplatesResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeTemplatesResponse struct {
}
