package vpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId              string
	PageNumber             int
	PageSize               int
	TotalCount             int
	VirtualBorderRouterSet DescribeVirtualBorderRoutersVirtualBorderRouterTypeList
}

type DescribeVirtualBorderRoutersVirtualBorderRouterType struct {
	VbrId                            string
	CreationTime                     string
	ActivationTime                   string
	TerminationTime                  string
	RecoveryTime                     string
	Status                           string
	VlanId                           int
	CircuitCode                      string
	RouteTableId                     string
	VlanInterfaceId                  string
	LocalGatewayIp                   string
	PeerGatewayIp                    string
	PeeringSubnetMask                string
	PhysicalConnectionId             string
	PhysicalConnectionStatus         string
	PhysicalConnectionBusinessStatus string
	PhysicalConnectionOwnerUid       string
	AccessPointId                    string
	Name                             string
	Description                      string
	AssociatedPhysicalConnections    DescribeVirtualBorderRoutersAssociatedPhysicalConnectionList
	AssociatedCens                   DescribeVirtualBorderRoutersAssociatedCenList
}

type DescribeVirtualBorderRoutersAssociatedPhysicalConnection struct {
	CircuitCode                      string
	VlanInterfaceId                  string
	LocalGatewayIp                   string
	PeerGatewayIp                    string
	PeeringSubnetMask                string
	PhysicalConnectionId             string
	PhysicalConnectionStatus         string
	PhysicalConnectionBusinessStatus string
	PhysicalConnectionOwnerUid       string
	VlanId                           string
}

type DescribeVirtualBorderRoutersAssociatedCen struct {
	CenId      string
	CenOwnerId int64
	CenStatus  string
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
