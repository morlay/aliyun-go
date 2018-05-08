package dds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SampleRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *SampleRequest) Invoke(client *sdk.Client) (resp *SampleResponse, err error) {
	req.InitWithApiInfo("Dds", "2015-12-01", "Sample", "dds", "")
	resp = &SampleResponse{}
	err = client.DoAction(req, resp)
	return
}

type SampleResponse struct {
	responses.BaseResponse
	RequestId        common.String
	SecurityIps      common.String
	SecurityIpGroups SampleSecurityIpGroupList
}

type SampleSecurityIpGroup struct {
	SecurityIpGroupName      common.String
	SecurityIpGroupAttribute common.String
	SecurityIpList           common.String
}

type SampleSecurityIpGroupList []SampleSecurityIpGroup

func (list *SampleSecurityIpGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SampleSecurityIpGroup)
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
