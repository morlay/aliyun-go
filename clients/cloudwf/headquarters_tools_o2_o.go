package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type HeadquartersToolsO2ORequest struct {
	requests.RpcRequest
	Bid int64 `position:"Query" name:"Bid"`
}

func (req *HeadquartersToolsO2ORequest) Invoke(client *sdk.Client) (resp *HeadquartersToolsO2OResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "HeadquartersToolsO2O", "", "")
	resp = &HeadquartersToolsO2OResponse{}
	err = client.DoAction(req, resp)
	return
}

type HeadquartersToolsO2OResponse struct {
	responses.BaseResponse
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
