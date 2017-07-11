package main

import (
	"context"
	"log"

	"github.com/coreos/etcd/clientv3"
)

func main() {
	c, _ := clientv3.New(clientv3.Config{Endpoints: []string{"http://localhost:2379"}})

	log.Println("writing keys")

	c.Put(context.TODO(), "mainkey", "a")
	c.Put(context.TODO(), "pfx/abc", "1")
	c.Put(context.TODO(), "pfx/def", "2")
	c.Put(context.TODO(), "pfx/xyz", "3")

	log.Println("submitting txn")
	cmp := clientv3.Version("pfx/").WithPrefix()
	txnresp, _ := c.Txn(context.TODO()).
		If(clientv3.Compare(cmp, "=", 0)).
		Then(clientv3.OpDelete("mainkey")).
		Commit()
	if txnresp.Succeeded {
		log.Fatal("deleted mainkey but pfx/ has keys")
	}
	log.Println("did not delete mainkey")

	log.Println("deleting prefix")
	c.Delete(context.TODO(), "pfx/", clientv3.WithPrefix())

	log.Println("submitting txn")
	txnresp, _ = c.Txn(context.TODO()).
		If(clientv3.Compare(cmp, "=", 0)).
		Then(clientv3.OpDelete("mainkey")).
		Commit()
	if !txnresp.Succeeded {
		log.Fatal("failed ot delete mainkey when pfx/ empty")
	}
	log.Println("deleted mainkey")
}
