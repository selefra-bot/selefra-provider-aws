

package ec2







import (




	"testing"





	"github.com/aws/aws-sdk-go-v2/service/ec2"




	"github.com/aws/aws-sdk-go-v2/service/ec2/types"




	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"




	"github.com/selefra/selefra-provider-aws/table_schema_generator"


)





func buildEc2VpnGateways(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {




	m := mocks.NewMockEc2Client(ctrl)


	l := types.VpnGateway{}


	err := faker.FakeObject(&l)
	if err != nil {
		t.Fatal(err)
	}




	m.EXPECT().DescribeVpnGateways(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(




		&ec2.DescribeVpnGatewaysOutput{


			VpnGateways: []types.VpnGateway{l},




		}, nil)
	return aws_client.AwsServices{




		Ec2: m,




	}
}









func TestEc2VpnGateways(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsEc2VpnGatewaysGenerator{}), buildEc2VpnGateways, aws_client.TestOptions{})


}


