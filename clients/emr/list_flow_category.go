package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListFlowCategoryRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Root            string `position:"Query" name:"Root"`
	PageSize        int    `position:"Query" name:"PageSize"`
	ProjectId       string `position:"Query" name:"ProjectId"`
	ParentId        string `position:"Query" name:"ParentId"`
	PageNumber      int    `position:"Query" name:"PageNumber"`
}

func (req *ListFlowCategoryRequest) Invoke(client *sdk.Client) (resp *ListFlowCategoryResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ListFlowCategory", "", "")
	resp = &ListFlowCategoryResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListFlowCategoryResponse struct {
	responses.BaseResponse
	RequestId  common.String
	PageNumber common.Integer
	PageSize   common.Integer
	Total      common.Integer
	Categories ListFlowCategoryCategoryList
}

type ListFlowCategoryCategory struct {
	Id           common.String
	GmtCreate    common.Long
	GmtModified  common.Long
	Name         common.String
	ParentId     common.String
	Type         common.String
	CategoryType common.String
	ObjectType   common.String
	ObjectId     common.String
	ProjectId    common.String
}

type ListFlowCategoryCategoryList []ListFlowCategoryCategory

func (list *ListFlowCategoryCategoryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListFlowCategoryCategory)
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
