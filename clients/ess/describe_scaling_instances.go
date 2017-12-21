package ess

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeScalingInstancesRequest struct {
	requests.RpcRequest
	InstanceId10           string `position:"Query" name:"InstanceId.10"`
	ResourceOwnerId        int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId12           string `position:"Query" name:"InstanceId.12"`
	InstanceId11           string `position:"Query" name:"InstanceId.11"`
	ScalingGroupId         string `position:"Query" name:"ScalingGroupId"`
	LifecycleState         string `position:"Query" name:"LifecycleState"`
	CreationType           string `position:"Query" name:"CreationType"`
	PageNumber             int    `position:"Query" name:"PageNumber"`
	PageSize               int    `position:"Query" name:"PageSize"`
	InstanceId20           string `position:"Query" name:"InstanceId.20"`
	InstanceId1            string `position:"Query" name:"InstanceId.1"`
	InstanceId3            string `position:"Query" name:"InstanceId.3"`
	ResourceOwnerAccount   string `position:"Query" name:"ResourceOwnerAccount"`
	InstanceId2            string `position:"Query" name:"InstanceId.2"`
	InstanceId5            string `position:"Query" name:"InstanceId.5"`
	InstanceId4            string `position:"Query" name:"InstanceId.4"`
	OwnerAccount           string `position:"Query" name:"OwnerAccount"`
	InstanceId7            string `position:"Query" name:"InstanceId.7"`
	InstanceId6            string `position:"Query" name:"InstanceId.6"`
	InstanceId9            string `position:"Query" name:"InstanceId.9"`
	InstanceId8            string `position:"Query" name:"InstanceId.8"`
	OwnerId                int64  `position:"Query" name:"OwnerId"`
	ScalingConfigurationId string `position:"Query" name:"ScalingConfigurationId"`
	HealthStatus           string `position:"Query" name:"HealthStatus"`
	InstanceId18           string `position:"Query" name:"InstanceId.18"`
	InstanceId17           string `position:"Query" name:"InstanceId.17"`
	InstanceId19           string `position:"Query" name:"InstanceId.19"`
	InstanceId14           string `position:"Query" name:"InstanceId.14"`
	InstanceId13           string `position:"Query" name:"InstanceId.13"`
	InstanceId16           string `position:"Query" name:"InstanceId.16"`
	InstanceId15           string `position:"Query" name:"InstanceId.15"`
}

func (req *DescribeScalingInstancesRequest) Invoke(client *sdk.Client) (resp *DescribeScalingInstancesResponse, err error) {
	req.InitWithApiInfo("Ess", "2014-08-28", "DescribeScalingInstances", "ess", "")
	resp = &DescribeScalingInstancesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeScalingInstancesResponse struct {
	responses.BaseResponse
	TotalCount       int
	PageNumber       int
	PageSize         int
	RequestId        string
	ScalingInstances DescribeScalingInstancesScalingInstanceList
}

type DescribeScalingInstancesScalingInstance struct {
	InstanceId             string
	ScalingConfigurationId string
	ScalingGroupId         string
	HealthStatus           string
	LoadBalancerWeight     int
	LifecycleState         string
	CreationTime           string
	CreationType           string
}

type DescribeScalingInstancesScalingInstanceList []DescribeScalingInstancesScalingInstance

func (list *DescribeScalingInstancesScalingInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeScalingInstancesScalingInstance)
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
