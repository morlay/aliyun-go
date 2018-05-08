package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListFlowProjectUserRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	PageSize        int    `position:"Query" name:"PageSize"`
	ProjectId       string `position:"Query" name:"ProjectId"`
	PageNumber      int    `position:"Query" name:"PageNumber"`
}

func (req *ListFlowProjectUserRequest) Invoke(client *sdk.Client) (resp *ListFlowProjectUserResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ListFlowProjectUser", "", "")
	resp = &ListFlowProjectUserResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListFlowProjectUserResponse struct {
	responses.BaseResponse
	RequestId  common.String
	PageNumber common.Integer
	PageSize   common.Integer
	Total      common.Integer
	Users      ListFlowProjectUserUserList
}

type ListFlowProjectUserUser struct {
	GmtCreate   common.Long
	GmtModified common.Long
	ProjectId   common.String
	OwnerId     common.String
	UserName    common.String
}

type ListFlowProjectUserUserList []ListFlowProjectUserUser

func (list *ListFlowProjectUserUserList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListFlowProjectUserUser)
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
