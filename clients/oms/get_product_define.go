package oms

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetProductDefineRequest struct {
	requests.RpcRequest
	ProductName string `position:"Query" name:"ProductName"`
	DataType    string `position:"Query" name:"DataType"`
}

func (req *GetProductDefineRequest) Invoke(client *sdk.Client) (resp *GetProductDefineResponse, err error) {
	req.InitWithApiInfo("Oms", "2015-02-12", "GetProductDefine", "", "")
	resp = &GetProductDefineResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetProductDefineResponse struct {
	responses.BaseResponse
	RequestId   common.String
	ProductName common.String
	DataType    common.String
	ProductList GetProductDefineProductList
}

type GetProductDefineProduct struct {
	ProductName common.String
	TypeList    GetProductDefineTypeList
}

type GetProductDefineType struct {
	DataType common.String
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

type GetProductDefineFieldList []common.String

func (list *GetProductDefineFieldList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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
