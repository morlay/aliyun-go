package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type QueryMediaWorkflowListRequest struct {
	requests.RpcRequest
	MediaWorkflowIds     string `position:"Query" name:"MediaWorkflowIds"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *QueryMediaWorkflowListRequest) Invoke(client *sdk.Client) (resp *QueryMediaWorkflowListResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "QueryMediaWorkflowList", "mts", "")
	resp = &QueryMediaWorkflowListResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryMediaWorkflowListResponse struct {
	responses.BaseResponse
	RequestId                common.String
	MediaWorkflowList        QueryMediaWorkflowListMediaWorkflowList
	NonExistMediaWorkflowIds QueryMediaWorkflowListNonExistMediaWorkflowIdList
}

type QueryMediaWorkflowListMediaWorkflow struct {
	MediaWorkflowId common.String
	Name            common.String
	Topology        common.String
	TriggerMode     common.String
	State           common.String
	CreationTime    common.String
}

type QueryMediaWorkflowListMediaWorkflowList []QueryMediaWorkflowListMediaWorkflow

func (list *QueryMediaWorkflowListMediaWorkflowList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMediaWorkflowListMediaWorkflow)
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

type QueryMediaWorkflowListNonExistMediaWorkflowIdList []common.String

func (list *QueryMediaWorkflowListNonExistMediaWorkflowIdList) UnmarshalJSON(data []byte) error {
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
