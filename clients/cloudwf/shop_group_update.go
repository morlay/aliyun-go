package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ShopGroupUpdateRequest struct {
	Gid         int64  `position:"Query" name:"Gid"`
	ShopIds     string `position:"Query" name:"ShopIds"`
	Name        string `position:"Query" name:"Name"`
	Description string `position:"Query" name:"Description"`
}

func (r ShopGroupUpdateRequest) Invoke(client *sdk.Client) (response *ShopGroupUpdateResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ShopGroupUpdateRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ShopGroupUpdate", "", "")

	resp := struct {
		*responses.BaseResponse
		ShopGroupUpdateResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ShopGroupUpdateResponse

	err = client.DoAction(&req, &resp)
	return
}

type ShopGroupUpdateResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
