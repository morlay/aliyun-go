package nas

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyMountTargetRequest struct {
	requests.RpcRequest
	MountTargetDomain string `position:"Query" name:"MountTargetDomain"`
	AccessGroupName   string `position:"Query" name:"AccessGroupName"`
	FileSystemId      string `position:"Query" name:"FileSystemId"`
	Status            string `position:"Query" name:"Status"`
}

func (req *ModifyMountTargetRequest) Invoke(client *sdk.Client) (resp *ModifyMountTargetResponse, err error) {
	req.InitWithApiInfo("NAS", "2017-06-26", "ModifyMountTarget", "nas", "")
	resp = &ModifyMountTargetResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyMountTargetResponse struct {
	responses.BaseResponse
	RequestId string
}
