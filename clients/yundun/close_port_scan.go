package yundun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ClosePortScanRequest struct {
	requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
}

func (req *ClosePortScanRequest) Invoke(client *sdk.Client) (resp *ClosePortScanResponse, err error) {
	req.InitWithApiInfo("Yundun", "2015-04-16", "ClosePortScan", "", "")
	resp = &ClosePortScanResponse{}
	err = client.DoAction(req, resp)
	return
}

type ClosePortScanResponse struct {
	responses.BaseResponse
	RequestId common.String
}
