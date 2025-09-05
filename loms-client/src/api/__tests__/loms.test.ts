import { createOrder, listOrder, cancelOrder, cntStocks } from '../loms';

beforeEach(() => {
    global.fetch = jest.fn();
});

afterEach(() => {
    jest.resetAllMocks();
});

test('createOrder returns orderID', async () => {
    (fetch as jest.Mock).mockResolvedValueOnce({
        ok: true,
        json: async () => ({ orderID: 123 }),
    });

    const res = await createOrder({ user: 1, items: [{ sku: 1, count: 2 }] });
    expect(res.orderID).toBe(123);
});

test('listOrder returns order info', async () => {
    (fetch as jest.Mock).mockResolvedValueOnce({
        ok: true,
        json: async () => ({ user: 1, status: 1, items: [{ sku: 1, count: 2 }] }),
    });

    const res = await listOrder(123);
    expect(res.user).toBe(1);
    expect(res.items.length).toBe(1);
});

test('cancelOrder does not throw on success', async () => {
    (fetch as jest.Mock).mockResolvedValueOnce({
        ok: true,
    });

    await expect(cancelOrder({ orderID: 123 })).resolves.not.toThrow();
});

test('cancelOrder throws on network error', async () => {
    (fetch as jest.Mock).mockResolvedValueOnce({
        ok: false,
    });

    await expect(cancelOrder({ orderID: 123 })).rejects.toThrow(
        "Ошибка сети при отмене заказа"
    );
});

test('cntStocks returns stock info', async () => {
    (fetch as jest.Mock).mockResolvedValueOnce({
        ok: true,
        json: async () => ({ stocks: [{ warehouseID: 1, count: 10 }] }),
    });

    const res = await cntStocks(1);
    expect(res.stocks[0].count).toBe(10);
});
