import testImage from './assets/coleus.jpg';
import './App.css';
import Home from './components/Home';
import About from './components/About';
import Contact from './components/Contact';
import {
  BrowserRouter as Router, Routes,
  Route, Navigate,
} from "react-router-dom";

function App() {
  return (
    <div className="App">
      <div className="navbar">
        <div className="title-box">
            <a href="/">Joe</a>
        </div>
        <div className='link-box'>
          <div className='link'>
            <a href="/">Home</a>
          </div>
          <div className='divider'>|</div>
          <div className='link'>
            <a href="/about">About</a>
          </div>
          <div className='divider'>|</div>
          <div className='link'>
            <a href="/contact">Contact</a>
          </div>
        </div>
        <div className="title-box">
        </div>
      </div >
      <div className="main">
        <Router>
          <Routes>
            <Route exact path="/" element={<Home />} />
            <Route path="/about" element={<About />} />
            <Route path="/contact" element={<Contact />} />
            <Route path="/*" element={<Home />} />
          </Routes>
        </Router>
      </div>
    </div>
  );
}

export default App;
