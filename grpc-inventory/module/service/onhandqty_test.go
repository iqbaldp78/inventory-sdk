package service

// import (
// 	"context"
// 	"testing"
// 	"log"

// 	"github.com/technical-assessment/iqbal/salestock/grpc-inventory/module/pb"
// 	"github.com/technical-assessment/iqbal/salestock/grpc-inventory/module/shared/lib/postgres"
// )

// func TestSearchSample(t *testing.T) {
// 	input := []struct {
// 		ctx   context.Context
// 		input *pb.SearchSampleReq
// 	}{
// 		{ctx: context.Background(), input: &pb.SearchSampleReq{Key: "zzz", Limit: 10, Offset: 0}},
// 		{ctx: context.Background(), input: &pb.SearchSampleReq{Key: "abc", Limit: 10, Offset: 0}},
// 	}
// 	output := []interface{}{"Data not found", nil}

// 	for index := range input {
// 		_, err := svc.SearchSample(input[index].ctx, input[index].input)
// 		log.Println(err)
// 		if err != output[index] {
// 			if err.Error() != output[index].(string) {
// 				t.Errorf("Expected result to be `%v`. Got `%v`", output[index], err.Error())
// 			}
// 		}
// 	}
// }

// func TestCleanUp(t *testing.T) {
// 	_, err := postgres.GenerateTestDB()
// 	if err != nil {
// 		t.Error(err.Error())
// 	}
// }
