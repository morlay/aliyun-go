package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeRouterInterfaces", "ecs", "")
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
	RequestId          string
	PageNumber         int
	PageSize           int
	TotalCount         int
	RouterInterfaceSet DescribeRouterInterfacesRouterInterfaceTypeList
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
