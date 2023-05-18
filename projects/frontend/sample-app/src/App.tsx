import './App.css';
import { useState, useEffect } from "react";
import Home from './pages/Home';
import About from './pages/About'
import { Link, Route, Routes, BrowserRouter } from 'react-router-dom'

function App() {
  const [isDesktop, setDesktop] = useState(window.innerWidth > 1200);

  const updateMedia = () => {
    setDesktop(window.innerWidth > 1200);
  };

  useEffect(() => {
    window.addEventListener("resize", updateMedia);
    return () => window.removeEventListener("resize", updateMedia);
  });
  return (
    <div className="App">
      <header className="App-header">
      <div className={isDesktop ? "Nav-bar" : "Nav-bar-mobile"}>
          <div className="Nav-bar-box">
            <Link to="/" className="Link">Home</Link>
            <Link to="/about" className="Link">About</Link>
          </div>
        </div>
        <Routes>
          <Route path='/' element={<Home isDesktop={isDesktop}/>} />
          <Route path='/about' element={<About />} />
        </Routes>
      </header >
    </div >
  );
};

export default App;
