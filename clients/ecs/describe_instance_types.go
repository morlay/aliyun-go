package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId     common.String
	InstanceTypes DescribeInstanceTypesInstanceTypeList
}

type DescribeInstanceTypesInstanceType struct {
	InstanceTypeId       common.String
	CpuCoreCount         common.Integer
	MemorySize           common.Float
	InstanceTypeFamily   common.String
	LocalStorageCapacity common.Long
	LocalStorageAmount   common.Integer
	LocalStorageCategory common.String
	GPUAmount            common.Integer
	GPUSpec              common.String
	InitialCredit        common.Integer
	BaselineCredit       common.Integer
	EniQuantity          common.Integer
	InstanceBandwidthRx  common.Integer
	InstanceBandwidthTx  common.Integer
	InstancePpsRx        common.Long
	InstancePpsTx        common.Long
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
