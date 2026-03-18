# RSLoyalty

Go-клиент для работы с RS Loyalty API v2. Библиотека предоставляет удобный интерфейс для интеграции с системой лояльности.

## Установка

```bash
go get github.com/rsl6/rsloyalty
```

## Быстрый старт

```go
package main

import (
	"context"
	"fmt"
	"log"
	"net/http/httptest"
	"time"

	"github.com/ramzes4rules/rsl6-integration-api/client"
	"github.com/ramzes4rules/rsl6-integration-api/mock"
	"github.com/ramzes4rules/rsl6-integration-api/models"
)

func main() {
	//Локальный mock-сервер позволяет сделать рабочий запрос без внешних зависимостей.
	server := httptest.NewServer(mock.NewServer())
	defer server.Close()

	cfg := client.DefaultConfig()
	cfg.BaseURL = server.URL
	api := client.NewClient(cfg)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := api.Countries.GetList(ctx, &models.GetListRequest{})
	if err != nil {
		log.Fatalf("countries get_list failed: %v", err)
	}

	fmt.Printf("Countries total: %d\n", result.Total)
	if len(result.Values) > 0 {
		fmt.Printf("First country: id=%s code=%s\n", result.Values[0].ID, result.Values[0].Code)
	}
}

```

## Структура проекта

```
├── client/                          # HTTP-клиент и сервисы
│   ├── client.go                    # Основной клиент с логированием
│   ├── accounts.go                  # Сервис работы со счетами
│   ├── countries.go                 # Сервис работы со странами
│   ├── currencies.go                # Сервис работы с валютами
│   ├── customers.go                 # Сервис работы с клиентами
│   ├── customer_cards.go            # Сервис привязки карт к клиентам
│   ├── customer_property.go         # Сервис свойств клиентов
│   ├── customer_segment.go          # Сервис сегментов клиентов
│   ├── loyalty_cards.go             # Сервис работы с картами лояльности
│   ├── loyalty_card_block_reason.go # Причины блокировки карт лояльности
│   ├── loyalty_card_group.go        # Группы карт лояльности
│   ├── loyalty_card_issue_reason.go # Причины выпуска карт лояльности
│   ├── gift_cards.go                # Сервис подарочных карт
│   ├── gift_card_block_reason.go    # Причины блокировки подарочных карт
│   ├── gift_card_group.go           # Группы подарочных карт
│   ├── gift_card_issue_reason.go    # Причины выпуска подарочных карт
│   ├── sponsored_cards.go           # Сервис спонсорских карт
│   ├── sponsored_card_block_reason.go # Причины блокировки спонсорских карт
│   ├── sponsored_card_group.go      # Группы спонсорских карт
│   ├── sponsored_card_issue_reason.go # Причины выпуска спонсорских карт
│   ├── sponsored_card_owner.go      # Владельцы спонсорских карт
│   ├── external_gift_card_series.go # Внешние серии подарочных карт
│   ├── external_loyalty_card_series.go # Внешние серии карт лояльности
│   ├── external_sponsored_card_series.go # Внешние серии спонсорских карт
│   ├── items.go                     # Сервис товаров
│   ├── item_categories.go           # Категории товаров
│   ├── item_groups.go               # Группы товаров
│   ├── item_properties.go           # Свойства товаров
│   ├── stores.go                    # Сервис магазинов
│   ├── store_cluster.go             # Кластеры магазинов
│   ├── store_format.go              # Форматы магазинов
│   ├── store_property.go            # Свойства магазинов
│   ├── pos.go                       # Сервис касс (POS)
│   ├── pos_type.go                  # Типы касс
│   ├── opening_hours.go             # Часы работы
│   ├── segment_group.go             # Группы сегментов
│   ├── static_segment.go            # Статические сегменты
│   ├── territorial_division.go      # Территориальные подразделения
│   └── manual_accrual_reason.go     # Причины ручного начисления
├── models/                          # Модели данных
│   ├── common.go                    # Общие типы, enums, базовые запросы
│   ├── accounts.go                  # Модели счетов и транзакций
│   ├── countries.go                 # Модели стран
│   ├── currencies.go                # Модели валют
│   ├── customers.go                 # Модели клиентов
│   ├── customer_cards.go            # Модели привязки карт
│   ├── customer_property.go         # Модели свойств клиентов
│   ├── customer_segment.go          # Модели сегментов клиентов
│   ├── loyalty_cards.go             # Модели карт лояльности
│   ├── gift_cards.go                # Модели подарочных карт
│   ├── sponsored_cards.go           # Модели спонсорских карт
│   ├── sponsored_card_owner.go      # Модели владельцев спонсорских карт
│   ├── external_card_series.go      # Модели внешних серий карт
│   ├── items.go                     # Модели товаров
│   ├── item_categories.go           # Модели категорий товаров
│   ├── item_groups.go               # Модели групп товаров
│   ├── item_properties.go           # Модели свойств товаров
│   ├── stores.go                    # Модели магазинов
│   ├── store_cluster.go             # Модели кластеров магазинов
│   ├── store_format.go              # Модели форматов магазинов
│   ├── store_property.go            # Модели свойств магазинов
│   ├── pos.go                       # Модели касс
│   ├── pos_type.go                  # Модели типов касс
│   ├── opening_hours.go             # Модели часов работы
│   ├── segment_group.go             # Модели групп сегментов
│   ├── static_segment.go            # Модели статических сегментов
│   └── territorial_division.go      # Модели территориальных подразделений
├── mock/                            # Mock-сервер для тестирования
│   ├── server.go                    # Основной mock-сервер
│   └── handlers_*.go                # Обработчики запросов
├── tests/                           # Тесты
│   ├── config.go                    # Конфигурация тестов
│   └── *_test.go                    # Тестовые файлы
├── go.mod                           # Go модуль
└── README.md                        # Документация
```

