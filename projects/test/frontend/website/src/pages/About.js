import logo from '../logo.svg';
import './css/Home.css';

function About() {
  return (
    <div className="Page">
      <header className="Page-header">
        <img src={logo} className="Page-logo" alt="logo" />
        <p>
          About
        </p>
        <a
          className="Page-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      </header>
    </div>
  );
}

export default About;