package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetOnlineApTimeSerRequest struct {
	ZoomStart int64 `position:"Query" name:"ZoomStart"`
	CompanyId int64 `position:"Query" name:"CompanyId"`
	ApgroupId int64 `position:"Query" name:"ApgroupId"`
	Start     int64 `position:"Query" name:"Start"`
	ZoomEnd   int64 `position:"Query" name:"ZoomEnd"`
	End       int64 `position:"Query" name:"End"`
}

func (r GetOnlineApTimeSerRequest) Invoke(client *sdk.Client) (response *GetOnlineApTimeSerResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetOnlineApTimeSerRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "GetOnlineApTimeSer", "", "")

	resp := struct {
		*responses.BaseResponse
		GetOnlineApTimeSerResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetOnlineApTimeSerResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetOnlineApTimeSerResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
