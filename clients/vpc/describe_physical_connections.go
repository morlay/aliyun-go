package vpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribePhysicalConnectionsRequest struct {
	requests.RpcRequest
	Filters              *DescribePhysicalConnectionsFilterList `position:"Query" type:"Repeated" name:"Filter"`
	ResourceOwnerId      int64                                  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string                                 `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string                                 `position:"Query" name:"ClientToken"`
	OwnerAccount         string                                 `position:"Query" name:"OwnerAccount"`
	PageSize             int                                    `position:"Query" name:"PageSize"`
	OwnerId              int64                                  `position:"Query" name:"OwnerId"`
	PageNumber           int                                    `position:"Query" name:"PageNumber"`
}

func (req *DescribePhysicalConnectionsRequest) Invoke(client *sdk.Client) (resp *DescribePhysicalConnectionsResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "DescribePhysicalConnections", "vpc", "")
	resp = &DescribePhysicalConnectionsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribePhysicalConnectionsFilter struct {
	Key    string                                `name:"Key"`
	Values *DescribePhysicalConnectionsValueList `type:"Repeated" name:"Value"`
}

type DescribePhysicalConnectionsResponse struct {
	responses.BaseResponse
	RequestId             common.String
	PageNumber            common.Integer
	PageSize              common.Integer
	TotalCount            common.Integer
	PhysicalConnectionSet DescribePhysicalConnectionsPhysicalConnectionTypeList
}

type DescribePhysicalConnectionsPhysicalConnectionType struct {
	PhysicalConnectionId          common.String
	AccessPointId                 common.String
	Type                          common.String
	Status                        common.String
	BusinessStatus                common.String
	CreationTime                  common.String
	EnabledTime                   common.String
	LineOperator                  common.String
	Spec                          common.String
	PeerLocation                  common.String
	PortType                      common.String
	RedundantPhysicalConnectionId common.String
	Name                          common.String
	Description                   common.String
	AdLocation                    common.String
	PortNumber                    common.String
	CircuitCode                   common.String
	Bandwidth                     common.Long
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
