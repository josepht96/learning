import './NavBar.css';

function NavBar() {
    return (
        <div className="NavBar">
            <div className="NavBar-body">
                {/* <div className='NavBar-title-box'>
                    <div className='NavBar-title'>
                        <a href="/">Plants</a>
                    </div>
                </div> */}
                <div className='NavBar-links-box'>
                    <div className='NavBar-link'>
                        <a href="/">Home</a>
                    </div>
                    <div className='NavBar-link'>
                        <a href="/about">About</a>
                    </div>
                </div>
            </div>
        </div >
    );
}

export default NavBar;
