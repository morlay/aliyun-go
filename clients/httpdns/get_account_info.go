package httpdns

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetAccountInfoRequest struct {
}

func (r GetAccountInfoRequest) Invoke(client *sdk.Client) (response *GetAccountInfoResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetAccountInfoRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Httpdns", "2016-02-01", "GetAccountInfo", "", "")

	resp := struct {
		*responses.BaseResponse
		GetAccountInfoResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetAccountInfoResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetAccountInfoResponse struct {
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
