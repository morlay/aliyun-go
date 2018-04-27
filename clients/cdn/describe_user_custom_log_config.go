package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeUserCustomLogConfigRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	Version       string `position:"Query" name:"Version"`
}

func (req *DescribeUserCustomLogConfigRequest) Invoke(client *sdk.Client) (resp *DescribeUserCustomLogConfigResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeUserCustomLogConfig", "", "")
	resp = &DescribeUserCustomLogConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeUserCustomLogConfigResponse struct {
	responses.BaseResponse
	RequestId string
	ConfigIds DescribeUserCustomLogConfigConfigIdList
}

type DescribeUserCustomLogConfigConfigIdList []string

func (list *DescribeUserCustomLogConfigConfigIdList) UnmarshalJSON(data []byte) error {
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
