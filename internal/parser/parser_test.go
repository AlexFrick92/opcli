package parser

import (
	"fmt"
	"testing"
)

// Mock variables for connectCommand and disconnectCommand
var (
	mockConnectCalled   bool
	mockConnectEndpoint string
	mockConnectError    error
	mockDisconnectCalled bool
	mockDisconnectError  error
)

// mockConnect is a mock implementation for connectCommand
func mockConnect(endpoint string) error {
	mockConnectCalled = true
	mockConnectEndpoint = endpoint
	return mockConnectError
}

// mockDisconnect is a mock implementation for disconnectCommand
func mockDisconnect() error {
	mockDisconnectCalled = true
	return mockDisconnectError
}

// resetMocks resets the state of all mock variables
func resetMocks() {
	mockConnectCalled = false
	mockConnectEndpoint = ""
	mockConnectError = nil
	mockDisconnectCalled = false
	mockDisconnectError = nil
}

// TestExecute проверяет функцию Execute для различных входных данных.
// Эта функция является основной точкой входа для обработки команд пользователя.
//
// Основные аспекты тестирования:
// - Обработка пустого ввода.
// - Корректная обработка известных команд (`help`, `exit`, `quit`).
// - Возврат ожидаемой ошибки для неизвестных команд.
// - Возврат ошибки при неверном использовании команды `connect` (например, без аргументов).
// - Корректная обработка команд `connect` и `disconnect` с использованием заглушек.
func TestExecute(t *testing.T) {
	// Сохраняем оригинальные функции и восстанавливаем их после выполнения всех тестов
	oldConnectCommand := connectCommand
	oldDisconnectCommand := disconnectCommand
	defer func() {
		connectCommand = oldConnectCommand
		disconnectCommand = oldDisconnectCommand
	}()

	tests := []struct {
		name                 string
		input                string
		setupMocks           func()
		checkMocks           func(*testing.T)
		wantErr              bool
		wantExit             bool
		errMsg               string
	}{
		{
			name:    "Пустой ввод не должен вызывать ошибок",
			input:   "",
			wantErr: false,
		},
		{
			name:    "Команда help не должна вызывать ошибок",
			input:   "help",
			wantErr: false,
		},
		{
			name:     "Команда exit должна вернуть ошибку 'exit'",
			input:    "exit",
			wantErr:  true,
			wantExit: true,
		},
		{
			name:     "Команда quit должна вернуть ошибку 'exit'",
			input:    "quit",
			wantErr:  true,
			wantExit: true,
		},
		{
			name:    "Неизвестная команда должна вернуть соответствующую ошибку",
			input:   "foobar",
			wantErr: true,
			errMsg:  "unknown command: foobar. Type 'help' for available commands",
		},
		{
			name:    "Команда connect без аргументов должна вернуть ошибку использования",
			input:   "connect",
			wantErr: true,
			errMsg:  "usage: connect <endpoint>",
		},
		{
			name:  "Команда connect с эндпоинтом должна вызвать mockConnect",
			input: "connect opc.tcp://localhost:4840",
			setupMocks: func() {
				connectCommand = mockConnect
				mockConnectError = nil // Устанавливаем успешное выполнение
			},
			checkMocks: func(t *testing.T) {
				if !mockConnectCalled {
					t.Errorf("mockConnect не был вызван")
				}
				if mockConnectEndpoint != "opc.tcp://localhost:4840" {
					t.Errorf("mockConnect вызван с неверным эндпоинтом: %s", mockConnectEndpoint)
				}
			},
			wantErr: false,
		},
		{
			name:  "Команда connect с эндпоинтом должна вернуть ошибку от mockConnect",
			input: "connect opc.tcp://remote:4840",
			setupMocks: func() {
				connectCommand = mockConnect
				mockConnectError = fmt.Errorf("mock connect failed")
			},
			checkMocks: func(t *testing.T) {
				if !mockConnectCalled {
					t.Errorf("mockConnect не был вызван")
				}
				if mockConnectEndpoint != "opc.tcp://remote:4840" {
					t.Errorf("mockConnect вызван с неверным эндпоинтом: %s", mockConnectEndpoint)
				}
			},
			wantErr: true,
			errMsg:  "mock connect failed",
		},
		{
			name:  "Команда disconnect должна вызвать mockDisconnect",
			input: "disconnect",
			setupMocks: func() {
				disconnectCommand = mockDisconnect
				mockDisconnectError = nil // Устанавливаем успешное выполнение
			},
			checkMocks: func(t *testing.T) {
				if !mockDisconnectCalled {
					t.Errorf("mockDisconnect не был вызван")
				}
			},
			wantErr: false,
		},
		{
			name:  "Команда disconnect должна вернуть ошибку от mockDisconnect",
			input: "disconnect",
			setupMocks: func() {
				disconnectCommand = mockDisconnect
				mockDisconnectError = fmt.Errorf("mock disconnect failed")
			},
			checkMocks: func(t *testing.T) {
				if !mockDisconnectCalled {
					t.Errorf("mockDisconnect не был вызван")
				}
			},
			wantErr: true,
			errMsg:  "mock disconnect failed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resetMocks() // Сброс моков перед каждым подтестом
			if tt.setupMocks != nil {
				tt.setupMocks()
			}

			err := Execute(tt.input)

			if tt.wantErr {
				if err == nil {
					t.Errorf("Execute() ожидалась ошибка, получено nil")
				} else if tt.wantExit && err.Error() != "exit" {
					t.Errorf("Execute() получено неожиданное сообщение об ошибке выхода = %v", err.Error())
				} else if !tt.wantExit && err.Error() != tt.errMsg {
					t.Errorf("Execute() получено неожиданное сообщение об ошибке = %v, ожидалось %v", err.Error(), tt.errMsg)
				}
			} else {
				if err != nil {
					t.Errorf("Execute() получена непредвиденная ошибка = %v", err)
				}
			}

			if tt.checkMocks != nil {
				tt.checkMocks(t)
			}
		})
	}
}

