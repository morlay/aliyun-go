package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListApDetailInfoRequest struct {
	ApAssetId int64 `position:"Query" name:"ApAssetId"`
}

func (r ListApDetailInfoRequest) Invoke(client *sdk.Client) (response *ListApDetailInfoResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ListApDetailInfoRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ListApDetailInfo", "", "")

	resp := struct {
		*responses.BaseResponse
		ListApDetailInfoResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ListApDetailInfoResponse

	err = client.DoAction(&req, &resp)
	return
}

type ListApDetailInfoResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
