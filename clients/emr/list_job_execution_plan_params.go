package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListJobExecutionPlanParamsRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	RelateId        string `position:"Query" name:"RelateId"`
	ParamBizType    string `position:"Query" name:"ParamBizType"`
}

func (req *ListJobExecutionPlanParamsRequest) Invoke(client *sdk.Client) (resp *ListJobExecutionPlanParamsResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ListJobExecutionPlanParams", "", "")
	resp = &ListJobExecutionPlanParamsResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListJobExecutionPlanParamsResponse struct {
	responses.BaseResponse
	RequestId     common.String
	Success       common.String
	ErrCode       common.String
	ErrMsg        common.String
	ParamInfoList ListJobExecutionPlanParamsParamInfoList
}

type ListJobExecutionPlanParamsParamInfo struct {
	ParamBizType         common.String
	RelateId             common.String
	ParamName            common.String
	ParamValue           common.String
	UtcCreateTimestamp   common.Long
	UtcModifiedTimestamp common.Long
}

type ListJobExecutionPlanParamsParamInfoList []ListJobExecutionPlanParamsParamInfo

func (list *ListJobExecutionPlanParamsParamInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListJobExecutionPlanParamsParamInfo)
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
