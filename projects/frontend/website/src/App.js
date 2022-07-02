import { Outlet, Link } from "react-router-dom";
import {
  BrowserRouter,
  Routes,
  Route,
} from "react-router-dom";
import About from "./pages/About";
import Home from "./pages/Home";
import './App.css';

function App() {
  return (
    <div className="App">
      <BrowserRouter>
        <nav class="Navbar">
          <Link class="Link" to="/">Home</Link>
          {/* <Link class="Link" to="/about">About</Link> */}
        </nav>
        <Routes>
          <Route path="/" element={<Home />} />
          {/* <Route path="about" element={<About />} /> */}
        </Routes>
        <Outlet />
      </BrowserRouter>
    </div>
  );
}

export default App;
