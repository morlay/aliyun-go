package bss

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeCashDetailRequest struct {
}

func (r DescribeCashDetailRequest) Invoke(client *sdk.Client) (response *DescribeCashDetailResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeCashDetailRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Bss", "2014-07-14", "DescribeCashDetail", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeCashDetailResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeCashDetailResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeCashDetailResponse struct {
	RequestId            string
	BalanceAmount        string
	AmountOwed           string
	EnableThresholdAlert string
	MiniAlertThreshold   int64
	FrozenAmount         string
	CreditCardAmount     string
	RemmitanceAmount     string
	CreditLimit          string
	AvailableCredit      string
}
