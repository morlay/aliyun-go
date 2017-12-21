package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryAsrPipelineListRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	PipelineIds          string `position:"Query" name:"PipelineIds"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *QueryAsrPipelineListRequest) Invoke(client *sdk.Client) (resp *QueryAsrPipelineListResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "QueryAsrPipelineList", "", "")
	resp = &QueryAsrPipelineListResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryAsrPipelineListResponse struct {
	responses.BaseResponse
	RequestId    string
	PipelineList QueryAsrPipelineListPipelineList
	NonExistIds  QueryAsrPipelineListNonExistIdList
}

type QueryAsrPipelineListPipeline struct {
	Id           string
	Name         string
	State        string
	Priority     string
	NotifyConfig QueryAsrPipelineListNotifyConfig
}

type QueryAsrPipelineListNotifyConfig struct {
	Topic     string
	QueueName string
}

type QueryAsrPipelineListPipelineList []QueryAsrPipelineListPipeline

func (list *QueryAsrPipelineListPipelineList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryAsrPipelineListPipeline)
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

type QueryAsrPipelineListNonExistIdList []string

func (list *QueryAsrPipelineListNonExistIdList) UnmarshalJSON(data []byte) error {
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
