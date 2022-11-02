import testImage from './assets/coleus.jpg';
import './App.css';
import NavBar from './components/NavBar/NavBar';
import Home from './components/Home';
import About from './components/About';
import {
  BrowserRouter as Router, Routes,
  Route, Navigate,
} from "react-router-dom";

function App() {
  return (
    <div className="App">
      <NavBar />
        <Router>
          <Routes>
            <Route exact path="/" element={<Home />} />
            <Route path="/about" element={<About />} />
            <Route path="/*" element={<Home />} />
          </Routes>
        </Router>
      </div>
  );
}

export default App;
