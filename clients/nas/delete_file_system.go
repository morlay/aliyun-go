package nas

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteFileSystemRequest struct {
	requests.RpcRequest
	FileSystemId string `position:"Query" name:"FileSystemId"`
}

func (req *DeleteFileSystemRequest) Invoke(client *sdk.Client) (resp *DeleteFileSystemResponse, err error) {
	req.InitWithApiInfo("NAS", "2017-06-26", "DeleteFileSystem", "nas", "")
	resp = &DeleteFileSystemResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteFileSystemResponse struct {
	responses.BaseResponse
	RequestId common.String
}
