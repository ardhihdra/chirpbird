import './App.css';

import AppRouter from '../Router/Router';

function App() {
  return (
    <div className="App">
      <header className="App-header">
      </header>
      <main > 
        <div className="App-main">
          <div className="main-section">
            <AppRouter ></AppRouter>
          </div>
        </div>
      </main>
      <footer>
        <div className="tribute">
            Icons made by 
            <a href="https://www.freepik.com" title="Freepik"> Freepik </a> 
            from 
            <a href="https://www.flaticon.com/" title="Flaticon"> www.flaticon.com</a>
          </div>

      </footer>
    </div>
  );
}

export default App;
