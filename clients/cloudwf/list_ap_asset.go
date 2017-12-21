package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListApAssetRequest struct {
	OrderCol       string `position:"Query" name:"OrderCol"`
	SearchName     string `position:"Query" name:"SearchName"`
	SearchSerialNo string `position:"Query" name:"SearchSerialNo"`
	Length         int    `position:"Query" name:"Length"`
	PageIndex      int    `position:"Query" name:"PageIndex"`
	SearchMac      string `position:"Query" name:"SearchMac"`
	OrderDir       string `position:"Query" name:"OrderDir"`
	SearchModel    string `position:"Query" name:"SearchModel"`
}

func (r ListApAssetRequest) Invoke(client *sdk.Client) (response *ListApAssetResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ListApAssetRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ListApAsset", "", "")

	resp := struct {
		*responses.BaseResponse
		ListApAssetResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ListApAssetResponse

	err = client.DoAction(&req, &resp)
	return
}

type ListApAssetResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
