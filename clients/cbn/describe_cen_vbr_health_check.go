package cbn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeCenVbrHealthCheckRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	CenId                string `position:"Query" name:"CenId"`
	VbrInstanceOwnerId   int64  `position:"Query" name:"VbrInstanceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	VbrInstanceId        string `position:"Query" name:"VbrInstanceId"`
	PageSize             int    `position:"Query" name:"PageSize"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	VbrInstanceRegionId  string `position:"Query" name:"VbrInstanceRegionId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
}

func (req *DescribeCenVbrHealthCheckRequest) Invoke(client *sdk.Client) (resp *DescribeCenVbrHealthCheckResponse, err error) {
	req.InitWithApiInfo("Cbn", "2017-09-12", "DescribeCenVbrHealthCheck", "cbn", "")
	resp = &DescribeCenVbrHealthCheckResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeCenVbrHealthCheckResponse struct {
	responses.BaseResponse
	RequestId       string
	TotalCount      int
	PageNumber      int
	PageSize        int
	VbrHealthChecks DescribeCenVbrHealthCheckVbrHealthCheckList
}

type DescribeCenVbrHealthCheckVbrHealthCheck struct {
	CenId               string
	VbrInstanceId       string
	LinkStatus          string
	PacketLoss          int64
	HealthCheckSourceIp string
	HealthCheckTargetIp string
	Delay               int64
}

type DescribeCenVbrHealthCheckVbrHealthCheckList []DescribeCenVbrHealthCheckVbrHealthCheck

func (list *DescribeCenVbrHealthCheckVbrHealthCheckList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCenVbrHealthCheckVbrHealthCheck)
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
