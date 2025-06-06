import { useState, useEffect, useRef } from "react";
import '../App.css';
import { CSSTransition } from 'react-transition-group';
import { motion } from 'framer-motion'
import { data, Message } from '../Data'

function Home(props: any) {
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
        <div className="Page">
            <div className={props.isDesktop ? "Scroll-pane" : "Scroll-pane-mobile"}>
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
                    {/* <div className="Dot-small">
                <motion.div
                  animate={{x: [-50, 50, -50] }}
                  transition={{ repeat: Infinity, type: "tween", duration: 3 }}
                >
                </motion.div>
              </div> */}
                </div>
            </div>
        </div>
    );
};

export default Home;
