package httpdns

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetAccountInfoRequest struct {
	requests.RpcRequest
}

func (req *GetAccountInfoRequest) Invoke(client *sdk.Client) (resp *GetAccountInfoResponse, err error) {
	req.InitWithApiInfo("Httpdns", "2016-02-01", "GetAccountInfo", "", "")
	resp = &GetAccountInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetAccountInfoResponse struct {
	responses.BaseResponse
	RequestId   common.String
	AccountInfo GetAccountInfoAccountInfo
}

type GetAccountInfoAccountInfo struct {
	AccountId              common.String
	MonthFreeCount         common.Integer
	MonthHttpsResolveCount common.Integer
	MonthResolveCount      common.Integer
	PackageCount           common.Integer
	UserStatus             common.Integer
	SignSecret             common.String
	UnsignedEnabled        bool
	SignedCount            common.Long
	UnsignedCount          common.Long
}
