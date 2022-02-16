import LoginForm from './LoginForm.jsx';
import './Login.css';

export default function Login() {
    return (
        <div className="ds-flow">
            {/* <div className="illustration">illustration</div> */}
            <div className="main-section pages-container">
                <div id="title" className="font-bold text-6xl pt-6 mt-6"><a href="/">Chirpbird</a></div>
                <div id="subtitle" className="text-lg mt-2 pb-6 mb-6">Make some friends from around the world</div>
                
                <div className="rounded-xl login-form-container border border-current shadow-2xl py-3 px-7">
                    <LoginForm/>
                </div>
            </div>     
        </div>
    )
}