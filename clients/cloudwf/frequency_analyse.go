package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type FrequencyAnalyseRequest struct {
	requests.RpcRequest
	Gsid int64 `position:"Query" name:"Gsid"`
}

func (req *FrequencyAnalyseRequest) Invoke(client *sdk.Client) (resp *FrequencyAnalyseResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "FrequencyAnalyse", "", "")
	resp = &FrequencyAnalyseResponse{}
	err = client.DoAction(req, resp)
	return
}

type FrequencyAnalyseResponse struct {
	responses.BaseResponse
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
