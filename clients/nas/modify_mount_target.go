package nas

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyMountTargetRequest struct {
	MountTargetDomain string `position:"Query" name:"MountTargetDomain"`
	AccessGroupName   string `position:"Query" name:"AccessGroupName"`
	FileSystemId      string `position:"Query" name:"FileSystemId"`
	Status            string `position:"Query" name:"Status"`
}

func (r ModifyMountTargetRequest) Invoke(client *sdk.Client) (response *ModifyMountTargetResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyMountTargetRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("NAS", "2017-06-26", "ModifyMountTarget", "nas", "")

	resp := struct {
		*responses.BaseResponse
		ModifyMountTargetResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyMountTargetResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyMountTargetResponse struct {
	RequestId string
}
