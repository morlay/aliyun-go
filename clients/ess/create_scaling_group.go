package ess

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateScalingGroupRequest struct {
	requests.RpcRequest
	MultiAZPolicy        string                           `position:"Query" name:"MultiAZPolicy"`
	DBInstanceIds        string                           `position:"Query" name:"DBInstanceIds"`
	LoadBalancerIds      string                           `position:"Query" name:"LoadBalancerIds"`
	ResourceOwnerAccount string                           `position:"Query" name:"ResourceOwnerAccount"`
	ScalingGroupName     string                           `position:"Query" name:"ScalingGroupName"`
	VSwitchIds           *CreateScalingGroupVSwitchIdList `position:"Query" type:"Repeated" name:"VSwitchId"`
	OwnerAccount         string                           `position:"Query" name:"OwnerAccount"`
	MinSize              int                              `position:"Query" name:"MinSize"`
	OwnerId              int64                            `position:"Query" name:"OwnerId"`
	VSwitchId            string                           `position:"Query" name:"VSwitchId"`
	MaxSize              int                              `position:"Query" name:"MaxSize"`
	DefaultCooldown      int                              `position:"Query" name:"DefaultCooldown"`
	RemovalPolicy1       string                           `position:"Query" name:"RemovalPolicy.1"`
	RemovalPolicy2       string                           `position:"Query" name:"RemovalPolicy.2"`
}

func (req *CreateScalingGroupRequest) Invoke(client *sdk.Client) (resp *CreateScalingGroupResponse, err error) {
	req.InitWithApiInfo("Ess", "2014-08-28", "CreateScalingGroup", "ess", "")
	resp = &CreateScalingGroupResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateScalingGroupResponse struct {
	responses.BaseResponse
	ScalingGroupId string
	RequestId      string
}

type CreateScalingGroupVSwitchIdList []string

func (list *CreateScalingGroupVSwitchIdList) UnmarshalJSON(data []byte) error {
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
