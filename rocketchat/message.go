package rocketchat

type Message struct {
	ID       string   `json:"_id"`
	T        string   `json:"t"`
	Ts       string   `json:"ts"`
	Rid      string   `json:"rid"`
	Msg      string   `json:"msg"`
	URL      []string `json:"url"`
	ExpireAt string   `json:"expireAt"`
	Mentions []string `json:"mentions"`
	U        User     `json:"u"`
	V        Visitor  `json:"v"`
}
