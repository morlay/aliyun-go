package ess

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ExitStandbyRequest struct {
	requests.RpcRequest
	InstanceIds          *ExitStandbyInstanceIdList `position:"Query" type:"Repeated" name:"InstanceId"`
	ResourceOwnerAccount string                     `position:"Query" name:"ResourceOwnerAccount"`
	ScalingGroupId       string                     `position:"Query" name:"ScalingGroupId"`
	OwnerId              int64                      `position:"Query" name:"OwnerId"`
}

func (req *ExitStandbyRequest) Invoke(client *sdk.Client) (resp *ExitStandbyResponse, err error) {
	req.InitWithApiInfo("Ess", "2014-08-28", "ExitStandby", "ess", "")
	resp = &ExitStandbyResponse{}
	err = client.DoAction(req, resp)
	return
}

type ExitStandbyResponse struct {
	responses.BaseResponse
	RequestId common.String
}

type ExitStandbyInstanceIdList []string

func (list *ExitStandbyInstanceIdList) UnmarshalJSON(data []byte) error {
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
