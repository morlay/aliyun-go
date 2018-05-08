package alidns

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CheckDomainRecordRequest struct {
	requests.RpcRequest
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	DomainName   string `position:"Query" name:"DomainName"`
	RR           string `position:"Query" name:"RR"`
	Type         string `position:"Query" name:"Type"`
	Value        string `position:"Query" name:"Value"`
}

func (req *CheckDomainRecordRequest) Invoke(client *sdk.Client) (resp *CheckDomainRecordResponse, err error) {
	req.InitWithApiInfo("Alidns", "2015-01-09", "CheckDomainRecord", "", "")
	resp = &CheckDomainRecordResponse{}
	err = client.DoAction(req, resp)
	return
}

type CheckDomainRecordResponse struct {
	responses.BaseResponse
	RequestId common.String
	IsExist   bool
}
