package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SearchPipelineRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	State                string `position:"Query" name:"State"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
}

func (req *SearchPipelineRequest) Invoke(client *sdk.Client) (resp *SearchPipelineResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "SearchPipeline", "mts", "")
	resp = &SearchPipelineResponse{}
	err = client.DoAction(req, resp)
	return
}

type SearchPipelineResponse struct {
	responses.BaseResponse
	RequestId    common.String
	TotalCount   common.Long
	PageNumber   common.Long
	PageSize     common.Long
	PipelineList SearchPipelinePipelineList
}

type SearchPipelinePipeline struct {
	Id           common.String
	Name         common.String
	State        common.String
	Speed        common.String
	SpeedLevel   common.Long
	Role         common.String
	NotifyConfig SearchPipelineNotifyConfig
}

type SearchPipelineNotifyConfig struct {
	Topic     common.String
	QueueName common.String
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
