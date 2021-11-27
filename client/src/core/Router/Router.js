import React from "react";
import {
    BrowserRouter as Router,
    Routes,
    Route,
    Link
} from "react-router-dom";
import './Router.css';
import routes from './routes';

let listRoutes = [];
for(let rt in routes) {
    Object.assign(listRoutes, routes[rt]);
}

export default function AppRouter() {
    let prevScrollpos = window.pageYOffset;
    window.onscroll = function() {
        let currentScrollPos = window.pageYOffset;
        const container = document.getElementById("router-container")
        if(container) {
            if(prevScrollpos > currentScrollPos) {
                container.style.bottom = "-5vh" // scroll to down
            } else {
                container.style.bottom = prevScrollpos === 0 ? "-5vh": ".5vh" // scroll to top 
            }
            prevScrollpos = currentScrollPos;
        }
    }
    return (
        <Router>
            <div className="ds-flow">
                <Routes>
                    {   
                        listRoutes.map((route, index) => (
                                <Route
                                    key={index}
                                    path={route.path}
                                    exact={route.exact}
                                    element={<route.main />}
                                />
                        ))
                    }
                </Routes>

                
                <div id="router-container" className="router-container fade-in">
                    <div className="router-list">
                        {routes.menus.map((route, index) => {
                            return (
                                <Link 
                                    className="router-box"
                                    key={"routes.menus" + index}
                                    to={route.path}
                                    exact={route.exact ? route.exact.toString(): "false"} title={route.title}>
                                        <img src={route.icon} alt="icon" className="main-icon" />
                                </Link>
                            )
                        })}
                    </div>
                </div>

                
            </div>
        </Router>
    )
}