package yundun

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type VulScanLogRequest struct {
	JstOwnerId int64  `position:"Query" name:"JstOwnerId"`
	PageNumber int    `position:"Query" name:"PageNumber"`
	PageSize   int    `position:"Query" name:"PageSize"`
	InstanceId string `position:"Query" name:"InstanceId"`
	VulStatus  int    `position:"Query" name:"VulStatus"`
}

func (r VulScanLogRequest) Invoke(client *sdk.Client) (response *VulScanLogResponse, err error) {
	req := struct {
		*requests.RpcRequest
		VulScanLogRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Yundun", "2015-04-16", "VulScanLog", "", "")

	resp := struct {
		*responses.BaseResponse
		VulScanLogResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.VulScanLogResponse

	err = client.DoAction(&req, &resp)
	return
}

type VulScanLogResponse struct {
	RequestId  string
	StartTime  string
	EndTime    string
	PageNumber int
	PageSize   int
	TotalCount int
	LogList    VulScanLogVulScanLogList
}

type VulScanLogVulScanLog struct {
	Id           int
	Type         string
	Url          string
	HelpAddress  string
	VulParameter string
	Status       int
}

type VulScanLogVulScanLogList []VulScanLogVulScanLog

func (list *VulScanLogVulScanLogList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]VulScanLogVulScanLog)
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
