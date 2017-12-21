package ubsms

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeBusinessStatusRequest struct {
	requests.RpcRequest
	CallerBid string `position:"Query" name:"CallerBid"`
	Password  string `position:"Query" name:"Password"`
}

func (req *DescribeBusinessStatusRequest) Invoke(client *sdk.Client) (resp *DescribeBusinessStatusResponse, err error) {
	req.InitWithApiInfo("Ubsms", "2015-06-23", "DescribeBusinessStatus", "", "")
	resp = &DescribeBusinessStatusResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeBusinessStatusResponse struct {
	responses.BaseResponse
	RequestId              string
	Success                bool
	UserBusinessStatusList DescribeBusinessStatusUserBusinessStatusList
}

type DescribeBusinessStatusUserBusinessStatus struct {
	Uid         string
	ServiceCode string
	Statuses    DescribeBusinessStatusStatusList
}

type DescribeBusinessStatusStatus struct {
	StatusKey   string
	StatusValue string
}

type DescribeBusinessStatusUserBusinessStatusList []DescribeBusinessStatusUserBusinessStatus

func (list *DescribeBusinessStatusUserBusinessStatusList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeBusinessStatusUserBusinessStatus)
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

type DescribeBusinessStatusStatusList []DescribeBusinessStatusStatus

func (list *DescribeBusinessStatusStatusList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeBusinessStatusStatus)
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
