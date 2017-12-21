package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeBandwidthPackagesRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	BandwidthPackageId   string `position:"Query" name:"BandwidthPackageId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int    `position:"Query" name:"PageSize"`
	NatGatewayId         string `position:"Query" name:"NatGatewayId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
}

func (req *DescribeBandwidthPackagesRequest) Invoke(client *sdk.Client) (resp *DescribeBandwidthPackagesResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeBandwidthPackages", "ecs", "")
	resp = &DescribeBandwidthPackagesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeBandwidthPackagesResponse struct {
	responses.BaseResponse
	RequestId         string
	TotalCount        int
	PageNumber        int
	PageSize          int
	BandwidthPackages DescribeBandwidthPackagesBandwidthPackageList
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
