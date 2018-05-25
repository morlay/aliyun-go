package ehpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyUserGroupsRequest struct {
	requests.RpcRequest
	ClusterId string                    `position:"Query" name:"ClusterId"`
	Users     *ModifyUserGroupsUserList `position:"Query" type:"Repeated" name:"User"`
}

func (req *ModifyUserGroupsRequest) Invoke(client *sdk.Client) (resp *ModifyUserGroupsResponse, err error) {
	req.InitWithApiInfo("EHPC", "2018-04-12", "ModifyUserGroups", "ehs", "")
	resp = &ModifyUserGroupsResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyUserGroupsUser struct {
	Name  string `name:"Name"`
	Group string `name:"Group"`
}

type ModifyUserGroupsResponse struct {
	responses.BaseResponse
	RequestId common.String
}

type ModifyUserGroupsUserList []ModifyUserGroupsUser

func (list *ModifyUserGroupsUserList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ModifyUserGroupsUser)
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
