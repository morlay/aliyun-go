package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteApiStageVariableRequest struct {
	GroupId      string `position:"Query" name:"GroupId"`
	StageId      string `position:"Query" name:"StageId"`
	VariableName string `position:"Query" name:"VariableName"`
}

func (r DeleteApiStageVariableRequest) Invoke(client *sdk.Client) (response *DeleteApiStageVariableResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteApiStageVariableRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DeleteApiStageVariable", "apigateway", "")

	resp := struct {
		*responses.BaseResponse
		DeleteApiStageVariableResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteApiStageVariableResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteApiStageVariableResponse struct {
	RequestId string
}
