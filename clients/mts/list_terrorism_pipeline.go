package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListTerrorismPipelineRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	State                string `position:"Query" name:"State"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
}

func (req *ListTerrorismPipelineRequest) Invoke(client *sdk.Client) (resp *ListTerrorismPipelineResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "ListTerrorismPipeline", "", "")
	resp = &ListTerrorismPipelineResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListTerrorismPipelineResponse struct {
	responses.BaseResponse
	RequestId    string
	TotalCount   int64
	PageNumber   int64
	PageSize     int64
	PipelineList ListTerrorismPipelinePipelineList
}

type ListTerrorismPipelinePipeline struct {
	Id           string
	Name         string
	State        string
	Priority     string
	NotifyConfig ListTerrorismPipelineNotifyConfig
}

type ListTerrorismPipelineNotifyConfig struct {
	Topic string
	Queue string
}

type ListTerrorismPipelinePipelineList []ListTerrorismPipelinePipeline

func (list *ListTerrorismPipelinePipelineList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListTerrorismPipelinePipeline)
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
