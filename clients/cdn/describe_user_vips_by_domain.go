package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeUserVipsByDomainRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	Available     string `position:"Query" name:"Available"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeUserVipsByDomainRequest) Invoke(client *sdk.Client) (response *DescribeUserVipsByDomainResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeUserVipsByDomainRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeUserVipsByDomain", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeUserVipsByDomainResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeUserVipsByDomainResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeUserVipsByDomainResponse struct {
	RequestId  string
	DomainName int64
	Vips       DescribeUserVipsByDomainVipList
}

type DescribeUserVipsByDomainVipList []string

func (list *DescribeUserVipsByDomainVipList) UnmarshalJSON(data []byte) error {
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
