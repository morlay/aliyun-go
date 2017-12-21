package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ShopGroupInfoRequest struct {
	Gid int64 `position:"Query" name:"Gid"`
}

func (r ShopGroupInfoRequest) Invoke(client *sdk.Client) (response *ShopGroupInfoResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ShopGroupInfoRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ShopGroupInfo", "", "")

	resp := struct {
		*responses.BaseResponse
		ShopGroupInfoResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ShopGroupInfoResponse

	err = client.DoAction(&req, &resp)
	return
}

type ShopGroupInfoResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
