package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeInstanceTypeFamiliesRequest struct {
	Generation           string `position:"Query" name:"Generation"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeInstanceTypeFamiliesRequest) Invoke(client *sdk.Client) (response *DescribeInstanceTypeFamiliesResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeInstanceTypeFamiliesRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeInstanceTypeFamilies", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		DescribeInstanceTypeFamiliesResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeInstanceTypeFamiliesResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeInstanceTypeFamiliesResponse struct {
	RequestId            string
	InstanceTypeFamilies DescribeInstanceTypeFamiliesInstanceTypeFamilyList
}

type DescribeInstanceTypeFamiliesInstanceTypeFamily struct {
	InstanceTypeFamilyId string
	Generation           string
}

type DescribeInstanceTypeFamiliesInstanceTypeFamilyList []DescribeInstanceTypeFamiliesInstanceTypeFamily

func (list *DescribeInstanceTypeFamiliesInstanceTypeFamilyList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInstanceTypeFamiliesInstanceTypeFamily)
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
