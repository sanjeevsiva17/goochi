package configs

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	tempDir := os.TempDir()

	_ = os.Chdir(tempDir)
	fd, err := os.Create(".env")
	if err != nil {
		t.Error(err)

		return
	}

	_, err = fd.Write([]byte(`
		NAME: test
		VERSION= 1
		#DATABASE: sql
	`))
	if err != nil {
		t.Error(err)

		return
	}

	c := NewConfigProvider(false, tempDir+"/.env", tempDir+"/.test.env")

	testCases := []struct {
		envName     string
		expectedVal string
	}{
		{"NAME", "test"},
		{"VERSION", "1"},
	}

	for i, tc := range testCases {
		val := c.Get(tc.envName)

		if val != tc.expectedVal {
			t.Errorf("FAILED[%v], expected: %v \t got: %v", i, tc.expectedVal, val)
		}
	}

}

func TestOverLoad(t *testing.T) {
	os.Setenv("TEST", "false")
	defer os.Unsetenv("TEST")

	tempDir := os.TempDir()

	_ = os.Chdir(tempDir)
	fd, err := os.Create(".env")
	if err != nil {
		t.Error(err)

		return
	}

	_, err = fd.Write([]byte(`
		NAME: test
		VERSION= 1
		#DATABASE: sql
	`))
	if err != nil {
		t.Error(err)

		return
	}

	c := NewConfigProvider(true, tempDir+"/.env", tempDir+"/.test.env")

	testCases := []struct {
		envName     string
		expectedVal string
	}{
		{"NAME", "test"},
		{"VERSION", "1"},
		{"TEST", "false"},
	}

	for i, tc := range testCases {
		val := c.Get(tc.envName)

		if val != tc.expectedVal {
			t.Errorf("FAILED[%v], expected: %v \t got: %v", i, tc.expectedVal, val)
		}
	}

}

func Test_config_GetOrDefault(t *testing.T) {
	os.Setenv("TEST", "true")
	defer os.Unsetenv("TEST")
	tests := []struct {
		name       string
		key        string
		defaultVal string
		want       string
	}{
		{"no default specified", "NAME", "", ""},
		{"default specified", "TESTING", "test", "test"},
		{"env exists", "TEST", "", "true"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &config{}
			if got := c.GetOrDefault(tt.key, tt.defaultVal); got != tt.want {
				t.Errorf("GetOrDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_config_Set_ExpandEnv(t *testing.T) {
	os.Setenv("NAME", "test")
	os.Setenv("DB_NAME", "DB_${NAME}")

	defer func() {
		os.Unsetenv("NAME")
		os.Unsetenv("DB_NAME")
	}()

	c := config{}

	val := c.ExpandEnv("DB_NAME")
	if val != "DB_test" {
		t.Errorf("Expected: DB_test, got: %v", val)
	}
}
