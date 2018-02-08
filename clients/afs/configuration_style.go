package afs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ConfigurationStyleRequest struct {
	requests.RpcRequest
	ResourceOwnerId     int64  `position:"Query" name:"ResourceOwnerId"`
	SourceIp            string `position:"Query" name:"SourceIp"`
	ConfigurationMethod string `position:"Query" name:"ConfigurationMethod"`
	ApplyType           string `position:"Query" name:"ApplyType"`
	Scene               string `position:"Query" name:"Scene"`
}

func (req *ConfigurationStyleRequest) Invoke(client *sdk.Client) (resp *ConfigurationStyleResponse, err error) {
	req.InitWithApiInfo("afs", "2018-01-12", "ConfigurationStyle", "", "")
	resp = &ConfigurationStyleResponse{}
	err = client.DoAction(req, resp)
	return
}

type ConfigurationStyleResponse struct {
	responses.BaseResponse
	RequestId string
	BizCode   string
	CodeData  ConfigurationStyleCodeData
}

type ConfigurationStyleCodeData struct {
	Html   string
	Net    string
	Php    string
	Python string
	Java   string
	NodeJs string
}