## Доступные сервисы

| Сервис | Описание |
|--------|----------|
| `Accounts` | Начисления и списания бонусов |
| `Countries` | Управление странами |
| `Currencies` | Управление валютами |
| `Customers` | Управление клиентами |
| `CustomerCards` | Привязка карт к клиентам |
| `CustomerProperty` | Свойства клиентов |
| `CustomerSegment` | Сегменты клиентов |
| `LoyaltyCards` | Карты лояльности |
| `LoyaltyCardBlockReason` | Причины блокировки карт лояльности |
| `LoyaltyCardGroup` | Группы карт лояльности |
| `LoyaltyCardIssueReason` | Причины выпуска карт лояльности |
| `GiftCards` | Подарочные карты |
| `GiftCardBlockReason` | Причины блокировки подарочных карт |
| `GiftCardGroup` | Группы подарочных карт |
| `GiftCardIssueReason` | Причины выпуска подарочных карт |
| `SponsoredCards` | Спонсорские карты |
| `SponsoredCardBlockReason` | Причины блокировки спонсорских карт |
| `SponsoredCardGroup` | Группы спонсорских карт |
| `SponsoredCardIssueReason` | Причины выпуска спонсорских карт |
| `SponsoredCardOwner` | Владельцы спонсорских карт |
| `ExternalGiftCardSeries` | Внешние серии подарочных карт |
| `ExternalLoyaltyCardSeries` | Внешние серии карт лояльности |
| `ExternalSponsoredCardSeries` | Внешние серии спонсорских карт |
| `Items` | Товары |
| `ItemCategories` | Категории товаров |
| `ItemGroups` | Группы товаров |
| `ItemProperties` | Свойства товаров |
| `Store` | Магазины |
| `StoreCluster` | Кластеры магазинов |
| `StoreFormat` | Форматы магазинов |
| `StoreProperty` | Свойства магазинов |
| `Pos` | Кассы (POS) |
| `PosType` | Типы касс |
| `OpeningHours` | Часы работы |
| `SegmentGroup` | Группы сегментов |
| `StaticSegment` | Статические сегменты |
| `TerritorialDivision` | Территориальные подразделения |
| `ManualAccrualReason` | Причины ручного начисления |

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

