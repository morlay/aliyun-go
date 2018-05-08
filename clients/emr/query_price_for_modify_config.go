package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type QueryPriceForModifyConfigRequest struct {
	requests.RpcRequest
	ResourceOwnerId   int64                                          `position:"Query" name:"ResourceOwnerId"`
	ModifyConfigSpecs *QueryPriceForModifyConfigModifyConfigSpecList `position:"Query" type:"Repeated" name:"ModifyConfigSpec"`
	ClusterId         string                                         `position:"Query" name:"ClusterId"`
}

func (req *QueryPriceForModifyConfigRequest) Invoke(client *sdk.Client) (resp *QueryPriceForModifyConfigResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "QueryPriceForModifyConfig", "", "")
	resp = &QueryPriceForModifyConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryPriceForModifyConfigModifyConfigSpec struct {
	HostGroupId     string `name:"HostGroupId"`
	NewInstanceType string `name:"NewInstanceType"`
	NewDiskSize     int    `name:"NewDiskSize"`
}

type QueryPriceForModifyConfigResponse struct {
	responses.BaseResponse
	RequestId  common.String
	EcsId      common.String
	EmrPriceDO QueryPriceForModifyConfigEmrPriceDO
	EcsPriceDO QueryPriceForModifyConfigEcsPriceDO
}

type QueryPriceForModifyConfigEmrPriceDO struct {
	OriginalPrice common.String
	DiscountPrice common.String
	TradePrice    common.String
	TaxPrice      common.String
	Currency      common.String
}

type QueryPriceForModifyConfigEcsPriceDO struct {
	OriginalPrice common.String
	DiscountPrice common.String
	TradePrice    common.String
	TaxPrice      common.String
	Currency      common.String
}

type QueryPriceForModifyConfigModifyConfigSpecList []QueryPriceForModifyConfigModifyConfigSpec

func (list *QueryPriceForModifyConfigModifyConfigSpecList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryPriceForModifyConfigModifyConfigSpec)
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
