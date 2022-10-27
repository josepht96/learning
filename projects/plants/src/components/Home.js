import testImage from '../assets/coleus.jpg';
import '../App.css';
import './Home.css';

function Home() {
  return (
    <div className="Home">
      <div className="SideBar">
        <div className="SideBar-item">
          <a href="/">Coleus</a>
        </div>
        <div className="SideBar-item">
          <a href="/">Coral Bells</a>
        </div>
        <div className="SideBar-item">
          <a href="/">Coleus</a>
        </div>
        <div className="SideBar-item">
          <a href="/">Coral Bells</a>
        </div>
        <div className="SideBar-item">
          <a href="/">Coleus</a>
        </div>
        <div className="SideBar-item">
          <a href="/">Coral Bells</a>
        </div>
        <div className="SideBar-item">
          <a href="/">Coleus</a>
        </div>
        <div className="SideBar-item">
          <a href="/">Coral Bells</a>
        </div>
        <div className="SideBar-item">
          <a href="/">Coleus</a>
        </div>
        <div className="SideBar-item">
          <a href="/">Coral Bells</a>
        </div>
        <div className="SideBar-item">
          <a href="/">Coleus</a>
        </div>
        <div className="SideBar-item">
          <a href="/">Coral Bells</a>
        </div>
        <div className="SideBar-item">
          <a href="/">Coleus</a>
        </div>
        <div className="SideBar-item">
          <a href="/">Coral Bells</a>
        </div>
        <div className="SideBar-item">
          <a href="/">Coleus</a>
        </div>
        <div className="SideBar-item">
          <a href="/">Coral Bells</a>
        </div>
        <div className="SideBar-item">
          <a href="/">Coleus</a>
        </div>
        <div className="SideBar-item">
          <a href="/">Coral Bells</a>
        </div>
        <div className="SideBar-item">
          <a href="/">Coleus</a>
        </div>
        <div className="SideBar-item">
          <a href="/">Coral Bells</a>
        </div>
        <div className="SideBar-item">
          <a href="/">Coleus</a>
        </div>
        <div className="SideBar-item">
          <a href="/">Coral Bells</a>
        </div>
        <div className="SideBar-item">
          <a href="/">Coleus</a>
        </div>
        <div className="SideBar-item">
          <a href="/">Coral Bells</a>
        </div>
      </div>
      <div className='PlantPanel'>
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

    </div>
  );
}

export default Home;
