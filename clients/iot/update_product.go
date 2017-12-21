package iot

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpdateProductRequest struct {
	CatId       int64  `position:"Query" name:"CatId"`
	ProductName string `position:"Query" name:"ProductName"`
	ExtProps    string `position:"Query" name:"ExtProps"`
	ProductKey  string `position:"Query" name:"ProductKey"`
	ProductDesc string `position:"Query" name:"ProductDesc"`
}

func (r UpdateProductRequest) Invoke(client *sdk.Client) (response *UpdateProductResponse, err error) {
	req := struct {
		*requests.RpcRequest
		UpdateProductRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Iot", "2017-04-20", "UpdateProduct", "", "")

	resp := struct {
		*responses.BaseResponse
		UpdateProductResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.UpdateProductResponse

	err = client.DoAction(&req, &resp)
	return
}

type UpdateProductResponse struct {
	RequestId    string
	Success      bool
	ErrorMessage string
}
