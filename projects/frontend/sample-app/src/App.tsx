import { useState, useEffect } from "react";
import './App.css';
import { CSSTransition } from 'react-transition-group';
import { motion } from 'framer-motion'

function App() {
  let data: Message[] = [
    {
      text: "Hello"
    },
    {
      text: "My name is Joe Thomas"
    },
    {
      text: "I work at Deloitte. This is a bunch of text that might run off the screen"
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
  const [rotate, setRotate] = useState(true)

  return (
    <div className="App">
      <header className="App-header">
        <div className="Scroll-pane">
          <CSSTransition
            in={tick}
            appear={true}
            timeout={2000}
            classNames="fadein"
            unmountOnExit={false}
          >
            <div className='Info-block' >
              {data[active].text}
              {/* <div className="Dot-small">
                <motion.div
                  animate={{x: [-50, 50, -50] }}
                  transition={{ repeat: Infinity, type: "tween", duration: 3 }}
                >
                </motion.div>
              </div> */}
            </div>
          </CSSTransition>
          <div className='Button-bar'>
            <div className='Dot-bar'>
              {data.map((message: Message, i: number) => (
                <div className='Dot-box' key={i}>
                  <div className={i == active ? 'Dot-selected' : 'Dot'} onMouseOver={() => mouserOver(i)}>
                    <motion.div
                      animate={{ scale: i == active ? 1.2 : 1 }}
                      transition={{ type: "spring " }}
                      onMouseOver={() => {
                        setRotate(!rotate);
                      }}
                    >
                    </motion.div>
                  </div>
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
