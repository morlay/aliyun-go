package drds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyReadOnlyAccountPasswordRequest struct {
	NewPasswd      string `position:"Query" name:"NewPasswd"`
	DbName         string `position:"Query" name:"DbName"`
	AccountName    string `position:"Query" name:"AccountName"`
	OriginPassword string `position:"Query" name:"OriginPassword"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
}

func (r ModifyReadOnlyAccountPasswordRequest) Invoke(client *sdk.Client) (response *ModifyReadOnlyAccountPasswordResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyReadOnlyAccountPasswordRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Drds", "2017-10-16", "ModifyReadOnlyAccountPassword", "", "")

	resp := struct {
		*responses.BaseResponse
		ModifyReadOnlyAccountPasswordResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyReadOnlyAccountPasswordResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyReadOnlyAccountPasswordResponse struct {
	RequestId string
	Success   bool
}
