package vpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeHaVipsRequest struct {
	requests.RpcRequest
	Filters              *DescribeHaVipsFilterList `position:"Query" type:"Repeated" name:"Filter"`
	ResourceOwnerId      int64                     `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string                    `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string                    `position:"Query" name:"OwnerAccount"`
	PageSize             int                       `position:"Query" name:"PageSize"`
	OwnerId              int64                     `position:"Query" name:"OwnerId"`
	PageNumber           int                       `position:"Query" name:"PageNumber"`
}

func (req *DescribeHaVipsRequest) Invoke(client *sdk.Client) (resp *DescribeHaVipsResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "DescribeHaVips", "vpc", "")
	resp = &DescribeHaVipsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeHaVipsFilter struct {
	Key    string                   `name:"Key"`
	Values *DescribeHaVipsValueList `type:"Repeated" name:"Value"`
}

type DescribeHaVipsResponse struct {
	responses.BaseResponse
	RequestId  common.String
	TotalCount common.Integer
	PageNumber common.Integer
	PageSize   common.Integer
	HaVips     DescribeHaVipsHaVipList
}

type DescribeHaVipsHaVip struct {
	HaVipId                common.String
	RegionId               common.String
	VpcId                  common.String
	VSwitchId              common.String
	IpAddress              common.String
	Status                 common.String
	MasterInstanceId       common.String
	Description            common.String
	CreateTime             common.String
	AssociatedInstances    DescribeHaVipsAssociatedInstanceList
	AssociatedEipAddresses DescribeHaVipsAssociatedEipAddressList
}

type DescribeHaVipsFilterList []DescribeHaVipsFilter

func (list *DescribeHaVipsFilterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeHaVipsFilter)
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

type DescribeHaVipsValueList []string

func (list *DescribeHaVipsValueList) UnmarshalJSON(data []byte) error {
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

type DescribeHaVipsHaVipList []DescribeHaVipsHaVip

func (list *DescribeHaVipsHaVipList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeHaVipsHaVip)
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

type DescribeHaVipsAssociatedInstanceList []common.String

func (list *DescribeHaVipsAssociatedInstanceList) UnmarshalJSON(data []byte) error {
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

type DescribeHaVipsAssociatedEipAddressList []common.String

func (list *DescribeHaVipsAssociatedEipAddressList) UnmarshalJSON(data []byte) error {
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
