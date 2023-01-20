import testImage from '../assets/coleus.jpg';
import '../App.css';
import './Home.css';

function Home() {
  return (
    <div className='PlantPanel'>
      <p>
        Coleus
      </p>
      <div className="PlantBox">
        <img src={testImage} className="PlantImage" alt="testImage" />
        <div className="PlantText">
          Whole bunch of text about stuff
          Whole bunch of text about stuff
          Whole bunch of text about stuff
          Whole bunch of text about stuff
          Whole bunch of text about stuff
        </div>
      </div>
    </div>
  );
}

export default Home;
