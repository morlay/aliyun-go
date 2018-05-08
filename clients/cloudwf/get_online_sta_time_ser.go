package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetOnlineStaTimeSerRequest struct {
	requests.RpcRequest
	ZoomStart int64 `position:"Query" name:"ZoomStart"`
	CompanyId int64 `position:"Query" name:"CompanyId"`
	ApgroupId int64 `position:"Query" name:"ApgroupId"`
	Start     int64 `position:"Query" name:"Start"`
	ZoomEnd   int64 `position:"Query" name:"ZoomEnd"`
	End       int64 `position:"Query" name:"End"`
}

func (req *GetOnlineStaTimeSerRequest) Invoke(client *sdk.Client) (resp *GetOnlineStaTimeSerResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "GetOnlineStaTimeSer", "", "")
	resp = &GetOnlineStaTimeSerResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetOnlineStaTimeSerResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
	Message   common.String
	Data      common.String
	ErrorCode common.Integer
	ErrorMsg  common.String
}
