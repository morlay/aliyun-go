package cs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeClusterCertsRequest struct {
	ClusterId string `position:"Path" name:"ClusterId"`
}

func (r DescribeClusterCertsRequest) Invoke(client *sdk.Client) (response *DescribeClusterCertsResponse, err error) {
	req := struct {
		*requests.RoaRequest
		DescribeClusterCertsRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("CS", "2015-12-15", "DescribeClusterCerts", "/clusters/[ClusterId]/certs", "", "")
	req.Method = "GET"

	resp := struct {
		*responses.BaseResponse
		DescribeClusterCertsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeClusterCertsResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeClusterCertsResponse struct {
}
