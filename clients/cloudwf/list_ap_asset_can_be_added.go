package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListApAssetCanBeAddedRequest struct {
	requests.RpcRequest
	SearchName  string `position:"Query" name:"SearchName"`
	ApgroupId   int64  `position:"Query" name:"ApgroupId"`
	Length      int    `position:"Query" name:"Length"`
	PageIndex   int    `position:"Query" name:"PageIndex"`
	SearchMac   string `position:"Query" name:"SearchMac"`
	SearchModel string `position:"Query" name:"SearchModel"`
}

func (req *ListApAssetCanBeAddedRequest) Invoke(client *sdk.Client) (resp *ListApAssetCanBeAddedResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ListApAssetCanBeAdded", "", "")
	resp = &ListApAssetCanBeAddedResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListApAssetCanBeAddedResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
