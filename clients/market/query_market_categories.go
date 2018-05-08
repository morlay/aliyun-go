package market

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type QueryMarketCategoriesRequest struct {
	requests.RpcRequest
}

func (req *QueryMarketCategoriesRequest) Invoke(client *sdk.Client) (resp *QueryMarketCategoriesResponse, err error) {
	req.InitWithApiInfo("Market", "2015-11-01", "QueryMarketCategories", "yunmarket", "")
	resp = &QueryMarketCategoriesResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryMarketCategoriesResponse struct {
	responses.BaseResponse
	PageNumber common.Integer
	PageSize   common.Integer
	TotalCount common.Integer
	RequestId  common.String
	Categories QueryMarketCategoriesCategoryList
}

type QueryMarketCategoriesCategory struct {
	Id              common.Long
	CategoryCode    common.String
	CategoryName    common.String
	ChildCategories QueryMarketCategoriesChildCategoryList
}

type QueryMarketCategoriesChildCategory struct {
	Id           common.Long
	CategoryCode common.String
	CategoryName common.String
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

type QueryMarketCategoriesChildCategoryList []QueryMarketCategoriesChildCategory

func (list *QueryMarketCategoriesChildCategoryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMarketCategoriesChildCategory)
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
