// import FutureFeature from '../components/FutureFeature';
import ContactInfo from '../components/ContactInfo';
import logo from '../assets/img/logo/logo-horizontal.png';

export default function About() {
    return (
        <div className="pages-container">
            <img src={logo} className="App-logo" alt="logo" />
            <div className="txt-header ds-pb-5">Tentang Kami</div>
            <div className="txt-subheader"></div>
            <div className="ds-border form-container ds-p-3">
                Kami adalah Kamu
                <ContactInfo />
            </div>
        </div>   
    )
}