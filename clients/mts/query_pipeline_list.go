package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type QueryPipelineListRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	PipelineIds          string `position:"Query" name:"PipelineIds"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *QueryPipelineListRequest) Invoke(client *sdk.Client) (resp *QueryPipelineListResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "QueryPipelineList", "mts", "")
	resp = &QueryPipelineListResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryPipelineListResponse struct {
	responses.BaseResponse
	RequestId    common.String
	PipelineList QueryPipelineListPipelineList
	NonExistPids QueryPipelineListNonExistPidList
}

type QueryPipelineListPipeline struct {
	Id           common.String
	Name         common.String
	State        common.String
	Speed        common.String
	SpeedLevel   common.Long
	Role         common.String
	NotifyConfig QueryPipelineListNotifyConfig
}

type QueryPipelineListNotifyConfig struct {
	Topic     common.String
	QueueName common.String
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

type QueryPipelineListNonExistPidList []common.String

func (list *QueryPipelineListNonExistPidList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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
