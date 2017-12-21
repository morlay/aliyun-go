package iot

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateProductRequest struct {
	requests.RpcRequest
	CatId          int64  `position:"Query" name:"CatId"`
	Name           string `position:"Query" name:"Name"`
	ExtProps       string `position:"Query" name:"ExtProps"`
	SecurityPolicy string `position:"Query" name:"SecurityPolicy"`
	Desc           string `position:"Query" name:"Desc"`
}

func (req *CreateProductRequest) Invoke(client *sdk.Client) (resp *CreateProductResponse, err error) {
	req.InitWithApiInfo("Iot", "2017-04-20", "CreateProduct", "", "")
	resp = &CreateProductResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateProductResponse struct {
	responses.BaseResponse
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
