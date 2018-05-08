package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeResourceDiagnosisRequest struct {
	requests.RpcRequest
	EndTime      string `position:"Query" name:"EndTime"`
	DBInstanceId string `position:"Query" name:"DBInstanceId"`
	StartTime    string `position:"Query" name:"StartTime"`
}

func (req *DescribeResourceDiagnosisRequest) Invoke(client *sdk.Client) (resp *DescribeResourceDiagnosisResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeResourceDiagnosis", "rds", "")
	resp = &DescribeResourceDiagnosisResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeResourceDiagnosisResponse struct {
	responses.BaseResponse
	RequestId  common.String
	StartTime  common.String
	EndTime    common.String
	CPU        DescribeResourceDiagnosisCPUList
	Memory     DescribeResourceDiagnosisMemoryList
	Storage    DescribeResourceDiagnosisStorageList
	IOPS       DescribeResourceDiagnosisIOPList
	Connection DescribeResourceDiagnosisConnectionList
}

type DescribeResourceDiagnosisCPUList []common.String

func (list *DescribeResourceDiagnosisCPUList) UnmarshalJSON(data []byte) error {
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

type DescribeResourceDiagnosisMemoryList []common.String

func (list *DescribeResourceDiagnosisMemoryList) UnmarshalJSON(data []byte) error {
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

type DescribeResourceDiagnosisStorageList []common.String

func (list *DescribeResourceDiagnosisStorageList) UnmarshalJSON(data []byte) error {
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

type DescribeResourceDiagnosisIOPList []common.String

func (list *DescribeResourceDiagnosisIOPList) UnmarshalJSON(data []byte) error {
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

type DescribeResourceDiagnosisConnectionList []common.String

func (list *DescribeResourceDiagnosisConnectionList) UnmarshalJSON(data []byte) error {
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
