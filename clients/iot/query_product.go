package iot

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId    common.String
	Success      bool
	ErrorMessage common.String
	Data         QueryProductData
}

type QueryProductData struct {
	GmtCreate   common.String
	DataFormat  common.Integer
	Description common.String
	DeviceCount common.Integer
	NodeType    common.Integer
	ProductKey  common.String
	ProductName common.String
}
