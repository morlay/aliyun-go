package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateApiStageVariableRequest struct {
	GroupId       string `position:"Query" name:"GroupId"`
	StageId       string `position:"Query" name:"StageId"`
	VariableName  string `position:"Query" name:"VariableName"`
	VariableValue string `position:"Query" name:"VariableValue"`
}

func (r CreateApiStageVariableRequest) Invoke(client *sdk.Client) (response *CreateApiStageVariableResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateApiStageVariableRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "CreateApiStageVariable", "apigateway", "")

	resp := struct {
		*responses.BaseResponse
		CreateApiStageVariableResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateApiStageVariableResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateApiStageVariableResponse struct {
	RequestId string
}
