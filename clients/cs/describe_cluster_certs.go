package cs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeClusterCertsRequest struct {
	requests.RoaRequest
	ClusterId string `position:"Path" name:"ClusterId"`
}

func (req *DescribeClusterCertsRequest) Invoke(client *sdk.Client) (resp *DescribeClusterCertsResponse, err error) {
	req.InitWithApiInfo("CS", "2015-12-15", "DescribeClusterCerts", "/clusters/[ClusterId]/certs", "", "")
	req.Method = "GET"

	resp = &DescribeClusterCertsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeClusterCertsResponse struct {
	responses.BaseResponse
}
