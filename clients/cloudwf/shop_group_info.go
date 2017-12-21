package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ShopGroupInfoRequest struct {
	requests.RpcRequest
	Gid int64 `position:"Query" name:"Gid"`
}

func (req *ShopGroupInfoRequest) Invoke(client *sdk.Client) (resp *ShopGroupInfoResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ShopGroupInfo", "", "")
	resp = &ShopGroupInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type ShopGroupInfoResponse struct {
	responses.BaseResponse
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
