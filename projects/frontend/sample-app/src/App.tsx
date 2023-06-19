import './App.css';
import { useState, useEffect } from "react";
import Home from './pages/Home';
import About from './pages/About'
import { Link, Route, Routes, BrowserRouter } from 'react-router-dom'
import * as THREE from "three";
import { Camera } from 'three';

function App() {
  const [isDesktop, setDesktop] = useState(window.innerWidth > 1200);

  const updateMedia = () => {
    setDesktop(window.innerWidth > 1200);
  };

  useEffect(() => {
    const scene = new THREE.Scene();
    const camera = new THREE.PerspectiveCamera(75, window.innerWidth / window.innerHeight, 0.1, 1000);

    const renderer = new THREE.WebGL1Renderer({
      canvas: document.querySelector('#bg'),
      alpha: true
    });

    renderer.setPixelRatio(window.devicePixelRatio);
    renderer.setSize(window.innerWidth / 2.5, window.innerHeight / 2.5);

    camera.position.setZ(30);
    renderer.render(scene, camera);

    const geometry = new THREE.TorusGeometry(15, 3, 10, 30);
    const material = new THREE.MeshBasicMaterial({
      color: 0x556345,
      wireframe: true
    });
    const torus = new THREE.Mesh(geometry, material);

    scene.add(torus)
    function animate() {
      requestAnimationFrame(animate);
      torus.rotation.x += 0.001;
      torus.rotation.y += 0.001;
      torus.rotation.z += 0.010;
      renderer.render(scene, camera)
    }
    animate()
  }, [])

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
        <div className='Three-canvas'>
          <canvas id='bg'></canvas>

        </div>
        <Routes>
          <Route path='/' element={<Home isDesktop={isDesktop} />} />
          <Route path='/about' element={<About />} />
        </Routes>
      </header >
    </div >
  );
};

export default App;
