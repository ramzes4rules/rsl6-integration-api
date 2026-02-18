# RS Loyalty API Client

Go-клиент для работы с RS Loyalty API v2. Библиотека предоставляет удобный интерфейс для интеграции с системой лояльности.

## Установка

```bash
go get github.com/rsl6/loyalty-client
```

## Быстрый старт

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/google/uuid"
    "github.com/rsl6/loyalty-client/client"
    "github.com/rsl6/loyalty-client/models"
)

func main() {
    // Создание клиента
    c := client.NewClient(&client.Config{
        BaseURL: "https://api.loyalty.example.com",
    })

    // Получение клиента по ID
    customerID := uuid.MustParse("11111111-1111-1111-1111-111111111111")
    customer, err := c.Customers.GetByID(context.Background(), customerID)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Customer: %s %s\n", *customer.FirstName, *customer.LastName)

    // Получение баланса клиента
    balance, err := c.Customers.GetBalanceByID(context.Background(), customerID)
    if err != nil {
        log.Fatal(err)
    }
    for _, b := range balance.Balances.Balances {
        fmt.Printf("Currency %s: %.2f\n", b.CurrencyID, b.Value)
    }
}
```

## Структура проекта

```
├── client/           # HTTP-клиент и сервисы
│   ├── client.go     # Основной клиент
│   ├── accounts.go   # Сервис работы со счетами
│   ├── countries.go  # Сервис работы со странами
│   ├── currencies.go # Сервис работы с валютами
│   ├── customers.go  # Сервис работы с клиентами
│   └── loyalty_cards.go # Сервис работы с картами лояльности
├── models/           # Модели данных
│   └── models.go     # Все структуры данных
├── mock/             # Mock-сервер для тестирования
│   ├── server.go     # Основной mock-сервер
│   └── handlers_*.go # Обработчики запросов
├── tests/            # Тесты
│   ├── config.go     # Конфигурация тестов
│   └── *_test.go     # Тестовые файлы
├── go.mod            # Go модуль
└── README.md         # Документация
```

## API Сервисы

### Customers (Клиенты)

```go
// Получить клиента по ID
customer, err := c.Customers.GetByID(ctx, customerID)

// Получить клиента по email или телефону
customer, err := c.Customers.GetByCommunicationValue(ctx, &models.GetByCommunicationValueRequest{
    CommunicationValueType: models.CommunicationEmail,
    Value:                  "email@example.com",
})

// Получить список клиентов
list, err := c.Customers.GetList(ctx, &models.GetListRequest{
    Take: &take,
    Skip: &skip,
})

// Получить баланс клиента
balance, err := c.Customers.GetBalanceByID(ctx, customerID)

// Переименовать клиента
err := c.Customers.Rename(ctx, &models.RenameCustomerRequest{
    BaseCommand: models.BaseCommand{ID: customerID},
    FirstName:   &newFirstName,
    LastName:    &newLastName,
}, nil)

// Установить email или телефон
err := c.Customers.SetCommunicationValue(ctx, &models.SetCustomerCommunicationValueRequest{
    BaseCommand: models.BaseCommand{ID: customerID},
    Value:       "new@email.com",
    ValueType:   models.CommunicationEmail,
}, nil)

// Подписать на рассылку
err := c.Customers.AllowSubscription(ctx, &models.AllowSubscriptionRequest{
    BaseCommand:      models.BaseCommand{ID: customerID},
    SubscriptionType: models.SubscriptionEmailMailing,
}, nil)

// Отписать от рассылки
err := c.Customers.DisallowSubscription(ctx, &models.DisallowSubscriptionRequest{
    BaseCommand:      models.BaseCommand{ID: customerID},
    SubscriptionType: models.SubscriptionEmailMailing,
}, nil)
```

### Accounts (Счета / Начисления)

```go
// Начислить бонусы клиенту
err := c.Accounts.AccrualToCustomer(ctx, &models.AccrualToCustomerRequest{
    BaseCommand: models.BaseCommand{ID: uuid.New()},
    CustomerID:  customerID,
    CurrencyID:  currencyID,
    Amount:      100.0,
}, nil)

// Начислить бонусы на карту лояльности
err := c.Accounts.AccrualToLoyaltyCard(ctx, &models.AccrualToLoyaltyCardRequest{
    BaseCommand:   models.BaseCommand{ID: uuid.New()},
    LoyaltyCardID: cardID,
    CurrencyID:    currencyID,
    Amount:        100.0,
}, nil)

