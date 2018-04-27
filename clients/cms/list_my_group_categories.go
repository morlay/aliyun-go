package cms

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListMyGroupCategoriesRequest struct {
	requests.RpcRequest
	GroupId int64 `position:"Query" name:"GroupId"`
}

func (req *ListMyGroupCategoriesRequest) Invoke(client *sdk.Client) (resp *ListMyGroupCategoriesResponse, err error) {
	req.InitWithApiInfo("Cms", "2018-03-08", "ListMyGroupCategories", "cms", "")
	resp = &ListMyGroupCategoriesResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListMyGroupCategoriesResponse struct {
	responses.BaseResponse
	RequestId    string
	Success      bool
	ErrorCode    int
	ErrorMessage string
	Category     ListMyGroupCategoriesCategory
}

type ListMyGroupCategoriesCategory struct {
	GroupId       int64
	CategoryItems ListMyGroupCategoriesCategoryItemList
}

type ListMyGroupCategoriesCategoryItem struct {
	Category string
	Count    int
}

type ListMyGroupCategoriesCategoryItemList []ListMyGroupCategoriesCategoryItem

func (list *ListMyGroupCategoriesCategoryItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListMyGroupCategoriesCategoryItem)
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
