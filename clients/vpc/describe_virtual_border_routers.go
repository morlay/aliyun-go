package vpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeVirtualBorderRoutersRequest struct {
	requests.RpcRequest
	Filters              *DescribeVirtualBorderRoutersFilterList `position:"Query" type:"Repeated" name:"Filter"`
	ResourceOwnerId      int64                                   `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string                                  `position:"Query" name:"ResourceOwnerAccount"`
	PageSize             int                                     `position:"Query" name:"PageSize"`
	OwnerId              int64                                   `position:"Query" name:"OwnerId"`
	PageNumber           int                                     `position:"Query" name:"PageNumber"`
}

func (req *DescribeVirtualBorderRoutersRequest) Invoke(client *sdk.Client) (resp *DescribeVirtualBorderRoutersResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "DescribeVirtualBorderRouters", "vpc", "")
	resp = &DescribeVirtualBorderRoutersResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeVirtualBorderRoutersFilter struct {
	Key    string                                 `name:"Key"`
	Values *DescribeVirtualBorderRoutersValueList `type:"Repeated" name:"Value"`
}

type DescribeVirtualBorderRoutersResponse struct {
	responses.BaseResponse
	RequestId              common.String
	PageNumber             common.Integer
	PageSize               common.Integer
	TotalCount             common.Integer
	VirtualBorderRouterSet DescribeVirtualBorderRoutersVirtualBorderRouterTypeList
}

type DescribeVirtualBorderRoutersVirtualBorderRouterType struct {
	VbrId                            common.String
	CreationTime                     common.String
	ActivationTime                   common.String
	TerminationTime                  common.String
	RecoveryTime                     common.String
	Status                           common.String
	VlanId                           common.Integer
	CircuitCode                      common.String
	RouteTableId                     common.String
	VlanInterfaceId                  common.String
	LocalGatewayIp                   common.String
	PeerGatewayIp                    common.String
	PeeringSubnetMask                common.String
	PhysicalConnectionId             common.String
	PhysicalConnectionStatus         common.String
	PhysicalConnectionBusinessStatus common.String
	PhysicalConnectionOwnerUid       common.String
	AccessPointId                    common.String
	Name                             common.String
	Description                      common.String
	AssociatedPhysicalConnections    DescribeVirtualBorderRoutersAssociatedPhysicalConnectionList
	AssociatedCens                   DescribeVirtualBorderRoutersAssociatedCenList
}

type DescribeVirtualBorderRoutersAssociatedPhysicalConnection struct {
	CircuitCode                      common.String
	VlanInterfaceId                  common.String
	LocalGatewayIp                   common.String
	PeerGatewayIp                    common.String
	PeeringSubnetMask                common.String
	PhysicalConnectionId             common.String
	PhysicalConnectionStatus         common.String
	PhysicalConnectionBusinessStatus common.String
	PhysicalConnectionOwnerUid       common.String
	VlanId                           common.String
}

type DescribeVirtualBorderRoutersAssociatedCen struct {
	CenId      common.String
	CenOwnerId common.Long
	CenStatus  common.String
}

type DescribeVirtualBorderRoutersFilterList []DescribeVirtualBorderRoutersFilter

func (list *DescribeVirtualBorderRoutersFilterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeVirtualBorderRoutersFilter)
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

type DescribeVirtualBorderRoutersValueList []string

func (list *DescribeVirtualBorderRoutersValueList) UnmarshalJSON(data []byte) error {
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

type DescribeVirtualBorderRoutersVirtualBorderRouterTypeList []DescribeVirtualBorderRoutersVirtualBorderRouterType

func (list *DescribeVirtualBorderRoutersVirtualBorderRouterTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeVirtualBorderRoutersVirtualBorderRouterType)
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

type DescribeVirtualBorderRoutersAssociatedPhysicalConnectionList []DescribeVirtualBorderRoutersAssociatedPhysicalConnection

func (list *DescribeVirtualBorderRoutersAssociatedPhysicalConnectionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeVirtualBorderRoutersAssociatedPhysicalConnection)
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

type DescribeVirtualBorderRoutersAssociatedCenList []DescribeVirtualBorderRoutersAssociatedCen

func (list *DescribeVirtualBorderRoutersAssociatedCenList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeVirtualBorderRoutersAssociatedCen)
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
