import testImage from '../assets/coleus.jpg';
import '../App.css';
import './Home.css';

function Home() {
  return (
    <div className="Home-box">
      <div className="SideBar">
        <div className="SideBar-item">
          <a href="/">Coleus</a>
        </div>
        <div className="SideBar-item">
          <a href="/">Coral Bells</a>
        </div>
      </div>
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
      <div className="SideBar-right">
        <div className="SideBar-hidden-item">
        </div>
      </div>
    </div>
  );
}

export default Home;
