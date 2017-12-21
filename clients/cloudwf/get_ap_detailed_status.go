package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
