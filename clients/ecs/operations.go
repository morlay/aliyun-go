package ecs

import "encoding/json"

func (c *EcsClient) DeletePhysicalConnection(req *DeletePhysicalConnectionArgs) (resp *DeletePhysicalConnectionResponse, err error) {
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

func (c *EcsClient) ImportKeyPair(req *ImportKeyPairArgs) (resp *ImportKeyPairResponse, err error) {
	resp = &ImportKeyPairResponse{}
	err = c.Request("ImportKeyPair", req, resp)
	return
}

type ImportKeyPairArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	PublicKeyBody        string
	KeyPairName          string
	OwnerId              int64
}
type ImportKeyPairResponse struct {
	RequestId          string
	KeyPairName        string
	KeyPairFingerPrint string
}

func (c *EcsClient) DescribeLimitation(req *DescribeLimitationArgs) (resp *DescribeLimitationResponse, err error) {
	resp = &DescribeLimitationResponse{}
	err = c.Request("DescribeLimitation", req, resp)
	return
}

type DescribeLimitationArgs struct {
	Limitation           string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
}
type DescribeLimitationResponse struct {
	RequestId  string
	Limitation string
	Value      string
}

func (c *EcsClient) AddIpRange(req *AddIpRangeArgs) (resp *AddIpRangeResponse, err error) {
	resp = &AddIpRangeResponse{}
	err = c.Request("AddIpRange", req, resp)
	return
}

type AddIpRangeArgs struct {
	IpAddress            string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	ZoneId               string
	OwnerId              int64
}
type AddIpRangeResponse struct {
	RequestId string
}

func (c *EcsClient) CreateKeyPair(req *CreateKeyPairArgs) (resp *CreateKeyPairResponse, err error) {
	resp = &CreateKeyPairResponse{}
	err = c.Request("CreateKeyPair", req, resp)
	return
}

type CreateKeyPairArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	KeyPairName          string
	OwnerId              int64
}
type CreateKeyPairResponse struct {
	RequestId          string
	KeyPairName        string
	KeyPairFingerPrint string
	PrivateKeyBody     string
}

func (c *EcsClient) ModifyInstanceSpec(req *ModifyInstanceSpecArgs) (resp *ModifyInstanceSpecResponse, err error) {
	resp = &ModifyInstanceSpecResponse{}
	err = c.Request("ModifyInstanceSpec", req, resp)
	return
}

type ModifyInstanceSpecArgs struct {
	ResourceOwnerId                  int64
	ResourceOwnerAccount             string
	ClientToken                      string
	AllowMigrateAcrossZone           bool
	OwnerAccount                     string
	InternetMaxBandwidthOut          int
	OwnerId                          int64
	TemporaryInternetMaxBandwidthOut int
	SystemDiskCategory               string
	TemporaryStartTime               string
	Async                            bool
	InstanceId                       string
	InstanceType                     string
	TemporaryEndTime                 string
	InternetMaxBandwidthIn           int
}
type ModifyInstanceSpecResponse struct {
	RequestId string
}

func (c *EcsClient) ModifyHpcClusterAttribute(req *ModifyHpcClusterAttributeArgs) (resp *ModifyHpcClusterAttributeResponse, err error) {
	resp = &ModifyHpcClusterAttributeResponse{}
	err = c.Request("ModifyHpcClusterAttribute", req, resp)
	return
}

type ModifyHpcClusterAttributeArgs struct {
	ResourceOwnerId      int64
	HpcClusterId         string
	ClientToken          string
	Description          string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	Name                 string
}
type ModifyHpcClusterAttributeResponse struct {
	RequestId string
}

