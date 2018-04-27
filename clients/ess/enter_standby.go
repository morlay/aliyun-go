package ess

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type EnterStandbyRequest struct {
	requests.RpcRequest
	InstanceIds          *EnterStandbyInstanceIdList `position:"Query" type:"Repeated" name:"InstanceId"`
	ResourceOwnerAccount string                      `position:"Query" name:"ResourceOwnerAccount"`
	ScalingGroupId       string                      `position:"Query" name:"ScalingGroupId"`
	OwnerId              int64                       `position:"Query" name:"OwnerId"`
}

func (req *EnterStandbyRequest) Invoke(client *sdk.Client) (resp *EnterStandbyResponse, err error) {
	req.InitWithApiInfo("Ess", "2014-08-28", "EnterStandby", "ess", "")
	resp = &EnterStandbyResponse{}
	err = client.DoAction(req, resp)
	return
}

type EnterStandbyResponse struct {
	responses.BaseResponse
	RequestId string
}

type EnterStandbyInstanceIdList []string

func (list *EnterStandbyInstanceIdList) UnmarshalJSON(data []byte) error {
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
