
package sas-api

import (
    "github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
)


type GetInstanceCountRequest struct {
OwnerId int64 `position:"Query" name:"OwnerId"`
}

func (r GetInstanceCountRequest) Invoke(client *sdk.Client) (response *GetInstanceCountResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetInstanceCountRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Sas-api", "2017-07-05", "GetInstanceCount", "", "")

	resp := struct {
		*responses.BaseResponse
		GetInstanceCountResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
    response = &resp.GetInstanceCountResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetInstanceCountResponse struct {
Code string
Message string
Success bool
RequestId string
Data int
}

