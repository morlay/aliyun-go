package market_inner

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type InnerQueryPurchaseImagesRequest struct {
	requests.RpcRequest
	Param  string `position:"Query" name:"Param"`
	AliUid int64  `position:"Query" name:"AliUid"`
}

func (req *InnerQueryPurchaseImagesRequest) Invoke(client *sdk.Client) (resp *InnerQueryPurchaseImagesResponse, err error) {
	req.InitWithApiInfo("Market-Inner", "2016-08-01", "InnerQueryPurchaseImages", "yunmarket", "")
	resp = &InnerQueryPurchaseImagesResponse{}
	err = client.DoAction(req, resp)
	return
}

type InnerQueryPurchaseImagesResponse struct {
	responses.BaseResponse
	TotalCount common.Integer
	RequestId  common.String
	ImageList  InnerQueryPurchaseImagesImageList
}

type InnerQueryPurchaseImagesImage struct {
	RequestImageId common.String
	ChargeType     common.String
	ProductCode    common.String
	ProductSKUCode common.String
	ImageId        common.String
	ImageVersion   common.String
	ImageStatus    common.String
	RegionId       common.String
	SupplierId     common.Long
	SupplierName   common.String
}

type InnerQueryPurchaseImagesImageList []InnerQueryPurchaseImagesImage

func (list *InnerQueryPurchaseImagesImageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]InnerQueryPurchaseImagesImage)
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
