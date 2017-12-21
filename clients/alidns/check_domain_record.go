package alidns

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CheckDomainRecordRequest struct {
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	DomainName   string `position:"Query" name:"DomainName"`
	RR           string `position:"Query" name:"RR"`
	Type         string `position:"Query" name:"Type"`
	Value        string `position:"Query" name:"Value"`
}

func (r CheckDomainRecordRequest) Invoke(client *sdk.Client) (response *CheckDomainRecordResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CheckDomainRecordRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Alidns", "2015-01-09", "CheckDomainRecord", "", "")

	resp := struct {
		*responses.BaseResponse
		CheckDomainRecordResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CheckDomainRecordResponse

	err = client.DoAction(&req, &resp)
	return
}

type CheckDomainRecordResponse struct {
	RequestId string
	IsExist   bool
}
