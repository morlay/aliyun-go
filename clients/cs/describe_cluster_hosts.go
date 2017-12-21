package cs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeClusterHostsRequest struct {
	ClusterId string `position:"Path" name:"ClusterId"`
}

func (r DescribeClusterHostsRequest) Invoke(client *sdk.Client) (response *DescribeClusterHostsResponse, err error) {
	req := struct {
		*requests.RoaRequest
		DescribeClusterHostsRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("CS", "2015-12-15", "DescribeClusterHosts", "/clusters/[ClusterId]/hosts", "", "")
	req.Method = "GET"

	resp := struct {
		*responses.BaseResponse
		DescribeClusterHostsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeClusterHostsResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeClusterHostsResponse struct {
}
