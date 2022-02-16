// import FutureFeature from '../components/FutureFeature';
import ContactInfo from '../components/ContactInfo';
import logo from '../assets/img/logo/logo-horizontal.png';

export default function About() {
    return (
        <div className="pages-container">
            <img src={logo} className="App-logo" alt="logo" />
            <div className="txt-header pb-5">Tentang Kami</div>
            <div className="txt-subheader"></div>
            <div className="rounded form-container p-3">
                Kami adalah Kamu
                <ContactInfo />
            </div>
        </div>   
    )
}