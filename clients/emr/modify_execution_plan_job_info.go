package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyExecutionPlanJobInfoRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64                                    `position:"Query" name:"ResourceOwnerId"`
	Id              string                                   `position:"Query" name:"Id"`
	JobIdLists      *ModifyExecutionPlanJobInfoJobIdListList `position:"Query" type:"Repeated" name:"JobIdList"`
}

func (req *ModifyExecutionPlanJobInfoRequest) Invoke(client *sdk.Client) (resp *ModifyExecutionPlanJobInfoResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ModifyExecutionPlanJobInfo", "", "")
	resp = &ModifyExecutionPlanJobInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyExecutionPlanJobInfoResponse struct {
	responses.BaseResponse
	RequestId common.String
}

type ModifyExecutionPlanJobInfoJobIdListList []string

func (list *ModifyExecutionPlanJobInfoJobIdListList) UnmarshalJSON(data []byte) error {
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
