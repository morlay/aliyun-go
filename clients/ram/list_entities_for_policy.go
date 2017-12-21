package ram

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListEntitiesForPolicyRequest struct {
	PolicyType string `position:"Query" name:"PolicyType"`
	PolicyName string `position:"Query" name:"PolicyName"`
}

func (r ListEntitiesForPolicyRequest) Invoke(client *sdk.Client) (response *ListEntitiesForPolicyResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ListEntitiesForPolicyRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ram", "2015-05-01", "ListEntitiesForPolicy", "", "")

	resp := struct {
		*responses.BaseResponse
		ListEntitiesForPolicyResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ListEntitiesForPolicyResponse

	err = client.DoAction(&req, &resp)
	return
}

type ListEntitiesForPolicyResponse struct {
	RequestId string
	Groups    ListEntitiesForPolicyGroupList
	Users     ListEntitiesForPolicyUserList
	Roles     ListEntitiesForPolicyRoleList
}

type ListEntitiesForPolicyGroup struct {
	GroupName  string
	Comments   string
	AttachDate string
}

type ListEntitiesForPolicyUser struct {
	UserId      string
	UserName    string
	DisplayName string
	AttachDate  string
}

type ListEntitiesForPolicyRole struct {
	RoleId      string
	RoleName    string
	Arn         string
	Description string
	AttachDate  string
}

type ListEntitiesForPolicyGroupList []ListEntitiesForPolicyGroup

func (list *ListEntitiesForPolicyGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListEntitiesForPolicyGroup)
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

type ListEntitiesForPolicyUserList []ListEntitiesForPolicyUser

func (list *ListEntitiesForPolicyUserList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListEntitiesForPolicyUser)
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

type ListEntitiesForPolicyRoleList []ListEntitiesForPolicyRole

func (list *ListEntitiesForPolicyRoleList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListEntitiesForPolicyRole)
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
