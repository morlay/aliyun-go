package imm

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type PutProjectRequest struct {
	requests.RpcRequest
	Indexers    string `position:"Query" name:"Indexers"`
	Engines     string `position:"Query" name:"Engines"`
	ServiceRole string `position:"Query" name:"ServiceRole"`
	Project     string `position:"Query" name:"Project"`
}

func (req *PutProjectRequest) Invoke(client *sdk.Client) (resp *PutProjectResponse, err error) {
	req.InitWithApiInfo("imm", "2017-09-06", "PutProject", "imm", "")
	resp = &PutProjectResponse{}
	err = client.DoAction(req, resp)
	return
}

type PutProjectResponse struct {
	responses.BaseResponse
	RequestId   string
	Project     string
	CreateTime  string
	ModifyTime  string
	ServiceRole string
	Engines     PutProjectEnginesItemList
	Indexers    PutProjectIndexersItemList
}

type PutProjectEnginesItem struct {
	Name   string
	JobTtl int64
}

type PutProjectIndexersItem struct {
	Name   string
	Status string
}

type PutProjectEnginesItemList []PutProjectEnginesItem

func (list *PutProjectEnginesItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]PutProjectEnginesItem)
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

type PutProjectIndexersItemList []PutProjectIndexersItem

func (list *PutProjectIndexersItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]PutProjectIndexersItem)
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
