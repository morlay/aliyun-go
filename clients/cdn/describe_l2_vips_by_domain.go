package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeL2VipsByDomainRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeL2VipsByDomainRequest) Invoke(client *sdk.Client) (resp *DescribeL2VipsByDomainResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeL2VipsByDomain", "", "")
	resp = &DescribeL2VipsByDomainResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeL2VipsByDomainResponse struct {
	responses.BaseResponse
	RequestId  common.String
	DomainName common.String
	Vips       DescribeL2VipsByDomainVipList
}

type DescribeL2VipsByDomainVipList []common.String

func (list *DescribeL2VipsByDomainVipList) UnmarshalJSON(data []byte) error {
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
