package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SearchMediaWorkflowRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	StateList            string `position:"Query" name:"StateList"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
}

func (req *SearchMediaWorkflowRequest) Invoke(client *sdk.Client) (resp *SearchMediaWorkflowResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "SearchMediaWorkflow", "mts", "")
	resp = &SearchMediaWorkflowResponse{}
	err = client.DoAction(req, resp)
	return
}

type SearchMediaWorkflowResponse struct {
	responses.BaseResponse
	RequestId         common.String
	TotalCount        common.Long
	PageNumber        common.Long
	PageSize          common.Long
	MediaWorkflowList SearchMediaWorkflowMediaWorkflowList
}

type SearchMediaWorkflowMediaWorkflow struct {
	MediaWorkflowId common.String
	Name            common.String
	Topology        common.String
	TriggerMode     common.String
	State           common.String
	CreationTime    common.String
}

type SearchMediaWorkflowMediaWorkflowList []SearchMediaWorkflowMediaWorkflow

func (list *SearchMediaWorkflowMediaWorkflowList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SearchMediaWorkflowMediaWorkflow)
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
