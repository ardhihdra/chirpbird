import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './core/App/App';
import reportWebVitals from './reportWebVitals';
import * as dotenv from 'dotenv'
import * as dotenvExpand from 'dotenv-expand'

global.MASTER_URL = `http://${process.env.REACT_APP_MASTER_URL}`

dotenvExpand(dotenv.config())

// to do
// axios.interceptors.request.use(function (config) {
//   const token = sessionStorage.getItem('token')
//   config.headers.Authorization = 'Bearer ' + token;

//   return config;
// });

ReactDOM.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
  document.getElementById('root')
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