// TestParseStartupArgs проверяет функцию ParseStartupArgs для различных аргументов запуска.
// Эта функция обрабатывает аргументы командной строки, переданные при старте приложения.
//
// Основные аспекты тестирования:
// - Обработка запуска без функциональных аргументов.
// - Обработка неполных аргументов для команды `connect`.
// - Корректная обработка аргументов для `connect` и IP-адресов с использованием заглушек.
func TestParseStartupArgs(t *testing.T) {
	// Сохраняем оригинальные функции и восстанавливаем их после выполнения всех тестов
	oldConnectCommand := connectCommand
	defer func() {
		connectCommand = oldConnectCommand
	}()

	tests := []struct {
		name                 string
		args                 []string
		setupMocks           func()
		checkMocks           func(*testing.T)
		wantErr              bool
		errMsg               string
	}{
		{
			name:    "Без аргументов после имени программы не должно быть ошибок",
			args:    []string{"opcli"},
			wantErr: false,
		},
		{
			name:    "Недостаточно аргументов для 'connect' должно вернуть ошибку использования",
			args:    []string{"opcli", "connect"},
			wantErr: true,
			errMsg:  "usage: connect <endpoint>",
		},
		{
			name: "Запуск с IP-адресом должен вызвать mockConnect",
			args: []string{"opcli", "127.0.0.1"},
			setupMocks: func() {
				connectCommand = mockConnect
				mockConnectError = nil
			},
			checkMocks: func(t *testing.T) {
				if !mockConnectCalled {
					t.Errorf("mockConnect не был вызван")
				}
				if mockConnectEndpoint != "opc.tcp://127.0.0.1:4840" {
					t.Errorf("mockConnect вызван с неверным эндпоинтом: %s", mockConnectEndpoint)
				}
			},
			wantErr: false,
		},
		{
			name: "Запуск с IP-адресом должен вернуть ошибку от mockConnect",
			args: []string{"opcli", "127.0.0.1"},
			setupMocks: func() {
				connectCommand = mockConnect
				mockConnectError = fmt.Errorf("mock startup connect failed")
			},
			checkMocks: func(t *testing.T) {
				if !mockConnectCalled {
					t.Errorf("mockConnect не был вызван")
				}
			},
			wantErr: true,
			errMsg:  "mock startup connect failed",
		},
		{
			name: "Запуск с 'connect <endpoint>' должен вызвать mockConnect",
			args: []string{"opcli", "connect", "opc.tcp://localhost:4840"},
			setupMocks: func() {
				connectCommand = mockConnect
				mockConnectError = nil
			},
			checkMocks: func(t *testing.T) {
				if !mockConnectCalled {
					t.Errorf("mockConnect не был вызван")
				}
				if mockConnectEndpoint != "opc.tcp://localhost:4840" {
					t.Errorf("mockConnect вызван с неверным эндпоинтом: %s", mockConnectEndpoint)
				}
			},
			wantErr: false,
		},
		{
			name: "Запуск с 'connect <endpoint>' должен вернуть ошибку от mockConnect",
			args: []string{"opcli", "connect", "opc.tcp://invalid:4840"},
			setupMocks: func() {
				connectCommand = mockConnect
				mockConnectError = fmt.Errorf("mock startup connect with endpoint failed")
			},
			checkMocks: func(t *testing.T) {
				if !mockConnectCalled {
					t.Errorf("mockConnect не был вызван")
				}
			},
			wantErr: true,
			errMsg:  "mock startup connect with endpoint failed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resetMocks() // Сброс моков перед каждым подтестом
			if tt.setupMocks != nil {
				tt.setupMocks()
			}

			err := ParseStartupArgs(tt.args)

			if tt.wantErr {
				if err == nil {
					t.Errorf("ParseStartupArgs() ожидалась ошибка, получено nil")
				} else if err.Error() != tt.errMsg {
					t.Errorf("ParseStartupArgs() получено неожиданное сообщение об ошибке = %v, ожидалось %v", err.Error(), tt.errMsg)
				}
			} else {
				if err != nil {
					t.Errorf("ParseStartupArgs() получена непредвиденная ошибка = %v", err)
				}
			}

			if tt.checkMocks != nil {
				tt.checkMocks(t)
			}
		})
	}
}

