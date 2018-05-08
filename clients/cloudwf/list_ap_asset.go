package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListApAssetRequest struct {
	requests.RpcRequest
	OrderCol       string `position:"Query" name:"OrderCol"`
	SearchName     string `position:"Query" name:"SearchName"`
	SearchSerialNo string `position:"Query" name:"SearchSerialNo"`
	Length         int    `position:"Query" name:"Length"`
	PageIndex      int    `position:"Query" name:"PageIndex"`
	SearchMac      string `position:"Query" name:"SearchMac"`
	OrderDir       string `position:"Query" name:"OrderDir"`
	SearchModel    string `position:"Query" name:"SearchModel"`
}

func (req *ListApAssetRequest) Invoke(client *sdk.Client) (resp *ListApAssetResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ListApAsset", "", "")
	resp = &ListApAssetResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListApAssetResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
	Message   common.String
	Data      common.String
	ErrorCode common.Integer
	ErrorMsg  common.String
}
