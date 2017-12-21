package oms

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetProductDefineRequest struct {
	ProductName string `position:"Query" name:"ProductName"`
	DataType    string `position:"Query" name:"DataType"`
}

func (r GetProductDefineRequest) Invoke(client *sdk.Client) (response *GetProductDefineResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetProductDefineRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Oms", "2015-02-12", "GetProductDefine", "", "")

	resp := struct {
		*responses.BaseResponse
		GetProductDefineResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetProductDefineResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetProductDefineResponse struct {
	RequestId   string
	ProductName string
	DataType    string
	ProductList GetProductDefineProductList
}

type GetProductDefineProduct struct {
	ProductName string
	TypeList    GetProductDefineTypeList
}

type GetProductDefineType struct {
	DataType string
	Fields   GetProductDefineFieldList
}

type GetProductDefineProductList []GetProductDefineProduct

func (list *GetProductDefineProductList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetProductDefineProduct)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}

type GetProductDefineTypeList []GetProductDefineType

func (list *GetProductDefineTypeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetProductDefineType)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}

type GetProductDefineFieldList []string

func (list *GetProductDefineFieldList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
