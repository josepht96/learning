import React from 'react';
import { useState, useEffect } from "react";
import logo from './logo.svg';
import './App.css';
import internal from 'stream';

function App() {
  const [count, setCount] = useState({counter: 3, countNum: 0});
  
  useEffect(() => {
    const interval = setInterval(() => {
      setCount((prevCount) => ({
        counter: prevCount.counter - 1,
        countNum: prevCount.countNum
      }));
    }, 1000);

    return () => clearInterval(interval);
  }, []);

  useEffect(() => {
    if (count.counter === 0) {
      setCount(() => ({
        counter: 3,
        countNum: count.countNum + 1
      }));
    }
  }, [count]);

  const sayHello = (): string => {
    return "Hello, from some function";
  };

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <div>{count.counter}</div>
        <div>{count.countNum}</div>
        <div>{sayHello()}</div>
        <p>
          Edit <code>src/App.tsx</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      </header>
    </div>
  );
};

export default App;
