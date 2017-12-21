package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteApRadioSsidConfigRequest struct {
	requests.RpcRequest
	InstantlyEffective int   `position:"Query" name:"InstantlyEffective"`
	Id                 int64 `position:"Query" name:"Id"`
}

func (req *DeleteApRadioSsidConfigRequest) Invoke(client *sdk.Client) (resp *DeleteApRadioSsidConfigResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "DeleteApRadioSsidConfig", "", "")
	resp = &DeleteApRadioSsidConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteApRadioSsidConfigResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
