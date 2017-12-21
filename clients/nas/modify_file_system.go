package nas

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyFileSystemRequest struct {
	Description  string `position:"Query" name:"Description"`
	FileSystemId string `position:"Query" name:"FileSystemId"`
}

func (r ModifyFileSystemRequest) Invoke(client *sdk.Client) (response *ModifyFileSystemResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyFileSystemRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("NAS", "2017-06-26", "ModifyFileSystem", "nas", "")

	resp := struct {
		*responses.BaseResponse
		ModifyFileSystemResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyFileSystemResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyFileSystemResponse struct {
	RequestId string
}
