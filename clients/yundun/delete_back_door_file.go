package yundun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteBackDoorFileRequest struct {
	JstOwnerId int64  `position:"Query" name:"JstOwnerId"`
	InstanceId string `position:"Query" name:"InstanceId"`
	Path       string `position:"Query" name:"Path"`
}

func (r DeleteBackDoorFileRequest) Invoke(client *sdk.Client) (response *DeleteBackDoorFileResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteBackDoorFileRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Yundun", "2015-04-16", "DeleteBackDoorFile", "", "")

	resp := struct {
		*responses.BaseResponse
		DeleteBackDoorFileResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteBackDoorFileResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteBackDoorFileResponse struct {
	RequestId string
}
