import { useState } from "react";
import './App.css';
import { CSSTransition } from 'react-transition-group';

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
  const [tick, setTick] = useState(true)

  function mouserOver(i: number) {
    if (i !== active) {
      setTick(!tick)
    }
    setActive(i)
  }

  return (
    <div className="App">
      <header className="App-header">
        <div>
          <CSSTransition
            in={tick}
            appear={true}
            timeout={2000}
            classNames="fadein"
            unmountOnExit={false}
          >
            <div id="text-body" className='Info-block' >
              {data[active].text}
            </div>
          </CSSTransition>
          <div className='Button-bar'>
            <div className='Dot-bar'>
              {data.map((message: Message, i: number) => (
                <div className='Dot-box' key={i}>
                  <div className={i == active ? 'Dot-selected' : 'Dot'} onMouseOver={() => mouserOver(i)}></div>
                </div>
              ))}
            </div>
          </div>
        </div>
      </header>
    </div>
  );
};

export default App;
