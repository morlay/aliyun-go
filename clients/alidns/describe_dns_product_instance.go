package alidns

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDnsProductInstanceRequest struct {
	requests.RpcRequest
	Lang         string `position:"Query" name:"Lang"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	InstanceId   string `position:"Query" name:"InstanceId"`
}

func (req *DescribeDnsProductInstanceRequest) Invoke(client *sdk.Client) (resp *DescribeDnsProductInstanceResponse, err error) {
	req.InitWithApiInfo("Alidns", "2015-01-09", "DescribeDnsProductInstance", "", "")
	resp = &DescribeDnsProductInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDnsProductInstanceResponse struct {
	responses.BaseResponse
	RequestId             string
	InstanceId            string
	VersionCode           string
	VersionName           string
	StartTime             string
	StartTimestamp        int64
	EndTime               string
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
	DnsServers            DescribeDnsProductInstanceDnsServerList
}

type DescribeDnsProductInstanceDnsServerList []string

func (list *DescribeDnsProductInstanceDnsServerList) UnmarshalJSON(data []byte) error {
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
