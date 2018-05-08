package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetOnlineApTimeSerRequest struct {
	requests.RpcRequest
	ZoomStart int64 `position:"Query" name:"ZoomStart"`
	CompanyId int64 `position:"Query" name:"CompanyId"`
	ApgroupId int64 `position:"Query" name:"ApgroupId"`
	Start     int64 `position:"Query" name:"Start"`
	ZoomEnd   int64 `position:"Query" name:"ZoomEnd"`
	End       int64 `position:"Query" name:"End"`
}

func (req *GetOnlineApTimeSerRequest) Invoke(client *sdk.Client) (resp *GetOnlineApTimeSerResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "GetOnlineApTimeSer", "", "")
	resp = &GetOnlineApTimeSerResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetOnlineApTimeSerResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
	Message   common.String
	Data      common.String
	ErrorCode common.Integer
	ErrorMsg  common.String
}
