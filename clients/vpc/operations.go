package vpc

import (
	"encoding/json"

	"github.com/morlay/aliyun-go/core"
)

func (c *VpcClient) DeletePhysicalConnection(req *DeletePhysicalConnectionArgs) (resp *DeletePhysicalConnectionResponse, err error) {
	resp = &DeletePhysicalConnectionResponse{}
	err = c.Request("DeletePhysicalConnection", req, resp)
	return
}

type DeletePhysicalConnectionArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ClientToken          string
	PhysicalConnectionId string
	OwnerAccount         string
	UserCidr             string
	OwnerId              int64
}
type DeletePhysicalConnectionResponse struct {
	RequestId string
}

func (c *VpcClient) DescribeNewProjectEipMonitorData(req *DescribeNewProjectEipMonitorDataArgs) (resp *DescribeNewProjectEipMonitorDataResponse, err error) {
	resp = &DescribeNewProjectEipMonitorDataResponse{}
	err = c.Request("DescribeNewProjectEipMonitorData", req, resp)
	return
}

type DescribeNewProjectEipMonitorDataEipMonitorData struct {
	EipRX        int
	EipTX        int
	EipFlow      int
	EipBandwidth int
	EipPackets   int
	TimeStamp    string
}
type DescribeNewProjectEipMonitorDataArgs struct {
	ResourceOwnerId      int64
	Period               int
	ResourceOwnerAccount string
	OwnerAccount         string
	EndTime              string
	AllocationId         string
	StartTime            string
	OwnerId              int64
}
type DescribeNewProjectEipMonitorDataResponse struct {
	RequestId       string
	EipMonitorDatas DescribeNewProjectEipMonitorDataEipMonitorDataList
}

type DescribeNewProjectEipMonitorDataEipMonitorDataList []DescribeNewProjectEipMonitorDataEipMonitorData

func (list *DescribeNewProjectEipMonitorDataEipMonitorDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeNewProjectEipMonitorDataEipMonitorData)
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

func (c *VpcClient) DescribePhysicalConnections(req *DescribePhysicalConnectionsArgs) (resp *DescribePhysicalConnectionsResponse, err error) {
	resp = &DescribePhysicalConnectionsResponse{}
	err = c.Request("DescribePhysicalConnections", req, resp)
	return
}

type DescribePhysicalConnectionsFilter struct {
	Key    string
	Values DescribePhysicalConnectionsValueList
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
type DescribePhysicalConnectionsArgs struct {
	Filters              DescribePhysicalConnectionsFilterList
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	PageSize             int
	UserCidr             string
	OwnerId              int64
	PageNumber           int
}
type DescribePhysicalConnectionsResponse struct {
	RequestId             string
	PageNumber            int
	PageSize              int
	TotalCount            int
	PhysicalConnectionSet DescribePhysicalConnectionsPhysicalConnectionTypeList
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

func (c *VpcClient) DescribeNqas(req *DescribeNqasArgs) (resp *DescribeNqasResponse, err error) {
	resp = &DescribeNqasResponse{}
	err = c.Request("DescribeNqas", req, resp)
	return
}

type DescribeNqasNqa struct {
	NqaId         string
	RegionId      string
	Status        string
	RouterId      string
	DestinationIp string
}
type DescribeNqasArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	RouterId             string
	OwnerAccount         string
	PageSize             int
	NqaId                string
	IsDefault            core.Bool
	OwnerId              int64
	PageNumber           int
}
type DescribeNqasResponse struct {
	RequestId  string
	TotalCount int
	PageNumber int
	PageSize   int
	Nqas       DescribeNqasNqaList
}

type DescribeNqasNqaList []DescribeNqasNqa

func (list *DescribeNqasNqaList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeNqasNqa)
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

func (c *VpcClient) DescribeRouteTables(req *DescribeRouteTablesArgs) (resp *DescribeRouteTablesResponse, err error) {
	resp = &DescribeRouteTablesResponse{}
	err = c.Request("DescribeRouteTables", req, resp)
	return
}

type DescribeRouteTablesRouteTable struct {
	VRouterId      string
	RouteTableId   string
	RouteTableType string
	CreationTime   string
	RouteEntrys    DescribeRouteTablesRouteEntryList
}

type DescribeRouteTablesRouteEntry struct {
	RouteTableId         string
	DestinationCidrBlock string
	Type                 string
	Status               string
	InstanceId           string
	NextHopType          string
	NextHops             DescribeRouteTablesNextHopList
}

type DescribeRouteTablesNextHop struct {
	NextHopType string
	NextHopId   string
	Enabled     int
	Weight      int
}
type DescribeRouteTablesArgs struct {
	RouterType           string
	ResourceOwnerId      int64
	VRouterId            string
	ResourceOwnerAccount string
	RouterId             string
	OwnerAccount         string
	PageSize             int
	OwnerId              int64
	Type                 string
	PageNumber           int
	RouteTableId         string
}
type DescribeRouteTablesResponse struct {
	RequestId   string
	TotalCount  int
	PageNumber  int
	PageSize    int
	RouteTables DescribeRouteTablesRouteTableList
}

type DescribeRouteTablesRouteEntryList []DescribeRouteTablesRouteEntry

func (list *DescribeRouteTablesRouteEntryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRouteTablesRouteEntry)
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

type DescribeRouteTablesNextHopList []DescribeRouteTablesNextHop

func (list *DescribeRouteTablesNextHopList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRouteTablesNextHop)
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

type DescribeRouteTablesRouteTableList []DescribeRouteTablesRouteTable

func (list *DescribeRouteTablesRouteTableList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRouteTablesRouteTable)
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

func (c *VpcClient) CreateVirtualBorderRouter(req *CreateVirtualBorderRouterArgs) (resp *CreateVirtualBorderRouterResponse, err error) {
	resp = &CreateVirtualBorderRouterResponse{}
	err = c.Request("CreateVirtualBorderRouter", req, resp)
	return
}

type CreateVirtualBorderRouterArgs struct {
	ResourceOwnerId      int64
	CircuitCode          string
	VlanId               int
	ClientToken          string
	ResourceOwnerAccount string
	OwnerAccount         string
	Description          string
	OwnerId              int64
	PeerGatewayIp        string
	PeeringSubnetMask    string
	PhysicalConnectionId string
	Name                 string
	LocalGatewayIp       string
	UserCidr             string
	VbrOwnerId           int64
}
type CreateVirtualBorderRouterResponse struct {
	RequestId string
	VbrId     string
}

func (c *VpcClient) UnassociateHaVip(req *UnassociateHaVipArgs) (resp *UnassociateHaVipResponse, err error) {
	resp = &UnassociateHaVipResponse{}
	err = c.Request("UnassociateHaVip", req, resp)
	return
}

type UnassociateHaVipArgs struct {
	HaVipId              string
	ResourceOwnerId      int64
	InstanceId           string
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	Force                string
	OwnerId              int64
}
type UnassociateHaVipResponse struct {
	RequestId string
}

func (c *VpcClient) DescribeCustomerGateway(req *DescribeCustomerGatewayArgs) (resp *DescribeCustomerGatewayResponse, err error) {
	resp = &DescribeCustomerGatewayResponse{}
	err = c.Request("DescribeCustomerGateway", req, resp)
	return
}

type DescribeCustomerGatewayArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	CustomerGatewayId    string
}
type DescribeCustomerGatewayResponse struct {
	RequestId         string
	CustomerGatewayId string
	IpAddress         string
	Name              string
	Description       string
	CreateTime        int64
}

func (c *VpcClient) DeleteRouterInterface(req *DeleteRouterInterfaceArgs) (resp *DeleteRouterInterfaceResponse, err error) {
	resp = &DeleteRouterInterfaceResponse{}
	err = c.Request("DeleteRouterInterface", req, resp)
	return
}

type DeleteRouterInterfaceArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	UserCidr             string
	RouterInterfaceId    string
	OwnerId              int64
}
type DeleteRouterInterfaceResponse struct {
	RequestId string
}

func (c *VpcClient) EnablePhysicalConnection(req *EnablePhysicalConnectionArgs) (resp *EnablePhysicalConnectionResponse, err error) {
	resp = &EnablePhysicalConnectionResponse{}
	err = c.Request("EnablePhysicalConnection", req, resp)
	return
}

type EnablePhysicalConnectionArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ClientToken          string
	PhysicalConnectionId string
	OwnerAccount         string
	UserCidr             string
	OwnerId              int64
}
type EnablePhysicalConnectionResponse struct {
	RequestId string
}

func (c *VpcClient) AddCommonBandwidthPackageIp(req *AddCommonBandwidthPackageIpArgs) (resp *AddCommonBandwidthPackageIpResponse, err error) {
	resp = &AddCommonBandwidthPackageIpResponse{}
	err = c.Request("AddCommonBandwidthPackageIp", req, resp)
	return
}

type AddCommonBandwidthPackageIpArgs struct {
	ResourceOwnerId      int64
	BandwidthPackageId   string
	ResourceOwnerAccount string
	OwnerAccount         string
	IpInstanceId         string
	OwnerId              int64
}
type AddCommonBandwidthPackageIpResponse struct {
	RequestId string
}

func (c *VpcClient) CreateBandwidthPackage(req *CreateBandwidthPackageArgs) (resp *CreateBandwidthPackageResponse, err error) {
	resp = &CreateBandwidthPackageResponse{}
	err = c.Request("CreateBandwidthPackage", req, resp)
	return
}

type CreateBandwidthPackageArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ClientToken          string
	Bandwidth            int
	OwnerAccount         string
	ISP                  string
	Description          string
	OwnerId              int64
	Zone                 string
	InternetChargeType   string
	Name                 string
	NatGatewayId         string
	IpCount              int
}
type CreateBandwidthPackageResponse struct {
	RequestId          string
	BandwidthPackageId string
}

func (c *VpcClient) UnassociateGlobalAccelerationInstance(req *UnassociateGlobalAccelerationInstanceArgs) (resp *UnassociateGlobalAccelerationInstanceResponse, err error) {
	resp = &UnassociateGlobalAccelerationInstanceResponse{}
	err = c.Request("UnassociateGlobalAccelerationInstance", req, resp)
	return
}

type UnassociateGlobalAccelerationInstanceArgs struct {
	ResourceOwnerId              int64
	ResourceOwnerAccount         string
	OwnerAccount                 string
	InstanceType                 string
	OwnerId                      int64
	GlobalAccelerationInstanceId string
}
type UnassociateGlobalAccelerationInstanceResponse struct {
	RequestId string
}

func (c *VpcClient) ModifyBandwidthPackageSpec(req *ModifyBandwidthPackageSpecArgs) (resp *ModifyBandwidthPackageSpecResponse, err error) {
	resp = &ModifyBandwidthPackageSpecResponse{}
	err = c.Request("ModifyBandwidthPackageSpec", req, resp)
	return
}

type ModifyBandwidthPackageSpecArgs struct {
	ResourceOwnerId      int64
	BandwidthPackageId   string
	ResourceOwnerAccount string
	Bandwidth            string
	OwnerAccount         string
	OwnerId              int64
}
type ModifyBandwidthPackageSpecResponse struct {
	RequestId string
}

func (c *VpcClient) DescribeBgpGroups(req *DescribeBgpGroupsArgs) (resp *DescribeBgpGroupsResponse, err error) {
	resp = &DescribeBgpGroupsResponse{}
	err = c.Request("DescribeBgpGroups", req, resp)
	return
}

type DescribeBgpGroupsBgpGroup struct {
	Name        string
	Description string
	BgpGroupId  string
	PeerAsn     string
	AuthKey     string
	RouterId    string
	Status      string
	Keepalive   string
	LocalAsn    string
	Hold        string
	IsFake      string
	RouteLimit  string
	RegionId    string
}
type DescribeBgpGroupsArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	RouterId             string
	OwnerAccount         string
	PageSize             int
	BgpGroupId           string
	IsDefault            core.Bool
	OwnerId              int64
	PageNumber           int
}
type DescribeBgpGroupsResponse struct {
	RequestId  string
	TotalCount int
	PageNumber int
	PageSize   int
	BgpGroups  DescribeBgpGroupsBgpGroupList
}

type DescribeBgpGroupsBgpGroupList []DescribeBgpGroupsBgpGroup

func (list *DescribeBgpGroupsBgpGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeBgpGroupsBgpGroup)
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

func (c *VpcClient) DeleteRouteEntry(req *DeleteRouteEntryArgs) (resp *DeleteRouteEntryResponse, err error) {
	resp = &DeleteRouteEntryResponse{}
	err = c.Request("DeleteRouteEntry", req, resp)
	return
}

type DeleteRouteEntryNextHopList struct {
	NextHopType string
	NextHopId   string
}
type DeleteRouteEntryArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	DestinationCidrBlock string
	OwnerAccount         string
	NextHopId            string
	OwnerId              int64
	NextHopLists         DeleteRouteEntryNextHopListList
	RouteTableId         string
}
type DeleteRouteEntryResponse struct {
	RequestId string
}

type DeleteRouteEntryNextHopListList []DeleteRouteEntryNextHopList

func (list *DeleteRouteEntryNextHopListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DeleteRouteEntryNextHopList)
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

func (c *VpcClient) CreateSnatEntry(req *CreateSnatEntryArgs) (resp *CreateSnatEntryResponse, err error) {
	resp = &CreateSnatEntryResponse{}
	err = c.Request("CreateSnatEntry", req, resp)
	return
}

type CreateSnatEntryArgs struct {
	ResourceOwnerId      int64
	SourceVSwitchId      string
	ResourceOwnerAccount string
	OwnerAccount         string
	SourceCIDR           string
	SnatTableId          string
	OwnerId              int64
	SnatIp               string
}
type CreateSnatEntryResponse struct {
	RequestId   string
	SnatEntryId string
}

