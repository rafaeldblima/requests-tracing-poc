package domain

// Message representa uma mensagem do sistema.
type Message struct {
	RequestID   string
	AccountID   string
	SessionID   string
	Request     string
	Response    string
	TargetURL   string
	ServiceName string
	Timestamp   string
}
