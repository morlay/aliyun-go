package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId  common.String
	PageNumber common.Integer
	PageSize   common.Integer
	Total      common.Integer
	Projects   ListFlowProjectProjectList
}

type ListFlowProjectProject struct {
	Id          common.String
	GmtCreate   common.Long
	GmtModified common.Long
	UserId      common.String
	Name        common.String
	Description common.String
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
