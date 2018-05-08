package yundun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type OpenPortScanRequest struct {
	requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
}

func (req *OpenPortScanRequest) Invoke(client *sdk.Client) (resp *OpenPortScanResponse, err error) {
	req.InitWithApiInfo("Yundun", "2015-04-16", "OpenPortScan", "", "")
	resp = &OpenPortScanResponse{}
	err = client.DoAction(req, resp)
	return
}

type OpenPortScanResponse struct {
	responses.BaseResponse
	RequestId common.String
}
