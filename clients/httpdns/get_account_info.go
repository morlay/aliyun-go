package httpdns

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId   string
	AccountInfo GetAccountInfoAccountInfo
}

type GetAccountInfoAccountInfo struct {
	AccountId              string
	MonthFreeCount         int
	MonthHttpsResolveCount int
	MonthResolveCount      int
	PackageCount           int
	UserStatus             int
	SignSecret             string
	UnsignedEnabled        bool
	SignedCount            int64
	UnsignedCount          int64
}
