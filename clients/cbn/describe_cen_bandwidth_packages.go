package cbn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId            string
	TotalCount           int
	PageNumber           int
	PageSize             int
	CenBandwidthPackages DescribeCenBandwidthPackagesCenBandwidthPackageList
}

type DescribeCenBandwidthPackagesCenBandwidthPackage struct {
	CenBandwidthPackageId      string
	Name                       string
	Description                string
	Bandwidth                  int64
	BandwidthPackageChargeType string
	GeographicRegionAId        string
	GeographicRegionBId        string
	BusinessStatus             string
	CreationTime               string
	ExpiredTime                string
	Status                     string
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

type DescribeCenBandwidthPackagesCenIdList []string

func (list *DescribeCenBandwidthPackagesCenIdList) UnmarshalJSON(data []byte) error {
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
