package market

import (
	"encoding/json"

	"github.com/morlay/aliyun-go/core"
)

func (c *MarketClient) DescribeProduct(req *DescribeProductArgs) (resp *DescribeProductResponse, err error) {
	resp = &DescribeProductResponse{}
	err = c.Request("DescribeProduct", req, resp)
	return
}

type DescribeProductProductSku struct {
	Name         string
	Code         string
	ChargeType   string
	Constraints  string
	OrderPeriods DescribeProductOrderPeriodList
	Modules      DescribeProductModuleList
}

type DescribeProductOrderPeriod struct {
	Name       string
	PeriodType string
}

type DescribeProductModule struct {
	Id         string
	Name       string
	Code       string
	Properties DescribeProductPropertyList
}

type DescribeProductProperty struct {
	Name           string
	Key            string
	PropertyValues DescribeProductPropertyValueList
}

type DescribeProductPropertyValue struct {
	Value       string
	DisplayName string
	Type        string
}

type DescribeProductProductExtra struct {
	Key    string
	Values string
	Label  string
	Order  int
	Type   string
}

type DescribeProductShopInfo struct {
	Id         int64
	Name       string
	Emails     string
	WangWangs  DescribeProductWangWangList
	Telephones DescribeProductTelephoneList
}

type DescribeProductWangWang struct {
	UserName string
	Remark   string
}
type DescribeProductArgs struct {
	AliUid string
	Code   string
}
type DescribeProductResponse struct {
	Code             string
	Name             string
	Type             string
	PicUrl           string
	Description      string
	ShortDescription string
	ProductSkus      DescribeProductProductSkuList
	ProductExtras    DescribeProductProductExtraList
	ShopInfo         DescribeProductShopInfo
}

type DescribeProductOrderPeriodList []DescribeProductOrderPeriod

func (list *DescribeProductOrderPeriodList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeProductOrderPeriod)
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

type DescribeProductModuleList []DescribeProductModule

func (list *DescribeProductModuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeProductModule)
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

type DescribeProductPropertyList []DescribeProductProperty

func (list *DescribeProductPropertyList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeProductProperty)
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

type DescribeProductPropertyValueList []DescribeProductPropertyValue

func (list *DescribeProductPropertyValueList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeProductPropertyValue)
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

type DescribeProductWangWangList []DescribeProductWangWang

func (list *DescribeProductWangWangList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeProductWangWang)
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

type DescribeProductTelephoneList []string

