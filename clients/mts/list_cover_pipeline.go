package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	req.InitWithApiInfo("Mts", "2014-06-18", "ListCoverPipeline", "", "")
	resp = &ListCoverPipelineResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListCoverPipelineResponse struct {
	responses.BaseResponse
	RequestId    string
	TotalCount   int64
	PageNumber   int64
	PageSize     int64
	PipelineList ListCoverPipelinePipelineList
}

type ListCoverPipelinePipeline struct {
	UserId       int64
	PipelineId   string
	Name         string
	State        string
	Priority     string
	QuotaNum     int
	QuotaUsed    int
	NotifyConfig string
	Role         string
	ExtendConfig string
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
