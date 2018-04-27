package iot

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId    string
	Success      bool
	ErrorMessage string
	Data         QueryProductListData
}

type QueryProductListData struct {
	CurrentPage int
	PageCount   int
	PageSize    int
	Total       int
	List        QueryProductListProductInfoList
}

type QueryProductListProductInfo struct {
	GmtCreate   string
	DataFormat  int
	Description string
	DeviceCount int
	NodeType    int
	ProductKey  string
	ProductName string
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
