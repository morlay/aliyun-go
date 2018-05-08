package yundun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CloseVulScanRequest struct {
	requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
}

func (req *CloseVulScanRequest) Invoke(client *sdk.Client) (resp *CloseVulScanResponse, err error) {
	req.InitWithApiInfo("Yundun", "2015-04-16", "CloseVulScan", "", "")
	resp = &CloseVulScanResponse{}
	err = client.DoAction(req, resp)
	return
}

type CloseVulScanResponse struct {
	responses.BaseResponse
	RequestId common.String
}
