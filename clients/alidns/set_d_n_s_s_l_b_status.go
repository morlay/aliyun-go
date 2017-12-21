package alidns

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetDNSSLBStatusRequest struct {
	requests.RpcRequest
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	SubDomain    string `position:"Query" name:"SubDomain"`
	Open         string `position:"Query" name:"Open"`
}

func (req *SetDNSSLBStatusRequest) Invoke(client *sdk.Client) (resp *SetDNSSLBStatusResponse, err error) {
	req.InitWithApiInfo("Alidns", "2015-01-09", "SetDNSSLBStatus", "", "")
	resp = &SetDNSSLBStatusResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetDNSSLBStatusResponse struct {
	responses.BaseResponse
	RequestId   string
	RecordCount int64
	Open        bool
}
