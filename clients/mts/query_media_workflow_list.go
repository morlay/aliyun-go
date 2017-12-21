package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryMediaWorkflowListRequest struct {
	MediaWorkflowIds     string `position:"Query" name:"MediaWorkflowIds"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r QueryMediaWorkflowListRequest) Invoke(client *sdk.Client) (response *QueryMediaWorkflowListResponse, err error) {
	req := struct {
		*requests.RpcRequest
		QueryMediaWorkflowListRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Mts", "2014-06-18", "QueryMediaWorkflowList", "", "")

	resp := struct {
		*responses.BaseResponse
		QueryMediaWorkflowListResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.QueryMediaWorkflowListResponse

	err = client.DoAction(&req, &resp)
	return
}

type QueryMediaWorkflowListResponse struct {
	RequestId                string
	MediaWorkflowList        QueryMediaWorkflowListMediaWorkflowList
	NonExistMediaWorkflowIds QueryMediaWorkflowListNonExistMediaWorkflowIdList
}

type QueryMediaWorkflowListMediaWorkflow struct {
	MediaWorkflowId string
	Name            string
	Topology        string
	State           string
	CreationTime    string
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

type QueryMediaWorkflowListNonExistMediaWorkflowIdList []string

func (list *QueryMediaWorkflowListNonExistMediaWorkflowIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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