// Списать бонусы с клиента
err := c.Accounts.SubtractFromCustomer(ctx, &models.SubtractFromCustomerRequest{
    BaseCommand: models.BaseCommand{ID: uuid.New()},
    CustomerID:  customerID,
    CurrencyID:  currencyID,
    Amount:      50.0,
}, nil)

// Списать бонусы с карты лояльности
err := c.Accounts.SubtractFromLoyaltyCard(ctx, &models.SubtractFromLoyaltyCardRequest{
    BaseCommand:   models.BaseCommand{ID: uuid.New()},
    LoyaltyCardID: cardID,
    CurrencyID:    currencyID,
    Amount:        50.0,
}, nil)

// Получить транзакции
transactions, err := c.Accounts.GetTransactions(ctx, &models.GetTransactionsRequest{
    FromDate: &fromDate,
    ToDate:   &toDate,
})
```

### LoyaltyCards (Карты лояльности)

```go
// Получить карту по ID
card, err := c.LoyaltyCards.GetByID(ctx, cardID)

// Получить карту по номеру
card, err := c.LoyaltyCards.GetByNumber(ctx, "1234567890")

// Получить список карт
list, err := c.LoyaltyCards.GetList(ctx, &models.GetListRequest{})

// Получить баланс карты
balance, err := c.LoyaltyCards.GetBalanceByID(ctx, cardID)

// Активировать карту
err := c.LoyaltyCards.Activate(ctx, cardID, nil)

// Заблокировать карту
err := c.LoyaltyCards.Block(ctx, cardID, nil)

// Разблокировать карту
err := c.LoyaltyCards.Unblock(ctx, cardID, nil)
```

### Countries (Страны)

```go
// Создать страну
err := c.Countries.Create(ctx, &models.CreateCountryRequest{
    BaseCommand: models.BaseCommand{ID: uuid.New()},
    Code:        "RU",
}, nil)

// Получить страну по ID
country, err := c.Countries.GetByID(ctx, countryID)

// Получить список стран
list, err := c.Countries.GetList(ctx, &models.GetListRequest{})

// Установить код страны
err := c.Countries.SetCode(ctx, &models.SetCountryCodeRequest{
    BaseCommand: models.BaseCommand{ID: countryID},
    Code:        "RUS",
}, nil)

// Удалить страну
err := c.Countries.Delete(ctx, &models.DeleteCountryRequest{
    BaseCommand: models.BaseCommand{ID: countryID},
}, nil)

// Восстановить страну
err := c.Countries.Restore(ctx, &models.RestoreCountryRequest{
    BaseCommand: models.BaseCommand{ID: countryID},
}, nil)
```

### Currencies (Валюты)

```go
// Создать валюту
err := c.Currencies.Create(ctx, &models.CreateCurrencyRequest{
    BaseCommand: models.BaseCommand{ID: uuid.New()},
    Name:        "BONUS",
}, nil)

// Получить валюту по ID
currency, err := c.Currencies.GetByID(ctx, currencyID)

// Получить список валют
list, err := c.Currencies.GetList(ctx, &models.GetListRequest{})

// Переименовать валюту
err := c.Currencies.Rename(ctx, &models.RenameCurrencyRequest{
    BaseCommand: models.BaseCommand{ID: currencyID},
    Name:        "POINTS",
}, nil)

// Установить курс валюты
err := c.Currencies.SetRate(ctx, &models.SetCurrencyRateRequest{
    BaseCommand: models.BaseCommand{ID: currencyID},
    Rate:        1.5,
}, nil)

// Активировать валюту
err := c.Currencies.Activate(ctx, &models.ActivateCurrencyRequest{
    BaseCommand: models.BaseCommand{ID: currencyID},
}, nil)

// Деактивировать валюту
err := c.Currencies.Deactivate(ctx, &models.DeactivateCurrencyRequest{
    BaseCommand: models.BaseCommand{ID: currencyID},
}, nil)
```

## Конфигурация клиента

```go
// Создание клиента с настройками
c := client.NewClient(&client.Config{
    BaseURL: "https://api.loyalty.example.com",
    Timeout: 30 * time.Second,
    Headers: map[string]string{
        "Authorization": "Bearer your-token",
    },
})

