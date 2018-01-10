package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SaveGroupApRadioConfigRequest struct {
	requests.RpcRequest
	JsonData string `position:"Query" name:"JsonData"`
}

func (req *SaveGroupApRadioConfigRequest) Invoke(client *sdk.Client) (resp *SaveGroupApRadioConfigResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "SaveGroupApRadioConfig", "", "")
	resp = &SaveGroupApRadioConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type SaveGroupApRadioConfigResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
