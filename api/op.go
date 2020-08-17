package api

const (
	// OpUnknown is an unknown OpCode and is the default in case our code doesn't exists
	OpUnknown = iota

	// OpHello used on the first connection
	// {  }
	OpHello

	// OpReqAuth request authentification using existing user token
	OpReqAuth
	// OpAckAuth acknowledges authetification and returns it's user id
	// { Snowflake }
	OpAckAuth

	// OpReqServerList request it's own ServerList from our Database
	// {  }
	OpReqServerList
	// OpAckServerList acknowledges the request and returns a server list
	// {  }
	OpAckServerList

	// OpReqUser request it's own User Information from our Database
	// {  }
	OpReqUser
	// OpAckUser acknowledges the request and returns our user information
	// {  }
	OpAckUser

	OpReqServer
	OpAckServer

	OpReqCategory
	OpAckCategory

	OpReqChannel
	OpAckChannel
)

type OpMessage struct {
	Op   int
	Cb   int
	Data []byte
}
