package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type HeadquartersToolsCoincideRequest struct {
	requests.RpcRequest
	Bid int64 `position:"Query" name:"Bid"`
}

func (req *HeadquartersToolsCoincideRequest) Invoke(client *sdk.Client) (resp *HeadquartersToolsCoincideResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "HeadquartersToolsCoincide", "", "")
	resp = &HeadquartersToolsCoincideResponse{}
	err = client.DoAction(req, resp)
	return
}

type HeadquartersToolsCoincideResponse struct {
	responses.BaseResponse
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