func (c *VpcClient) DescribeVpnConnections(req *DescribeVpnConnectionsArgs) (resp *DescribeVpnConnectionsResponse, err error) {
	resp = &DescribeVpnConnectionsResponse{}
	err = c.Request("DescribeVpnConnections", req, resp)
	return
}

type DescribeVpnConnectionsVpnConnection struct {
	VpnConnectionId   string
	CustomerGatewayId string
	VpnGatewayId      string
	Name              string
	LocalSubnet       string
	RemoteSubnet      string
	CreateTime        int64
	EffectImmediately core.Bool
	Status            string
	IkeConfig         DescribeVpnConnectionsIkeConfig
	IpsecConfig       DescribeVpnConnectionsIpsecConfig
}

type DescribeVpnConnectionsIkeConfig struct {
	Psk         string
	IkeVersion  string
	IkeMode     string
	IkeEncAlg   string
	IkeAuthAlg  string
	IkePfs      string
	IkeLifetime int64
	LocalId     string
	RemoteId    string
}

type DescribeVpnConnectionsIpsecConfig struct {
	IpsecEncAlg   string
	IpsecAuthAlg  string
	IpsecPfs      string
	IpsecLifetime int64
}
type DescribeVpnConnectionsArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	PageSize             int
	VpnGatewayId         string
	OwnerId              int64
	CustomerGatewayId    string
	PageNumber           int
}
type DescribeVpnConnectionsResponse struct {
	RequestId      string
	TotalCount     int
	PageNumber     int
	PageSize       int
	VpnConnections DescribeVpnConnectionsVpnConnectionList
}

type DescribeVpnConnectionsVpnConnectionList []DescribeVpnConnectionsVpnConnection

func (list *DescribeVpnConnectionsVpnConnectionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeVpnConnectionsVpnConnection)
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

func (c *VpcClient) DeleteVpc(req *DeleteVpcArgs) (resp *DeleteVpcResponse, err error) {
	resp = &DeleteVpcResponse{}
	err = c.Request("DeleteVpc", req, resp)
	return
}

type DeleteVpcArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	VpcId                string
	OwnerAccount         string
	OwnerId              int64
}
type DeleteVpcResponse struct {
	RequestId string
}

func (c *VpcClient) DescribeEipAddresses(req *DescribeEipAddressesArgs) (resp *DescribeEipAddressesResponse, err error) {
	resp = &DescribeEipAddressesResponse{}
	err = c.Request("DescribeEipAddresses", req, resp)
	return
}

type DescribeEipAddressesEipAddress struct {
	RegionId           string
	IpAddress          string
	AllocationId       string
	Status             string
	InstanceId         string
	Bandwidth          string
	InternetChargeType string
	AllocationTime     string
	InstanceType       string
	ChargeType         string
	ExpiredTime        string
	Name               string
	Descritpion        string
	BandwidthPackageId string
	OperationLocks     DescribeEipAddressesLockReasonList
}

type DescribeEipAddressesLockReason struct {
	LockReason string
}
type DescribeEipAddressesArgs struct {
	ResourceOwnerId        int64
	ResourceOwnerAccount   string
	Filter2Value           string
	OwnerAccount           string
	AllocationId           string
	Filter1Value           string
	Filter2Key             string
	OwnerId                int64
	EipAddress             string
	PageNumber             int
	LockReason             string
	Filter1Key             string
	AssociatedInstanceType string
	PageSize               int
	ChargeType             string
	AssociatedInstanceId   string
	Status                 string
}
type DescribeEipAddressesResponse struct {
	RequestId    string
	TotalCount   int
	PageNumber   int
	PageSize     int
	EipAddresses DescribeEipAddressesEipAddressList
}

type DescribeEipAddressesLockReasonList []DescribeEipAddressesLockReason

func (list *DescribeEipAddressesLockReasonList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeEipAddressesLockReason)
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

type DescribeEipAddressesEipAddressList []DescribeEipAddressesEipAddress

func (list *DescribeEipAddressesEipAddressList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeEipAddressesEipAddress)
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

func (c *VpcClient) DescribeGlobalAccelerationInstances(req *DescribeGlobalAccelerationInstancesArgs) (resp *DescribeGlobalAccelerationInstancesResponse, err error) {
	resp = &DescribeGlobalAccelerationInstancesResponse{}
	err = c.Request("DescribeGlobalAccelerationInstances", req, resp)
	return
}

type DescribeGlobalAccelerationInstancesGlobalAccelerationInstance struct {
	RegionId                     string
	GlobalAccelerationInstanceId string
	IpAddress                    string
	Status                       string
	Bandwidth                    string
	InternetChargeType           string
	ChargeType                   string
	AccelerationLocation         string
	ServiceLocation              string
	Name                         string
	Description                  string
	ExpiredTime                  string
	CreationTime                 string
	OperationLocks               DescribeGlobalAccelerationInstancesLockReasonList
	BackendServers               DescribeGlobalAccelerationInstancesBackendServerList
}

type DescribeGlobalAccelerationInstancesLockReason struct {
	LockReason string
}

type DescribeGlobalAccelerationInstancesBackendServer struct {
	RegionId        string
	ServerId        string
	ServerIpAddress string
	ServerType      string
}
type DescribeGlobalAccelerationInstancesArgs struct {
	IpAddress                    string
	ResourceOwnerId              int64
	ResourceOwnerAccount         string
	ServiceLocation              string
	OwnerAccount                 string
	OwnerId                      int64
	GlobalAccelerationInstanceId string
	ServerId                     string
	PageNumber                   int
	Name                         string
	PageSize                     int
	Status                       string
}
type DescribeGlobalAccelerationInstancesResponse struct {
	RequestId                   string
	TotalCount                  int
	PageNumber                  int
	PageSize                    int
	GlobalAccelerationInstances DescribeGlobalAccelerationInstancesGlobalAccelerationInstanceList
}

type DescribeGlobalAccelerationInstancesLockReasonList []DescribeGlobalAccelerationInstancesLockReason

func (list *DescribeGlobalAccelerationInstancesLockReasonList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeGlobalAccelerationInstancesLockReason)
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

type DescribeGlobalAccelerationInstancesBackendServerList []DescribeGlobalAccelerationInstancesBackendServer

func (list *DescribeGlobalAccelerationInstancesBackendServerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeGlobalAccelerationInstancesBackendServer)
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

type DescribeGlobalAccelerationInstancesGlobalAccelerationInstanceList []DescribeGlobalAccelerationInstancesGlobalAccelerationInstance

func (list *DescribeGlobalAccelerationInstancesGlobalAccelerationInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeGlobalAccelerationInstancesGlobalAccelerationInstance)
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

func (c *VpcClient) DescribeVpnGateways(req *DescribeVpnGatewaysArgs) (resp *DescribeVpnGatewaysResponse, err error) {
	resp = &DescribeVpnGatewaysResponse{}
	err = c.Request("DescribeVpnGateways", req, resp)
	return
}

type DescribeVpnGatewaysVpnGateway struct {
	VpnGatewayId   string
	VpcId          string
	VSwitchId      string
	InternetIp     string
	CreateTime     int64
	EndTime        int64
	Spec           string
	Name           string
	Description    string
	Status         string
	BusinessStatus string
	ChargeType     string
}
type DescribeVpnGatewaysArgs struct {
	BusinessStatus       string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	VpcId                string
	PageSize             int
	VpnGatewayId         string
	OwnerId              int64
	PageNumber           int
	Status               string
}
type DescribeVpnGatewaysResponse struct {
	RequestId   string
	TotalCount  int
	PageNumber  int
	PageSize    int
	VpnGateways DescribeVpnGatewaysVpnGatewayList
}

type DescribeVpnGatewaysVpnGatewayList []DescribeVpnGatewaysVpnGateway

func (list *DescribeVpnGatewaysVpnGatewayList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeVpnGatewaysVpnGateway)
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

func (c *VpcClient) CancelPhysicalConnection(req *CancelPhysicalConnectionArgs) (resp *CancelPhysicalConnectionResponse, err error) {
	resp = &CancelPhysicalConnectionResponse{}
	err = c.Request("CancelPhysicalConnection", req, resp)
	return
}

type CancelPhysicalConnectionArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ClientToken          string
	PhysicalConnectionId string
	OwnerAccount         string
	UserCidr             string
	OwnerId              int64
}
type CancelPhysicalConnectionResponse struct {
	RequestId string
}

func (c *VpcClient) ModifyGlobalAccelerationInstanceSpec(req *ModifyGlobalAccelerationInstanceSpecArgs) (resp *ModifyGlobalAccelerationInstanceSpecResponse, err error) {
	resp = &ModifyGlobalAccelerationInstanceSpecResponse{}
	err = c.Request("ModifyGlobalAccelerationInstanceSpec", req, resp)
	return
}

type ModifyGlobalAccelerationInstanceSpecArgs struct {
	ResourceOwnerId              int64
	ResourceOwnerAccount         string
	Bandwidth                    string
	OwnerAccount                 string
	OwnerId                      int64
	GlobalAccelerationInstanceId string
}
type ModifyGlobalAccelerationInstanceSpecResponse struct {
	RequestId string
}

func (c *VpcClient) ConnectRouterInterface(req *ConnectRouterInterfaceArgs) (resp *ConnectRouterInterfaceResponse, err error) {
	resp = &ConnectRouterInterfaceResponse{}
	err = c.Request("ConnectRouterInterface", req, resp)
	return
}

type ConnectRouterInterfaceArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerId              int64
	RouterInterfaceId    string
}
type ConnectRouterInterfaceResponse struct {
	RequestId string
}

func (c *VpcClient) DescribeVRouters(req *DescribeVRoutersArgs) (resp *DescribeVRoutersResponse, err error) {
	resp = &DescribeVRoutersResponse{}
	err = c.Request("DescribeVRouters", req, resp)
	return
}

type DescribeVRoutersVRouter struct {
	RegionId      string
	VpcId         string
	VRouterName   string
	Description   string
	VRouterId     string
	CreationTime  string
	RouteTableIds DescribeVRoutersRouteTableIdList
}
type DescribeVRoutersArgs struct {
	ResourceOwnerId      int64
	VRouterId            string
	ResourceOwnerAccount string
	OwnerAccount         string
	PageSize             int
	OwnerId              int64
	PageNumber           int
}
type DescribeVRoutersResponse struct {
	RequestId  string
	TotalCount int
	PageNumber int
	PageSize   int
	VRouters   DescribeVRoutersVRouterList
}

type DescribeVRoutersRouteTableIdList []string

func (list *DescribeVRoutersRouteTableIdList) UnmarshalJSON(data []byte) error {
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

type DescribeVRoutersVRouterList []DescribeVRoutersVRouter

func (list *DescribeVRoutersVRouterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeVRoutersVRouter)
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

func (c *VpcClient) CreateRouterInterface(req *CreateRouterInterfaceArgs) (resp *CreateRouterInterfaceResponse, err error) {
	resp = &CreateRouterInterfaceResponse{}
	err = c.Request("CreateRouterInterface", req, resp)
	return
}

type CreateRouterInterfaceArgs struct {
	AccessPointId            string
	OppositeRouterId         string
	OppositeAccessPointId    string
	ResourceOwnerId          int64
	Role                     string
	ClientToken              string
	ResourceOwnerAccount     string
	OppositeRegionId         string
	OwnerAccount             string
	HealthCheckTargetIp      string
	Description              string
	OwnerId                  int64
	Spec                     string
	OppositeInterfaceOwnerId string
	RouterType               string
	HealthCheckSourceIp      string
	RouterId                 string
	OppositeRouterType       string
	Name                     string
	UserCidr                 string
	OppositeInterfaceId      string
}
type CreateRouterInterfaceResponse struct {
	RequestId         string
	RouterInterfaceId string
}

func (c *VpcClient) CreateForwardEntry(req *CreateForwardEntryArgs) (resp *CreateForwardEntryResponse, err error) {
	resp = &CreateForwardEntryResponse{}
	err = c.Request("CreateForwardEntry", req, resp)
	return
}

type CreateForwardEntryArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	IpProtocol           string
	InternalPort         string
	OwnerAccount         string
	ForwardTableId       string
	OwnerId              int64
	ExternalIp           string
	ExternalPort         string
	InternalIp           string
}
type CreateForwardEntryResponse struct {
	RequestId      string
	ForwardEntryId string
}

func (c *VpcClient) AssociateHaVip(req *AssociateHaVipArgs) (resp *AssociateHaVipResponse, err error) {
	resp = &AssociateHaVipResponse{}
	err = c.Request("AssociateHaVip", req, resp)
	return
}

type AssociateHaVipArgs struct {
	HaVipId              string
	ResourceOwnerId      int64
	InstanceId           string
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	OwnerId              int64
}
type AssociateHaVipResponse struct {
	RequestId string
}

func (c *VpcClient) DescribeCustomerGateways(req *DescribeCustomerGatewaysArgs) (resp *DescribeCustomerGatewaysResponse, err error) {
	resp = &DescribeCustomerGatewaysResponse{}
	err = c.Request("DescribeCustomerGateways", req, resp)
	return
}

type DescribeCustomerGatewaysCustomerGateway struct {
	CustomerGatewayId string
	Name              string
	IpAddress         string
	Description       string
	CreateTime        int64
}
type DescribeCustomerGatewaysArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	PageSize             int
	OwnerId              int64
	CustomerGatewayId    string
	PageNumber           int
}
type DescribeCustomerGatewaysResponse struct {
	RequestId        string
	TotalCount       int
	PageNumber       int
	PageSize         int
	CustomerGateways DescribeCustomerGatewaysCustomerGatewayList
}

