package drds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyDrdsInstanceDescriptionRequest struct {
	Description    string `position:"Query" name:"Description"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
}

func (r ModifyDrdsInstanceDescriptionRequest) Invoke(client *sdk.Client) (response *ModifyDrdsInstanceDescriptionResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyDrdsInstanceDescriptionRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Drds", "2017-10-16", "ModifyDrdsInstanceDescription", "", "")

	resp := struct {
		*responses.BaseResponse
		ModifyDrdsInstanceDescriptionResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyDrdsInstanceDescriptionResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyDrdsInstanceDescriptionResponse struct {
	RequestId string
	Success   bool
}
