package domain

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CheckDomainRequest struct {
	DomainName string `position:"Query" name:"DomainName"`
}

func (r CheckDomainRequest) Invoke(client *sdk.Client) (response *CheckDomainResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CheckDomainRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Domain", "2016-05-11", "CheckDomain", "", "")

	resp := struct {
		*responses.BaseResponse
		CheckDomainResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CheckDomainResponse

	err = client.DoAction(&req, &resp)
	return
}

type CheckDomainResponse struct {
	RequestId   string
	Name        string
	Avail       int
	Reason      string
	FeeCurrency string
	FeePeriod   int
	FeeFee      string
	RmbFee      string
	FeeCommand  string
}
