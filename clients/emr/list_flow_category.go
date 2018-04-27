package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId  string
	PageNumber int
	PageSize   int
	Total      int
	Categories ListFlowCategoryCategoryList
}

type ListFlowCategoryCategory struct {
	Id           string
	GmtCreate    int64
	GmtModified  int64
	Name         string
	ParentId     string
	Type         string
	CategoryType string
	ObjectType   string
	ObjectId     string
	ProjectId    string
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
