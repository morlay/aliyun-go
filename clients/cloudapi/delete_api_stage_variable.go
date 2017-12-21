package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteApiStageVariableRequest struct {
	requests.RpcRequest
	GroupId      string `position:"Query" name:"GroupId"`
	StageId      string `position:"Query" name:"StageId"`
	VariableName string `position:"Query" name:"VariableName"`
}

func (req *DeleteApiStageVariableRequest) Invoke(client *sdk.Client) (resp *DeleteApiStageVariableResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DeleteApiStageVariable", "apigateway", "")
	resp = &DeleteApiStageVariableResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteApiStageVariableResponse struct {
	responses.BaseResponse
	RequestId string
}
