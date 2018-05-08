package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListVideoSummaryPipelineRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	State                string `position:"Query" name:"State"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
}

func (req *ListVideoSummaryPipelineRequest) Invoke(client *sdk.Client) (resp *ListVideoSummaryPipelineResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "ListVideoSummaryPipeline", "mts", "")
	resp = &ListVideoSummaryPipelineResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListVideoSummaryPipelineResponse struct {
	responses.BaseResponse
	RequestId    common.String
	TotalCount   common.Long
	PageNumber   common.Long
	PageSize     common.Long
	PipelineList ListVideoSummaryPipelinePipelineList
}

type ListVideoSummaryPipelinePipeline struct {
	Id           common.String
	Name         common.String
	State        common.String
	Priority     common.String
	NotifyConfig ListVideoSummaryPipelineNotifyConfig
}

type ListVideoSummaryPipelineNotifyConfig struct {
	Topic     common.String
	QueueName common.String
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
