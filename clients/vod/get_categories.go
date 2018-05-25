package vod

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetCategoriesRequest struct {
	requests.RpcRequest
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	CateId               int64  `position:"Query" name:"CateId"`
	PageNo               int64  `position:"Query" name:"PageNo"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	OwnerId              string `position:"Query" name:"OwnerId"`
}

func (req *GetCategoriesRequest) Invoke(client *sdk.Client) (resp *GetCategoriesResponse, err error) {
	req.InitWithApiInfo("vod", "2017-03-21", "GetCategories", "vod", "")
	resp = &GetCategoriesResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetCategoriesResponse struct {
	responses.BaseResponse
	RequestId     common.String
	SubTotal      common.Long
	SubCategories GetCategoriesCategoryList
	Category1     GetCategoriesCategory1
}

type GetCategoriesCategory struct {
	CateId   common.Long
	CateName common.String
	Level    common.Long
	ParentId common.Long
}

type GetCategoriesCategory1 struct {
	CateId   common.Long
	CateName common.String
	Level    common.Long
	ParentId common.Long
}

type GetCategoriesCategoryList []GetCategoriesCategory

func (list *GetCategoriesCategoryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetCategoriesCategory)
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
