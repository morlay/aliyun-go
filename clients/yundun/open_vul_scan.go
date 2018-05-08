package yundun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type OpenVulScanRequest struct {
	requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
}

func (req *OpenVulScanRequest) Invoke(client *sdk.Client) (resp *OpenVulScanResponse, err error) {
	req.InitWithApiInfo("Yundun", "2015-04-16", "OpenVulScan", "", "")
	resp = &OpenVulScanResponse{}
	err = client.DoAction(req, resp)
	return
}

type OpenVulScanResponse struct {
	responses.BaseResponse
	RequestId common.String
}
