import LoginForm from './LoginForm.jsx';
import './Home.css';

export default function Home() {
    return (
        <div className="pages-container">
            <div id="title" className="search-title ds-pt-6"><a href="/">Chirpbird</a></div>
            <div id="subtitle" className="search-subtitle">Chat and make some friends from around the world</div>
           
            <div className="ds-border form-container ds-p-3">
                <LoginForm/>
            </div>
        </div>
    )
}