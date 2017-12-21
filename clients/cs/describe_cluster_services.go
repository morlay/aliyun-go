package cs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeClusterServicesRequest struct {
	requests.RoaRequest
	ClusterId string `position:"Path" name:"ClusterId"`
}

func (req *DescribeClusterServicesRequest) Invoke(client *sdk.Client) (resp *DescribeClusterServicesResponse, err error) {
	req.InitWithApiInfo("CS", "2015-12-15", "DescribeClusterServices", "/clusters/[ClusterId]/services", "", "")
	req.Method = "GET"

	resp = &DescribeClusterServicesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeClusterServicesResponse struct {
	responses.BaseResponse
}
