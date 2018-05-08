package market

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type QueryMarketImagesRequest struct {
	requests.RpcRequest
	Param string `position:"Query" name:"Param"`
}

func (req *QueryMarketImagesRequest) Invoke(client *sdk.Client) (resp *QueryMarketImagesResponse, err error) {
	req.InitWithApiInfo("Market", "2015-11-01", "QueryMarketImages", "yunmarket", "")
	resp = &QueryMarketImagesResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryMarketImagesResponse struct {
	responses.BaseResponse
	PageNumber common.Integer
	PageSize   common.Integer
	TotalCount common.Integer
	RequestId  common.String
	Result     QueryMarketImagesImageProductList
}

type QueryMarketImagesImageProduct struct {
	ImageProductCode common.String
	ProductName      common.String
	CategoryName     common.String
	SupplierName     common.String
	BaseSystem       common.String
	OsKind           common.String
	OsBit            common.Integer
	PictureUrl       common.String
	SmallPicUrl      common.String
	ShortDescription common.String
	AgreementUrl     common.String
	DetailUrl        common.String
	BuyUrl           common.String
	StoreUrl         common.String
	Score            common.Float
	UserCount        common.Long
	SupportIO        bool
	CreatedOn        common.Long
	Images           QueryMarketImagesImageList
	SkuCodes         QueryMarketImagesSkuCodeList
	Quota            QueryMarketImagesQuota
	PriceInfo        QueryMarketImagesPriceInfo
}

type QueryMarketImagesImage struct {
	Version            common.String
	VersionDescription common.String
	ImageId            common.String
	ImageSize          common.Integer
	Region             common.String
	IsDefault          bool
	SupportIO          bool
	DiskDeviceMappings QueryMarketImagesDiskDeviceMappingList
}

type QueryMarketImagesDiskDeviceMapping struct {
	DiskType        common.String
	Format          common.String
	SnapshotId      common.String
	Size            common.Integer
	Device          common.String
	ImportOSSBucket common.String
	ImportOSSObject common.String
}

type QueryMarketImagesQuota struct {
	TotalQuota  common.Long
	UsingQuota  common.Long
	UnusedQuota common.Long
}

type QueryMarketImagesPriceInfo struct {
	Rules QueryMarketImagesRuleList
	Order QueryMarketImagesOrder
}

type QueryMarketImagesRule struct {
	RuleId common.Long
	Title  common.String
	Name   common.String
}

type QueryMarketImagesOrder struct {
	OriginalPrice common.Float
	DiscountPrice common.Float
	TradePrice    common.Float
	Currency      common.String
	Period        common.Integer
	PriceUnit     common.String
	RuleIdSet     QueryMarketImagesRuleIdSetList
}

type QueryMarketImagesImageProductList []QueryMarketImagesImageProduct

func (list *QueryMarketImagesImageProductList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMarketImagesImageProduct)
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

type QueryMarketImagesImageList []QueryMarketImagesImage

func (list *QueryMarketImagesImageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMarketImagesImage)
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

type QueryMarketImagesSkuCodeList []common.String

func (list *QueryMarketImagesSkuCodeList) UnmarshalJSON(data []byte) error {
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

type QueryMarketImagesDiskDeviceMappingList []QueryMarketImagesDiskDeviceMapping

func (list *QueryMarketImagesDiskDeviceMappingList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMarketImagesDiskDeviceMapping)
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

type QueryMarketImagesRuleList []QueryMarketImagesRule

func (list *QueryMarketImagesRuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMarketImagesRule)
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

type QueryMarketImagesRuleIdSetList []common.String

func (list *QueryMarketImagesRuleIdSetList) UnmarshalJSON(data []byte) error {
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
