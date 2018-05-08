package imm

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListProjectsRequest struct {
	requests.RpcRequest
	MaxKeys int    `position:"Query" name:"MaxKeys"`
	Marker  string `position:"Query" name:"Marker"`
}

func (req *ListProjectsRequest) Invoke(client *sdk.Client) (resp *ListProjectsResponse, err error) {
	req.InitWithApiInfo("imm", "2017-09-06", "ListProjects", "imm", "")
	resp = &ListProjectsResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListProjectsResponse struct {
	responses.BaseResponse
	RequestId  common.String
	NextMarker common.String
	Projects   ListProjectsProjectsItemList
}

type ListProjectsProjectsItem struct {
	Project     common.String
	Endpoint    common.String
	ServiceRole common.String
	CreateTime  common.String
	ModifyTime  common.String
	Engines     ListProjectsEnginesItemList
	Indexers    ListProjectsIndexersItemList
}

type ListProjectsEnginesItem struct {
	Name   common.String
	JobTtl common.Long
}

type ListProjectsIndexersItem struct {
	Name   common.String
	Status common.String
}

type ListProjectsProjectsItemList []ListProjectsProjectsItem

func (list *ListProjectsProjectsItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListProjectsProjectsItem)
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

type ListProjectsEnginesItemList []ListProjectsEnginesItem

func (list *ListProjectsEnginesItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListProjectsEnginesItem)
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

type ListProjectsIndexersItemList []ListProjectsIndexersItem

func (list *ListProjectsIndexersItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListProjectsIndexersItem)
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
