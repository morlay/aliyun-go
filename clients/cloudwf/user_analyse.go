package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UserAnalyseRequest struct {
	requests.RpcRequest
	Gsid int64 `position:"Query" name:"Gsid"`
}

func (req *UserAnalyseRequest) Invoke(client *sdk.Client) (resp *UserAnalyseResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "UserAnalyse", "", "")
	resp = &UserAnalyseResponse{}
	err = client.DoAction(req, resp)
	return
}

type UserAnalyseResponse struct {
	responses.BaseResponse
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
