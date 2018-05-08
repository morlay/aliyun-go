package ccc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId      common.String
	Success        bool
	Code           common.String
	Message        common.String
	HttpStatusCode common.Integer
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
