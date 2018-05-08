package nas

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateFileSystemRequest struct {
	requests.RpcRequest
	Description  string `position:"Query" name:"Description"`
	ProtocolType string `position:"Query" name:"ProtocolType"`
	StorageType  string `position:"Query" name:"StorageType"`
}

func (req *CreateFileSystemRequest) Invoke(client *sdk.Client) (resp *CreateFileSystemResponse, err error) {
	req.InitWithApiInfo("NAS", "2017-06-26", "CreateFileSystem", "nas", "")
	resp = &CreateFileSystemResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateFileSystemResponse struct {
	responses.BaseResponse
	RequestId    common.String
	FileSystemId common.String
}
