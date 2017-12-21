package bss

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeCashDetailRequest struct {
	requests.RpcRequest
}

func (req *DescribeCashDetailRequest) Invoke(client *sdk.Client) (resp *DescribeCashDetailResponse, err error) {
	req.InitWithApiInfo("Bss", "2014-07-14", "DescribeCashDetail", "", "")
	resp = &DescribeCashDetailResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeCashDetailResponse struct {
	responses.BaseResponse
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
