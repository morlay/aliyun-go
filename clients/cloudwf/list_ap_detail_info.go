package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListApDetailInfoRequest struct {
	requests.RpcRequest
	ApAssetId int64 `position:"Query" name:"ApAssetId"`
}

func (req *ListApDetailInfoRequest) Invoke(client *sdk.Client) (resp *ListApDetailInfoResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ListApDetailInfo", "", "")
	resp = &ListApDetailInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListApDetailInfoResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