type DescribeCustomerGatewaysCustomerGatewayList []DescribeCustomerGatewaysCustomerGateway

func (list *DescribeCustomerGatewaysCustomerGatewayList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCustomerGatewaysCustomerGateway)
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

func (c *VpcClient) AssociatePhysicalConnectionToVirtualBorderRouter(req *AssociatePhysicalConnectionToVirtualBorderRouterArgs) (resp *AssociatePhysicalConnectionToVirtualBorderRouterResponse, err error) {
	resp = &AssociatePhysicalConnectionToVirtualBorderRouterResponse{}
	err = c.Request("AssociatePhysicalConnectionToVirtualBorderRouter", req, resp)
	return
}

type AssociatePhysicalConnectionToVirtualBorderRouterArgs struct {
	ResourceOwnerId      int64
	CircuitCode          string
	VlanId               string
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	VbrId                string
	OwnerId              int64
	PeerGatewayIp        string
	PeeringSubnetMask    string
	PhysicalConnectionId string
	LocalGatewayIp       string
	UserCidr             string
}
type AssociatePhysicalConnectionToVirtualBorderRouterResponse struct {
	RequestId string
}

func (c *VpcClient) ModifyHaVipAttribute(req *ModifyHaVipAttributeArgs) (resp *ModifyHaVipAttributeResponse, err error) {
	resp = &ModifyHaVipAttributeResponse{}
	err = c.Request("ModifyHaVipAttribute", req, resp)
	return
}

type ModifyHaVipAttributeArgs struct {
	HaVipId              string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	Description          string
	OwnerId              int64
}
type ModifyHaVipAttributeResponse struct {
	RequestId string
}

func (c *VpcClient) CreateGlobalAccelerationInstance(req *CreateGlobalAccelerationInstanceArgs) (resp *CreateGlobalAccelerationInstanceResponse, err error) {
	resp = &CreateGlobalAccelerationInstanceResponse{}
	err = c.Request("CreateGlobalAccelerationInstance", req, resp)
	return
}

type CreateGlobalAccelerationInstanceArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ServiceLocation      string
	Bandwidth            string
	ClientToken          string
	InternetChargeType   string
	OwnerAccount         string
	Name                 string
	Description          string
	OwnerId              int64
}
type CreateGlobalAccelerationInstanceResponse struct {
	RequestId                    string
	GlobalAccelerationInstanceId string
	IpAddress                    string
}

func (c *VpcClient) RemoveBandwidthPackageIps(req *RemoveBandwidthPackageIpsArgs) (resp *RemoveBandwidthPackageIpsResponse, err error) {
	resp = &RemoveBandwidthPackageIpsResponse{}
	err = c.Request("RemoveBandwidthPackageIps", req, resp)
	return
}

type RemoveBandwidthPackageIpsArgs struct {
	RemovedIpAddressess  RemoveBandwidthPackageIpsRemovedIpAddressesList
	ResourceOwnerId      int64
	BandwidthPackageId   string
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	OwnerId              int64
}
type RemoveBandwidthPackageIpsResponse struct {
	RequestId string
}

type RemoveBandwidthPackageIpsRemovedIpAddressesList []string

func (list *RemoveBandwidthPackageIpsRemovedIpAddressesList) UnmarshalJSON(data []byte) error {
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

func (c *VpcClient) DeleteVirtualBorderRouter(req *DeleteVirtualBorderRouterArgs) (resp *DeleteVirtualBorderRouterResponse, err error) {
	resp = &DeleteVirtualBorderRouterResponse{}
	err = c.Request("DeleteVirtualBorderRouter", req, resp)
	return
}

type DeleteVirtualBorderRouterArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	UserCidr             string
	VbrId                string
	OwnerId              int64
}
type DeleteVirtualBorderRouterResponse struct {
	RequestId string
}

func (c *VpcClient) ActivateRouterInterface(req *ActivateRouterInterfaceArgs) (resp *ActivateRouterInterfaceResponse, err error) {
	resp = &ActivateRouterInterfaceResponse{}
	err = c.Request("ActivateRouterInterface", req, resp)
	return
}

type ActivateRouterInterfaceArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerId              int64
	RouterInterfaceId    string
}
type ActivateRouterInterfaceResponse struct {
	RequestId string
}

func (c *VpcClient) ModifyVRouterAttribute(req *ModifyVRouterAttributeArgs) (resp *ModifyVRouterAttributeResponse, err error) {
	resp = &ModifyVRouterAttributeResponse{}
	err = c.Request("ModifyVRouterAttribute", req, resp)
	return
}

type ModifyVRouterAttributeArgs struct {
	VRouterName          string
	ResourceOwnerId      int64
	VRouterId            string
	ResourceOwnerAccount string
	OwnerAccount         string
	Description          string
	OwnerId              int64
}
type ModifyVRouterAttributeResponse struct {
	RequestId string
}

func (c *VpcClient) DeleteVSwitch(req *DeleteVSwitchArgs) (resp *DeleteVSwitchResponse, err error) {
	resp = &DeleteVSwitchResponse{}
	err = c.Request("DeleteVSwitch", req, resp)
	return
}

type DeleteVSwitchArgs struct {
	VSwitchId            string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
}
type DeleteVSwitchResponse struct {
	RequestId string
}

func (c *VpcClient) ModifyVSwitchAttribute(req *ModifyVSwitchAttributeArgs) (resp *ModifyVSwitchAttributeResponse, err error) {
	resp = &ModifyVSwitchAttributeResponse{}
	err = c.Request("ModifyVSwitchAttribute", req, resp)
	return
}

type ModifyVSwitchAttributeArgs struct {
	VSwitchId            string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	VSwitchName          string
	OwnerAccount         string
	Description          string
	OwnerId              int64
}
type ModifyVSwitchAttributeResponse struct {
	RequestId string
}

func (c *VpcClient) CreateVSwitch(req *CreateVSwitchArgs) (resp *CreateVSwitchResponse, err error) {
	resp = &CreateVSwitchResponse{}
	err = c.Request("CreateVSwitch", req, resp)
	return
}

type CreateVSwitchArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ClientToken          string
	VpcId                string
	VSwitchName          string
	OwnerAccount         string
	CidrBlock            string
	ZoneId               string
	Description          string
	OwnerId              int64
}
type CreateVSwitchResponse struct {
	RequestId string
	VSwitchId string
}

func (c *VpcClient) EnableVpcClassicLink(req *EnableVpcClassicLinkArgs) (resp *EnableVpcClassicLinkResponse, err error) {
	resp = &EnableVpcClassicLinkResponse{}
	err = c.Request("EnableVpcClassicLink", req, resp)
	return
}

type EnableVpcClassicLinkArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ClientToken          string
	VpcId                string
	OwnerAccount         string
	OwnerId              int64
}
type EnableVpcClassicLinkResponse struct {
	RequestId string
}

func (c *VpcClient) DeleteNatGateway(req *DeleteNatGatewayArgs) (resp *DeleteNatGatewayResponse, err error) {
	resp = &DeleteNatGatewayResponse{}
	err = c.Request("DeleteNatGateway", req, resp)
	return
}

type DeleteNatGatewayArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	Force                core.Bool
	NatGatewayId         string
	OwnerId              int64
}
type DeleteNatGatewayResponse struct {
	RequestId string
}

func (c *VpcClient) CreateNqa(req *CreateNqaArgs) (resp *CreateNqaResponse, err error) {
	resp = &CreateNqaResponse{}
	err = c.Request("CreateNqa", req, resp)
	return
}

type CreateNqaArgs struct {
	DestinationIp        string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ClientToken          string
	RouterId             string
	OwnerAccount         string
	OwnerId              int64
}
type CreateNqaResponse struct {
	RequestId string
	NqaId     string
}

func (c *VpcClient) ModifyNatGatewaySpec(req *ModifyNatGatewaySpecArgs) (resp *ModifyNatGatewaySpecResponse, err error) {
	resp = &ModifyNatGatewaySpecResponse{}
	err = c.Request("ModifyNatGatewaySpec", req, resp)
	return
}

type ModifyNatGatewaySpecArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	NatGatewayId         string
	OwnerId              int64
	Spec                 string
}
type ModifyNatGatewaySpecResponse struct {
	RequestId string
}

func (c *VpcClient) ModifySnatEntry(req *ModifySnatEntryArgs) (resp *ModifySnatEntryResponse, err error) {
	resp = &ModifySnatEntryResponse{}
	err = c.Request("ModifySnatEntry", req, resp)
	return
}

type ModifySnatEntryArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	SnatTableId          string
	SnatEntryId          string
	OwnerId              int64
	SnatIp               string
}
type ModifySnatEntryResponse struct {
	RequestId string
}

func (c *VpcClient) RemoveCommonBandwidthPackageIp(req *RemoveCommonBandwidthPackageIpArgs) (resp *RemoveCommonBandwidthPackageIpResponse, err error) {
	resp = &RemoveCommonBandwidthPackageIpResponse{}
	err = c.Request("RemoveCommonBandwidthPackageIp", req, resp)
	return
}

type RemoveCommonBandwidthPackageIpArgs struct {
	ResourceOwnerId      int64
	BandwidthPackageId   string
	ResourceOwnerAccount string
	OwnerAccount         string
	IpInstanceId         string
	OwnerId              int64
}
type RemoveCommonBandwidthPackageIpResponse struct {
	RequestId string
}

func (c *VpcClient) DeleteCustomerGateway(req *DeleteCustomerGatewayArgs) (resp *DeleteCustomerGatewayResponse, err error) {
	resp = &DeleteCustomerGatewayResponse{}
	err = c.Request("DeleteCustomerGateway", req, resp)
	return
}

type DeleteCustomerGatewayArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	OwnerId              int64
	CustomerGatewayId    string
}
type DeleteCustomerGatewayResponse struct {
	RequestId string
}

func (c *VpcClient) CreateBgpGroup(req *CreateBgpGroupArgs) (resp *CreateBgpGroupResponse, err error) {
	resp = &CreateBgpGroupResponse{}
	err = c.Request("CreateBgpGroup", req, resp)
	return
}

type CreateBgpGroupArgs struct {
	AuthKey              string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	Description          string
	OwnerId              int64
	PeerAsn              int64
	IsFakeAsn            core.Bool
	RouterId             string
	Name                 string
}
type CreateBgpGroupResponse struct {
	RequestId  string
	BgpGroupId string
}

func (c *VpcClient) ModifyForwardEntry(req *ModifyForwardEntryArgs) (resp *ModifyForwardEntryResponse, err error) {
	resp = &ModifyForwardEntryResponse{}
	err = c.Request("ModifyForwardEntry", req, resp)
	return
}

type ModifyForwardEntryArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	IpProtocol           string
	OwnerAccount         string
	ForwardTableId       string
	OwnerId              int64
	InternalIp           string
	ForwardEntryId       string
	InternalPort         string
	ExternalIp           string
	ExternalPort         string
}
type ModifyForwardEntryResponse struct {
	RequestId string
}

func (c *VpcClient) ModifyRouterInterfaceAttribute(req *ModifyRouterInterfaceAttributeArgs) (resp *ModifyRouterInterfaceAttributeResponse, err error) {
	resp = &ModifyRouterInterfaceAttributeResponse{}
	err = c.Request("ModifyRouterInterfaceAttribute", req, resp)
	return
}

type ModifyRouterInterfaceAttributeArgs struct {
	OppositeRouterId         string
	ResourceOwnerId          int64
	ResourceOwnerAccount     string
	Description              string
	HealthCheckTargetIp      string
	OwnerId                  int64
	RouterInterfaceId        string
	OppositeInterfaceOwnerId int64
	HealthCheckSourceIp      string
	Name                     string
	OppositeRouterType       string
	OppositeInterfaceId      string
}
type ModifyRouterInterfaceAttributeResponse struct {
	RequestId string
}

func (c *VpcClient) DeleteHaVip(req *DeleteHaVipArgs) (resp *DeleteHaVipResponse, err error) {
	resp = &DeleteHaVipResponse{}
	err = c.Request("DeleteHaVip", req, resp)
	return
}

type DeleteHaVipArgs struct {
	HaVipId              string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	OwnerId              int64
}
type DeleteHaVipResponse struct {
	RequestId string
}

func (c *VpcClient) DescribeRouterInterfacesForGlobal(req *DescribeRouterInterfacesForGlobalArgs) (resp *DescribeRouterInterfacesForGlobalResponse, err error) {
	resp = &DescribeRouterInterfacesForGlobalResponse{}
	err = c.Request("DescribeRouterInterfacesForGlobal", req, resp)
	return
}

type DescribeRouterInterfacesForGlobalRouterInterfaceType struct {
	BusinessStatus                  string
	AccessPointId                   string
	ChargeType                      string
	ConnectedTime                   string
	CreationTime                    string
	RouterInterfaceId               string
	OppositeInterfaceBusinessStatus string
	OppositeInterfaceId             string
	OppositeInterfaceOwnerId        int64
	OppositeInterfaceSpec           string
	OppositeInterfaceStatus         string
	OppositeRegionId                string
	OppositeAccessPointId           string
	OppositeRouterId                string
	OppositeRouterType              string
	OppositeVpcInstanceId           string
	RegionId                        string
	Role                            string
	RouterId                        string
	RouterType                      string
	Spec                            string
	Status                          string
	VpcInstanceId                   string
	Name                            string
	Description                     string
	HealthCheckSourceIp             string
	HealthCheckTargetIp             string
}
type DescribeRouterInterfacesForGlobalArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	PageSize             int
	OwnerId              int64
	PageNumber           int
	Status               string
}
type DescribeRouterInterfacesForGlobalResponse struct {
	RequestId          string
	Code               string
	Message            string
	Desc               string
	Success            core.Bool
	PageSize           int
	PageNumber         int
	TotalCount         int
	RouterInterfaceSet DescribeRouterInterfacesForGlobalRouterInterfaceTypeList
}

