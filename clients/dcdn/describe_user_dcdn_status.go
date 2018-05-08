package dcdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeUserDcdnStatusRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeUserDcdnStatusRequest) Invoke(client *sdk.Client) (resp *DescribeUserDcdnStatusResponse, err error) {
	req.InitWithApiInfo("dcdn", "2018-01-15", "DescribeUserDcdnStatus", "dcdn", "")
	resp = &DescribeUserDcdnStatusResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeUserDcdnStatusResponse struct {
	responses.BaseResponse
	RequestId     string
	Enabled       bool
	OnService     bool
	InDebt        bool
	InDebtOverdue bool
}
