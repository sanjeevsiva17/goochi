package configs

type MockConfig struct {
	Data map[string]string
}

func (m MockConfig) Get(key string) string {
	return m.Data[key]
}

func (m MockConfig) GetOrDefault(key, defaultVal string) string {
	if v, ok := m.Data[key]; ok {
		return v
	}

	return defaultVal
}

func (m MockConfig) ExpandEnv(key string) string {
	return m.Get(key)
}
