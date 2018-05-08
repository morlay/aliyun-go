package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type QueryTerrorismPipelineListRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	PipelineIds          string `position:"Query" name:"PipelineIds"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *QueryTerrorismPipelineListRequest) Invoke(client *sdk.Client) (resp *QueryTerrorismPipelineListResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "QueryTerrorismPipelineList", "mts", "")
	resp = &QueryTerrorismPipelineListResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryTerrorismPipelineListResponse struct {
	responses.BaseResponse
	RequestId    common.String
	PipelineList QueryTerrorismPipelineListPipelineList
	NonExistIds  QueryTerrorismPipelineListNonExistIdList
}

type QueryTerrorismPipelineListPipeline struct {
	Id           common.String
	Name         common.String
	State        common.String
	Priority     common.String
	NotifyConfig QueryTerrorismPipelineListNotifyConfig
}

type QueryTerrorismPipelineListNotifyConfig struct {
	Topic common.String
	Queue common.String
}

type QueryTerrorismPipelineListPipelineList []QueryTerrorismPipelineListPipeline

func (list *QueryTerrorismPipelineListPipelineList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryTerrorismPipelineListPipeline)
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

type QueryTerrorismPipelineListNonExistIdList []common.String

func (list *QueryTerrorismPipelineListNonExistIdList) UnmarshalJSON(data []byte) error {
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
