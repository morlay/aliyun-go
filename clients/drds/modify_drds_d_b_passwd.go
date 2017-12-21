package drds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyDrdsDBPasswdRequest struct {
	NewPasswd      string `position:"Query" name:"NewPasswd"`
	DbName         string `position:"Query" name:"DbName"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
}

func (r ModifyDrdsDBPasswdRequest) Invoke(client *sdk.Client) (response *ModifyDrdsDBPasswdResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyDrdsDBPasswdRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Drds", "2017-10-16", "ModifyDrdsDBPasswd", "", "")

	resp := struct {
		*responses.BaseResponse
		ModifyDrdsDBPasswdResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyDrdsDBPasswdResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyDrdsDBPasswdResponse struct {
	RequestId string
	Success   bool
}
