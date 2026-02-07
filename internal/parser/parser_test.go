package parser

import (

	"testing"


)

// TestExecute проверяет функцию Execute для различных входных данных.
// Эта функция является основной точкой входа для обработки команд пользователя.
//
// Основные аспекты тестирования:
// - Обработка пустого ввода.
// - Корректная обработка известных команд (`help`, `exit`, `quit`).
// - Возврат ожидаемой ошибки для неизвестных команд.
// - Возврат ошибки при неверном использовании команды `connect` (например, без аргументов).
//
// Ограничения:
// Тестирование команд, которые инициируют фактические сетевые подключения (`connect <endpoint>`, `disconnect`),
// требует заглушек (mocks) для функций `commands.Connect` и `commands.Disconnect`.
// Без модификации пакета `commands` (например, чтобы сделать `Connect` и `Disconnect` переменными)
// невозможно изолировать эти тесты от реальных сетевых операций.
// Поэтому такие сценарии здесь не тестируются.
func TestExecute(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantErr  bool
		wantExit bool
		errMsg   string
	}{
		{
			name:    "Пустой ввод не должен вызывать ошибок",
			input:   "",
			wantErr: false,
		},
		{
			name:    "Команда help не должна вызывать ошибок",
			input:   "help",
			wantErr: false, // PrintHelp просто печатает в консоль, не возвращает ошибок
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
		// Сценарии `connect <endpoint>` и `disconnect` здесь не тестируются, см. Ограничения выше.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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
		})
	}
}

// TestParseStartupArgs проверяет функцию ParseStartupArgs для различных аргументов запуска.
// Эта функция обрабатывает аргументы командной строки, переданные при старте приложения.
//
// Основные аспекты тестирования:
// - Обработка запуска без функциональных аргументов.
// - Обработка неполных аргументов для команды `connect`.
//
// Ограничения:
// Сценарии, которые приводят к фактическому вызову `commands.Connect`
// (например, `opcli 127.0.0.1` или `opcli connect opc.tcp://...`), не тестируются напрямую.
// Это связано с тем, что `commands.Connect` является функцией, а не переменной,
// и не может быть легко переопределена для целей тестирования без модификации пакета `commands`.
// Полноценное тестирование этих сценариев потребует рефакторинга для обеспечения тестопригодности (например, внедрение зависимостей через интерфейсы).
func TestParseStartupArgs(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantErr bool
		errMsg  string
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
			errMsg:  "usage: connect <endpoint>", // Эта ошибка происходит из handleConnect, вызываемого ParseStartupArgs
		},
		// Сценарии с действительным IP или точкой подключения для `connect` здесь не тестируются, см. Ограничения выше.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// В этом тесте мы не переопределяем commands.Connect, так как это невозможно
			// без изменения пакета commands. Поэтому тестируются только те сценарии,
			// которые не приводят к фактическому вызову commands.Connect или
			// возвращают ошибку раньше.

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
