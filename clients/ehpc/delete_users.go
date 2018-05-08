package ehpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteUsersRequest struct {
	requests.RpcRequest
	ClusterId string               `position:"Query" name:"ClusterId"`
	Users     *DeleteUsersUserList `position:"Query" type:"Repeated" name:"User"`
}

func (req *DeleteUsersRequest) Invoke(client *sdk.Client) (resp *DeleteUsersResponse, err error) {
	req.InitWithApiInfo("EHPC", "2017-07-14", "DeleteUsers", "ehs", "")
	resp = &DeleteUsersResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteUsersUser struct {
	Name string `name:"Name"`
}

type DeleteUsersResponse struct {
	responses.BaseResponse
	RequestId common.String
}

type DeleteUsersUserList []DeleteUsersUser

func (list *DeleteUsersUserList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DeleteUsersUser)
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
