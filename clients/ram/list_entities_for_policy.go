package ram

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListEntitiesForPolicyRequest struct {
	requests.RpcRequest
	PolicyType string `position:"Query" name:"PolicyType"`
	PolicyName string `position:"Query" name:"PolicyName"`
}

func (req *ListEntitiesForPolicyRequest) Invoke(client *sdk.Client) (resp *ListEntitiesForPolicyResponse, err error) {
	req.InitWithApiInfo("Ram", "2015-05-01", "ListEntitiesForPolicy", "", "")
	resp = &ListEntitiesForPolicyResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListEntitiesForPolicyResponse struct {
	responses.BaseResponse
	RequestId common.String
	Groups    ListEntitiesForPolicyGroupList
	Users     ListEntitiesForPolicyUserList
	Roles     ListEntitiesForPolicyRoleList
}

type ListEntitiesForPolicyGroup struct {
	GroupName  common.String
	Comments   common.String
	AttachDate common.String
}

type ListEntitiesForPolicyUser struct {
	UserId      common.String
	UserName    common.String
	DisplayName common.String
	AttachDate  common.String
}

type ListEntitiesForPolicyRole struct {
	RoleId      common.String
	RoleName    common.String
	Arn         common.String
	Description common.String
	AttachDate  common.String
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
