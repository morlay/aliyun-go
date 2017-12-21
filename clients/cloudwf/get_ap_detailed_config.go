package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetApDetailedConfigRequest struct {
	requests.RpcRequest
	Id int64 `position:"Query" name:"Id"`
}

func (req *GetApDetailedConfigRequest) Invoke(client *sdk.Client) (resp *GetApDetailedConfigResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "GetApDetailedConfig", "", "")
	resp = &GetApDetailedConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetApDetailedConfigResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
