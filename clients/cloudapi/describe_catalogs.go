package cloudapi

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeCatalogsRequest struct {
}

func (r DescribeCatalogsRequest) Invoke(client *sdk.Client) (response *DescribeCatalogsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeCatalogsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeCatalogs", "apigateway", "")

	resp := struct {
		*responses.BaseResponse
		DescribeCatalogsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeCatalogsResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeCatalogsResponse struct {
	RequestId         string
	TotalCount        int
	PageSize          int
	PageNumber        int
	CatalogAttributes DescribeCatalogsCatalogAttributeList
}

type DescribeCatalogsCatalogAttribute struct {
	CatalogId    string
	CatalogName  string
	Description  string
	ParentId     string
	CreatedTime  string
	ModifiedTime string
	RegionId     string
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
