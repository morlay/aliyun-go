package bss

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId            common.String
	BalanceAmount        common.String
	AmountOwed           common.String
	EnableThresholdAlert common.String
	MiniAlertThreshold   common.Long
	FrozenAmount         common.String
	CreditCardAmount     common.String
	RemmitanceAmount     common.String
	CreditLimit          common.String
	AvailableCredit      common.String
}
