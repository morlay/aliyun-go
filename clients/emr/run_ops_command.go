package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RunOpsCommandRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64                        `position:"Query" name:"ResourceOwnerId"`
	OpsCommandName  string                       `position:"Query" name:"OpsCommandName"`
	Comment         string                       `position:"Query" name:"Comment"`
	CustomParams    string                       `position:"Query" name:"CustomParams"`
	ClusterId       string                       `position:"Query" name:"ClusterId"`
	HostIdLists     *RunOpsCommandHostIdListList `position:"Query" type:"Repeated" name:"HostIdList"`
	Dimension       string                       `position:"Query" name:"Dimension"`
}

func (req *RunOpsCommandRequest) Invoke(client *sdk.Client) (resp *RunOpsCommandResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "RunOpsCommand", "", "")
	resp = &RunOpsCommandResponse{}
	err = client.DoAction(req, resp)
	return
}

type RunOpsCommandResponse struct {
	responses.BaseResponse
	RequestId   string
	OperationId int64
}

type RunOpsCommandHostIdListList []int64

func (list *RunOpsCommandHostIdListList) UnmarshalJSON(data []byte) error {
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
