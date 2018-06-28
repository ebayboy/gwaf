package main

import (
	"database/sql"
	"fmt"
	"github.com/cihub/seelog"
	"github.com/flier/gohs/hyperscan"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/golang/groupcache"
	proto "github.com/golang/protobuf/proto"
	"github.com/kardianos/service"
	"github.com/valyala/fasthttp"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

func main() {
	fmt.Println("vim-go")
}
