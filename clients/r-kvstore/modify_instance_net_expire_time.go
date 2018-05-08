package r_kvstore

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyInstanceNetExpireTimeRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ConnectionString     string `position:"Query" name:"ConnectionString"`
	ClassicExpiredDays   int    `position:"Query" name:"ClassicExpiredDays"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *ModifyInstanceNetExpireTimeRequest) Invoke(client *sdk.Client) (resp *ModifyInstanceNetExpireTimeResponse, err error) {
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "ModifyInstanceNetExpireTime", "redisa", "")
	resp = &ModifyInstanceNetExpireTimeResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyInstanceNetExpireTimeResponse struct {
	responses.BaseResponse
	RequestId    common.String
	InstanceId   common.String
	NetInfoItems ModifyInstanceNetExpireTimeNetInfoItemList
}

type ModifyInstanceNetExpireTimeNetInfoItem struct {
	DBInstanceNetType common.String
	Port              common.String
	ExpiredTime       common.String
	ConnectionString  common.String
	IPAddress         common.String
}

type ModifyInstanceNetExpireTimeNetInfoItemList []ModifyInstanceNetExpireTimeNetInfoItem

func (list *ModifyInstanceNetExpireTimeNetInfoItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ModifyInstanceNetExpireTimeNetInfoItem)
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
