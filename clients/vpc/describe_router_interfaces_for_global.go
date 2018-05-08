package vpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeRouterInterfacesForGlobalRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int    `position:"Query" name:"PageSize"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
	Status               string `position:"Query" name:"Status"`
}

func (req *DescribeRouterInterfacesForGlobalRequest) Invoke(client *sdk.Client) (resp *DescribeRouterInterfacesForGlobalResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "DescribeRouterInterfacesForGlobal", "vpc", "")
	resp = &DescribeRouterInterfacesForGlobalResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeRouterInterfacesForGlobalResponse struct {
	responses.BaseResponse
	RequestId          common.String
	Code               common.String
	Message            common.String
	Desc               common.String
	Success            bool
	PageSize           common.Integer
	PageNumber         common.Integer
	TotalCount         common.Integer
	RouterInterfaceSet DescribeRouterInterfacesForGlobalRouterInterfaceTypeList
}

type DescribeRouterInterfacesForGlobalRouterInterfaceType struct {
	BusinessStatus                  common.String
	AccessPointId                   common.String
	ChargeType                      common.String
	ConnectedTime                   common.String
	CreationTime                    common.String
	RouterInterfaceId               common.String
	OppositeInterfaceBusinessStatus common.String
	OppositeInterfaceId             common.String
	OppositeInterfaceOwnerId        common.Long
	OppositeInterfaceSpec           common.String
	OppositeInterfaceStatus         common.String
	OppositeRegionId                common.String
	OppositeAccessPointId           common.String
	OppositeRouterId                common.String
	OppositeRouterType              common.String
	OppositeVpcInstanceId           common.String
	RegionId                        common.String
	Role                            common.String
	RouterId                        common.String
	RouterType                      common.String
	Spec                            common.String
	Status                          common.String
	VpcInstanceId                   common.String
	Name                            common.String
	Description                     common.String
	HealthCheckSourceIp             common.String
	HealthCheckTargetIp             common.String
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
