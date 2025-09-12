const BASE_URL = "http://loms:8082/v1";

export interface Item {
    sku: number;
    count: number;
}

export interface CreateOrderIn {
    user: number;
    items: Item[];
}

export interface CreateOrderOut {
    orderID: number;
}

export interface ListOrderOut {
    user: number;
    status: number;
    items: Item[];
}

export interface CancelOrderIn {
    orderID: number;
}

export interface StocksOut {
    stocks: { warehouseID: number; count: number }[];
}

async function request<T>(url: string, options?: RequestInit): Promise<T> {
    const res = await fetch(url, {
        headers: { "Content-Type": "application/json" },
        ...options,
    });
    if (!res.ok) throw new Error("Ошибка сети");
    return res.json();
}

export function createOrder(data: CreateOrderIn) {
    return request<CreateOrderOut>(`${BASE_URL}/orders`, {
        method: "POST",
        body: JSON.stringify(data),
    });
}

export function listOrder(orderID: number) {
    return request<ListOrderOut>(`${BASE_URL}/orders/${orderID}`);
}

export async function cancelOrder({ orderID }: CancelOrderIn): Promise<void> {
    const res = await fetch(`${BASE_URL}/orders/${orderID}`, {
        method: "DELETE",
        headers: { "Content-Type": "application/json" },
    });
    if (!res.ok) throw new Error("Ошибка сети при отмене заказа");
}

export function cntStocks(sku: number) {
    return request<StocksOut>(`${BASE_URL}/stocks/${sku}`);
}
