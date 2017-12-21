package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListCensorPipelineRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	State                string `position:"Query" name:"State"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
}

func (req *ListCensorPipelineRequest) Invoke(client *sdk.Client) (resp *ListCensorPipelineResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "ListCensorPipeline", "", "")
	resp = &ListCensorPipelineResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListCensorPipelineResponse struct {
	responses.BaseResponse
	RequestId    string
	TotalCount   int64
	PageNumber   int64
	PageSize     int64
	PipelineList ListCensorPipelinePipelineList
}

type ListCensorPipelinePipeline struct {
	Id           string
	Name         string
	State        string
	Priority     string
	NotifyConfig ListCensorPipelineNotifyConfig
}

type ListCensorPipelineNotifyConfig struct {
	Topic string
	Queue string
}

type ListCensorPipelinePipelineList []ListCensorPipelinePipeline

func (list *ListCensorPipelinePipelineList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListCensorPipelinePipeline)
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
