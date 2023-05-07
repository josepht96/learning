import React from 'react';
import { useState, useEffect } from "react";
import logo from './logo.svg';
import './App.css';
import internal from 'stream';
import styled from "@emotion/styled";
import { config } from "react-spring";
import { Spring, animated } from "react-spring";


function App() {
  let data: Message[] = [
    {
      text: "Hello"
    },
    {
      text: "My name is Joe Thomas"
    },
    {
      text: "I work at Deloitte"
    },
    {
      text: "My email is joebthomas4@gmail.com"
    }
  ]

  interface Message {
    text: string;
  }

  const [active, setActive] = useState(0);

  const Dot = styled(animated.button)`
  outline: none;
  border: none;
  width: 15px;
  height: 15px;
  background: #fff;
  border-radius: 50%;
  margin: 0 16px;
  cursor: pointer;
`;

  return (
    <div className="App">
      <header className="App-header">
        <div className='Scroll-pane'>
          <div className='Info-block'>
            <p>{data[active].text}</p>
          </div>
          <div className='Button-bar'>
              {data.map((message: Message, i: number) => (
                <Spring
                  config={config.wobbly}
                  from={{ transform: `scale(1)` }}
                  to={{ transform: active === i ? `scale(1.25)` : `scale(1)` }}
                  key={message.text}
                >
                  {({ transform }: { transform: any }) => (
                    <Dot style={{ transform }} onClick={() => setActive(i)} />
                  )}
                </Spring>
              ))}
          </div>
        </div>
      </header>
    </div>
  );
};

export default App;
