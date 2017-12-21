package cs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeClusterNodeInfoWithInstanceRequest struct {
	Token      string `position:"Path" name:"Token"`
	InstanceId string `position:"Path" name:"InstanceId"`
}

func (r DescribeClusterNodeInfoWithInstanceRequest) Invoke(client *sdk.Client) (response *DescribeClusterNodeInfoWithInstanceResponse, err error) {
	req := struct {
		*requests.RoaRequest
		DescribeClusterNodeInfoWithInstanceRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("CS", "2015-12-15", "DescribeClusterNodeInfoWithInstance", "/token/[Token]/instance/[InstanceId]/node_info", "", "")
	req.Method = "GET"

	resp := struct {
		*responses.BaseResponse
		DescribeClusterNodeInfoWithInstanceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeClusterNodeInfoWithInstanceResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeClusterNodeInfoWithInstanceResponse struct {
}
