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
}

type GetQuotaHistoryInfoCpuMaxQuota struct {
	Min int
	Max int
	Avg int
}

type GetQuotaHistoryInfoCpuMinQuota struct {
	Min int
	Max int
	Avg int
}

type GetQuotaHistoryInfoMemUsed struct {
	Min int
	Max int
	Avg int
}

type GetQuotaHistoryInfoCpuUsed struct {
	Min int
	Max int
	Avg int
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
