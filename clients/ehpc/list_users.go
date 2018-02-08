package ehpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListUsersRequest struct {
	requests.RpcRequest
	PageSize   int    `position:"Query" name:"PageSize"`
	ClusterId  string `position:"Query" name:"ClusterId"`
	PageNumber int    `position:"Query" name:"PageNumber"`
}

func (req *ListUsersRequest) Invoke(client *sdk.Client) (resp *ListUsersResponse, err error) {
	req.InitWithApiInfo("EHPC", "2017-07-14", "ListUsers", "ehs", "")
	resp = &ListUsersResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListUsersResponse struct {
	responses.BaseResponse
	RequestId  string
	TotalCount int
	PageNumber int
	PageSize   int
	Users      ListUsersUserInfoList
}

type ListUsersUserInfo struct {
	Name    string
	Group   string
	AddTime string
}

type ListUsersUserInfoList []ListUsersUserInfo

func (list *ListUsersUserInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListUsersUserInfo)
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
