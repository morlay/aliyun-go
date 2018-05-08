package cloudapi

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeCatalogsRequest struct {
	requests.RpcRequest
}

func (req *DescribeCatalogsRequest) Invoke(client *sdk.Client) (resp *DescribeCatalogsResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeCatalogs", "apigateway", "")
	resp = &DescribeCatalogsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeCatalogsResponse struct {
	responses.BaseResponse
	RequestId         common.String
	TotalCount        common.Integer
	PageSize          common.Integer
	PageNumber        common.Integer
	CatalogAttributes DescribeCatalogsCatalogAttributeList
}

type DescribeCatalogsCatalogAttribute struct {
	CatalogId    common.String
	CatalogName  common.String
	Description  common.String
	ParentId     common.String
	CreatedTime  common.String
	ModifiedTime common.String
	RegionId     common.String
}

type DescribeCatalogsCatalogAttributeList []DescribeCatalogsCatalogAttribute

func (list *DescribeCatalogsCatalogAttributeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCatalogsCatalogAttribute)
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
