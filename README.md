# Logistics and Order Management System

Сервис учета заказов и логистики

## 1. createOrder

- Создает новый заказ пользователя
- Резервирует товары на складе

Request:
```
{
  user int64,
  items []{
    sku uint32, 
    cnt uint16,
  },
}
```
Response:
```
{
  orderID int64,
}
```

## 2. listOrder

- Показывает информацию по заказу

Request:
```
{
  orderID int64,
}
```
Response:
```
{
    status string, // (new | awaiting_payment | failed | payed | cancelled)
    user int64,
    items []{
        sku  uint32,
        count uint16,
    },
}
```

## 3. cancelOrder

- Отменяет заказ
- Снимает резерв со всех товаров в заказе

Request:
```
{
  orderID int64,
}
```
Response:
```
{}
```

## 4. cntStocks

- Возвращает кол-во товаров, которые можно купить с разных складов
- Гарантирует, что товар, который зарезервирован у кого-то в заказе и ждет оплаты, купить нельзя

Request:
```
{
  sku uint32,
}
```
Response:
```
{
    stocks []{
        warehouseID int64,
        cnt uint64,
    },
}
```