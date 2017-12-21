package ram

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListGroupsForUserRequest struct {
	UserName string `position:"Query" name:"UserName"`
}

func (r ListGroupsForUserRequest) Invoke(client *sdk.Client) (response *ListGroupsForUserResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ListGroupsForUserRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ram", "2015-05-01", "ListGroupsForUser", "", "")

	resp := struct {
		*responses.BaseResponse
		ListGroupsForUserResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ListGroupsForUserResponse

	err = client.DoAction(&req, &resp)
	return
}

type ListGroupsForUserResponse struct {
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
