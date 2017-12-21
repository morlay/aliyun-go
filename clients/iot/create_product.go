package iot

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateProductRequest struct {
	CatId          int64  `position:"Query" name:"CatId"`
	Name           string `position:"Query" name:"Name"`
	ExtProps       string `position:"Query" name:"ExtProps"`
	SecurityPolicy string `position:"Query" name:"SecurityPolicy"`
	Desc           string `position:"Query" name:"Desc"`
}

func (r CreateProductRequest) Invoke(client *sdk.Client) (response *CreateProductResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateProductRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Iot", "2017-04-20", "CreateProduct", "", "")

	resp := struct {
		*responses.BaseResponse
		CreateProductResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateProductResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateProductResponse struct {
	RequestId    string
	Success      bool
	ErrorMessage string
	ProductInfo  CreateProductProductInfo
}

type CreateProductProductInfo struct {
	ProductKey    string
	ProductName   string
	CatId         int64
	CreateUserId  int64
	ProductDesc   string
	FromSource    string
	ExtProps      string
	GmtCreate     string
	GmtModified   string
	ProductSecret string
}
