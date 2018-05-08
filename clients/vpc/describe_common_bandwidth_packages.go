package vpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeCommonBandwidthPackagesRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	BandwidthPackageId   string `position:"Query" name:"BandwidthPackageId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	PageSize             int    `position:"Query" name:"PageSize"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
}

func (req *DescribeCommonBandwidthPackagesRequest) Invoke(client *sdk.Client) (resp *DescribeCommonBandwidthPackagesResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "DescribeCommonBandwidthPackages", "vpc", "")
	resp = &DescribeCommonBandwidthPackagesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeCommonBandwidthPackagesResponse struct {
	responses.BaseResponse
	RequestId               common.String
	TotalCount              common.Integer
	PageNumber              common.Integer
	PageSize                common.Integer
	CommonBandwidthPackages DescribeCommonBandwidthPackagesCommonBandwidthPackageList
}

type DescribeCommonBandwidthPackagesCommonBandwidthPackage struct {
	BandwidthPackageId common.String
	RegionId           common.String
	Name               common.String
	Description        common.String
	Bandwidth          common.String
	InstanceChargeType common.String
	InternetChargeType common.String
	BusinessStatus     common.String
	CreationTime       common.String
	ExpiredTime        common.String
	Status             common.String
	Ratio              common.Integer
	PublicIpAddresses  DescribeCommonBandwidthPackagesPublicIpAddresseList
}

type DescribeCommonBandwidthPackagesPublicIpAddresse struct {
	AllocationId common.String
	IpAddress    common.String
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
