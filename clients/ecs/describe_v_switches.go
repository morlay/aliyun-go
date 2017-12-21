package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeVSwitchesRequest struct {
	VSwitchId            string `position:"Query" name:"VSwitchId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	VpcId                string `position:"Query" name:"VpcId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int    `position:"Query" name:"PageSize"`
	ZoneId               string `position:"Query" name:"ZoneId"`
	IsDefault            string `position:"Query" name:"IsDefault"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
}

func (r DescribeVSwitchesRequest) Invoke(client *sdk.Client) (response *DescribeVSwitchesResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeVSwitchesRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeVSwitches", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		DescribeVSwitchesResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeVSwitchesResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeVSwitchesResponse struct {
	RequestId  string
	TotalCount int
	PageNumber int
	PageSize   int
	VSwitches  DescribeVSwitchesVSwitchList
}

type DescribeVSwitchesVSwitch struct {
	VSwitchId               string
	VpcId                   string
	Status                  string
	CidrBlock               string
	ZoneId                  string
	AvailableIpAddressCount int64
	Description             string
	VSwitchName             string
	CreationTime            string
	IsDefault               bool
}

type DescribeVSwitchesVSwitchList []DescribeVSwitchesVSwitch

func (list *DescribeVSwitchesVSwitchList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeVSwitchesVSwitch)
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
