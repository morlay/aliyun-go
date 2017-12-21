package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type HeadquartersToolsContrastRequest struct {
	requests.RpcRequest
	Bid int64 `position:"Query" name:"Bid"`
}

func (req *HeadquartersToolsContrastRequest) Invoke(client *sdk.Client) (resp *HeadquartersToolsContrastResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "HeadquartersToolsContrast", "", "")
	resp = &HeadquartersToolsContrastResponse{}
	err = client.DoAction(req, resp)
	return
}

type HeadquartersToolsContrastResponse struct {
	responses.BaseResponse
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
