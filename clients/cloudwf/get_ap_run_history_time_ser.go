package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetApRunHistoryTimeSerRequest struct {
	Start int64 `position:"Query" name:"Start"`
	End   int64 `position:"Query" name:"End"`
	Id    int64 `position:"Query" name:"Id"`
}

func (r GetApRunHistoryTimeSerRequest) Invoke(client *sdk.Client) (response *GetApRunHistoryTimeSerResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetApRunHistoryTimeSerRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "GetApRunHistoryTimeSer", "", "")

	resp := struct {
		*responses.BaseResponse
		GetApRunHistoryTimeSerResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetApRunHistoryTimeSerResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetApRunHistoryTimeSerResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
