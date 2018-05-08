package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListAllCategoryRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *ListAllCategoryRequest) Invoke(client *sdk.Client) (resp *ListAllCategoryResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "ListAllCategory", "mts", "")
	resp = &ListAllCategoryResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListAllCategoryResponse struct {
	responses.BaseResponse
	RequestId    common.String
	CategoryList ListAllCategoryCategoryList
}

type ListAllCategoryCategory struct {
	CateId   common.String
	CateName common.String
	ParentId common.String
	Level    common.String
}

type ListAllCategoryCategoryList []ListAllCategoryCategory

func (list *ListAllCategoryCategoryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAllCategoryCategory)
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
