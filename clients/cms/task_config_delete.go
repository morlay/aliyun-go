package cms

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type TaskConfigDeleteRequest struct {
	requests.RpcRequest
	IdLists *TaskConfigDeleteIdListList `position:"Query" type:"Repeated" name:"IdList"`
}

func (req *TaskConfigDeleteRequest) Invoke(client *sdk.Client) (resp *TaskConfigDeleteResponse, err error) {
	req.InitWithApiInfo("Cms", "2018-03-08", "TaskConfigDelete", "cms", "")
	resp = &TaskConfigDeleteResponse{}
	err = client.DoAction(req, resp)
	return
}

type TaskConfigDeleteResponse struct {
	responses.BaseResponse
	ErrorCode    common.Integer
	ErrorMessage common.String
	Success      bool
	RequestId    common.String
}

type TaskConfigDeleteIdListList []int64

func (list *TaskConfigDeleteIdListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]int64)
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
