package main

import (
	"fmt"

	pb "github.com/coreos/etcd/etcdserver/etcdserverpb"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
)

func main() {
	txn := &pb.TxnRequest{
		Compare: []*pb.Compare{
			&pb.Compare{
				Key: []byte("foo"),
				Result: pb.Compare_EQUAL,
				Target: pb.Compare_CREATE,
				TargetUnion: &pb.Compare_CreateRevision{0},
			},
		},
		Success: []*pb.RequestOp{
			{
				Request: &pb.RequestOp_RequestPut{
					RequestPut: &pb.PutRequest{
						Key: []byte("foo"),
						Value: []byte("bar"),
					},
				},
			},
		},
	}
	jsonpb := &runtime.JSONPb{}
	jsonpbOrig := &runtime.JSONPb{OrigName: true}
	jsonb := &runtime.JSONBuiltin{}

	bDat,  _ := jsonb.Marshal(txn)
	bTxn := &pb.TxnRequest{}
	if err := jsonb.Unmarshal(bDat, bTxn); err != nil {
		fmt.Println(err)
	}

	pbDat, _ := jsonpb.Marshal(txn)
	pbTxn := &pb.TxnRequest{}
	if err := jsonpb.Unmarshal(pbDat, pbTxn); err != nil {
		fmt.Println(err)
	}

	pbOrigDat, _ := jsonpbOrig.Marshal(txn)
	pbOrigTxn := &pb.TxnRequest{}
	if err := jsonpbOrig.Unmarshal(pbOrigDat, pbOrigTxn); err != nil {
		fmt.Println(err)
	}
	fmt.Println("jsonBuiltin json:", string(bDat))
	fmt.Printf("jsonBuiltin struct: %+v\n", bTxn)

	fmt.Println("jsonPB json:", string(pbDat))
	fmt.Printf("jsonPB struct: %+v\n", pbTxn)

	fmt.Println("jsonPBOrig json:", string(pbOrigDat))
	fmt.Printf("jsonPBOrig struct: %+v\n", pbOrigTxn)
}
