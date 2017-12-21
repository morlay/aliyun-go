package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListExecutionPlansRequest struct {
	ResourceOwnerId int64                             `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string                            `position:"Query" name:"ClusterId"`
	JobId           string                            `position:"Query" name:"JobId"`
	Strategy        string                            `position:"Query" name:"Strategy"`
	IsDesc          string                            `position:"Query" name:"IsDesc"`
	PageNumber      int                               `position:"Query" name:"PageNumber"`
	PageSize        int                               `position:"Query" name:"PageSize"`
	StatusLists     *ListExecutionPlansStatusListList `position:"Query" type:"Repeated" name:"StatusList"`
}

func (r ListExecutionPlansRequest) Invoke(client *sdk.Client) (response *ListExecutionPlansResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ListExecutionPlansRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Emr", "2016-04-08", "ListExecutionPlans", "", "")

	resp := struct {
		*responses.BaseResponse
		ListExecutionPlansResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ListExecutionPlansResponse

	err = client.DoAction(&req, &resp)
	return
}

type ListExecutionPlansResponse struct {
	RequestId      string
	TotalCount     int
	PageNumber     int
	PageSize       int
	ExecutionPlans ListExecutionPlansExecutionPlanInfoList
}

type ListExecutionPlansExecutionPlanInfo struct {
	Id                    string
	Name                  string
	CreateClusterOnDemand bool
	Stragety              string
	Status                string
	TimeInterval          int
	StartTime             int64
	TimeUnit              string
}

type ListExecutionPlansStatusListList []string

func (list *ListExecutionPlansStatusListList) UnmarshalJSON(data []byte) error {
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

type ListExecutionPlansExecutionPlanInfoList []ListExecutionPlansExecutionPlanInfo

func (list *ListExecutionPlansExecutionPlanInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListExecutionPlansExecutionPlanInfo)
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
