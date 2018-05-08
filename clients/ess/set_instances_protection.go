package ess

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetInstancesProtectionRequest struct {
	requests.RpcRequest
	InstanceIds          *SetInstancesProtectionInstanceIdList `position:"Query" type:"Repeated" name:"InstanceId"`
	ResourceOwnerAccount string                                `position:"Query" name:"ResourceOwnerAccount"`
	ScalingGroupId       string                                `position:"Query" name:"ScalingGroupId"`
	OwnerId              int64                                 `position:"Query" name:"OwnerId"`
	ProtectedFromScaleIn string                                `position:"Query" name:"ProtectedFromScaleIn"`
}

func (req *SetInstancesProtectionRequest) Invoke(client *sdk.Client) (resp *SetInstancesProtectionResponse, err error) {
	req.InitWithApiInfo("Ess", "2014-08-28", "SetInstancesProtection", "ess", "")
	resp = &SetInstancesProtectionResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetInstancesProtectionResponse struct {
	responses.BaseResponse
	RequestId string
}

type SetInstancesProtectionInstanceIdList []string

func (list *SetInstancesProtectionInstanceIdList) UnmarshalJSON(data []byte) error {
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
