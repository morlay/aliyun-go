package ccc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetServiceExtensionsRequest struct {
	requests.RpcRequest
	ServiceType string `position:"Query" name:"ServiceType"`
	InstanceId  string `position:"Query" name:"InstanceId"`
}

func (req *GetServiceExtensionsRequest) Invoke(client *sdk.Client) (resp *GetServiceExtensionsResponse, err error) {
	req.InitWithApiInfo("CCC", "2017-07-05", "GetServiceExtensions", "CCC", "")
	resp = &GetServiceExtensionsResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetServiceExtensionsResponse struct {
	responses.BaseResponse
	RequestId         string
	Success           bool
	Code              string
	Message           string
	HttpStatusCode    int
	ServiceExtensions GetServiceExtensionsServiceExtensionList
}

type GetServiceExtensionsServiceExtension struct {
	Name   string
	Number string
}

type GetServiceExtensionsServiceExtensionList []GetServiceExtensionsServiceExtension

func (list *GetServiceExtensionsServiceExtensionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetServiceExtensionsServiceExtension)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
