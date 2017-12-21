package drds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyFullTableScanRequest struct {
	DbName         string `position:"Query" name:"DbName"`
	TableNames     string `position:"Query" name:"TableNames"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
	FullTableScan  string `position:"Query" name:"FullTableScan"`
}

func (r ModifyFullTableScanRequest) Invoke(client *sdk.Client) (response *ModifyFullTableScanResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyFullTableScanRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Drds", "2017-10-16", "ModifyFullTableScan", "", "")

	resp := struct {
		*responses.BaseResponse
		ModifyFullTableScanResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyFullTableScanResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyFullTableScanResponse struct {
	RequestId string
	Success   bool
}
