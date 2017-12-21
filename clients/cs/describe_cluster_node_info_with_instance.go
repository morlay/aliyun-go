package cs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeClusterNodeInfoWithInstanceRequest struct {
	requests.RoaRequest
	Token      string `position:"Path" name:"Token"`
	InstanceId string `position:"Path" name:"InstanceId"`
}

func (req *DescribeClusterNodeInfoWithInstanceRequest) Invoke(client *sdk.Client) (resp *DescribeClusterNodeInfoWithInstanceResponse, err error) {
	req.InitWithApiInfo("CS", "2015-12-15", "DescribeClusterNodeInfoWithInstance", "/token/[Token]/instance/[InstanceId]/node_info", "", "")
	req.Method = "GET"

	resp = &DescribeClusterNodeInfoWithInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeClusterNodeInfoWithInstanceResponse struct {
	responses.BaseResponse
}
