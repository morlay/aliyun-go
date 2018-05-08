package vod

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	req.InitWithApiInfo("vod", "2017-03-21", "SearchEditingProject", "vod", "")
	resp = &SearchEditingProjectResponse{}
	err = client.DoAction(req, resp)
	return
}

type SearchEditingProjectResponse struct {
	responses.BaseResponse
	RequestId   common.String
	Total       common.Integer
	ProjectList SearchEditingProjectProjectList
}

type SearchEditingProjectProject struct {
	ProjectId    common.String
	CreationTime common.String
	ModifiedTime common.String
	Status       common.String
	Description  common.String
	Title        common.String
	CoverURL     common.String
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
