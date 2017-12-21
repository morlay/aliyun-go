package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListVideoSummaryPipelineRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	State                string `position:"Query" name:"State"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
}

func (r ListVideoSummaryPipelineRequest) Invoke(client *sdk.Client) (response *ListVideoSummaryPipelineResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ListVideoSummaryPipelineRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Mts", "2014-06-18", "ListVideoSummaryPipeline", "", "")

	resp := struct {
		*responses.BaseResponse
		ListVideoSummaryPipelineResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ListVideoSummaryPipelineResponse

	err = client.DoAction(&req, &resp)
	return
}

type ListVideoSummaryPipelineResponse struct {
	RequestId    string
	TotalCount   int64
	PageNumber   int64
	PageSize     int64
	PipelineList ListVideoSummaryPipelinePipelineList
}

type ListVideoSummaryPipelinePipeline struct {
	Id           string
	Name         string
	State        string
	Priority     string
	NotifyConfig ListVideoSummaryPipelineNotifyConfig
}

type ListVideoSummaryPipelineNotifyConfig struct {
	Topic     string
	QueueName string
}

type ListVideoSummaryPipelinePipelineList []ListVideoSummaryPipelinePipeline

func (list *ListVideoSummaryPipelinePipelineList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListVideoSummaryPipelinePipeline)
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
