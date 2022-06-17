import { Outlet, Link } from "react-router-dom";
import {
  BrowserRouter,
  Routes,
  Route,
} from "react-router-dom";
import Expenses from "./routes/expenses";
import Invoices from "./routes/invoices";
import Home from "./routes/home";
import './App.css';

function App() {
  return (
    <div className="App">
      <BrowserRouter>
        <nav
          style={{
            borderBottom: "solid 1px",
            paddingBottom: "1rem",
          }}>
          <Link to="/">Home</Link> |{" "}
          <Link to="/invoices">Invoices</Link> |{" "}
          <Link to="/expenses">Expenses</Link>
        </nav>
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="expenses" element={<Expenses />} />
          <Route path="invoices" element={<Invoices />} />
        </Routes>
        <Outlet />
      </BrowserRouter>
    </div>
  );
}

export default App;
