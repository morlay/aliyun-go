package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateJobExecutionPlanParamRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64                                         `position:"Query" name:"ResourceOwnerId"`
	RelateId        string                                        `position:"Query" name:"RelateId"`
	WorkParamPairs  *CreateJobExecutionPlanParamWorkParamPairList `position:"Query" type:"Repeated" name:"WorkParamPair"`
	ParamBizType    string                                        `position:"Query" name:"ParamBizType"`
}

func (req *CreateJobExecutionPlanParamRequest) Invoke(client *sdk.Client) (resp *CreateJobExecutionPlanParamResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "CreateJobExecutionPlanParam", "", "")
	resp = &CreateJobExecutionPlanParamResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateJobExecutionPlanParamWorkParamPair struct {
	Name  string `name:"Name"`
	Value string `name:"Value"`
}

type CreateJobExecutionPlanParamResponse struct {
	responses.BaseResponse
	RequestId string
	Success   string
	ErrCode   string
	ErrMsg    string
	Ids       CreateJobExecutionPlanParamIdList
}

type CreateJobExecutionPlanParamWorkParamPairList []CreateJobExecutionPlanParamWorkParamPair

func (list *CreateJobExecutionPlanParamWorkParamPairList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]CreateJobExecutionPlanParamWorkParamPair)
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

type CreateJobExecutionPlanParamIdList []string

func (list *CreateJobExecutionPlanParamIdList) UnmarshalJSON(data []byte) error {
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
