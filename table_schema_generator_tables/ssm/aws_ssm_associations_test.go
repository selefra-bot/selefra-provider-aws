

package ssm







import (




	"testing"









	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"


	"github.com/selefra/selefra-provider-aws/aws_client/mocks"




	"github.com/selefra/selefra-provider-aws/faker"


	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)



func buildAssociations(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {


	mock := mocks.NewMockSsmClient(ctrl)







	var i types.Association




	if err := faker.FakeObject(&i); err != nil {
		t.Fatal(err)
	}




	mock.EXPECT().ListAssociations(
		gomock.Any(),




		gomock.Any(),
		gomock.Any(),


	).AnyTimes().Return(


		&ssm.ListAssociationsOutput{Associations: []types.Association{i}},




		nil,


	)





	return aws_client.AwsServices{Ssm: mock}
}







func TestAssociations(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsSsmAssociationsGenerator{}), buildAssociations, aws_client.TestOptions{})


}


