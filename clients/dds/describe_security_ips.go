package dds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeSecurityIpsRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeSecurityIpsRequest) Invoke(client *sdk.Client) (resp *DescribeSecurityIpsResponse, err error) {
	req.InitWithApiInfo("Dds", "2015-12-01", "DescribeSecurityIps", "dds", "")
	resp = &DescribeSecurityIpsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeSecurityIpsResponse struct {
	responses.BaseResponse
	RequestId        common.String
	SecurityIps      common.String
	SecurityIpGroups DescribeSecurityIpsSecurityIpGroupList
}

type DescribeSecurityIpsSecurityIpGroup struct {
	SecurityIpGroupName      common.String
	SecurityIpGroupAttribute common.String
	SecurityIpList           common.String
}

type DescribeSecurityIpsSecurityIpGroupList []DescribeSecurityIpsSecurityIpGroup

func (list *DescribeSecurityIpsSecurityIpGroupList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeSecurityIpsSecurityIpGroup)
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
