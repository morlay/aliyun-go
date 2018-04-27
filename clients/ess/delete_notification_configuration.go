package ess

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteNotificationConfigurationRequest struct {
	requests.RpcRequest
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ScalingGroupId       string `position:"Query" name:"ScalingGroupId"`
	NotificationArn      string `position:"Query" name:"NotificationArn"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DeleteNotificationConfigurationRequest) Invoke(client *sdk.Client) (resp *DeleteNotificationConfigurationResponse, err error) {
	req.InitWithApiInfo("Ess", "2014-08-28", "DeleteNotificationConfiguration", "ess", "")
	resp = &DeleteNotificationConfigurationResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteNotificationConfigurationResponse struct {
	responses.BaseResponse
	RequestId string
}