type DescribeRouterInterfacesForGlobalRouterInterfaceTypeList []DescribeRouterInterfacesForGlobalRouterInterfaceType

func (list *DescribeRouterInterfacesForGlobalRouterInterfaceTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRouterInterfacesForGlobalRouterInterfaceType)
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

func (c *VpcClient) CreatePhysicalConnectionNew(req *CreatePhysicalConnectionNewArgs) (resp *CreatePhysicalConnectionNewResponse, err error) {
	resp = &CreatePhysicalConnectionNewResponse{}
	err = c.Request("CreatePhysicalConnectionNew", req, resp)
	return
}

type CreatePhysicalConnectionNewArgs struct {
	AccessPointId                 string
	RedundantPhysicalConnectionId string
	PeerLocation                  string
	ResourceOwnerId               int64
	PortType                      string
	CircuitCode                   string
	Bandwidth                     int
	ClientToken                   string
	ResourceOwnerAccount          string
	OwnerAccount                  string
	Description                   string
	InterfaceName                 string
	Type                          string
	OwnerId                       int64
	LineOperator                  string
	Name                          string
	DeviceName                    string
	UserCidr                      string
}
type CreatePhysicalConnectionNewResponse struct {
	RequestId            string
	PhysicalConnectionId string
}

func (c *VpcClient) DescribeCommonBandwidthPackages(req *DescribeCommonBandwidthPackagesArgs) (resp *DescribeCommonBandwidthPackagesResponse, err error) {
	resp = &DescribeCommonBandwidthPackagesResponse{}
	err = c.Request("DescribeCommonBandwidthPackages", req, resp)
	return
}

type DescribeCommonBandwidthPackagesCommonBandwidthPackage struct {
	BandwidthPackageId string
	RegionId           string
	Name               string
	Description        string
	Bandwidth          string
	InstanceChargeType string
	InternetChargeType string
	BusinessStatus     string
	CreationTime       string
	ExpiredTime        string
	Status             string
	Ratio              int
	PublicIpAddresses  DescribeCommonBandwidthPackagesPublicIpAddresseList
}

type DescribeCommonBandwidthPackagesPublicIpAddresse struct {
	AllocationId string
	IpAddress    string
}
type DescribeCommonBandwidthPackagesArgs struct {
	ResourceOwnerId      int64
	BandwidthPackageId   string
	ResourceOwnerAccount string
	OwnerAccount         string
	Name                 string
	PageSize             int
	OwnerId              int64
	PageNumber           int
}
type DescribeCommonBandwidthPackagesResponse struct {
	RequestId               string
	TotalCount              int
	PageNumber              int
	PageSize                int
	CommonBandwidthPackages DescribeCommonBandwidthPackagesCommonBandwidthPackageList
}

type DescribeCommonBandwidthPackagesPublicIpAddresseList []DescribeCommonBandwidthPackagesPublicIpAddresse

func (list *DescribeCommonBandwidthPackagesPublicIpAddresseList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCommonBandwidthPackagesPublicIpAddresse)
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

type DescribeCommonBandwidthPackagesCommonBandwidthPackageList []DescribeCommonBandwidthPackagesCommonBandwidthPackage

func (list *DescribeCommonBandwidthPackagesCommonBandwidthPackageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCommonBandwidthPackagesCommonBandwidthPackage)
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

func (c *VpcClient) DescribeVpcAttribute(req *DescribeVpcAttributeArgs) (resp *DescribeVpcAttributeResponse, err error) {
	resp = &DescribeVpcAttributeResponse{}
	err = c.Request("DescribeVpcAttribute", req, resp)
	return
}

type DescribeVpcAttributeCloudResourceSetType struct {
	ResourceType  string
	ResourceCount int
}
type DescribeVpcAttributeArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	VpcId                string
	OwnerAccount         string
	IsDefault            core.Bool
	OwnerId              int64
}
type DescribeVpcAttributeResponse struct {
	RequestId          string
	VpcId              string
	RegionId           string
	Status             string
	VpcName            string
	CreationTime       string
	CidrBlock          string
	VRouterId          string
	Description        string
	IsDefault          core.Bool
	ClassicLinkEnabled core.Bool
	CloudResources     DescribeVpcAttributeCloudResourceSetTypeList
	VSwitchIds         DescribeVpcAttributeVSwitchIdList
	UserCidrs          DescribeVpcAttributeUserCidrList
}

type DescribeVpcAttributeCloudResourceSetTypeList []DescribeVpcAttributeCloudResourceSetType

func (list *DescribeVpcAttributeCloudResourceSetTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeVpcAttributeCloudResourceSetType)
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

type DescribeVpcAttributeVSwitchIdList []string

func (list *DescribeVpcAttributeVSwitchIdList) UnmarshalJSON(data []byte) error {
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

type DescribeVpcAttributeUserCidrList []string

func (list *DescribeVpcAttributeUserCidrList) UnmarshalJSON(data []byte) error {
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

func (c *VpcClient) ModifyCommonBandwidthPackageSpec(req *ModifyCommonBandwidthPackageSpecArgs) (resp *ModifyCommonBandwidthPackageSpecResponse, err error) {
	resp = &ModifyCommonBandwidthPackageSpecResponse{}
	err = c.Request("ModifyCommonBandwidthPackageSpec", req, resp)
	return
}

type ModifyCommonBandwidthPackageSpecArgs struct {
	ResourceOwnerId      int64
	BandwidthPackageId   string
	ResourceOwnerAccount string
	Bandwidth            string
	OwnerAccount         string
	OwnerId              int64
}
type ModifyCommonBandwidthPackageSpecResponse struct {
	RequestId string
}

func (c *VpcClient) DescribeVirtualBorderRouters(req *DescribeVirtualBorderRoutersArgs) (resp *DescribeVirtualBorderRoutersResponse, err error) {
	resp = &DescribeVirtualBorderRoutersResponse{}
	err = c.Request("DescribeVirtualBorderRouters", req, resp)
	return
}

type DescribeVirtualBorderRoutersFilter struct {
	Key    string
	Values DescribeVirtualBorderRoutersValueList
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
type DescribeVirtualBorderRoutersArgs struct {
	Filters              DescribeVirtualBorderRoutersFilterList
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	PageSize             int
	OwnerId              int64
	PageNumber           int
}
type DescribeVirtualBorderRoutersResponse struct {
	RequestId              string
	PageNumber             int
	PageSize               int
	TotalCount             int
	VirtualBorderRouterSet DescribeVirtualBorderRoutersVirtualBorderRouterTypeList
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

func (c *VpcClient) AddBandwidthPackageIps(req *AddBandwidthPackageIpsArgs) (resp *AddBandwidthPackageIpsResponse, err error) {
	resp = &AddBandwidthPackageIpsResponse{}
	err = c.Request("AddBandwidthPackageIps", req, resp)
	return
}

type AddBandwidthPackageIpsArgs struct {
	ResourceOwnerId      int64
	BandwidthPackageId   string
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	OwnerId              int64
	IpCount              string
}
type AddBandwidthPackageIpsResponse struct {
	RequestId string
}

func (c *VpcClient) CreateCustomerGateway(req *CreateCustomerGatewayArgs) (resp *CreateCustomerGatewayResponse, err error) {
	resp = &CreateCustomerGatewayResponse{}
	err = c.Request("CreateCustomerGateway", req, resp)
	return
}

type CreateCustomerGatewayArgs struct {
	IpAddress            string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	Name                 string
	Description          string
	OwnerId              int64
}
type CreateCustomerGatewayResponse struct {
	RequestId         string
	CustomerGatewayId string
	IpAddress         string
	Name              string
	Description       string
	CreateTime        int64
}

func (c *VpcClient) CreateBgpPeer(req *CreateBgpPeerArgs) (resp *CreateBgpPeerResponse, err error) {
	resp = &CreateBgpPeerResponse{}
	err = c.Request("CreateBgpPeer", req, resp)
	return
}

type CreateBgpPeerArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	BgpGroupId           string
	OwnerId              int64
	PeerIpAddress        string
}
type CreateBgpPeerResponse struct {
	RequestId string
	BgpPeerId string
}

func (c *VpcClient) DeleteSnatEntry(req *DeleteSnatEntryArgs) (resp *DeleteSnatEntryResponse, err error) {
	resp = &DeleteSnatEntryResponse{}
	err = c.Request("DeleteSnatEntry", req, resp)
	return
}

type DeleteSnatEntryArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	SnatTableId          string
	SnatEntryId          string
	OwnerId              int64
}
type DeleteSnatEntryResponse struct {
	RequestId string
}

func (c *VpcClient) DeleteBgpGroup(req *DeleteBgpGroupArgs) (resp *DeleteBgpGroupResponse, err error) {
	resp = &DeleteBgpGroupResponse{}
	err = c.Request("DeleteBgpGroup", req, resp)
	return
}

type DeleteBgpGroupArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	BgpGroupId           string
	OwnerId              int64
}
type DeleteBgpGroupResponse struct {
	RequestId string
}

func (c *VpcClient) DeleteBgpPeer(req *DeleteBgpPeerArgs) (resp *DeleteBgpPeerResponse, err error) {
	resp = &DeleteBgpPeerResponse{}
	err = c.Request("DeleteBgpPeer", req, resp)
	return
}

type DeleteBgpPeerArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	BgpPeerId            string
	OwnerId              int64
}
type DeleteBgpPeerResponse struct {
	RequestId string
}

func (c *VpcClient) UnassociateEipAddress(req *UnassociateEipAddressArgs) (resp *UnassociateEipAddressResponse, err error) {
	resp = &UnassociateEipAddressResponse{}
	err = c.Request("UnassociateEipAddress", req, resp)
	return
}

type UnassociateEipAddressArgs struct {
	ResourceOwnerId      int64
	InstanceId           string
	ResourceOwnerAccount string
	OwnerAccount         string
	InstanceType         string
	AllocationId         string
	OwnerId              int64
}
type UnassociateEipAddressResponse struct {
	RequestId string
}

func (c *VpcClient) ModifyNqa(req *ModifyNqaArgs) (resp *ModifyNqaResponse, err error) {
	resp = &ModifyNqaResponse{}
	err = c.Request("ModifyNqa", req, resp)
	return
}

type ModifyNqaArgs struct {
	DestinationIp        string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	NqaId                string
	OwnerId              int64
}
type ModifyNqaResponse struct {
	RequestId string
}

func (c *VpcClient) DescribeBandwidthPackages(req *DescribeBandwidthPackagesArgs) (resp *DescribeBandwidthPackagesResponse, err error) {
	resp = &DescribeBandwidthPackagesResponse{}
	err = c.Request("DescribeBandwidthPackages", req, resp)
	return
}

type DescribeBandwidthPackagesBandwidthPackage struct {
	BandwidthPackageId string
	RegionId           string
	Name               string
	Description        string
	ZoneId             string
	NatGatewayId       string
	Bandwidth          string
	InstanceChargeType string
	InternetChargeType string
	BusinessStatus     string
	IpCount            string
	CreationTime       string
	Status             string
	ISP                string
	PublicIpAddresses  DescribeBandwidthPackagesPublicIpAddresseList
}

type DescribeBandwidthPackagesPublicIpAddresse struct {
	AllocationId    string
	IpAddress       string
	UsingStatus     string
	ApAccessEnabled core.Bool
}
type DescribeBandwidthPackagesArgs struct {
	ResourceOwnerId      int64
	BandwidthPackageId   string
	ResourceOwnerAccount string
	OwnerAccount         string
	PageSize             int
	NatGatewayId         string
	OwnerId              int64
	PageNumber           int
}
type DescribeBandwidthPackagesResponse struct {
	RequestId         string
	TotalCount        int
	PageNumber        int
	PageSize          int
	BandwidthPackages DescribeBandwidthPackagesBandwidthPackageList
}

type DescribeBandwidthPackagesPublicIpAddresseList []DescribeBandwidthPackagesPublicIpAddresse

func (list *DescribeBandwidthPackagesPublicIpAddresseList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeBandwidthPackagesPublicIpAddresse)
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

type DescribeBandwidthPackagesBandwidthPackageList []DescribeBandwidthPackagesBandwidthPackage

func (list *DescribeBandwidthPackagesBandwidthPackageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeBandwidthPackagesBandwidthPackage)
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

func (c *VpcClient) DescribeServerRelatedGlobalAccelerationInstances(req *DescribeServerRelatedGlobalAccelerationInstancesArgs) (resp *DescribeServerRelatedGlobalAccelerationInstancesResponse, err error) {
	resp = &DescribeServerRelatedGlobalAccelerationInstancesResponse{}
	err = c.Request("DescribeServerRelatedGlobalAccelerationInstances", req, resp)
	return
}

type DescribeServerRelatedGlobalAccelerationInstancesGlobalAccelerationInstance struct {
	RegionId                     string
	GlobalAccelerationInstanceId string
	IpAddress                    string
	ServerIpAddress              string
}
type DescribeServerRelatedGlobalAccelerationInstancesArgs struct {
	ResourceOwnerId      int64
	ServerType           string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	ServerId             string
}
type DescribeServerRelatedGlobalAccelerationInstancesResponse struct {
	RequestId                   string
	GlobalAccelerationInstances DescribeServerRelatedGlobalAccelerationInstancesGlobalAccelerationInstanceList
}

type DescribeServerRelatedGlobalAccelerationInstancesGlobalAccelerationInstanceList []DescribeServerRelatedGlobalAccelerationInstancesGlobalAccelerationInstance

