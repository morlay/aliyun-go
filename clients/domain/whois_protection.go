package domain

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type WhoisProtectionRequest struct {
	requests.RpcRequest
	WhoisProtect string `position:"Query" name:"WhoisProtect"`
	DataSource   int    `position:"Query" name:"DataSource"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	DataContent  string `position:"Query" name:"DataContent"`
	Lang         string `position:"Query" name:"Lang"`
}

func (req *WhoisProtectionRequest) Invoke(client *sdk.Client) (resp *WhoisProtectionResponse, err error) {
	req.InitWithApiInfo("Domain", "2016-05-11", "WhoisProtection", "", "")
	resp = &WhoisProtectionResponse{}
	err = client.DoAction(req, resp)
	return
}

type WhoisProtectionResponse struct {
	responses.BaseResponse
	RequestId string
	Result    int
}
