package vpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeVirtualBorderRoutersForPhysicalConnectionRequest struct {
	requests.RpcRequest
	Filters              *DescribeVirtualBorderRoutersForPhysicalConnectionFilterList `position:"Query" type:"Repeated" name:"Filter"`
	ResourceOwnerId      int64                                                        `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string                                                       `position:"Query" name:"ResourceOwnerAccount"`
	PhysicalConnectionId string                                                       `position:"Query" name:"PhysicalConnectionId"`
	PageSize             int                                                          `position:"Query" name:"PageSize"`
	OwnerId              int64                                                        `position:"Query" name:"OwnerId"`
	PageNumber           int                                                          `position:"Query" name:"PageNumber"`
}

func (req *DescribeVirtualBorderRoutersForPhysicalConnectionRequest) Invoke(client *sdk.Client) (resp *DescribeVirtualBorderRoutersForPhysicalConnectionResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "DescribeVirtualBorderRoutersForPhysicalConnection", "vpc", "")
	resp = &DescribeVirtualBorderRoutersForPhysicalConnectionResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeVirtualBorderRoutersForPhysicalConnectionFilter struct {
	Key    string                                                      `name:"Key"`
	Values *DescribeVirtualBorderRoutersForPhysicalConnectionValueList `type:"Repeated" name:"Value"`
}

type DescribeVirtualBorderRoutersForPhysicalConnectionResponse struct {
	responses.BaseResponse
	RequestId                                   common.String
	PageNumber                                  common.Integer
	PageSize                                    common.Integer
	TotalCount                                  common.Integer
	VirtualBorderRouterForPhysicalConnectionSet DescribeVirtualBorderRoutersForPhysicalConnectionVirtualBorderRouterForPhysicalConnectionTypeList
}

type DescribeVirtualBorderRoutersForPhysicalConnectionVirtualBorderRouterForPhysicalConnectionType struct {
	VbrId           common.String
	VbrOwnerUid     common.Long
	CreationTime    common.String
	ActivationTime  common.String
	TerminationTime common.String
	RecoveryTime    common.String
	VlanId          common.Integer
	CircuitCode     common.String
}

type DescribeVirtualBorderRoutersForPhysicalConnectionFilterList []DescribeVirtualBorderRoutersForPhysicalConnectionFilter

func (list *DescribeVirtualBorderRoutersForPhysicalConnectionFilterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeVirtualBorderRoutersForPhysicalConnectionFilter)
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

type DescribeVirtualBorderRoutersForPhysicalConnectionValueList []string

func (list *DescribeVirtualBorderRoutersForPhysicalConnectionValueList) UnmarshalJSON(data []byte) error {
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

type DescribeVirtualBorderRoutersForPhysicalConnectionVirtualBorderRouterForPhysicalConnectionTypeList []DescribeVirtualBorderRoutersForPhysicalConnectionVirtualBorderRouterForPhysicalConnectionType

func (list *DescribeVirtualBorderRoutersForPhysicalConnectionVirtualBorderRouterForPhysicalConnectionTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeVirtualBorderRoutersForPhysicalConnectionVirtualBorderRouterForPhysicalConnectionType)
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
