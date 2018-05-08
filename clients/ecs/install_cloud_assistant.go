package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type InstallCloudAssistantRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64                                `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string                               `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string                               `position:"Query" name:"OwnerAccount"`
	OwnerId              int64                                `position:"Query" name:"OwnerId"`
	InstanceIds          *InstallCloudAssistantInstanceIdList `position:"Query" type:"Repeated" name:"InstanceId"`
}

func (req *InstallCloudAssistantRequest) Invoke(client *sdk.Client) (resp *InstallCloudAssistantResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "InstallCloudAssistant", "ecs", "")
	resp = &InstallCloudAssistantResponse{}
	err = client.DoAction(req, resp)
	return
}

type InstallCloudAssistantResponse struct {
	responses.BaseResponse
	RequestId common.String
}

type InstallCloudAssistantInstanceIdList []string

func (list *InstallCloudAssistantInstanceIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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
