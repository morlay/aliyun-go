package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListApAssetCanBeAddedRequest struct {
	SearchName  string `position:"Query" name:"SearchName"`
	ApgroupId   int64  `position:"Query" name:"ApgroupId"`
	Length      int    `position:"Query" name:"Length"`
	PageIndex   int    `position:"Query" name:"PageIndex"`
	SearchMac   string `position:"Query" name:"SearchMac"`
	SearchModel string `position:"Query" name:"SearchModel"`
}

func (r ListApAssetCanBeAddedRequest) Invoke(client *sdk.Client) (response *ListApAssetCanBeAddedResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ListApAssetCanBeAddedRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ListApAssetCanBeAdded", "", "")

	resp := struct {
		*responses.BaseResponse
		ListApAssetCanBeAddedResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ListApAssetCanBeAddedResponse

	err = client.DoAction(&req, &resp)
	return
}

type ListApAssetCanBeAddedResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
