package alidns

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetDNSSLBStatusRequest struct {
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	SubDomain    string `position:"Query" name:"SubDomain"`
	Open         string `position:"Query" name:"Open"`
}

func (r SetDNSSLBStatusRequest) Invoke(client *sdk.Client) (response *SetDNSSLBStatusResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SetDNSSLBStatusRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Alidns", "2015-01-09", "SetDNSSLBStatus", "", "")

	resp := struct {
		*responses.BaseResponse
		SetDNSSLBStatusResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SetDNSSLBStatusResponse

	err = client.DoAction(&req, &resp)
	return
}

type SetDNSSLBStatusResponse struct {
	RequestId   string
	RecordCount int64
	Open        bool
}
