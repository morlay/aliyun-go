package ccc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetConfigRequest struct {
	requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	Name       string `position:"Query" name:"Name"`
	ObjectType string `position:"Query" name:"ObjectType"`
	ObjectId   string `position:"Query" name:"ObjectId"`
}

func (req *GetConfigRequest) Invoke(client *sdk.Client) (resp *GetConfigResponse, err error) {
	req.InitWithApiInfo("CCC", "2017-07-05", "GetConfig", "CCC", "")
	resp = &GetConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetConfigResponse struct {
	responses.BaseResponse
	RequestId      string
	Success        bool
	Code           string
	Message        string
	HttpStatusCode int
	ConfigItem     GetConfigConfigItem
}

type GetConfigConfigItem struct {
	Name  string
	Value string
}
