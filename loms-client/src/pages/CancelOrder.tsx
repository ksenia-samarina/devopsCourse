import { useState } from "react";
import { cancelOrder } from "../api/loms";

export default function CancelOrder() {
    const [orderID, setOrderID] = useState<number | null>(null);
    const [status, setStatus] = useState<string>("");

    const handleCancel = async () => {
        if (orderID === null) return alert("Введите ID заказа");
        try {
            // Передаем объект с полем orderID
            await cancelOrder({ orderID });
            setStatus(`Order ${orderID} cancelled successfully`);
        } catch (err) {
            console.error(err);
            setStatus("Ошибка при отмене заказа");
        }
    };

    return (
        <div className="p-8">
            <h1 className="text-2xl font-bold mb-4">Cancel Order</h1>
            <input
                type="number"
                placeholder="Order ID"
                value={orderID ?? ""}
                onChange={e => setOrderID(Number(e.target.value))}
                className="border p-2 mb-2"
            />
            <button onClick={handleCancel} className="bg-red-600 text-white p-2">
                Cancel Order
            </button>
            {status && <p className="mt-4">{status}</p>}
        </div>
    );
}
