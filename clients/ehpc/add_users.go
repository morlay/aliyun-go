package ehpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type AddUsersRequest struct {
	requests.RpcRequest
	ReleaseInstance string            `position:"Query" name:"ReleaseInstance"`
	ClusterId       string            `position:"Query" name:"ClusterId"`
	Users           *AddUsersUserList `position:"Query" type:"Repeated" name:"User"`
}

func (req *AddUsersRequest) Invoke(client *sdk.Client) (resp *AddUsersResponse, err error) {
	req.InitWithApiInfo("EHPC", "2018-04-12", "AddUsers", "ehs", "")
	resp = &AddUsersResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddUsersUser struct {
	Name     string `name:"Name"`
	Group    string `name:"Group"`
	Password string `name:"Password"`
}

type AddUsersResponse struct {
	responses.BaseResponse
	RequestId common.String
}

type AddUsersUserList []AddUsersUser

func (list *AddUsersUserList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]AddUsersUser)
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
