package market_inner

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type InnerQueryPurchaseImageRequest struct {
	requests.RpcRequest
	ImageId           string `position:"Query" name:"ImageId"`
	ImageProductCode  string `position:"Query" name:"ImageProductCode"`
	AliUid            int64  `position:"Query" name:"AliUid"`
	ChargeType        string `position:"Query" name:"ChargeType"`
	ImagePurchaseType string `position:"Query" name:"ImagePurchaseType"`
	ImageCategory     string `position:"Query" name:"ImageCategory"`
	RegionNo          string `position:"Query" name:"RegionNo"`
}

func (req *InnerQueryPurchaseImageRequest) Invoke(client *sdk.Client) (resp *InnerQueryPurchaseImageResponse, err error) {
	req.InitWithApiInfo("Market-Inner", "2016-08-01", "InnerQueryPurchaseImage", "yunmarket", "")
	resp = &InnerQueryPurchaseImageResponse{}
	err = client.DoAction(req, resp)
	return
}

type InnerQueryPurchaseImageResponse struct {
	responses.BaseResponse
	Image InnerQueryPurchaseImageImage
}

type InnerQueryPurchaseImageImage struct {
	ProductCode        common.String
	ProductSKUCode     common.String
	ImageId            common.String
	ImageVersion       common.String
	RegionId           common.String
	SupplierId         common.Long
	SupplierName       common.String
	DiskDeviceMappings InnerQueryPurchaseImageDiskDeviceMappingList
}

type InnerQueryPurchaseImageDiskDeviceMapping struct {
	DiskType        common.String
	Format          common.String
	SnapshotId      common.String
	Size            common.Integer
	Device          common.String
	ImportOSSBucket common.String
	ImportOSSObject common.String
}

type InnerQueryPurchaseImageDiskDeviceMappingList []InnerQueryPurchaseImageDiskDeviceMapping

func (list *InnerQueryPurchaseImageDiskDeviceMappingList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]InnerQueryPurchaseImageDiskDeviceMapping)
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
