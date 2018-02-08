package ehpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyUserPasswordsRequest struct {
	requests.RpcRequest
	ClusterId string                       `position:"Query" name:"ClusterId"`
	Users     *ModifyUserPasswordsUserList `position:"Query" type:"Repeated" name:"User"`
}

func (req *ModifyUserPasswordsRequest) Invoke(client *sdk.Client) (resp *ModifyUserPasswordsResponse, err error) {
	req.InitWithApiInfo("EHPC", "2017-07-14", "ModifyUserPasswords", "ehs", "")
	resp = &ModifyUserPasswordsResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyUserPasswordsUser struct {
	Name     string `name:"Name"`
	Password string `name:"Password"`
}

type ModifyUserPasswordsResponse struct {
	responses.BaseResponse
	RequestId string
}

type ModifyUserPasswordsUserList []ModifyUserPasswordsUser

func (list *ModifyUserPasswordsUserList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ModifyUserPasswordsUser)
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
