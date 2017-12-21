package nas

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateMountTargetRequest struct {
	VSwitchId       string `position:"Query" name:"VSwitchId"`
	VpcId           string `position:"Query" name:"VpcId"`
	NetworkType     string `position:"Query" name:"NetworkType"`
	AccessGroupName string `position:"Query" name:"AccessGroupName"`
	FileSystemId    string `position:"Query" name:"FileSystemId"`
}

func (r CreateMountTargetRequest) Invoke(client *sdk.Client) (response *CreateMountTargetResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateMountTargetRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("NAS", "2017-06-26", "CreateMountTarget", "nas", "")

	resp := struct {
		*responses.BaseResponse
		CreateMountTargetResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateMountTargetResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateMountTargetResponse struct {
	RequestId         string
	MountTargetDomain string
}
