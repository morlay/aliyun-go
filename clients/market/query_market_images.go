package market

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	PageNumber int
	PageSize   int
	TotalCount int
	RequestId  string
	Result     QueryMarketImagesImageProductList
}

type QueryMarketImagesImageProduct struct {
	ImageProductCode string
	ProductName      string
	CategoryName     string
	SupplierName     string
	BaseSystem       string
	OsKind           string
	OsBit            int
	PictureUrl       string
	SmallPicUrl      string
	ShortDescription string
	AgreementUrl     string
	DetailUrl        string
	BuyUrl           string
	StoreUrl         string
	Score            float32
	UserCount        int64
	SupportIO        bool
	CreatedOn        int64
	Images           QueryMarketImagesImageList
	SkuCodes         QueryMarketImagesSkuCodeList
	Quota            QueryMarketImagesQuota
	PriceInfo        QueryMarketImagesPriceInfo
}

type QueryMarketImagesImage struct {
	Version            string
	VersionDescription string
	ImageId            string
	ImageSize          int
	Region             string
	IsDefault          bool
	SupportIO          bool
	DiskDeviceMappings QueryMarketImagesDiskDeviceMappingList
}

type QueryMarketImagesDiskDeviceMapping struct {
	DiskType        string
	Format          string
	SnapshotId      string
	Size            int
	Device          string
	ImportOSSBucket string
	ImportOSSObject string
}

type QueryMarketImagesQuota struct {
	TotalQuota  int64
	UsingQuota  int64
	UnusedQuota int64
}

type QueryMarketImagesPriceInfo struct {
	Rules QueryMarketImagesRuleList
	Order QueryMarketImagesOrder
}

type QueryMarketImagesRule struct {
	RuleId int64
	Title  string
	Name   string
}

type QueryMarketImagesOrder struct {
	OriginalPrice float32
	DiscountPrice float32
	TradePrice    float32
	Currency      string
	Period        int
	PriceUnit     string
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

type QueryMarketImagesSkuCodeList []string

func (list *QueryMarketImagesSkuCodeList) UnmarshalJSON(data []byte) error {
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

type QueryMarketImagesRuleIdSetList []string

func (list *QueryMarketImagesRuleIdSetList) UnmarshalJSON(data []byte) error {
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
