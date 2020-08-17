package api

func (client *APIClient) SendAuthComplete(cb int, user *User) {
	client.Send(OpAckAuth, cb, user.ID)
}

func (client *APIClient) SendServerList(cb int, serverList []*Server) {
	var srvList []int64
	for _, server := range serverList {
		srvList = append(srvList, server.ID)
	}

	client.Send(OpAckServerList, cb, srvList)
}

func (client *APIClient) SendUser(cb int, user *User) {
	client.Send(OpAckUser, cb, *user)
}

func (client *APIClient) SendServer(cb int, server *Server) {
	client.Send(OpAckServer, cb, *server)
}

func (client *APIClient) SendCategory(cb int, category *Category) {
	client.Send(OpAckCategory, cb, *category)
}

func (client *APIClient) SendChannel(cb int, channel *Channel) {
	client.Send(OpAckChannel, cb, *channel)
}
