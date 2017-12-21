package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ShopGroupShowListRequest struct {
	Page int   `position:"Query" name:"Page"`
	Bid  int64 `position:"Query" name:"Bid"`
	Per  int   `position:"Query" name:"Per"`
}

func (r ShopGroupShowListRequest) Invoke(client *sdk.Client) (response *ShopGroupShowListResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ShopGroupShowListRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ShopGroupShowList", "", "")

	resp := struct {
		*responses.BaseResponse
		ShopGroupShowListResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ShopGroupShowListResponse

	err = client.DoAction(&req, &resp)
	return
}

type ShopGroupShowListResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
