package market_inner

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type InnerReviewHistoryImageRequest struct {
	requests.RpcRequest
	ProductCode  string `position:"Query" name:"ProductCode"`
	ReviewType   string `position:"Query" name:"ReviewType"`
	ImageVersion string `position:"Query" name:"ImageVersion"`
	ImageNo      string `position:"Query" name:"ImageNo"`
	PageIndex    int    `position:"Query" name:"PageIndex"`
	RegionNo     string `position:"Query" name:"RegionNo"`
}

func (req *InnerReviewHistoryImageRequest) Invoke(client *sdk.Client) (resp *InnerReviewHistoryImageResponse, err error) {
	req.InitWithApiInfo("Market-Inner", "2016-08-01", "InnerReviewHistoryImage", "yunmarket", "")
	resp = &InnerReviewHistoryImageResponse{}
	err = client.DoAction(req, resp)
	return
}

type InnerReviewHistoryImageResponse struct {
	responses.BaseResponse
	TotalCount common.Integer
	RequestId  common.String
	IsSuccess  bool
	ImageList  InnerReviewHistoryImageImageList
}

type InnerReviewHistoryImageImage struct {
	RegionNo        common.String
	ProductCode     common.String
	ImageNo         common.String
	ImageVersion    common.String
	SupplierId      common.Long
	ProductName     common.String
	InstanceId      common.String
	InstanceAddress common.String
	OsType          common.String
	UserName        common.String
	SupplierName    common.String
	Password        common.String
	OsKind          common.String
	OsBit           common.Integer
	InstanceStatus  common.String
	ImageCreateTime common.String
}

type InnerReviewHistoryImageImageList []InnerReviewHistoryImageImage

func (list *InnerReviewHistoryImageImageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]InnerReviewHistoryImageImage)
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
