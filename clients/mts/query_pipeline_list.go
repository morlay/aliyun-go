package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryPipelineListRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	PipelineIds          string `position:"Query" name:"PipelineIds"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r QueryPipelineListRequest) Invoke(client *sdk.Client) (response *QueryPipelineListResponse, err error) {
	req := struct {
		*requests.RpcRequest
		QueryPipelineListRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Mts", "2014-06-18", "QueryPipelineList", "", "")

	resp := struct {
		*responses.BaseResponse
		QueryPipelineListResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.QueryPipelineListResponse

	err = client.DoAction(&req, &resp)
	return
}

type QueryPipelineListResponse struct {
	RequestId    string
	PipelineList QueryPipelineListPipelineList
	NonExistPids QueryPipelineListNonExistPidList
}

type QueryPipelineListPipeline struct {
	Id           string
	Name         string
	State        string
	Speed        string
	SpeedLevel   int64
	Role         string
	NotifyConfig QueryPipelineListNotifyConfig
}

type QueryPipelineListNotifyConfig struct {
	Topic     string
	QueueName string
}

type QueryPipelineListPipelineList []QueryPipelineListPipeline

func (list *QueryPipelineListPipelineList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryPipelineListPipeline)
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

type QueryPipelineListNonExistPidList []string

func (list *QueryPipelineListNonExistPidList) UnmarshalJSON(data []byte) error {
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
