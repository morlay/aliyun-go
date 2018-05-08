package alidns

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeDnsProductInstancesRequest struct {
	requests.RpcRequest
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	PageNumber   int64  `position:"Query" name:"PageNumber"`
	PageSize     int64  `position:"Query" name:"PageSize"`
	VersionCode  string `position:"Query" name:"VersionCode"`
}

func (req *DescribeDnsProductInstancesRequest) Invoke(client *sdk.Client) (resp *DescribeDnsProductInstancesResponse, err error) {
	req.InitWithApiInfo("Alidns", "2015-01-09", "DescribeDnsProductInstances", "", "")
	resp = &DescribeDnsProductInstancesResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDnsProductInstancesResponse struct {
	responses.BaseResponse
	RequestId   common.String
	TotalCount  common.Long
	PageNumber  common.Long
	PageSize    common.Long
	DnsProducts DescribeDnsProductInstancesDnsProductList
}

type DescribeDnsProductInstancesDnsProduct struct {
	InstanceId            common.String
	VersionCode           common.String
	VersionName           common.String
	StartTime             common.String
	EndTime               common.String
	StartTimestamp        common.Long
	EndTimestamp          common.Long
	Domain                common.String
	BindCount             common.Long
	BindUsedCount         common.Long
	TTLMinValue           common.Long
	SubDomainLevel        common.Long
	DnsSLBCount           common.Long
	URLForwardCount       common.Long
	DDosDefendFlow        common.Long
	DDosDefendQuery       common.Long
	OverseaDDosDefendFlow common.Long
	SearchEngineLines     common.String
	ISPLines              common.String
	ISPRegionLines        common.String
	OverseaLine           common.String
	MonitorNodeCount      common.Long
	MonitorFrequency      common.Long
	MonitorTaskCount      common.Long
}

type DescribeDnsProductInstancesDnsProductList []DescribeDnsProductInstancesDnsProduct

func (list *DescribeDnsProductInstancesDnsProductList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDnsProductInstancesDnsProduct)
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
