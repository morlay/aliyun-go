package market_inner

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type InnerListSecurityAuditingImagesRequest struct {
	requests.RpcRequest
	Channel   int `position:"Query" name:"Channel"`
	PageIndex int `position:"Query" name:"PageIndex"`
}

func (req *InnerListSecurityAuditingImagesRequest) Invoke(client *sdk.Client) (resp *InnerListSecurityAuditingImagesResponse, err error) {
	req.InitWithApiInfo("Market-Inner", "2016-08-01", "InnerListSecurityAuditingImages", "yunmarket", "")
	resp = &InnerListSecurityAuditingImagesResponse{}
	err = client.DoAction(req, resp)
	return
}

type InnerListSecurityAuditingImagesResponse struct {
	responses.BaseResponse
	TotalCount common.Integer
	RequestId  common.String
	ImageList  InnerListSecurityAuditingImagesImageList
}

type InnerListSecurityAuditingImagesImage struct {
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
}

type InnerListSecurityAuditingImagesImageList []InnerListSecurityAuditingImagesImage

func (list *InnerListSecurityAuditingImagesImageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]InnerListSecurityAuditingImagesImage)
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
