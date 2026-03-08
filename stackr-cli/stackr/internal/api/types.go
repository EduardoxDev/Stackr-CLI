package api

type App struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Status    string `json:"status"`
	MemoryMb  int    `json:"memoryMb"`
	RAM       string `json:"ram"`
	Language  string `json:"language"`
	Type      string `json:"type"`
	OOMKilled bool   `json:"oomKilled"`
}

type AppStats struct {
	Running   bool   `json:"running"`
	CPU       string `json:"cpu"`
	RAM       string `json:"ram"`
	NetworkRx string `json:"networkRx"`
	NetworkTx string `json:"networkTx"`
}

type AppLogs struct {
	Logs []string `json:"logs"`
}

type MessageResp struct {
	Message string `json:"message"`
	BotID   string `json:"botId"`
	Status  string `json:"status"`
}

type Database struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Engine           string `json:"engine"`
	Status           string `json:"status"`
	MemoryMb         int    `json:"memoryMb"`
	Host             string `json:"host"`
	Port             int    `json:"port"`
	Database         string `json:"database"`
	Username         string `json:"username"`
	Password         string `json:"password"`
	ConnectionString string `json:"connectionString"`
	CreatedAt        string `json:"createdAt"`
}

type DBStats struct {
	Running    bool   `json:"running"`
	CPUPercent string `json:"cpuPercent"`
	MemUsage   string `json:"memUsage"`
	MemPercent string `json:"memPercent"`
	NetRx      string `json:"netRx"`
	NetTx      string `json:"netTx"`
	PIDs       int    `json:"pids"`
}

type DBLogs struct {
	Logs string `json:"logs"`
}

type FileEntry struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Size int64  `json:"size"`
}

type FileContent struct {
	Content string `json:"content"`
}