func (c *EcsClient) DescribeNewProjectEipMonitorData(req *DescribeNewProjectEipMonitorDataArgs) (resp *DescribeNewProjectEipMonitorDataResponse, err error) {
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

func (c *EcsClient) JoinSecurityGroup(req *JoinSecurityGroupArgs) (resp *JoinSecurityGroupResponse, err error) {
	resp = &JoinSecurityGroupResponse{}
	err = c.Request("JoinSecurityGroup", req, resp)
	return
}

type JoinSecurityGroupArgs struct {
	ResourceOwnerId      int64
	InstanceId           string
	ResourceOwnerAccount string
	OwnerAccount         string
	SecurityGroupId      string
	OwnerId              int64
}
type JoinSecurityGroupResponse struct {
	RequestId string
}

func (c *EcsClient) DescribePhysicalConnections(req *DescribePhysicalConnectionsArgs) (resp *DescribePhysicalConnectionsResponse, err error) {
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

func (c *EcsClient) ModifyInstanceAttribute(req *ModifyInstanceAttributeArgs) (resp *ModifyInstanceAttributeResponse, err error) {
	resp = &ModifyInstanceAttributeResponse{}
	err = c.Request("ModifyInstanceAttribute", req, resp)
	return
}

type ModifyInstanceAttributeArgs struct {
	UserData             string
	ResourceOwnerId      int64
	Password             string
	HostName             string
	InstanceId           string
	InstanceName         string
	ResourceOwnerAccount string
	Recyclable           bool
	OwnerAccount         string
	Description          string
	OwnerId              int64
}
type ModifyInstanceAttributeResponse struct {
	RequestId string
}

func (c *EcsClient) DescribeRouteTables(req *DescribeRouteTablesArgs) (resp *DescribeRouteTablesResponse, err error) {
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
	RouteTableName       string
	VRouterId            string
	ResourceOwnerAccount string
	RouterId             string
	OwnerAccount         string
	PageSize             int
	OwnerId              int64
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

func (c *EcsClient) CreateVirtualBorderRouter(req *CreateVirtualBorderRouterArgs) (resp *CreateVirtualBorderRouterResponse, err error) {
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

func (c *EcsClient) UnassociateHaVip(req *UnassociateHaVipArgs) (resp *UnassociateHaVipResponse, err error) {
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

func (c *EcsClient) ModifySnapshotAttribute(req *ModifySnapshotAttributeArgs) (resp *ModifySnapshotAttributeResponse, err error) {
	resp = &ModifySnapshotAttributeResponse{}
	err = c.Request("ModifySnapshotAttribute", req, resp)
	return
}

type ModifySnapshotAttributeArgs struct {
	ResourceOwnerId      int64
	SnapshotId           string
	ResourceOwnerAccount string
	OwnerAccount         string
	Description          string
	SnapshotName         string
	OwnerId              int64
}
type ModifySnapshotAttributeResponse struct {
	RequestId string
}

func (c *EcsClient) DeleteRouterInterface(req *DeleteRouterInterfaceArgs) (resp *DeleteRouterInterfaceResponse, err error) {
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

func (c *EcsClient) EnablePhysicalConnection(req *EnablePhysicalConnectionArgs) (resp *EnablePhysicalConnectionResponse, err error) {
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

func (c *EcsClient) DescribeTags(req *DescribeTagsArgs) (resp *DescribeTagsResponse, err error) {
	resp = &DescribeTagsResponse{}
	err = c.Request("DescribeTags", req, resp)
	return
}

type DescribeTagsTag struct {
	TagKey            string
	TagValue          string
	ResourceTypeCount DescribeTagsResourceTypeCount
}

type DescribeTagsResourceTypeCount struct {
	Instance      int
	Disk          int
	Volume        int
	Image         int
	Snapshot      int
	Securitygroup int
}
type DescribeTagsArgs struct {
	Tag4Value            string
	ResourceOwnerId      int64
	ResourceId           string
	Tag2Key              string
	Tag5Key              string
	ResourceOwnerAccount string
	Tag3Key              string
	OwnerId              int64
	ResourceType         string
	Tag5Value            string
	PageNumber           int
	Tag1Key              string
	Tag1Value            string
	Tag2Value            string
	PageSize             int
	Tag4Key              string
	Tag3Value            string
}
type DescribeTagsResponse struct {
	RequestId  string
	PageSize   int
	PageNumber int
	TotalCount int
	Tags       DescribeTagsTagList
}

type DescribeTagsTagList []DescribeTagsTag

func (list *DescribeTagsTagList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeTagsTag)
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

func (c *EcsClient) CreateHpcCluster(req *CreateHpcClusterArgs) (resp *CreateHpcClusterResponse, err error) {
	resp = &CreateHpcClusterResponse{}
	err = c.Request("CreateHpcCluster", req, resp)
	return
}

type CreateHpcClusterArgs struct {
	ResourceOwnerId      int64
	ClientToken          string
	Description          string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	Name                 string
}
type CreateHpcClusterResponse struct {
	RequestId    string
	HpcClusterId string
}

func (c *EcsClient) DescribeSpotPriceHistory(req *DescribeSpotPriceHistoryArgs) (resp *DescribeSpotPriceHistoryResponse, err error) {
	resp = &DescribeSpotPriceHistoryResponse{}
	err = c.Request("DescribeSpotPriceHistory", req, resp)
	return
}

type DescribeSpotPriceHistorySpotPriceType struct {
	ZoneId       string
	InstanceType string
	IoOptimized  string
	Timestamp    string
	NetworkType  string
	SpotPrice    float32
	OriginPrice  float32
}
type DescribeSpotPriceHistoryArgs struct {
	ResourceOwnerId      int64
	IoOptimized          string
	NetworkType          string
	StartTime            string
	InstanceType         string
	Offset               int
	ResourceOwnerAccount string
	OwnerAccount         string
	EndTime              string
	OSType               string
	OwnerId              int64
	ZoneId               string
}
type DescribeSpotPriceHistoryResponse struct {
	RequestId  string
	NextOffset int
	Currency   string
	SpotPrices DescribeSpotPriceHistorySpotPriceTypeList
}

type DescribeSpotPriceHistorySpotPriceTypeList []DescribeSpotPriceHistorySpotPriceType

func (list *DescribeSpotPriceHistorySpotPriceTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSpotPriceHistorySpotPriceType)
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

func (c *EcsClient) CancelAgreement(req *CancelAgreementArgs) (resp *CancelAgreementResponse, err error) {
	resp = &CancelAgreementResponse{}
	err = c.Request("CancelAgreement", req, resp)
	return
}

type CancelAgreementArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	AgreementType        string
}
type CancelAgreementResponse struct {
	RequestId string
}

func (c *EcsClient) DescribeInstanceAttribute(req *DescribeInstanceAttributeArgs) (resp *DescribeInstanceAttributeResponse, err error) {
	resp = &DescribeInstanceAttributeResponse{}
	err = c.Request("DescribeInstanceAttribute", req, resp)
	return
}

type DescribeInstanceAttributeLockReason struct {
	LockReason string
}

type DescribeInstanceAttributeVpcAttributes struct {
	VpcId            string
	VSwitchId        string
	NatIpAddress     string
	PrivateIpAddress DescribeInstanceAttributePrivateIpAddresList
}

type DescribeInstanceAttributeEipAddress struct {
	AllocationId       string
	IpAddress          string
	Bandwidth          int
	InternetChargeType string
}
type DescribeInstanceAttributeArgs struct {
	ResourceOwnerId      int64
	InstanceId           string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
}
type DescribeInstanceAttributeResponse struct {
	RequestId               string
	InstanceId              string
	InstanceName            string
	ImageId                 string
	RegionId                string
	ZoneId                  string
	ClusterId               string
	InstanceType            string
	Cpu                     int
	Memory                  int
	HostName                string
	Status                  string
	InternetChargeType      string
	InternetMaxBandwidthIn  int
	InternetMaxBandwidthOut int
	VlanId                  string
	SerialNumber            string
	CreationTime            string
	Description             string
	InstanceNetworkType     string
	IoOptimized             string
	InstanceChargeType      string
	ExpiredTime             string
	StoppedMode             string
	OperationLocks          DescribeInstanceAttributeLockReasonList
	SecurityGroupIds        DescribeInstanceAttributeSecurityGroupIdList
	PublicIpAddress         DescribeInstanceAttributePublicIpAddresList
	InnerIpAddress          DescribeInstanceAttributeInnerIpAddresList
	VpcAttributes           DescribeInstanceAttributeVpcAttributes
	EipAddress              DescribeInstanceAttributeEipAddress
}

type DescribeInstanceAttributePrivateIpAddresList []string

func (list *DescribeInstanceAttributePrivateIpAddresList) UnmarshalJSON(data []byte) error {
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

type DescribeInstanceAttributeLockReasonList []DescribeInstanceAttributeLockReason

func (list *DescribeInstanceAttributeLockReasonList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInstanceAttributeLockReason)
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

type DescribeInstanceAttributeSecurityGroupIdList []string

func (list *DescribeInstanceAttributeSecurityGroupIdList) UnmarshalJSON(data []byte) error {
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

type DescribeInstanceAttributePublicIpAddresList []string

func (list *DescribeInstanceAttributePublicIpAddresList) UnmarshalJSON(data []byte) error {
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

type DescribeInstanceAttributeInnerIpAddresList []string

func (list *DescribeInstanceAttributeInnerIpAddresList) UnmarshalJSON(data []byte) error {
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

func (c *EcsClient) CreateSnapshot(req *CreateSnapshotArgs) (resp *CreateSnapshotResponse, err error) {
	resp = &CreateSnapshotResponse{}
	err = c.Request("CreateSnapshot", req, resp)
	return
}

type CreateSnapshotArgs struct {
	Tag4Value            string
	ResourceOwnerId      int64
	Tag2Key              string
	Tag5Key              string
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	Description          string
	SnapshotName         string
	Tag3Key              string
	OwnerId              int64
	Tag5Value            string
	Tag1Key              string
	Tag1Value            string
	Tag2Value            string
	Tag4Key              string
	DiskId               string
	Tag3Value            string
}
type CreateSnapshotResponse struct {
	RequestId  string
	SnapshotId string
}

func (c *EcsClient) RenewInstance(req *RenewInstanceArgs) (resp *RenewInstanceResponse, err error) {
	resp = &RenewInstanceResponse{}
	err = c.Request("RenewInstance", req, resp)
	return
}

type RenewInstanceArgs struct {
	ResourceOwnerId      int64
	Period               int
	PeriodUnit           string
	InstanceId           string
	ClientToken          string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
}
type RenewInstanceResponse struct {
	RequestId string
}

func (c *EcsClient) ModifyBandwidthPackageSpec(req *ModifyBandwidthPackageSpecArgs) (resp *ModifyBandwidthPackageSpecResponse, err error) {
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

func (c *EcsClient) DeleteRouteEntry(req *DeleteRouteEntryArgs) (resp *DeleteRouteEntryResponse, err error) {
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

func (c *EcsClient) ReleasePublicIpAddress(req *ReleasePublicIpAddressArgs) (resp *ReleasePublicIpAddressResponse, err error) {
	resp = &ReleasePublicIpAddressResponse{}
	err = c.Request("ReleasePublicIpAddress", req, resp)
	return
}

type ReleasePublicIpAddressArgs struct {
	ResourceOwnerId      int64
	PublicIpAddress      string
	InstanceId           string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
}
type ReleasePublicIpAddressResponse struct {
	RequestId string
}

func (c *EcsClient) AuthorizeSecurityGroup(req *AuthorizeSecurityGroupArgs) (resp *AuthorizeSecurityGroupResponse, err error) {
	resp = &AuthorizeSecurityGroupResponse{}
	err = c.Request("AuthorizeSecurityGroup", req, resp)
	return
}

type AuthorizeSecurityGroupArgs struct {
	NicType                 string
	ResourceOwnerId         int64
	SourcePortRange         string
	ClientToken             string
	SecurityGroupId         string
	Description             string
	SourceGroupOwnerId      int64
	SourceGroupOwnerAccount string
	Policy                  string
	PortRange               string
	ResourceOwnerAccount    string
	IpProtocol              string
	OwnerAccount            string
	SourceCidrIp            string
	OwnerId                 int64
	Priority                string
	DestCidrIp              string
	SourceGroupId           string
}
type AuthorizeSecurityGroupResponse struct {
	RequestId string
}

func (c *EcsClient) DescribeInstanceTypes(req *DescribeInstanceTypesArgs) (resp *DescribeInstanceTypesResponse, err error) {
	resp = &DescribeInstanceTypesResponse{}
	err = c.Request("DescribeInstanceTypes", req, resp)
	return
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
}
type DescribeInstanceTypesArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	InstanceTypeFamily   string
	OwnerId              int64
}
type DescribeInstanceTypesResponse struct {
	RequestId     string
	InstanceTypes DescribeInstanceTypesInstanceTypeList
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

func (c *EcsClient) SignAgreement(req *SignAgreementArgs) (resp *SignAgreementResponse, err error) {
	resp = &SignAgreementResponse{}
	err = c.Request("SignAgreement", req, resp)
	return
}

type SignAgreementArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	AgreementType        string
}
type SignAgreementResponse struct {
	RequestId string
}

func (c *EcsClient) ModifyImageSharePermission(req *ModifyImageSharePermissionArgs) (resp *ModifyImageSharePermissionResponse, err error) {
	resp = &ModifyImageSharePermissionResponse{}
	err = c.Request("ModifyImageSharePermission", req, resp)
	return
}

type ModifyImageSharePermissionArgs struct {
	AddAccount1          string
	ResourceOwnerId      int64
	ImageId              string
	AddAccount9          string
	AddAccount8          string
	AddAccount7          string
	AddAccount6          string
	AddAccount5          string
	AddAccount10         string
	AddAccount4          string
	AddAccount3          string
	AddAccount2          string
	ResourceOwnerAccount string
	OwnerAccount         string
	RemoveAccount1       string
	RemoveAccount2       string
	RemoveAccount3       string
	RemoveAccount4       string
	OwnerId              int64
	RemoveAccount9       string
	RemoveAccount5       string
	RemoveAccount6       string
	RemoveAccount7       string
	RemoveAccount8       string
	RemoveAccount10      string
}
type ModifyImageSharePermissionResponse struct {
	RequestId string
}

func (c *EcsClient) CopyImage(req *CopyImageArgs) (resp *CopyImageResponse, err error) {
	resp = &CopyImageResponse{}
	err = c.Request("CopyImage", req, resp)
	return
}

type CopyImageArgs struct {
	Tag4Value              string
	ResourceOwnerId        int64
	ImageId                string
	Tag2Key                string
	Tag5Key                string
	ResourceOwnerAccount   string
	DestinationImageName   string
	DestinationRegionId    string
	OwnerAccount           string
	Tag3Key                string
	OwnerId                int64
	Tag5Value              string
	Tag1Key                string
	Tag1Value              string
	Encrypted              bool
	Tag2Value              string
	Tag4Key                string
	DestinationDescription string
	Tag3Value              string
}
type CopyImageResponse struct {
	RequestId string
	ImageId   string
}

func (c *EcsClient) DeleteVpc(req *DeleteVpcArgs) (resp *DeleteVpcResponse, err error) {
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

func (c *EcsClient) DescribeEipAddresses(req *DescribeEipAddressesArgs) (resp *DescribeEipAddressesResponse, err error) {
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
	EipBandwidth       string
	InternetChargeType string
	AllocationTime     string
	InstanceType       string
	ChargeType         string
	ExpiredTime        string
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

func (c *EcsClient) RevokeSecurityGroup(req *RevokeSecurityGroupArgs) (resp *RevokeSecurityGroupResponse, err error) {
	resp = &RevokeSecurityGroupResponse{}
	err = c.Request("RevokeSecurityGroup", req, resp)
	return
}

type RevokeSecurityGroupArgs struct {
	NicType                 string
	ResourceOwnerId         int64
	SourcePortRange         string
	ClientToken             string
	SecurityGroupId         string
	Description             string
	SourceGroupOwnerId      int64
	SourceGroupOwnerAccount string
	Policy                  string
	PortRange               string
	ResourceOwnerAccount    string
	IpProtocol              string
	OwnerAccount            string
	SourceCidrIp            string
	OwnerId                 int64
	Priority                string
	DestCidrIp              string
	SourceGroupId           string
}
type RevokeSecurityGroupResponse struct {
	RequestId string
}

func (c *EcsClient) CreateAutoSnapshotPolicy(req *CreateAutoSnapshotPolicyArgs) (resp *CreateAutoSnapshotPolicyResponse, err error) {
	resp = &CreateAutoSnapshotPolicyResponse{}
	err = c.Request("CreateAutoSnapshotPolicy", req, resp)
	return
}

type CreateAutoSnapshotPolicyArgs struct {
	ResourceOwnerId        int64
	ResourceOwnerAccount   string
	TimePoints             string
	RetentionDays          int
	OwnerId                int64
	RepeatWeekdays         string
	AutoSnapshotPolicyName string
}
type CreateAutoSnapshotPolicyResponse struct {
	RequestId            string
	AutoSnapshotPolicyId string
}

func (c *EcsClient) CancelPhysicalConnection(req *CancelPhysicalConnectionArgs) (resp *CancelPhysicalConnectionResponse, err error) {
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

func (c *EcsClient) ConnectRouterInterface(req *ConnectRouterInterfaceArgs) (resp *ConnectRouterInterfaceResponse, err error) {
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

func (c *EcsClient) DescribeVRouters(req *DescribeVRoutersArgs) (resp *DescribeVRoutersResponse, err error) {
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

func (c *EcsClient) CancelCopyImage(req *CancelCopyImageArgs) (resp *CancelCopyImageResponse, err error) {
	resp = &CancelCopyImageResponse{}
	err = c.Request("CancelCopyImage", req, resp)
	return
}

type CancelCopyImageArgs struct {
	ResourceOwnerId      int64
	ImageId              string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
}
type CancelCopyImageResponse struct {
	RequestId string
}

func (c *EcsClient) CreateRouterInterface(req *CreateRouterInterfaceArgs) (resp *CreateRouterInterfaceResponse, err error) {
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

func (c *EcsClient) CreateForwardEntry(req *CreateForwardEntryArgs) (resp *CreateForwardEntryResponse, err error) {
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

func (c *EcsClient) AssociateHaVip(req *AssociateHaVipArgs) (resp *AssociateHaVipResponse, err error) {
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

func (c *EcsClient) DescribeInvocationResults(req *DescribeInvocationResultsArgs) (resp *DescribeInvocationResultsResponse, err error) {
	resp = &DescribeInvocationResultsResponse{}
	err = c.Request("DescribeInvocationResults", req, resp)
	return
}

type DescribeInvocationResultsInvocation struct {
	InvokeId    string
	PageSize    int64
	PageNumber  int64
	TotalCount  int64
	Status      string
	ResultLists DescribeInvocationResultsResultItemList
}

type DescribeInvocationResultsResultItem struct {
	InstanceId   string
	FinishedTime string
	Output       string
	ExitCode     int64
}
type DescribeInvocationResultsArgs struct {
	ResourceOwnerId      int64
	PageNumber           int64
	PageSize             int64
	InvokeId             string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	InstanceId           string
}
type DescribeInvocationResultsResponse struct {
	RequestId  string
	Invocation DescribeInvocationResultsInvocation
}

type DescribeInvocationResultsResultItemList []DescribeInvocationResultsResultItem

func (list *DescribeInvocationResultsResultItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInvocationResultsResultItem)
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

func (c *EcsClient) DescribeTagKeys(req *DescribeTagKeysArgs) (resp *DescribeTagKeysResponse, err error) {
	resp = &DescribeTagKeysResponse{}
	err = c.Request("DescribeTagKeys", req, resp)
	return
}

type DescribeTagKeysArgs struct {
	ResourceOwnerId      int64
	ResourceId           string
	ResourceOwnerAccount string
	PageSize             int
	OwnerId              int64
	ResourceType         string
	PageNumber           int
}
type DescribeTagKeysResponse struct {
	RequestId  string
	PageSize   int
	PageNumber int
	TotalCount int
	TagKeys    DescribeTagKeysTagKeyList
}

type DescribeTagKeysTagKeyList []string

func (list *DescribeTagKeysTagKeyList) UnmarshalJSON(data []byte) error {
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

func (c *EcsClient) DescribeCommands(req *DescribeCommandsArgs) (resp *DescribeCommandsResponse, err error) {
	resp = &DescribeCommandsResponse{}
	err = c.Request("DescribeCommands", req, resp)
	return
}

type DescribeCommandsCommand struct {
	CommandId      string
	Name           string
	Type           string
	Description    string
	CommandContent string
	WorkingDir     string
	Timeout        int64
}
type DescribeCommandsArgs struct {
	ResourceOwnerId      int64
	Description          string
	Type                 string
	CommandId            string
	PageNumber           int64
	PageSize             int64
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	Name                 string
}
type DescribeCommandsResponse struct {
	RequestId   string
	TotalCount  int64
	PageNumber  int64
	PageSize    int64
	CommandList DescribeCommandsCommandList
}

type DescribeCommandsCommandList []DescribeCommandsCommand

func (list *DescribeCommandsCommandList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCommandsCommand)
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

func (c *EcsClient) DeleteCommand(req *DeleteCommandArgs) (resp *DeleteCommandResponse, err error) {
	resp = &DeleteCommandResponse{}
	err = c.Request("DeleteCommand", req, resp)
	return
}

type DeleteCommandArgs struct {
	ResourceOwnerId      int64
	CommandId            string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
}
type DeleteCommandResponse struct {
	RequestId string
}

func (c *EcsClient) ResetDisk(req *ResetDiskArgs) (resp *ResetDiskResponse, err error) {
	resp = &ResetDiskResponse{}
	err = c.Request("ResetDisk", req, resp)
	return
}

type ResetDiskArgs struct {
	ResourceOwnerId      int64
	SnapshotId           string
	ResourceOwnerAccount string
	OwnerAccount         string
	DiskId               string
	OwnerId              int64
}
type ResetDiskResponse struct {
	RequestId string
}

func (c *EcsClient) DescribeHpcClusters(req *DescribeHpcClustersArgs) (resp *DescribeHpcClustersResponse, err error) {
	resp = &DescribeHpcClustersResponse{}
	err = c.Request("DescribeHpcClusters", req, resp)
	return
}

type DescribeHpcClustersHpcCluster struct {
	HpcClusterId string
	Name         string
	Description  string
}
type DescribeHpcClustersArgs struct {
	ResourceOwnerId      int64
	ClientToken          string
	PageNumber           int
	PageSize             int
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	HpcClusterIds        string
}
type DescribeHpcClustersResponse struct {
	RequestId   string
	TotalCount  int
	PageNumber  int
	PageSize    int
	HpcClusters DescribeHpcClustersHpcClusterList
}

type DescribeHpcClustersHpcClusterList []DescribeHpcClustersHpcCluster

func (list *DescribeHpcClustersHpcClusterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeHpcClustersHpcCluster)
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

func (c *EcsClient) ModifySecurityGroupPolicy(req *ModifySecurityGroupPolicyArgs) (resp *ModifySecurityGroupPolicyResponse, err error) {
	resp = &ModifySecurityGroupPolicyResponse{}
	err = c.Request("ModifySecurityGroupPolicy", req, resp)
	return
}

type ModifySecurityGroupPolicyArgs struct {
	ResourceOwnerId      int64
	ClientToken          string
	ResourceOwnerAccount string
	OwnerAccount         string
	SecurityGroupId      string
	OwnerId              int64
	InnerAccessPolicy    string
}
type ModifySecurityGroupPolicyResponse struct {
	RequestId string
}

func (c *EcsClient) ModifyHaVipAttribute(req *ModifyHaVipAttributeArgs) (resp *ModifyHaVipAttributeResponse, err error) {
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

func (c *EcsClient) DeleteSnapshot(req *DeleteSnapshotArgs) (resp *DeleteSnapshotResponse, err error) {
	resp = &DeleteSnapshotResponse{}
	err = c.Request("DeleteSnapshot", req, resp)
	return
}

type DeleteSnapshotArgs struct {
	ResourceOwnerId      int64
	SnapshotId           string
	ResourceOwnerAccount string
	OwnerAccount         string
	Force                bool
	OwnerId              int64
}
type DeleteSnapshotResponse struct {
	RequestId string
}

func (c *EcsClient) CreateCommand(req *CreateCommandArgs) (resp *CreateCommandResponse, err error) {
	resp = &CreateCommandResponse{}
	err = c.Request("CreateCommand", req, resp)
	return
}

type CreateCommandArgs struct {
	ResourceOwnerId      int64
	WorkingDir           string
	Description          string
	Type                 string
	CommandContent       string
	Timeout              int64
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	Name                 string
}
type CreateCommandResponse struct {
	RequestId string
	CommandId string
}

func (c *EcsClient) RemoveBandwidthPackageIps(req *RemoveBandwidthPackageIpsArgs) (resp *RemoveBandwidthPackageIpsResponse, err error) {
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

func (c *EcsClient) DeleteVirtualBorderRouter(req *DeleteVirtualBorderRouterArgs) (resp *DeleteVirtualBorderRouterResponse, err error) {
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

func (c *EcsClient) DescribeKeyPairs(req *DescribeKeyPairsArgs) (resp *DescribeKeyPairsResponse, err error) {
	resp = &DescribeKeyPairsResponse{}
	err = c.Request("DescribeKeyPairs", req, resp)
	return
}

type DescribeKeyPairsKeyPair struct {
	KeyPairName        string
	KeyPairFingerPrint string
}
type DescribeKeyPairsArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	KeyPairFingerPrint   string
	PageSize             int
	KeyPairName          string
	OwnerId              int64
	PageNumber           int
}
type DescribeKeyPairsResponse struct {
	RequestId  string
	TotalCount int
	PageNumber int
	PageSize   int
	KeyPairs   DescribeKeyPairsKeyPairList
}

type DescribeKeyPairsKeyPairList []DescribeKeyPairsKeyPair

func (list *DescribeKeyPairsKeyPairList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeKeyPairsKeyPair)
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

func (c *EcsClient) ActivateRouterInterface(req *ActivateRouterInterfaceArgs) (resp *ActivateRouterInterfaceResponse, err error) {
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

func (c *EcsClient) RunInstances(req *RunInstancesArgs) (resp *RunInstancesResponse, err error) {
	resp = &RunInstancesResponse{}
	err = c.Request("RunInstances", req, resp)
	return
}

type RunInstancesTag struct {
	Key   string
	Value string
}

type RunInstancesNetworkInterface struct {
	PrimaryIpAddress     string
	VSwitchId            string
	SecurityGroupId      string
	NetworkInterfaceName string
	Description          string
}

type RunInstancesDataDisk struct {
	Size               int
	SnapshotId         string
	Category           string
	Encrypted          string
	DiskName           string
	Description        string
	DeleteWithInstance bool
}
type RunInstancesArgs struct {
	ResourceOwnerId             int64
	SecurityEnhancementStrategy string
	KeyPairName                 string
	SpotPriceLimit              float32
	HostName                    string
	Password                    string
	Tags                        RunInstancesTagList
	OwnerId                     int64
	VSwitchId                   string
	SpotStrategy                string
	InstanceName                string
	InternetChargeType          string
	ZoneId                      string
	InternetMaxBandwidthIn      int
	ImageId                     string
	ClientToken                 string
	IoOptimized                 string
	SecurityGroupId             string
	InternetMaxBandwidthOut     int
	Description                 string
	SystemDiskCategory          string
	UserData                    string
	InstanceType                string
	NetworkInterfaces           RunInstancesNetworkInterfaceList
	Amount                      int
	ResourceOwnerAccount        string
	OwnerAccount                string
	SystemDiskDiskName          string
	RamRoleName                 string
	AutoReleaseTime             string
	DataDisks                   RunInstancesDataDiskList
	SystemDiskSize              string
	SystemDiskDescription       string
}
type RunInstancesResponse struct {
	RequestId      string
	InstanceIdSets RunInstancesInstanceIdSetList
}

type RunInstancesTagList []RunInstancesTag

func (list *RunInstancesTagList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]RunInstancesTag)
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

type RunInstancesNetworkInterfaceList []RunInstancesNetworkInterface

func (list *RunInstancesNetworkInterfaceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]RunInstancesNetworkInterface)
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

type RunInstancesDataDiskList []RunInstancesDataDisk

func (list *RunInstancesDataDiskList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]RunInstancesDataDisk)
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

type RunInstancesInstanceIdSetList []string

func (list *RunInstancesInstanceIdSetList) UnmarshalJSON(data []byte) error {
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

func (c *EcsClient) ModifyImageShareGroupPermission(req *ModifyImageShareGroupPermissionArgs) (resp *ModifyImageShareGroupPermissionResponse, err error) {
	resp = &ModifyImageShareGroupPermissionResponse{}
	err = c.Request("ModifyImageShareGroupPermission", req, resp)
	return
}

type ModifyImageShareGroupPermissionArgs struct {
	ResourceOwnerId      int64
	ImageId              string
	AddGroup1            string
	ResourceOwnerAccount string
	OwnerAccount         string
	RemoveGroup1         string
	OwnerId              int64
}
type ModifyImageShareGroupPermissionResponse struct {
	RequestId string
}

func (c *EcsClient) ModifyVRouterAttribute(req *ModifyVRouterAttributeArgs) (resp *ModifyVRouterAttributeResponse, err error) {
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

func (c *EcsClient) DescribeIpRanges(req *DescribeIpRangesArgs) (resp *DescribeIpRangesResponse, err error) {
	resp = &DescribeIpRangesResponse{}
	err = c.Request("DescribeIpRanges", req, resp)
	return
}

type DescribeIpRangesIpRange struct {
	IpAddress string
	NicType   string
}
type DescribeIpRangesArgs struct {
	NicType              string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	PageSize             int
	ClusterId            string
	OwnerId              int64
	PageNumber           int
}
type DescribeIpRangesResponse struct {
	RequestId  string
	RegionId   string
	ClusterId  string
	TotalCount int
	PageNumber int
	PageSize   int
	IpRanges   DescribeIpRangesIpRangeList
}

type DescribeIpRangesIpRangeList []DescribeIpRangesIpRange

func (list *DescribeIpRangesIpRangeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeIpRangesIpRange)
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

func (c *EcsClient) DeleteVSwitch(req *DeleteVSwitchArgs) (resp *DeleteVSwitchResponse, err error) {
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

func (c *EcsClient) ModifyVSwitchAttribute(req *ModifyVSwitchAttributeArgs) (resp *ModifyVSwitchAttributeResponse, err error) {
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

func (c *EcsClient) CheckAutoSnapshotPolicy(req *CheckAutoSnapshotPolicyArgs) (resp *CheckAutoSnapshotPolicyResponse, err error) {
	resp = &CheckAutoSnapshotPolicyResponse{}
	err = c.Request("CheckAutoSnapshotPolicy", req, resp)
	return
}

type CheckAutoSnapshotPolicyArgs struct {
	DataDiskPolicyEnabled             bool
	ResourceOwnerId                   int64
	DataDiskPolicyRetentionDays       int
	ResourceOwnerAccount              string
	SystemDiskPolicyRetentionLastWeek bool
	OwnerAccount                      string
	SystemDiskPolicyTimePeriod        int
	OwnerId                           int64
	DataDiskPolicyRetentionLastWeek   bool
	SystemDiskPolicyRetentionDays     int
	DataDiskPolicyTimePeriod          int
	SystemDiskPolicyEnabled           bool
}
type CheckAutoSnapshotPolicyResponse struct {
	RequestId              string
	AutoSnapshotOccupation int
	IsPermittedModify      string
}

func (c *EcsClient) DescribeInstanceVncUrl(req *DescribeInstanceVncUrlArgs) (resp *DescribeInstanceVncUrlResponse, err error) {
	resp = &DescribeInstanceVncUrlResponse{}
	err = c.Request("DescribeInstanceVncUrl", req, resp)
	return
}

type DescribeInstanceVncUrlArgs struct {
	ResourceOwnerId      int64
	InstanceId           string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
}
type DescribeInstanceVncUrlResponse struct {
	RequestId string
	VncUrl    string
}

func (c *EcsClient) CreateVSwitch(req *CreateVSwitchArgs) (resp *CreateVSwitchResponse, err error) {
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

func (c *EcsClient) ModifyInstanceVncPasswd(req *ModifyInstanceVncPasswdArgs) (resp *ModifyInstanceVncPasswdResponse, err error) {
	resp = &ModifyInstanceVncPasswdResponse{}
	err = c.Request("ModifyInstanceVncPasswd", req, resp)
	return
}

type ModifyInstanceVncPasswdArgs struct {
	ResourceOwnerId      int64
	InstanceId           string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	VncPassword          string
}
type ModifyInstanceVncPasswdResponse struct {
	RequestId string
}

func (c *EcsClient) DetachDisk(req *DetachDiskArgs) (resp *DetachDiskResponse, err error) {
	resp = &DetachDiskResponse{}
	err = c.Request("DetachDisk", req, resp)
	return
}

type DetachDiskArgs struct {
	ResourceOwnerId      int64
	InstanceId           string
	ResourceOwnerAccount string
	OwnerAccount         string
	DiskId               string
	OwnerId              int64
}
type DetachDiskResponse struct {
	RequestId string
}

func (c *EcsClient) ConvertNatPublicIpToEip(req *ConvertNatPublicIpToEipArgs) (resp *ConvertNatPublicIpToEipResponse, err error) {
	resp = &ConvertNatPublicIpToEipResponse{}
	err = c.Request("ConvertNatPublicIpToEip", req, resp)
	return
}

type ConvertNatPublicIpToEipArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerId              int64
	InstanceId           string
}
type ConvertNatPublicIpToEipResponse struct {
	RequestId string
}

func (c *EcsClient) DescribeInstanceHistoryEvents(req *DescribeInstanceHistoryEventsArgs) (resp *DescribeInstanceHistoryEventsResponse, err error) {
	resp = &DescribeInstanceHistoryEventsResponse{}
	err = c.Request("DescribeInstanceHistoryEvents", req, resp)
	return
}

type DescribeInstanceHistoryEventsInstanceSystemEventType struct {
	InstanceId       string
	EventId          string
	EventPublishTime string
	NotBefore        string
	EventType        DescribeInstanceHistoryEventsEventType
}

type DescribeInstanceHistoryEventsEventType struct {
	Code int
	Name string
}
type DescribeInstanceHistoryEventsArgs struct {
	EventIds              DescribeInstanceHistoryEventsEventIdList
	ResourceOwnerId       int64
	EventCycleStatus      string
	PageNumber            int
	PageSize              int
	EventPublishTimeEnd   string
	ResourceOwnerAccount  string
	OwnerAccount          string
	NotBeforeStart        string
	OwnerId               int64
	EventPublishTimeStart string
	InstanceId            string
	NotBeforeEnd          string
	EventType             string
}
type DescribeInstanceHistoryEventsResponse struct {
	RequestId              string
	TotalCount             int
	PageNumber             int
	PageSize               int
	InstanceSystemEventSet DescribeInstanceHistoryEventsInstanceSystemEventTypeList
}

type DescribeInstanceHistoryEventsEventIdList []string

func (list *DescribeInstanceHistoryEventsEventIdList) UnmarshalJSON(data []byte) error {
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

type DescribeInstanceHistoryEventsInstanceSystemEventTypeList []DescribeInstanceHistoryEventsInstanceSystemEventType

func (list *DescribeInstanceHistoryEventsInstanceSystemEventTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInstanceHistoryEventsInstanceSystemEventType)
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

func (c *EcsClient) ModifyCommand(req *ModifyCommandArgs) (resp *ModifyCommandResponse, err error) {
	resp = &ModifyCommandResponse{}
	err = c.Request("ModifyCommand", req, resp)
	return
}

type ModifyCommandArgs struct {
	ResourceOwnerId      int64
	WorkingDir           string
	Description          string
	CommandId            string
	CommandContent       string
	Timeout              int64
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	Name                 string
}
type ModifyCommandResponse struct {
	RequestId string
}

func (c *EcsClient) DescribeUserData(req *DescribeUserDataArgs) (resp *DescribeUserDataResponse, err error) {
	resp = &DescribeUserDataResponse{}
	err = c.Request("DescribeUserData", req, resp)
	return
}

type DescribeUserDataArgs struct {
	ResourceOwnerId      int64
	InstanceId           string
	ResourceOwnerAccount string
	OwnerId              int64
}
type DescribeUserDataResponse struct {
	RequestId  string
	RegionId   string
	InstanceId string
	UserData   string
}

func (c *EcsClient) AttachDisk(req *AttachDiskArgs) (resp *AttachDiskResponse, err error) {
	resp = &AttachDiskResponse{}
	err = c.Request("AttachDisk", req, resp)
	return
}

type AttachDiskArgs struct {
	ResourceOwnerId      int64
	InstanceId           string
	ResourceOwnerAccount string
	OwnerAccount         string
	DiskId               string
	OwnerId              int64
	Device               string
	DeleteWithInstance   bool
}
type AttachDiskResponse struct {
	RequestId string
}

func (c *EcsClient) DescribeInstances(req *DescribeInstancesArgs) (resp *DescribeInstancesResponse, err error) {
	resp = &DescribeInstancesResponse{}
	err = c.Request("DescribeInstances", req, resp)
	return
}

type DescribeInstancesInstance struct {
	InstanceId              string
	InstanceName            string
	Description             string
	ImageId                 string
	OSName                  string
	OSType                  string
	RegionId                string
	ZoneId                  string
	ClusterId               string
	InstanceType            string
	Cpu                     int
	Memory                  int
	HostName                string
	Status                  string
	SerialNumber            string
	InternetChargeType      string
	InternetMaxBandwidthIn  int
	InternetMaxBandwidthOut int
	VlanId                  string
	CreationTime            string
	InstanceNetworkType     string
	InstanceChargeType      string
	SaleCycle               string
	ExpiredTime             string
	AutoReleaseTime         string
	IoOptimized             bool
	DeviceAvailable         bool
	InstanceTypeFamily      string
	LocalStorageCapacity    int64
	LocalStorageAmount      int
	GPUAmount               int
	GPUSpec                 string
	SpotStrategy            string
	SpotPriceLimit          float32
	ResourceGroupId         string
	KeyPairName             string
	Recyclable              bool
	HpcClusterId            string
	StoppedMode             string
	NetworkInterfaces       DescribeInstancesNetworkInterfaceList
	OperationLocks          DescribeInstancesLockReasonList
	Tags                    DescribeInstancesTagList
	SecurityGroupIds        DescribeInstancesSecurityGroupIdList
	PublicIpAddress         DescribeInstancesPublicIpAddresList
	InnerIpAddress          DescribeInstancesInnerIpAddresList
	RdmaIpAddress           DescribeInstancesRdmaIpAddresList
	VpcAttributes           DescribeInstancesVpcAttributes
	EipAddress              DescribeInstancesEipAddress
}

type DescribeInstancesNetworkInterface struct {
	NetworkInterfaceId string
	MacAddress         string
	PrimaryIpAddress   string
}

type DescribeInstancesLockReason struct {
	LockReason string
	LockMsg    string
}

type DescribeInstancesTag struct {
	TagKey   string
	TagValue string
}

type DescribeInstancesVpcAttributes struct {
	VpcId            string
	VSwitchId        string
	NatIpAddress     string
	PrivateIpAddress DescribeInstancesPrivateIpAddresList
}

type DescribeInstancesEipAddress struct {
	AllocationId         string
	IpAddress            string
	Bandwidth            int
	InternetChargeType   string
	IsSupportUnassociate bool
}
type DescribeInstancesArgs struct {
	Tag4Value            string
	InnerIpAddresses     string
	ResourceOwnerId      int64
	Tag2Key              string
	PrivateIpAddresses   string
	HpcClusterId         string
	Filter2Value         string
	Tag3Key              string
	KeyPairName          string
	Tag1Value            string
	ResourceGroupId      string
	LockReason           string
	Filter1Key           string
	DeviceAvailable      bool
	Filter3Value         string
	Tag5Key              string
	Filter1Value         string
	OwnerId              int64
	VSwitchId            string
	InstanceName         string
	InstanceIds          string
	InternetChargeType   string
	ZoneId               string
	Tag4Key              string
	InstanceNetworkType  string
	Status               string
	ImageId              string
	Filter4Value         string
	IoOptimized          bool
	SecurityGroupId      string
	Filter4Key           string
	PageNumber           int
	RdmaIpAddresses      string
	PageSize             int
	PublicIpAddresses    string
	InstanceType         string
	InstanceChargeType   string
	Tag3Value            string
	ResourceOwnerAccount string
	OwnerAccount         string
	InstanceTypeFamily   string
	Filter2Key           string
	Tag5Value            string
	Tag1Key              string
	EipAddresses         string
	VpcId                string
	Tag2Value            string
	Filter3Key           string
}
type DescribeInstancesResponse struct {
	RequestId  string
	TotalCount int
	PageNumber int
	PageSize   int
	Instances  DescribeInstancesInstanceList
}

type DescribeInstancesNetworkInterfaceList []DescribeInstancesNetworkInterface

func (list *DescribeInstancesNetworkInterfaceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInstancesNetworkInterface)
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

type DescribeInstancesLockReasonList []DescribeInstancesLockReason

func (list *DescribeInstancesLockReasonList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInstancesLockReason)
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

type DescribeInstancesTagList []DescribeInstancesTag

func (list *DescribeInstancesTagList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInstancesTag)
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

type DescribeInstancesSecurityGroupIdList []string

func (list *DescribeInstancesSecurityGroupIdList) UnmarshalJSON(data []byte) error {
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

type DescribeInstancesPublicIpAddresList []string

func (list *DescribeInstancesPublicIpAddresList) UnmarshalJSON(data []byte) error {
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

type DescribeInstancesInnerIpAddresList []string

func (list *DescribeInstancesInnerIpAddresList) UnmarshalJSON(data []byte) error {
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

type DescribeInstancesRdmaIpAddresList []string

func (list *DescribeInstancesRdmaIpAddresList) UnmarshalJSON(data []byte) error {
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

type DescribeInstancesPrivateIpAddresList []string

func (list *DescribeInstancesPrivateIpAddresList) UnmarshalJSON(data []byte) error {
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

type DescribeInstancesInstanceList []DescribeInstancesInstance

func (list *DescribeInstancesInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInstancesInstance)
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

func (c *EcsClient) DeleteNatGateway(req *DeleteNatGatewayArgs) (resp *DeleteNatGatewayResponse, err error) {
	resp = &DeleteNatGatewayResponse{}
	err = c.Request("DeleteNatGateway", req, resp)
	return
}

type DeleteNatGatewayArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	NatGatewayId         string
	OwnerId              int64
}
type DeleteNatGatewayResponse struct {
	RequestId string
}

func (c *EcsClient) DescribeInstanceAutoRenewAttribute(req *DescribeInstanceAutoRenewAttributeArgs) (resp *DescribeInstanceAutoRenewAttributeResponse, err error) {
	resp = &DescribeInstanceAutoRenewAttributeResponse{}
	err = c.Request("DescribeInstanceAutoRenewAttribute", req, resp)
	return
}

type DescribeInstanceAutoRenewAttributeInstanceRenewAttribute struct {
	InstanceId       string
	AutoRenewEnabled bool
	Duration         int
	PeriodUnit       string
	RenewalStatus    string
}
type DescribeInstanceAutoRenewAttributeArgs struct {
	ResourceOwnerId      int64
	InstanceId           string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
}
type DescribeInstanceAutoRenewAttributeResponse struct {
	RequestId               string
	InstanceRenewAttributes DescribeInstanceAutoRenewAttributeInstanceRenewAttributeList
}

type DescribeInstanceAutoRenewAttributeInstanceRenewAttributeList []DescribeInstanceAutoRenewAttributeInstanceRenewAttribute

func (list *DescribeInstanceAutoRenewAttributeInstanceRenewAttributeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInstanceAutoRenewAttributeInstanceRenewAttribute)
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

func (c *EcsClient) InvokeCommand(req *InvokeCommandArgs) (resp *InvokeCommandResponse, err error) {
	resp = &InvokeCommandResponse{}
	err = c.Request("InvokeCommand", req, resp)
	return
}

type InvokeCommandArgs struct {
	ResourceOwnerId      int64
	CommandId            string
	Frequency            string
	Timed                bool
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	InstanceIds          InvokeCommandInstanceIdList
}
type InvokeCommandResponse struct {
	RequestId string
	InvokeId  string
}

type InvokeCommandInstanceIdList []string

func (list *InvokeCommandInstanceIdList) UnmarshalJSON(data []byte) error {
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

func (c *EcsClient) DescribeInstanceVncPasswd(req *DescribeInstanceVncPasswdArgs) (resp *DescribeInstanceVncPasswdResponse, err error) {
	resp = &DescribeInstanceVncPasswdResponse{}
	err = c.Request("DescribeInstanceVncPasswd", req, resp)
	return
}

type DescribeInstanceVncPasswdArgs struct {
	ResourceOwnerId      int64
	InstanceId           string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
}
type DescribeInstanceVncPasswdResponse struct {
	RequestId string
	VncPasswd string
}

func (c *EcsClient) DescribeUserBusinessBehavior(req *DescribeUserBusinessBehaviorArgs) (resp *DescribeUserBusinessBehaviorResponse, err error) {
	resp = &DescribeUserBusinessBehaviorResponse{}
	err = c.Request("DescribeUserBusinessBehavior", req, resp)
	return
}

type DescribeUserBusinessBehaviorArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	StatusKey            string
}
type DescribeUserBusinessBehaviorResponse struct {
	RequestId   string
	StatusValue string
}

func (c *EcsClient) DescribeAutoSnapshotPolicyEx(req *DescribeAutoSnapshotPolicyExArgs) (resp *DescribeAutoSnapshotPolicyExResponse, err error) {
	resp = &DescribeAutoSnapshotPolicyExResponse{}
	err = c.Request("DescribeAutoSnapshotPolicyEx", req, resp)
	return
}

type DescribeAutoSnapshotPolicyExAutoSnapshotPolicy struct {
	AutoSnapshotPolicyId   string
	RegionId               string
	AutoSnapshotPolicyName string
	TimePoints             string
	RepeatWeekdays         string
	RetentionDays          int
	DiskNums               int
	VolumeNums             int
	CreationTime           string
	Status                 string
}
type DescribeAutoSnapshotPolicyExArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	AutoSnapshotPolicyId string
	OwnerAccount         string
	PageSize             int
	OwnerId              int64
	PageNumber           int
}
type DescribeAutoSnapshotPolicyExResponse struct {
	RequestId            string
	TotalCount           int
	PageNumber           int
	PageSize             int
	AutoSnapshotPolicies DescribeAutoSnapshotPolicyExAutoSnapshotPolicyList
}

type DescribeAutoSnapshotPolicyExAutoSnapshotPolicyList []DescribeAutoSnapshotPolicyExAutoSnapshotPolicy

func (list *DescribeAutoSnapshotPolicyExAutoSnapshotPolicyList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeAutoSnapshotPolicyExAutoSnapshotPolicy)
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

func (c *EcsClient) ModifyForwardEntry(req *ModifyForwardEntryArgs) (resp *ModifyForwardEntryResponse, err error) {
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

func (c *EcsClient) ModifyRouterInterfaceAttribute(req *ModifyRouterInterfaceAttributeArgs) (resp *ModifyRouterInterfaceAttributeResponse, err error) {
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

func (c *EcsClient) DeleteHaVip(req *DeleteHaVipArgs) (resp *DeleteHaVipResponse, err error) {
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

func (c *EcsClient) AuthorizeSecurityGroupEgress(req *AuthorizeSecurityGroupEgressArgs) (resp *AuthorizeSecurityGroupEgressResponse, err error) {
	resp = &AuthorizeSecurityGroupEgressResponse{}
	err = c.Request("AuthorizeSecurityGroupEgress", req, resp)
	return
}

type AuthorizeSecurityGroupEgressArgs struct {
	NicType               string
	ResourceOwnerId       int64
	SourcePortRange       string
	ClientToken           string
	SecurityGroupId       string
	Description           string
	Policy                string
	PortRange             string
	ResourceOwnerAccount  string
	IpProtocol            string
	OwnerAccount          string
	SourceCidrIp          string
	DestGroupId           string
	OwnerId               int64
	DestGroupOwnerAccount string
	Priority              string
	DestCidrIp            string
	DestGroupOwnerId      int64
}
type AuthorizeSecurityGroupEgressResponse struct {
	RequestId string
}

func (c *EcsClient) ModifyPrepayInstanceSpec(req *ModifyPrepayInstanceSpecArgs) (resp *ModifyPrepayInstanceSpecResponse, err error) {
	resp = &ModifyPrepayInstanceSpecResponse{}
	err = c.Request("ModifyPrepayInstanceSpec", req, resp)
	return
}

type ModifyPrepayInstanceSpecArgs struct {
	ResourceOwnerId      int64
	InstanceId           string
	AutoPay              bool
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	InstanceType         string
	OwnerId              int64
}
type ModifyPrepayInstanceSpecResponse struct {
	RequestId string
	OrderId   string
}

func (c *EcsClient) AllocatePublicIpAddress(req *AllocatePublicIpAddressArgs) (resp *AllocatePublicIpAddressResponse, err error) {
	resp = &AllocatePublicIpAddressResponse{}
	err = c.Request("AllocatePublicIpAddress", req, resp)
	return
}

type AllocatePublicIpAddressArgs struct {
	IpAddress            string
	ResourceOwnerId      int64
	InstanceId           string
	ResourceOwnerAccount string
	VlanId               string
	OwnerAccount         string
	OwnerId              int64
}
type AllocatePublicIpAddressResponse struct {
	RequestId string
	IpAddress string
}

func (c *EcsClient) DescribeImageSupportInstanceTypes(req *DescribeImageSupportInstanceTypesArgs) (resp *DescribeImageSupportInstanceTypesResponse, err error) {
	resp = &DescribeImageSupportInstanceTypesResponse{}
	err = c.Request("DescribeImageSupportInstanceTypes", req, resp)
	return
}

type DescribeImageSupportInstanceTypesInstanceType struct {
	InstanceTypeId     string
	CpuCoreCount       int
	MemorySize         float32
	InstanceTypeFamily string
}
type DescribeImageSupportInstanceTypesArgs struct {
	ResourceOwnerId      int64
	ImageId              string
	ResourceOwnerAccount string
	OwnerId              int64
}
type DescribeImageSupportInstanceTypesResponse struct {
	RequestId     string
	RegionId      string
	ImageId       string
	InstanceTypes DescribeImageSupportInstanceTypesInstanceTypeList
}

type DescribeImageSupportInstanceTypesInstanceTypeList []DescribeImageSupportInstanceTypesInstanceType

func (list *DescribeImageSupportInstanceTypesInstanceTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeImageSupportInstanceTypesInstanceType)
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

func (c *EcsClient) AttachNetworkInterface(req *AttachNetworkInterfaceArgs) (resp *AttachNetworkInterfaceResponse, err error) {
	resp = &AttachNetworkInterfaceResponse{}
	err = c.Request("AttachNetworkInterface", req, resp)
	return
}

type AttachNetworkInterfaceArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	InstanceId           string
	NetworkInterfaceId   string
}
type AttachNetworkInterfaceResponse struct {
	RequestId string
}

func (c *EcsClient) ModifyNetworkInterfaceAttribute(req *ModifyNetworkInterfaceAttributeArgs) (resp *ModifyNetworkInterfaceAttributeResponse, err error) {
	resp = &ModifyNetworkInterfaceAttributeResponse{}
	err = c.Request("ModifyNetworkInterfaceAttribute", req, resp)
	return
}

type ModifyNetworkInterfaceAttributeArgs struct {
	ResourceOwnerId      int64
	SecurityGroupIds     ModifyNetworkInterfaceAttributeSecurityGroupIdList
	Description          string
	NetworkInterfaceName string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	NetworkInterfaceId   string
}
type ModifyNetworkInterfaceAttributeResponse struct {
	RequestId string
}

type ModifyNetworkInterfaceAttributeSecurityGroupIdList []string

func (list *ModifyNetworkInterfaceAttributeSecurityGroupIdList) UnmarshalJSON(data []byte) error {
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

func (c *EcsClient) DescribeClassicLinkInstances(req *DescribeClassicLinkInstancesArgs) (resp *DescribeClassicLinkInstancesResponse, err error) {
	resp = &DescribeClassicLinkInstancesResponse{}
	err = c.Request("DescribeClassicLinkInstances", req, resp)
	return
}

type DescribeClassicLinkInstancesLink struct {
	InstanceId string
	VpcId      string
}
type DescribeClassicLinkInstancesArgs struct {
	ResourceOwnerId      int64
	InstanceId           string
	ResourceOwnerAccount string
	VpcId                string
	PageSize             string
	OwnerId              int64
	PageNumber           string
}
type DescribeClassicLinkInstancesResponse struct {
	RequestId  string
	TotalCount int
	PageNumber int
	PageSize   int
	Links      DescribeClassicLinkInstancesLinkList
}

type DescribeClassicLinkInstancesLinkList []DescribeClassicLinkInstancesLink

func (list *DescribeClassicLinkInstancesLinkList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClassicLinkInstancesLink)
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

func (c *EcsClient) ModifyDiskChargeType(req *ModifyDiskChargeTypeArgs) (resp *ModifyDiskChargeTypeResponse, err error) {
	resp = &ModifyDiskChargeTypeResponse{}
	err = c.Request("ModifyDiskChargeType", req, resp)
	return
}

type ModifyDiskChargeTypeArgs struct {
	ResourceOwnerId      int64
	InstanceId           string
	AutoPay              bool
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	DiskIds              string
	OwnerId              int64
}
type ModifyDiskChargeTypeResponse struct {
	RequestId string
	OrderId   string
}

func (c *EcsClient) CreateNetworkInterface(req *CreateNetworkInterfaceArgs) (resp *CreateNetworkInterfaceResponse, err error) {
	resp = &CreateNetworkInterfaceResponse{}
	err = c.Request("CreateNetworkInterface", req, resp)
	return
}

type CreateNetworkInterfaceArgs struct {
	ResourceOwnerId      int64
	ClientToken          string
	SecurityGroupId      string
	Description          string
	NetworkInterfaceName string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	VSwitchId            string
	PrimaryIpAddress     string
}
type CreateNetworkInterfaceResponse struct {
	RequestId          string
	NetworkInterfaceId string
}

func (c *EcsClient) DescribeDeploymentSetTopology(req *DescribeDeploymentSetTopologyArgs) (resp *DescribeDeploymentSetTopologyResponse, err error) {
	resp = &DescribeDeploymentSetTopologyResponse{}
	err = c.Request("DescribeDeploymentSetTopology", req, resp)
	return
}

type DescribeDeploymentSetTopology_Switch struct {
	SwitchId string
	Hosts    DescribeDeploymentSetTopologyHostList
}

type DescribeDeploymentSetTopologyHost struct {
	HostId      string
	InstanceIds DescribeDeploymentSetTopologyInstanceIdList
}

type DescribeDeploymentSetTopologyRack struct {
	RackId string
	Hosts1 DescribeDeploymentSetTopologyHost2List
}

type DescribeDeploymentSetTopologyHost2 struct {
	HostId       string
	InstanceIds3 DescribeDeploymentSetTopologyInstanceIds3List
}
type DescribeDeploymentSetTopologyArgs struct {
	DeploymentSetId      string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	Granularity          string
	Domain               string
	NetworkType          string
	DeploymentSetName    string
	OwnerId              int64
	Strategy             string
}
type DescribeDeploymentSetTopologyResponse struct {
	RequestId string
	Switchs   DescribeDeploymentSetTopology_SwitchList
	Racks     DescribeDeploymentSetTopologyRackList
}

type DescribeDeploymentSetTopologyHostList []DescribeDeploymentSetTopologyHost

func (list *DescribeDeploymentSetTopologyHostList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDeploymentSetTopologyHost)
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

type DescribeDeploymentSetTopologyInstanceIdList []string

func (list *DescribeDeploymentSetTopologyInstanceIdList) UnmarshalJSON(data []byte) error {
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

type DescribeDeploymentSetTopologyHost2List []DescribeDeploymentSetTopologyHost2

func (list *DescribeDeploymentSetTopologyHost2List) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDeploymentSetTopologyHost2)
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

type DescribeDeploymentSetTopologyInstanceIds3List []string

func (list *DescribeDeploymentSetTopologyInstanceIds3List) UnmarshalJSON(data []byte) error {
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

type DescribeDeploymentSetTopology_SwitchList []DescribeDeploymentSetTopology_Switch

func (list *DescribeDeploymentSetTopology_SwitchList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDeploymentSetTopology_Switch)
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

type DescribeDeploymentSetTopologyRackList []DescribeDeploymentSetTopologyRack

func (list *DescribeDeploymentSetTopologyRackList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDeploymentSetTopologyRack)
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

func (c *EcsClient) DetachNetworkInterface(req *DetachNetworkInterfaceArgs) (resp *DetachNetworkInterfaceResponse, err error) {
	resp = &DetachNetworkInterfaceResponse{}
	err = c.Request("DetachNetworkInterface", req, resp)
	return
}

type DetachNetworkInterfaceArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	InstanceId           string
	NetworkInterfaceId   string
}
type DetachNetworkInterfaceResponse struct {
	RequestId string
}

func (c *EcsClient) ModifyInstanceAutoReleaseTime(req *ModifyInstanceAutoReleaseTimeArgs) (resp *ModifyInstanceAutoReleaseTimeResponse, err error) {
	resp = &ModifyInstanceAutoReleaseTimeResponse{}
	err = c.Request("ModifyInstanceAutoReleaseTime", req, resp)
	return
}

type ModifyInstanceAutoReleaseTimeArgs struct {
	ResourceOwnerId      int64
	InstanceId           string
	ResourceOwnerAccount string
	OwnerAccount         string
	AutoReleaseTime      string
	OwnerId              int64
}
type ModifyInstanceAutoReleaseTimeResponse struct {
	RequestId string
}

func (c *EcsClient) ModifyInstanceVpcAttribute(req *ModifyInstanceVpcAttributeArgs) (resp *ModifyInstanceVpcAttributeResponse, err error) {
	resp = &ModifyInstanceVpcAttributeResponse{}
	err = c.Request("ModifyInstanceVpcAttribute", req, resp)
	return
}

type ModifyInstanceVpcAttributeArgs struct {
	VSwitchId            string
	PrivateIpAddress     string
	ResourceOwnerId      int64
	InstanceId           string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
}
type ModifyInstanceVpcAttributeResponse struct {
	RequestId string
}

func (c *EcsClient) RebootInstance(req *RebootInstanceArgs) (resp *RebootInstanceResponse, err error) {
	resp = &RebootInstanceResponse{}
	err = c.Request("RebootInstance", req, resp)
	return
}

type RebootInstanceArgs struct {
	ResourceOwnerId      int64
	InstanceId           string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	ForceStop            bool
}
type RebootInstanceResponse struct {
	RequestId string
}

func (c *EcsClient) DescribeVirtualBorderRouters(req *DescribeVirtualBorderRoutersArgs) (resp *DescribeVirtualBorderRoutersResponse, err error) {
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

func (c *EcsClient) AddBandwidthPackageIps(req *AddBandwidthPackageIpsArgs) (resp *AddBandwidthPackageIpsResponse, err error) {
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

func (c *EcsClient) DescribeResourceByTags(req *DescribeResourceByTagsArgs) (resp *DescribeResourceByTagsResponse, err error) {
	resp = &DescribeResourceByTagsResponse{}
	err = c.Request("DescribeResourceByTags", req, resp)
	return
}

type DescribeResourceByTagsResource struct {
	ResourceId   string
	ResourceType string
	RegionId     string
}
type DescribeResourceByTagsArgs struct {
	Tag4Value            string
	ResourceOwnerId      int64
	Tag2Key              string
	Tag5Key              string
	ResourceOwnerAccount string
	Tag3Key              string
	OwnerId              int64
	ResourceType         string
	Tag5Value            string
	PageNumber           int
	Tag1Key              string
	Tag1Value            string
	Tag2Value            string
	PageSize             int
	Tag4Key              string
	Tag3Value            string
}
type DescribeResourceByTagsResponse struct {
	RequestId  string
	PageSize   int
	PageNumber int
	TotalCount int
	Resources  DescribeResourceByTagsResourceList
}

type DescribeResourceByTagsResourceList []DescribeResourceByTagsResource

func (list *DescribeResourceByTagsResourceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeResourceByTagsResource)
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

func (c *EcsClient) ModifyInstanceAutoRenewAttribute(req *ModifyInstanceAutoRenewAttributeArgs) (resp *ModifyInstanceAutoRenewAttributeResponse, err error) {
	resp = &ModifyInstanceAutoRenewAttributeResponse{}
	err = c.Request("ModifyInstanceAutoRenewAttribute", req, resp)
	return
}

type ModifyInstanceAutoRenewAttributeArgs struct {
	Duration             int
	ResourceOwnerId      int64
	InstanceId           string
	AutoRenew            bool
	ResourceOwnerAccount string
	OwnerAccount         string
	RenewalStatus        string
	OwnerId              int64
}
type ModifyInstanceAutoRenewAttributeResponse struct {
	RequestId string
}

func (c *EcsClient) DescribeDisks(req *DescribeDisksArgs) (resp *DescribeDisksResponse, err error) {
	resp = &DescribeDisksResponse{}
	err = c.Request("DescribeDisks", req, resp)
	return
}

type DescribeDisksDisk struct {
	DiskId                        string
	RegionId                      string
	ZoneId                        string
	DiskName                      string
	Description                   string
	Type                          string
	Category                      string
	Size                          int
	ImageId                       string
	SourceSnapshotId              string
	AutoSnapshotPolicyId          string
	ProductCode                   string
	Portable                      bool
	Status                        string
	InstanceId                    string
	Device                        string
	DeleteWithInstance            bool
	DeleteAutoSnapshot            bool
	EnableAutoSnapshot            bool
	EnableAutomatedSnapshotPolicy bool
	CreationTime                  string
	AttachedTime                  string
	DetachedTime                  string
	DiskChargeType                string
	ExpiredTime                   string
	ResourceGroupId               string
	Encrypted                     bool
	IOPS                          int
	IOPSRead                      int
	IOPSWrite                     int
	OperationLocks                DescribeDisksOperationLockList
	Tags                          DescribeDisksTagList
}

type DescribeDisksOperationLock struct {
	LockReason string
}

type DescribeDisksTag struct {
	TagKey   string
	TagValue string
}
type DescribeDisksArgs struct {
	Tag4Value                     string
	ResourceOwnerId               int64
	SnapshotId                    string
	Tag2Key                       string
	Filter2Value                  string
	AutoSnapshotPolicyId          string
	Tag3Key                       string
	PageNumber                    int
	DiskName                      string
	Tag1Value                     string
	DeleteAutoSnapshot            bool
	ResourceGroupId               string
	DiskChargeType                string
	LockReason                    string
	Filter1Key                    string
	PageSize                      int
	DiskIds                       string
	DeleteWithInstance            bool
	Tag3Value                     string
	EnableAutoSnapshot            bool
	Tag5Key                       string
	ResourceOwnerAccount          string
	OwnerAccount                  string
	Filter1Value                  string
	Portable                      bool
	EnableAutomatedSnapshotPolicy bool
	Filter2Key                    string
	OwnerId                       int64
	DiskType                      string
	Tag5Value                     string
	Tag1Key                       string
	AdditionalAttributess         DescribeDisksAdditionalAttributesList
	EnableShared                  bool
	InstanceId                    string
	Encrypted                     bool
	Tag2Value                     string
	ZoneId                        string
	Tag4Key                       string
	Category                      string
	Status                        string
}
type DescribeDisksResponse struct {
	RequestId  string
	TotalCount int
	PageNumber int
	PageSize   int
	Disks      DescribeDisksDiskList
}

type DescribeDisksOperationLockList []DescribeDisksOperationLock

func (list *DescribeDisksOperationLockList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDisksOperationLock)
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

type DescribeDisksTagList []DescribeDisksTag

func (list *DescribeDisksTagList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDisksTag)
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

type DescribeDisksAdditionalAttributesList []int64

func (list *DescribeDisksAdditionalAttributesList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]int64)
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

type DescribeDisksDiskList []DescribeDisksDisk

func (list *DescribeDisksDiskList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDisksDisk)
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

func (c *EcsClient) DeleteImage(req *DeleteImageArgs) (resp *DeleteImageResponse, err error) {
	resp = &DeleteImageResponse{}
	err = c.Request("DeleteImage", req, resp)
	return
}

type DeleteImageArgs struct {
	ResourceOwnerId      int64
	ImageId              string
	ResourceOwnerAccount string
	OwnerAccount         string
	Force                bool
	OwnerId              int64
}
type DeleteImageResponse struct {
	RequestId string
}

func (c *EcsClient) DescribeDeploymentSets(req *DescribeDeploymentSetsArgs) (resp *DescribeDeploymentSetsResponse, err error) {
	resp = &DescribeDeploymentSetsResponse{}
	err = c.Request("DescribeDeploymentSets", req, resp)
	return
}

type DescribeDeploymentSetsDeploymentSet struct {
	DeploymentSetId          string
	DeploymentSetDescription string
	DeploymentSetName        string
	Strategy                 string
	Domain                   string
	Granularity              string
	InstanceAmount           int
	CreationTime             string
}
type DescribeDeploymentSetsArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	NetworkType          string
	DeploymentSetName    string
	OwnerId              int64
	PageNumber           int
	DeploymentSetIds     string
	Granularity          string
	Domain               string
	PageSize             int
	Strategy             string
}
type DescribeDeploymentSetsResponse struct {
	RequestId      string
	RegionId       string
	TotalCount     int
	PageNumber     int
	PageSize       int
	DeploymentSets DescribeDeploymentSetsDeploymentSetList
}

type DescribeDeploymentSetsDeploymentSetList []DescribeDeploymentSetsDeploymentSet

func (list *DescribeDeploymentSetsDeploymentSetList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDeploymentSetsDeploymentSet)
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

func (c *EcsClient) StartInstance(req *StartInstanceArgs) (resp *StartInstanceResponse, err error) {
	resp = &StartInstanceResponse{}
	err = c.Request("StartInstance", req, resp)
	return
}

type StartInstanceArgs struct {
	InitLocalDisk        bool
	ResourceOwnerId      int64
	InstanceId           string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
}
type StartInstanceResponse struct {
	RequestId string
}

func (c *EcsClient) DescribeSnapshotLinks(req *DescribeSnapshotLinksArgs) (resp *DescribeSnapshotLinksResponse, err error) {
	resp = &DescribeSnapshotLinksResponse{}
	err = c.Request("DescribeSnapshotLinks", req, resp)
	return
}

type DescribeSnapshotLinksSnapshotLink struct {
	SnapshotLinkId string
	RegionId       string
	InstanceId     string
	InstanceName   string
	SourceDiskId   string
	SourceDiskSize int
	SourceDiskType string
	TotalSize      int
	TotalCount     int
}
type DescribeSnapshotLinksArgs struct {
	ResourceOwnerId      int64
	InstanceId           string
	ResourceOwnerAccount string
	OwnerAccount         string
	PageSize             int
	DiskIds              string
	SnapshotLinkIds      string
	OwnerId              int64
	PageNumber           int
}
type DescribeSnapshotLinksResponse struct {
	RequestId     string
	TotalCount    int
	PageNumber    int
	PageSize      int
	SnapshotLinks DescribeSnapshotLinksSnapshotLinkList
}

type DescribeSnapshotLinksSnapshotLinkList []DescribeSnapshotLinksSnapshotLink

func (list *DescribeSnapshotLinksSnapshotLinkList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSnapshotLinksSnapshotLink)
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

func (c *EcsClient) RemoveTags(req *RemoveTagsArgs) (resp *RemoveTagsResponse, err error) {
	resp = &RemoveTagsResponse{}
	err = c.Request("RemoveTags", req, resp)
	return
}

type RemoveTagsArgs struct {
	Tag4Value            string
	ResourceOwnerId      int64
	ResourceId           string
	Tag2Key              string
	Tag5Key              string
	ResourceOwnerAccount string
	Tag3Key              string
	OwnerId              int64
	ResourceType         string
	Tag5Value            string
	Tag1Key              string
	Tag1Value            string
	Tag2Value            string
	Tag4Key              string
	Tag3Value            string
}
type RemoveTagsResponse struct {
	RequestId string
}

func (c *EcsClient) UnassociateEipAddress(req *UnassociateEipAddressArgs) (resp *UnassociateEipAddressResponse, err error) {
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

func (c *EcsClient) DescribeSnapshotMonitorData(req *DescribeSnapshotMonitorDataArgs) (resp *DescribeSnapshotMonitorDataResponse, err error) {
	resp = &DescribeSnapshotMonitorDataResponse{}
	err = c.Request("DescribeSnapshotMonitorData", req, resp)
	return
}

type DescribeSnapshotMonitorDataDataPoint struct {
	TimeStamp string
	Size      int64
}
type DescribeSnapshotMonitorDataArgs struct {
	ResourceOwnerId      int64
	Period               int
	ResourceOwnerAccount string
	OwnerAccount         string
	EndTime              string
	StartTime            string
	OwnerId              int64
}
type DescribeSnapshotMonitorDataResponse struct {
	RequestId   string
	MonitorData DescribeSnapshotMonitorDataDataPointList
}

type DescribeSnapshotMonitorDataDataPointList []DescribeSnapshotMonitorDataDataPoint

func (list *DescribeSnapshotMonitorDataDataPointList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSnapshotMonitorDataDataPoint)
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

func (c *EcsClient) RevokeSecurityGroupEgress(req *RevokeSecurityGroupEgressArgs) (resp *RevokeSecurityGroupEgressResponse, err error) {
	resp = &RevokeSecurityGroupEgressResponse{}
	err = c.Request("RevokeSecurityGroupEgress", req, resp)
	return
}

type RevokeSecurityGroupEgressArgs struct {
	NicType               string
	ResourceOwnerId       int64
	SourcePortRange       string
	ClientToken           string
	SecurityGroupId       string
	Description           string
	Policy                string
	PortRange             string
	ResourceOwnerAccount  string
	IpProtocol            string
	OwnerAccount          string
	SourceCidrIp          string
	DestGroupId           string
	OwnerId               int64
	DestGroupOwnerAccount string
	Priority              string
	DestCidrIp            string
	DestGroupOwnerId      int64
}
type RevokeSecurityGroupEgressResponse struct {
	RequestId string
}

func (c *EcsClient) BindIpRange(req *BindIpRangeArgs) (resp *BindIpRangeResponse, err error) {
	resp = &BindIpRangeResponse{}
	err = c.Request("BindIpRange", req, resp)
	return
}

type BindIpRangeArgs struct {
	IpAddress            string
	ResourceOwnerId      int64
	InstanceId           string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
}
type BindIpRangeResponse struct {
	RequestId string
}

func (c *EcsClient) DescribeBandwidthPackages(req *DescribeBandwidthPackagesArgs) (resp *DescribeBandwidthPackagesResponse, err error) {
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
	PublicIpAddresses  DescribeBandwidthPackagesPublicIpAddresseList
}

type DescribeBandwidthPackagesPublicIpAddresse struct {
	AllocationId string
	IpAddress    string
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

func (c *EcsClient) DescribeTasks(req *DescribeTasksArgs) (resp *DescribeTasksResponse, err error) {
	resp = &DescribeTasksResponse{}
	err = c.Request("DescribeTasks", req, resp)
	return
}

type DescribeTasksTask struct {
	TaskId        string
	TaskAction    string
	TaskStatus    string
	SupportCancel string
	CreationTime  string
	FinishedTime  string
}
type DescribeTasksArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	EndTime              string
	StartTime            string
	OwnerId              int64
	TaskIds              string
	PageNumber           int
	TaskStatus           string
	PageSize             int
	TaskAction           string
}
type DescribeTasksResponse struct {
	RequestId  string
	RegionId   string
	TotalCount int
	PageNumber int
	PageSize   int
	TaskSet    DescribeTasksTaskList
}

type DescribeTasksTaskList []DescribeTasksTask

func (list *DescribeTasksTaskList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeTasksTask)
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

func (c *EcsClient) DescribeInstanceRamRole(req *DescribeInstanceRamRoleArgs) (resp *DescribeInstanceRamRoleResponse, err error) {
	resp = &DescribeInstanceRamRoleResponse{}
	err = c.Request("DescribeInstanceRamRole", req, resp)
	return
}

type DescribeInstanceRamRoleInstanceRamRoleSet struct {
	InstanceId  string
	RamRoleName string
}
type DescribeInstanceRamRoleArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	InstanceIds          string
	PageSize             int
	RamRoleName          string
	OwnerId              int64
	PageNumber           int
}
type DescribeInstanceRamRoleResponse struct {
	RequestId           string
	RegionId            string
	TotalCount          int
	InstanceRamRoleSets DescribeInstanceRamRoleInstanceRamRoleSetList
}

type DescribeInstanceRamRoleInstanceRamRoleSetList []DescribeInstanceRamRoleInstanceRamRoleSet

func (list *DescribeInstanceRamRoleInstanceRamRoleSetList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInstanceRamRoleInstanceRamRoleSet)
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

func (c *EcsClient) DeleteRecycleBin(req *DeleteRecycleBinArgs) (resp *DeleteRecycleBinResponse, err error) {
	resp = &DeleteRecycleBinResponse{}
	err = c.Request("DeleteRecycleBin", req, resp)
	return
}

type DeleteRecycleBinArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	ResourceIds          string
}
type DeleteRecycleBinResponse struct {
	RequestId string
}

func (c *EcsClient) DeleteAutoSnapshotPolicy(req *DeleteAutoSnapshotPolicyArgs) (resp *DeleteAutoSnapshotPolicyResponse, err error) {
	resp = &DeleteAutoSnapshotPolicyResponse{}
	err = c.Request("DeleteAutoSnapshotPolicy", req, resp)
	return
}

type DeleteAutoSnapshotPolicyArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	AutoSnapshotPolicyId string
	OwnerId              int64
}
type DeleteAutoSnapshotPolicyResponse struct {
	RequestId string
}

func (c *EcsClient) DescribeNetworkInterfaces(req *DescribeNetworkInterfacesArgs) (resp *DescribeNetworkInterfacesResponse, err error) {
	resp = &DescribeNetworkInterfacesResponse{}
	err = c.Request("DescribeNetworkInterfaces", req, resp)
	return
}

type DescribeNetworkInterfacesNetworkInterfaceSet struct {
	NetworkInterfaceId   string
	Status               string
	Type                 string
	VpcId                string
	VSwitchId            string
	ZoneId               string
	PrivateIpAddress     string
	MacAddress           string
	NetworkInterfaceName string
	Description          string
	InstanceId           string
	CreationTime         string
	PrivateIpSets        DescribeNetworkInterfacesPrivateIpSetList
	SecurityGroupIds     DescribeNetworkInterfacesSecurityGroupIdList
	AssociatedPublicIp   DescribeNetworkInterfacesAssociatedPublicIp
}

type DescribeNetworkInterfacesPrivateIpSet struct {
	PrivateIpAddress    string
	Primary             bool
	AssociatedPublicIp1 DescribeNetworkInterfacesAssociatedPublicIp1
}

type DescribeNetworkInterfacesAssociatedPublicIp1 struct {
	PublicIpAddress string
	AllocationId    string
}

type DescribeNetworkInterfacesAssociatedPublicIp struct {
	PublicIpAddress string
	AllocationId    string
}
type DescribeNetworkInterfacesArgs struct {
	ResourceOwnerId      int64
	SecurityGroupId      string
	Type                 string
	PageNumber           int
	PageSize             int
	NetworkInterfaceName string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	VSwitchId            string
	InstanceId           string
	PrimaryIpAddress     string
	NetworkInterfaceIds  DescribeNetworkInterfacesNetworkInterfaceIdList
}
type DescribeNetworkInterfacesResponse struct {
	RequestId            string
	TotalCount           int
	PageNumber           int
	PageSize             int
	NetworkInterfaceSets DescribeNetworkInterfacesNetworkInterfaceSetList
}

type DescribeNetworkInterfacesPrivateIpSetList []DescribeNetworkInterfacesPrivateIpSet

func (list *DescribeNetworkInterfacesPrivateIpSetList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeNetworkInterfacesPrivateIpSet)
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

type DescribeNetworkInterfacesSecurityGroupIdList []string

func (list *DescribeNetworkInterfacesSecurityGroupIdList) UnmarshalJSON(data []byte) error {
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

type DescribeNetworkInterfacesNetworkInterfaceIdList []string

func (list *DescribeNetworkInterfacesNetworkInterfaceIdList) UnmarshalJSON(data []byte) error {
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

type DescribeNetworkInterfacesNetworkInterfaceSetList []DescribeNetworkInterfacesNetworkInterfaceSet

func (list *DescribeNetworkInterfacesNetworkInterfaceSetList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeNetworkInterfacesNetworkInterfaceSet)
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

func (c *EcsClient) DescribeSnapshotsUsage(req *DescribeSnapshotsUsageArgs) (resp *DescribeSnapshotsUsageResponse, err error) {
	resp = &DescribeSnapshotsUsageResponse{}
	err = c.Request("DescribeSnapshotsUsage", req, resp)
	return
}

type DescribeSnapshotsUsageArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
}
type DescribeSnapshotsUsageResponse struct {
	RequestId     string
	SnapshotCount int
	SnapshotSize  int64
}

func (c *EcsClient) DeleteBandwidthPackage(req *DeleteBandwidthPackageArgs) (resp *DeleteBandwidthPackageResponse, err error) {
	resp = &DeleteBandwidthPackageResponse{}
	err = c.Request("DeleteBandwidthPackage", req, resp)
	return
}

type DeleteBandwidthPackageArgs struct {
	ResourceOwnerId      int64
	BandwidthPackageId   string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
}
type DeleteBandwidthPackageResponse struct {
	RequestId string
}

func (c *EcsClient) DescribeInvocations(req *DescribeInvocationsArgs) (resp *DescribeInvocationsResponse, err error) {
	resp = &DescribeInvocationsResponse{}
	err = c.Request("DescribeInvocations", req, resp)
	return
}

type DescribeInvocationsInvocationItem struct {
	InvokeId     string
	CommandId    string
	CommandType  string
	CommandName  string
	Timed        bool
	InvokeStatus string
	InvokeItem   DescribeInvocationsInvokeItemItemList
}

type DescribeInvocationsInvokeItemItem struct {
	InstanceId string
	Status     string
}
type DescribeInvocationsArgs struct {
	ResourceOwnerId      int64
	InvokeStatus         string
	CommandId            string
	PageNumber           int64
	PageSize             int64
	InvokeId             string
	Timed                bool
	CommandName          string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	CommandType          string
	InstanceId           string
}
type DescribeInvocationsResponse struct {
	RequestId  string
	TotalCount int64
	PageNumber int64
	PageSize   int64
	Invocation DescribeInvocationsInvocationItemList
}

type DescribeInvocationsInvokeItemItemList []DescribeInvocationsInvokeItemItem

func (list *DescribeInvocationsInvokeItemItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInvocationsInvokeItemItem)
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

type DescribeInvocationsInvocationItemList []DescribeInvocationsInvocationItem

func (list *DescribeInvocationsInvocationItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInvocationsInvocationItem)
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

func (c *EcsClient) DetachClassicLinkVpc(req *DetachClassicLinkVpcArgs) (resp *DetachClassicLinkVpcResponse, err error) {
	resp = &DetachClassicLinkVpcResponse{}
	err = c.Request("DetachClassicLinkVpc", req, resp)
	return
}

type DetachClassicLinkVpcArgs struct {
	ResourceOwnerId      int64
	InstanceId           string
	ResourceOwnerAccount string
	VpcId                string
	OwnerId              int64
}
type DetachClassicLinkVpcResponse struct {
	RequestId string
}

func (c *EcsClient) AllocateEipAddress(req *AllocateEipAddressArgs) (resp *AllocateEipAddressResponse, err error) {
	resp = &AllocateEipAddressResponse{}
	err = c.Request("AllocateEipAddress", req, resp)
	return
}

type AllocateEipAddressArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	Bandwidth            string
	ClientToken          string
	InternetChargeType   string
	OwnerAccount         string
	OwnerId              int64
}
type AllocateEipAddressResponse struct {
	RequestId    string
	AllocationId string
	EipAddress   string
}

func (c *EcsClient) DescribeDiskMonitorData(req *DescribeDiskMonitorDataArgs) (resp *DescribeDiskMonitorDataResponse, err error) {
	resp = &DescribeDiskMonitorDataResponse{}
	err = c.Request("DescribeDiskMonitorData", req, resp)
	return
}

type DescribeDiskMonitorDataDiskMonitorData struct {
	DiskId       string
	IOPSRead     int
	IOPSWrite    int
	IOPSTotal    int
	BPSRead      int
	BPSWrite     int
	BPSTotal     int
	LatencyRead  int
	LatencyWrite int
	TimeStamp    string
}
type DescribeDiskMonitorDataArgs struct {
	ResourceOwnerId      int64
	Period               int
	ResourceOwnerAccount string
	OwnerAccount         string
	EndTime              string
	DiskId               string
	StartTime            string
	OwnerId              int64
}
type DescribeDiskMonitorDataResponse struct {
	RequestId   string
	TotalCount  int
	MonitorData DescribeDiskMonitorDataDiskMonitorDataList
}

type DescribeDiskMonitorDataDiskMonitorDataList []DescribeDiskMonitorDataDiskMonitorData

func (list *DescribeDiskMonitorDataDiskMonitorDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDiskMonitorDataDiskMonitorData)
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

func (c *EcsClient) CreateHaVip(req *CreateHaVipArgs) (resp *CreateHaVipResponse, err error) {
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

func (c *EcsClient) CreateInstance(req *CreateInstanceArgs) (resp *CreateInstanceResponse, err error) {
	resp = &CreateInstanceResponse{}
	err = c.Request("CreateInstance", req, resp)
	return
}

type CreateInstanceDataDisk struct {
	Size               int
	SnapshotId         string
	Category           string
	DiskName           string
	Description        string
	Device             string
	DeleteWithInstance bool
	Encrypted          bool
}
type CreateInstanceArgs struct {
	Tag4Value                   string
	ResourceOwnerId             int64
	Tag2Key                     string
	HpcClusterId                string
	Tag3Key                     string
	SecurityEnhancementStrategy string
	KeyPairName                 string
	SpotPriceLimit              float32
	Tag1Value                   string
	ResourceGroupId             string
	HostName                    string
	Password                    string
	AutoRenewPeriod             int
	NodeControllerId            string
	Period                      int
	Tag5Key                     string
	OwnerId                     int64
	VSwitchId                   string
	PrivateIpAddress            string
	SpotStrategy                string
	PeriodUnit                  string
	InstanceName                string
	AutoRenew                   bool
	InternetChargeType          string
	ZoneId                      string
	Tag4Key                     string
	InternetMaxBandwidthIn      int
	UseAdditionalService        bool
	ImageId                     string
	ClientToken                 string
	VlanId                      string
	IoOptimized                 string
	SecurityGroupId             string
	InternetMaxBandwidthOut     int
	Description                 string
	SystemDiskCategory          string
	UserData                    string
	InstanceType                string
	InstanceChargeType          string
	Tag3Value                   string
	DeploymentSetId             string
	InnerIpAddress              string
	ResourceOwnerAccount        string
	OwnerAccount                string
	SystemDiskDiskName          string
	RamRoleName                 string
	ClusterId                   string
	DataDisks                   CreateInstanceDataDiskList
	Tag5Value                   string
	Tag1Key                     string
	SystemDiskSize              int
	Tag2Value                   string
	SystemDiskDescription       string
}
type CreateInstanceResponse struct {
	RequestId  string
	InstanceId string
}

type CreateInstanceDataDiskList []CreateInstanceDataDisk

func (list *CreateInstanceDataDiskList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateInstanceDataDisk)
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

func (c *EcsClient) DescribeRenewalPrice(req *DescribeRenewalPriceArgs) (resp *DescribeRenewalPriceResponse, err error) {
	resp = &DescribeRenewalPriceResponse{}
	err = c.Request("DescribeRenewalPrice", req, resp)
	return
}

type DescribeRenewalPricePriceInfo struct {
	Rules DescribeRenewalPriceRuleList
	Price DescribeRenewalPricePrice
}

type DescribeRenewalPriceRule struct {
	RuleId      int64
	Description string
}

type DescribeRenewalPricePrice struct {
	OriginalPrice float32
	DiscountPrice float32
	TradePrice    float32
	Currency      string
}
type DescribeRenewalPriceArgs struct {
	ResourceOwnerId      int64
	ResourceId           string
	Period               int
	ResourceOwnerAccount string
	OwnerAccount         string
	PriceUnit            string
	OwnerId              int64
	ResourceType         string
}
type DescribeRenewalPriceResponse struct {
	RequestId string
	PriceInfo DescribeRenewalPricePriceInfo
}

type DescribeRenewalPriceRuleList []DescribeRenewalPriceRule

func (list *DescribeRenewalPriceRuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRenewalPriceRule)
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

func (c *EcsClient) ModifyDeploymentSetAttribute(req *ModifyDeploymentSetAttributeArgs) (resp *ModifyDeploymentSetAttributeResponse, err error) {
	resp = &ModifyDeploymentSetAttributeResponse{}
	err = c.Request("ModifyDeploymentSetAttribute", req, resp)
	return
}

type ModifyDeploymentSetAttributeArgs struct {
	DeploymentSetId      string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	Description          string
	DeploymentSetName    string
	OwnerId              int64
}
type ModifyDeploymentSetAttributeResponse struct {
	RequestId string
}

func (c *EcsClient) DeleteKeyPairs(req *DeleteKeyPairsArgs) (resp *DeleteKeyPairsResponse, err error) {
	resp = &DeleteKeyPairsResponse{}
	err = c.Request("DeleteKeyPairs", req, resp)
	return
}

type DeleteKeyPairsArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	KeyPairNames         string
	OwnerId              int64
}
type DeleteKeyPairsResponse struct {
	RequestId string
}

func (c *EcsClient) ModifyVirtualBorderRouterAttribute(req *ModifyVirtualBorderRouterAttributeArgs) (resp *ModifyVirtualBorderRouterAttributeResponse, err error) {
	resp = &ModifyVirtualBorderRouterAttributeResponse{}
	err = c.Request("ModifyVirtualBorderRouterAttribute", req, resp)
	return
}

type ModifyVirtualBorderRouterAttributeArgs struct {
	ResourceOwnerId      int64
	CircuitCode          string
	VlanId               int
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	Description          string
	VbrId                string
	OwnerId              int64
	PeerGatewayIp        string
	PeeringSubnetMask    string
	Name                 string
	LocalGatewayIp       string
	UserCidr             string
}
type ModifyVirtualBorderRouterAttributeResponse struct {
	RequestId string
}

func (c *EcsClient) DescribeAutoSnapshotPolicy(req *DescribeAutoSnapshotPolicyArgs) (resp *DescribeAutoSnapshotPolicyResponse, err error) {
	resp = &DescribeAutoSnapshotPolicyResponse{}
	err = c.Request("DescribeAutoSnapshotPolicy", req, resp)
	return
}

type DescribeAutoSnapshotPolicyAutoSnapshotPolicy struct {
	SystemDiskPolicyEnabled           string
	SystemDiskPolicyTimePeriod        string
	SystemDiskPolicyRetentionDays     string
	SystemDiskPolicyRetentionLastWeek string
	DataDiskPolicyEnabled             string
	DataDiskPolicyTimePeriod          string
	DataDiskPolicyRetentionDays       string
	DataDiskPolicyRetentionLastWeek   string
}

type DescribeAutoSnapshotPolicyAutoSnapshotExcutionStatus struct {
	SystemDiskExcutionStatus string
	DataDiskExcutionStatus   string
}
type DescribeAutoSnapshotPolicyArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
}
type DescribeAutoSnapshotPolicyResponse struct {
	RequestId                  string
	AutoSnapshotOccupation     int
	AutoSnapshotPolicy         DescribeAutoSnapshotPolicyAutoSnapshotPolicy
	AutoSnapshotExcutionStatus DescribeAutoSnapshotPolicyAutoSnapshotExcutionStatus
}

func (c *EcsClient) AttachKeyPair(req *AttachKeyPairArgs) (resp *AttachKeyPairResponse, err error) {
	resp = &AttachKeyPairResponse{}
	err = c.Request("AttachKeyPair", req, resp)
	return
}

type AttachKeyPairResult struct {
	InstanceId string
	Success    string
	Code       string
	Message    string
}
type AttachKeyPairArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	InstanceIds          string
	KeyPairName          string
	OwnerId              int64
}
type AttachKeyPairResponse struct {
	RequestId   string
	TotalCount  string
	FailCount   string
	KeyPairName string
	Results     AttachKeyPairResultList
}

type AttachKeyPairResultList []AttachKeyPairResult

func (list *AttachKeyPairResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]AttachKeyPairResult)
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

func (c *EcsClient) CancelAutoSnapshotPolicy(req *CancelAutoSnapshotPolicyArgs) (resp *CancelAutoSnapshotPolicyResponse, err error) {
	resp = &CancelAutoSnapshotPolicyResponse{}
	err = c.Request("CancelAutoSnapshotPolicy", req, resp)
	return
}

type CancelAutoSnapshotPolicyArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	DiskIds              string
	OwnerId              int64
}
type CancelAutoSnapshotPolicyResponse struct {
	RequestId string
}

func (c *EcsClient) CreatePhysicalConnection(req *CreatePhysicalConnectionArgs) (resp *CreatePhysicalConnectionResponse, err error) {
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

func (c *EcsClient) DescribeRouterInterfaces(req *DescribeRouterInterfacesArgs) (resp *DescribeRouterInterfacesResponse, err error) {
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

func (c *EcsClient) CreateNatGateway(req *CreateNatGatewayArgs) (resp *CreateNatGatewayResponse, err error) {
	resp = &CreateNatGatewayResponse{}
	err = c.Request("CreateNatGateway", req, resp)
	return
}

type CreateNatGatewayBandwidthPackage struct {
	IpCount   int
	Bandwidth int
	Zone      string
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
}
type CreateNatGatewayResponse struct {
	RequestId           string
	NatGatewayId        string
	ForwardTableIds     CreateNatGatewayForwardTableIdList
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

func (c *EcsClient) DescribeVpcs(req *DescribeVpcsArgs) (resp *DescribeVpcsResponse, err error) {
	resp = &DescribeVpcsResponse{}
	err = c.Request("DescribeVpcs", req, resp)
	return
}

type DescribeVpcsVpc struct {
	VpcId        string
	RegionId     string
	Status       string
	VpcName      string
	CreationTime string
	CidrBlock    string
	VRouterId    string
	Description  string
	IsDefault    bool
	VSwitchIds   DescribeVpcsVSwitchIdList
	UserCidrs    DescribeVpcsUserCidrList
}
type DescribeVpcsArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	VpcId                string
	OwnerAccount         string
	PageSize             int
	IsDefault            bool
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

func (c *EcsClient) DescribeDisksFullStatus(req *DescribeDisksFullStatusArgs) (resp *DescribeDisksFullStatusResponse, err error) {
	resp = &DescribeDisksFullStatusResponse{}
	err = c.Request("DescribeDisksFullStatus", req, resp)
	return
}

type DescribeDisksFullStatusDiskFullStatusType struct {
	DiskId       string
	DiskEventSet DescribeDisksFullStatusDiskEventTypeList
	Status       DescribeDisksFullStatusStatus
	HealthStatus DescribeDisksFullStatusHealthStatus
}

type DescribeDisksFullStatusDiskEventType struct {
	EventId   string
	EventTime string
	EventType DescribeDisksFullStatusEventType
}

type DescribeDisksFullStatusEventType struct {
	Code int
	Name string
}

type DescribeDisksFullStatusStatus struct {
	Code int
	Name string
}

type DescribeDisksFullStatusHealthStatus struct {
	Code int
	Name string
}
type DescribeDisksFullStatusArgs struct {
	EventIds             DescribeDisksFullStatusEventIdList
	ResourceOwnerId      int64
	PageNumber           int
	EventTimeStart       string
	PageSize             int
	DiskIds              DescribeDisksFullStatusDiskIdList
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	EventTimeEnd         string
	HealthStatus         string
	EventType            string
	Status               string
}
type DescribeDisksFullStatusResponse struct {
	RequestId         string
	TotalCount        int
	PageNumber        int
	PageSize          int
	DiskFullStatusSet DescribeDisksFullStatusDiskFullStatusTypeList
}

type DescribeDisksFullStatusDiskEventTypeList []DescribeDisksFullStatusDiskEventType

func (list *DescribeDisksFullStatusDiskEventTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDisksFullStatusDiskEventType)
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

type DescribeDisksFullStatusEventIdList []string

func (list *DescribeDisksFullStatusEventIdList) UnmarshalJSON(data []byte) error {
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

type DescribeDisksFullStatusDiskIdList []string

func (list *DescribeDisksFullStatusDiskIdList) UnmarshalJSON(data []byte) error {
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

type DescribeDisksFullStatusDiskFullStatusTypeList []DescribeDisksFullStatusDiskFullStatusType

func (list *DescribeDisksFullStatusDiskFullStatusTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDisksFullStatusDiskFullStatusType)
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

func (c *EcsClient) ModifyInstanceNetworkSpec(req *ModifyInstanceNetworkSpecArgs) (resp *ModifyInstanceNetworkSpecResponse, err error) {
	resp = &ModifyInstanceNetworkSpecResponse{}
	err = c.Request("ModifyInstanceNetworkSpec", req, resp)
	return
}

type ModifyInstanceNetworkSpecArgs struct {
	ResourceOwnerId         int64
	AutoPay                 bool
	ResourceOwnerAccount    string
	ClientToken             string
	OwnerAccount            string
	InternetMaxBandwidthOut int
	EndTime                 string
	StartTime               string
	OwnerId                 int64
	InstanceId              string
	NetworkChargeType       string
	InternetMaxBandwidthIn  int
	AllocatePublicIp        bool
}
type ModifyInstanceNetworkSpecResponse struct {
	RequestId string
	OrderId   string
}

func (c *EcsClient) CheckDiskEnableAutoSnapshotValidation(req *CheckDiskEnableAutoSnapshotValidationArgs) (resp *CheckDiskEnableAutoSnapshotValidationResponse, err error) {
	resp = &CheckDiskEnableAutoSnapshotValidationResponse{}
	err = c.Request("CheckDiskEnableAutoSnapshotValidation", req, resp)
	return
}

type CheckDiskEnableAutoSnapshotValidationArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	DiskIds              string
	OwnerId              int64
}
type CheckDiskEnableAutoSnapshotValidationResponse struct {
	RequestId              string
	IsPermitted            string
	AutoSnapshotOccupation int
}

func (c *EcsClient) DeleteDisk(req *DeleteDiskArgs) (resp *DeleteDiskResponse, err error) {
	resp = &DeleteDiskResponse{}
	err = c.Request("DeleteDisk", req, resp)
	return
}

type DeleteDiskArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	DiskId               string
	OwnerId              int64
}
type DeleteDiskResponse struct {
	RequestId string
}

func (c *EcsClient) JoinResourceGroup(req *JoinResourceGroupArgs) (resp *JoinResourceGroupResponse, err error) {
	resp = &JoinResourceGroupResponse{}
	err = c.Request("JoinResourceGroup", req, resp)
	return
}

type JoinResourceGroupArgs struct {
	ResourceGroupId      string
	ResourceOwnerId      int64
	ResourceId           string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	ResourceType         string
}
type JoinResourceGroupResponse struct {
	RequestId string
}

func (c *EcsClient) CreateDisk(req *CreateDiskArgs) (resp *CreateDiskResponse, err error) {
	resp = &CreateDiskResponse{}
	err = c.Request("CreateDisk", req, resp)
	return
}

type CreateDiskArgs struct {
	Tag4Value            string
	ResourceOwnerId      int64
	SnapshotId           string
	Tag2Key              string
	ClientToken          string
	Description          string
	Tag3Key              string
	DiskName             string
	Tag1Value            string
	ResourceGroupId      string
	DiskCategory         string
	Tag3Value            string
	Tag5Key              string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	Tag5Value            string
	Tag1Key              string
	Size                 int
	Encrypted            bool
	Tag2Value            string
	ZoneId               string
	Tag4Key              string
}
type CreateDiskResponse struct {
	RequestId string
	DiskId    string
}

func (c *EcsClient) ModifyIntranetBandwidthKb(req *ModifyIntranetBandwidthKbArgs) (resp *ModifyIntranetBandwidthKbResponse, err error) {
	resp = &ModifyIntranetBandwidthKbResponse{}
	err = c.Request("ModifyIntranetBandwidthKb", req, resp)
	return
}

type ModifyIntranetBandwidthKbArgs struct {
	ResourceOwnerId         int64
	IntranetMaxBandwidthOut int
	InstanceId              string
	IntranetMaxBandwidthIn  int
	ResourceOwnerAccount    string
	OwnerAccount            string
	OwnerId                 int64
}
type ModifyIntranetBandwidthKbResponse struct {
	RequestId string
}

func (c *EcsClient) ModifyRouterInterfaceSpec(req *ModifyRouterInterfaceSpecArgs) (resp *ModifyRouterInterfaceSpecResponse, err error) {
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

func (c *EcsClient) ModifyAutoSnapshotPolicyEx(req *ModifyAutoSnapshotPolicyExArgs) (resp *ModifyAutoSnapshotPolicyExResponse, err error) {
	resp = &ModifyAutoSnapshotPolicyExResponse{}
	err = c.Request("ModifyAutoSnapshotPolicyEx", req, resp)
	return
}

type ModifyAutoSnapshotPolicyExArgs struct {
	ResourceOwnerId        int64
	ResourceOwnerAccount   string
	AutoSnapshotPolicyId   string
	TimePoints             string
	RetentionDays          int
	OwnerId                int64
	RepeatWeekdays         string
	AutoSnapshotPolicyName string
}
type ModifyAutoSnapshotPolicyExResponse struct {
	RequestId string
}

func (c *EcsClient) CreateImage(req *CreateImageArgs) (resp *CreateImageResponse, err error) {
	resp = &CreateImageResponse{}
	err = c.Request("CreateImage", req, resp)
	return
}

type CreateImageDiskDeviceMapping struct {
	Size       int
	SnapshotId string
	Device     string
	DiskType   string
}
type CreateImageArgs struct {
	DiskDeviceMappings   CreateImageDiskDeviceMappingList
	Tag4Value            string
	ResourceOwnerId      int64
	SnapshotId           string
	Tag2Key              string
	ClientToken          string
	Description          string
	Tag3Key              string
	Platform             string
	Tag1Value            string
	ImageName            string
	Tag3Value            string
	Architecture         string
	Tag5Key              string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	Tag5Value            string
	Tag1Key              string
	InstanceId           string
	Tag2Value            string
	ImageVersion         string
	Tag4Key              string
}
type CreateImageResponse struct {
	RequestId string
	ImageId   string
}

type CreateImageDiskDeviceMappingList []CreateImageDiskDeviceMapping

func (list *CreateImageDiskDeviceMappingList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateImageDiskDeviceMapping)
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

func (c *EcsClient) ModifyUserBusinessBehavior(req *ModifyUserBusinessBehaviorArgs) (resp *ModifyUserBusinessBehaviorResponse, err error) {
	resp = &ModifyUserBusinessBehaviorResponse{}
	err = c.Request("ModifyUserBusinessBehavior", req, resp)
	return
}

type ModifyUserBusinessBehaviorArgs struct {
	ResourceOwnerId      int64
	StatusValue          string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	StatusKey            string
}
type ModifyUserBusinessBehaviorResponse struct {
	RequestId string
}

func (c *EcsClient) ModifyImageAttribute(req *ModifyImageAttributeArgs) (resp *ModifyImageAttributeResponse, err error) {
	resp = &ModifyImageAttributeResponse{}
	err = c.Request("ModifyImageAttribute", req, resp)
	return
}

type ModifyImageAttributeArgs struct {
	ResourceOwnerId      int64
	ImageId              string
	ResourceOwnerAccount string
	ImageName            string
	OwnerAccount         string
	Description          string
	OwnerId              int64
}
type ModifyImageAttributeResponse struct {
	RequestId string
}

func (c *EcsClient) DeleteInstance(req *DeleteInstanceArgs) (resp *DeleteInstanceResponse, err error) {
	resp = &DeleteInstanceResponse{}
	err = c.Request("DeleteInstance", req, resp)
	return
}

type DeleteInstanceArgs struct {
	ResourceOwnerId       int64
	InstanceId            string
	ResourceOwnerAccount  string
	OwnerAccount          string
	TerminateSubscription bool
	Force                 bool
	OwnerId               int64
}
type DeleteInstanceResponse struct {
	RequestId string
}

func (c *EcsClient) CancelTask(req *CancelTaskArgs) (resp *CancelTaskResponse, err error) {
	resp = &CancelTaskResponse{}
	err = c.Request("CancelTask", req, resp)
	return
}

type CancelTaskArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerId              int64
	TaskId               string
}
type CancelTaskResponse struct {
	RequestId string
}

func (c *EcsClient) AttachInstanceRamRole(req *AttachInstanceRamRoleArgs) (resp *AttachInstanceRamRoleResponse, err error) {
	resp = &AttachInstanceRamRoleResponse{}
	err = c.Request("AttachInstanceRamRole", req, resp)
	return
}

type AttachInstanceRamRoleAttachInstanceRamRoleResult struct {
	InstanceId string
	Success    bool
	Code       string
	Message    string
}
type AttachInstanceRamRoleArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	InstanceIds          string
	RamRoleName          string
	OwnerId              int64
}
type AttachInstanceRamRoleResponse struct {
	RequestId                    string
	TotalCount                   int
	FailCount                    int
	RamRoleName                  string
	AttachInstanceRamRoleResults AttachInstanceRamRoleAttachInstanceRamRoleResultList
}

type AttachInstanceRamRoleAttachInstanceRamRoleResultList []AttachInstanceRamRoleAttachInstanceRamRoleResult

func (list *AttachInstanceRamRoleAttachInstanceRamRoleResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]AttachInstanceRamRoleAttachInstanceRamRoleResult)
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

func (c *EcsClient) DescribeInstanceStatus(req *DescribeInstanceStatusArgs) (resp *DescribeInstanceStatusResponse, err error) {
	resp = &DescribeInstanceStatusResponse{}
	err = c.Request("DescribeInstanceStatus", req, resp)
	return
}

type DescribeInstanceStatusInstanceStatus struct {
	InstanceId string
	Status     string
}
type DescribeInstanceStatusArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	PageSize             int
	ZoneId               string
	ClusterId            string
	OwnerId              int64
	PageNumber           int
}
type DescribeInstanceStatusResponse struct {
	RequestId        string
	TotalCount       int
	PageNumber       int
	PageSize         int
	InstanceStatuses DescribeInstanceStatusInstanceStatusList
}

type DescribeInstanceStatusInstanceStatusList []DescribeInstanceStatusInstanceStatus

func (list *DescribeInstanceStatusInstanceStatusList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInstanceStatusInstanceStatus)
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

func (c *EcsClient) DescribeRecommendInstanceType(req *DescribeRecommendInstanceTypeArgs) (resp *DescribeRecommendInstanceTypeResponse, err error) {
	resp = &DescribeRecommendInstanceTypeResponse{}
	err = c.Request("DescribeRecommendInstanceType", req, resp)
	return
}

type DescribeRecommendInstanceTypeRecommendInstanceType struct {
	RegionNo      string
	CommodityCode string
	Scene         string
	Zones         DescribeRecommendInstanceTypeZoneList
	InstanceType  DescribeRecommendInstanceTypeInstanceType
}

type DescribeRecommendInstanceTypeZone struct {
	ZoneNo       string
	NetworkTypes DescribeRecommendInstanceTypeNetworkTypeList
}

type DescribeRecommendInstanceTypeInstanceType struct {
	Generation         string
	InstanceTypeFamily string
	InstanceType       string
	SupportIoOptimized string
	Cores              int
	Memory             int
}
type DescribeRecommendInstanceTypeArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	Channel              string
	NetworkType          string
	OwnerId              int64
	Operator             string
	Token                string
	Scene                string
	InstanceType         string
	ProxyId              string
}
type DescribeRecommendInstanceTypeResponse struct {
	RequestId string
	Data      DescribeRecommendInstanceTypeRecommendInstanceTypeList
}

type DescribeRecommendInstanceTypeZoneList []DescribeRecommendInstanceTypeZone

func (list *DescribeRecommendInstanceTypeZoneList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRecommendInstanceTypeZone)
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

type DescribeRecommendInstanceTypeNetworkTypeList []string

func (list *DescribeRecommendInstanceTypeNetworkTypeList) UnmarshalJSON(data []byte) error {
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

type DescribeRecommendInstanceTypeRecommendInstanceTypeList []DescribeRecommendInstanceTypeRecommendInstanceType

func (list *DescribeRecommendInstanceTypeRecommendInstanceTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRecommendInstanceTypeRecommendInstanceType)
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

func (c *EcsClient) DescribeHaVips(req *DescribeHaVipsArgs) (resp *DescribeHaVipsResponse, err error) {
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

func (c *EcsClient) DeleteForwardEntry(req *DeleteForwardEntryArgs) (resp *DeleteForwardEntryResponse, err error) {
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

func (c *EcsClient) ReInitDisk(req *ReInitDiskArgs) (resp *ReInitDiskResponse, err error) {
	resp = &ReInitDiskResponse{}
	err = c.Request("ReInitDisk", req, resp)
	return
}

type ReInitDiskArgs struct {
	ResourceOwnerId             int64
	Password                    string
	ResourceOwnerAccount        string
	AutoStartInstance           bool
	OwnerAccount                string
	DiskId                      string
	SecurityEnhancementStrategy string
	KeyPairName                 string
	OwnerId                     int64
}
type ReInitDiskResponse struct {
	RequestId string
}

func (c *EcsClient) ExportImage(req *ExportImageArgs) (resp *ExportImageResponse, err error) {
	resp = &ExportImageResponse{}
	err = c.Request("ExportImage", req, resp)
	return
}

type ExportImageArgs struct {
	ResourceOwnerId      int64
	ImageId              string
	OSSBucket            string
	ResourceOwnerAccount string
	OSSPrefix            string
	RoleName             string
	OwnerId              int64
	ImageFormat          string
}
type ExportImageResponse struct {
	RequestId string
	TaskId    string
	RegionId  string
}

func (c *EcsClient) DeleteSecurityGroup(req *DeleteSecurityGroupArgs) (resp *DeleteSecurityGroupResponse, err error) {
	resp = &DeleteSecurityGroupResponse{}
	err = c.Request("DeleteSecurityGroup", req, resp)
	return
}

type DeleteSecurityGroupArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	SecurityGroupId      string
	OwnerId              int64
}
type DeleteSecurityGroupResponse struct {
	RequestId string
}

func (c *EcsClient) DescribeNatGateways(req *DescribeNatGatewaysArgs) (resp *DescribeNatGatewaysResponse, err error) {
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
	ForwardTableIds     DescribeNatGatewaysForwardTableIdList
	BandwidthPackageIds DescribeNatGatewaysBandwidthPackageIdList
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

func (c *EcsClient) DescribeInstanceTypeFamilies(req *DescribeInstanceTypeFamiliesArgs) (resp *DescribeInstanceTypeFamiliesResponse, err error) {
	resp = &DescribeInstanceTypeFamiliesResponse{}
	err = c.Request("DescribeInstanceTypeFamilies", req, resp)
	return
}

type DescribeInstanceTypeFamiliesInstanceTypeFamily struct {
	InstanceTypeFamilyId string
	Generation           string
}
type DescribeInstanceTypeFamiliesArgs struct {
	Generation           string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
}
type DescribeInstanceTypeFamiliesResponse struct {
	RequestId            string
	InstanceTypeFamilies DescribeInstanceTypeFamiliesInstanceTypeFamilyList
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

func (c *EcsClient) DeleteDeploymentSet(req *DeleteDeploymentSetArgs) (resp *DeleteDeploymentSetResponse, err error) {
	resp = &DeleteDeploymentSetResponse{}
	err = c.Request("DeleteDeploymentSet", req, resp)
	return
}

type DeleteDeploymentSetArgs struct {
	DeploymentSetId      string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
}
type DeleteDeploymentSetResponse struct {
	RequestId string
}

func (c *EcsClient) TerminateVirtualBorderRouter(req *TerminateVirtualBorderRouterArgs) (resp *TerminateVirtualBorderRouterResponse, err error) {
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

func (c *EcsClient) ApplyAutoSnapshotPolicy(req *ApplyAutoSnapshotPolicyArgs) (resp *ApplyAutoSnapshotPolicyResponse, err error) {
	resp = &ApplyAutoSnapshotPolicyResponse{}
	err = c.Request("ApplyAutoSnapshotPolicy", req, resp)
	return
}

type ApplyAutoSnapshotPolicyArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	AutoSnapshotPolicyId string
	DiskIds              string
	OwnerId              int64
}
type ApplyAutoSnapshotPolicyResponse struct {
	RequestId string
}

func (c *EcsClient) ModifyInstanceChargeType(req *ModifyInstanceChargeTypeArgs) (resp *ModifyInstanceChargeTypeResponse, err error) {
	resp = &ModifyInstanceChargeTypeResponse{}
	err = c.Request("ModifyInstanceChargeType", req, resp)
	return
}

type ModifyInstanceChargeTypeArgs struct {
	ResourceOwnerId      int64
	Period               int
	DryRun               bool
	AutoPay              bool
	IncludeDataDisks     bool
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	OwnerId              int64
	PeriodUnit           string
	InstanceIds          string
}
type ModifyInstanceChargeTypeResponse struct {
	RequestId string
	OrderId   string
}

func (c *EcsClient) ModifyEipAddressAttribute(req *ModifyEipAddressAttributeArgs) (resp *ModifyEipAddressAttributeResponse, err error) {
	resp = &ModifyEipAddressAttributeResponse{}
	err = c.Request("ModifyEipAddressAttribute", req, resp)
	return
}

type ModifyEipAddressAttributeArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	Bandwidth            string
	OwnerAccount         string
	AllocationId         string
	OwnerId              int64
}
type ModifyEipAddressAttributeResponse struct {
	RequestId string
}

func (c *EcsClient) DescribeClusters(req *DescribeClustersArgs) (resp *DescribeClustersResponse, err error) {
	resp = &DescribeClustersResponse{}
	err = c.Request("DescribeClusters", req, resp)
	return
}

type DescribeClustersCluster struct {
	ClusterId string
}
type DescribeClustersArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
}
type DescribeClustersResponse struct {
	RequestId string
	Clusters  DescribeClustersClusterList
}

type DescribeClustersClusterList []DescribeClustersCluster

func (list *DescribeClustersClusterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClustersCluster)
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

func (c *EcsClient) DeactivateRouterInterface(req *DeactivateRouterInterfaceArgs) (resp *DeactivateRouterInterfaceResponse, err error) {
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

func (c *EcsClient) DescribeZones(req *DescribeZonesArgs) (resp *DescribeZonesResponse, err error) {
	resp = &DescribeZonesResponse{}
	err = c.Request("DescribeZones", req, resp)
	return
}

type DescribeZonesZone struct {
	ZoneId                    string
	LocalName                 string
	AvailableResources        DescribeZonesResourcesInfoList
	AvailableResourceCreation DescribeZonesAvailableResourceCreationList
	AvailableDiskCategories   DescribeZonesAvailableDiskCategoryList
	AvailableInstanceTypes    DescribeZonesAvailableInstanceTypeList
	AvailableVolumeCategories DescribeZonesAvailableVolumeCategoryList
}

type DescribeZonesResourcesInfo struct {
	IoOptimized          bool
	SystemDiskCategories DescribeZonesSystemDiskCategoryList
	DataDiskCategories   DescribeZonesDataDiskCategoryList
	NetworkTypes         DescribeZonesNetworkTypeList
	InstanceTypes        DescribeZonesInstanceTypeList
	InstanceTypeFamilies DescribeZonesInstanceTypeFamilyList
	InstanceGenerations  DescribeZonesInstanceGenerationList
}
type DescribeZonesArgs struct {
	SpotStrategy         string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	InstanceChargeType   string
	Verbose              bool
}
type DescribeZonesResponse struct {
	RequestId string
	Zones     DescribeZonesZoneList
}

type DescribeZonesResourcesInfoList []DescribeZonesResourcesInfo

func (list *DescribeZonesResourcesInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeZonesResourcesInfo)
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

type DescribeZonesAvailableResourceCreationList []string

func (list *DescribeZonesAvailableResourceCreationList) UnmarshalJSON(data []byte) error {
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

type DescribeZonesAvailableDiskCategoryList []string

func (list *DescribeZonesAvailableDiskCategoryList) UnmarshalJSON(data []byte) error {
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

type DescribeZonesAvailableInstanceTypeList []string

func (list *DescribeZonesAvailableInstanceTypeList) UnmarshalJSON(data []byte) error {
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

type DescribeZonesAvailableVolumeCategoryList []string

func (list *DescribeZonesAvailableVolumeCategoryList) UnmarshalJSON(data []byte) error {
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

type DescribeZonesSystemDiskCategoryList []string

func (list *DescribeZonesSystemDiskCategoryList) UnmarshalJSON(data []byte) error {
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

type DescribeZonesDataDiskCategoryList []string

func (list *DescribeZonesDataDiskCategoryList) UnmarshalJSON(data []byte) error {
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

type DescribeZonesNetworkTypeList []string

func (list *DescribeZonesNetworkTypeList) UnmarshalJSON(data []byte) error {
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

type DescribeZonesInstanceTypeList []string

func (list *DescribeZonesInstanceTypeList) UnmarshalJSON(data []byte) error {
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

type DescribeZonesInstanceTypeFamilyList []string

func (list *DescribeZonesInstanceTypeFamilyList) UnmarshalJSON(data []byte) error {
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

type DescribeZonesInstanceGenerationList []string

func (list *DescribeZonesInstanceGenerationList) UnmarshalJSON(data []byte) error {
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

func (c *EcsClient) AttachClassicLinkVpc(req *AttachClassicLinkVpcArgs) (resp *AttachClassicLinkVpcResponse, err error) {
	resp = &AttachClassicLinkVpcResponse{}
	err = c.Request("AttachClassicLinkVpc", req, resp)
	return
}

type AttachClassicLinkVpcArgs struct {
	ResourceOwnerId      int64
	InstanceId           string
	ResourceOwnerAccount string
	VpcId                string
	OwnerId              int64
}
type AttachClassicLinkVpcResponse struct {
	RequestId string
}

func (c *EcsClient) DescribeVirtualBorderRoutersForPhysicalConnection(req *DescribeVirtualBorderRoutersForPhysicalConnectionArgs) (resp *DescribeVirtualBorderRoutersForPhysicalConnectionResponse, err error) {
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

func (c *EcsClient) StopInstance(req *StopInstanceArgs) (resp *StopInstanceResponse, err error) {
	resp = &StopInstanceResponse{}
	err = c.Request("StopInstance", req, resp)
	return
}

type StopInstanceArgs struct {
	ResourceOwnerId      int64
	InstanceId           string
	ResourceOwnerAccount string
	ConfirmStop          bool
	OwnerAccount         string
	StoppedMode          string
	OwnerId              int64
	ForceStop            bool
}
type StopInstanceResponse struct {
	RequestId string
}

func (c *EcsClient) UnbindIpRange(req *UnbindIpRangeArgs) (resp *UnbindIpRangeResponse, err error) {
	resp = &UnbindIpRangeResponse{}
	err = c.Request("UnbindIpRange", req, resp)
	return
}

type UnbindIpRangeArgs struct {
	IpAddress            string
	ResourceOwnerId      int64
	InstanceId           string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
}
type UnbindIpRangeResponse struct {
	RequestId string
}

func (c *EcsClient) DescribeEipMonitorData(req *DescribeEipMonitorDataArgs) (resp *DescribeEipMonitorDataResponse, err error) {
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

func (c *EcsClient) ModifyVpcAttribute(req *ModifyVpcAttributeArgs) (resp *ModifyVpcAttributeResponse, err error) {
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

func (c *EcsClient) DescribeInstancesFullStatus(req *DescribeInstancesFullStatusArgs) (resp *DescribeInstancesFullStatusResponse, err error) {
	resp = &DescribeInstancesFullStatusResponse{}
	err = c.Request("DescribeInstancesFullStatus", req, resp)
	return
}

type DescribeInstancesFullStatusInstanceFullStatusType struct {
	InstanceId              string
	ScheduledSystemEventSet DescribeInstancesFullStatusScheduledSystemEventTypeList
	Status                  DescribeInstancesFullStatusStatus
	HealthStatus            DescribeInstancesFullStatusHealthStatus
}

type DescribeInstancesFullStatusScheduledSystemEventType struct {
	EventId          string
	EventPublishTime string
	NotBefore        string
	EventCycleStatus DescribeInstancesFullStatusEventCycleStatus
	EventType        DescribeInstancesFullStatusEventType
}

type DescribeInstancesFullStatusEventCycleStatus struct {
	Code int
	Name string
}

type DescribeInstancesFullStatusEventType struct {
	Code int
	Name string
}

type DescribeInstancesFullStatusStatus struct {
	Code int
	Name string
}

type DescribeInstancesFullStatusHealthStatus struct {
	Code int
	Name string
}
type DescribeInstancesFullStatusArgs struct {
	EventIds              DescribeInstancesFullStatusEventIdList
	ResourceOwnerId       int64
	PageNumber            int
	PageSize              int
	EventPublishTimeEnd   string
	ResourceOwnerAccount  string
	OwnerAccount          string
	NotBeforeStart        string
	OwnerId               int64
	EventPublishTimeStart string
	InstanceIds           DescribeInstancesFullStatusInstanceIdList
	NotBeforeEnd          string
	HealthStatus          string
	EventType             string
	Status                string
}
type DescribeInstancesFullStatusResponse struct {
	RequestId             string
	TotalCount            int
	PageNumber            int
	PageSize              int
	InstanceFullStatusSet DescribeInstancesFullStatusInstanceFullStatusTypeList
}

type DescribeInstancesFullStatusScheduledSystemEventTypeList []DescribeInstancesFullStatusScheduledSystemEventType

func (list *DescribeInstancesFullStatusScheduledSystemEventTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInstancesFullStatusScheduledSystemEventType)
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

type DescribeInstancesFullStatusEventIdList []string

func (list *DescribeInstancesFullStatusEventIdList) UnmarshalJSON(data []byte) error {
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

type DescribeInstancesFullStatusInstanceIdList []string

func (list *DescribeInstancesFullStatusInstanceIdList) UnmarshalJSON(data []byte) error {
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

type DescribeInstancesFullStatusInstanceFullStatusTypeList []DescribeInstancesFullStatusInstanceFullStatusType

func (list *DescribeInstancesFullStatusInstanceFullStatusTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInstancesFullStatusInstanceFullStatusType)
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

func (c *EcsClient) ModifyAutoSnapshotPolicy(req *ModifyAutoSnapshotPolicyArgs) (resp *ModifyAutoSnapshotPolicyResponse, err error) {
	resp = &ModifyAutoSnapshotPolicyResponse{}
	err = c.Request("ModifyAutoSnapshotPolicy", req, resp)
	return
}

type ModifyAutoSnapshotPolicyArgs struct {
	DataDiskPolicyEnabled             bool
	ResourceOwnerId                   int64
	DataDiskPolicyRetentionDays       int
	ResourceOwnerAccount              string
	SystemDiskPolicyRetentionLastWeek bool
	OwnerAccount                      string
	SystemDiskPolicyTimePeriod        int
	OwnerId                           int64
	DataDiskPolicyRetentionLastWeek   bool
	SystemDiskPolicyRetentionDays     int
	DataDiskPolicyTimePeriod          int
	SystemDiskPolicyEnabled           bool
}
type ModifyAutoSnapshotPolicyResponse struct {
	RequestId string
}

func (c *EcsClient) DescribeAccessPoints(req *DescribeAccessPointsArgs) (resp *DescribeAccessPointsResponse, err error) {
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
	ResourceOwnerAccount string
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

func (c *EcsClient) LeaveSecurityGroup(req *LeaveSecurityGroupArgs) (resp *LeaveSecurityGroupResponse, err error) {
	resp = &LeaveSecurityGroupResponse{}
	err = c.Request("LeaveSecurityGroup", req, resp)
	return
}

type LeaveSecurityGroupArgs struct {
	ResourceOwnerId      int64
	InstanceId           string
	ResourceOwnerAccount string
	OwnerAccount         string
	SecurityGroupId      string
	OwnerId              int64
}
type LeaveSecurityGroupResponse struct {
	RequestId string
}

func (c *EcsClient) DescribeSnapshots(req *DescribeSnapshotsArgs) (resp *DescribeSnapshotsResponse, err error) {
	resp = &DescribeSnapshotsResponse{}
	err = c.Request("DescribeSnapshots", req, resp)
	return
}

type DescribeSnapshotsSnapshot struct {
	SnapshotId        string
	SnapshotName      string
	Progress          string
	ProductCode       string
	SourceDiskId      string
	SourceDiskType    string
	RetentionDays     int
	Encrypted         bool
	SourceDiskSize    string
	Description       string
	CreationTime      string
	Status            string
	Usage             string
	SourceStorageType string
	Tags              DescribeSnapshotsTagList
}

type DescribeSnapshotsTag struct {
	TagKey   string
	TagValue string
}
type DescribeSnapshotsArgs struct {
	Tag4Value            string
	ResourceOwnerId      int64
	Tag2Key              string
	Filter2Value         string
	SnapshotIds          string
	Usage                string
	SnapshotLinkId       string
	SnapshotName         string
	Tag3Key              string
	PageNumber           int
	Tag1Value            string
	Filter1Key           string
	PageSize             int
	DiskId               string
	Tag3Value            string
	Tag5Key              string
	ResourceOwnerAccount string
	OwnerAccount         string
	SourceDiskType       string
	Filter1Value         string
	Filter2Key           string
	OwnerId              int64
	Tag5Value            string
	Tag1Key              string
	InstanceId           string
	Encrypted            bool
	SnapshotType         string
	Tag2Value            string
	Tag4Key              string
	Status               string
}
type DescribeSnapshotsResponse struct {
	RequestId  string
	TotalCount int
	PageNumber int
	PageSize   int
	Snapshots  DescribeSnapshotsSnapshotList
}

type DescribeSnapshotsTagList []DescribeSnapshotsTag

func (list *DescribeSnapshotsTagList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSnapshotsTag)
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

type DescribeSnapshotsSnapshotList []DescribeSnapshotsSnapshot

func (list *DescribeSnapshotsSnapshotList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSnapshotsSnapshot)
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

func (c *EcsClient) DescribeInstancePhysicalAttribute(req *DescribeInstancePhysicalAttributeArgs) (resp *DescribeInstancePhysicalAttributeResponse, err error) {
	resp = &DescribeInstancePhysicalAttributeResponse{}
	err = c.Request("DescribeInstancePhysicalAttribute", req, resp)
	return
}

type DescribeInstancePhysicalAttributeArgs struct {
	ResourceOwnerId      int64
	InstanceId           string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
}
type DescribeInstancePhysicalAttributeResponse struct {
	RequestId        string
	InstanceId       string
	VlanId           string
	NodeControllerId string
	RackId           string
}

func (c *EcsClient) AssociateEipAddress(req *AssociateEipAddressArgs) (resp *AssociateEipAddressResponse, err error) {
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

func (c *EcsClient) DescribeRegions(req *DescribeRegionsArgs) (resp *DescribeRegionsResponse, err error) {
	resp = &DescribeRegionsResponse{}
	err = c.Request("DescribeRegions", req, resp)
	return
}

type DescribeRegionsRegion struct {
	RegionId  string
	LocalName string
	Status    string
}
type DescribeRegionsArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	InstanceChargeType   string
	ResourceType         string
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

func (c *EcsClient) TerminatePhysicalConnection(req *TerminatePhysicalConnectionArgs) (resp *TerminatePhysicalConnectionResponse, err error) {
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

func (c *EcsClient) DeleteHpcCluster(req *DeleteHpcClusterArgs) (resp *DeleteHpcClusterResponse, err error) {
	resp = &DeleteHpcClusterResponse{}
	err = c.Request("DeleteHpcCluster", req, resp)
	return
}

type DeleteHpcClusterArgs struct {
	ResourceOwnerId      int64
	HpcClusterId         string
	ClientToken          string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
}
type DeleteHpcClusterResponse struct {
	RequestId string
}

func (c *EcsClient) DescribeImages(req *DescribeImagesArgs) (resp *DescribeImagesResponse, err error) {
	resp = &DescribeImagesResponse{}
	err = c.Request("DescribeImages", req, resp)
	return
}

type DescribeImagesImage struct {
	Progress             string
	ImageId              string
	ImageName            string
	ImageVersion         string
	Description          string
	Size                 int
	ImageOwnerAlias      string
	IsSupportIoOptimized bool
	IsSupportCloudinit   bool
	OSName               string
	Architecture         string
	Status               string
	ProductCode          string
	IsSubscribed         bool
	CreationTime         string
	IsSelfShared         string
	OSType               string
	Platform             string
	Usage                string
	IsCopied             bool
	DiskDeviceMappings   DescribeImagesDiskDeviceMappingList
	Tags                 DescribeImagesTagList
}

type DescribeImagesDiskDeviceMapping struct {
	SnapshotId      string
	Size            string
	Device          string
	Type            string
	Format          string
	ImportOSSBucket string
	ImportOSSObject string
}

type DescribeImagesTag struct {
	TagKey   string
	TagValue string
}
type DescribeImagesArgs struct {
	Tag4Value            string
	ResourceOwnerId      int64
	ImageId              string
	SnapshotId           string
	Tag2Key              string
	Filter2Value         string
	Usage                string
	Tag3Key              string
	PageNumber           int
	ImageOwnerAlias      string
	Tag1Value            string
	IsSupportIoOptimized bool
	Filter1Key           string
	ImageName            string
	IsSupportCloudinit   bool
	PageSize             int
	InstanceType         string
	Tag3Value            string
	Architecture         string
	Tag5Key              string
	ResourceOwnerAccount string
	OwnerAccount         string
	ShowExpired          bool
	Filter1Value         string
	OSType               string
	Filter2Key           string
	OwnerId              int64
	Tag5Value            string
	Tag1Key              string
	Tag2Value            string
	Tag4Key              string
	Status               string
}
type DescribeImagesResponse struct {
	RequestId  string
	RegionId   string
	TotalCount int
	PageNumber int
	PageSize   int
	Images     DescribeImagesImageList
}

type DescribeImagesDiskDeviceMappingList []DescribeImagesDiskDeviceMapping

func (list *DescribeImagesDiskDeviceMappingList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeImagesDiskDeviceMapping)
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

type DescribeImagesTagList []DescribeImagesTag

func (list *DescribeImagesTagList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeImagesTag)
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

type DescribeImagesImageList []DescribeImagesImage

func (list *DescribeImagesImageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeImagesImage)
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

func (c *EcsClient) CreateRouteEntry(req *CreateRouteEntryArgs) (resp *CreateRouteEntryResponse, err error) {
	resp = &CreateRouteEntryResponse{}
	err = c.Request("CreateRouteEntry", req, resp)
	return
}

type CreateRouteEntryNextHopList struct {
	NextHopType string
	NextHopId   string
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

func (c *EcsClient) DescribeSecurityGroupReferences(req *DescribeSecurityGroupReferencesArgs) (resp *DescribeSecurityGroupReferencesResponse, err error) {
	resp = &DescribeSecurityGroupReferencesResponse{}
	err = c.Request("DescribeSecurityGroupReferences", req, resp)
	return
}

type DescribeSecurityGroupReferencesSecurityGroupReference struct {
	SecurityGroupId           string
	ReferencingSecurityGroups DescribeSecurityGroupReferencesReferencingSecurityGroupList
}

type DescribeSecurityGroupReferencesReferencingSecurityGroup struct {
	AliUid          string
	SecurityGroupId string
}
type DescribeSecurityGroupReferencesArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	SecurityGroupIds     DescribeSecurityGroupReferencesSecurityGroupIdList
	OwnerId              int64
}
type DescribeSecurityGroupReferencesResponse struct {
	RequestId               string
	SecurityGroupReferences DescribeSecurityGroupReferencesSecurityGroupReferenceList
}

type DescribeSecurityGroupReferencesReferencingSecurityGroupList []DescribeSecurityGroupReferencesReferencingSecurityGroup

func (list *DescribeSecurityGroupReferencesReferencingSecurityGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSecurityGroupReferencesReferencingSecurityGroup)
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

type DescribeSecurityGroupReferencesSecurityGroupIdList []string

func (list *DescribeSecurityGroupReferencesSecurityGroupIdList) UnmarshalJSON(data []byte) error {
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

type DescribeSecurityGroupReferencesSecurityGroupReferenceList []DescribeSecurityGroupReferencesSecurityGroupReference

func (list *DescribeSecurityGroupReferencesSecurityGroupReferenceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSecurityGroupReferencesSecurityGroupReference)
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

func (c *EcsClient) CreateVpc(req *CreateVpcArgs) (resp *CreateVpcResponse, err error) {
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

func (c *EcsClient) EipFillParams(req *EipFillParamsArgs) (resp *EipFillParamsResponse, err error) {
	resp = &EipFillParamsResponse{}
	err = c.Request("EipFillParams", req, resp)
	return
}

type EipFillParamsArgs struct {
	ResourceOwnerId      int64
	Data                 string
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	UserCidr             string
	OwnerId              int64
}
type EipFillParamsResponse struct {
	RequestId string
	Data      string
	Code      string
	Success   bool
	Message   string
}

func (c *EcsClient) ResizeDisk(req *ResizeDiskArgs) (resp *ResizeDiskResponse, err error) {
	resp = &ResizeDiskResponse{}
	err = c.Request("ResizeDisk", req, resp)
	return
}

type ResizeDiskArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	NewSize              int
	DiskId               string
	OwnerId              int64
}
type ResizeDiskResponse struct {
	RequestId string
}

func (c *EcsClient) DescribeSecurityGroupAttribute(req *DescribeSecurityGroupAttributeArgs) (resp *DescribeSecurityGroupAttributeResponse, err error) {
	resp = &DescribeSecurityGroupAttributeResponse{}
	err = c.Request("DescribeSecurityGroupAttribute", req, resp)
	return
}

type DescribeSecurityGroupAttributePermission struct {
	IpProtocol              string
	PortRange               string
	SourceGroupId           string
	SourceGroupName         string
	SourceCidrIp            string
	Policy                  string
	NicType                 string
	SourceGroupOwnerAccount string
	DestGroupId             string
	DestGroupName           string
	DestCidrIp              string
	DestGroupOwnerAccount   string
	Priority                string
	Direction               string
	Description             string
	CreateTime              string
}
type DescribeSecurityGroupAttributeArgs struct {
	NicType              string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	SecurityGroupId      string
	OwnerId              int64
	Direction            string
}
type DescribeSecurityGroupAttributeResponse struct {
	RequestId         string
	RegionId          string
	SecurityGroupId   string
	Description       string
	SecurityGroupName string
	VpcId             string
	InnerAccessPolicy string
	Permissions       DescribeSecurityGroupAttributePermissionList
}

type DescribeSecurityGroupAttributePermissionList []DescribeSecurityGroupAttributePermission

func (list *DescribeSecurityGroupAttributePermissionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSecurityGroupAttributePermission)
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

func (c *EcsClient) ModifyPhysicalConnectionAttribute(req *ModifyPhysicalConnectionAttributeArgs) (resp *ModifyPhysicalConnectionAttributeResponse, err error) {
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

func (c *EcsClient) DetachKeyPair(req *DetachKeyPairArgs) (resp *DetachKeyPairResponse, err error) {
	resp = &DetachKeyPairResponse{}
	err = c.Request("DetachKeyPair", req, resp)
	return
}

type DetachKeyPairResult struct {
	InstanceId string
	Success    string
	Code       string
	Message    string
}
type DetachKeyPairArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	InstanceIds          string
	KeyPairName          string
	OwnerId              int64
}
type DetachKeyPairResponse struct {
	RequestId   string
	TotalCount  string
	FailCount   string
	KeyPairName string
	Results     DetachKeyPairResultList
}

type DetachKeyPairResultList []DetachKeyPairResult

func (list *DetachKeyPairResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DetachKeyPairResult)
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

func (c *EcsClient) ModifySecurityGroupRule(req *ModifySecurityGroupRuleArgs) (resp *ModifySecurityGroupRuleResponse, err error) {
	resp = &ModifySecurityGroupRuleResponse{}
	err = c.Request("ModifySecurityGroupRule", req, resp)
	return
}

type ModifySecurityGroupRuleArgs struct {
	NicType                 string
	ResourceOwnerId         int64
	SourcePortRange         string
	ClientToken             string
	SecurityGroupId         string
	Description             string
	SourceGroupOwnerId      int64
	SourceGroupOwnerAccount string
	Policy                  string
	PortRange               string
	ResourceOwnerAccount    string
	IpProtocol              string
	OwnerAccount            string
	SourceCidrIp            string
	OwnerId                 int64
	Priority                string
	DestCidrIp              string
	SourceGroupId           string
}
type ModifySecurityGroupRuleResponse struct {
	RequestId string
}

func (c *EcsClient) DescribeForwardTableEntries(req *DescribeForwardTableEntriesArgs) (resp *DescribeForwardTableEntriesResponse, err error) {
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

func (c *EcsClient) DescribeVSwitches(req *DescribeVSwitchesArgs) (resp *DescribeVSwitchesResponse, err error) {
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
	IsDefault               bool
}
type DescribeVSwitchesArgs struct {
	VSwitchId            string
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	VpcId                string
	OwnerAccount         string
	PageSize             int
	ZoneId               string
	IsDefault            bool
	OwnerId              int64
	PageNumber           int
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

func (c *EcsClient) DescribeInstanceMonitorData(req *DescribeInstanceMonitorDataArgs) (resp *DescribeInstanceMonitorDataResponse, err error) {
	resp = &DescribeInstanceMonitorDataResponse{}
	err = c.Request("DescribeInstanceMonitorData", req, resp)
	return
}

type DescribeInstanceMonitorDataInstanceMonitorData struct {
	InstanceId        string
	CPU               int
	IntranetRX        int
	IntranetTX        int
	IntranetBandwidth int
	InternetRX        int
	InternetTX        int
	InternetBandwidth int
	IOPSRead          int
	IOPSWrite         int
	BPSRead           int
	BPSWrite          int
	CPUCreditUsage    float32
	CPUCreditBalance  float32
	TimeStamp         string
}
type DescribeInstanceMonitorDataArgs struct {
	ResourceOwnerId      int64
	Period               int
	InstanceId           string
	ResourceOwnerAccount string
	OwnerAccount         string
	EndTime              string
	StartTime            string
	OwnerId              int64
}
type DescribeInstanceMonitorDataResponse struct {
	RequestId   string
	MonitorData DescribeInstanceMonitorDataInstanceMonitorDataList
}

type DescribeInstanceMonitorDataInstanceMonitorDataList []DescribeInstanceMonitorDataInstanceMonitorData

func (list *DescribeInstanceMonitorDataInstanceMonitorDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeInstanceMonitorDataInstanceMonitorData)
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

func (c *EcsClient) DescribeSnapshotPackage(req *DescribeSnapshotPackageArgs) (resp *DescribeSnapshotPackageResponse, err error) {
	resp = &DescribeSnapshotPackageResponse{}
	err = c.Request("DescribeSnapshotPackage", req, resp)
	return
}

type DescribeSnapshotPackageSnapshotPackage struct {
	StartTime    string
	EndTime      string
	InitCapacity int64
	DisplayName  string
}
type DescribeSnapshotPackageArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	PageSize             int
	OwnerId              int64
	PageNumber           int
}
type DescribeSnapshotPackageResponse struct {
	RequestId        string
	TotalCount       int
	PageNumber       int
	PageSize         int
	SnapshotPackages DescribeSnapshotPackageSnapshotPackageList
}

type DescribeSnapshotPackageSnapshotPackageList []DescribeSnapshotPackageSnapshotPackage

func (list *DescribeSnapshotPackageSnapshotPackageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSnapshotPackageSnapshotPackage)
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

func (c *EcsClient) EipFillProduct(req *EipFillProductArgs) (resp *EipFillProductResponse, err error) {
	resp = &EipFillProductResponse{}
	err = c.Request("EipFillProduct", req, resp)
	return
}

type EipFillProductArgs struct {
	ResourceOwnerId      int64
	Data                 string
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	UserCidr             string
	OwnerId              int64
}
type EipFillProductResponse struct {
	RequestId string
	Data      string
	Code      string
	Success   bool
	Message   string
}

func (c *EcsClient) ReplaceSystemDisk(req *ReplaceSystemDiskArgs) (resp *ReplaceSystemDiskResponse, err error) {
	resp = &ReplaceSystemDiskResponse{}
	err = c.Request("ReplaceSystemDisk", req, resp)
	return
}

type ReplaceSystemDiskArgs struct {
	ResourceOwnerId             int64
	ImageId                     string
	ResourceOwnerAccount        string
	ClientToken                 string
	OwnerAccount                string
	SecurityEnhancementStrategy string
	KeyPairName                 string
	OwnerId                     int64
	Password                    string
	InstanceId                  string
	SystemDiskSize              int
	UseAdditionalService        bool
}
type ReplaceSystemDiskResponse struct {
	RequestId string
	DiskId    string
}

func (c *EcsClient) CreateSecurityGroup(req *CreateSecurityGroupArgs) (resp *CreateSecurityGroupResponse, err error) {
	resp = &CreateSecurityGroupResponse{}
	err = c.Request("CreateSecurityGroup", req, resp)
	return
}

type CreateSecurityGroupArgs struct {
	Tag4Value            string
	ResourceOwnerId      int64
	Tag2Key              string
	Tag5Key              string
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	Description          string
	Tag3Key              string
	OwnerId              int64
	SecurityGroupName    string
	Tag5Value            string
	Tag1Key              string
	Tag1Value            string
	VpcId                string
	Tag2Value            string
	Tag4Key              string
	Tag3Value            string
}
type CreateSecurityGroupResponse struct {
	RequestId       string
	SecurityGroupId string
}

func (c *EcsClient) DetachInstanceRamRole(req *DetachInstanceRamRoleArgs) (resp *DetachInstanceRamRoleResponse, err error) {
	resp = &DetachInstanceRamRoleResponse{}
	err = c.Request("DetachInstanceRamRole", req, resp)
	return
}

type DetachInstanceRamRoleDetachInstanceRamRoleResult struct {
	InstanceId          string
	Success             bool
	Code                string
	Message             string
	InstanceRamRoleSets DetachInstanceRamRoleInstanceRamRoleSetList
}

type DetachInstanceRamRoleInstanceRamRoleSet struct {
	InstanceId  string
	RamRoleName string
}
type DetachInstanceRamRoleArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	InstanceIds          string
	RamRoleName          string
	OwnerId              int64
}
type DetachInstanceRamRoleResponse struct {
	RequestId                    string
	TotalCount                   int
	FailCount                    int
	RamRoleName                  string
	DetachInstanceRamRoleResults DetachInstanceRamRoleDetachInstanceRamRoleResultList
}

type DetachInstanceRamRoleInstanceRamRoleSetList []DetachInstanceRamRoleInstanceRamRoleSet

func (list *DetachInstanceRamRoleInstanceRamRoleSetList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DetachInstanceRamRoleInstanceRamRoleSet)
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

type DetachInstanceRamRoleDetachInstanceRamRoleResultList []DetachInstanceRamRoleDetachInstanceRamRoleResult

func (list *DetachInstanceRamRoleDetachInstanceRamRoleResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DetachInstanceRamRoleDetachInstanceRamRoleResult)
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

func (c *EcsClient) StopInvocation(req *StopInvocationArgs) (resp *StopInvocationResponse, err error) {
	resp = &StopInvocationResponse{}
	err = c.Request("StopInvocation", req, resp)
	return
}

type StopInvocationArgs struct {
	ResourceOwnerId      int64
	InvokeId             string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	InstanceIds          StopInvocationInstanceIdList
}
type StopInvocationResponse struct {
	RequestId string
}

type StopInvocationInstanceIdList []string

func (list *StopInvocationInstanceIdList) UnmarshalJSON(data []byte) error {
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

func (c *EcsClient) DescribeRecycleBin(req *DescribeRecycleBinArgs) (resp *DescribeRecycleBinResponse, err error) {
	resp = &DescribeRecycleBinResponse{}
	err = c.Request("DescribeRecycleBin", req, resp)
	return
}

type DescribeRecycleBinRecycleBinModel struct {
	ResourceId        string
	RegionId          string
	ResourceType      string
	CreationTime      string
	Status            string
	RelationResources DescribeRecycleBinRelationResourceList
}

type DescribeRecycleBinRelationResource struct {
	RelationResourceId   string
	RelationResourceType string
}
type DescribeRecycleBinArgs struct {
	ResourceOwnerId      int64
	ResourceId           string
	ResourceOwnerAccount string
	OwnerAccount         string
	PageSize             int
	OwnerId              int64
	PageNumber           int
	Status               string
}
type DescribeRecycleBinResponse struct {
	RequestId        string
	TotalCount       int
	RecycleBinModels DescribeRecycleBinRecycleBinModelList
}

type DescribeRecycleBinRelationResourceList []DescribeRecycleBinRelationResource

func (list *DescribeRecycleBinRelationResourceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRecycleBinRelationResource)
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

type DescribeRecycleBinRecycleBinModelList []DescribeRecycleBinRecycleBinModel

func (list *DescribeRecycleBinRecycleBinModelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeRecycleBinRecycleBinModel)
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

func (c *EcsClient) DescribePrice(req *DescribePriceArgs) (resp *DescribePriceResponse, err error) {
	resp = &DescribePriceResponse{}
	err = c.Request("DescribePrice", req, resp)
	return
}

type DescribePricePriceInfo struct {
	Rules DescribePriceRuleList
	Price DescribePricePrice
}

type DescribePriceRule struct {
	RuleId      int64
	Description string
}

type DescribePricePrice struct {
	OriginalPrice float32
	DiscountPrice float32
	TradePrice    float32
	Currency      string
}
type DescribePriceArgs struct {
	DataDisk3Size           int
	ResourceOwnerId         int64
	ImageId                 string
	DataDisk3Category       string
	IoOptimized             string
	InternetMaxBandwidthOut int
	SystemDiskCategory      string
	DataDisk4Category       string
	DataDisk4Size           int
	PriceUnit               string
	InstanceType            string
	DataDisk2Category       string
	DataDisk1Size           int
	Period                  int
	Amount                  int
	ResourceOwnerAccount    string
	OwnerAccount            string
	DataDisk2Size           int
	OwnerId                 int64
	ResourceType            string
	DataDisk1Category       string
	SystemDiskSize          int
	InternetChargeType      string
	InstanceNetworkType     string
}
type DescribePriceResponse struct {
	RequestId string
	PriceInfo DescribePricePriceInfo
}

type DescribePriceRuleList []DescribePriceRule

func (list *DescribePriceRuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribePriceRule)
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

func (c *EcsClient) DescribeTaskAttribute(req *DescribeTaskAttributeArgs) (resp *DescribeTaskAttributeResponse, err error) {
	resp = &DescribeTaskAttributeResponse{}
	err = c.Request("DescribeTaskAttribute", req, resp)
	return
}

type DescribeTaskAttributeOperationProgress struct {
	OperationStatus string
	ErrorCode       string
	ErrorMsg        string
	RelatedItemSet  DescribeTaskAttributeRelatedItemList
}

type DescribeTaskAttributeRelatedItem struct {
	Name  string
	Value string
}
type DescribeTaskAttributeArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerId              int64
	TaskId               string
}
type DescribeTaskAttributeResponse struct {
	RequestId            string
	TaskId               string
	RegionId             string
	TaskAction           string
	TaskStatus           string
	TaskProcess          string
	SupportCancel        string
	TotalCount           int
	SuccessCount         int
	FailedCount          int
	CreationTime         string
	FinishedTime         string
	OperationProgressSet DescribeTaskAttributeOperationProgressList
}

type DescribeTaskAttributeRelatedItemList []DescribeTaskAttributeRelatedItem

func (list *DescribeTaskAttributeRelatedItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeTaskAttributeRelatedItem)
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

type DescribeTaskAttributeOperationProgressList []DescribeTaskAttributeOperationProgress

func (list *DescribeTaskAttributeOperationProgressList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeTaskAttributeOperationProgress)
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

func (c *EcsClient) AddTags(req *AddTagsArgs) (resp *AddTagsResponse, err error) {
	resp = &AddTagsResponse{}
	err = c.Request("AddTags", req, resp)
	return
}

type AddTagsArgs struct {
	Tag4Value            string
	ResourceOwnerId      int64
	ResourceId           string
	Tag2Key              string
	Tag5Key              string
	ResourceOwnerAccount string
	Tag3Key              string
	OwnerId              int64
	ResourceType         string
	Tag5Value            string
	Tag1Key              string
	Tag1Value            string
	Tag2Value            string
	Tag4Key              string
	Tag3Value            string
}
type AddTagsResponse struct {
	RequestId string
}

func (c *EcsClient) ReleaseEipAddress(req *ReleaseEipAddressArgs) (resp *ReleaseEipAddressResponse, err error) {
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

func (c *EcsClient) CreateDeploymentSet(req *CreateDeploymentSetArgs) (resp *CreateDeploymentSetResponse, err error) {
	resp = &CreateDeploymentSetResponse{}
	err = c.Request("CreateDeploymentSet", req, resp)
	return
}

type CreateDeploymentSetArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	Description          string
	DeploymentSetName    string
	OwnerId              int64
	Granularity          string
	Domain               string
	ZoneId               string
	Strategy             string
}
type CreateDeploymentSetResponse struct {
	RequestId       string
	DeploymentSetId string
}

func (c *EcsClient) ModifySecurityGroupEgressRule(req *ModifySecurityGroupEgressRuleArgs) (resp *ModifySecurityGroupEgressRuleResponse, err error) {
	resp = &ModifySecurityGroupEgressRuleResponse{}
	err = c.Request("ModifySecurityGroupEgressRule", req, resp)
	return
}

type ModifySecurityGroupEgressRuleArgs struct {
	NicType               string
	ResourceOwnerId       int64
	SourcePortRange       string
	ClientToken           string
	SecurityGroupId       string
	Description           string
	Policy                string
	PortRange             string
	ResourceOwnerAccount  string
	IpProtocol            string
	OwnerAccount          string
	SourceCidrIp          string
	DestGroupId           string
	OwnerId               int64
	DestGroupOwnerAccount string
	Priority              string
	DestCidrIp            string
	DestGroupOwnerId      int64
}
type ModifySecurityGroupEgressRuleResponse struct {
	RequestId string
}

func (c *EcsClient) ModifyDiskAttribute(req *ModifyDiskAttributeArgs) (resp *ModifyDiskAttributeResponse, err error) {
	resp = &ModifyDiskAttributeResponse{}
	err = c.Request("ModifyDiskAttribute", req, resp)
	return
}

type ModifyDiskAttributeArgs struct {
	DiskName             string
	DeleteAutoSnapshot   bool
	ResourceOwnerId      int64
	EnableAutoSnapshot   bool
	ResourceOwnerAccount string
	OwnerAccount         string
	Description          string
	DiskId               string
	OwnerId              int64
	DeleteWithInstance   bool
}
type ModifyDiskAttributeResponse struct {
	RequestId string
}

func (c *EcsClient) DeleteNetworkInterface(req *DeleteNetworkInterfaceArgs) (resp *DeleteNetworkInterfaceResponse, err error) {
	resp = &DeleteNetworkInterfaceResponse{}
	err = c.Request("DeleteNetworkInterface", req, resp)
	return
}

type DeleteNetworkInterfaceArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	NetworkInterfaceId   string
}
type DeleteNetworkInterfaceResponse struct {
	RequestId string
}

func (c *EcsClient) DescribeSecurityGroups(req *DescribeSecurityGroupsArgs) (resp *DescribeSecurityGroupsResponse, err error) {
	resp = &DescribeSecurityGroupsResponse{}
	err = c.Request("DescribeSecurityGroups", req, resp)
	return
}

type DescribeSecurityGroupsSecurityGroup struct {
	SecurityGroupId         string
	Description             string
	SecurityGroupName       string
	VpcId                   string
	CreationTime            string
	AvailableInstanceAmount int
	EcsCount                int
	Tags                    DescribeSecurityGroupsTagList
}

type DescribeSecurityGroupsTag struct {
	TagKey   string
	TagValue string
}
type DescribeSecurityGroupsArgs struct {
	Tag4Value            string
	ResourceOwnerId      int64
	Tag2Key              string
	FuzzyQuery           bool
	SecurityGroupId      string
	Tag3Key              string
	IsQueryEcsCount      bool
	NetworkType          string
	SecurityGroupName    string
	PageNumber           int
	Tag1Value            string
	PageSize             int
	Tag3Value            string
	Tag5Key              string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
	SecurityGroupIds     string
	Tag5Value            string
	Tag1Key              string
	VpcId                string
	Tag2Value            string
	Tag4Key              string
}
type DescribeSecurityGroupsResponse struct {
	RequestId      string
	RegionId       string
	TotalCount     int
	PageNumber     int
	PageSize       int
	SecurityGroups DescribeSecurityGroupsSecurityGroupList
}

type DescribeSecurityGroupsTagList []DescribeSecurityGroupsTag

func (list *DescribeSecurityGroupsTagList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSecurityGroupsTag)
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

type DescribeSecurityGroupsSecurityGroupList []DescribeSecurityGroupsSecurityGroup

func (list *DescribeSecurityGroupsSecurityGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSecurityGroupsSecurityGroup)
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

func (c *EcsClient) EipNotifyPaid(req *EipNotifyPaidArgs) (resp *EipNotifyPaidResponse, err error) {
	resp = &EipNotifyPaidResponse{}
	err = c.Request("EipNotifyPaid", req, resp)
	return
}

type EipNotifyPaidArgs struct {
	ResourceOwnerId      int64
	Data                 string
	ResourceOwnerAccount string
	ClientToken          string
	OwnerAccount         string
	UserCidr             string
	OwnerId              int64
}
type EipNotifyPaidResponse struct {
	RequestId string
	Data      string
	Code      string
	Message   string
	Success   bool
}

func (c *EcsClient) ImportImage(req *ImportImageArgs) (resp *ImportImageResponse, err error) {
	resp = &ImportImageResponse{}
	err = c.Request("ImportImage", req, resp)
	return
}

type ImportImageDiskDeviceMapping struct {
	Format        string
	OSSBucket     string
	OSSObject     string
	DiskImSize    int
	DiskImageSize int
	Device        string
}
type ImportImageArgs struct {
	DiskDeviceMappings   ImportImageDiskDeviceMappingList
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	ImageName            string
	RoleName             string
	Description          string
	OSType               string
	OwnerId              int64
	Platform             string
	Architecture         string
}
type ImportImageResponse struct {
	RequestId string
	TaskId    string
	RegionId  string
	ImageId   string
}

type ImportImageDiskDeviceMappingList []ImportImageDiskDeviceMapping

func (list *ImportImageDiskDeviceMappingList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ImportImageDiskDeviceMapping)
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

func (c *EcsClient) RecoverVirtualBorderRouter(req *RecoverVirtualBorderRouterArgs) (resp *RecoverVirtualBorderRouterResponse, err error) {
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

func (c *EcsClient) DescribeImageSharePermission(req *DescribeImageSharePermissionArgs) (resp *DescribeImageSharePermissionResponse, err error) {
	resp = &DescribeImageSharePermissionResponse{}
	err = c.Request("DescribeImageSharePermission", req, resp)
	return
}

type DescribeImageSharePermissionShareGroup struct {
	Group string
}

type DescribeImageSharePermissionAccount struct {
	AliyunId string
}
type DescribeImageSharePermissionArgs struct {
	ResourceOwnerId      int64
	ImageId              string
	ResourceOwnerAccount string
	OwnerAccount         string
	PageSize             int
	OwnerId              int64
	PageNumber           int
}
type DescribeImageSharePermissionResponse struct {
	RequestId   string
	RegionId    string
	TotalCount  int
	PageNumber  int
	PageSize    int
	ImageId     string
	ShareGroups DescribeImageSharePermissionShareGroupList
	Accounts    DescribeImageSharePermissionAccountList
}

type DescribeImageSharePermissionShareGroupList []DescribeImageSharePermissionShareGroup

func (list *DescribeImageSharePermissionShareGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeImageSharePermissionShareGroup)
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

type DescribeImageSharePermissionAccountList []DescribeImageSharePermissionAccount

func (list *DescribeImageSharePermissionAccountList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeImageSharePermissionAccount)
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

func (c *EcsClient) DescribeIntranetAttributeKb(req *DescribeIntranetAttributeKbArgs) (resp *DescribeIntranetAttributeKbResponse, err error) {
	resp = &DescribeIntranetAttributeKbResponse{}
	err = c.Request("DescribeIntranetAttributeKb", req, resp)
	return
}

type DescribeIntranetAttributeKbArgs struct {
	ResourceOwnerId      int64
	InstanceId           string
	ResourceOwnerAccount string
	OwnerAccount         string
	OwnerId              int64
}
type DescribeIntranetAttributeKbResponse struct {
	RequestId               string
	InstanceId              string
	VlanId                  string
	IntranetIpAddress       string
	IntranetMaxBandwidthIn  int
	IntranetMaxBandwidthOut int
}

func (c *EcsClient) ModifySecurityGroupAttribute(req *ModifySecurityGroupAttributeArgs) (resp *ModifySecurityGroupAttributeResponse, err error) {
	resp = &ModifySecurityGroupAttributeResponse{}
	err = c.Request("ModifySecurityGroupAttribute", req, resp)
	return
}

type ModifySecurityGroupAttributeArgs struct {
	ResourceOwnerId      int64
	ResourceOwnerAccount string
	OwnerAccount         string
	SecurityGroupId      string
	Description          string
	OwnerId              int64
	SecurityGroupName    string
}
type ModifySecurityGroupAttributeResponse struct {
	RequestId string
}
