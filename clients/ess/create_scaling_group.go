package ess

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateScalingGroupRequest struct {
	DBInstanceIds        string                            `position:"Query" name:"DBInstanceIds"`
	LoadBalancerIds      string                            `position:"Query" name:"LoadBalancerIds"`
	ResourceOwnerAccount string                            `position:"Query" name:"ResourceOwnerAccount"`
	ScalingGroupName     string                            `position:"Query" name:"ScalingGroupName"`
	VSwitchIdss          *CreateScalingGroupVSwitchIdsList `position:"Query" type:"Repeated" name:"VSwitchIds"`
	OwnerAccount         string                            `position:"Query" name:"OwnerAccount"`
	MinSize              int                               `position:"Query" name:"MinSize"`
	OwnerId              int64                             `position:"Query" name:"OwnerId"`
	VSwitchId            string                            `position:"Query" name:"VSwitchId"`
	MaxSize              int                               `position:"Query" name:"MaxSize"`
	DefaultCooldown      int                               `position:"Query" name:"DefaultCooldown"`
	RemovalPolicy1       string                            `position:"Query" name:"RemovalPolicy.1"`
	RemovalPolicy2       string                            `position:"Query" name:"RemovalPolicy.2"`
}

func (r CreateScalingGroupRequest) Invoke(client *sdk.Client) (response *CreateScalingGroupResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateScalingGroupRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ess", "2014-08-28", "CreateScalingGroup", "ess", "")

	resp := struct {
		*responses.BaseResponse
		CreateScalingGroupResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateScalingGroupResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateScalingGroupResponse struct {
	ScalingGroupId string
	RequestId      string
}

type CreateScalingGroupVSwitchIdsList []string

func (list *CreateScalingGroupVSwitchIdsList) UnmarshalJSON(data []byte) error {
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
