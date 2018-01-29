package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryCoverPipelineListRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	PipelineIds          string `position:"Query" name:"PipelineIds"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *QueryCoverPipelineListRequest) Invoke(client *sdk.Client) (resp *QueryCoverPipelineListResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "QueryCoverPipelineList", "mts", "")
	resp = &QueryCoverPipelineListResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryCoverPipelineListResponse struct {
	responses.BaseResponse
	RequestId    string
	PipelineList QueryCoverPipelineListPipelineList
	NonExistIds  QueryCoverPipelineListNonExistIdList
}

type QueryCoverPipelineListPipeline struct {
	Id           string
	Name         string
	State        string
	Priority     string
	Role         string
	NotifyConfig QueryCoverPipelineListNotifyConfig
}

type QueryCoverPipelineListNotifyConfig struct {
	Topic string
	Queue string
}

type QueryCoverPipelineListPipelineList []QueryCoverPipelineListPipeline

func (list *QueryCoverPipelineListPipelineList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryCoverPipelineListPipeline)
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

type QueryCoverPipelineListNonExistIdList []string

func (list *QueryCoverPipelineListNonExistIdList) UnmarshalJSON(data []byte) error {
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