func (list *DescribeServerRelatedGlobalAccelerationInstancesGlobalAccelerationInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeServerRelatedGlobalAccelerationInstancesGlobalAccelerationInstance)
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

func (c *VpcClient) DeleteGlobalAccelerationInstance(req *DeleteGlobalAccelerationInstanceArgs) (resp *DeleteGlobalAccelerationInstanceResponse, err error) {
	resp = &DeleteGlobalAccelerationInstanceResponse{}
	err = c.Request("DeleteGlobalAccelerationInstance", req, resp)
	return
}

type DeleteGlobalAccelerationInstanceArgs struct {
	ResourceOwnerId              int64
	ResourceOwnerAccount         string
	OwnerAccount                 string
	OwnerId                      int64
	GlobalAccelerationInstanceId string
}
type DeleteGlobalAccelerationInstanceResponse struct {
	RequestId string
}

func (c *VpcClient) DeleteBandwidthPackage(req *DeleteBandwidthPackageArgs) (resp *DeleteBandwidthPackageResponse, err error) {
	resp = &DeleteBandwidthPackageResponse{}
	err = c.Request("DeleteBandwidthPackage", req, resp)
	return
}

type DeleteBandwidthPackageArgs struct {
	ResourceOwnerId      int64
	BandwidthPackageId   string
	ResourceOwnerAccount string
	OwnerAccount         string
	Force                core.Bool
	OwnerId              int64
}
type DeleteBandwidthPackageResponse struct {
	RequestId string
}

func (c *VpcClient) ModifyRouteTableAttributes(req *ModifyRouteTableAttributesArgs) (resp *ModifyRouteTableAttributesResponse, err error) {
	resp = &ModifyRouteTableAttributesResponse{}
	err = c.Request("ModifyRouteTableAttributes", req, resp)
	return
}

type ModifyRouteTableAttributesArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	Bandwidth            string
	OwnerAccount         string
	Description          string
	OwnerId              int64
	KbpsBandwidth        string
	RouteTableName       string
	ResourceUid          int64
	ResourceBid          string
	RouteTableId         string
}
type ModifyRouteTableAttributesResponse struct {
	RequestId string
	Code      string
	Message   string
	Success   core.Bool
}

func (c *VpcClient) AllocateEipAddress(req *AllocateEipAddressArgs) (resp *AllocateEipAddressResponse, err error) {
	resp = &AllocateEipAddressResponse{}
	err = c.Request("AllocateEipAddress", req, resp)
	return
}

type AllocateEipAddressArgs struct {
	ResourceOwnerId      int64
	Period               int
	AutoPay              core.Bool
	ResourceOwnerAccount string
	Bandwidth            string
	ClientToken          string
	ISP                  string
	OwnerAccount         string
	OwnerId              int64
	InternetChargeType   string
	Netmode              string
	PricingCycle         string
	InstanceChargeType   string
}
type AllocateEipAddressResponse struct {
	RequestId    string
	AllocationId string
	EipAddress   string
	OrderId      int64
}

func (c *VpcClient) DescribeVpnConnection(req *DescribeVpnConnectionArgs) (resp *DescribeVpnConnectionResponse, err error) {
	resp = &DescribeVpnConnectionResponse{}
	err = c.Request("DescribeVpnConnection", req, resp)
	return
}

type DescribeVpnConnectionIkeConfig struct {
	Psk         string
	IkeVersion  string
	IkeMode     string
	IkeEncAlg   string
	IkeAuthAlg  string
	IkePfs      string
	IkeLifetime int64
	LocalId     string
	RemoteId    string
}

type DescribeVpnConnectionIpsecConfig struct {
	IpsecEncAlg   string
	IpsecAuthAlg  string
	IpsecPfs      string
	IpsecLifetime int64
}
type DescribeVpnConnectionArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	VpnConnectionId      string
	OwnerAccount         string
	OwnerId              int64
}
type DescribeVpnConnectionResponse struct {
	RequestId         string
	VpnConnectionId   string
	CustomerGatewayId string
	VpnGatewayId      string
	Name              string
	LocalSubnet       string
	RemoteSubnet      string
	CreateTime        int64
	EffectImmediately core.Bool
	Status            string
	IkeConfig         DescribeVpnConnectionIkeConfig
	IpsecConfig       DescribeVpnConnectionIpsecConfig
}

func (c *VpcClient) DeleteVpnGateway(req *DeleteVpnGatewayArgs) (resp *DeleteVpnGatewayResponse, err error) {
	resp = &DeleteVpnGatewayResponse{}
	err = c.Request("DeleteVpnGateway", req, resp)
	return
}

type DeleteVpnGatewayArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	VpnGatewayId         string
	OwnerId              int64
}
type DeleteVpnGatewayResponse struct {
	RequestId string
}

func (c *VpcClient) CreateHaVip(req *CreateHaVipArgs) (resp *CreateHaVipResponse, err error) {
	resp = &CreateHaVipResponse{}
	err = c.Request("CreateHaVip", req, resp)
	return
}

type CreateHaVipArgs struct {
	VSwitchId            string
	IpAddress            string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	Description          string
	OwnerId              int64
}
type CreateHaVipResponse struct {
	RequestId string
	HaVipId   string
}

func (c *VpcClient) ModifyBgpGroupAttribute(req *ModifyBgpGroupAttributeArgs) (resp *ModifyBgpGroupAttributeResponse, err error) {
	resp = &ModifyBgpGroupAttributeResponse{}
	err = c.Request("ModifyBgpGroupAttribute", req, resp)
	return
}

type ModifyBgpGroupAttributeArgs struct {
	AuthKey              string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	BgpGroupId           string
	Description          string
	OwnerId              int64
	PeerAsn              int64
	IsFakeAsn            core.Bool
	Name                 string
}
type ModifyBgpGroupAttributeResponse struct {
	RequestId string
}

func (c *VpcClient) ModifyVirtualBorderRouterAttribute(req *ModifyVirtualBorderRouterAttributeArgs) (resp *ModifyVirtualBorderRouterAttributeResponse, err error) {
	resp = &ModifyVirtualBorderRouterAttributeResponse{}
	err = c.Request("ModifyVirtualBorderRouterAttribute", req, resp)
	return
}

type ModifyVirtualBorderRouterAttributeArgs struct {
	ResourceOwnerId               int64
	CircuitCode                   string
	AssociatedPhysicalConnections string
	VlanId                        int
	ResourceOwnerAccount          string
	ClientToken                   string
	OwnerAccount                  string
	Description                   string
	VbrId                         string
	OwnerId                       int64
	PeerGatewayIp                 string
	PeeringSubnetMask             string
	Name                          string
	LocalGatewayIp                string
	UserCidr                      string
}
type ModifyVirtualBorderRouterAttributeResponse struct {
	RequestId string
}

func (c *VpcClient) DescribeSnatTableEntries(req *DescribeSnatTableEntriesArgs) (resp *DescribeSnatTableEntriesResponse, err error) {
	resp = &DescribeSnatTableEntriesResponse{}
	err = c.Request("DescribeSnatTableEntries", req, resp)
	return
}

type DescribeSnatTableEntriesSnatTableEntry struct {
	SnatTableId     string
	SnatEntryId     string
	SourceVSwitchId string
	SourceCIDR      string
	SnatIp          string
	Status          string
}
type DescribeSnatTableEntriesArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	PageSize             int
	SnatTableId          string
	SnatEntryId          string
	OwnerId              int64
	PageNumber           int
}
type DescribeSnatTableEntriesResponse struct {
	RequestId        string
	TotalCount       int
	PageNumber       int
	PageSize         int
	SnatTableEntries DescribeSnatTableEntriesSnatTableEntryList
}

type DescribeSnatTableEntriesSnatTableEntryList []DescribeSnatTableEntriesSnatTableEntry

func (list *DescribeSnatTableEntriesSnatTableEntryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSnatTableEntriesSnatTableEntry)
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

func (c *VpcClient) CreatePhysicalConnection(req *CreatePhysicalConnectionArgs) (resp *CreatePhysicalConnectionResponse, err error) {
	resp = &CreatePhysicalConnectionResponse{}
	err = c.Request("CreatePhysicalConnection", req, resp)
	return
}

type CreatePhysicalConnectionArgs struct {
	AccessPointId                 string
	RedundantPhysicalConnectionId string
	PeerLocation                  string
	ResourceOwnerId               int64
	PortType                      string
	CircuitCode                   string
	Bandwidth                     int
	ClientToken                   string
	ResourceOwnerAccount          string
	OwnerAccount                  string
	Description                   string
	Type                          string
	OwnerId                       int64
	LineOperator                  string
	Name                          string
	UserCidr                      string
}
type CreatePhysicalConnectionResponse struct {
	RequestId            string
	PhysicalConnectionId string
}

func (c *VpcClient) DescribeRouterInterfaces(req *DescribeRouterInterfacesArgs) (resp *DescribeRouterInterfacesResponse, err error) {
	resp = &DescribeRouterInterfacesResponse{}
	err = c.Request("DescribeRouterInterfaces", req, resp)
	return
}

type DescribeRouterInterfacesFilter struct {
	Key    string
	Values DescribeRouterInterfacesValueList
}

type DescribeRouterInterfacesRouterInterfaceType struct {
	RouterInterfaceId               string
	OppositeRegionId                string
	Role                            string
	Spec                            string
	Name                            string
	Description                     string
	RouterId                        string
	RouterType                      string
	CreationTime                    string
	EndTime                         string
	ChargeType                      string
	Status                          string
	BusinessStatus                  string
	ConnectedTime                   string
	OppositeInterfaceId             string
	OppositeInterfaceSpec           string
	OppositeInterfaceStatus         string
	OppositeInterfaceBusinessStatus string
	OppositeRouterId                string
	OppositeRouterType              string
	OppositeInterfaceOwnerId        string
	AccessPointId                   string
	OppositeAccessPointId           string
	HealthCheckSourceIp             string
	HealthCheckTargetIp             string
	OppositeVpcInstanceId           string
	VpcInstanceId                   string
}
type DescribeRouterInterfacesArgs struct {
	Filters              DescribeRouterInterfacesFilterList
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	PageSize             int
	OwnerId              int64
	PageNumber           int
}
type DescribeRouterInterfacesResponse struct {
	RequestId          string
	PageNumber         int
	PageSize           int
	TotalCount         int
	RouterInterfaceSet DescribeRouterInterfacesRouterInterfaceTypeList
}

type DescribeRouterInterfacesValueList []string

func (list *DescribeRouterInterfacesValueList) UnmarshalJSON(data []byte) error {
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

type DescribeRouterInterfacesFilterList []DescribeRouterInterfacesFilter

func (list *DescribeRouterInterfacesFilterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRouterInterfacesFilter)
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

type DescribeRouterInterfacesRouterInterfaceTypeList []DescribeRouterInterfacesRouterInterfaceType

func (list *DescribeRouterInterfacesRouterInterfaceTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRouterInterfacesRouterInterfaceType)
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

func (c *VpcClient) CreateNatGateway(req *CreateNatGatewayArgs) (resp *CreateNatGatewayResponse, err error) {
	resp = &CreateNatGatewayResponse{}
	err = c.Request("CreateNatGateway", req, resp)
	return
}

type CreateNatGatewayBandwidthPackage struct {
	IpCount            int
	Bandwidth          int
	Zone               string
	ISP                string
	InternetChargeType string
}
type CreateNatGatewayArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	VpcId                string
	Name                 string
	Description          string
	OwnerId              int64
	BandwidthPackages    CreateNatGatewayBandwidthPackageList
	Spec                 string
}
type CreateNatGatewayResponse struct {
	RequestId           string
	NatGatewayId        string
	ForwardTableIds     CreateNatGatewayForwardTableIdList
	SnatTableIds        CreateNatGatewaySnatTableIdList
	BandwidthPackageIds CreateNatGatewayBandwidthPackageIdList
}

type CreateNatGatewayBandwidthPackageList []CreateNatGatewayBandwidthPackage

func (list *CreateNatGatewayBandwidthPackageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateNatGatewayBandwidthPackage)
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

type CreateNatGatewayForwardTableIdList []string

