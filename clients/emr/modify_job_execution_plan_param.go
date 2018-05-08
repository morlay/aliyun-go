package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyJobExecutionPlanParamRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ParamName       string `position:"Query" name:"ParamName"`
	ParamValue      string `position:"Query" name:"ParamValue"`
	Id              int64  `position:"Query" name:"Id"`
}

func (req *ModifyJobExecutionPlanParamRequest) Invoke(client *sdk.Client) (resp *ModifyJobExecutionPlanParamResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ModifyJobExecutionPlanParam", "", "")
	resp = &ModifyJobExecutionPlanParamResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyJobExecutionPlanParamResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   common.String
	ErrCode   common.String
	ErrMsg    common.String
}
