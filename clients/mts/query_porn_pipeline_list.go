package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryPornPipelineListRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	PipelineIds          string `position:"Query" name:"PipelineIds"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *QueryPornPipelineListRequest) Invoke(client *sdk.Client) (resp *QueryPornPipelineListResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "QueryPornPipelineList", "mts", "")
	resp = &QueryPornPipelineListResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryPornPipelineListResponse struct {
	responses.BaseResponse
	RequestId    string
	PipelineList QueryPornPipelineListPipelineList
	NonExistIds  QueryPornPipelineListNonExistIdList
}

type QueryPornPipelineListPipeline struct {
	Id           string
	Name         string
	State        string
	Priority     string
	NotifyConfig QueryPornPipelineListNotifyConfig
}

type QueryPornPipelineListNotifyConfig struct {
	Topic string
	Queue string
}

type QueryPornPipelineListPipelineList []QueryPornPipelineListPipeline

func (list *QueryPornPipelineListPipelineList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryPornPipelineListPipeline)
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

type QueryPornPipelineListNonExistIdList []string

func (list *QueryPornPipelineListNonExistIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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
