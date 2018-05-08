package ram

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListUsersForGroupRequest struct {
	requests.RpcRequest
	GroupName string `position:"Query" name:"GroupName"`
}

func (req *ListUsersForGroupRequest) Invoke(client *sdk.Client) (resp *ListUsersForGroupResponse, err error) {
	req.InitWithApiInfo("Ram", "2015-05-01", "ListUsersForGroup", "", "")
	resp = &ListUsersForGroupResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListUsersForGroupResponse struct {
	responses.BaseResponse
	RequestId common.String
	Users     ListUsersForGroupUserList
}

type ListUsersForGroupUser struct {
	UserName    common.String
	DisplayName common.String
	JoinDate    common.String
}

type ListUsersForGroupUserList []ListUsersForGroupUser

func (list *ListUsersForGroupUserList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListUsersForGroupUser)
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
