import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Navbar from "./components/Navbar";
import Dashboard from "./pages/Dashboard";
import CreateOrder from "./pages/CreateOrder";
import Order from "./pages/Order";
import Stocks from "./pages/Stocks";
import CancelOrder from "./pages/CancelOrder";

function App() {
    return (
        <Router>
            <Navbar />
            <Routes>
                <Route path="/" element={<Dashboard />} />
                <Route path="/create" element={<CreateOrder />} />
                <Route path="/order" element={<Order />} />
                <Route path="/stocks" element={<Stocks />} />
                <Route path="/cancel" element={<CancelOrder />} />
            </Routes>
        </Router>
    );
}

export default App;
