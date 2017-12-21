package domain

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type WhoisProtectionRequest struct {
	WhoisProtect string `position:"Query" name:"WhoisProtect"`
	DataSource   int    `position:"Query" name:"DataSource"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	DataContent  string `position:"Query" name:"DataContent"`
	Lang         string `position:"Query" name:"Lang"`
}

func (r WhoisProtectionRequest) Invoke(client *sdk.Client) (response *WhoisProtectionResponse, err error) {
	req := struct {
		*requests.RpcRequest
		WhoisProtectionRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Domain", "2016-05-11", "WhoisProtection", "", "")

	resp := struct {
		*responses.BaseResponse
		WhoisProtectionResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.WhoisProtectionResponse

	err = client.DoAction(&req, &resp)
	return
}

type WhoisProtectionResponse struct {
	RequestId string
	Result    int
}
