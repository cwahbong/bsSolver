package bs

import (
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"net/http"
)

/**
  We must export:
  1. Classes of service, args, and reply.
  2. Rpc method name.
  Otherwise gorilla will skip that.
*/
type BsService struct{}

type BsSolveArgs struct {
	Board [][]int `json:"board"`
}

type BsSolveReply struct {
	Lines []Line `json:"lines"`
}

func (*BsService) Solve(request *http.Request, args *BsSolveArgs, reply *BsSolveReply) (err error) {
	reply.Lines, _ = Solve(args.Board)
	return
}

func RpcServer() http.Handler {
	rpcServer := rpc.NewServer()
	rpcServer.RegisterCodec(json.NewCodec(), "application/json")
	rpcServer.RegisterService(new(BsService), "bs")
	return rpcServer
}
