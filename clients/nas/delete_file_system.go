package nas

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteFileSystemRequest struct {
	FileSystemId string `position:"Query" name:"FileSystemId"`
}

func (r DeleteFileSystemRequest) Invoke(client *sdk.Client) (response *DeleteFileSystemResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteFileSystemRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("NAS", "2017-06-26", "DeleteFileSystem", "nas", "")

	resp := struct {
		*responses.BaseResponse
		DeleteFileSystemResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteFileSystemResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteFileSystemResponse struct {
	RequestId string
}
