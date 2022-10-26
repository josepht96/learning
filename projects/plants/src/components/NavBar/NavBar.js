import './NavBar.css';

function NavBar() {
    return (
        <div className="NavBar">
            <div className="NavBar-body">
            <div className='NavBar-title-box'>
                    <div className='NavBar-title'>
                        Plants and stuff
                    </div>
                </div>
                <div className='NavBar-links-box'>
                    <div className='NavBar-link'>
                        Home
                    </div>
                    <div className='NavBar-link'>
                        About
                    </div>
                </div>
            </div>
        </div>
    );
}

export default NavBar;
