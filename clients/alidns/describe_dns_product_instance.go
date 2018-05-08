package alidns

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId             common.String
	InstanceId            common.String
	VersionCode           common.String
	VersionName           common.String
	StartTime             common.String
	StartTimestamp        common.Long
	EndTime               common.String
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
	DnsServers            DescribeDnsProductInstanceDnsServerList
}

type DescribeDnsProductInstanceDnsServerList []common.String

func (list *DescribeDnsProductInstanceDnsServerList) UnmarshalJSON(data []byte) error {
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
