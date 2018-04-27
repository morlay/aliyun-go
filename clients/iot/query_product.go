package iot

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryProductRequest struct {
	requests.RpcRequest
	ProductKey string `position:"Query" name:"ProductKey"`
}

func (req *QueryProductRequest) Invoke(client *sdk.Client) (resp *QueryProductResponse, err error) {
	req.InitWithApiInfo("Iot", "2018-01-20", "QueryProduct", "", "")
	resp = &QueryProductResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryProductResponse struct {
	responses.BaseResponse
	RequestId    string
	Success      bool
	ErrorMessage string
	Data         QueryProductData
}

type QueryProductData struct {
	GmtCreate   string
	DataFormat  int
	Description string
	DeviceCount int
	NodeType    int
	ProductKey  string
	ProductName string
}
