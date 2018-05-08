package cbn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeCensRequest struct {
	requests.RpcRequest
	Filters              *DescribeCensFilterList `position:"Query" type:"Repeated" name:"Filter"`
	ResourceOwnerId      int64                   `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string                  `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string                  `position:"Query" name:"OwnerAccount"`
	PageSize             int                     `position:"Query" name:"PageSize"`
	OwnerId              int64                   `position:"Query" name:"OwnerId"`
	PageNumber           int                     `position:"Query" name:"PageNumber"`
}

func (req *DescribeCensRequest) Invoke(client *sdk.Client) (resp *DescribeCensResponse, err error) {
	req.InitWithApiInfo("Cbn", "2017-09-12", "DescribeCens", "cbn", "")
	resp = &DescribeCensResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeCensFilter struct {
	Key    string                 `name:"Key"`
	Values *DescribeCensValueList `type:"Repeated" name:"Value"`
}

type DescribeCensResponse struct {
	responses.BaseResponse
	RequestId  common.String
	TotalCount common.Integer
	PageNumber common.Integer
	PageSize   common.Integer
	Cens       DescribeCensCenList
}

type DescribeCensCen struct {
	CenId                  common.String
	Name                   common.String
	Description            common.String
	Status                 common.String
	CreationTime           common.String
	CenBandwidthPackageIds DescribeCensCenBandwidthPackageIdList
}

type DescribeCensFilterList []DescribeCensFilter

func (list *DescribeCensFilterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCensFilter)
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

type DescribeCensValueList []string

func (list *DescribeCensValueList) UnmarshalJSON(data []byte) error {
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

type DescribeCensCenList []DescribeCensCen

func (list *DescribeCensCenList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCensCen)
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

type DescribeCensCenBandwidthPackageIdList []common.String

func (list *DescribeCensCenBandwidthPackageIdList) UnmarshalJSON(data []byte) error {
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
