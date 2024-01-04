package notification

type INotification interface {
	Send() error
}

type Notification struct {
	To      string
	Content interface{}
	UserID  int64
}
