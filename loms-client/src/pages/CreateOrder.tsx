import { useState } from "react";
import { createOrder, Item } from "../api/loms";

export default function CreateOrder() {
    const [user, setUser] = useState(0);
    const [items, setItems] = useState<Item[]>([{ sku: 0, count: 1 }]);
    const [orderID, setOrderID] = useState<number | null>(null);

    const handleAddItem = () => setItems([...items, { sku: 0, count: 1 }]);

    const handleSubmit = async () => {
        try {
            const res = await createOrder({ user, items });
            setOrderID(res.orderID);
        } catch (err) {
            alert("Ошибка при создании заказа");
        }
    };

    return (
        <div className="p-8">
            <h1 className="text-2xl font-bold mb-4">Create Order</h1>
            <input
                type="number"
                placeholder="User ID"
                value={user}
                onChange={e => setUser(Number(e.target.value))}
                className="border p-2 mb-2"
            />
            {items.map((item, idx) => (
                <div key={idx} className="flex gap-2 mb-2">
                    <input
                        type="number"
                        placeholder="SKU"
                        value={item.sku}
                        onChange={e => {
                            const newItems = [...items];
                            newItems[idx].sku = Number(e.target.value);
                            setItems(newItems);
                        }}
                        className="border p-2"
                    />
                    <input
                        type="number"
                        placeholder="Count"
                        value={item.count}
                        onChange={e => {
                            const newItems = [...items];
                            newItems[idx].count = Number(e.target.value);
                            setItems(newItems);
                        }}
                        className="border p-2"
                    />
                </div>
            ))}
            <button
                onClick={handleAddItem}
                className="bg-gray-200 p-2 mb-2"
            >
                Add Item
            </button>
            <button
                onClick={handleSubmit}
                className="bg-blue-600 text-white p-2"
            >
                Create Order
            </button>
            {orderID && <p className="mt-4">Order created with ID: {orderID}</p>}
        </div>
    );
}
