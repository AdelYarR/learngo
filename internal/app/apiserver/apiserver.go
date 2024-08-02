package apiserver

// APIServer ...
type APIServer struct{
	config *Config
}

// Инициализация APIServer ...
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
	}
}

// Запуск HTTP Сервера, подключение к БД ...
func (s *APIServer) Start() error {
	return nil
}