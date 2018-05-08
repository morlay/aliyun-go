package vpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeRouterInterfacesRequest struct {
	requests.RpcRequest
	Filters              *DescribeRouterInterfacesFilterList `position:"Query" type:"Repeated" name:"Filter"`
	ResourceOwnerId      int64                               `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string                              `position:"Query" name:"ResourceOwnerAccount"`
	PageSize             int                                 `position:"Query" name:"PageSize"`
	OwnerId              int64                               `position:"Query" name:"OwnerId"`
	PageNumber           int                                 `position:"Query" name:"PageNumber"`
}

func (req *DescribeRouterInterfacesRequest) Invoke(client *sdk.Client) (resp *DescribeRouterInterfacesResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "DescribeRouterInterfaces", "vpc", "")
	resp = &DescribeRouterInterfacesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeRouterInterfacesFilter struct {
	Key    string                             `name:"Key"`
	Values *DescribeRouterInterfacesValueList `type:"Repeated" name:"Value"`
}

type DescribeRouterInterfacesResponse struct {
	responses.BaseResponse
	RequestId          common.String
	PageNumber         common.Integer
	PageSize           common.Integer
	TotalCount         common.Integer
	RouterInterfaceSet DescribeRouterInterfacesRouterInterfaceTypeList
}

type DescribeRouterInterfacesRouterInterfaceType struct {
	RouterInterfaceId               common.String
	OppositeRegionId                common.String
	Role                            common.String
	Spec                            common.String
	Name                            common.String
	Description                     common.String
	RouterId                        common.String
	RouterType                      common.String
	CreationTime                    common.String
	EndTime                         common.String
	ChargeType                      common.String
	Status                          common.String
	BusinessStatus                  common.String
	ConnectedTime                   common.String
	OppositeInterfaceId             common.String
	OppositeInterfaceSpec           common.String
	OppositeInterfaceStatus         common.String
	OppositeInterfaceBusinessStatus common.String
	OppositeRouterId                common.String
	OppositeRouterType              common.String
	OppositeInterfaceOwnerId        common.String
	AccessPointId                   common.String
	OppositeAccessPointId           common.String
	HealthCheckSourceIp             common.String
	HealthCheckTargetIp             common.String
	OppositeVpcInstanceId           common.String
	VpcInstanceId                   common.String
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