func (list *DescribeProductTelephoneList) UnmarshalJSON(data []byte) error {
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

type DescribeProductProductSkuList []DescribeProductProductSku

func (list *DescribeProductProductSkuList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeProductProductSku)
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

type DescribeProductProductExtraList []DescribeProductProductExtra

func (list *DescribeProductProductExtraList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeProductProductExtra)
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

func (c *MarketClient) BindImagePackage(req *BindImagePackageArgs) (resp *BindImagePackageResponse, err error) {
	resp = &BindImagePackageResponse{}
	err = c.Request("BindImagePackage", req, resp)
	return
}

type BindImagePackageArgs struct {
	EcsInstanceId          string
	ImagePackageInstanceId string
}
type BindImagePackageResponse struct {
	RequestId string
	Success   core.Bool
}

func (c *MarketClient) DescribeInstance(req *DescribeInstanceArgs) (resp *DescribeInstanceResponse, err error) {
	resp = &DescribeInstanceResponse{}
	err = c.Request("DescribeInstance", req, resp)
	return
}

type DescribeInstanceArgs struct {
	InstanceId string
}
type DescribeInstanceResponse struct {
	InstanceId     int64
	OrderId        int64
	SupplierName   string
	ProductCode    string
	ProductSkuCode string
	ProductName    string
	ProductType    string
	Status         string
	BeganOn        int64
	EndOn          int64
	CreatedOn      int64
	ExtendJson     string
	HostJson       string
	AppJson        string
}

func (c *MarketClient) QueryMarketImages(req *QueryMarketImagesArgs) (resp *QueryMarketImagesResponse, err error) {
	resp = &QueryMarketImagesResponse{}
	err = c.Request("QueryMarketImages", req, resp)
	return
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
	SupportIO        core.Bool
	CreatedOn        int64
	Images           QueryMarketImagesImageList
	SkuCodes         QueryMarketImagesSkuCodeList
	Quota            QueryMarketImagesQuota
	PriceInfo        QueryMarketImagesPriceInfo
}

type QueryMarketImagesImage struct {
	Version            string
	ImageId            string
	ImageSize          int
	Region             string
	IsDefault          core.Bool
	SupportIO          core.Bool
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
type QueryMarketImagesArgs struct {
	Param string
}
type QueryMarketImagesResponse struct {
	PageNumber int
	PageSize   int
	TotalCount int
	RequestId  string
	Result     QueryMarketImagesImageProductList
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

func (c *MarketClient) DescribeLicense(req *DescribeLicenseArgs) (resp *DescribeLicenseResponse, err error) {
	resp = &DescribeLicenseResponse{}
	err = c.Request("DescribeLicense", req, resp)
	return
}

type DescribeLicenseLicense struct {
	LicenseStatus string
	LicenseCode   string
	InstanceId    string
	CreateTime    string
	ExpiredTime   string
	ActivateTime  string
	ProductSkuId  string
	ProductCode   string
	ProductName   string
	ExtendInfo    DescribeLicenseExtendInfo
}

type DescribeLicenseExtendInfo struct {
	AliUid          int64
	Email           string
	Mobile          string
	AccountQuantity int64
}
type DescribeLicenseArgs struct {
	LicenseCode string
}
type DescribeLicenseResponse struct {
	RequestId string
	License   DescribeLicenseLicense
}

func (c *MarketClient) DescribeOrder(req *DescribeOrderArgs) (resp *DescribeOrderResponse, err error) {
	resp = &DescribeOrderResponse{}
	err = c.Request("DescribeOrder", req, resp)
	return
}

type DescribeOrderArgs struct {
	OrderId string
}
type DescribeOrderResponse struct {
	OrderId             int64
	AliUid              int64
	SupplierCompanyName string
	ProductCode         string
	ProductSkuCode      string
	ProductName         string
	PeriodType          string
	Quantity            int
	AccountQuantity     int64
	OrderType           string
	OrderStatus         string
	PayStatus           string
	PaidOn              int64
	CreatedOn           int64
	OriginalPrice       float32
	TotalPrice          float32
	PaymentPrice        float32
	CouponPrice         float32
	SupplierTelephones  DescribeOrderSupplierTelephoneList
	InstanceIds         DescribeOrderInstanceIdList
}

type DescribeOrderSupplierTelephoneList []string

func (list *DescribeOrderSupplierTelephoneList) UnmarshalJSON(data []byte) error {
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

type DescribeOrderInstanceIdList []string

func (list *DescribeOrderInstanceIdList) UnmarshalJSON(data []byte) error {
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

func (c *MarketClient) QueryMarketCategories(req *QueryMarketCategoriesArgs) (resp *QueryMarketCategoriesResponse, err error) {
	resp = &QueryMarketCategoriesResponse{}
	err = c.Request("QueryMarketCategories", req, resp)
	return
}

type QueryMarketCategoriesCategory struct {
	Id           int64
	CategoryCode string
	CategoryName string
}
type QueryMarketCategoriesArgs struct {
}
type QueryMarketCategoriesResponse struct {
	PageNumber int
	PageSize   int
	TotalCount int
	RequestId  string
	Categories QueryMarketCategoriesCategoryList
}

type QueryMarketCategoriesCategoryList []QueryMarketCategoriesCategory

func (list *QueryMarketCategoriesCategoryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMarketCategoriesCategory)
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

func (c *MarketClient) DescribePrice(req *DescribePriceArgs) (resp *DescribePriceResponse, err error) {
	resp = &DescribePriceResponse{}
	err = c.Request("DescribePrice", req, resp)
	return
}

type DescribePricePromotionRule struct {
	RuleId string
	Name   string
	Title  string
}
type DescribePriceArgs struct {
	OrderType string
	Commodity string
}
type DescribePriceResponse struct {
	ProductCode    string
	OriginalPrice  float32
	TradePrice     float32
	PromotionRules DescribePricePromotionRuleList
}

type DescribePricePromotionRuleList []DescribePricePromotionRule

func (list *DescribePricePromotionRuleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribePricePromotionRule)
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

func (c *MarketClient) ActivateLicense(req *ActivateLicenseArgs) (resp *ActivateLicenseResponse, err error) {
	resp = &ActivateLicenseResponse{}
	err = c.Request("ActivateLicense", req, resp)
	return
}

type ActivateLicenseArgs struct {
	LicenseCode    string
	Identification string
}
type ActivateLicenseResponse struct {
	RequestId string
	Success   core.Bool
}

func (c *MarketClient) SubscribeImage(req *SubscribeImageArgs) (resp *SubscribeImageResponse, err error) {
	resp = &SubscribeImageResponse{}
	err = c.Request("SubscribeImage", req, resp)
	return
}

type SubscribeImageArgs struct {
	ProductCode string
}
type SubscribeImageResponse struct {
	RequestId string
	Success   core.Bool
}

func (c *MarketClient) CreateOrder(req *CreateOrderArgs) (resp *CreateOrderResponse, err error) {
	resp = &CreateOrderResponse{}
	err = c.Request("CreateOrder", req, resp)
	return
}

type CreateOrderArgs struct {
	OrderType   string
	Commodity   string
	OrderSouce  string
	PaymentType string
	ClientToken string
}
type CreateOrderResponse struct {
	OrderId string
}
