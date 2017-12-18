package oms

import "encoding/json"

func (c *OmsClient) GetProductDefine(req *GetProductDefineArgs) (resp *GetProductDefineResponse, err error) {
	resp = &GetProductDefineResponse{}
	err = c.Request("GetProductDefine", req, resp)
	return
}

type GetProductDefineProduct struct {
	ProductName string
	TypeList    GetProductDefineTypeList
}

type GetProductDefineType struct {
	DataType string
	Fields   GetProductDefineFieldList
}
type GetProductDefineArgs struct {
	ProductName string
	DataType    string
}
type GetProductDefineResponse struct {
	RequestId   string
	ProductName string
	DataType    string
	ProductList GetProductDefineProductList
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

func (c *OmsClient) GetUserData(req *GetUserDataArgs) (resp *GetUserDataResponse, err error) {
	resp = &GetUserDataResponse{}
	err = c.Request("GetUserData", req, resp)
	return
}

type GetUserDataData struct {
	DataItems GetUserDataDataItemList
}

type GetUserDataDataItem struct {
	Name  string
	Value string
}
type GetUserDataArgs struct {
	OwnerId      int64
	OwnerAccount string
	ProductName  string
	DataType     string
	StartTime    string
	EndTime      string
	NextToken    string
	MaxResult    int
}
type GetUserDataResponse struct {
	RequestId   string
	ProductName string
	DataType    string
	NextToken   string
	DataList    GetUserDataDataList
}

type GetUserDataDataItemList []GetUserDataDataItem

func (list *GetUserDataDataItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetUserDataDataItem)
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

type GetUserDataDataList []GetUserDataData

func (list *GetUserDataDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetUserDataData)
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
