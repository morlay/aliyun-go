package alidns

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDnsProductInstancesRequest struct {
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	PageNumber   int64  `position:"Query" name:"PageNumber"`
	PageSize     int64  `position:"Query" name:"PageSize"`
	VersionCode  string `position:"Query" name:"VersionCode"`
}

func (r DescribeDnsProductInstancesRequest) Invoke(client *sdk.Client) (response *DescribeDnsProductInstancesResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeDnsProductInstancesRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Alidns", "2015-01-09", "DescribeDnsProductInstances", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeDnsProductInstancesResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeDnsProductInstancesResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeDnsProductInstancesResponse struct {
	RequestId   string
	TotalCount  int64
	PageNumber  int64
	PageSize    int64
	DnsProducts DescribeDnsProductInstancesDnsProductList
}

type DescribeDnsProductInstancesDnsProduct struct {
	InstanceId            string
	VersionCode           string
	VersionName           string
	StartTime             string
	EndTime               string
	StartTimestamp        int64
	EndTimestamp          int64
	Domain                string
	BindCount             int64
	BindUsedCount         int64
	TTLMinValue           int64
	SubDomainLevel        int64
	DnsSLBCount           int64
	URLForwardCount       int64
	DDosDefendFlow        int64
	DDosDefendQuery       int64
	OverseaDDosDefendFlow int64
	SearchEngineLines     string
	ISPLines              string
	ISPRegionLines        string
	OverseaLine           string
	MonitorNodeCount      int64
	MonitorFrequency      int64
	MonitorTaskCount      int64
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
