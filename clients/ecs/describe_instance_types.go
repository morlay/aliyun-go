package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeInstanceTypesRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	InstanceTypeFamily   string `position:"Query" name:"InstanceTypeFamily"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeInstanceTypesRequest) Invoke(client *sdk.Client) (resp *DescribeInstanceTypesResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeInstanceTypes", "ecs", "")
	resp = &DescribeInstanceTypesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeInstanceTypesResponse struct {
	responses.BaseResponse
	RequestId     string
	InstanceTypes DescribeInstanceTypesInstanceTypeList
}

type DescribeInstanceTypesInstanceType struct {
	InstanceTypeId       string
	CpuCoreCount         int
	MemorySize           float32
	InstanceTypeFamily   string
	LocalStorageCapacity int64
	LocalStorageAmount   int
	LocalStorageCategory string
	GPUAmount            int
	GPUSpec              string
	InitialCredit        int
	BaselineCredit       int
	EniQuantity          int
	InstanceBandwidthRx  int
	InstanceBandwidthTx  int
}

type DescribeInstanceTypesInstanceTypeList []DescribeInstanceTypesInstanceType

func (list *DescribeInstanceTypesInstanceTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInstanceTypesInstanceType)
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
