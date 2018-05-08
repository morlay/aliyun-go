package ess

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateNotificationConfigurationRequest struct {
	requests.RpcRequest
	ResourceOwnerAccount string                                               `position:"Query" name:"ResourceOwnerAccount"`
	ScalingGroupId       string                                               `position:"Query" name:"ScalingGroupId"`
	NotificationArn      string                                               `position:"Query" name:"NotificationArn"`
	NotificationTypes    *CreateNotificationConfigurationNotificationTypeList `position:"Query" type:"Repeated" name:"NotificationType"`
	OwnerId              int64                                                `position:"Query" name:"OwnerId"`
}

func (req *CreateNotificationConfigurationRequest) Invoke(client *sdk.Client) (resp *CreateNotificationConfigurationResponse, err error) {
	req.InitWithApiInfo("Ess", "2014-08-28", "CreateNotificationConfiguration", "ess", "")
	resp = &CreateNotificationConfigurationResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateNotificationConfigurationResponse struct {
	responses.BaseResponse
	RequestId common.String
}

type CreateNotificationConfigurationNotificationTypeList []string

func (list *CreateNotificationConfigurationNotificationTypeList) UnmarshalJSON(data []byte) error {
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
