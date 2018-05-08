package teslamaxcompute

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetQuotaHistoryInfoRequest struct {
	requests.RpcRequest
	Cluster   string `position:"Query" name:"Cluster"`
	EndTime   int    `position:"Query" name:"EndTime"`
	StartTime int    `position:"Query" name:"StartTime"`
	Region    string `position:"Query" name:"Region"`
	QuotaName string `position:"Query" name:"QuotaName"`
}

func (req *GetQuotaHistoryInfoRequest) Invoke(client *sdk.Client) (resp *GetQuotaHistoryInfoResponse, err error) {
	req.InitWithApiInfo("TeslaMaxCompute", "2018-01-04", "GetQuotaHistoryInfo", "", "")
	resp = &GetQuotaHistoryInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetQuotaHistoryInfoResponse struct {
	responses.BaseResponse
	Code      common.Integer
	Message   common.String
	RequestId common.String
	Data      GetQuotaHistoryInfoDataItemList
}

type GetQuotaHistoryInfoDataItem struct {
	Times common.Integer
	Point GetQuotaHistoryInfoPoint
}

type GetQuotaHistoryInfoPoint struct {
	CpuMaxQuota GetQuotaHistoryInfoCpuMaxQuota
	CpuMinQuota GetQuotaHistoryInfoCpuMinQuota
	MemUsed     GetQuotaHistoryInfoMemUsed
	CpuUsed     GetQuotaHistoryInfoCpuUsed
	MemMaxQuota GetQuotaHistoryInfoMemMaxQuota
	MemMinQuota GetQuotaHistoryInfoMemMinQuota
}

type GetQuotaHistoryInfoCpuMaxQuota struct {
	Min common.Float
	Max common.Float
	Avg common.Float
}

type GetQuotaHistoryInfoCpuMinQuota struct {
	Min common.Float
	Max common.Float
	Avg common.Float
}

type GetQuotaHistoryInfoMemUsed struct {
	Min common.Float
	Max common.Float
	Avg common.Float
}

type GetQuotaHistoryInfoCpuUsed struct {
	Min common.Float
	Max common.Float
	Avg common.Float
}

type GetQuotaHistoryInfoMemMaxQuota struct {
	Min common.Float
	Max common.Float
	Avg common.Float
}

type GetQuotaHistoryInfoMemMinQuota struct {
	Min common.Float
	Max common.Float
	Avg common.Float
}

type GetQuotaHistoryInfoDataItemList []GetQuotaHistoryInfoDataItem

func (list *GetQuotaHistoryInfoDataItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetQuotaHistoryInfoDataItem)
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
