package vod

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SearchEditingProjectRequest struct {
	requests.RpcRequest
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	EndTime              string `position:"Query" name:"EndTime"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	Title                string `position:"Query" name:"Title"`
	PageNo               int    `position:"Query" name:"PageNo"`
	PageSize             int    `position:"Query" name:"PageSize"`
	SortBy               string `position:"Query" name:"SortBy"`
	Status               string `position:"Query" name:"Status"`
}

func (req *SearchEditingProjectRequest) Invoke(client *sdk.Client) (resp *SearchEditingProjectResponse, err error) {
	req.InitWithApiInfo("vod", "2017-03-21", "SearchEditingProject", "", "")
	resp = &SearchEditingProjectResponse{}
	err = client.DoAction(req, resp)
	return
}

type SearchEditingProjectResponse struct {
	responses.BaseResponse
	RequestId   string
	Total       int
	ProjectList SearchEditingProjectProjectList
}

type SearchEditingProjectProject struct {
	ProjectId    string
	CreationTime string
	ModifiedTime string
	Status       string
	Description  string
	Title        string
	CoverURL     string
}

type SearchEditingProjectProjectList []SearchEditingProjectProject

func (list *SearchEditingProjectProjectList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SearchEditingProjectProject)
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
