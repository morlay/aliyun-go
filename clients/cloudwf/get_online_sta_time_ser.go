package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetOnlineStaTimeSerRequest struct {
	ZoomStart int64 `position:"Query" name:"ZoomStart"`
	CompanyId int64 `position:"Query" name:"CompanyId"`
	ApgroupId int64 `position:"Query" name:"ApgroupId"`
	Start     int64 `position:"Query" name:"Start"`
	ZoomEnd   int64 `position:"Query" name:"ZoomEnd"`
	End       int64 `position:"Query" name:"End"`
}

func (r GetOnlineStaTimeSerRequest) Invoke(client *sdk.Client) (response *GetOnlineStaTimeSerResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetOnlineStaTimeSerRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "GetOnlineStaTimeSer", "", "")

	resp := struct {
		*responses.BaseResponse
		GetOnlineStaTimeSerResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetOnlineStaTimeSerResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetOnlineStaTimeSerResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
