package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryCensorPipelineListRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	PipelineIds          string `position:"Query" name:"PipelineIds"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r QueryCensorPipelineListRequest) Invoke(client *sdk.Client) (response *QueryCensorPipelineListResponse, err error) {
	req := struct {
		*requests.RpcRequest
		QueryCensorPipelineListRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Mts", "2014-06-18", "QueryCensorPipelineList", "", "")

	resp := struct {
		*responses.BaseResponse
		QueryCensorPipelineListResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.QueryCensorPipelineListResponse

	err = client.DoAction(&req, &resp)
	return
}

type QueryCensorPipelineListResponse struct {
	RequestId    string
	PipelineList QueryCensorPipelineListPipelineList
	NonExistIds  QueryCensorPipelineListNonExistIdList
}

type QueryCensorPipelineListPipeline struct {
	Id           string
	Name         string
	State        string
	Priority     string
	NotifyConfig QueryCensorPipelineListNotifyConfig
}

type QueryCensorPipelineListNotifyConfig struct {
	Topic string
	Queue string
}

type QueryCensorPipelineListPipelineList []QueryCensorPipelineListPipeline

func (list *QueryCensorPipelineListPipelineList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryCensorPipelineListPipeline)
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

type QueryCensorPipelineListNonExistIdList []string

func (list *QueryCensorPipelineListNonExistIdList) UnmarshalJSON(data []byte) error {
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