### Store (Магазины)

```go
// Создать магазин
err := c.Store.Create(ctx, &models.CreateStoreRequest{
    BaseCommand: models.BaseCommand{ID: uuid.New()},
    Name:        "Магазин #1",
}, nil)

// Получить магазин по ID
store, err := c.Store.GetByID(ctx, &models.GetByIdRequest{ID: storeID}, nil)

// Получить список магазинов
list, err := c.Store.GetList(ctx, &models.GetListRequest{}, nil)

// Установить адрес магазина
err := c.Store.SetAddress(ctx, &models.SetStoreAddressRequest{
    BaseCommand: models.BaseCommand{ID: storeID},
    Address:     "ул. Примерная, д. 1",
}, nil)

// Установить координаты
err := c.Store.SetLocationCoordinates(ctx, &models.SetStoreLocationCoordinatesRequest{
    BaseCommand: models.BaseCommand{ID: storeID},
    Latitude:    &lat,
    Longitude:   &lon,
}, nil)

// Закрыть магазин
err := c.Store.Close(ctx, &models.CloseStoreRequest{
    BaseCommand: models.BaseCommand{ID: storeID},
}, nil)

// Открыть магазин
err := c.Store.Open(ctx, &models.OpenStoreRequest{
    BaseCommand: models.BaseCommand{ID: storeID},
}, nil)
```

### SponsoredCards (Спонсорские карты)

```go
// Создать спонсорскую карту
err := c.SponsoredCards.Create(ctx, &models.CreateSponsoredCardRequest{
    BaseCommand: models.BaseCommand{ID: uuid.New()},
    SeriesID:    seriesID,
}, nil)

// Выпустить карту
err := c.SponsoredCards.Issue(ctx, &models.IssueSponsoredCardRequest{
    BaseCommand: models.BaseCommand{ID: cardID},
    OwnerID:     ownerID,
}, nil)

// Начислить на карту
err := c.SponsoredCards.Accrual(ctx, &models.SponsoredCardAccrualRequest{
    BaseCommand: models.BaseCommand{ID: cardID},
    CurrencyID:  currencyID,
    Amount:      500.0,
}, nil)

// Списать с карты
err := c.SponsoredCards.Subtract(ctx, &models.SponsoredCardSubtractRequest{
    BaseCommand: models.BaseCommand{ID: cardID},
    CurrencyID:  currencyID,
    Amount:      100.0,
}, nil)

// Заблокировать карту
err := c.SponsoredCards.Block(ctx, &models.BlockSponsoredCardRequest{
    BaseCommand:   models.BaseCommand{ID: cardID},
    BlockReasonID: &reasonID,
}, nil)

// Получить транзакции карты
transactions, err := c.SponsoredCards.GetTransactions(ctx, &models.GetByIdRequest{ID: cardID}, nil)

// Получить чеки карты
cheques, err := c.SponsoredCards.GetCheques(ctx, &models.GetByIdRequest{ID: cardID}, nil)
```

### StaticSegment (Статические сегменты)

```go
// Создать статический сегмент
err := c.StaticSegment.Create(ctx, &models.CreateStaticSegmentRequest{
    BaseCommand: models.BaseCommand{ID: uuid.New()},
    Name:        "VIP клиенты",
}, nil)

// Установить группу сегмента
err := c.StaticSegment.SetGroup(ctx, &models.SetStaticSegmentGroupRequest{
    BaseCommand: models.BaseCommand{ID: segmentID},
    GroupID:     &groupID,
}, nil)

// Добавить клиента в сегмент (через CustomerSegment)
err := c.CustomerSegment.AddToStaticSegment(ctx, &models.AddToStaticSegmentRequest{
    BaseCommand:     models.BaseCommand{ID: customerID},
    StaticSegmentID: segmentID,
}, nil)

// Удалить клиента из сегмента
err := c.CustomerSegment.RemoveFromStaticSegment(ctx, &models.RemoveFromStaticSegmentRequest{
    BaseCommand:     models.BaseCommand{ID: customerID},
    StaticSegmentID: segmentID,
}, nil)
```

