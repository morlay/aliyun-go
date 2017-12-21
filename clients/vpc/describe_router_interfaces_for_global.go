package vpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeRouterInterfacesForGlobalRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int    `position:"Query" name:"PageSize"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
	Status               string `position:"Query" name:"Status"`
}

func (r DescribeRouterInterfacesForGlobalRequest) Invoke(client *sdk.Client) (response *DescribeRouterInterfacesForGlobalResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeRouterInterfacesForGlobalRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Vpc", "2016-04-28", "DescribeRouterInterfacesForGlobal", "vpc", "")

	resp := struct {
		*responses.BaseResponse
		DescribeRouterInterfacesForGlobalResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeRouterInterfacesForGlobalResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeRouterInterfacesForGlobalResponse struct {
	RequestId          string
	Code               string
	Message            string
	Desc               string
	Success            bool
	PageSize           int
	PageNumber         int
	TotalCount         int
	RouterInterfaceSet DescribeRouterInterfacesForGlobalRouterInterfaceTypeList
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
