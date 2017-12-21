package iot

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpdateProductRequest struct {
	requests.RpcRequest
	CatId       int64  `position:"Query" name:"CatId"`
	ProductName string `position:"Query" name:"ProductName"`
	ExtProps    string `position:"Query" name:"ExtProps"`
	ProductKey  string `position:"Query" name:"ProductKey"`
	ProductDesc string `position:"Query" name:"ProductDesc"`
}

func (req *UpdateProductRequest) Invoke(client *sdk.Client) (resp *UpdateProductResponse, err error) {
	req.InitWithApiInfo("Iot", "2017-04-20", "UpdateProduct", "", "")
	resp = &UpdateProductResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpdateProductResponse struct {
	responses.BaseResponse
	RequestId    string
	Success      bool
	ErrorMessage string
}