### TerritorialDivision (Территориальные подразделения)

```go
// Создать территориальное подразделение
err := c.TerritorialDivision.Create(ctx, &models.CreateTerritorialDivisionRequest{
    BaseCommand: models.BaseCommand{ID: uuid.New()},
    Name:        "Москва",
    ParentID:    nil, // или &parentID для вложенных
}, nil)

// Установить родителя
err := c.TerritorialDivision.SetParent(ctx, &models.SetParentRequest{
    BaseCommand: models.BaseCommand{ID: divisionID},
    ParentID:    &parentID,
}, nil)
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

### Batch-операции

Для выполнения нескольких операций за один запрос можно использовать batch-методы:

```go
// Batch для валют
err := c.Currencies.Batch(ctx, &models.BatchRequest{
    Commands: []interface{}{
        map[string]interface{}{"type": "create", "data": {...}},
        map[string]interface{}{"type": "activate", "data": {...}},
    },
}, nil)

// Batch для стран
err := c.Countries.Batch(ctx, &models.BatchRequest{
    Commands: []interface{}{...},
}, nil)

// Batch для счетов
err := c.Accounts.Batch(ctx, &models.BatchRequest{
    Commands: []interface{}{...},
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

## Логирование

Клиент автоматически логирует все запросы с использованием стандартного логгера Go (`log`). В логах фиксируется:

- **[REQUEST]** - отправка запроса (метод, endpoint, тип объекта, ID)
- **[RESPONSE]** - получение ответа (HTTP статус, тип объекта, ID, время выполнения)
- **[SUCCESS]** - успешное выполнение операции
- **[ERROR]** - ошибки (статус, детали ошибки)

Пример вывода:
```
2024/01/15 12:00:00 [REQUEST] Method: POST, Endpoint: /api/v2/stores/create, ObjectType: CreateStoreRequest, ObjectID: 550e8400-e29b-41d4-a716-446655440000
2024/01/15 12:00:00 [RESPONSE] Status: 200, ObjectType: CreateStoreRequest, ObjectID: 550e8400-e29b-41d4-a716-446655440000, Duration: 150ms
2024/01/15 12:00:00 [SUCCESS] Command executed: ObjectType: CreateStoreRequest, ObjectID: 550e8400-e29b-41d4-a716-446655440000
```

При ошибках:
```
2024/01/15 12:00:00 [ERROR] API command failed: Status: 404, ObjectType: GetByIdRequest, ObjectID: 550e8400-e29b-41d4-a716-446655440000, Error: API command failed with status 404, Details: {"error": "not found"}
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

    "github.com/rsl6/rsloyalty/client"
    "github.com/rsl6/rsloyalty/mock"
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

// Статусы подарочных карт
models.GiftCardStatusCreated       // "Created"
models.GiftCardStatusActive        // "Active"
models.GiftCardStatusBlocked       // "Blocked"
models.GiftCardStatusExpired       // "Expired"
models.GiftCardStatusFullyRedeemed // "FullyRedeemed"

// Статусы спонсорских карт
models.SponsoredCardStatusCreated       // "Created"
models.SponsoredCardStatusReadyToIssued // "ReadyToIssued"
models.SponsoredCardStatusIssued        // "Issued"
models.SponsoredCardStatusActive        // "Active"
models.SponsoredCardStatusBlocked       // "Blocked"
models.SponsoredCardStatusExpired       // "Expired"

// Типы транзакций
models.TransactionAccrual  // "Accrual"
models.TransactionSubtract // "Subtract"
models.TransactionBurning  // "Burning"

// Типы свойств
models.PropertyTypeInteger      // "Integer"
models.PropertyTypeString       // "String"
models.PropertyTypeDate         // "Date"
models.PropertyTypeEnum         // "Enum"
models.PropertyTypeBoolean      // "Boolean"
models.PropertyTypeEnumMultiple // "EnumMultiple"
```

## Лицензия

MIT License
