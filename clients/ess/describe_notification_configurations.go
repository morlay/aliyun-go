package ess

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeNotificationConfigurationsRequest struct {
	requests.RpcRequest
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ScalingGroupId       string `position:"Query" name:"ScalingGroupId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeNotificationConfigurationsRequest) Invoke(client *sdk.Client) (resp *DescribeNotificationConfigurationsResponse, err error) {
	req.InitWithApiInfo("Ess", "2014-08-28", "DescribeNotificationConfigurations", "ess", "")
	resp = &DescribeNotificationConfigurationsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeNotificationConfigurationsResponse struct {
	responses.BaseResponse
	RequestId                       string
	NotificationConfigurationModels DescribeNotificationConfigurationsNotificationConfigurationModelList
}

type DescribeNotificationConfigurationsNotificationConfigurationModel struct {
	ScalingGroupId    string
	NotificationArn   string
	NotificationTypes DescribeNotificationConfigurationsNotificationTypeList
}

type DescribeNotificationConfigurationsNotificationConfigurationModelList []DescribeNotificationConfigurationsNotificationConfigurationModel

func (list *DescribeNotificationConfigurationsNotificationConfigurationModelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeNotificationConfigurationsNotificationConfigurationModel)
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

type DescribeNotificationConfigurationsNotificationTypeList []string

func (list *DescribeNotificationConfigurationsNotificationTypeList) UnmarshalJSON(data []byte) error {
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
