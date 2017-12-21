package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type PeripheryAnalyseRequest struct {
	requests.RpcRequest
	Gsid int64 `position:"Query" name:"Gsid"`
}

func (req *PeripheryAnalyseRequest) Invoke(client *sdk.Client) (resp *PeripheryAnalyseResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "PeripheryAnalyse", "", "")
	resp = &PeripheryAnalyseResponse{}
	err = client.DoAction(req, resp)
	return
}

type PeripheryAnalyseResponse struct {
	responses.BaseResponse
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
