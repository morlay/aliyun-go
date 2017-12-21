package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ShopGroupUpdateRequest struct {
	requests.RpcRequest
	Gid         int64  `position:"Query" name:"Gid"`
	ShopIds     string `position:"Query" name:"ShopIds"`
	Name        string `position:"Query" name:"Name"`
	Description string `position:"Query" name:"Description"`
}

func (req *ShopGroupUpdateRequest) Invoke(client *sdk.Client) (resp *ShopGroupUpdateResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ShopGroupUpdate", "", "")
	resp = &ShopGroupUpdateResponse{}
	err = client.DoAction(req, resp)
	return
}

type ShopGroupUpdateResponse struct {
	responses.BaseResponse
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
