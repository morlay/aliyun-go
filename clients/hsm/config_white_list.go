package hsm

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ConfigWhiteListRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId      string `position:"Query" name:"InstanceId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	WhiteList       string `position:"Query" name:"WhiteList"`
}

func (req *ConfigWhiteListRequest) Invoke(client *sdk.Client) (resp *ConfigWhiteListResponse, err error) {
	req.InitWithApiInfo("hsm", "2018-01-11", "ConfigWhiteList", "hsm", "")
	resp = &ConfigWhiteListResponse{}
	err = client.DoAction(req, resp)
	return
}

type ConfigWhiteListResponse struct {
	responses.BaseResponse
	RequestId string
}
