package nas

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteMountTargetRequest struct {
	MountTargetDomain string `position:"Query" name:"MountTargetDomain"`
	FileSystemId      string `position:"Query" name:"FileSystemId"`
}

func (r DeleteMountTargetRequest) Invoke(client *sdk.Client) (response *DeleteMountTargetResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteMountTargetRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("NAS", "2017-06-26", "DeleteMountTarget", "nas", "")

	resp := struct {
		*responses.BaseResponse
		DeleteMountTargetResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteMountTargetResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteMountTargetResponse struct {
	RequestId string
}
