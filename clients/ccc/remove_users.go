package ccc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RemoveUsersRequest struct {
	requests.RpcRequest
	InstanceId string                 `position:"Query" name:"InstanceId"`
	UserIds    *RemoveUsersUserIdList `position:"Query" type:"Repeated" name:"UserId"`
}

func (req *RemoveUsersRequest) Invoke(client *sdk.Client) (resp *RemoveUsersResponse, err error) {
	req.InitWithApiInfo("CCC", "2017-07-05", "RemoveUsers", "ccc", "")
	resp = &RemoveUsersResponse{}
	err = client.DoAction(req, resp)
	return
}

type RemoveUsersResponse struct {
	responses.BaseResponse
	RequestId      string
	Success        bool
	Code           string
	Message        string
	HttpStatusCode int
}

type RemoveUsersUserIdList []string

func (list *RemoveUsersUserIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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
