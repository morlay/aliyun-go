package nas

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyFileSystemRequest struct {
	requests.RpcRequest
	Description  string `position:"Query" name:"Description"`
	FileSystemId string `position:"Query" name:"FileSystemId"`
}

func (req *ModifyFileSystemRequest) Invoke(client *sdk.Client) (resp *ModifyFileSystemResponse, err error) {
	req.InitWithApiInfo("NAS", "2017-06-26", "ModifyFileSystem", "nas", "")
	resp = &ModifyFileSystemResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyFileSystemResponse struct {
	responses.BaseResponse
	RequestId common.String
}
