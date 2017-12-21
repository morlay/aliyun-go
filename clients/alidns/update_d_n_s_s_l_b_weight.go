package alidns

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpdateDNSSLBWeightRequest struct {
	requests.RpcRequest
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	RecordId     string `position:"Query" name:"RecordId"`
	Weight       int    `position:"Query" name:"Weight"`
}

func (req *UpdateDNSSLBWeightRequest) Invoke(client *sdk.Client) (resp *UpdateDNSSLBWeightResponse, err error) {
	req.InitWithApiInfo("Alidns", "2015-01-09", "UpdateDNSSLBWeight", "", "")
	resp = &UpdateDNSSLBWeightResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpdateDNSSLBWeightResponse struct {
	responses.BaseResponse
	RequestId string
	RecordId  string
	Weight    int
}
