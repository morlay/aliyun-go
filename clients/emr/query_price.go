package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type QueryPriceRequest struct {
	requests.RpcRequest
	ResourceOwnerId        int64                    `position:"Query" name:"ResourceOwnerId"`
	Period                 int                      `position:"Query" name:"Period"`
	TaskInstanceType       string                   `position:"Query" name:"TaskInstanceType"`
	TaskDiskType           string                   `position:"Query" name:"TaskDiskType"`
	IoOptimized            string                   `position:"Query" name:"IoOptimized"`
	MasterDiskType         string                   `position:"Query" name:"MasterDiskType"`
	TaskInstanceQuantity   int                      `position:"Query" name:"TaskInstanceQuantity"`
	MasterInstanceType     string                   `position:"Query" name:"MasterInstanceType"`
	CoreInstanceQuantity   int                      `position:"Query" name:"CoreInstanceQuantity"`
	Duration               int                      `position:"Query" name:"Duration"`
	MasterDiskQuantity     int                      `position:"Query" name:"MasterDiskQuantity"`
	CoreDiskQuantity       int                      `position:"Query" name:"CoreDiskQuantity"`
	CoreInstanceType       string                   `position:"Query" name:"CoreInstanceType"`
	NetType                string                   `position:"Query" name:"NetType"`
	HostGroups             *QueryPriceHostGroupList `position:"Query" type:"Repeated" name:"HostGroup"`
	ZoneId                 string                   `position:"Query" name:"ZoneId"`
	CoreDiskType           string                   `position:"Query" name:"CoreDiskType"`
	ChargeType             string                   `position:"Query" name:"ChargeType"`
	MasterInstanceQuantity int                      `position:"Query" name:"MasterInstanceQuantity"`
	TaskDiskQuantity       int                      `position:"Query" name:"TaskDiskQuantity"`
}

func (req *QueryPriceRequest) Invoke(client *sdk.Client) (resp *QueryPriceResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "QueryPrice", "", "")
	resp = &QueryPriceResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryPriceHostGroup struct {
	HostGroupId     string `name:"HostGroupId"`
	HostGroupName   string `name:"HostGroupName"`
	HostGroupType   string `name:"HostGroupType"`
	ClusterId       string `name:"ClusterId"`
	Comment         string `name:"Comment"`
	CreateType      string `name:"CreateType"`
	ChargeType      string `name:"ChargeType"`
	Period          int    `name:"Period"`
	NodeCount       int    `name:"NodeCount"`
	InstanceType    string `name:"InstanceType"`
	DiskType        string `name:"DiskType"`
	DiskCapacity    int    `name:"DiskCapacity"`
	DiskCount       int    `name:"DiskCount"`
	SysDiskType     string `name:"SysDiskType"`
	SysDiskCapacity int    `name:"SysDiskCapacity"`
	AutoRenew       string `name:"AutoRenew"`
	VSwitchId       string `name:"VSwitchId"`
}

type QueryPriceResponse struct {
	responses.BaseResponse
	RequestId  common.String
	EmrPrice   common.String
	EcsPrice   common.String
	EmrPriceDO QueryPriceEmrPriceDO
	EcsPriceDO QueryPriceEcsPriceDO
}

type QueryPriceEmrPriceDO struct {
	OriginalPrice common.String
	DiscountPrice common.String
	TradePrice    common.String
	TaxPrice      common.String
	Currency      common.String
}

type QueryPriceEcsPriceDO struct {
	OriginalPrice common.String
	DiscountPrice common.String
	TradePrice    common.String
	TaxPrice      common.String
	Currency      common.String
}

type QueryPriceHostGroupList []QueryPriceHostGroup

func (list *QueryPriceHostGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryPriceHostGroup)
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
