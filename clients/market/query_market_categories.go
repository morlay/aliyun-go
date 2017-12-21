package market

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryMarketCategoriesRequest struct {
}

func (r QueryMarketCategoriesRequest) Invoke(client *sdk.Client) (response *QueryMarketCategoriesResponse, err error) {
	req := struct {
		*requests.RpcRequest
		QueryMarketCategoriesRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Market", "2015-11-01", "QueryMarketCategories", "", "")

	resp := struct {
		*responses.BaseResponse
		QueryMarketCategoriesResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.QueryMarketCategoriesResponse

	err = client.DoAction(&req, &resp)
	return
}

type QueryMarketCategoriesResponse struct {
	PageNumber int
	PageSize   int
	TotalCount int
	RequestId  string
	Categories QueryMarketCategoriesCategoryList
}

type QueryMarketCategoriesCategory struct {
	Id           int64
	CategoryCode string
	CategoryName string
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