// TestIsIPv4 проверяет вспомогательную функцию isIPv4.
// Эта функция является неэкспортированной, но её логика достаточно важна и самодостаточна,
// чтобы быть протестированной напрямую.
//
// Основные аспекты тестирования:
// - Корректное определение действительных IPv4-адресов.
// - Корректное определение недействительных форматов, включая IPv6, неполные адреса,
//   адреса с некорректными значениями октетов и текстовые строки.
func TestIsIPv4(t *testing.T) {
	tests := []struct {
		name string
		ip   string
		want bool
	}{
		{
			name: "Действительный IPv4-адрес (loopback)",
			ip:   "127.0.0.1",
			want: true,
		},
		{
			name: "Действительный IPv4-адрес (приватный)",
			ip:   "192.168.1.1",
			want: true,
		},
		{
			name: "Действительный IPv4-адрес (публичный)",
			ip:   "8.8.8.8",
			want: true,
		},
		{
			name: "Недействительный IPv4-адрес (слишком много сегментов)",
			ip:   "1.2.3.4.5",
			want: false,
		},
		{
			name: "Недействительный IPv4-адрес (текст)",
			ip:   "localhost",
			want: false,
		},
		{
			name: "Недействительный IPv4-адрес (пустая строка)",
			ip:   "",
			want: false,
		},
		{
			name: "Недействительный IPv4-адрес (частично действительный)",
			ip:   "192.168.1",
			want: false,
		},
		{
			name: "Недействительный IPv4-адрес (значение октета вне диапазона)",
			ip:   "256.0.0.1",
			want: false,
		},
		{
			name: "Действительный IPv6-адрес (должен быть ложным для isIPv4)",
			ip:   "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isIPv4(tt.ip); got != tt.want {
				t.Errorf("isIPv4(%q) = %v, ожидалось %v", tt.ip, got, tt.want)
			}
		})
	}
}
