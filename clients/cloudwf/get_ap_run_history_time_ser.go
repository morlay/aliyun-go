package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetApRunHistoryTimeSerRequest struct {
	requests.RpcRequest
	Start int64 `position:"Query" name:"Start"`
	End   int64 `position:"Query" name:"End"`
	Id    int64 `position:"Query" name:"Id"`
}

func (req *GetApRunHistoryTimeSerRequest) Invoke(client *sdk.Client) (resp *GetApRunHistoryTimeSerResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "GetApRunHistoryTimeSer", "", "")
	resp = &GetApRunHistoryTimeSerResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetApRunHistoryTimeSerResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
