package cs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeClusterNodeInfoRequest struct {
	Token string `position:"Path" name:"Token"`
}

func (r DescribeClusterNodeInfoRequest) Invoke(client *sdk.Client) (response *DescribeClusterNodeInfoResponse, err error) {
	req := struct {
		*requests.RoaRequest
		DescribeClusterNodeInfoRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("CS", "2015-12-15", "DescribeClusterNodeInfo", "/token/[Token]/node_info", "", "")
	req.Method = "GET"

	resp := struct {
		*responses.BaseResponse
		DescribeClusterNodeInfoResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeClusterNodeInfoResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeClusterNodeInfoResponse struct {
}
