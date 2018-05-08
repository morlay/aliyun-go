package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type QueryPornPipelineListRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	PipelineIds          string `position:"Query" name:"PipelineIds"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *QueryPornPipelineListRequest) Invoke(client *sdk.Client) (resp *QueryPornPipelineListResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "QueryPornPipelineList", "mts", "")
	resp = &QueryPornPipelineListResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryPornPipelineListResponse struct {
	responses.BaseResponse
	RequestId    common.String
	PipelineList QueryPornPipelineListPipelineList
	NonExistIds  QueryPornPipelineListNonExistIdList
}

type QueryPornPipelineListPipeline struct {
	Id           common.String
	Name         common.String
	State        common.String
	Priority     common.String
	NotifyConfig QueryPornPipelineListNotifyConfig
}

type QueryPornPipelineListNotifyConfig struct {
	Topic common.String
	Queue common.String
}

type QueryPornPipelineListPipelineList []QueryPornPipelineListPipeline

func (list *QueryPornPipelineListPipelineList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryPornPipelineListPipeline)
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

type QueryPornPipelineListNonExistIdList []common.String

func (list *QueryPornPipelineListNonExistIdList) UnmarshalJSON(data []byte) error {
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
