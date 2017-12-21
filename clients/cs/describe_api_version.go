package cs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeApiVersionRequest struct {
}

func (r DescribeApiVersionRequest) Invoke(client *sdk.Client) (response *DescribeApiVersionResponse, err error) {
	req := struct {
		*requests.RoaRequest
		DescribeApiVersionRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("CS", "2015-12-15", "DescribeApiVersion", "/version", "", "")
	req.Method = "GET"

	resp := struct {
		*responses.BaseResponse
		DescribeApiVersionResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeApiVersionResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeApiVersionResponse struct {
}
