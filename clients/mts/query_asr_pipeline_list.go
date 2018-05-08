package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type QueryAsrPipelineListRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	PipelineIds          string `position:"Query" name:"PipelineIds"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *QueryAsrPipelineListRequest) Invoke(client *sdk.Client) (resp *QueryAsrPipelineListResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "QueryAsrPipelineList", "mts", "")
	resp = &QueryAsrPipelineListResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryAsrPipelineListResponse struct {
	responses.BaseResponse
	RequestId    common.String
	PipelineList QueryAsrPipelineListPipelineList
	NonExistIds  QueryAsrPipelineListNonExistIdList
}

type QueryAsrPipelineListPipeline struct {
	Id           common.String
	Name         common.String
	State        common.String
	Priority     common.String
	NotifyConfig QueryAsrPipelineListNotifyConfig
}

type QueryAsrPipelineListNotifyConfig struct {
	Topic     common.String
	QueueName common.String
}

type QueryAsrPipelineListPipelineList []QueryAsrPipelineListPipeline

func (list *QueryAsrPipelineListPipelineList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryAsrPipelineListPipeline)
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

type QueryAsrPipelineListNonExistIdList []common.String

func (list *QueryAsrPipelineListNonExistIdList) UnmarshalJSON(data []byte) error {
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
