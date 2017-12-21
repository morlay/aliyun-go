package nas

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateMountTargetRequest struct {
	requests.RpcRequest
	VSwitchId       string `position:"Query" name:"VSwitchId"`
	VpcId           string `position:"Query" name:"VpcId"`
	NetworkType     string `position:"Query" name:"NetworkType"`
	AccessGroupName string `position:"Query" name:"AccessGroupName"`
	FileSystemId    string `position:"Query" name:"FileSystemId"`
}

func (req *CreateMountTargetRequest) Invoke(client *sdk.Client) (resp *CreateMountTargetResponse, err error) {
	req.InitWithApiInfo("NAS", "2017-06-26", "CreateMountTarget", "nas", "")
	resp = &CreateMountTargetResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateMountTargetResponse struct {
	responses.BaseResponse
	RequestId         string
	MountTargetDomain string
}
