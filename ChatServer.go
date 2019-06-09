package main

import (
	"net"
)

// ChatServer : A wrapper around net.Listener to provide chat convenience methods
type ChatServer struct {
	net.Listener
	channels []*ChatChannel
	users    []*ChatClient
}

func (server *ChatServer) createChannel(name string, user *ChatClient) *ChatChannel {
	channel := ChatChannel{
		name:   name,
		server: server,
		users:  []*ChatClient{user},
	}
	channel.init()
	server.channels = append(server.channels, &channel)
	return &channel
}

func (server *ChatServer) findChannel(name string) *ChatChannel {
	for _, channel := range server.channels {
		if channel.name == name {
			return channel
		}
	}
	return nil
}

func (server *ChatServer) joinChannel(name string, user *ChatClient) *ChatChannel {
	channel := server.findChannel(name)

	if channel == nil {
		channel = server.createChannel(name, user)
	}

	channel.join(user)

	return channel
}
