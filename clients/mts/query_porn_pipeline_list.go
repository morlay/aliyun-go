package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryPornPipelineListRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	PipelineIds          string `position:"Query" name:"PipelineIds"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r QueryPornPipelineListRequest) Invoke(client *sdk.Client) (response *QueryPornPipelineListResponse, err error) {
	req := struct {
		*requests.RpcRequest
		QueryPornPipelineListRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Mts", "2014-06-18", "QueryPornPipelineList", "", "")

	resp := struct {
		*responses.BaseResponse
		QueryPornPipelineListResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.QueryPornPipelineListResponse

	err = client.DoAction(&req, &resp)
	return
}

type QueryPornPipelineListResponse struct {
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
