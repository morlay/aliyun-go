package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListCoverPipelineRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	State                string `position:"Query" name:"State"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
}

func (req *ListCoverPipelineRequest) Invoke(client *sdk.Client) (resp *ListCoverPipelineResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "ListCoverPipeline", "mts", "")
	resp = &ListCoverPipelineResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListCoverPipelineResponse struct {
	responses.BaseResponse
	RequestId    common.String
	TotalCount   common.Long
	PageNumber   common.Long
	PageSize     common.Long
	PipelineList ListCoverPipelinePipelineList
}

type ListCoverPipelinePipeline struct {
	UserId       common.Long
	PipelineId   common.String
	Name         common.String
	State        common.String
	Priority     common.String
	QuotaNum     common.Integer
	QuotaUsed    common.Integer
	NotifyConfig common.String
	Role         common.String
	ExtendConfig common.String
}

type ListCoverPipelinePipelineList []ListCoverPipelinePipeline

func (list *ListCoverPipelinePipelineList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListCoverPipelinePipeline)
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
