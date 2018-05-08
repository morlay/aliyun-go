package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetApDetailedStatusRequest struct {
	requests.RpcRequest
	Id int64 `position:"Query" name:"Id"`
}

func (req *GetApDetailedStatusRequest) Invoke(client *sdk.Client) (resp *GetApDetailedStatusResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "GetApDetailedStatus", "", "")
	resp = &GetApDetailedStatusResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetApDetailedStatusResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
	Message   common.String
	Data      common.String
	ErrorCode common.Integer
	ErrorMsg  common.String
}
