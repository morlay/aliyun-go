package cms

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type TaskConfigEnableRequest struct {
	requests.RpcRequest
	IdLists *TaskConfigEnableIdListList `position:"Query" type:"Repeated" name:"IdList"`
	Enabled string                      `position:"Query" name:"Enabled"`
}

func (req *TaskConfigEnableRequest) Invoke(client *sdk.Client) (resp *TaskConfigEnableResponse, err error) {
	req.InitWithApiInfo("Cms", "2018-03-08", "TaskConfigEnable", "cms", "")
	resp = &TaskConfigEnableResponse{}
	err = client.DoAction(req, resp)
	return
}

type TaskConfigEnableResponse struct {
	responses.BaseResponse
	ErrorCode    int
	ErrorMessage string
	Success      bool
	RequestId    string
}

type TaskConfigEnableIdListList []int64

func (list *TaskConfigEnableIdListList) UnmarshalJSON(data []byte) error {
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
