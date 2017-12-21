package cs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeClusterServicesRequest struct {
	ClusterId string `position:"Path" name:"ClusterId"`
}

func (r DescribeClusterServicesRequest) Invoke(client *sdk.Client) (response *DescribeClusterServicesResponse, err error) {
	req := struct {
		*requests.RoaRequest
		DescribeClusterServicesRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("CS", "2015-12-15", "DescribeClusterServices", "/clusters/[ClusterId]/services", "", "")
	req.Method = "GET"

	resp := struct {
		*responses.BaseResponse
		DescribeClusterServicesResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeClusterServicesResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeClusterServicesResponse struct {
}
