package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetRadioRunHistoryTimeSerRequest struct {
	requests.RpcRequest
	Id int64 `position:"Query" name:"Id"`
}

func (req *GetRadioRunHistoryTimeSerRequest) Invoke(client *sdk.Client) (resp *GetRadioRunHistoryTimeSerResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "GetRadioRunHistoryTimeSer", "", "")
	resp = &GetRadioRunHistoryTimeSerResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetRadioRunHistoryTimeSerResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
