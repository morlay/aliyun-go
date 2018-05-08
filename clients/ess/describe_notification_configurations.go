package ess

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId                       common.String
	NotificationConfigurationModels DescribeNotificationConfigurationsNotificationConfigurationModelList
}

type DescribeNotificationConfigurationsNotificationConfigurationModel struct {
	ScalingGroupId    common.String
	NotificationArn   common.String
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

type DescribeNotificationConfigurationsNotificationTypeList []common.String

func (list *DescribeNotificationConfigurationsNotificationTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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
