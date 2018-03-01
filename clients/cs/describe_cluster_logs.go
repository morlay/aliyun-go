package cs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeClusterLogsRequest struct {
	requests.RoaRequest
	ClusterId string `position:"Path" name:"ClusterId"`
}

func (req *DescribeClusterLogsRequest) Invoke(client *sdk.Client) (resp *DescribeClusterLogsResponse, err error) {
	req.InitWithApiInfo("CS", "2015-12-15", "DescribeClusterLogs", "/clusters/[ClusterId]/logs", "", "")
	req.Method = "GET"

	resp = &DescribeClusterLogsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeClusterLogsResponse struct {
	responses.BaseResponse
}
