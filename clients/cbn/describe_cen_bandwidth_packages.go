package cbn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeCenBandwidthPackagesRequest struct {
	requests.RpcRequest
	Filters              *DescribeCenBandwidthPackagesFilterList `position:"Query" type:"Repeated" name:"Filter"`
	ResourceOwnerId      int64                                   `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string                                  `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string                                  `position:"Query" name:"OwnerAccount"`
	PageSize             int                                     `position:"Query" name:"PageSize"`
	OwnerId              int64                                   `position:"Query" name:"OwnerId"`
	PageNumber           int                                     `position:"Query" name:"PageNumber"`
	IsOrKey              string                                  `position:"Query" name:"IsOrKey"`
}

func (req *DescribeCenBandwidthPackagesRequest) Invoke(client *sdk.Client) (resp *DescribeCenBandwidthPackagesResponse, err error) {
	req.InitWithApiInfo("Cbn", "2017-09-12", "DescribeCenBandwidthPackages", "cbn", "")
	resp = &DescribeCenBandwidthPackagesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeCenBandwidthPackagesFilter struct {
	Key    string                                 `name:"Key"`
	Values *DescribeCenBandwidthPackagesValueList `type:"Repeated" name:"Value"`
}

type DescribeCenBandwidthPackagesResponse struct {
	responses.BaseResponse
	RequestId            common.String
	TotalCount           common.Integer
	PageNumber           common.Integer
	PageSize             common.Integer
	CenBandwidthPackages DescribeCenBandwidthPackagesCenBandwidthPackageList
}

type DescribeCenBandwidthPackagesCenBandwidthPackage struct {
	CenBandwidthPackageId      common.String
	Name                       common.String
	Description                common.String
	Bandwidth                  common.Long
	BandwidthPackageChargeType common.String
	GeographicRegionAId        common.String
	GeographicRegionBId        common.String
	BusinessStatus             common.String
	CreationTime               common.String
	ExpiredTime                common.String
	Status                     common.String
	CenIds                     DescribeCenBandwidthPackagesCenIdList
}

type DescribeCenBandwidthPackagesFilterList []DescribeCenBandwidthPackagesFilter

func (list *DescribeCenBandwidthPackagesFilterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCenBandwidthPackagesFilter)
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

type DescribeCenBandwidthPackagesValueList []string

func (list *DescribeCenBandwidthPackagesValueList) UnmarshalJSON(data []byte) error {
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

type DescribeCenBandwidthPackagesCenBandwidthPackageList []DescribeCenBandwidthPackagesCenBandwidthPackage

func (list *DescribeCenBandwidthPackagesCenBandwidthPackageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCenBandwidthPackagesCenBandwidthPackage)
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

type DescribeCenBandwidthPackagesCenIdList []common.String

func (list *DescribeCenBandwidthPackagesCenIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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
