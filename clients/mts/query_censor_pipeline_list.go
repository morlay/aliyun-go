package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type QueryCensorPipelineListRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	PipelineIds          string `position:"Query" name:"PipelineIds"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *QueryCensorPipelineListRequest) Invoke(client *sdk.Client) (resp *QueryCensorPipelineListResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "QueryCensorPipelineList", "mts", "")
	resp = &QueryCensorPipelineListResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryCensorPipelineListResponse struct {
	responses.BaseResponse
	RequestId    common.String
	PipelineList QueryCensorPipelineListPipelineList
	NonExistIds  QueryCensorPipelineListNonExistIdList
}

type QueryCensorPipelineListPipeline struct {
	Id           common.String
	Name         common.String
	State        common.String
	Priority     common.String
	NotifyConfig QueryCensorPipelineListNotifyConfig
}

type QueryCensorPipelineListNotifyConfig struct {
	Topic common.String
	Queue common.String
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

type QueryCensorPipelineListNonExistIdList []common.String

func (list *QueryCensorPipelineListNonExistIdList) UnmarshalJSON(data []byte) error {
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
