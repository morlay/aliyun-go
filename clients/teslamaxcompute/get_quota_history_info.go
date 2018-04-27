package teslamaxcompute

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	Code      int
	Message   string
	RequestId string
	Data      GetQuotaHistoryInfoDataItemList
}

type GetQuotaHistoryInfoDataItem struct {
	Times int
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
	Min float32
	Max float32
	Avg float32
}

type GetQuotaHistoryInfoCpuMinQuota struct {
	Min float32
	Max float32
	Avg float32
}

type GetQuotaHistoryInfoMemUsed struct {
	Min float32
	Max float32
	Avg float32
}

type GetQuotaHistoryInfoCpuUsed struct {
	Min float32
	Max float32
	Avg float32
}

type GetQuotaHistoryInfoMemMaxQuota struct {
	Min float32
	Max float32
	Avg float32
}

type GetQuotaHistoryInfoMemMinQuota struct {
	Min float32
	Max float32
	Avg float32
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
