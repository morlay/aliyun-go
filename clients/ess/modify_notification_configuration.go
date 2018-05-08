package ess

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyNotificationConfigurationRequest struct {
	requests.RpcRequest
	ResourceOwnerAccount string                                               `position:"Query" name:"ResourceOwnerAccount"`
	ScalingGroupId       string                                               `position:"Query" name:"ScalingGroupId"`
	NotificationArn      string                                               `position:"Query" name:"NotificationArn"`
	NotificationTypes    *ModifyNotificationConfigurationNotificationTypeList `position:"Query" type:"Repeated" name:"NotificationType"`
	OwnerId              int64                                                `position:"Query" name:"OwnerId"`
}

func (req *ModifyNotificationConfigurationRequest) Invoke(client *sdk.Client) (resp *ModifyNotificationConfigurationResponse, err error) {
	req.InitWithApiInfo("Ess", "2014-08-28", "ModifyNotificationConfiguration", "ess", "")
	resp = &ModifyNotificationConfigurationResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyNotificationConfigurationResponse struct {
	responses.BaseResponse
	RequestId common.String
}

type ModifyNotificationConfigurationNotificationTypeList []string

func (list *ModifyNotificationConfigurationNotificationTypeList) UnmarshalJSON(data []byte) error {
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
