package ram

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListGroupsForUserRequest struct {
	requests.RpcRequest
	UserName string `position:"Query" name:"UserName"`
}

func (req *ListGroupsForUserRequest) Invoke(client *sdk.Client) (resp *ListGroupsForUserResponse, err error) {
	req.InitWithApiInfo("Ram", "2015-05-01", "ListGroupsForUser", "", "")
	resp = &ListGroupsForUserResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListGroupsForUserResponse struct {
	responses.BaseResponse
	RequestId string
	Groups    ListGroupsForUserGroupList
}

type ListGroupsForUserGroup struct {
	GroupName string
	Comments  string
	JoinDate  string
}

type ListGroupsForUserGroupList []ListGroupsForUserGroup

func (list *ListGroupsForUserGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListGroupsForUserGroup)
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
