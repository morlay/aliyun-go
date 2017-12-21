package rds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDBInstanceTDERequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeDBInstanceTDERequest) Invoke(client *sdk.Client) (response *DescribeDBInstanceTDEResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeDBInstanceTDERequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "DescribeDBInstanceTDE", "rds", "")

	resp := struct {
		*responses.BaseResponse
		DescribeDBInstanceTDEResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeDBInstanceTDEResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeDBInstanceTDEResponse struct {
	RequestId string
	TDEStatus string
	Databases DescribeDBInstanceTDEDatabaseList
}

type DescribeDBInstanceTDEDatabase struct {
	DBName    string
	TDEStatus string
}

type DescribeDBInstanceTDEDatabaseList []DescribeDBInstanceTDEDatabase

func (list *DescribeDBInstanceTDEDatabaseList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDBInstanceTDEDatabase)
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
