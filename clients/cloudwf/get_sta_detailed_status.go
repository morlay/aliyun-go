package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
