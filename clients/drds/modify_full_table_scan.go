package drds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyFullTableScanRequest struct {
	requests.RpcRequest
	DbName         string `position:"Query" name:"DbName"`
	TableNames     string `position:"Query" name:"TableNames"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
	FullTableScan  string `position:"Query" name:"FullTableScan"`
}

func (req *ModifyFullTableScanRequest) Invoke(client *sdk.Client) (resp *ModifyFullTableScanResponse, err error) {
	req.InitWithApiInfo("Drds", "2017-10-16", "ModifyFullTableScan", "", "")
	resp = &ModifyFullTableScanResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyFullTableScanResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
}
