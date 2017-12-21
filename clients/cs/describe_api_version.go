package cs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeApiVersionRequest struct {
	requests.RoaRequest
}

func (req *DescribeApiVersionRequest) Invoke(client *sdk.Client) (resp *DescribeApiVersionResponse, err error) {
	req.InitWithApiInfo("CS", "2015-12-15", "DescribeApiVersion", "/version", "", "")
	req.Method = "GET"

	resp = &DescribeApiVersionResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeApiVersionResponse struct {
	responses.BaseResponse
}
