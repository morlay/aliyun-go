package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetStaDetailedStatusRequest struct {
	requests.RpcRequest
	Id int64 `position:"Query" name:"Id"`
}

func (req *GetStaDetailedStatusRequest) Invoke(client *sdk.Client) (resp *GetStaDetailedStatusResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "GetStaDetailedStatus", "", "")
	resp = &GetStaDetailedStatusResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetStaDetailedStatusResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
	Message   common.String
	Data      common.String
	ErrorCode common.Integer
	ErrorMsg  common.String
}
