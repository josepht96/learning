import testImage from './assets/coleus.jpg';
import './App.css';
import NavBar from './components/NavBar/NavBar';

function App() {
  return (
    <div className="App">
      <NavBar />
      <header className="App-header">
        <img src={testImage} className="App-image" alt="testImage" />
        <p>
          Edit and save to reload.
        </p>
      </header>
      <header className="App-header">
        <img src={testImage} className="App-image" alt="logo" />
        <p>
          Edit <code>src/App.js</code> and save to reload.
        </p>
      </header>
    </div>
  );
}

export default App;
