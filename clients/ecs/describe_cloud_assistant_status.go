package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeCloudAssistantStatusRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64                                       `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string                                      `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string                                      `position:"Query" name:"OwnerAccount"`
	OwnerId              int64                                       `position:"Query" name:"OwnerId"`
	InstanceIds          *DescribeCloudAssistantStatusInstanceIdList `position:"Query" type:"Repeated" name:"InstanceId"`
}

func (req *DescribeCloudAssistantStatusRequest) Invoke(client *sdk.Client) (resp *DescribeCloudAssistantStatusResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeCloudAssistantStatus", "ecs", "")
	resp = &DescribeCloudAssistantStatusResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeCloudAssistantStatusResponse struct {
	responses.BaseResponse
	RequestId                       string
	InstanceCloudAssistantStatusSet DescribeCloudAssistantStatusInstanceCloudAssistantStatusList
}

type DescribeCloudAssistantStatusInstanceCloudAssistantStatus struct {
	InstanceId           string
	CloudAssistantStatus string
}

type DescribeCloudAssistantStatusInstanceIdList []string

func (list *DescribeCloudAssistantStatusInstanceIdList) UnmarshalJSON(data []byte) error {
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

type DescribeCloudAssistantStatusInstanceCloudAssistantStatusList []DescribeCloudAssistantStatusInstanceCloudAssistantStatus

func (list *DescribeCloudAssistantStatusInstanceCloudAssistantStatusList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCloudAssistantStatusInstanceCloudAssistantStatus)
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
