package alidns

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpdateDNSSLBWeightRequest struct {
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	RecordId     string `position:"Query" name:"RecordId"`
	Weight       int    `position:"Query" name:"Weight"`
}

func (r UpdateDNSSLBWeightRequest) Invoke(client *sdk.Client) (response *UpdateDNSSLBWeightResponse, err error) {
	req := struct {
		*requests.RpcRequest
		UpdateDNSSLBWeightRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Alidns", "2015-01-09", "UpdateDNSSLBWeight", "", "")

	resp := struct {
		*responses.BaseResponse
		UpdateDNSSLBWeightResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.UpdateDNSSLBWeightResponse

	err = client.DoAction(&req, &resp)
	return
}

type UpdateDNSSLBWeightResponse struct {
	RequestId string
	RecordId  string
	Weight    int
}
