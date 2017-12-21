package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetRadioRunHistoryTimeSerRequest struct {
	Id int64 `position:"Query" name:"Id"`
}

func (r GetRadioRunHistoryTimeSerRequest) Invoke(client *sdk.Client) (response *GetRadioRunHistoryTimeSerResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetRadioRunHistoryTimeSerRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "GetRadioRunHistoryTimeSer", "", "")

	resp := struct {
		*responses.BaseResponse
		GetRadioRunHistoryTimeSerResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetRadioRunHistoryTimeSerResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetRadioRunHistoryTimeSerResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
