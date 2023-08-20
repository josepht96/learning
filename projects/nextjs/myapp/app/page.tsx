'use client';

import { useState, useEffect, useRef } from "react";
import './app.css';
import { CSSTransition } from 'react-transition-group';
import { motion } from 'framer-motion';

interface Message {
  text: string;
}

const data: Message[] = [
  {
      text: "Hello"
  },
  {
      text: 'My name is Joe Thomas'
  },
  {
      text: "I am currently an engineer at Deloitte, in the federal government space. I work on devOps, SRE, and infrastructure related issues"
  },
  {
      text: "I am interested primarily in SRE and full-stack development work"
  }
]

export default function HomePage(props: any) {
  const [active, setActive] = useState(0);
  const [tick, setTick] = useState(true);
  const [rotate, setRotate] = useState(true)
  const nodeRef = useRef(null);

  function mouserOver(i: number) {
    if (i !== active) {
      setTick(!tick)
    }
    setActive(i)
  }
  return (
      <div className="Scroll-pane">
        <CSSTransition
          in={tick}
          appear={true}
          timeout={2000}
          classNames="fadein"
          unmountOnExit={false}
          nodeRef={nodeRef}
        >
          <div className='Info-block' ref={nodeRef}>
            <p>
              {data[active].text}
            </p>
          </div>
        </CSSTransition>
        <div className='Button-bar'>
          <div className='Dot-bar'>
            {data.map((message: Message, i: number) => (
              <div className='Dot-box' key={i}>
                <div className={i == active ? 'Dot-selected' : 'Dot'} onMouseOver={() => mouserOver(i)}>
                  {/* <motion.div
                    animate={{ scale: i == active ? 1.2 : 1 }}
                    transition={{ type: "spring " }}
                    onMouseOver={() => {
                      setRotate(!rotate);
                    }}
                  >
                  </motion.div> */}
                </div>
              </div>
            ))}
          </div>
        </div>
      </div>
  );
}