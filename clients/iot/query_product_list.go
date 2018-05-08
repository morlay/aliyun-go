package iot

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type QueryProductListRequest struct {
	requests.RpcRequest
	PageSize            int    `position:"Query" name:"PageSize"`
	CurrentPage         int    `position:"Query" name:"CurrentPage"`
	AliyunCommodityCode string `position:"Query" name:"AliyunCommodityCode"`
}

func (req *QueryProductListRequest) Invoke(client *sdk.Client) (resp *QueryProductListResponse, err error) {
	req.InitWithApiInfo("Iot", "2018-01-20", "QueryProductList", "", "")
	resp = &QueryProductListResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryProductListResponse struct {
	responses.BaseResponse
	RequestId    common.String
	Success      bool
	ErrorMessage common.String
	Data         QueryProductListData
}

type QueryProductListData struct {
	CurrentPage common.Integer
	PageCount   common.Integer
	PageSize    common.Integer
	Total       common.Integer
	List        QueryProductListProductInfoList
}

type QueryProductListProductInfo struct {
	GmtCreate   common.String
	DataFormat  common.Integer
	Description common.String
	DeviceCount common.Integer
	NodeType    common.Integer
	ProductKey  common.String
	ProductName common.String
}

type QueryProductListProductInfoList []QueryProductListProductInfo

func (list *QueryProductListProductInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryProductListProductInfo)
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
