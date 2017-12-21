package nas

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateFileSystemRequest struct {
	Description  string `position:"Query" name:"Description"`
	ProtocolType string `position:"Query" name:"ProtocolType"`
	StorageType  string `position:"Query" name:"StorageType"`
}

func (r CreateFileSystemRequest) Invoke(client *sdk.Client) (response *CreateFileSystemResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateFileSystemRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("NAS", "2017-06-26", "CreateFileSystem", "nas", "")

	resp := struct {
		*responses.BaseResponse
		CreateFileSystemResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateFileSystemResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateFileSystemResponse struct {
	RequestId    string
	FileSystemId string
}
