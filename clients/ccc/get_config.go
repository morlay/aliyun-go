package ccc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetConfigRequest struct {
	InstanceId string `position:"Query" name:"InstanceId"`
	Name       string `position:"Query" name:"Name"`
	ObjectType string `position:"Query" name:"ObjectType"`
	ObjectId   string `position:"Query" name:"ObjectId"`
}

func (r GetConfigRequest) Invoke(client *sdk.Client) (response *GetConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CCC", "2017-07-05", "GetConfig", "ccc", "")

	resp := struct {
		*responses.BaseResponse
		GetConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetConfigResponse struct {
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
