package ccc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetConfigRequest struct {
	requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	Name       string `position:"Query" name:"Name"`
	ObjectType string `position:"Query" name:"ObjectType"`
	ObjectId   string `position:"Query" name:"ObjectId"`
}

func (req *GetConfigRequest) Invoke(client *sdk.Client) (resp *GetConfigResponse, err error) {
	req.InitWithApiInfo("CCC", "2017-07-05", "GetConfig", "ccc", "")
	resp = &GetConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetConfigResponse struct {
	responses.BaseResponse
	RequestId      common.String
	Success        bool
	Code           common.String
	Message        common.String
	HttpStatusCode common.Integer
	ConfigItem     GetConfigConfigItem
}

type GetConfigConfigItem struct {
	Name  common.String
	Value common.String
}
