package imm

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetProjectRequest struct {
	requests.RpcRequest
	Project string `position:"Query" name:"Project"`
}

func (req *GetProjectRequest) Invoke(client *sdk.Client) (resp *GetProjectResponse, err error) {
	req.InitWithApiInfo("imm", "2017-09-06", "GetProject", "imm", "")
	resp = &GetProjectResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetProjectResponse struct {
	responses.BaseResponse
	RequestId   string
	Project     string
	ServiceRole string
	Endpoint    string
	CreateTime  string
	ModifyTime  string
	Indexers    GetProjectIndexersItemList
	Engines     GetProjectEnginesItemList
}

type GetProjectIndexersItem struct {
	Name   string
	Status string
}

type GetProjectEnginesItem struct {
	Name   string
	JobTtl int64
}

type GetProjectIndexersItemList []GetProjectIndexersItem

func (list *GetProjectIndexersItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetProjectIndexersItem)
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

type GetProjectEnginesItemList []GetProjectEnginesItem

func (list *GetProjectEnginesItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetProjectEnginesItem)
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
