package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeUserVipsByDomainRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	Available     string `position:"Query" name:"Available"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeUserVipsByDomainRequest) Invoke(client *sdk.Client) (resp *DescribeUserVipsByDomainResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeUserVipsByDomain", "", "")
	resp = &DescribeUserVipsByDomainResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeUserVipsByDomainResponse struct {
	responses.BaseResponse
	RequestId  common.String
	DomainName common.Long
	Vips       DescribeUserVipsByDomainVipList
}

type DescribeUserVipsByDomainVipList []common.String

func (list *DescribeUserVipsByDomainVipList) UnmarshalJSON(data []byte) error {
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
