package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ShopShowListRequest struct {
	Gid        int64  `position:"Query" name:"Gid"`
	Address    string `position:"Query" name:"Address"`
	Name       string `position:"Query" name:"Name"`
	Dirc       string `position:"Query" name:"Dirc"`
	Page       int    `position:"Query" name:"Page"`
	Bid        int64  `position:"Query" name:"Bid"`
	Per        int    `position:"Query" name:"Per"`
	ShopStatus int    `position:"Query" name:"ShopStatus"`
}

func (r ShopShowListRequest) Invoke(client *sdk.Client) (response *ShopShowListResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ShopShowListRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ShopShowList", "", "")

	resp := struct {
		*responses.BaseResponse
		ShopShowListResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ShopShowListResponse

	err = client.DoAction(&req, &resp)
	return
}

type ShopShowListResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
