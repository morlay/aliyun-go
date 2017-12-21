package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeResourceDiagnosisRequest struct {
	EndTime      string `position:"Query" name:"EndTime"`
	DBInstanceId string `position:"Query" name:"DBInstanceId"`
	StartTime    string `position:"Query" name:"StartTime"`
}

func (r DescribeResourceDiagnosisRequest) Invoke(client *sdk.Client) (response *DescribeResourceDiagnosisResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeResourceDiagnosisRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeResourceDiagnosis", "rds", "")

	resp := struct {
		*responses.BaseResponse
		DescribeResourceDiagnosisResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeResourceDiagnosisResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeResourceDiagnosisResponse struct {
	RequestId  string
	StartTime  string
	EndTime    string
	CPU        DescribeResourceDiagnosisCPUList
	Memory     DescribeResourceDiagnosisMemoryList
	Storage    DescribeResourceDiagnosisStorageList
	IOPS       DescribeResourceDiagnosisIOPList
	Connection DescribeResourceDiagnosisConnectionList
}

type DescribeResourceDiagnosisCPUList []string

func (list *DescribeResourceDiagnosisCPUList) UnmarshalJSON(data []byte) error {
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

type DescribeResourceDiagnosisMemoryList []string

func (list *DescribeResourceDiagnosisMemoryList) UnmarshalJSON(data []byte) error {
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

type DescribeResourceDiagnosisStorageList []string

func (list *DescribeResourceDiagnosisStorageList) UnmarshalJSON(data []byte) error {
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

type DescribeResourceDiagnosisIOPList []string

func (list *DescribeResourceDiagnosisIOPList) UnmarshalJSON(data []byte) error {
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

type DescribeResourceDiagnosisConnectionList []string

func (list *DescribeResourceDiagnosisConnectionList) UnmarshalJSON(data []byte) error {
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