func (list *CreateNatGatewayForwardTableIdList) UnmarshalJSON(data []byte) error {
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

type CreateNatGatewaySnatTableIdList []string

func (list *CreateNatGatewaySnatTableIdList) UnmarshalJSON(data []byte) error {
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

type CreateNatGatewayBandwidthPackageIdList []string

func (list *CreateNatGatewayBandwidthPackageIdList) UnmarshalJSON(data []byte) error {
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

func (c *VpcClient) DescribeVpcs(req *DescribeVpcsArgs) (resp *DescribeVpcsResponse, err error) {
	resp = &DescribeVpcsResponse{}
	err = c.Request("DescribeVpcs", req, resp)
	return
}

type DescribeVpcsVpc struct {
	VpcId          string
	RegionId       string
	Status         string
	VpcName        string
	CreationTime   string
	CidrBlock      string
	VRouterId      string
	Description    string
	IsDefault      core.Bool
	VSwitchIds     DescribeVpcsVSwitchIdList
	UserCidrs      DescribeVpcsUserCidrList
	NatGatewayIds  DescribeVpcsNatGatewayIdList
	RouterTableIds DescribeVpcsRouterTableIdList
}
type DescribeVpcsArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	VpcId                string
	OwnerAccount         string
	Name                 string
	PageSize             int
	IsDefault            core.Bool
	OwnerId              int64
	PageNumber           int
}
type DescribeVpcsResponse struct {
	RequestId  string
	TotalCount int
	PageNumber int
	PageSize   int
	Vpcs       DescribeVpcsVpcList
}

type DescribeVpcsVSwitchIdList []string

func (list *DescribeVpcsVSwitchIdList) UnmarshalJSON(data []byte) error {
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

type DescribeVpcsUserCidrList []string

func (list *DescribeVpcsUserCidrList) UnmarshalJSON(data []byte) error {
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

type DescribeVpcsNatGatewayIdList []string

func (list *DescribeVpcsNatGatewayIdList) UnmarshalJSON(data []byte) error {
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

type DescribeVpcsRouterTableIdList []string

func (list *DescribeVpcsRouterTableIdList) UnmarshalJSON(data []byte) error {
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

type DescribeVpcsVpcList []DescribeVpcsVpc

func (list *DescribeVpcsVpcList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeVpcsVpc)
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

func (c *VpcClient) AddBgpNetwork(req *AddBgpNetworkArgs) (resp *AddBgpNetworkResponse, err error) {
	resp = &AddBgpNetworkResponse{}
	err = c.Request("AddBgpNetwork", req, resp)
	return
}

type AddBgpNetworkArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ClientToken          string
	RouterId             string
	VpcId                string
	OwnerAccount         string
	OwnerId              int64
	DstCidrBlock         string
}
type AddBgpNetworkResponse struct {
	RequestId string
}

func (c *VpcClient) DeleteNqa(req *DeleteNqaArgs) (resp *DeleteNqaResponse, err error) {
	resp = &DeleteNqaResponse{}
	err = c.Request("DeleteNqa", req, resp)
	return
}

type DeleteNqaArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	NqaId                string
	OwnerId              int64
}
type DeleteNqaResponse struct {
	RequestId string
}

func (c *VpcClient) CreateCommonBandwidthPackage(req *CreateCommonBandwidthPackageArgs) (resp *CreateCommonBandwidthPackageResponse, err error) {
	resp = &CreateCommonBandwidthPackageResponse{}
	err = c.Request("CreateCommonBandwidthPackage", req, resp)
	return
}

type CreateCommonBandwidthPackageArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ClientToken          string
	Bandwidth            int
	InternetChargeType   string
	OwnerAccount         string
	Name                 string
	Description          string
	OwnerId              int64
	Ratio                int
}
type CreateCommonBandwidthPackageResponse struct {
	RequestId          string
	BandwidthPackageId string
}

func (c *VpcClient) ModifyRouterInterfaceSpec(req *ModifyRouterInterfaceSpecArgs) (resp *ModifyRouterInterfaceSpecResponse, err error) {
	resp = &ModifyRouterInterfaceSpecResponse{}
	err = c.Request("ModifyRouterInterfaceSpec", req, resp)
	return
}

type ModifyRouterInterfaceSpecArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	UserCidr             string
	RouterInterfaceId    string
	OwnerId              int64
	Spec                 string
}
type ModifyRouterInterfaceSpecResponse struct {
	RequestId string
	Spec      string
}

func (c *VpcClient) DeleteCommonBandwidthPackage(req *DeleteCommonBandwidthPackageArgs) (resp *DeleteCommonBandwidthPackageResponse, err error) {
	resp = &DeleteCommonBandwidthPackageResponse{}
	err = c.Request("DeleteCommonBandwidthPackage", req, resp)
	return
}

type DeleteCommonBandwidthPackageArgs struct {
	ResourceOwnerId      int64
	BandwidthPackageId   string
	ResourceOwnerAccount string
	OwnerAccount         string
	Force                string
	OwnerId              int64
}
type DeleteCommonBandwidthPackageResponse struct {
	RequestId string
}

func (c *VpcClient) ModifyBandwidthPackageAttribute(req *ModifyBandwidthPackageAttributeArgs) (resp *ModifyBandwidthPackageAttributeResponse, err error) {
	resp = &ModifyBandwidthPackageAttributeResponse{}
	err = c.Request("ModifyBandwidthPackageAttribute", req, resp)
	return
}

type ModifyBandwidthPackageAttributeArgs struct {
	ResourceOwnerId      int64
	BandwidthPackageId   string
	ResourceOwnerAccount string
	OwnerAccount         string
	Name                 string
	Description          string
	OwnerId              int64
}
type ModifyBandwidthPackageAttributeResponse struct {
	RequestId string
}

func (c *VpcClient) DeleteBgpNetwork(req *DeleteBgpNetworkArgs) (resp *DeleteBgpNetworkResponse, err error) {
	resp = &DeleteBgpNetworkResponse{}
	err = c.Request("DeleteBgpNetwork", req, resp)
	return
}

type DeleteBgpNetworkArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ClientToken          string
	RouterId             string
	OwnerAccount         string
	OwnerId              int64
	DstCidrBlock         string
}
type DeleteBgpNetworkResponse struct {
	RequestId string
}

func (c *VpcClient) ModifyNatGatewayAttribute(req *ModifyNatGatewayAttributeArgs) (resp *ModifyNatGatewayAttributeResponse, err error) {
	resp = &ModifyNatGatewayAttributeResponse{}
	err = c.Request("ModifyNatGatewayAttribute", req, resp)
	return
}

type ModifyNatGatewayAttributeArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	Name                 string
	Description          string
	NatGatewayId         string
	OwnerId              int64
}
type ModifyNatGatewayAttributeResponse struct {
	RequestId string
}

func (c *VpcClient) DescribeHaVips(req *DescribeHaVipsArgs) (resp *DescribeHaVipsResponse, err error) {
	resp = &DescribeHaVipsResponse{}
	err = c.Request("DescribeHaVips", req, resp)
	return
}

type DescribeHaVipsFilter struct {
	Key    string
	Values DescribeHaVipsValueList
}

type DescribeHaVipsHaVip struct {
	HaVipId                string
	RegionId               string
	VpcId                  string
	VSwitchId              string
	IpAddress              string
	Status                 string
	MasterInstanceId       string
	Description            string
	CreateTime             string
	AssociatedInstances    DescribeHaVipsAssociatedInstanceList
	AssociatedEipAddresses DescribeHaVipsAssociatedEipAddressList
}
type DescribeHaVipsArgs struct {
	Filters              DescribeHaVipsFilterList
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	PageSize             int
	OwnerId              int64
	PageNumber           int
}
type DescribeHaVipsResponse struct {
	RequestId  string
	TotalCount int
	PageNumber int
	PageSize   int
	HaVips     DescribeHaVipsHaVipList
}

type DescribeHaVipsValueList []string

func (list *DescribeHaVipsValueList) UnmarshalJSON(data []byte) error {
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

type DescribeHaVipsAssociatedInstanceList []string

func (list *DescribeHaVipsAssociatedInstanceList) UnmarshalJSON(data []byte) error {
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

type DescribeHaVipsAssociatedEipAddressList []string

func (list *DescribeHaVipsAssociatedEipAddressList) UnmarshalJSON(data []byte) error {
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

type DescribeHaVipsFilterList []DescribeHaVipsFilter

func (list *DescribeHaVipsFilterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeHaVipsFilter)
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

type DescribeHaVipsHaVipList []DescribeHaVipsHaVip

func (list *DescribeHaVipsHaVipList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeHaVipsHaVip)
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

func (c *VpcClient) UnassociatePhysicalConnectionFromVirtualBorderRouter(req *UnassociatePhysicalConnectionFromVirtualBorderRouterArgs) (resp *UnassociatePhysicalConnectionFromVirtualBorderRouterResponse, err error) {
	resp = &UnassociatePhysicalConnectionFromVirtualBorderRouterResponse{}
	err = c.Request("UnassociatePhysicalConnectionFromVirtualBorderRouter", req, resp)
	return
}

type UnassociatePhysicalConnectionFromVirtualBorderRouterArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ClientToken          string
	PhysicalConnectionId string
	OwnerAccount         string
	UserCidr             string
	VbrId                string
	OwnerId              int64
}
type UnassociatePhysicalConnectionFromVirtualBorderRouterResponse struct {
	RequestId string
}

func (c *VpcClient) DeleteForwardEntry(req *DeleteForwardEntryArgs) (resp *DeleteForwardEntryResponse, err error) {
	resp = &DeleteForwardEntryResponse{}
	err = c.Request("DeleteForwardEntry", req, resp)
	return
}

type DeleteForwardEntryArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ForwardEntryId       string
	OwnerAccount         string
	ForwardTableId       string
	OwnerId              int64
}
type DeleteForwardEntryResponse struct {
	RequestId string
}

func (c *VpcClient) CreateVpnConnection(req *CreateVpnConnectionArgs) (resp *CreateVpnConnectionResponse, err error) {
	resp = &CreateVpnConnectionResponse{}
	err = c.Request("CreateVpnConnection", req, resp)
	return
}

type CreateVpnConnectionArgs struct {
	IkeConfig            string
	ResourceOwnerId      int64
	RemoteSubnet         string
	EffectImmediately    core.Bool
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	IpsecConfig          string
	VpnGatewayId         string
	OwnerId              int64
	CustomerGatewayId    string
	LocalSubnet          string
	Name                 string
}
type CreateVpnConnectionResponse struct {
	RequestId       string
	VpnConnectionId string
	Name            string
	CreateTime      int64
}

func (c *VpcClient) DescribeNatGateways(req *DescribeNatGatewaysArgs) (resp *DescribeNatGatewaysResponse, err error) {
	resp = &DescribeNatGatewaysResponse{}
	err = c.Request("DescribeNatGateways", req, resp)
	return
}

type DescribeNatGatewaysNatGateway struct {
	NatGatewayId        string
	RegionId            string
	Name                string
	Description         string
	VpcId               string
	Spec                string
	InstanceChargeType  string
	BusinessStatus      string
	CreationTime        string
	Status              string
	IpLists             DescribeNatGatewaysIpListList
	ForwardTableIds     DescribeNatGatewaysForwardTableIdList
	SnatTableIds        DescribeNatGatewaysSnatTableIdList
	BandwidthPackageIds DescribeNatGatewaysBandwidthPackageIdList
}

type DescribeNatGatewaysIpList struct {
	AllocationId    string
	IpAddress       string
	UsingStatus     string
	ApAccessEnabled core.Bool
}
type DescribeNatGatewaysArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	VpcId                string
	PageSize             int
	NatGatewayId         string
	OwnerId              int64
	PageNumber           int
}
type DescribeNatGatewaysResponse struct {
	RequestId   string
	TotalCount  int
	PageNumber  int
	PageSize    int
	NatGateways DescribeNatGatewaysNatGatewayList
}

type DescribeNatGatewaysIpListList []DescribeNatGatewaysIpList

func (list *DescribeNatGatewaysIpListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeNatGatewaysIpList)
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

type DescribeNatGatewaysForwardTableIdList []string

func (list *DescribeNatGatewaysForwardTableIdList) UnmarshalJSON(data []byte) error {
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

type DescribeNatGatewaysSnatTableIdList []string

func (list *DescribeNatGatewaysSnatTableIdList) UnmarshalJSON(data []byte) error {
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

type DescribeNatGatewaysBandwidthPackageIdList []string

func (list *DescribeNatGatewaysBandwidthPackageIdList) UnmarshalJSON(data []byte) error {
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

type DescribeNatGatewaysNatGatewayList []DescribeNatGatewaysNatGateway

func (list *DescribeNatGatewaysNatGatewayList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeNatGatewaysNatGateway)
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

func (c *VpcClient) DescribeBandwidthPackagePublicIpMonitorData(req *DescribeBandwidthPackagePublicIpMonitorDataArgs) (resp *DescribeBandwidthPackagePublicIpMonitorDataResponse, err error) {
	resp = &DescribeBandwidthPackagePublicIpMonitorDataResponse{}
	err = c.Request("DescribeBandwidthPackagePublicIpMonitorData", req, resp)
	return
}

type DescribeBandwidthPackagePublicIpMonitorDataMonitorData struct {
	RX                   int64
	TX                   int64
	ReceivedBandwidth    int64
	TransportedBandwidth int64
	Flow                 int64
	Bandwidth            int64
	Packets              int64
	TimeStamp            string
}
type DescribeBandwidthPackagePublicIpMonitorDataArgs struct {
	ResourceOwnerId      int64
	Period               int
	ResourceOwnerAccount string
	OwnerAccount         string
	EndTime              string
	AllocationId         string
	StartTime            string
	OwnerId              int64
}
type DescribeBandwidthPackagePublicIpMonitorDataResponse struct {
	RequestId    string
	MonitorDatas DescribeBandwidthPackagePublicIpMonitorDataMonitorDataList
}

type DescribeBandwidthPackagePublicIpMonitorDataMonitorDataList []DescribeBandwidthPackagePublicIpMonitorDataMonitorData

func (list *DescribeBandwidthPackagePublicIpMonitorDataMonitorDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeBandwidthPackagePublicIpMonitorDataMonitorData)
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

func (c *VpcClient) TerminateVirtualBorderRouter(req *TerminateVirtualBorderRouterArgs) (resp *TerminateVirtualBorderRouterResponse, err error) {
	resp = &TerminateVirtualBorderRouterResponse{}
	err = c.Request("TerminateVirtualBorderRouter", req, resp)
	return
}

type TerminateVirtualBorderRouterArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	UserCidr             string
	VbrId                string
	OwnerId              int64
}
type TerminateVirtualBorderRouterResponse struct {
	RequestId string
}

func (c *VpcClient) ModifyGlobalAccelerationInstanceAttributes(req *ModifyGlobalAccelerationInstanceAttributesArgs) (resp *ModifyGlobalAccelerationInstanceAttributesResponse, err error) {
	resp = &ModifyGlobalAccelerationInstanceAttributesResponse{}
	err = c.Request("ModifyGlobalAccelerationInstanceAttributes", req, resp)
	return
}

type ModifyGlobalAccelerationInstanceAttributesArgs struct {
	ResourceOwnerId              int64
	ResourceOwnerAccount         string
	OwnerAccount                 string
	Name                         string
	Description                  string
	OwnerId                      int64
	GlobalAccelerationInstanceId string
}
type ModifyGlobalAccelerationInstanceAttributesResponse struct {
	RequestId string
}

func (c *VpcClient) ModifyEipAddressAttribute(req *ModifyEipAddressAttributeArgs) (resp *ModifyEipAddressAttributeResponse, err error) {
	resp = &ModifyEipAddressAttributeResponse{}
	err = c.Request("ModifyEipAddressAttribute", req, resp)
	return
}

type ModifyEipAddressAttributeArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	Bandwidth            string
	OwnerAccount         string
	Name                 string
	Description          string
	AllocationId         string
	OwnerId              int64
}
type ModifyEipAddressAttributeResponse struct {
	RequestId string
}

func (c *VpcClient) DeactivateRouterInterface(req *DeactivateRouterInterfaceArgs) (resp *DeactivateRouterInterfaceResponse, err error) {
	resp = &DeactivateRouterInterfaceResponse{}
	err = c.Request("DeactivateRouterInterface", req, resp)
	return
}

type DeactivateRouterInterfaceArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerId              int64
	RouterInterfaceId    string
}
type DeactivateRouterInterfaceResponse struct {
	RequestId string
}

