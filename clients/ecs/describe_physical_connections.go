package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribePhysicalConnectionsRequest struct {
	Filters              *DescribePhysicalConnectionsFilterList `position:"Query" type:"Repeated" name:"Filter"`
	ResourceOwnerId      int64                                  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string                                 `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string                                 `position:"Query" name:"ClientToken"`
	OwnerAccount         string                                 `position:"Query" name:"OwnerAccount"`
	PageSize             int                                    `position:"Query" name:"PageSize"`
	UserCidr             string                                 `position:"Query" name:"UserCidr"`
	OwnerId              int64                                  `position:"Query" name:"OwnerId"`
	PageNumber           int                                    `position:"Query" name:"PageNumber"`
}

func (r DescribePhysicalConnectionsRequest) Invoke(client *sdk.Client) (response *DescribePhysicalConnectionsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribePhysicalConnectionsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribePhysicalConnections", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		DescribePhysicalConnectionsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribePhysicalConnectionsResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribePhysicalConnectionsFilter struct {
	Key    string                                `name:"Key"`
	Values *DescribePhysicalConnectionsValueList `type:"Repeated" name:"Value"`
}

type DescribePhysicalConnectionsResponse struct {
	RequestId             string
	PageNumber            int
	PageSize              int
	TotalCount            int
	PhysicalConnectionSet DescribePhysicalConnectionsPhysicalConnectionTypeList
}

type DescribePhysicalConnectionsPhysicalConnectionType struct {
	PhysicalConnectionId          string
	AccessPointId                 string
	Type                          string
	Status                        string
	BusinessStatus                string
	CreationTime                  string
	EnabledTime                   string
	LineOperator                  string
	Spec                          string
	PeerLocation                  string
	PortType                      string
	RedundantPhysicalConnectionId string
	Name                          string
	Description                   string
	AdLocation                    string
	PortNumber                    string
	CircuitCode                   string
	Bandwidth                     int64
}

type DescribePhysicalConnectionsFilterList []DescribePhysicalConnectionsFilter

func (list *DescribePhysicalConnectionsFilterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribePhysicalConnectionsFilter)
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

type DescribePhysicalConnectionsValueList []string

func (list *DescribePhysicalConnectionsValueList) UnmarshalJSON(data []byte) error {
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

type DescribePhysicalConnectionsPhysicalConnectionTypeList []DescribePhysicalConnectionsPhysicalConnectionType

func (list *DescribePhysicalConnectionsPhysicalConnectionTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribePhysicalConnectionsPhysicalConnectionType)
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
