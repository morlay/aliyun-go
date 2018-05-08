package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateApiStageVariableRequest struct {
	requests.RpcRequest
	GroupId       string `position:"Query" name:"GroupId"`
	StageId       string `position:"Query" name:"StageId"`
	VariableName  string `position:"Query" name:"VariableName"`
	VariableValue string `position:"Query" name:"VariableValue"`
}

func (req *CreateApiStageVariableRequest) Invoke(client *sdk.Client) (resp *CreateApiStageVariableResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "CreateApiStageVariable", "apigateway", "")
	resp = &CreateApiStageVariableResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateApiStageVariableResponse struct {
	responses.BaseResponse
	RequestId common.String
}