func (c *VpcClient) DescribeZones(req *DescribeZonesArgs) (resp *DescribeZonesResponse, err error) {
	resp = &DescribeZonesResponse{}
	err = c.Request("DescribeZones", req, resp)
	return
}

type DescribeZonesZone struct {
	ZoneId    string
	LocalName string
}
type DescribeZonesArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
}
type DescribeZonesResponse struct {
	RequestId string
	Zones     DescribeZonesZoneList
}

type DescribeZonesZoneList []DescribeZonesZone

func (list *DescribeZonesZoneList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeZonesZone)
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

func (c *VpcClient) DescribeVirtualBorderRoutersForPhysicalConnection(req *DescribeVirtualBorderRoutersForPhysicalConnectionArgs) (resp *DescribeVirtualBorderRoutersForPhysicalConnectionResponse, err error) {
	resp = &DescribeVirtualBorderRoutersForPhysicalConnectionResponse{}
	err = c.Request("DescribeVirtualBorderRoutersForPhysicalConnection", req, resp)
	return
}

type DescribeVirtualBorderRoutersForPhysicalConnectionFilter struct {
	Key    string
	Values DescribeVirtualBorderRoutersForPhysicalConnectionValueList
}

type DescribeVirtualBorderRoutersForPhysicalConnectionVirtualBorderRouterForPhysicalConnectionType struct {
	VbrId           string
	VbrOwnerUid     int64
	CreationTime    string
	ActivationTime  string
	TerminationTime string
	RecoveryTime    string
	VlanId          int
	CircuitCode     string
}
type DescribeVirtualBorderRoutersForPhysicalConnectionArgs struct {
	Filters              DescribeVirtualBorderRoutersForPhysicalConnectionFilterList
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	PhysicalConnectionId string
	PageSize             int
	OwnerId              int64
	PageNumber           int
}
type DescribeVirtualBorderRoutersForPhysicalConnectionResponse struct {
	RequestId                                   string
	PageNumber                                  int
	PageSize                                    int
	TotalCount                                  int
	VirtualBorderRouterForPhysicalConnectionSet DescribeVirtualBorderRoutersForPhysicalConnectionVirtualBorderRouterForPhysicalConnectionTypeList
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

func (c *VpcClient) DownloadVpnConnectionConfig(req *DownloadVpnConnectionConfigArgs) (resp *DownloadVpnConnectionConfigResponse, err error) {
	resp = &DownloadVpnConnectionConfigResponse{}
	err = c.Request("DownloadVpnConnectionConfig", req, resp)
	return
}

type DownloadVpnConnectionConfigVpnConnectionConfig struct {
	LocalSubnet  string
	RemoteSubnet string
	Local        string
	Remote       string
	IkeConfig    DownloadVpnConnectionConfigIkeConfig
	IpsecConfig  DownloadVpnConnectionConfigIpsecConfig
}

type DownloadVpnConnectionConfigIkeConfig struct {
	Psk         string
	IkeVersion  string
	IkeMode     string
	IkeEncAlg   string
	IkeAuthAlg  string
	IkePfs      string
	IkeLifetime int64
	LocalId     string
	RemoteId    string
}

type DownloadVpnConnectionConfigIpsecConfig struct {
	IpsecEncAlg   string
	IpsecAuthAlg  string
	IpsecPfs      string
	IpsecLifetime int64
}
type DownloadVpnConnectionConfigArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	VpnConnectionId      string
	OwnerAccount         string
	OwnerId              int64
}
type DownloadVpnConnectionConfigResponse struct {
	RequestId           string
	VpnConnectionConfig DownloadVpnConnectionConfigVpnConnectionConfig
}

func (c *VpcClient) DescribeEipMonitorData(req *DescribeEipMonitorDataArgs) (resp *DescribeEipMonitorDataResponse, err error) {
	resp = &DescribeEipMonitorDataResponse{}
	err = c.Request("DescribeEipMonitorData", req, resp)
	return
}

type DescribeEipMonitorDataEipMonitorData struct {
	EipRX        int
	EipTX        int
	EipFlow      int
	EipBandwidth int
	EipPackets   int
	TimeStamp    string
}
type DescribeEipMonitorDataArgs struct {
	ResourceOwnerId      int64
	Period               int
	ResourceOwnerAccount string
	OwnerAccount         string
	EndTime              string
	AllocationId         string
	StartTime            string
	OwnerId              int64
}
type DescribeEipMonitorDataResponse struct {
	RequestId       string
	EipMonitorDatas DescribeEipMonitorDataEipMonitorDataList
}

type DescribeEipMonitorDataEipMonitorDataList []DescribeEipMonitorDataEipMonitorData

func (list *DescribeEipMonitorDataEipMonitorDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeEipMonitorDataEipMonitorData)
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

func (c *VpcClient) ModifyVpcAttribute(req *ModifyVpcAttributeArgs) (resp *ModifyVpcAttributeResponse, err error) {
	resp = &ModifyVpcAttributeResponse{}
	err = c.Request("ModifyVpcAttribute", req, resp)
	return
}

type ModifyVpcAttributeArgs struct {
	VpcName              string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	VpcId                string
	OwnerAccount         string
	Description          string
	UserCidr             string
	OwnerId              int64
}
type ModifyVpcAttributeResponse struct {
	RequestId string
}

func (c *VpcClient) DescribeVSwitchAttributes(req *DescribeVSwitchAttributesArgs) (resp *DescribeVSwitchAttributesResponse, err error) {
	resp = &DescribeVSwitchAttributesResponse{}
	err = c.Request("DescribeVSwitchAttributes", req, resp)
	return
}

type DescribeVSwitchAttributesCloudResourceSetType struct {
	ResourceType  string
	ResourceCount int
}
type DescribeVSwitchAttributesArgs struct {
	VSwitchId            string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
}
type DescribeVSwitchAttributesResponse struct {
	RequestId               string
	VSwitchId               string
	VpcId                   string
	Status                  string
	CidrBlock               string
	ZoneId                  string
	AvailableIpAddressCount int64
	Description             string
	VSwitchName             string
	CreationTime            string
	IsDefault               core.Bool
	CloudResources          DescribeVSwitchAttributesCloudResourceSetTypeList
}

type DescribeVSwitchAttributesCloudResourceSetTypeList []DescribeVSwitchAttributesCloudResourceSetType

func (list *DescribeVSwitchAttributesCloudResourceSetTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeVSwitchAttributesCloudResourceSetType)
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

func (c *VpcClient) DescribeAccessPoints(req *DescribeAccessPointsArgs) (resp *DescribeAccessPointsResponse, err error) {
	resp = &DescribeAccessPointsResponse{}
	err = c.Request("DescribeAccessPoints", req, resp)
	return
}

type DescribeAccessPointsFilter struct {
	Key    string
	Values DescribeAccessPointsValueList
}

type DescribeAccessPointsAccessPointType struct {
	AccessPointId    string
	Status           string
	Type             string
	AttachedRegionNo string
	Location         string
	HostOperator     string
	Name             string
	Description      string
}
type DescribeAccessPointsArgs struct {
	Filters              DescribeAccessPointsFilterList
	ResourceOwnerId      int64
	HostOperator         string
	ResourceOwnerAccount string
	Name                 string
	PageSize             int
	OwnerId              int64
	Type                 string
	PageNumber           int
}
type DescribeAccessPointsResponse struct {
	RequestId      string
	PageNumber     int
	PageSize       int
	TotalCount     int
	AccessPointSet DescribeAccessPointsAccessPointTypeList
}

type DescribeAccessPointsValueList []string

func (list *DescribeAccessPointsValueList) UnmarshalJSON(data []byte) error {
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

type DescribeAccessPointsFilterList []DescribeAccessPointsFilter

func (list *DescribeAccessPointsFilterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAccessPointsFilter)
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

type DescribeAccessPointsAccessPointTypeList []DescribeAccessPointsAccessPointType

func (list *DescribeAccessPointsAccessPointTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAccessPointsAccessPointType)
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

func (c *VpcClient) DescribeBgpPeers(req *DescribeBgpPeersArgs) (resp *DescribeBgpPeersResponse, err error) {
	resp = &DescribeBgpPeersResponse{}
	err = c.Request("DescribeBgpPeers", req, resp)
	return
}

type DescribeBgpPeersBgpPeer struct {
	Name          string
	Description   string
	BgpPeerId     string
	BgpGroupId    string
	PeerIpAddress string
	PeerAsn       string
	AuthKey       string
	RouterId      string
	BgpStatus     string
	Status        string
	Keepalive     string
	LocalAsn      string
	Hold          string
	IsFake        string
	RouteLimit    string
	RegionId      string
}
type DescribeBgpPeersArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	RouterId             string
	OwnerAccount         string
	PageSize             int
	BgpGroupId           string
	BgpPeerId            string
	IsDefault            core.Bool
	OwnerId              int64
	PageNumber           int
}
type DescribeBgpPeersResponse struct {
	RequestId  string
	TotalCount int
	PageNumber int
	PageSize   int
	BgpPeers   DescribeBgpPeersBgpPeerList
}

type DescribeBgpPeersBgpPeerList []DescribeBgpPeersBgpPeer

func (list *DescribeBgpPeersBgpPeerList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeBgpPeersBgpPeer)
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

func (c *VpcClient) AssociateEipAddress(req *AssociateEipAddressArgs) (resp *AssociateEipAddressResponse, err error) {
	resp = &AssociateEipAddressResponse{}
	err = c.Request("AssociateEipAddress", req, resp)
	return
}

type AssociateEipAddressArgs struct {
	ResourceOwnerId      int64
	InstanceId           string
	ResourceOwnerAccount string
	OwnerAccount         string
	InstanceType         string
	AllocationId         string
	OwnerId              int64
}
type AssociateEipAddressResponse struct {
	RequestId string
}

func (c *VpcClient) DescribeRegions(req *DescribeRegionsArgs) (resp *DescribeRegionsResponse, err error) {
	resp = &DescribeRegionsResponse{}
	err = c.Request("DescribeRegions", req, resp)
	return
}

type DescribeRegionsRegion struct {
	RegionId  string
	LocalName string
}
type DescribeRegionsArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	ProductType          string
}
type DescribeRegionsResponse struct {
	RequestId string
	Regions   DescribeRegionsRegionList
}

type DescribeRegionsRegionList []DescribeRegionsRegion

func (list *DescribeRegionsRegionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRegionsRegion)
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

func (c *VpcClient) ModifyVpnGatewayAttribute(req *ModifyVpnGatewayAttributeArgs) (resp *ModifyVpnGatewayAttributeResponse, err error) {
	resp = &ModifyVpnGatewayAttributeResponse{}
	err = c.Request("ModifyVpnGatewayAttribute", req, resp)
	return
}

type ModifyVpnGatewayAttributeArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	Name                 string
	Description          string
	VpnGatewayId         string
	OwnerId              int64
}
type ModifyVpnGatewayAttributeResponse struct {
	RequestId      string
	VpnGatewayId   string
	VpcId          string
	VSwitchId      string
	InternetIp     string
	IntranetIp     string
	CreateTime     int64
	EndTime        int64
	Spec           string
	Name           string
	Description    string
	Status         string
	BusinessStatus string
}

func (c *VpcClient) TerminatePhysicalConnection(req *TerminatePhysicalConnectionArgs) (resp *TerminatePhysicalConnectionResponse, err error) {
	resp = &TerminatePhysicalConnectionResponse{}
	err = c.Request("TerminatePhysicalConnection", req, resp)
	return
}

type TerminatePhysicalConnectionArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ClientToken          string
	PhysicalConnectionId string
	OwnerAccount         string
	UserCidr             string
	OwnerId              int64
}
type TerminatePhysicalConnectionResponse struct {
	RequestId string
}

func (c *VpcClient) DescribeForwardTables(req *DescribeForwardTablesArgs) (resp *DescribeForwardTablesResponse, err error) {
	resp = &DescribeForwardTablesResponse{}
	err = c.Request("DescribeForwardTables", req, resp)
	return
}

type DescribeForwardTablesForwardTable struct {
	NatGatewayId   string
	ForwardTableId string
	CreationTime   string
	ForwardEntrys  DescribeForwardTablesForwardEntryList
}

