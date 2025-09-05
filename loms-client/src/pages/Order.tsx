import { useState } from "react";
import { listOrder, cancelOrder } from "../api/loms";

export default function Order() {
    const [orderID, setOrderID] = useState(0);
    const [order, setOrder] = useState<any>(null);

    const handleFetch = async () => {
        try {
            const res = await listOrder(orderID);
            setOrder(res);
        } catch {
            alert("Не удалось загрузить заказ");
        }
    };

    const handleCancel = async () => {
        try {
            await cancelOrder({orderID});
            alert("Заказ отменён");
            setOrder(null);
        } catch {
            alert("Ошибка при отмене");
        }
    };

    return (
        <div className="p-8">
            <h1 className="text-2xl font-bold mb-4">Order Info</h1>
            <input
                type="number"
                placeholder="Order ID"
                value={orderID}
                onChange={e => setOrderID(Number(e.target.value))}
                className="border p-2 mb-2"
            />
            <button onClick={handleFetch} className="bg-blue-600 text-white p-2 mb-2">
                Fetch Order
            </button>
            <button onClick={handleCancel} className="bg-red-600 text-white p-2 ml-2">
                Cancel Order
            </button>

            {order && (
                <div className="mt-4">
                    <p>User ID: {order.user}</p>
                    <p>Status: {order.status}</p>
                    <ul className="list-disc pl-5">
                        {order.items.map((i: any, idx: number) => (
                            <li key={idx}>
                                SKU: {i.sku}, Count: {i.count}
                            </li>
                        ))}
                    </ul>
                </div>
            )}
        </div>
    );
}
