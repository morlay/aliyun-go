package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SearchPipelineRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	State                string `position:"Query" name:"State"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
}

func (r SearchPipelineRequest) Invoke(client *sdk.Client) (response *SearchPipelineResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SearchPipelineRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Mts", "2014-06-18", "SearchPipeline", "", "")

	resp := struct {
		*responses.BaseResponse
		SearchPipelineResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SearchPipelineResponse

	err = client.DoAction(&req, &resp)
	return
}

type SearchPipelineResponse struct {
	RequestId    string
	TotalCount   int64
	PageNumber   int64
	PageSize     int64
	PipelineList SearchPipelinePipelineList
}

type SearchPipelinePipeline struct {
	Id           string
	Name         string
	State        string
	Speed        string
	SpeedLevel   int64
	Role         string
	NotifyConfig SearchPipelineNotifyConfig
}

type SearchPipelineNotifyConfig struct {
	Topic     string
	QueueName string
}

type SearchPipelinePipelineList []SearchPipelinePipeline

func (list *SearchPipelinePipelineList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SearchPipelinePipeline)
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
