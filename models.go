package main

//WSMessage is the message format from WS service
type WSMessage struct {
	Topic   string  `json:"topic"`
	Ref     string  `json:"ref"`
	Payload Machine `json:"payload"`
	JoinRef string  `json:"join_ref"`
	Event   string  `json:"event"`
}

//Machine is the machine model
type Machine struct {
	Timestamp string `json:"timestamp"`
	Status    string `json:"status"`
	MachineID string `json:"machine_id"`
	ID        string `json:"id"`
}
