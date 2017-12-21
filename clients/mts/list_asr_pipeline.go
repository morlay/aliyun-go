package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListAsrPipelineRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	State                string `position:"Query" name:"State"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
}

func (r ListAsrPipelineRequest) Invoke(client *sdk.Client) (response *ListAsrPipelineResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ListAsrPipelineRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Mts", "2014-06-18", "ListAsrPipeline", "", "")

	resp := struct {
		*responses.BaseResponse
		ListAsrPipelineResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ListAsrPipelineResponse

	err = client.DoAction(&req, &resp)
	return
}

type ListAsrPipelineResponse struct {
	RequestId    string
	TotalCount   int64
	PageNumber   int64
	PageSize     int64
	PipelineList ListAsrPipelinePipelineList
}

type ListAsrPipelinePipeline struct {
	Id           string
	Name         string
	State        string
	Priority     string
	NotifyConfig ListAsrPipelineNotifyConfig
}

type ListAsrPipelineNotifyConfig struct {
	Topic     string
	QueueName string
}

type ListAsrPipelinePipelineList []ListAsrPipelinePipeline

func (list *ListAsrPipelinePipelineList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAsrPipelinePipeline)
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
