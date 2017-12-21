package ram

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListUsersForGroupRequest struct {
	GroupName string `position:"Query" name:"GroupName"`
}

func (r ListUsersForGroupRequest) Invoke(client *sdk.Client) (response *ListUsersForGroupResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ListUsersForGroupRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ram", "2015-05-01", "ListUsersForGroup", "", "")

	resp := struct {
		*responses.BaseResponse
		ListUsersForGroupResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ListUsersForGroupResponse

	err = client.DoAction(&req, &resp)
	return
}

type ListUsersForGroupResponse struct {
	RequestId string
	Users     ListUsersForGroupUserList
}

type ListUsersForGroupUser struct {
	UserName    string
	DisplayName string
	JoinDate    string
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
