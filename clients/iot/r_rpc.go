package iot

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RRpcRequest struct {
	requests.RpcRequest
	RequestBase64Byte string `position:"Query" name:"RequestBase.64.Byte"`
	DeviceName        string `position:"Query" name:"DeviceName"`
	ProductKey        string `position:"Query" name:"ProductKey"`
	Timeout           int    `position:"Query" name:"Timeout"`
}

func (req *RRpcRequest) Invoke(client *sdk.Client) (resp *RRpcResponse, err error) {
	req.InitWithApiInfo("Iot", "2017-04-20", "RRpc", "", "")
	resp = &RRpcResponse{}
	err = client.DoAction(req, resp)
	return
}

type RRpcResponse struct {
	responses.BaseResponse
	RequestId         string
	Success           bool
	ErrorMessage      string
	RrpcCode          string
	PayloadBase64Byte string
	MessageId         int64
}
