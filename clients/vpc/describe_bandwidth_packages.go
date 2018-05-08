package vpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	req.InitWithApiInfo("Vpc", "2016-04-28", "DescribeBandwidthPackages", "vpc", "")
	resp = &DescribeBandwidthPackagesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeBandwidthPackagesResponse struct {
	responses.BaseResponse
	RequestId         common.String
	TotalCount        common.Integer
	PageNumber        common.Integer
	PageSize          common.Integer
	BandwidthPackages DescribeBandwidthPackagesBandwidthPackageList
}

type DescribeBandwidthPackagesBandwidthPackage struct {
	BandwidthPackageId common.String
	RegionId           common.String
	Name               common.String
	Description        common.String
	ZoneId             common.String
	NatGatewayId       common.String
	Bandwidth          common.String
	InstanceChargeType common.String
	InternetChargeType common.String
	BusinessStatus     common.String
	IpCount            common.String
	CreationTime       common.String
	Status             common.String
	ISP                common.String
	PublicIpAddresses  DescribeBandwidthPackagesPublicIpAddresseList
}

type DescribeBandwidthPackagesPublicIpAddresse struct {
	AllocationId    common.String
	IpAddress       common.String
	UsingStatus     common.String
	ApAccessEnabled bool
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
