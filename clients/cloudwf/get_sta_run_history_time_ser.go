package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetStaRunHistoryTimeSerRequest struct {
	Id int64 `position:"Query" name:"Id"`
}

func (r GetStaRunHistoryTimeSerRequest) Invoke(client *sdk.Client) (response *GetStaRunHistoryTimeSerResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetStaRunHistoryTimeSerRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "GetStaRunHistoryTimeSer", "", "")

	resp := struct {
		*responses.BaseResponse
		GetStaRunHistoryTimeSerResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetStaRunHistoryTimeSerResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetStaRunHistoryTimeSerResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
