package market_inner

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type InnerDescribeImageInstancesRequest struct {
	requests.RpcRequest
	EcsInstanceId        string                                              `position:"Query" name:"EcsInstanceId"`
	PageSize             int                                                 `position:"Query" name:"PageSize"`
	ImageNo              string                                              `position:"Query" name:"ImageNo"`
	ProductName          string                                              `position:"Query" name:"ProductName"`
	PageNumber           int                                                 `position:"Query" name:"PageNumber"`
	ImageInstanceIdLists *InnerDescribeImageInstancesImageInstanceIdListList `position:"Query" type:"Repeated" name:"ImageInstanceIdList"`
	RegionNo             string                                              `position:"Query" name:"RegionNo"`
}

func (req *InnerDescribeImageInstancesRequest) Invoke(client *sdk.Client) (resp *InnerDescribeImageInstancesResponse, err error) {
	req.InitWithApiInfo("Market-Inner", "2016-08-01", "InnerDescribeImageInstances", "yunmarket", "")
	resp = &InnerDescribeImageInstancesResponse{}
	err = client.DoAction(req, resp)
	return
}

type InnerDescribeImageInstancesResponse struct {
	responses.BaseResponse
	PageNumber common.Integer
	PageSize   common.Integer
	TotalCount common.Integer
	RequestId  common.String
	ImageList  InnerDescribeImageInstancesImageList
}

type InnerDescribeImageInstancesImage struct {
	InstanceId     common.Long
	OrderId        common.Long
	SupplierId     common.Long
	SupplierName   common.String
	ProductCode    common.String
	ProductSkuCode common.String
	ProductName    common.String
	Status         common.String
	ChargeType     common.String
	BeganOn        common.Long
	EndOn          common.Long
	CreatedOn      common.Long
	RemainTime     common.Long
	RegionNo       common.String
	ImageNo        common.String
	ImageVersion   common.String
	EcsInstanceId  common.String
}

type InnerDescribeImageInstancesImageInstanceIdListList []int64

func (list *InnerDescribeImageInstancesImageInstanceIdListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]int64)
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

type InnerDescribeImageInstancesImageList []InnerDescribeImageInstancesImage

func (list *InnerDescribeImageInstancesImageList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]InnerDescribeImageInstancesImage)
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
