package cms

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId    common.String
	Success      bool
	ErrorCode    common.Integer
	ErrorMessage common.String
	Category     ListMyGroupCategoriesCategory
}

type ListMyGroupCategoriesCategory struct {
	GroupId       common.Long
	CategoryItems ListMyGroupCategoriesCategoryItemList
}

type ListMyGroupCategoriesCategoryItem struct {
	Category common.String
	Count    common.Integer
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