type DescribeForwardTablesForwardEntry struct {
	ForwardTableId string
	ForwardEntryId string
	ExternalIp     string
	ExternalPort   string
	IpProtocol     string
	InternalIp     string
	InternalPort   string
	Status         string
}
type DescribeForwardTablesArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	ForwardTableId       string
	PageSize             int
	OwnerId              int64
	PageNumber           int
}
type DescribeForwardTablesResponse struct {
	RequestId     string
	TotalCount    int
	PageNumber    int
	PageSize      int
	ForwardTables DescribeForwardTablesForwardTableList
}

type DescribeForwardTablesForwardEntryList []DescribeForwardTablesForwardEntry

func (list *DescribeForwardTablesForwardEntryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeForwardTablesForwardEntry)
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

type DescribeForwardTablesForwardTableList []DescribeForwardTablesForwardTable

func (list *DescribeForwardTablesForwardTableList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeForwardTablesForwardTable)
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

func (c *VpcClient) DescribeVpnGateway(req *DescribeVpnGatewayArgs) (resp *DescribeVpnGatewayResponse, err error) {
	resp = &DescribeVpnGatewayResponse{}
	err = c.Request("DescribeVpnGateway", req, resp)
	return
}

type DescribeVpnGatewayArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	VpnGatewayId         string
	OwnerId              int64
}
type DescribeVpnGatewayResponse struct {
	RequestId      string
	VpnGatewayId   string
	VpcId          string
	VSwitchId      string
	InternetIp     string
	CreateTime     int64
	EndTime        int64
	Spec           string
	Name           string
	Description    string
	Status         string
	BusinessStatus string
	ChargeType     string
}

func (c *VpcClient) CreateRouteEntry(req *CreateRouteEntryArgs) (resp *CreateRouteEntryResponse, err error) {
	resp = &CreateRouteEntryResponse{}
	err = c.Request("CreateRouteEntry", req, resp)
	return
}

type CreateRouteEntryNextHopList struct {
	NextHopType string
	NextHopId   string
	Weight      int
}
type CreateRouteEntryArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ClientToken          string
	DestinationCidrBlock string
	OwnerAccount         string
	NextHopId            string
	OwnerId              int64
	NextHopType          string
	NextHopLists         CreateRouteEntryNextHopListList
	RouteTableId         string
}
type CreateRouteEntryResponse struct {
	RequestId string
}

type CreateRouteEntryNextHopListList []CreateRouteEntryNextHopList

func (list *CreateRouteEntryNextHopListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateRouteEntryNextHopList)
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

func (c *VpcClient) ModifyCustomerGatewayAttribute(req *ModifyCustomerGatewayAttributeArgs) (resp *ModifyCustomerGatewayAttributeResponse, err error) {
	resp = &ModifyCustomerGatewayAttributeResponse{}
	err = c.Request("ModifyCustomerGatewayAttribute", req, resp)
	return
}

type ModifyCustomerGatewayAttributeArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	Name                 string
	Description          string
	OwnerId              int64
	CustomerGatewayId    string
}
type ModifyCustomerGatewayAttributeResponse struct {
	RequestId         string
	CustomerGatewayId string
	IpAddress         string
	Name              string
	Description       string
	CreateTime        int64
}

func (c *VpcClient) CreateVpc(req *CreateVpcArgs) (resp *CreateVpcResponse, err error) {
	resp = &CreateVpcResponse{}
	err = c.Request("CreateVpc", req, resp)
	return
}

type CreateVpcArgs struct {
	VpcName              string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	CidrBlock            string
	Description          string
	UserCidr             string
	OwnerId              int64
}
type CreateVpcResponse struct {
	RequestId    string
	VpcId        string
	VRouterId    string
	RouteTableId string
}

func (c *VpcClient) ModifyPhysicalConnectionAttribute(req *ModifyPhysicalConnectionAttributeArgs) (resp *ModifyPhysicalConnectionAttributeResponse, err error) {
	resp = &ModifyPhysicalConnectionAttributeResponse{}
	err = c.Request("ModifyPhysicalConnectionAttribute", req, resp)
	return
}

type ModifyPhysicalConnectionAttributeArgs struct {
	RedundantPhysicalConnectionId string
	PeerLocation                  string
	ResourceOwnerId               int64
	PortType                      string
	CircuitCode                   string
	Bandwidth                     int
	ClientToken                   string
	ResourceOwnerAccount          string
	OwnerAccount                  string
	Description                   string
	OwnerId                       int64
	LineOperator                  string
	PhysicalConnectionId          string
	Name                          string
	UserCidr                      string
}
type ModifyPhysicalConnectionAttributeResponse struct {
	RequestId string
}

func (c *VpcClient) DescribeForwardTableEntries(req *DescribeForwardTableEntriesArgs) (resp *DescribeForwardTableEntriesResponse, err error) {
	resp = &DescribeForwardTableEntriesResponse{}
	err = c.Request("DescribeForwardTableEntries", req, resp)
	return
}

type DescribeForwardTableEntriesForwardTableEntry struct {
	ForwardTableId string
	ForwardEntryId string
	ExternalIp     string
	ExternalPort   string
	IpProtocol     string
	InternalIp     string
	InternalPort   string
	Status         string
}
type DescribeForwardTableEntriesArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ForwardEntryId       string
	OwnerAccount         string
	ForwardTableId       string
	PageSize             int
	OwnerId              int64
	PageNumber           int
}
type DescribeForwardTableEntriesResponse struct {
	RequestId           string
	TotalCount          int
	PageNumber          int
	PageSize            int
	ForwardTableEntries DescribeForwardTableEntriesForwardTableEntryList
}

type DescribeForwardTableEntriesForwardTableEntryList []DescribeForwardTableEntriesForwardTableEntry

func (list *DescribeForwardTableEntriesForwardTableEntryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeForwardTableEntriesForwardTableEntry)
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

func (c *VpcClient) DescribeVSwitches(req *DescribeVSwitchesArgs) (resp *DescribeVSwitchesResponse, err error) {
	resp = &DescribeVSwitchesResponse{}
	err = c.Request("DescribeVSwitches", req, resp)
	return
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
	IsDefault               core.Bool
}
type DescribeVSwitchesArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	PageNumber           int
	VSwitchId            string
	VpcId                string
	Name                 string
	PageSize             int
	ZoneId               string
	IsDefault            core.Bool
}
type DescribeVSwitchesResponse struct {
	RequestId  string
	TotalCount int
	PageNumber int
	PageSize   int
	VSwitches  DescribeVSwitchesVSwitchList
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

func (c *VpcClient) ModifyVpnConnectionAttribute(req *ModifyVpnConnectionAttributeArgs) (resp *ModifyVpnConnectionAttributeResponse, err error) {
	resp = &ModifyVpnConnectionAttributeResponse{}
	err = c.Request("ModifyVpnConnectionAttribute", req, resp)
	return
}

type ModifyVpnConnectionAttributeIkeConfig struct {
	Psk         string
	IkeVersion  string
	IkeMode     string
	IkeEncAlg   string
	IkeAuthAlg  string
	IkePfs      string
	IkeLifetime int64
	LocalId     string
	RemoteId    string
}

type ModifyVpnConnectionAttributeIpsecConfig struct {
	IpsecEncAlg   string
	IpsecAuthAlg  string
	IpsecPfs      string
	IpsecLifetime int64
}
type ModifyVpnConnectionAttributeArgs struct {
	IkeConfig            string
	ResourceOwnerId      int64
	RemoteSubnet         string
	EffectImmediately    core.Bool
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	IpsecConfig          string
	OwnerId              int64
	LocalSubnet          string
	VpnConnectionId      string
	Name                 string
}
type ModifyVpnConnectionAttributeResponse struct {
	RequestId         string
	VpnConnectionId   string
	CustomerGatewayId string
	VpnGatewayId      string
	Name              string
	Description       string
	LocalSubnet       string
	RemoteSubnet      string
	CreateTime        int64
	EffectImmediately core.Bool
	IkeConfig         ModifyVpnConnectionAttributeIkeConfig
	IpsecConfig       ModifyVpnConnectionAttributeIpsecConfig
}

func (c *VpcClient) DeleteVpnConnection(req *DeleteVpnConnectionArgs) (resp *DeleteVpnConnectionResponse, err error) {
	resp = &DeleteVpnConnectionResponse{}
	err = c.Request("DeleteVpnConnection", req, resp)
	return
}

type DeleteVpnConnectionArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ClientToken          string
	VpnConnectionId      string
	OwnerAccount         string
	OwnerId              int64
}
type DeleteVpnConnectionResponse struct {
	RequestId string
}

func (c *VpcClient) AssociateGlobalAccelerationInstance(req *AssociateGlobalAccelerationInstanceArgs) (resp *AssociateGlobalAccelerationInstanceResponse, err error) {
	resp = &AssociateGlobalAccelerationInstanceResponse{}
	err = c.Request("AssociateGlobalAccelerationInstance", req, resp)
	return
}

type AssociateGlobalAccelerationInstanceArgs struct {
	ResourceOwnerId              int64
	ResourceOwnerAccount         string
	OwnerAccount                 string
	BackendServerId              string
	OwnerId                      int64
	GlobalAccelerationInstanceId string
	BackendServerRegionId        string
	BackendServerType            string
}
type AssociateGlobalAccelerationInstanceResponse struct {
	RequestId string
}

func (c *VpcClient) ModifyCommonBandwidthPackageAttribute(req *ModifyCommonBandwidthPackageAttributeArgs) (resp *ModifyCommonBandwidthPackageAttributeResponse, err error) {
	resp = &ModifyCommonBandwidthPackageAttributeResponse{}
	err = c.Request("ModifyCommonBandwidthPackageAttribute", req, resp)
	return
}

type ModifyCommonBandwidthPackageAttributeArgs struct {
	ResourceOwnerId      int64
	BandwidthPackageId   string
	ResourceOwnerAccount string
	OwnerAccount         string
	Name                 string
	Description          string
	OwnerId              int64
}
type ModifyCommonBandwidthPackageAttributeResponse struct {
	RequestId string
}

func (c *VpcClient) DescribeRouteTableList(req *DescribeRouteTableListArgs) (resp *DescribeRouteTableListResponse, err error) {
	resp = &DescribeRouteTableListResponse{}
	err = c.Request("DescribeRouteTableList", req, resp)
	return
}

type DescribeRouteTableListRouterTableListType struct {
	VpcId          string
	RouterType     string
	RouterId       string
	RouteTableId   string
	RouteTableName string
	RouteTableType string
	Description    string
	CreationTime   string
}
type DescribeRouteTableListArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	Bandwidth            string
	OwnerAccount         string
	OwnerId              int64
	PageNumber           int
	RouterType           string
	KbpsBandwidth        string
	RouteTableName       string
	RouterId             string
	VpcId                string
	ResourceUid          int64
	PageSize             int
	ResourceBid          string
	RouteTableId         string
}
type DescribeRouteTableListResponse struct {
	RequestId       string
	Code            string
	Message         string
	Success         core.Bool
	PageSize        int
	PageNumber      int
	TotalCount      int
	RouterTableList DescribeRouteTableListRouterTableListTypeList
}

type DescribeRouteTableListRouterTableListTypeList []DescribeRouteTableListRouterTableListType

func (list *DescribeRouteTableListRouterTableListTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRouteTableListRouterTableListType)
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

func (c *VpcClient) ReleaseEipAddress(req *ReleaseEipAddressArgs) (resp *ReleaseEipAddressResponse, err error) {
	resp = &ReleaseEipAddressResponse{}
	err = c.Request("ReleaseEipAddress", req, resp)
	return
}

type ReleaseEipAddressArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	AllocationId         string
	OwnerId              int64
}
type ReleaseEipAddressResponse struct {
	RequestId string
}

func (c *VpcClient) DisableVpcClassicLink(req *DisableVpcClassicLinkArgs) (resp *DisableVpcClassicLinkResponse, err error) {
	resp = &DisableVpcClassicLinkResponse{}
	err = c.Request("DisableVpcClassicLink", req, resp)
	return
}

type DisableVpcClassicLinkArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ClientToken          string
	VpcId                string
	OwnerAccount         string
	OwnerId              int64
}
type DisableVpcClassicLinkResponse struct {
	RequestId string
}

func (c *VpcClient) RecoverVirtualBorderRouter(req *RecoverVirtualBorderRouterArgs) (resp *RecoverVirtualBorderRouterResponse, err error) {
	resp = &RecoverVirtualBorderRouterResponse{}
	err = c.Request("RecoverVirtualBorderRouter", req, resp)
	return
}

type RecoverVirtualBorderRouterArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	UserCidr             string
	VbrId                string
	OwnerId              int64
}
type RecoverVirtualBorderRouterResponse struct {
	RequestId string
}

func (c *VpcClient) ModifyCommonBandwidthPackagePayType(req *ModifyCommonBandwidthPackagePayTypeArgs) (resp *ModifyCommonBandwidthPackagePayTypeResponse, err error) {
	resp = &ModifyCommonBandwidthPackagePayTypeResponse{}
	err = c.Request("ModifyCommonBandwidthPackagePayType", req, resp)
	return
}

type ModifyCommonBandwidthPackagePayTypeArgs struct {
	ResourceOwnerId      int64
	BandwidthPackageId   string
	AutoPay              core.Bool
	ResourceOwnerAccount string
	Bandwidth            string
	OwnerAccount         string
	OwnerId              int64
	Duration             int
	KbpsBandwidth        string
	ResourceUid          int64
	ResourceBid          string
	PayType              string
	PricingCycle         string
}
type ModifyCommonBandwidthPackagePayTypeResponse struct {
	RequestId string
	OrderId   int64
	Code      string
	Message   string
}
