package api

const (
	ChannelText = iota
)

const (
	UserStatusActive = iota
	UserStatusBusy
	UserStatusAway
	UserStatusOffline
)

type Channel struct {
	ID    int64
	Name  string
	Topic string

	VoiceConnectedUsers []int64 // ID of Users

	MessageCount   int
	MessagesRead   int
	MessageHistory []Message // Maximum of 50 at a time
}

type Message struct {
	ID        int64
	User      int64 // ID of user
	Message   string
	Timestamp int64
}

type Category struct {
	ID       int64
	Name     string
	Channels []int64 // ID of channels
}

type User struct {
	ID         int64
	Name       string
	Identifier int16
	Status     int8
}

type Image struct {
	File string // Nullable
	Hash string
}

type Icon struct {
	ID   int64
	Name string

	Emote string // Nullable
	Image Image  // Nullable
}

type Server struct {
	ID   int64
	Name string

	Categories []int64 // ID of Categories
	Users      []int64 // ID of users

	Logo int64 // ID of the Icon
}
