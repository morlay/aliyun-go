package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListFlowProjectRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Name            string `position:"Query" name:"Name"`
	PageSize        int    `position:"Query" name:"PageSize"`
	ProjectId       string `position:"Query" name:"ProjectId"`
	PageNumber      int    `position:"Query" name:"PageNumber"`
}

func (req *ListFlowProjectRequest) Invoke(client *sdk.Client) (resp *ListFlowProjectResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ListFlowProject", "", "")
	resp = &ListFlowProjectResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListFlowProjectResponse struct {
	responses.BaseResponse
	RequestId  string
	PageNumber int
	PageSize   int
	Total      int
	Projects   ListFlowProjectProjectList
}

type ListFlowProjectProject struct {
	Id          string
	GmtCreate   int64
	GmtModified int64
	UserId      string
	Name        string
	Description string
}

type ListFlowProjectProjectList []ListFlowProjectProject

func (list *ListFlowProjectProjectList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListFlowProjectProject)
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
