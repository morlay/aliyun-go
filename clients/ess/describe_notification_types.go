package ess

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeNotificationTypesRequest struct {
	requests.RpcRequest
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeNotificationTypesRequest) Invoke(client *sdk.Client) (resp *DescribeNotificationTypesResponse, err error) {
	req.InitWithApiInfo("Ess", "2014-08-28", "DescribeNotificationTypes", "ess", "")
	resp = &DescribeNotificationTypesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeNotificationTypesResponse struct {
	responses.BaseResponse
	RequestId         string
	NotificationTypes DescribeNotificationTypesNotificationTypeList
}

type DescribeNotificationTypesNotificationTypeList []string

func (list *DescribeNotificationTypesNotificationTypeList) UnmarshalJSON(data []byte) error {
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
