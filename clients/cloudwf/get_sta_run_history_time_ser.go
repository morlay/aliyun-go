package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetStaRunHistoryTimeSerRequest struct {
	requests.RpcRequest
	Id int64 `position:"Query" name:"Id"`
}

func (req *GetStaRunHistoryTimeSerRequest) Invoke(client *sdk.Client) (resp *GetStaRunHistoryTimeSerResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "GetStaRunHistoryTimeSer", "", "")
	resp = &GetStaRunHistoryTimeSerResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetStaRunHistoryTimeSerResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
