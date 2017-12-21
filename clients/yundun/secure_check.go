package yundun

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SecureCheckRequest struct {
	JstOwnerId  int64  `position:"Query" name:"JstOwnerId"`
	InstanceIds string `position:"Query" name:"InstanceIds"`
}

func (r SecureCheckRequest) Invoke(client *sdk.Client) (response *SecureCheckResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SecureCheckRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Yundun", "2015-04-16", "SecureCheck", "", "")

	resp := struct {
		*responses.BaseResponse
		SecureCheckResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SecureCheckResponse

	err = client.DoAction(&req, &resp)
	return
}

type SecureCheckResponse struct {
	RequestId        string
	RecentInstanceId string
	ProblemList      SecureCheckInfoList
	NoProblemList    SecureCheckInfoList
	NoScanList       SecureCheckInfoList
	ScanningList     SecureCheckInfoList
	InnerIpList      SecureCheckInfoList
}

type SecureCheckInfo struct {
	Ip         string
	Status     string
	VulNum     string
	InstanceId string
}

type SecureCheckInfoList []SecureCheckInfo

func (list *SecureCheckInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SecureCheckInfo)
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
