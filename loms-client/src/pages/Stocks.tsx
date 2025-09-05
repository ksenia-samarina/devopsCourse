import { useState } from "react";
import { cntStocks } from "../api/loms";

export default function Stocks() {
    const [sku, setSku] = useState(0);
    const [stocks, setStocks] = useState<any[]>([]);

    const handleFetch = async () => {
        try {
            const res = await cntStocks(sku);
            setStocks(res.stocks);
        } catch {
            alert("Ошибка при получении остатков");
        }
    };

    return (
        <div className="p-8">
            <h1 className="text-2xl font-bold mb-4">Stocks</h1>
            <input
                type="number"
                placeholder="SKU"
                value={sku}
                onChange={e => setSku(Number(e.target.value))}
                className="border p-2 mb-2"
            />
            <button onClick={handleFetch} className="bg-blue-600 text-white p-2 mb-2">
                Check Stocks
            </button>

            <ul className="list-disc pl-5 mt-4">
                {stocks.map((s, idx) => (
                    <li key={idx}>
                        Warehouse {s.warehouseID}: {s.count}
                    </li>
                ))}
            </ul>
        </div>
    );
}
