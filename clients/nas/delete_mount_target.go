package nas

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteMountTargetRequest struct {
	requests.RpcRequest
	MountTargetDomain string `position:"Query" name:"MountTargetDomain"`
	FileSystemId      string `position:"Query" name:"FileSystemId"`
}

func (req *DeleteMountTargetRequest) Invoke(client *sdk.Client) (resp *DeleteMountTargetResponse, err error) {
	req.InitWithApiInfo("NAS", "2017-06-26", "DeleteMountTarget", "nas", "")
	resp = &DeleteMountTargetResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteMountTargetResponse struct {
	responses.BaseResponse
	RequestId common.String
}