// Установка заголовка после создания
c.SetHeader("X-Custom-Header", "value")
```

## Передача заголовков в запросах

Для команд (операций изменения данных) можно передавать дополнительные заголовки:

```go
headers := &models.RequestHeaders{
    CommandID:          stringPtr("unique-command-id"),
    OperationDate:      stringPtr("2024-01-15T10:30:00+03:00"),
    UserID:             stringPtr("operator-123"),
    InteractionChannel: stringPtr("API"),
}

err := c.Customers.Rename(ctx, req, headers)
```

## Тестирование

### Запуск тестов с mock-сервером (по умолчанию)

```bash
cd tests
go test -v ./...
```

### Запуск тестов с реальным сервером

```bash
# Установить переменные окружения
$env:USE_MOCK_SERVER = "false"
$env:REAL_SERVER_URL = "https://api.loyalty.example.com"

# Запустить тесты
cd tests
go test -v ./...
```

### Переменные окружения для тестов

| Переменная | Описание | Значение по умолчанию |
|------------|----------|----------------------|
| `USE_MOCK_SERVER` | Использовать mock-сервер | `true` |
| `REAL_SERVER_URL` | URL реального сервера | `http://localhost:8080` |

## Mock-сервер

Библиотека включает встроенный mock-сервер для тестирования. Он содержит предзаполненные тестовые данные:

### Тестовые данные

| Сущность | ID | Описание |
|----------|-----|----------|
| Customer | `11111111-1111-1111-1111-111111111111` | Тестовый клиент "Иван Петров" |
| Country | `22222222-2222-2222-2222-222222222222` | Страна "RU" |
| Currency | `33333333-3333-3333-3333-333333333333` | Валюта "BONUS" |
| LoyaltyCard | `44444444-4444-4444-4444-444444444444` | Карта "1234567890" |

### Использование mock-сервера в своих тестах

```go
package mytest

import (
    "net/http/httptest"
    "testing"

    "github.com/rsl6/loyalty-client/client"
    "github.com/rsl6/loyalty-client/mock"
)

func TestMyFeature(t *testing.T) {
    // Создать mock-сервер
    mockServer := mock.NewServer()
    httpServer := httptest.NewServer(mockServer)
    defer httpServer.Close()

    // Создать клиент
    c := client.NewClient(&client.Config{
        BaseURL: httpServer.URL,
    })

    // Использовать клиент для тестов
    // ...
}
```

### Добавление данных в mock-сервер

```go
mockServer := mock.NewServer()

// Добавить клиента
mockServer.AddCustomer(&models.CustomerDto{
    ID:        uuid.New(),
    FirstName: stringPtr("Test"),
    Gender:    models.GenderMale,
    // ...
})

// Добавить баланс
mockServer.SetBalance(customerID, []models.BalanceDto{
    {CurrencyID: currencyID, Value: 1000.0},
})

// Сбросить данные к начальному состоянию
mockServer.Reset()
```

## Обработка ошибок

```go
customer, err := c.Customers.GetByID(ctx, customerID)
if err != nil {
    if apiErr, ok := err.(*models.APIError); ok {
        fmt.Printf("API Error: %d - %s\n", apiErr.StatusCode, apiErr.Message)
        fmt.Printf("Details: %s\n", apiErr.Details)
    } else {
        fmt.Printf("Error: %v\n", err)
    }
}
```

## Модели данных

### Перечисления

```go
// Типы пола
models.GenderMale    // "Male"
models.GenderFemale  // "Female"
models.GenderUnknown // "Unknown"

// Типы коммуникации
models.CommunicationPhone // "Phone"
models.CommunicationEmail // "Email"

// Типы подписок
models.SubscriptionEmailMailing        // "EmailMailing"
models.SubscriptionSmsMailing          // "SmsMailing"
models.SubscriptionPushMailing         // "PushMailing"
models.SubscriptionViberMailing        // "ViberMailing"
models.SubscriptionPhoneCallMailing    // "PhoneCallMailing"
models.SubscriptionClientPortalMailing // "ClientPortalMailing"

// Статусы карт лояльности
models.LoyaltyCardStatusActive   // "Active"
models.LoyaltyCardStatusBlocked  // "Blocked"
models.LoyaltyCardStatusCreated  // "Created"
models.LoyaltyCardStatusIssued   // "Issued"
models.LoyaltyCardStatusExpired  // "Expired"
// ... и другие

// Типы транзакций
models.TransactionAccrual  // "Accrual"
models.TransactionSubtract // "Subtract"
models.TransactionBurning  // "Burning"
```

## Лицензия

MIT License

