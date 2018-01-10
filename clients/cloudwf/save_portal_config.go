package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SavePortalConfigRequest struct {
	requests.RpcRequest
	JsonData string `position:"Query" name:"JsonData"`
}

func (req *SavePortalConfigRequest) Invoke(client *sdk.Client) (resp *SavePortalConfigResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "SavePortalConfig", "", "")
	resp = &SavePortalConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type SavePortalConfigResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
