package raft

type Options struct {
	DataDir        string // data directory
	HttpAddress    string // http server address
	RaftTCPAddress string // construct Raft Address
	Bootstrap      bool   // start as master or not
	JoinAddress    string // peer address to join
}

func NewOptions() *Options {
	opts := &Options{}

	//var HttpAddress = flag.String("http", "127.0.0.1:6000", "Http address")
	//var RaftTCPAddress = flag.String("RaftNode", "127.0.0.1:7000", "RaftNode tcp address")
	//var node = flag.String("node", "node1", "RaftNode node name")
	//var Bootstrap = flag.Bool("Bootstrap", false, "start as RaftNode cluster")
	//var JoinAddress = flag.String("join", "", "join address for RaftNode cluster")
	//flag.Parse()
	//
	//Opts.DataDir = "./" + *node
	//Opts.HttpAddress = *HttpAddress
	//Opts.Bootstrap = *Bootstrap
	//Opts.RaftTCPAddress = *RaftTCPAddress
	//Opts.JoinAddress = *JoinAddress
	//return Opts

	opts.DataDir = "./node"
	opts.HttpAddress = "127.0.0.1:2869"
	opts.Bootstrap = true
	opts.RaftTCPAddress = "127.0.0.1:2868"
	opts.JoinAddress = ""
	return opts
}
