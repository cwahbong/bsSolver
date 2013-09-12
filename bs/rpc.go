package bs

import (
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"net/http"
)

type bsService struct{}

type bsSolveArgs struct {
	Board [][]int `json:"board"`
}

type bsSolveReply struct {
	Lines []Line `json:"lines"`
}

func (*bsService) Solve(request *http.Request, args *bsSolveArgs, reply *bsSolveReply) (err error) {
	reply.Lines, _ = Solve(args.Board)
	return
}

func RpcServer() http.Handler {
	rpcServer := rpc.NewServer()
	rpcServer.RegisterCodec(json.NewCodec(), "application/json")
	rpcServer.RegisterService(new(bsService), "bs")
	return rpcServer
}
