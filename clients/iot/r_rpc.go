package iot

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RRpcRequest struct {
	RequestBase64Byte string `position:"Query" name:"RequestBase.64.Byte"`
	DeviceName        string `position:"Query" name:"DeviceName"`
	ProductKey        string `position:"Query" name:"ProductKey"`
	Timeout           int    `position:"Query" name:"Timeout"`
}

func (r RRpcRequest) Invoke(client *sdk.Client) (response *RRpcResponse, err error) {
	req := struct {
		*requests.RpcRequest
		RRpcRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Iot", "2017-04-20", "RRpc", "", "")

	resp := struct {
		*responses.BaseResponse
		RRpcResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.RRpcResponse

	err = client.DoAction(&req, &resp)
	return
}

type RRpcResponse struct {
	RequestId         string
	Success           bool
	ErrorMessage      string
	RrpcCode          string
	PayloadBase64Byte string
	MessageId         int64
}
