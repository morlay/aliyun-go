package cbn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId       common.String
	TotalCount      common.Integer
	PageNumber      common.Integer
	PageSize        common.Integer
	VbrHealthChecks DescribeCenVbrHealthCheckVbrHealthCheckList
}

type DescribeCenVbrHealthCheckVbrHealthCheck struct {
	CenId               common.String
	VbrInstanceId       common.String
	LinkStatus          common.String
	PacketLoss          common.Long
	HealthCheckSourceIp common.String
	HealthCheckTargetIp common.String
	Delay               common.Long
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
