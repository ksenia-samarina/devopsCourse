import { Link } from "react-router-dom";

export default function Navbar() {
    return (
        <nav className="bg-blue-600 text-white p-4 flex gap-4">
            <Link to="/" className="hover:underline">Dashboard</Link>
            <Link to="/create" className="hover:underline">Create Order</Link>
            <Link to="/order" className="hover:underline">Order Info</Link>
            <Link to="/stocks" className="hover:underline">Stocks</Link>
        </nav>
    );
}
